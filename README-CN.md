English | [简体中文](README-CN.md)

# Meoo Open API Go SDK

Meoo 开放平台的官方 Go SDK。纯 Go 标准库实现、无第三方依赖；提供同步 `client.Client`，覆盖 Bearer/API Key 认证、统一错误分层、有限重试、项目分页以及 Agent Run/SSE 事件流。所有网络方法都接收 `context.Context`，超时/取消由调用方的 ctx 控制。

## 要求

- Go 1.22 及以上。

## 安装

```sh
go get gitlab.alibaba-inc.com/oneday/meoo-sdk-go/client
```

> 内网私有 module：消费方需设置 `GOPRIVATE=gitlab.alibaba-inc.com`（让 `go get` 跳过公共 proxy/sumdb 直连 VCS），并配置好 Git 凭证；若 GitLab 走 http 或自签证书，视情况再配 `GOINSECURE`。

## 结构

module 根即仓库根，下辖两个包，按“是否由契约驱动”划分职责：

| 包 / 目录 | 来源 | 说明 |
|---|---|---|
| `generated/api_*.go`、`generated/model_*.go` | **OpenAPI Generator 自动生成**（禁止手改） | 契约驱动：operation 的 request-builder、参数、schema 对应的结构体 |
| `generated/{client,configuration,response,utils,transport,codec,params,errors}.go` | **手写静态基础设施** | 与契约无关的 HTTP 发起调用、序列化反序列化、参数编码与通用异常处理；提供生成 api/model 依赖的同包符号面 |
| `client/**` | **手写高层 Runtime** | 认证/超时/有限重试/SSE/分页/错误分层的正式运行时语义，以及公共客户端 facade |

`client` 包的业务模型直接复用 `generated` 的模型**类型**（并 re-export 为 `client.Project`、`client.AgentRun` 等别名），避免字段重复定义；但列表信封一律用局部结构 + 显式 `json` tag 解码，不引用生成模型的 Go 字段名，把“生成器字段命名变化”的编译期风险降到最低。

## 使用

门面代码在 `client/` 目录（`package client`）。因 `client` 与常见局部变量名冲突，示例按阿里云 SDK 惯例用别名 `meoo` 引入（`import meoo "…/meoo-sdk-go/client"`）；不起别名时，把客户端变量命名为 `client` 以外的名字即可。

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	meoo "gitlab.alibaba-inc.com/oneday/meoo-sdk-go/client"
)

func main() {
	c, err := meoo.NewClient(meoo.WithAPIKey(os.Getenv("MEOO_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	ctx := context.Background()

	// 惰性翻页（对齐 sql.Rows）：PageToken 由迭代器内部管理
	it := c.Projects().Iter(ctx, meoo.ListProjectsParams{}, nil)
	for it.Next() {
		fmt.Println(it.Item().UrlId)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

`Projects()` 提供 `Create`/`List`/`Iter`，`Agent()` 提供 `Start`/`Current`/`Cancel`/`Events`。

### 消费 SSE 事件流

`Agent().Events` 返回持有连接的 `*EventStream`，必须 `Close`（通常用 `defer`）。逐行惰性解析，收到契约声明的终态事件（`run.completed`、`run.failed`、`run.canceled`、`run.interrupted`、`run.superseded`）或流自然结束后 `Next` 返回 `io.EOF`，且不做重连：

```go
stream, err := c.Agent().Events(ctx, projectID, runID, nil)
if err != nil {
	return err
}
defer stream.Close()

for {
	event, err := stream.Next()
	if errors.Is(err, io.EOF) {
		break
	}
	if err != nil {
		return err
	}
	switch event.Event {
	case "message.delta":
		var delta generated.AgentMessageDeltaEventData
		if err := event.DataAs(&delta); err != nil {
			return err
		}
		// 处理 delta
	}
	// 契约要求忽略未知的 SSE 事件名：未知 event 原样透出且不报错
}
```

`AgentEvent.Data` 是所有 `data:` 行以 `\n` 拼接后的原始 JSON，用 `DataAs(&v)` 映射到生成模型或任意结构。

### 使用生成客户端覆盖全部 operation

facade 未封装的 operation 用 `Generated()` 的生成客户端（request-builder 模式）。它与本客户端共享 `baseURL` 与 `*http.Client`；凭证需先用 `Context(ctx)` 注入到 ctx：

```go
ctx, err := c.Context(context.Background())
if err != nil {
	return err
}
user, _, err := c.Generated().UserApi.GetUser(ctx).Execute()
if err != nil {
	return err
}
```

也可以直接用低层 `Transport()` 发原始契约路径的请求，认证/重试/错误语义与 facade 完全一致：

```go
raw, err := c.Transport().Request(ctx, "GET", "/open/v1/user", nil, nil)
```

### 错误分层

所有错误实现 `client.Error` 接口（`Message()`/`Unwrap()`），用 `errors.As` 判定具体类型，`errors.Is` 做链式判定：

- `*client.TransportError`：连接失败、超时、取消，以及请求体序列化 / 响应体解析失败。`Unwrap()` 保留底层 cause，支持 `errors.Is(err, context.Canceled)`。
- `*client.APIError`：服务端返回非 2xx。暴露 `Status`、`Code`、`TraceID`、原始 `Problem`（RFC 7807）与 `Headers`；`problem` 不可解析时按 HTTP status 兜底，`TraceID` 缺失时回退到响应头 `X-Meoo-Trace-Id`。
- 其余 SDK 级错误（如 SSE data 非法 JSON）仅通过 `Error` 接口对外暴露。

```go
var apiErr *client.APIError
if errors.As(err, &apiErr) && apiErr.Status == 404 {
	// 资源不存在
}
```

### 认证与重试

- **认证**：`Authorization: Bearer <credential>`，凭证每次请求动态取值。固定凭证用 `WithAPIKey`/`WithAccessToken`；动态凭证（OAuth refresh、成员 Token 重签发）用 `WithCredentialProvider` + `CredentialProviderFunc`，single-flight 与缓存由实现方负责。SSE 的凭证只走 Header，绝不进 query。
- **重试**：默认 `maxRetries=2`，仅对 429/502/503/504 且“可重试”的请求退避重试。GET/HEAD 或携带 `IdempotencyKey` 的写请求可重试；用 `RequestOptions.Retry`（配 `client.Bool(...)`）可逐请求覆盖。退避为指数退避（上限 8s），`Retry-After` 头优先。连接失败/超时直接归为 `*TransportError` 不重试。

## 构建与测试

```sh
go build ./...     # 编译 generated + client 两个包
go test ./...      # 运行 client 包的单元测试（httptest 打桩，无需网络）
go vet ./...
```

## 来源与生成

本仓是**发布产物**：`generated/api_*.go`、`model_*.go` 由 OpenAPI Generator（版本锁定）从契约生成，连同手写静态基础设施与 `client` Runtime，统一在源仓库 `meoo-open-sdk`（monorepo 的 `sdks/go/`）中生成、验证，再镜像到此处。**生成器脚本、契约副本、实网 E2E（覆盖全部 35 operation）与跨语言一致性校验都保留在源仓库**，本仓不含它们，以保持发布产物精简。

`generated/api_*.go`、`model_*.go` **禁止手改**；需要变更时请在源仓库改契约或生成约束后重新生成，再镜像过来。

## 发行说明

各版本的详细变更记录在 [ChangeLog.txt](./ChangeLog.txt)。

## 许可证

Apache-2.0（详见 [LICENSE](./LICENSE)）。

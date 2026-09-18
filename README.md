# Go SDK

module：`gitlab.alibaba-inc.com/oneday/meoo-sdk-go`。当前 MVP 已提供同步 `meoo.Client`，覆盖 Bearer/API Key、统一错误分层、有限重试、项目分页以及 Agent Run/SSE。仅依赖 Go 标准库（`go 1.22`，无第三方 `require`）：认证收敛为 Bearer，OAuth refresh 的 single-flight 与缓存交由调用方的 `CredentialProvider` 实现，因此不引入 `golang.org/x/oauth2`，无需拉取外部模块。

所有网络方法都接收 `context.Context`，超时/取消由调用方的 ctx 控制。

## 安装

```bash
go get gitlab.alibaba-inc.com/oneday/meoo-sdk-go/meoo
```

> 内网私有 module：消费方需设置 `GOPRIVATE=gitlab.alibaba-inc.com`（让 `go get` 跳过公共 proxy/sumdb 直连 VCS），并配置好 Git 凭证；若 GitLab 走 http 或自签证书，视情况再配 `GOINSECURE`。

## 两层结构

本仓是 Go SDK 的独立发布仓，module 根即仓库根，下辖两个包，按“是否由契约驱动”划分职责：

| 包 / 目录 | 来源 | 说明 |
|---|---|---|
| `generated/api_*.go`、`generated/model_*.go` | **OpenAPI Generator 自动生成**（禁止手改） | 真正由契约驱动的部分：operation 的 request-builder、参数、schema 对应的结构体 |
| `generated/{client,configuration,response,utils,transport,codec,params,errors}.go` | **手写静态基础设施** | 与契约无关的 HTTP 发起调用、序列化反序列化、参数编码与通用异常处理；提供生成 api/model 依赖的同包符号面 |
| `meoo/**` | **手写高层 Runtime** | 认证/超时/有限重试/SSE/分页/错误分层的正式运行时语义，以及公共客户端 facade |

`meoo` 包的业务模型直接复用 `generated` 的模型**类型**（如 `generated.Project`、`generated.AgentRun`），避免字段重复定义；但列表信封一律用局部结构 + 显式 `json` tag 解码，不引用生成模型的 Go 字段名，把“生成器字段命名变化”的编译期风险降到最低。

> **同包手写核心**：Go 生成器把 `client.go`/`configuration.go`/`response.go`/`utils.go` 与 `api_*.go`/`model_*.go` 放在同一个 `package generated`。这四个支撑文件正是 HTTP 发起、序列化反序列化与通用异常的所在，因此整组抽回手写，由手写文件维持生成 api/model 依赖的同包符号面（`APIClient`/`Configuration`/`APIResponse`/`GenericOpenAPIError`/`Nullable*`/`Ptr*`/`MappedNullable`/`reportError`/`parameterAddToHeaderOrQuery`/`selectHeaderContentType`/`selectHeaderAccept`/`decode`/`prepareRequest`/`callAPI` 等）。符号面若有偏差会在编译期暴露。

## 快速开始（facade）

facade 覆盖最高频的 operation，网络行为（认证、每请求动态取凭证、有限重试、RFC 7807 错误分层、SSE）全部收敛在手写 `Transport`：

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gitlab.alibaba-inc.com/oneday/meoo-sdk-go/meoo"
)

func main() {
	client, err := meoo.NewClient(meoo.WithAPIKey(os.Getenv("MEOO_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	// 惰性翻页（对齐 sql.Rows）：PageToken 由迭代器内部管理
	it := client.Projects().Iter(ctx, meoo.ListProjectsParams{}, nil)
	for it.Next() {
		fmt.Println(it.Item().UrlId)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

`Projects()` 提供 `Create`/`List`/`Iter`，`Agent()` 提供 `Start`/`Current`/`Cancel`/`Events`。

## 消费 SSE 事件流

`Agent().Events` 返回持有连接的 `*EventStream`，必须 `Close`（通常用 `defer`）。逐行惰性解析，收到契约声明的终态事件（`run.completed`、`run.failed`、`run.canceled`、`run.interrupted`、`run.superseded`）或流自然结束后 `Next` 返回 `io.EOF`，且不做重连：

```go
stream, err := client.Agent().Events(ctx, projectID, runID, nil)
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

## 使用生成客户端覆盖全部 operation

facade 未封装的 operation 用 `Generated()` 的生成客户端（request-builder 模式）。它与本客户端共享 `baseURL` 与 `*http.Client`；凭证需先用 `Context(ctx)` 注入到 ctx（生成层用 `ContextAccessToken` 承载 Bearer）：

```go
ctx, err := client.Context(context.Background())
if err != nil {
	return err
}
user, _, err := client.Generated().UserApi.GetUser(ctx).Execute()
if err != nil {
	return err
}
```

也可以直接用低层 `Transport()` 发原始契约路径的请求，认证/重试/错误语义与 facade 完全一致：

```go
raw, err := client.Transport().Request(ctx, "GET", "/open/v1/user", nil, nil)
```

## 错误分层

所有错误实现 `meoo.Error` 接口（`Message()`/`Unwrap()`），用 `errors.As` 判定具体类型，`errors.Is` 做链式判定：

- `*meoo.TransportError`：连接失败、超时、取消，以及请求体序列化 / 响应体解析失败。`Unwrap()` 保留底层 cause，支持 `errors.Is(err, context.Canceled)`。
- `*meoo.APIError`：服务端返回非 2xx。暴露 `Status`、`Code`、`TraceID`、原始 `Problem`（RFC 7807）与 `Headers`；`problem` 不可解析时按 HTTP status 兜底，`TraceID` 缺失时回退到响应头 `X-Meoo-Trace-Id`。
- 其余 SDK 级错误（如 SSE data 非法 JSON）仅通过 `Error` 接口对外暴露。

```go
var apiErr *meoo.APIError
if errors.As(err, &apiErr) && apiErr.Status == 404 {
	// 资源不存在
}
```

## 认证与重试

- **认证**：`Authorization: Bearer <credential>`，凭证每次请求动态取值。固定凭证用 `WithAPIKey`/`WithAccessToken`；动态凭证（OAuth refresh、成员 Token 重签发）用 `WithCredentialProvider` + `CredentialProviderFunc`，single-flight 与缓存由实现方负责。SSE 的凭证只走 Header，绝不进 query。
- **重试**：默认 `maxRetries=2`，仅对 429/502/503/504 且“可重试”的请求退避重试。GET/HEAD 或携带 `IdempotencyKey` 的写请求可重试；用 `RequestOptions.Retry`（配 `meoo.Bool(...)`）可逐请求覆盖。退避为指数退避（上限 8s），`Retry-After` 头优先。连接失败/超时直接归为 `*TransportError` 不重试，避免多语言行为漂移。

## 构建与测试

```bash
go build ./...     # 编译 generated + meoo 两个包
go test ./...      # 运行 meoo 包的单元测试（httptest 打桩，无需网络）
go vet ./...
```

## 来源与生成

本仓是**发布产物**：`generated/api_*.go`、`model_*.go` 由 OpenAPI Generator（版本锁定）从契约生成，连同手写静态基础设施与 `meoo` Runtime，统一在源仓库 `meoo-open-sdk`（monorepo 的 `sdks/go/`）中生成、验证，再镜像到此处。**生成器脚本、契约副本、实网 E2E（覆盖全部 35 operation）与跨语言一致性校验都保留在源仓库**，本仓不含它们，以保持发布产物精简。

`generated/api_*.go`、`model_*.go` 是契约驱动的生成产物，**禁止手改**；需要变更时请在源仓库改契约或生成约束后重新生成，再镜像过来。手写静态基础设施与 `meoo` Runtime 的修改同样在源仓库进行，避免两仓漂移。

## 约定

- **module 根即仓库根**，覆盖 `generated/` 与 `meoo/` 两个包，由手写的 `go.mod` 负责；`generated/` 不单独提供 `go.mod`。
- **`Transport` 自带 `*http.Client` 与请求构建，不经过 `generated.APIClient`**，以便统一控制超时/重试/SSE/认证；只在 facade（`projects.go`/`agent.go`）里复用 `generated` 的模型类型。`WithHTTPClient` 注入的 client 不应设置 `Timeout` 字段（会截断 SSE 长连接），超时由 SDK 通过 context 逐请求控制；默认 client 只设 `ResponseHeaderTimeout` 兜底响应头等待。
- **path 参数编码分两处**：生成层 `api_*.go` 用标准库 `url.PathEscape`（硬编码，不可改）；facade 层用 `encodeSegment`（跨语言契约口径的百分号编码）。二者只在 id 含保留字（`/ ; , ?` 等）时分叉，而真实的 `url_id`/`run_id` 都是 URL-safe 的，故实际一致。
- **异步**：Go 用 goroutine + `context` 表达并发，不单独提供异步门面；所有方法都是同步阻塞的，取消/超时交给 ctx。

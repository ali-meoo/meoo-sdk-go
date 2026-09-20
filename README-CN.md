[English](README.md) | 简体中文

# Meoo Open API Go SDK

Meoo 开放平台官方 Go SDK。**纯 Go 标准库实现，无任何第三方依赖**。提供并发安全的同步客户端，覆盖 Bearer / API Key 认证、统一错误分层、有限重试、项目分页与 Agent Run/SSE 事件流，并通过逃生舱客户端覆盖全部 **35 个** Open API operation。

## 环境要求

- Go 1.22 及以上。

## 安装

```sh
go get gitlab.alibaba-inc.com/oneday/meoo-sdk-go/client
```

> **内网私有 module**：消费方需设置 `GOPRIVATE=gitlab.alibaba-inc.com`（让 `go get` 跳过公共 proxy/sumdb 直连 VCS），并配置好 Git 凭证；若走 http 或自签证书，视情况再配 `GOINSECURE`。

## 快速开始

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
	// 包名是 client，与常见局部变量名冲突，惯例用别名 meoo 引入。
	c, err := meoo.NewClient(meoo.WithAPIKey(os.Getenv("MEOO_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	ctx := context.Background()

	// 创建项目
	project, err := c.Projects().Create(ctx, meoo.CreateProjectParams{Name: "demo"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created:", project.Name, project.UrlId)

	// 惰性翻页遍历所有项目（用法对齐 sql.Rows）
	it := c.Projects().Iter(ctx, meoo.ListProjectsParams{}, nil)
	for it.Next() {
		fmt.Println(it.Item().Name)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

## 构造客户端

`NewClient` 接受函数式选项；`WithAPIKey` / `WithAccessToken` / `WithCredentialProvider` 至少提供其一，否则返回错误。

| 选项 | 说明 | 默认值 |
|---|---|---|
| `WithAPIKey(key string)` | 固定 API Key（`meoo_ak`），以 Bearer 承载 | — |
| `WithAccessToken(token string)` | 固定 OAuth Access Token | — |
| `WithCredentialProvider(p CredentialProvider)` | 动态凭证来源；OAuth refresh / 成员 Token 重签发的 single-flight 与缓存由实现方负责 | — |
| `WithBaseURL(url string)` | 服务地址（末尾斜杠自动去除） | `https://meoo.com` |
| `WithTimeout(d time.Duration)` | 普通请求超时（不约束 SSE 长连接整体时长） | `30s` |
| `WithMaxRetries(n int)` | 最大重试次数（不含首次尝试） | `2` |
| `WithHTTPClient(hc *http.Client)` | 自定义 `*http.Client`（代理 / 连接池 / 测试打桩）；不要设其 `Timeout` 字段，否则会截断 SSE | 内置 |
| `WithBackoff(b Backoff)` | 自定义退避策略 | 指数退避（上限 8s） |

便捷构造：`NewWithAPIKey(key)`、`NewWithAccessToken(token)`。`Client` 实例**并发安全**，可跨 goroutine 复用；用完调用 `Close()` 释放空闲连接。

## 接口一览

### 高层门面（推荐日常使用）

**`client.Projects()` — 项目**

| 方法 | 签名 |
|---|---|
| `Create` | `(ctx, CreateProjectParams, *RequestOptions) (*Project, error)` |
| `List` | `(ctx, ListProjectsParams, *RequestOptions) (*Page[Project], error)` |
| `Iter` | `(ctx, ListProjectsParams, *RequestOptions) *ProjectIterator` — 惰性翻页 |

**`client.Agent()` — Agent Run**

| 方法 | 签名 |
|---|---|
| `Start` | `(ctx, projectID string, input any, *RequestOptions) (*AgentRun, error)` |
| `Current` | `(ctx, projectID string, *RequestOptions) (*AgentRun, error)` |
| `Cancel` | `(ctx, projectID, runID string, *RequestOptions) (*AgentRun, error)` |
| `Events` | `(ctx, projectID, runID string, *RequestOptions) (*EventStream, error)` — SSE 事件流 |

### 全量 operation（`Generated()` 逃生舱）

门面只封装最高频的 operation；其余全部 operation 通过 `Generated()` 返回的 request-builder 客户端调用，覆盖 **12 组 35 个** operation。它与主客户端共享 `baseURL` 与 `*http.Client`，调用前先用 `Context(ctx)` 注入凭证：

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

| 分组（`Generated()` 字段） | operation |
|---|---|
| `ProjectsApi` | CreateProject、CreateProjectToken、GetProjectWatermarkRemoval、ListProjects、UpdateProjectWatermarkRemoval |
| `AgentRunsApi` | StartAgentRun、GetCurrentAgentRun、CancelAgentRun、StreamAgentRunEvents、RespondToAgentAction、CreateAgentUpload |
| `AgentHistoryApi` | ListAgentConversations、ListAgentConversationMessages |
| `CloudDatabaseApi` | ExecuteCloudDatabaseQuery、GetCloudDatabaseStatus、ListCloudDatabaseTables |
| `CloudFunctionsApi` | ListCloudFunctions、ListCloudFunctionLogs |
| `CloudSecretsApi` | ListCloudSecrets、PutCloudSecret、DeleteCloudSecret |
| `CloudStorageApi` | ListCloudStorageBuckets、ListCloudStorageObjects |
| `PreviewApi` | CreateAgentPreviewLink、OpenPreviewShell |
| `ReleasesApi` | CreateRelease、GetCurrentRelease、ListReleases、PrepareReleaseUpload、CompleteReleaseUpload、UnpublishRelease |
| `SkillsApi` | ListSelectableSkills、UploadSkill |
| `SourceApi` | CreateCurrentProjectSourceExport |
| `UserApi` | GetUser |

所有请求 / 响应模型（如 `Project`、`AgentRun`、`CloudFunction`、`AgentMessageDeltaEventData`）都在 `client` 包内，直接以 `meoo.<类型名>` 使用，无需引入任何内部包。

### 低层传输（`Transport()`）

需要对未封装路径完全掌控时，用 `Transport()` 直接发请求，认证 / 超时 / 重试 / 错误语义与门面**完全一致**：

```go
raw, err := c.Transport().Request(ctx, "GET", "/open/v1/user", nil, nil)
```

## 消费 SSE 事件流

`Agent().Events` 返回持有连接的 `*EventStream`，**必须 `Close`**（通常 `defer`）。逐行惰性解析；收到契约声明的终态事件（`run.completed`、`run.failed`、`run.canceled`、`run.interrupted`、`run.superseded`）或流自然结束后，`Next` 返回 `io.EOF` 且不再重连：

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
		var delta meoo.AgentMessageDeltaEventData
		if err := event.DataAs(&delta); err != nil {
			return err
		}
		// 处理增量消息
	case "run.completed":
		var terminal meoo.AgentRunTerminalEvent
		_ = event.DataAs(&terminal)
		// 处理终态
	}
	// 契约要求忽略未知事件名：未知 event 原样透出、不报错
}
```

`AgentEvent` 字段：`ID`（诊断用）、`Event`（事件名，缺省 `message`）、`Data`（所有 `data:` 行以 `\n` 拼接的原始 JSON）；用 `DataAs(&v)` 映射到具体模型或任意结构。`IsTerminalEvent(name)` 判断是否终态事件。

## 错误处理

所有错误实现 `client.Error` 接口（`Message()` / `Unwrap()`），用 `errors.As` 判定具体类型、`errors.Is` 做链式判定：

| 类型 | 触发场景 | 关键字段 |
|---|---|---|
| `*client.APIError` | 服务端返回非 2xx | `Status`、`Code`、`TraceID`、`Problem`（RFC 7807 原始 JSON）、`Headers` |
| `*client.TransportError` | 连接失败、超时、取消，请求体序列化 / 响应体解析失败 | `Unwrap()` 保留底层 cause |
| `client.Error`（其余） | SDK 级错误（如 SSE data 非法 JSON） | 仅经接口暴露 |

```go
var apiErr *client.APIError
if errors.As(err, &apiErr) {
	switch apiErr.Status {
	case 404: // 资源不存在
	case 429: // 触发限流
	}
	fmt.Println("trace:", apiErr.TraceID)
}
if errors.Is(err, context.Canceled) {
	// 调用方主动取消
}
```

`APIError` 的 `problem` 不可解析时按 HTTP status 兜底（契约允许服务端新增错误码）；`TraceID` 缺失时回退到响应头 `X-Meoo-Trace-Id`。

## 分页

`Page[T]` 含 `Items []T`（永不为 nil）与 `NextPageToken`，`HasNextPage()` 判断是否还有下一页。`Iter` 返回的 `ProjectIterator` 自动管理翻页游标，用法对齐 `sql.Rows`：`for it.Next() { it.Item() }`，结束后 `it.Err()` 取错误。迭代器**非并发安全**，应由单个 goroutine 使用。

## 认证与重试

- **认证**：`Authorization: Bearer <credential>`，凭证**每次请求动态取值**（`CredentialProvider.Token`）。固定凭证用 `WithAPIKey` / `WithAccessToken`；动态凭证（OAuth refresh、成员 Token 重签发）用 `WithCredentialProvider` + `CredentialProviderFunc`，single-flight 与缓存由实现方负责。SSE 的凭证只走 Header，绝不进 query。
- **重试**：默认 `maxRetries=2`，仅对 `429/502/503/504` 且「可重试」的请求退避重试。GET/HEAD 或携带 `IdempotencyKey` 的写请求可重试；用 `RequestOptions.Retry`（配 `client.Bool(...)`）逐请求覆盖。退避为指数退避（上限 8s），`Retry-After` 头优先。连接失败 / 超时直接归为 `*TransportError`，不重试。

单次请求可用 `RequestOptions` 覆盖：`IdempotencyKey`（写入 `Idempotency-Key` 头）、`Timeout`、`Retry`；传 `nil` 表示沿用默认。

## 构建与测试

```sh
go build ./...
go test ./...     # 单元测试用 httptest 打桩，无需网络
go vet ./...
```

## 发行说明

各版本变更记录见 [ChangeLog.txt](./ChangeLog.txt)。

## 许可证

内部二方 SDK，许可证待定。

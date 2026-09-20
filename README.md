English | [简体中文](README-CN.md)

# Meoo Open API Go SDK

The official Go SDK for the Meoo Open API. Written with the **Go standard library only —
no third-party dependencies.** It provides a concurrency-safe synchronous client covering
Bearer / API-Key authentication, a unified error taxonomy, bounded retries, project
pagination, and Agent Run/SSE streaming, plus a full client that reaches every one of the
**35** Open API operations.

## Requirements

- Go 1.22 or later.

## Installation

```sh
go get github.com/ali-meoo/meoo-sdk-go/client
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	meoo "github.com/ali-meoo/meoo-sdk-go/client"
)

func main() {
	// The package is named client, which collides with a common local variable name,
	// so this example imports it under the alias meoo.
	c, err := meoo.NewClient(meoo.WithAPIKey(os.Getenv("MEOO_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	ctx := context.Background()

	project, err := c.Projects().Create(ctx, meoo.CreateProjectParams{Name: "demo"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created:", project.Name, project.UrlId)

	// Lazy pagination (sql.Rows-style).
	it := c.Projects().Iter(ctx, meoo.ListProjectsParams{}, nil)
	for it.Next() {
		fmt.Println(it.Item().Name)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

## Constructing a client

`NewClient` takes functional options. Provide at least one of `WithAPIKey` /
`WithAccessToken` / `WithCredentialProvider`, otherwise it returns an error.

| Option | Purpose | Default |
|---|---|---|
| `WithAPIKey(key string)` | Static API key, sent as a Bearer credential | — |
| `WithAccessToken(token string)` | Static OAuth access token | — |
| `WithCredentialProvider(p CredentialProvider)` | Dynamic credential source; single-flight & caching for OAuth refresh / token reissue are the caller's responsibility | — |
| `WithBaseURL(url string)` | Service base URL (trailing slash trimmed) | `https://meoo.com` |
| `WithTimeout(d time.Duration)` | Per-request timeout for normal calls (does not bound SSE streams) | `30s` |
| `WithMaxRetries(n int)` | Max retries (excluding the first attempt) | `2` |
| `WithHTTPClient(hc *http.Client)` | Custom `*http.Client` (proxy / pool / test stub); do not set its `Timeout` field or it will truncate SSE | built-in |
| `WithBackoff(b Backoff)` | Custom backoff strategy | exponential (capped at 8s) |

Convenience constructors: `NewWithAPIKey(key)`, `NewWithAccessToken(token)`. A `Client`
is **concurrency-safe** and meant to be reused across goroutines; call `Close()` to
release idle connections.

## API surface

### High-level resources (recommended)

**`client.Projects()`**

| Method | Signature |
|---|---|
| `Create` | `(ctx, CreateProjectParams, *RequestOptions) (*Project, error)` |
| `List` | `(ctx, ListProjectsParams, *RequestOptions) (*Page[Project], error)` |
| `Iter` | `(ctx, ListProjectsParams, *RequestOptions) *ProjectIterator` — lazy pagination |

**`client.Agent()`**

| Method | Signature |
|---|---|
| `Start` | `(ctx, projectID string, input any, *RequestOptions) (*AgentRun, error)` |
| `Current` | `(ctx, projectID string, *RequestOptions) (*AgentRun, error)` |
| `Cancel` | `(ctx, projectID, runID string, *RequestOptions) (*AgentRun, error)` |
| `Events` | `(ctx, projectID, runID string, *RequestOptions) (*EventStream, error)` — SSE stream |

### Every operation (`Generated()`)

The high-level resources wrap the most frequent operations. Everything else is reachable
through the request-builder client returned by `Generated()`, covering **12 groups / 35
operations**. It shares the base URL and `*http.Client` with the main client; inject
credentials first via `Context(ctx)`:

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

| Group (`Generated()` field) | Operations |
|---|---|
| `ProjectsApi` | CreateProject, CreateProjectToken, GetProjectWatermarkRemoval, ListProjects, UpdateProjectWatermarkRemoval |
| `AgentRunsApi` | StartAgentRun, GetCurrentAgentRun, CancelAgentRun, StreamAgentRunEvents, RespondToAgentAction, CreateAgentUpload |
| `AgentHistoryApi` | ListAgentConversations, ListAgentConversationMessages |
| `CloudDatabaseApi` | ExecuteCloudDatabaseQuery, GetCloudDatabaseStatus, ListCloudDatabaseTables |
| `CloudFunctionsApi` | ListCloudFunctions, ListCloudFunctionLogs |
| `CloudSecretsApi` | ListCloudSecrets, PutCloudSecret, DeleteCloudSecret |
| `CloudStorageApi` | ListCloudStorageBuckets, ListCloudStorageObjects |
| `PreviewApi` | CreateAgentPreviewLink, OpenPreviewShell |
| `ReleasesApi` | CreateRelease, GetCurrentRelease, ListReleases, PrepareReleaseUpload, CompleteReleaseUpload, UnpublishRelease |
| `SkillsApi` | ListSelectableSkills, UploadSkill |
| `SourceApi` | CreateCurrentProjectSourceExport |
| `UserApi` | GetUser |

Every request/response model (e.g. `Project`, `AgentRun`, `CloudFunction`,
`AgentMessageDeltaEventData`) lives in the `client` package — use it directly as
`meoo.<Type>`.

### Low-level transport (`Transport()`)

For full control over an unwrapped path, call `Transport()` directly. Auth, timeout,
retry, and error semantics are **identical** to the high-level resources:

```go
raw, err := c.Transport().Request(ctx, "GET", "/open/v1/user", nil, nil)
```

## Consuming SSE streams

`Agent().Events` returns a connection-holding `*EventStream` that **must be closed**
(usually via `defer`). Frames are parsed lazily line by line. After a terminal event
(`run.completed`, `run.failed`, `run.canceled`, `run.interrupted`, `run.superseded`) or a
natural end of stream, `Next` returns `io.EOF` and never reconnects:

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
		// handle incremental message
	case "run.completed":
		var terminal meoo.AgentRunTerminalEvent
		_ = event.DataAs(&terminal)
		// handle terminal state
	}
	// Unknown event names are passed through as-is and never fail the stream.
}
```

`AgentEvent` fields: `ID` (diagnostics), `Event` (name, defaults to `message`), `Data`
(raw JSON of all `data:` lines joined by `\n`). Use `DataAs(&v)` to map it onto a model
or any struct. `IsTerminalEvent(name)` reports whether a name is terminal.

## Error handling

Every error implements the `client.Error` interface (`Message()` / `Unwrap()`). Use
`errors.As` to match a concrete type and `errors.Is` for chained checks:

| Type | When | Key fields |
|---|---|---|
| `*client.APIError` | Server returned non-2xx | `Status`, `Code`, `TraceID`, `Problem` (raw RFC 7807 JSON), `Headers` |
| `*client.TransportError` | Connect failure, timeout, cancellation, request/response (de)serialization failure | `Unwrap()` keeps the underlying cause |
| `client.Error` (other) | SDK-level errors (e.g. malformed SSE data JSON) | exposed via the interface only |

```go
var apiErr *client.APIError
if errors.As(err, &apiErr) {
	switch apiErr.Status {
	case 404: // not found
	case 429: // rate limited
	}
	fmt.Println("trace:", apiErr.TraceID)
}
if errors.Is(err, context.Canceled) {
	// caller canceled
}
```

When `problem` cannot be parsed, `APIError` falls back to the HTTP status (the server may
introduce new error codes); when `TraceID` is absent it falls back to the
`X-Meoo-Trace-Id` response header.

## Pagination

`Page[T]` carries `Items []T` (never nil) and `NextPageToken`; `HasNextPage()` reports
whether more pages remain. The `ProjectIterator` from `Iter` manages the cursor for you,
`sql.Rows`-style: `for it.Next() { it.Item() }`, then `it.Err()`. The iterator is **not**
concurrency-safe and should be driven by a single goroutine.

## Authentication & retries

- **Auth:** `Authorization: Bearer <credential>`, resolved **per request**
  (`CredentialProvider.Token`). Static credentials via `WithAPIKey` / `WithAccessToken`;
  dynamic credentials (OAuth refresh, member-token reissue) via `WithCredentialProvider`
  + `CredentialProviderFunc` — single-flight and caching are the caller's job. SSE
  credentials go only in headers, never in the query string.
- **Retries:** default `maxRetries=2`, backing off only on `429/502/503/504` for
  retryable requests. GET/HEAD, or writes carrying an `IdempotencyKey`, are retryable;
  override per request with `RequestOptions.Retry` (use `client.Bool(...)`). Backoff is
  exponential (capped at 8s) and honors `Retry-After` first. Connect failures / timeouts
  become `*TransportError` and are not retried.

Per-request overrides live in `RequestOptions`: `IdempotencyKey` (sent as
`Idempotency-Key`), `Timeout`, `Retry`. Pass `nil` to use the defaults.

## Build & test

```sh
go build ./...
go test ./...     # unit tests use httptest stubs; no network required
go vet ./...
```

## Changelog

See [ChangeLog.txt](./ChangeLog.txt) for per-version changes.

## License

To be determined; a `LICENSE` file will be added before the first public release.

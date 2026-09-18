English | [简体中文](README-CN.md)

# Meoo Open API Go SDK

Official Go SDK for the Meoo Open API. Pure Go standard library, no third-party
dependencies. Provides a synchronous `client.Client` covering Bearer/API-Key auth,
a unified error taxonomy, bounded retries, project pagination and Agent Run/SSE streaming.

## Requirements

- Go 1.22 or later.

## Installation

```sh
go get gitlab.alibaba-inc.com/oneday/meoo-sdk-go/client
```

> Internal private module: consumers must set `GOPRIVATE=gitlab.alibaba-inc.com`
> (so `go get` bypasses the public proxy/sumdb and fetches directly from VCS) and
> have Git credentials configured.

## Layout

The module root is the repository root, containing two packages:

| Package / dir | Source | Purpose |
|---|---|---|
| `generated/api_*.go`, `generated/model_*.go` | OpenAPI Generator (do not edit) | Contract-driven request-builders, params, schema structs |
| `generated/{client,configuration,response,utils,transport,codec,params,errors}.go` | Handwritten static core | HTTP dispatch, (de)serialization, param encoding, common errors |
| `client/**` | Handwritten high-level runtime | Auth / timeout / bounded retry / SSE / pagination / error taxonomy, and the public facade |

## Usage

The facade lives in `client/` (`package client`). Because `client` collides with the
common local variable name, examples import it under the alias `meoo` (the convention
used by Alibaba Cloud SDKs):

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
	it := c.Projects().Iter(ctx, meoo.ListProjectsParams{}, nil)
	for it.Next() {
		fmt.Println(it.Item().UrlId)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

- `Projects()`: `Create` / `List` / `Iter` (lazy pagination, mirrors `sql.Rows`)
- `Agent()`: `Start` / `Current` / `Cancel` / `Events` (SSE stream)
- `Generated()`: escape hatch to the generated client covering every operation
- `Transport()`: low-level raw request outlet

Errors implement `client.Error`; use `errors.As` to inspect `*client.APIError`
(server non-2xx, RFC 7807) or `*client.TransportError` (connect/timeout/cancel and
(de)serialization failures). See [README-CN.md](README-CN.md) for the full guide.

## Build & Test

```sh
go build ./...
go test ./...
go vet ./...
```

## Provenance

This repository is a published artifact: `generated/api_*.go` and `model_*.go` are
produced by OpenAPI Generator from the contract, then mirrored here from the source
monorepo `meoo-open-sdk` (`sdks/go/`). Generator scripts, the contract copy and the
live E2E suite stay in the source monorepo. Never edit `generated/api_*.go` or
`model_*.go` by hand.

## Changelog

See [ChangeLog.txt](./ChangeLog.txt) for per-release changes.

## License

Apache-2.0 (see `LICENSE`).

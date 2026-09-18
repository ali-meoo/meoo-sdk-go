// Meoo Open API Go SDK。
//
// 本仓是 Go SDK 的独立发布仓，module 根即仓库根，下辖两个包：
//   - generated/：OpenAPI Generator 产出的 api/model（动态、契约驱动）
//     + 手写静态基础设施（HTTP 发起、序列化反序列化、通用异常，见 .openapi-generator-ignore）
//   - meoo/     ：手写高层 Runtime 与公共客户端 facade（认证、超时、有限重试、SSE、分页）
//
// 仅依赖 Go 标准库，无第三方 require：认证收敛为 Bearer + API Key，OAuth refresh 由
// meoo 包的 CredentialProvider 承载，因此不引入 golang.org/x/oauth2，CI 无需拉取外部模块。
module gitlab.alibaba-inc.com/oneday/meoo-sdk-go

go 1.22

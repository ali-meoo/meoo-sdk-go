/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * 动态凭证来源，落实 runtime-spec/auth.md 第 2 条：普通资源 API 使用
 * Authorization: Bearer <credential>，且凭证必须支持每次请求动态取值——OAuth refresh 与
 * 成员 Token 重签发所需的 single-flight 与缓存由实现方负责，SDK 只在发请求前取一次值。
 *
 * 对齐 Java com.meoo.runtime.CredentialProvider 与 TypeScript CredentialProvider。
 */

package client

import "context"

// CredentialProvider 在每次请求前提供一个 Bearer 凭证（API Key 或 OAuth Access Token）。
// 返回 error 时，Transport 将其包装为 *TransportError 并中止本次请求。
type CredentialProvider interface {
	Token(ctx context.Context) (string, error)
}

// CredentialProviderFunc 让普通函数满足 CredentialProvider，便于内联实现动态取值，例如：
//
//	meoo.WithCredentialProvider(meoo.CredentialProviderFunc(func(ctx context.Context) (string, error) {
//	    return oauthCache.Get(ctx) // 实现方自行保证 single-flight 与缓存
//	}))
type CredentialProviderFunc func(ctx context.Context) (string, error)

// Token 实现 CredentialProvider。
func (f CredentialProviderFunc) Token(ctx context.Context) (string, error) { return f(ctx) }

// staticCredential 返回一个固定凭证的 Provider。空值校验统一推迟到 NewClient，
// 以便和其它选项一起给出一致的错误信息。
func staticCredential(token string) CredentialProvider {
	return CredentialProviderFunc(func(context.Context) (string, error) { return token, nil })
}

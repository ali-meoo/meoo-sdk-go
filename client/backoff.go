/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * 重试等待策略，语义见 runtime-spec/retry.md，与 Java com.meoo.runtime.Backoff、
 * TypeScript retryDelay、Python Runtime 逐条对齐：Retry-After（秒）优先，否则
 * min(500ms * 2^attempt, 8s)。契约要求 SSE 的 429 在等待基础上增加抖动；需要时通过
 * WithBackoff 注入自定义实现，不要改默认值，否则三语言退避行为会漂移
 * （改默认值必须先改 runtime-spec 并同步 Java/TS/Python/Go 四语言）。
 */

package client

import (
	"strconv"
	"strings"
	"time"
)

// Backoff 计算第 attempt 次重试前的等待时长，attempt 从 0 开始表示「已完成的尝试次数」。
// retryAfter 是响应的 Retry-After 头取值，可能为空。
type Backoff func(attempt int, retryAfter []string) time.Duration

// maxBackoffShift 限制指数增长的位移，避免 1<<attempt 在 attempt 很大时溢出；与 Java 的
// Math.min(attempt, 20) 一致。
const maxBackoffShift = 20

// ExponentialBackoff 是默认退避策略。
func ExponentialBackoff(attempt int, retryAfter []string) time.Duration {
	if len(retryAfter) > 0 {
		if seconds, err := strconv.ParseFloat(strings.TrimSpace(retryAfter[0]), 64); err == nil && seconds >= 0 {
			return time.Duration(seconds * float64(time.Second))
		}
		// HTTP-date 形式的 Retry-After 不解析，退回指数退避（与 Java/TypeScript 一致）。
	}
	shift := attempt
	if shift < 0 {
		shift = 0
	}
	if shift > maxBackoffShift {
		shift = maxBackoffShift
	}
	millis := int64(500) * (int64(1) << shift)
	if millis > 8000 {
		millis = 8000
	}
	return time.Duration(millis) * time.Millisecond
}

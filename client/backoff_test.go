/*
 * Meoo Open API Go SDK —— package client 测试。
 *
 * 镜像 Java com.meoo.runtime.BackoffTest：默认退避必须与 TypeScript retryDelay、Python
 * request 的等待时长完全一致（runtime-spec/retry.md）。断言逐条对应，改默认值前必须同步四语言。
 */

package client

import (
	"testing"
	"time"
)

func TestExponentialBackoffIsCappedAtEightSeconds(t *testing.T) {
	cases := []struct {
		attempt int
		want    time.Duration
	}{
		{0, 500 * time.Millisecond},
		{1, 1000 * time.Millisecond},
		{2, 2000 * time.Millisecond},
		{9, 8 * time.Second},
	}
	for _, c := range cases {
		if got := ExponentialBackoff(c.attempt, nil); got != c.want {
			t.Errorf("ExponentialBackoff(%d, nil) = %v, want %v", c.attempt, got, c.want)
		}
	}
}

func TestRetryAfterHeaderWins(t *testing.T) {
	cases := []struct {
		attempt    int
		retryAfter []string
		want       time.Duration
	}{
		{0, []string{"60"}, 60 * time.Second},
		{3, []string{"2.5"}, 2500 * time.Millisecond},
		// Retry-After: 0 表示立即可重试，此时不再退回指数退避
		{0, []string{"0"}, 0},
	}
	for _, c := range cases {
		if got := ExponentialBackoff(c.attempt, c.retryAfter); got != c.want {
			t.Errorf("ExponentialBackoff(%d, %q) = %v, want %v", c.attempt, c.retryAfter, got, c.want)
		}
	}
}

func TestUnsupportedRetryAfterFallsBackToExponential(t *testing.T) {
	// HTTP-date 形式不解析，与 TypeScript 的 Number() 行为一致
	if got := ExponentialBackoff(0, []string{"Wed, 21 Oct 2015 07:28:00 GMT"}); got != 500*time.Millisecond {
		t.Errorf("HTTP-date Retry-After should fall back to exponential, got %v, want %v", got, 500*time.Millisecond)
	}
	// 负数秒不合法，同样退回指数退避
	if got := ExponentialBackoff(1, []string{"-5"}); got != 1000*time.Millisecond {
		t.Errorf("negative Retry-After should fall back to exponential, got %v, want %v", got, 1000*time.Millisecond)
	}
}

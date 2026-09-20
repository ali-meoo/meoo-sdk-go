/*
 * Meoo Open API Go SDK — 同步 SSE 事件流测试。
 *
 * 验证：按行惰性交付、收到终态事件后 Next 返回 io.EOF 且不再重连、流自然结束返回 io.EOF、
 * 非法 data 透出解析错误、Close 幂等。
 */

package client

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestIsTerminalEvent(t *testing.T) {
	terminal := []string{"run.completed", "run.failed", "run.canceled", "run.interrupted", "run.superseded"}
	for _, name := range terminal {
		if !IsTerminalEvent(name) {
			t.Errorf("IsTerminalEvent(%q) = false, want true", name)
		}
	}
	nonTerminal := []string{"message.delta", "message.snapshot", "run.working", "run.some_future_event", ""}
	for _, name := range nonTerminal {
		if IsTerminalEvent(name) {
			t.Errorf("IsTerminalEvent(%q) = true, want false", name)
		}
	}
}

func TestEventStreamStopsAfterTerminalEventWithoutReconnect(t *testing.T) {
	frames := "id: 1\nevent: message.delta\ndata: {\"delta\":\"hi\"}\n\n" +
		"id: 2\nevent: run.completed\ndata: {\"status\":\"completed\"}\n\n" +
		"id: 3\nevent: message.delta\ndata: {\"delta\":\"after-terminal\"}\n\n"
	stream := newEventStream(io.NopCloser(strings.NewReader(frames)))
	defer stream.Close()

	first, err := stream.Next()
	if err != nil {
		t.Fatalf("first Next() error: %v", err)
	}
	if first.Event != "message.delta" {
		t.Errorf("first event = %q, want %q", first.Event, "message.delta")
	}

	second, err := stream.Next()
	if err != nil {
		t.Fatalf("second Next() error: %v", err)
	}
	if second.Event != "run.completed" {
		t.Errorf("second event = %q, want %q (terminal event must still be delivered)", second.Event, "run.completed")
	}

	// 终态事件之后不再交付后续帧，也不重连：直接 io.EOF
	if _, err := stream.Next(); !errors.Is(err, io.EOF) {
		t.Errorf("Next() after terminal event = %v, want io.EOF", err)
	}
}

func TestEventStreamNaturalEndReturnsEOF(t *testing.T) {
	frames := "event: message.delta\ndata: {\"delta\":\"a\"}\n\n" +
		"event: message.delta\ndata: {\"delta\":\"b\"}\n\n"
	stream := newEventStream(io.NopCloser(strings.NewReader(frames)))
	defer stream.Close()

	count := 0
	for {
		event, err := stream.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("Next() error: %v", err)
		}
		if event != nil {
			count++
		}
	}
	if count != 2 {
		t.Errorf("delivered %d events, want 2", count)
	}
}

func TestEventStreamTrailingFrameWithoutBlankLineIsDelivered(t *testing.T) {
	// 服务端可能在最后一帧后直接断开，未必带收尾空行
	stream := newEventStream(io.NopCloser(strings.NewReader("event: run.working\ndata: {\"status\":\"working\"}")))
	defer stream.Close()

	event, err := stream.Next()
	if err != nil {
		t.Fatalf("Next() error: %v", err)
	}
	if event == nil || event.Event != "run.working" {
		t.Fatalf("trailing frame not delivered, got %v", event)
	}
	if _, err := stream.Next(); !errors.Is(err, io.EOF) {
		t.Errorf("Next() after trailing frame = %v, want io.EOF", err)
	}
}

func TestEventStreamMalformedDataSurfacesError(t *testing.T) {
	stream := newEventStream(io.NopCloser(strings.NewReader("event: message.delta\ndata: {\n\n")))
	defer stream.Close()

	_, err := stream.Next()
	if err == nil {
		t.Fatal("malformed data should surface an error, got nil")
	}
	if errors.Is(err, io.EOF) {
		t.Errorf("malformed data should not be reported as io.EOF")
	}
}

func TestEventStreamCloseIsIdempotent(t *testing.T) {
	stream := newEventStream(io.NopCloser(strings.NewReader("")))
	if err := stream.Close(); err != nil {
		t.Errorf("first Close() = %v, want nil", err)
	}
	if err := stream.Close(); err != nil {
		t.Errorf("second Close() = %v, want nil (Close must be idempotent)", err)
	}
}

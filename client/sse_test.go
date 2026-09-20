/*
 * Meoo Open API Go SDK — SSE 帧解析测试。
 */

package client

import (
	"encoding/json"
	"errors"
	"testing"
)

// dataField 从事件 data 里取一个字符串字段（仅测试用）。
func dataField(t *testing.T, event *AgentEvent, key string) string {
	t.Helper()
	var fields map[string]interface{}
	if err := json.Unmarshal(event.Data, &fields); err != nil {
		t.Fatalf("event data is not a JSON object: %v (data=%s)", err, event.Data)
	}
	value, ok := fields[key].(string)
	if !ok {
		t.Fatalf("event data field %q missing or not a string (data=%s)", key, event.Data)
	}
	return value
}

func TestParseSingleFrame(t *testing.T) {
	event, err := ParseFrame("id: 7\nevent: message.delta\ndata: {\"delta\":\"hi\"}")
	if err != nil {
		t.Fatalf("ParseFrame returned error: %v", err)
	}
	if event == nil {
		t.Fatal("ParseFrame returned nil event for a frame with data")
	}
	if event.ID != "7" {
		t.Errorf("event.ID = %q, want %q", event.ID, "7")
	}
	if event.Event != "message.delta" {
		t.Errorf("event.Event = %q, want %q", event.Event, "message.delta")
	}
	if got := dataField(t, event, "delta"); got != "hi" {
		t.Errorf("event data delta = %q, want %q", got, "hi")
	}
}

func TestParseSharedFixture(t *testing.T) {
	events, err := ParseStream(readFixture(t, "sse/basic.txt"))
	if err != nil {
		t.Fatalf("ParseStream returned error: %v", err)
	}
	if len(events) != 3 {
		t.Fatalf("fixture parsed into %d events, want 3", len(events))
	}

	wantNames := []string{"message.delta", "message.snapshot", "run.completed"}
	wantIDs := []string{"1", "2", "3"}
	for i, event := range events {
		if event.Event != wantNames[i] {
			t.Errorf("events[%d].Event = %q, want %q", i, event.Event, wantNames[i])
		}
		if event.ID != wantIDs[i] {
			t.Errorf("events[%d].ID = %q, want %q", i, event.ID, wantIDs[i])
		}
	}
	if got := dataField(t, events[0], "delta"); got != "hello" {
		t.Errorf("events[0] delta = %q, want %q", got, "hello")
	}
	if got := dataField(t, events[1], "content"); got != "hello world" {
		t.Errorf("events[1] content = %q, want %q", got, "hello world")
	}
	if got := dataField(t, events[2], "status"); got != "completed" {
		t.Errorf("events[2] status = %q, want %q", got, "completed")
	}
}

func TestIgnoresHeartbeatCommentsAndKeepsUnknownEvents(t *testing.T) {
	// heartbeat 用 SSE 注释帧承载；未知事件名不得导致流失败（前向兼容）
	events, err := ParseStream(": heartbeat\n\nevent: run.some_future_event\ndata: {\"run_id\":\"r1\"}\n\n")
	if err != nil {
		t.Fatalf("ParseStream returned error: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("parsed %d events, want 1 (heartbeat comment frame must be dropped)", len(events))
	}
	if events[0].Event != "run.some_future_event" {
		t.Errorf("event name = %q, want %q", events[0].Event, "run.some_future_event")
	}
	if events[0].ID != "" {
		t.Errorf("event.ID = %q, want empty", events[0].ID)
	}
}

func TestDefaultsEventNameToMessageAndJoinsDataLines(t *testing.T) {
	event, err := ParseFrame("data: {\"run_id\":\"r1\",\ndata: \"status\":\"working\"}")
	if err != nil {
		t.Fatalf("ParseFrame returned error: %v", err)
	}
	if event == nil {
		t.Fatal("ParseFrame returned nil event")
	}
	if event.Event != "message" {
		t.Errorf("default event name = %q, want %q", event.Event, "message")
	}
	if got := dataField(t, event, "status"); got != "working" {
		t.Errorf("status = %q, want %q (multi-line data must be joined with \\n)", got, "working")
	}
}

func TestDropsFramesWithoutDataAndKeepsTrailingFrame(t *testing.T) {
	if event, err := ParseFrame("event: message.delta"); err != nil || event != nil {
		t.Errorf("frame without data = (%v, %v), want (nil, nil)", event, err)
	}
	if events, err := ParseStream(": keep-alive\n\n"); err != nil || len(events) != 0 {
		t.Errorf("comment-only stream = (%d events, %v), want (0, nil)", len(events), err)
	}
	// 流结束时最后一帧没有空行也要交付
	events, err := ParseStream("event: run.working\ndata: {\"status\":\"working\"}")
	if err != nil {
		t.Fatalf("ParseStream returned error: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("trailing frame without blank line parsed into %d events, want 1", len(events))
	}
}

func TestMalformedDataFailsTheFrame(t *testing.T) {
	event, err := ParseFrame("event: message.delta\ndata: {")
	if err == nil {
		t.Fatalf("malformed data should fail, got event=%v err=nil", event)
	}
	if event != nil {
		t.Errorf("malformed data should not return an event, got %v", event)
	}
	var meooErr *meooError
	if !errors.As(err, &meooErr) {
		t.Errorf("malformed data error should be a *meooError, got %T", err)
	}
}

func TestMapsEventDataToModels(t *testing.T) {
	event, err := ParseFrame("event: run.completed\ndata: {\"run_id\":\"r1\",\"status\":\"completed\"}")
	if err != nil {
		t.Fatalf("ParseFrame returned error: %v", err)
	}
	// 用局部结构映射，避免耦合模型的 Go 字段名
	var data struct {
		RunID  string `json:"run_id"`
		Status string `json:"status"`
	}
	if err := event.DataAs(&data); err != nil {
		t.Fatalf("DataAs returned error: %v", err)
	}
	if data.RunID != "r1" || data.Status != "completed" {
		t.Errorf("DataAs mapped to %+v, want RunID=r1 Status=completed", data)
	}
}

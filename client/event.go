/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * SSE 事件帧模型，对齐 Java com.meoo.runtime.AgentEvent 与 TypeScript AgentEvent。
 * 契约前向兼容条款要求客户端忽略未知的 SSE 事件名，因此未知 event 原样透出且不报错。
 */

package client

import "encoding/json"

// AgentEvent 是一个 SSE 事件帧（runtime-spec/streaming.md 第 2 条）。
type AgentEvent struct {
	// ID 是事件帧 id（heartbeat 注释帧不带 id），仅用于诊断——服务端不支持 Last-Event-ID 续传。
	// 缺省时为空串。
	ID string
	// Event 是事件名，缺省为 "message"。
	Event string
	// Data 是所有 data: 行以 "\n" 拼接后的原始 JSON 字节，交给 DataAs 映射到具体模型。
	Data json.RawMessage
}

// DataAs 把 Data 反序列化到生成模型或任意结构，例如：
//
//	var terminal AgentRunTerminalEvent
//	if err := event.DataAs(&terminal); err != nil { ... }
//
// 映射失败返回 *meooError（包装底层 json 错误）。
func (e *AgentEvent) DataAs(v interface{}) error {
	if err := json.Unmarshal(e.Data, v); err != nil {
		return newMeooError("failed to map event data", err)
	}
	return nil
}

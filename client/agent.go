/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * Agent Run 资源，语义对齐 Java com.meoo.runtime.AgentResource 与 TypeScript AgentResource。
 * Events 返回的 *EventStream 持有连接，必须 Close。仅复用 AgentRun 类型，不引用其字段名。
 */

package client

import (
	"context"
	"encoding/json"
)

// AgentResource 封装 Agent Run 的高频 operation。
type AgentResource struct {
	transport *Transport
}

func newAgentResource(transport *Transport) *AgentResource {
	return &AgentResource{transport: transport}
}

// Start 启动一个 Run；input 可以是 map[string]interface{} 或生成的请求模型
// （例如 AgentRunStartRequest）。不携带 conversation_id 时由服务端隐式创建会话。
func (r *AgentResource) Start(ctx context.Context, projectID string, input interface{}, opts *RequestOptions) (*AgentRun, error) {
	raw, err := r.transport.Request(ctx, "POST", runsPath(projectID), input, opts)
	if err != nil {
		return nil, err
	}
	return decodeAgentRun(raw)
}

// Current 返回当前活跃 Run；没有活跃 Run 时返回最近一次终态 Run（契约行为，调用方无需区分）。
func (r *AgentResource) Current(ctx context.Context, projectID string, opts *RequestOptions) (*AgentRun, error) {
	raw, err := r.transport.Request(ctx, "GET", runsPath(projectID)+"/current", nil, opts)
	if err != nil {
		return nil, err
	}
	return decodeAgentRun(raw)
}

// Cancel 取消指定 Run；契约未声明请求体，因此不发送 body（与生成的 AgentRunsApi 一致）；
// 已是 canceled 时幂等返回 200。
func (r *AgentResource) Cancel(ctx context.Context, projectID, runID string, opts *RequestOptions) (*AgentRun, error) {
	path := runsPath(projectID) + "/" + encodeSegment(runID) + "/cancellations"
	raw, err := r.transport.Request(ctx, "POST", path, nil, opts)
	if err != nil {
		return nil, err
	}
	return decodeAgentRun(raw)
}

// Events 打开 Run 的 SSE 事件流（runtime-spec/streaming.md）；返回的 *EventStream 必须 Close。
func (r *AgentResource) Events(ctx context.Context, projectID, runID string, opts *RequestOptions) (*EventStream, error) {
	path := runsPath(projectID) + "/" + encodeSegment(runID) + "/events"
	return r.transport.Stream(ctx, path, opts)
}

// runsPath 拼接项目下的 agent runs 集合路径，projectID 整体编码。
func runsPath(projectID string) string {
	return "/open/v1/projects/" + encodeSegment(projectID) + "/agent/runs"
}

// decodeAgentRun 把响应 JSON 映射为 AgentRun；raw 为空时返回 (nil, nil)。
func decodeAgentRun(raw json.RawMessage) (*AgentRun, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	var run AgentRun
	if err := json.Unmarshal(raw, &run); err != nil {
		return nil, newMeooError("failed to map response to AgentRun", err)
	}
	return &run, nil
}

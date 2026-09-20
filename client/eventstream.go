/*
 * Meoo Open API Go SDK — 同步 SSE 事件流。
 *
 * 按行惰性读取，收到终态事件或流结束后 Next 返回 io.EOF，且不做重连。用法为 sql.Rows
 * 风格的 Next()/Close()。
 *
 * 用 bufio.Reader 而非 bufio.Scanner：Scanner 单行默认上限 64KB，会截断较大的 message.snapshot
 * 帧；Reader.ReadString 无此限制。
 */

package client

import (
	"bufio"
	"io"
	"strings"
)

// terminalEvents 是终态与 superseded 事件名。
var terminalEvents = map[string]struct{}{
	"run.completed":   {},
	"run.failed":      {},
	"run.canceled":    {},
	"run.interrupted": {},
	"run.superseded":  {},
}

// IsTerminalEvent 判断事件名是否为终态事件；未知事件名一律视为非终态。
func IsTerminalEvent(event string) bool {
	_, ok := terminalEvents[event]
	return ok
}

// EventStream 是持有连接的同步 SSE 事件流，必须调用 Close 释放（通常用 defer）。用法：
//
//	stream, err := client.Agent().Events(ctx, projectID, runID, nil)
//	if err != nil {
//	    return err
//	}
//	defer stream.Close()
//	for {
//	    event, err := stream.Next()
//	    if errors.Is(err, io.EOF) {
//	        break
//	    }
//	    if err != nil {
//	        return err
//	    }
//	    // 处理 event
//	}
type EventStream struct {
	reader   *bufio.Reader
	body     io.ReadCloser
	frame    []string
	pending  *AgentEvent
	finished bool
	closed   bool
}

// newEventStream 包装响应体为事件流；body 的关闭由 EventStream 负责。
func newEventStream(body io.ReadCloser) *EventStream {
	return &EventStream{reader: bufio.NewReader(body), body: body}
}

// Next 返回下一个事件：
//   - 流自然结束，或已交付终态事件后再次调用，返回 io.EOF；
//   - 读取失败返回 *TransportError；
//   - SSE data 非法 JSON 返回 *meooError。
func (s *EventStream) Next() (*AgentEvent, error) {
	for {
		if s.pending != nil {
			event := s.pending
			s.pending = nil
			return event, nil
		}
		if s.finished {
			return nil, io.EOF
		}

		line, err := s.reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")
		}

		if err != nil {
			if err == io.EOF {
				// 服务端可能在终态事件后直接断开，最后一帧未必带空行。
				if line != "" {
					s.frame = append(s.frame, line)
				}
				event, flushErr := s.flush()
				if flushErr != nil {
					s.finish()
					return nil, flushErr
				}
				if event != nil {
					s.setPending(event)
					continue
				}
				s.finish()
				return nil, io.EOF
			}
			s.finish()
			return nil, newTransportError("failed to read event stream", err)
		}

		if line == "" {
			// 空行是一帧的结束。
			event, flushErr := s.flush()
			if flushErr != nil {
				s.finish()
				return nil, flushErr
			}
			if event != nil {
				s.setPending(event)
			}
			continue
		}
		s.frame = append(s.frame, line)
	}
}

// Close 释放底层连接；多次调用安全。
func (s *EventStream) Close() error {
	if s.closed {
		return nil
	}
	s.closed = true
	if s.body != nil {
		return s.body.Close()
	}
	return nil
}

// setPending 暂存待交付事件；若是终态事件则标记结束并立即关闭连接（不再重连）。
func (s *EventStream) setPending(event *AgentEvent) {
	s.pending = event
	if IsTerminalEvent(event.Event) {
		s.finish()
	}
}

// flush 解析当前累积的帧；帧为空返回 (nil, nil)。
func (s *EventStream) flush() (*AgentEvent, error) {
	if len(s.frame) == 0 {
		return nil, nil
	}
	event, err := ParseFrame(strings.Join(s.frame, "\n"))
	s.frame = nil
	return event, err
}

// finish 标记流结束并关闭连接。
func (s *EventStream) finish() {
	s.finished = true
	_ = s.Close()
}

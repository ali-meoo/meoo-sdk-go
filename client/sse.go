/*
 * Meoo Open API Go SDK — SSE 帧解析。
 *
 * 按 SSE 规范解析事件帧：忽略注释帧（heartbeat 用注释帧承载）与空行、按首个冒号分割字段、
 * 字段值只去掉一个前导空格、多行 data 以 "\n" 拼接、默认事件名 message、没有 data 的帧丢弃、
 * data 非法 JSON 时报错。
 */

package client

import (
	"encoding/json"
	"strings"
)

// ParseFrame 解析单个事件帧；帧内没有 data 行时返回 (nil, nil)，data 非法 JSON 时返回错误。
func ParseFrame(frame string) (*AgentEvent, error) {
	id := ""
	event := "message"
	var data []string

	for _, line := range strings.Split(frame, "\n") {
		line = strings.TrimSuffix(line, "\r")
		// 注释帧（以 ":" 开头）与空行直接跳过。
		if line == "" || strings.HasPrefix(line, ":") {
			continue
		}
		field, value := line, ""
		if at := strings.Index(line, ":"); at >= 0 {
			field = line[:at]
			value = line[at+1:]
		}
		// SSE 规范：字段值只去掉一个前导空格。
		value = strings.TrimPrefix(value, " ")
		switch field {
		case "id":
			id = value
		case "event":
			event = value
		case "data":
			data = append(data, value)
		}
	}

	if len(data) == 0 {
		return nil, nil
	}
	raw := strings.Join(data, "\n")
	if !json.Valid([]byte(raw)) {
		return nil, newMeooError("malformed SSE data for event "+event, nil)
	}
	return &AgentEvent{ID: id, Event: event, Data: json.RawMessage(raw)}, nil
}

// ParseStream 解析完整响应体文本，按空行分帧，返回全部事件。
func ParseStream(body string) ([]*AgentEvent, error) {
	return ParseLines(strings.Split(body, "\n"))
}

// ParseLines 逐行解析；结尾未以空行收尾的最后一帧同样会被交付。每行末尾的 "\r" 会被去除。
func ParseLines(lines []string) ([]*AgentEvent, error) {
	var events []*AgentEvent
	var frame []string

	flush := func() error {
		if len(frame) == 0 {
			return nil
		}
		event, err := ParseFrame(strings.Join(frame, "\n"))
		frame = nil
		if err != nil {
			return err
		}
		if event != nil {
			events = append(events, event)
		}
		return nil
	}

	for _, line := range lines {
		if normalized := strings.TrimSuffix(line, "\r"); normalized == "" {
			if err := flush(); err != nil {
				return nil, err
			}
		} else {
			frame = append(frame, normalized)
		}
	}
	if err := flush(); err != nil {
		return nil, err
	}
	return events, nil
}

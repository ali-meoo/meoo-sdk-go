package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMessageSnapshotEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMessageSnapshotEventData{}

// AgentMessageSnapshotEventData struct for AgentMessageSnapshotEventData
type AgentMessageSnapshotEventData struct {
	RunId string `json:"run_id"`
	// 当前回复的全量快照；整段替换该 Run 已累计的 delta 文本。
	Content string `json:"content"`
	// 存在时是当前轮 Read/Write/Edit/Skill 及 Bash 图片/视频生成调用公开状态的权威全量快照，客户端整体替换； 缺省表示瞬态文本快照，不得据此清空已有工具状态。
	ToolCalls []AgentSnapshotToolCall `json:"tool_calls,omitempty"`
}

type _AgentMessageSnapshotEventData AgentMessageSnapshotEventData

// NewAgentMessageSnapshotEventData instantiates a new AgentMessageSnapshotEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMessageSnapshotEventData(runId string, content string) *AgentMessageSnapshotEventData {
	this := AgentMessageSnapshotEventData{}
	this.RunId = runId
	this.Content = content
	return &this
}

// NewAgentMessageSnapshotEventDataWithDefaults instantiates a new AgentMessageSnapshotEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMessageSnapshotEventDataWithDefaults() *AgentMessageSnapshotEventData {
	this := AgentMessageSnapshotEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentMessageSnapshotEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentMessageSnapshotEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentMessageSnapshotEventData) SetRunId(v string) {
	o.RunId = v
}

// GetContent returns the Content field value
func (o *AgentMessageSnapshotEventData) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AgentMessageSnapshotEventData) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AgentMessageSnapshotEventData) SetContent(v string) {
	o.Content = v
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *AgentMessageSnapshotEventData) GetToolCalls() []AgentSnapshotToolCall {
	if o == nil || IsNil(o.ToolCalls) {
		var ret []AgentSnapshotToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentMessageSnapshotEventData) GetToolCallsOk() ([]AgentSnapshotToolCall, bool) {
	if o == nil || IsNil(o.ToolCalls) {
		return nil, false
	}
	return o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *AgentMessageSnapshotEventData) HasToolCalls() bool {
	if o != nil && !IsNil(o.ToolCalls) {
		return true
	}

	return false
}

// SetToolCalls gets a reference to the given []AgentSnapshotToolCall and assigns it to the ToolCalls field.
func (o *AgentMessageSnapshotEventData) SetToolCalls(v []AgentSnapshotToolCall) {
	o.ToolCalls = v
}

func (o AgentMessageSnapshotEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMessageSnapshotEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["content"] = o.Content
	if !IsNil(o.ToolCalls) {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	return toSerialize, nil
}

func (o *AgentMessageSnapshotEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
		"content",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAgentMessageSnapshotEventData := _AgentMessageSnapshotEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentMessageSnapshotEventData)

	if err != nil {
		return err
	}

	*o = AgentMessageSnapshotEventData(varAgentMessageSnapshotEventData)

	return err
}

type NullableAgentMessageSnapshotEventData struct {
	value *AgentMessageSnapshotEventData
	isSet bool
}

func (v NullableAgentMessageSnapshotEventData) Get() *AgentMessageSnapshotEventData {
	return v.value
}

func (v *NullableAgentMessageSnapshotEventData) Set(val *AgentMessageSnapshotEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMessageSnapshotEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMessageSnapshotEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMessageSnapshotEventData(val *AgentMessageSnapshotEventData) *NullableAgentMessageSnapshotEventData {
	return &NullableAgentMessageSnapshotEventData{value: val, isSet: true}
}

func (v NullableAgentMessageSnapshotEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMessageSnapshotEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

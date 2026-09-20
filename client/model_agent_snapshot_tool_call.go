package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentSnapshotToolCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentSnapshotToolCall{}

// AgentSnapshotToolCall struct for AgentSnapshotToolCall
type AgentSnapshotToolCall struct {
	ToolCallId string `json:"tool_call_id"`
	// Bash 仅包含 meoo-cli image-generate/video-generate 调用，其他 Bash 调用不进入快照。
	ToolName string `json:"tool_name"`
	// 有配对的持久化工具结果时为 completed，否则为 started。
	Status string `json:"status"`
	// Read/Write/Edit 可携带的安全 basename；不含目录、URL 参数或读取内容。
	FileName *string             `json:"file_name,omitempty"`
	Skill    *AgentToolCallSkill `json:"skill,omitempty"`
	// 与完成态 tool.call.result 相同的生成媒体结果。
	Result *AgentMediaResult           `json:"result,omitempty"`
	Input  *AgentSnapshotToolCallInput `json:"input,omitempty"`
	// 仅 completed 可携带；由配对工具结果的状态安全归一化，不包含原始结果。
	Outcome *string `json:"outcome,omitempty"`
	// 出现时表示公开输入投影达到大小限制或未能完整解析。
	Truncated            *bool `json:"truncated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentSnapshotToolCall AgentSnapshotToolCall

// NewAgentSnapshotToolCall instantiates a new AgentSnapshotToolCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentSnapshotToolCall(toolCallId string, toolName string, status string) *AgentSnapshotToolCall {
	this := AgentSnapshotToolCall{}
	this.ToolCallId = toolCallId
	this.ToolName = toolName
	this.Status = status
	return &this
}

// NewAgentSnapshotToolCallWithDefaults instantiates a new AgentSnapshotToolCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentSnapshotToolCallWithDefaults() *AgentSnapshotToolCall {
	this := AgentSnapshotToolCall{}
	return &this
}

// GetToolCallId returns the ToolCallId field value
func (o *AgentSnapshotToolCall) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value
func (o *AgentSnapshotToolCall) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetToolName returns the ToolName field value
func (o *AgentSnapshotToolCall) GetToolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolName, true
}

// SetToolName sets field value
func (o *AgentSnapshotToolCall) SetToolName(v string) {
	o.ToolName = v
}

// GetStatus returns the Status field value
func (o *AgentSnapshotToolCall) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentSnapshotToolCall) SetStatus(v string) {
	o.Status = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AgentSnapshotToolCall) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AgentSnapshotToolCall) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AgentSnapshotToolCall) SetFileName(v string) {
	o.FileName = &v
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *AgentSnapshotToolCall) GetSkill() AgentToolCallSkill {
	if o == nil || IsNil(o.Skill) {
		var ret AgentToolCallSkill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetSkillOk() (*AgentToolCallSkill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *AgentSnapshotToolCall) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given AgentToolCallSkill and assigns it to the Skill field.
func (o *AgentSnapshotToolCall) SetSkill(v AgentToolCallSkill) {
	o.Skill = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AgentSnapshotToolCall) GetResult() AgentMediaResult {
	if o == nil || IsNil(o.Result) {
		var ret AgentMediaResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetResultOk() (*AgentMediaResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AgentSnapshotToolCall) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AgentMediaResult and assigns it to the Result field.
func (o *AgentSnapshotToolCall) SetResult(v AgentMediaResult) {
	o.Result = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *AgentSnapshotToolCall) GetInput() AgentSnapshotToolCallInput {
	if o == nil || IsNil(o.Input) {
		var ret AgentSnapshotToolCallInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetInputOk() (*AgentSnapshotToolCallInput, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *AgentSnapshotToolCall) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given AgentSnapshotToolCallInput and assigns it to the Input field.
func (o *AgentSnapshotToolCall) SetInput(v AgentSnapshotToolCallInput) {
	o.Input = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *AgentSnapshotToolCall) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *AgentSnapshotToolCall) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *AgentSnapshotToolCall) SetOutcome(v string) {
	o.Outcome = &v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *AgentSnapshotToolCall) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCall) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *AgentSnapshotToolCall) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *AgentSnapshotToolCall) SetTruncated(v bool) {
	o.Truncated = &v
}

func (o AgentSnapshotToolCall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentSnapshotToolCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tool_call_id"] = o.ToolCallId
	toSerialize["tool_name"] = o.ToolName
	toSerialize["status"] = o.Status
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.Truncated) {
		toSerialize["truncated"] = o.Truncated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentSnapshotToolCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tool_call_id",
		"tool_name",
		"status",
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

	varAgentSnapshotToolCall := _AgentSnapshotToolCall{}

	err = json.Unmarshal(data, &varAgentSnapshotToolCall)

	if err != nil {
		return err
	}

	*o = AgentSnapshotToolCall(varAgentSnapshotToolCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tool_call_id")
		delete(additionalProperties, "tool_name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "file_name")
		delete(additionalProperties, "skill")
		delete(additionalProperties, "result")
		delete(additionalProperties, "input")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "truncated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentSnapshotToolCall struct {
	value *AgentSnapshotToolCall
	isSet bool
}

func (v NullableAgentSnapshotToolCall) Get() *AgentSnapshotToolCall {
	return v.value
}

func (v *NullableAgentSnapshotToolCall) Set(val *AgentSnapshotToolCall) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSnapshotToolCall) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSnapshotToolCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSnapshotToolCall(val *AgentSnapshotToolCall) *NullableAgentSnapshotToolCall {
	return &NullableAgentSnapshotToolCall{value: val, isSet: true}
}

func (v NullableAgentSnapshotToolCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSnapshotToolCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

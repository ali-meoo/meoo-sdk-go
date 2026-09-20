package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentToolInputSnapshotEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolInputSnapshotEventData{}

// AgentToolInputSnapshotEventData struct for AgentToolInputSnapshotEventData
type AgentToolInputSnapshotEventData struct {
	RunId      string `json:"run_id"`
	ToolCallId string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// 从内部路径派生的安全 basename；可能在无法安全解析时省略。
	FileName *string                              `json:"file_name,omitempty"`
	Input    AgentToolInputSnapshotEventDataInput `json:"input"`
	// 参数是否生成完成；true 不表示工具执行成功，实际结果见 tool.call.outcome。
	Complete bool `json:"complete"`
	// 出现时表示公开投影达到大小或并发限制，内部工具仍按完整参数执行。
	Truncated *bool `json:"truncated,omitempty"`
}

type _AgentToolInputSnapshotEventData AgentToolInputSnapshotEventData

// NewAgentToolInputSnapshotEventData instantiates a new AgentToolInputSnapshotEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolInputSnapshotEventData(runId string, toolCallId string, toolName string, input AgentToolInputSnapshotEventDataInput, complete bool) *AgentToolInputSnapshotEventData {
	this := AgentToolInputSnapshotEventData{}
	this.RunId = runId
	this.ToolCallId = toolCallId
	this.ToolName = toolName
	this.Input = input
	this.Complete = complete
	return &this
}

// NewAgentToolInputSnapshotEventDataWithDefaults instantiates a new AgentToolInputSnapshotEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolInputSnapshotEventDataWithDefaults() *AgentToolInputSnapshotEventData {
	this := AgentToolInputSnapshotEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentToolInputSnapshotEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentToolInputSnapshotEventData) SetRunId(v string) {
	o.RunId = v
}

// GetToolCallId returns the ToolCallId field value
func (o *AgentToolInputSnapshotEventData) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value
func (o *AgentToolInputSnapshotEventData) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetToolName returns the ToolName field value
func (o *AgentToolInputSnapshotEventData) GetToolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolName, true
}

// SetToolName sets field value
func (o *AgentToolInputSnapshotEventData) SetToolName(v string) {
	o.ToolName = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AgentToolInputSnapshotEventData) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AgentToolInputSnapshotEventData) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AgentToolInputSnapshotEventData) SetFileName(v string) {
	o.FileName = &v
}

// GetInput returns the Input field value
func (o *AgentToolInputSnapshotEventData) GetInput() AgentToolInputSnapshotEventDataInput {
	if o == nil {
		var ret AgentToolInputSnapshotEventDataInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetInputOk() (*AgentToolInputSnapshotEventDataInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *AgentToolInputSnapshotEventData) SetInput(v AgentToolInputSnapshotEventDataInput) {
	o.Input = v
}

// GetComplete returns the Complete field value
func (o *AgentToolInputSnapshotEventData) GetComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetCompleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Complete, true
}

// SetComplete sets field value
func (o *AgentToolInputSnapshotEventData) SetComplete(v bool) {
	o.Complete = v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *AgentToolInputSnapshotEventData) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventData) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *AgentToolInputSnapshotEventData) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *AgentToolInputSnapshotEventData) SetTruncated(v bool) {
	o.Truncated = &v
}

func (o AgentToolInputSnapshotEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolInputSnapshotEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["tool_call_id"] = o.ToolCallId
	toSerialize["tool_name"] = o.ToolName
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	toSerialize["input"] = o.Input
	toSerialize["complete"] = o.Complete
	if !IsNil(o.Truncated) {
		toSerialize["truncated"] = o.Truncated
	}
	return toSerialize, nil
}

func (o *AgentToolInputSnapshotEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
		"tool_call_id",
		"tool_name",
		"input",
		"complete",
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

	varAgentToolInputSnapshotEventData := _AgentToolInputSnapshotEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentToolInputSnapshotEventData)

	if err != nil {
		return err
	}

	*o = AgentToolInputSnapshotEventData(varAgentToolInputSnapshotEventData)

	return err
}

type NullableAgentToolInputSnapshotEventData struct {
	value *AgentToolInputSnapshotEventData
	isSet bool
}

func (v NullableAgentToolInputSnapshotEventData) Get() *AgentToolInputSnapshotEventData {
	return v.value
}

func (v *NullableAgentToolInputSnapshotEventData) Set(val *AgentToolInputSnapshotEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolInputSnapshotEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolInputSnapshotEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolInputSnapshotEventData(val *AgentToolInputSnapshotEventData) *NullableAgentToolInputSnapshotEventData {
	return &NullableAgentToolInputSnapshotEventData{value: val, isSet: true}
}

func (v NullableAgentToolInputSnapshotEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolInputSnapshotEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

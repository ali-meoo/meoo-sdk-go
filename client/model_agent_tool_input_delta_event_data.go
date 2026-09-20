package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentToolInputDeltaEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolInputDeltaEventData{}

// AgentToolInputDeltaEventData struct for AgentToolInputDeltaEventData
type AgentToolInputDeltaEventData struct {
	RunId      string `json:"run_id"`
	ToolCallId string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// Write 仅使用 content；Edit 仅使用 old_string 或 new_string。
	Field string `json:"field"`
	// 追加到同一 tool_call_id、同一 field 当前值末尾的文本增量。
	Delta string `json:"delta"`
}

type _AgentToolInputDeltaEventData AgentToolInputDeltaEventData

// NewAgentToolInputDeltaEventData instantiates a new AgentToolInputDeltaEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolInputDeltaEventData(runId string, toolCallId string, toolName string, field string, delta string) *AgentToolInputDeltaEventData {
	this := AgentToolInputDeltaEventData{}
	this.RunId = runId
	this.ToolCallId = toolCallId
	this.ToolName = toolName
	this.Field = field
	this.Delta = delta
	return &this
}

// NewAgentToolInputDeltaEventDataWithDefaults instantiates a new AgentToolInputDeltaEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolInputDeltaEventDataWithDefaults() *AgentToolInputDeltaEventData {
	this := AgentToolInputDeltaEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentToolInputDeltaEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentToolInputDeltaEventData) SetRunId(v string) {
	o.RunId = v
}

// GetToolCallId returns the ToolCallId field value
func (o *AgentToolInputDeltaEventData) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEventData) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value
func (o *AgentToolInputDeltaEventData) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetToolName returns the ToolName field value
func (o *AgentToolInputDeltaEventData) GetToolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEventData) GetToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolName, true
}

// SetToolName sets field value
func (o *AgentToolInputDeltaEventData) SetToolName(v string) {
	o.ToolName = v
}

// GetField returns the Field field value
func (o *AgentToolInputDeltaEventData) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEventData) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AgentToolInputDeltaEventData) SetField(v string) {
	o.Field = v
}

// GetDelta returns the Delta field value
func (o *AgentToolInputDeltaEventData) GetDelta() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEventData) GetDeltaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *AgentToolInputDeltaEventData) SetDelta(v string) {
	o.Delta = v
}

func (o AgentToolInputDeltaEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolInputDeltaEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["tool_call_id"] = o.ToolCallId
	toSerialize["tool_name"] = o.ToolName
	toSerialize["field"] = o.Field
	toSerialize["delta"] = o.Delta
	return toSerialize, nil
}

func (o *AgentToolInputDeltaEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
		"tool_call_id",
		"tool_name",
		"field",
		"delta",
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

	varAgentToolInputDeltaEventData := _AgentToolInputDeltaEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentToolInputDeltaEventData)

	if err != nil {
		return err
	}

	*o = AgentToolInputDeltaEventData(varAgentToolInputDeltaEventData)

	return err
}

type NullableAgentToolInputDeltaEventData struct {
	value *AgentToolInputDeltaEventData
	isSet bool
}

func (v NullableAgentToolInputDeltaEventData) Get() *AgentToolInputDeltaEventData {
	return v.value
}

func (v *NullableAgentToolInputDeltaEventData) Set(val *AgentToolInputDeltaEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolInputDeltaEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolInputDeltaEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolInputDeltaEventData(val *AgentToolInputDeltaEventData) *NullableAgentToolInputDeltaEventData {
	return &NullableAgentToolInputDeltaEventData{value: val, isSet: true}
}

func (v NullableAgentToolInputDeltaEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolInputDeltaEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

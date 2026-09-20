package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunWorkingEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunWorkingEventData{}

// AgentRunWorkingEventData Run 正在处理；公开进度阶段与文案仅由 tool.call 事件承载。
type AgentRunWorkingEventData struct {
	RunId string `json:"run_id"`
	// Run 状态；客户端遇到未知取值时一律视为非终态。
	Status string `json:"status"`
}

type _AgentRunWorkingEventData AgentRunWorkingEventData

// NewAgentRunWorkingEventData instantiates a new AgentRunWorkingEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunWorkingEventData(runId string, status string) *AgentRunWorkingEventData {
	this := AgentRunWorkingEventData{}
	this.RunId = runId
	this.Status = status
	return &this
}

// NewAgentRunWorkingEventDataWithDefaults instantiates a new AgentRunWorkingEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunWorkingEventDataWithDefaults() *AgentRunWorkingEventData {
	this := AgentRunWorkingEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentRunWorkingEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentRunWorkingEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentRunWorkingEventData) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value
func (o *AgentRunWorkingEventData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentRunWorkingEventData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentRunWorkingEventData) SetStatus(v string) {
	o.Status = v
}

func (o AgentRunWorkingEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunWorkingEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *AgentRunWorkingEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
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

	varAgentRunWorkingEventData := _AgentRunWorkingEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRunWorkingEventData)

	if err != nil {
		return err
	}

	*o = AgentRunWorkingEventData(varAgentRunWorkingEventData)

	return err
}

type NullableAgentRunWorkingEventData struct {
	value *AgentRunWorkingEventData
	isSet bool
}

func (v NullableAgentRunWorkingEventData) Get() *AgentRunWorkingEventData {
	return v.value
}

func (v *NullableAgentRunWorkingEventData) Set(val *AgentRunWorkingEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunWorkingEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunWorkingEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunWorkingEventData(val *AgentRunWorkingEventData) *NullableAgentRunWorkingEventData {
	return &NullableAgentRunWorkingEventData{value: val, isSet: true}
}

func (v NullableAgentRunWorkingEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunWorkingEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

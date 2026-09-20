package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunTerminalEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunTerminalEventData{}

// AgentRunTerminalEventData struct for AgentRunTerminalEventData
type AgentRunTerminalEventData struct {
	RunId string `json:"run_id"`
	// Run 状态；客户端遇到未知取值时一律视为非终态。
	Status string `json:"status"`
}

type _AgentRunTerminalEventData AgentRunTerminalEventData

// NewAgentRunTerminalEventData instantiates a new AgentRunTerminalEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunTerminalEventData(runId string, status string) *AgentRunTerminalEventData {
	this := AgentRunTerminalEventData{}
	this.RunId = runId
	this.Status = status
	return &this
}

// NewAgentRunTerminalEventDataWithDefaults instantiates a new AgentRunTerminalEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunTerminalEventDataWithDefaults() *AgentRunTerminalEventData {
	this := AgentRunTerminalEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentRunTerminalEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentRunTerminalEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentRunTerminalEventData) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value
func (o *AgentRunTerminalEventData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentRunTerminalEventData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentRunTerminalEventData) SetStatus(v string) {
	o.Status = v
}

func (o AgentRunTerminalEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunTerminalEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *AgentRunTerminalEventData) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunTerminalEventData := _AgentRunTerminalEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentRunTerminalEventData)

	if err != nil {
		return err
	}

	*o = AgentRunTerminalEventData(varAgentRunTerminalEventData)

	return err
}

type NullableAgentRunTerminalEventData struct {
	value *AgentRunTerminalEventData
	isSet bool
}

func (v NullableAgentRunTerminalEventData) Get() *AgentRunTerminalEventData {
	return v.value
}

func (v *NullableAgentRunTerminalEventData) Set(val *AgentRunTerminalEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunTerminalEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunTerminalEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunTerminalEventData(val *AgentRunTerminalEventData) *NullableAgentRunTerminalEventData {
	return &NullableAgentRunTerminalEventData{value: val, isSet: true}
}

func (v NullableAgentRunTerminalEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunTerminalEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

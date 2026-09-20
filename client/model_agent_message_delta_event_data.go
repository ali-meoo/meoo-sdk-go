package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMessageDeltaEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMessageDeltaEventData{}

// AgentMessageDeltaEventData struct for AgentMessageDeltaEventData
type AgentMessageDeltaEventData struct {
	RunId string `json:"run_id"`
	// 追加到当前回复末尾的文本增量。
	Delta string `json:"delta"`
}

type _AgentMessageDeltaEventData AgentMessageDeltaEventData

// NewAgentMessageDeltaEventData instantiates a new AgentMessageDeltaEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMessageDeltaEventData(runId string, delta string) *AgentMessageDeltaEventData {
	this := AgentMessageDeltaEventData{}
	this.RunId = runId
	this.Delta = delta
	return &this
}

// NewAgentMessageDeltaEventDataWithDefaults instantiates a new AgentMessageDeltaEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMessageDeltaEventDataWithDefaults() *AgentMessageDeltaEventData {
	this := AgentMessageDeltaEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentMessageDeltaEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentMessageDeltaEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentMessageDeltaEventData) SetRunId(v string) {
	o.RunId = v
}

// GetDelta returns the Delta field value
func (o *AgentMessageDeltaEventData) GetDelta() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *AgentMessageDeltaEventData) GetDeltaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *AgentMessageDeltaEventData) SetDelta(v string) {
	o.Delta = v
}

func (o AgentMessageDeltaEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMessageDeltaEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["delta"] = o.Delta
	return toSerialize, nil
}

func (o *AgentMessageDeltaEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
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

	varAgentMessageDeltaEventData := _AgentMessageDeltaEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentMessageDeltaEventData)

	if err != nil {
		return err
	}

	*o = AgentMessageDeltaEventData(varAgentMessageDeltaEventData)

	return err
}

type NullableAgentMessageDeltaEventData struct {
	value *AgentMessageDeltaEventData
	isSet bool
}

func (v NullableAgentMessageDeltaEventData) Get() *AgentMessageDeltaEventData {
	return v.value
}

func (v *NullableAgentMessageDeltaEventData) Set(val *AgentMessageDeltaEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMessageDeltaEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMessageDeltaEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMessageDeltaEventData(val *AgentMessageDeltaEventData) *NullableAgentMessageDeltaEventData {
	return &NullableAgentMessageDeltaEventData{value: val, isSet: true}
}

func (v NullableAgentMessageDeltaEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMessageDeltaEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

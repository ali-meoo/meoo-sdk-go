package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunSupersededEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunSupersededEventData{}

// AgentRunSupersededEventData struct for AgentRunSupersededEventData
type AgentRunSupersededEventData struct {
	RunId  string `json:"run_id"`
	Status string `json:"status"`
	// 接管的 Run ID；不改变 status 所表达的旧 Run 真实终态，无法确定时省略。
	CurrentRunId *string `json:"current_run_id,omitempty"`
}

type _AgentRunSupersededEventData AgentRunSupersededEventData

// NewAgentRunSupersededEventData instantiates a new AgentRunSupersededEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunSupersededEventData(runId string, status string) *AgentRunSupersededEventData {
	this := AgentRunSupersededEventData{}
	this.RunId = runId
	this.Status = status
	return &this
}

// NewAgentRunSupersededEventDataWithDefaults instantiates a new AgentRunSupersededEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunSupersededEventDataWithDefaults() *AgentRunSupersededEventData {
	this := AgentRunSupersededEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentRunSupersededEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentRunSupersededEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentRunSupersededEventData) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value
func (o *AgentRunSupersededEventData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentRunSupersededEventData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentRunSupersededEventData) SetStatus(v string) {
	o.Status = v
}

// GetCurrentRunId returns the CurrentRunId field value if set, zero value otherwise.
func (o *AgentRunSupersededEventData) GetCurrentRunId() string {
	if o == nil || IsNil(o.CurrentRunId) {
		var ret string
		return ret
	}
	return *o.CurrentRunId
}

// GetCurrentRunIdOk returns a tuple with the CurrentRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunSupersededEventData) GetCurrentRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentRunId) {
		return nil, false
	}
	return o.CurrentRunId, true
}

// HasCurrentRunId returns a boolean if a field has been set.
func (o *AgentRunSupersededEventData) HasCurrentRunId() bool {
	if o != nil && !IsNil(o.CurrentRunId) {
		return true
	}

	return false
}

// SetCurrentRunId gets a reference to the given string and assigns it to the CurrentRunId field.
func (o *AgentRunSupersededEventData) SetCurrentRunId(v string) {
	o.CurrentRunId = &v
}

func (o AgentRunSupersededEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunSupersededEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["status"] = o.Status
	if !IsNil(o.CurrentRunId) {
		toSerialize["current_run_id"] = o.CurrentRunId
	}
	return toSerialize, nil
}

func (o *AgentRunSupersededEventData) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunSupersededEventData := _AgentRunSupersededEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentRunSupersededEventData)

	if err != nil {
		return err
	}

	*o = AgentRunSupersededEventData(varAgentRunSupersededEventData)

	return err
}

type NullableAgentRunSupersededEventData struct {
	value *AgentRunSupersededEventData
	isSet bool
}

func (v NullableAgentRunSupersededEventData) Get() *AgentRunSupersededEventData {
	return v.value
}

func (v *NullableAgentRunSupersededEventData) Set(val *AgentRunSupersededEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunSupersededEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunSupersededEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunSupersededEventData(val *AgentRunSupersededEventData) *NullableAgentRunSupersededEventData {
	return &NullableAgentRunSupersededEventData{value: val, isSet: true}
}

func (v NullableAgentRunSupersededEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunSupersededEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

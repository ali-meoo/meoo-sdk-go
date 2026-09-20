package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunInputRequiredEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunInputRequiredEventData{}

// AgentRunInputRequiredEventData struct for AgentRunInputRequiredEventData
type AgentRunInputRequiredEventData struct {
	RunId  string         `json:"run_id"`
	Status AgentRunStatus `json:"status"`
	Action *AgentAction   `json:"action,omitempty"`
	// 无 action 时的原因：当前阻塞工具不支持开放端回复，需在 Meoo Web 处理。
	Detail *string `json:"detail,omitempty"`
}

type _AgentRunInputRequiredEventData AgentRunInputRequiredEventData

// NewAgentRunInputRequiredEventData instantiates a new AgentRunInputRequiredEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunInputRequiredEventData(runId string, status AgentRunStatus) *AgentRunInputRequiredEventData {
	this := AgentRunInputRequiredEventData{}
	this.RunId = runId
	this.Status = status
	return &this
}

// NewAgentRunInputRequiredEventDataWithDefaults instantiates a new AgentRunInputRequiredEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunInputRequiredEventDataWithDefaults() *AgentRunInputRequiredEventData {
	this := AgentRunInputRequiredEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentRunInputRequiredEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentRunInputRequiredEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentRunInputRequiredEventData) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value
func (o *AgentRunInputRequiredEventData) GetStatus() AgentRunStatus {
	if o == nil {
		var ret AgentRunStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentRunInputRequiredEventData) GetStatusOk() (*AgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentRunInputRequiredEventData) SetStatus(v AgentRunStatus) {
	o.Status = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AgentRunInputRequiredEventData) GetAction() AgentAction {
	if o == nil || IsNil(o.Action) {
		var ret AgentAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunInputRequiredEventData) GetActionOk() (*AgentAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AgentRunInputRequiredEventData) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given AgentAction and assigns it to the Action field.
func (o *AgentRunInputRequiredEventData) SetAction(v AgentAction) {
	o.Action = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AgentRunInputRequiredEventData) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunInputRequiredEventData) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AgentRunInputRequiredEventData) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *AgentRunInputRequiredEventData) SetDetail(v string) {
	o.Detail = &v
}

func (o AgentRunInputRequiredEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunInputRequiredEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["status"] = o.Status
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

func (o *AgentRunInputRequiredEventData) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunInputRequiredEventData := _AgentRunInputRequiredEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentRunInputRequiredEventData)

	if err != nil {
		return err
	}

	*o = AgentRunInputRequiredEventData(varAgentRunInputRequiredEventData)

	return err
}

type NullableAgentRunInputRequiredEventData struct {
	value *AgentRunInputRequiredEventData
	isSet bool
}

func (v NullableAgentRunInputRequiredEventData) Get() *AgentRunInputRequiredEventData {
	return v.value
}

func (v *NullableAgentRunInputRequiredEventData) Set(val *AgentRunInputRequiredEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunInputRequiredEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunInputRequiredEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunInputRequiredEventData(val *AgentRunInputRequiredEventData) *NullableAgentRunInputRequiredEventData {
	return &NullableAgentRunInputRequiredEventData{value: val, isSet: true}
}

func (v NullableAgentRunInputRequiredEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunInputRequiredEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

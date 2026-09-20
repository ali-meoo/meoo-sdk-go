package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunUsageEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunUsageEventData{}

// AgentRunUsageEventData struct for AgentRunUsageEventData
type AgentRunUsageEventData struct {
	RunId string `json:"run_id"`
	// 单次结算的稳定标识；断线重连重放时保持不变，客户端据此去重。
	UsageId string `json:"usage_id" validate:"regexp=^usage_[a-f0-9]{32}$"`
	// 本次结算消耗的积分，以非负十进制整数字符串表示；不是 Run 累计值。
	CreditsUsed string `json:"credits_used" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _AgentRunUsageEventData AgentRunUsageEventData

// NewAgentRunUsageEventData instantiates a new AgentRunUsageEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunUsageEventData(runId string, usageId string, creditsUsed string) *AgentRunUsageEventData {
	this := AgentRunUsageEventData{}
	this.RunId = runId
	this.UsageId = usageId
	this.CreditsUsed = creditsUsed
	return &this
}

// NewAgentRunUsageEventDataWithDefaults instantiates a new AgentRunUsageEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunUsageEventDataWithDefaults() *AgentRunUsageEventData {
	this := AgentRunUsageEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentRunUsageEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentRunUsageEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentRunUsageEventData) SetRunId(v string) {
	o.RunId = v
}

// GetUsageId returns the UsageId field value
func (o *AgentRunUsageEventData) GetUsageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageId
}

// GetUsageIdOk returns a tuple with the UsageId field value
// and a boolean to check if the value has been set.
func (o *AgentRunUsageEventData) GetUsageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageId, true
}

// SetUsageId sets field value
func (o *AgentRunUsageEventData) SetUsageId(v string) {
	o.UsageId = v
}

// GetCreditsUsed returns the CreditsUsed field value
func (o *AgentRunUsageEventData) GetCreditsUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreditsUsed
}

// GetCreditsUsedOk returns a tuple with the CreditsUsed field value
// and a boolean to check if the value has been set.
func (o *AgentRunUsageEventData) GetCreditsUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsUsed, true
}

// SetCreditsUsed sets field value
func (o *AgentRunUsageEventData) SetCreditsUsed(v string) {
	o.CreditsUsed = v
}

func (o AgentRunUsageEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunUsageEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["usage_id"] = o.UsageId
	toSerialize["credits_used"] = o.CreditsUsed
	return toSerialize, nil
}

func (o *AgentRunUsageEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
		"usage_id",
		"credits_used",
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

	varAgentRunUsageEventData := _AgentRunUsageEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRunUsageEventData)

	if err != nil {
		return err
	}

	*o = AgentRunUsageEventData(varAgentRunUsageEventData)

	return err
}

type NullableAgentRunUsageEventData struct {
	value *AgentRunUsageEventData
	isSet bool
}

func (v NullableAgentRunUsageEventData) Get() *AgentRunUsageEventData {
	return v.value
}

func (v *NullableAgentRunUsageEventData) Set(val *AgentRunUsageEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunUsageEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunUsageEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunUsageEventData(val *AgentRunUsageEventData) *NullableAgentRunUsageEventData {
	return &NullableAgentRunUsageEventData{value: val, isSet: true}
}

func (v NullableAgentRunUsageEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunUsageEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

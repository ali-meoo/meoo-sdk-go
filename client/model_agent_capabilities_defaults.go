package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentCapabilitiesDefaults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentCapabilitiesDefaults{}

// AgentCapabilitiesDefaults struct for AgentCapabilitiesDefaults
type AgentCapabilitiesDefaults struct {
	// 每次 Run 没有显式 model 时使用的当前默认模型 ID。
	Model     string `json:"model"`
	SpeedTier string `json:"speed_tier"`
}

type _AgentCapabilitiesDefaults AgentCapabilitiesDefaults

// NewAgentCapabilitiesDefaults instantiates a new AgentCapabilitiesDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentCapabilitiesDefaults(model string, speedTier string) *AgentCapabilitiesDefaults {
	this := AgentCapabilitiesDefaults{}
	this.Model = model
	this.SpeedTier = speedTier
	return &this
}

// NewAgentCapabilitiesDefaultsWithDefaults instantiates a new AgentCapabilitiesDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentCapabilitiesDefaultsWithDefaults() *AgentCapabilitiesDefaults {
	this := AgentCapabilitiesDefaults{}
	return &this
}

// GetModel returns the Model field value
func (o *AgentCapabilitiesDefaults) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilitiesDefaults) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *AgentCapabilitiesDefaults) SetModel(v string) {
	o.Model = v
}

// GetSpeedTier returns the SpeedTier field value
func (o *AgentCapabilitiesDefaults) GetSpeedTier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpeedTier
}

// GetSpeedTierOk returns a tuple with the SpeedTier field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilitiesDefaults) GetSpeedTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpeedTier, true
}

// SetSpeedTier sets field value
func (o *AgentCapabilitiesDefaults) SetSpeedTier(v string) {
	o.SpeedTier = v
}

func (o AgentCapabilitiesDefaults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentCapabilitiesDefaults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["speed_tier"] = o.SpeedTier
	return toSerialize, nil
}

func (o *AgentCapabilitiesDefaults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"speed_tier",
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

	varAgentCapabilitiesDefaults := _AgentCapabilitiesDefaults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentCapabilitiesDefaults)

	if err != nil {
		return err
	}

	*o = AgentCapabilitiesDefaults(varAgentCapabilitiesDefaults)

	return err
}

type NullableAgentCapabilitiesDefaults struct {
	value *AgentCapabilitiesDefaults
	isSet bool
}

func (v NullableAgentCapabilitiesDefaults) Get() *AgentCapabilitiesDefaults {
	return v.value
}

func (v *NullableAgentCapabilitiesDefaults) Set(val *AgentCapabilitiesDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentCapabilitiesDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentCapabilitiesDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentCapabilitiesDefaults(val *AgentCapabilitiesDefaults) *NullableAgentCapabilitiesDefaults {
	return &NullableAgentCapabilitiesDefaults{value: val, isSet: true}
}

func (v NullableAgentCapabilitiesDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentCapabilitiesDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

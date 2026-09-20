package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentCapabilitySpeedTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentCapabilitySpeedTier{}

// AgentCapabilitySpeedTier struct for AgentCapabilitySpeedTier
type AgentCapabilitySpeedTier struct {
	Id string `json:"id"`
	// 面向用户的档位展示名称。
	Name string `json:"name"`
}

type _AgentCapabilitySpeedTier AgentCapabilitySpeedTier

// NewAgentCapabilitySpeedTier instantiates a new AgentCapabilitySpeedTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentCapabilitySpeedTier(id string, name string) *AgentCapabilitySpeedTier {
	this := AgentCapabilitySpeedTier{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAgentCapabilitySpeedTierWithDefaults instantiates a new AgentCapabilitySpeedTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentCapabilitySpeedTierWithDefaults() *AgentCapabilitySpeedTier {
	this := AgentCapabilitySpeedTier{}
	return &this
}

// GetId returns the Id field value
func (o *AgentCapabilitySpeedTier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilitySpeedTier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentCapabilitySpeedTier) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AgentCapabilitySpeedTier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilitySpeedTier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AgentCapabilitySpeedTier) SetName(v string) {
	o.Name = v
}

func (o AgentCapabilitySpeedTier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentCapabilitySpeedTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *AgentCapabilitySpeedTier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varAgentCapabilitySpeedTier := _AgentCapabilitySpeedTier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentCapabilitySpeedTier)

	if err != nil {
		return err
	}

	*o = AgentCapabilitySpeedTier(varAgentCapabilitySpeedTier)

	return err
}

type NullableAgentCapabilitySpeedTier struct {
	value *AgentCapabilitySpeedTier
	isSet bool
}

func (v NullableAgentCapabilitySpeedTier) Get() *AgentCapabilitySpeedTier {
	return v.value
}

func (v *NullableAgentCapabilitySpeedTier) Set(val *AgentCapabilitySpeedTier) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentCapabilitySpeedTier) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentCapabilitySpeedTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentCapabilitySpeedTier(val *AgentCapabilitySpeedTier) *NullableAgentCapabilitySpeedTier {
	return &NullableAgentCapabilitySpeedTier{value: val, isSet: true}
}

func (v NullableAgentCapabilitySpeedTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentCapabilitySpeedTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

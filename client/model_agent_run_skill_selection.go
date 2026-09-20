package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunSkillSelection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunSkillSelection{}

// AgentRunSkillSelection struct for AgentRunSkillSelection
type AgentRunSkillSelection struct {
	SkillId string `json:"skill_id" validate:"regexp=^[1-9][0-9]*$"`
}

type _AgentRunSkillSelection AgentRunSkillSelection

// NewAgentRunSkillSelection instantiates a new AgentRunSkillSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunSkillSelection(skillId string) *AgentRunSkillSelection {
	this := AgentRunSkillSelection{}
	this.SkillId = skillId
	return &this
}

// NewAgentRunSkillSelectionWithDefaults instantiates a new AgentRunSkillSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunSkillSelectionWithDefaults() *AgentRunSkillSelection {
	this := AgentRunSkillSelection{}
	return &this
}

// GetSkillId returns the SkillId field value
func (o *AgentRunSkillSelection) GetSkillId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkillId
}

// GetSkillIdOk returns a tuple with the SkillId field value
// and a boolean to check if the value has been set.
func (o *AgentRunSkillSelection) GetSkillIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkillId, true
}

// SetSkillId sets field value
func (o *AgentRunSkillSelection) SetSkillId(v string) {
	o.SkillId = v
}

func (o AgentRunSkillSelection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunSkillSelection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["skill_id"] = o.SkillId
	return toSerialize, nil
}

func (o *AgentRunSkillSelection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"skill_id",
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

	varAgentRunSkillSelection := _AgentRunSkillSelection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRunSkillSelection)

	if err != nil {
		return err
	}

	*o = AgentRunSkillSelection(varAgentRunSkillSelection)

	return err
}

type NullableAgentRunSkillSelection struct {
	value *AgentRunSkillSelection
	isSet bool
}

func (v NullableAgentRunSkillSelection) Get() *AgentRunSkillSelection {
	return v.value
}

func (v *NullableAgentRunSkillSelection) Set(val *AgentRunSkillSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunSkillSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunSkillSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunSkillSelection(val *AgentRunSkillSelection) *NullableAgentRunSkillSelection {
	return &NullableAgentRunSkillSelection{value: val, isSet: true}
}

func (v NullableAgentRunSkillSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunSkillSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentToolCallSkill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolCallSkill{}

// AgentToolCallSkill struct for AgentToolCallSkill
type AgentToolCallSkill struct {
	// 服务端解析后的 canonical Skill 名称。
	SkillName string `json:"skill_name"`
	// 可选用户展示名；缺失时客户端回退显示 skill_name。
	DisplayName          *string `json:"display_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentToolCallSkill AgentToolCallSkill

// NewAgentToolCallSkill instantiates a new AgentToolCallSkill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolCallSkill(skillName string) *AgentToolCallSkill {
	this := AgentToolCallSkill{}
	this.SkillName = skillName
	return &this
}

// NewAgentToolCallSkillWithDefaults instantiates a new AgentToolCallSkill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolCallSkillWithDefaults() *AgentToolCallSkill {
	this := AgentToolCallSkill{}
	return &this
}

// GetSkillName returns the SkillName field value
func (o *AgentToolCallSkill) GetSkillName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkillName
}

// GetSkillNameOk returns a tuple with the SkillName field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallSkill) GetSkillNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkillName, true
}

// SetSkillName sets field value
func (o *AgentToolCallSkill) SetSkillName(v string) {
	o.SkillName = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AgentToolCallSkill) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolCallSkill) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AgentToolCallSkill) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AgentToolCallSkill) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o AgentToolCallSkill) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolCallSkill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["skill_name"] = o.SkillName
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentToolCallSkill) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"skill_name",
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

	varAgentToolCallSkill := _AgentToolCallSkill{}

	err = json.Unmarshal(data, &varAgentToolCallSkill)

	if err != nil {
		return err
	}

	*o = AgentToolCallSkill(varAgentToolCallSkill)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "skill_name")
		delete(additionalProperties, "display_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentToolCallSkill struct {
	value *AgentToolCallSkill
	isSet bool
}

func (v NullableAgentToolCallSkill) Get() *AgentToolCallSkill {
	return v.value
}

func (v *NullableAgentToolCallSkill) Set(val *AgentToolCallSkill) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolCallSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolCallSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolCallSkill(val *AgentToolCallSkill) *NullableAgentToolCallSkill {
	return &NullableAgentToolCallSkill{value: val, isSet: true}
}

func (v NullableAgentToolCallSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolCallSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

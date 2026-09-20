package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SelectableSkill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectableSkill{}

// SelectableSkill struct for SelectableSkill
type SelectableSkill struct {
	SkillId     string `json:"skill_id" validate:"regexp=^[1-9][0-9]*$"`
	SkillName   string `json:"skill_name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Visibility  string `json:"visibility"`
}

type _SelectableSkill SelectableSkill

// NewSelectableSkill instantiates a new SelectableSkill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectableSkill(skillId string, skillName string, displayName string, description string, source string, visibility string) *SelectableSkill {
	this := SelectableSkill{}
	this.SkillId = skillId
	this.SkillName = skillName
	this.DisplayName = displayName
	this.Description = description
	this.Source = source
	this.Visibility = visibility
	return &this
}

// NewSelectableSkillWithDefaults instantiates a new SelectableSkill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectableSkillWithDefaults() *SelectableSkill {
	this := SelectableSkill{}
	return &this
}

// GetSkillId returns the SkillId field value
func (o *SelectableSkill) GetSkillId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkillId
}

// GetSkillIdOk returns a tuple with the SkillId field value
// and a boolean to check if the value has been set.
func (o *SelectableSkill) GetSkillIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkillId, true
}

// SetSkillId sets field value
func (o *SelectableSkill) SetSkillId(v string) {
	o.SkillId = v
}

// GetSkillName returns the SkillName field value
func (o *SelectableSkill) GetSkillName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkillName
}

// GetSkillNameOk returns a tuple with the SkillName field value
// and a boolean to check if the value has been set.
func (o *SelectableSkill) GetSkillNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkillName, true
}

// SetSkillName sets field value
func (o *SelectableSkill) SetSkillName(v string) {
	o.SkillName = v
}

// GetDisplayName returns the DisplayName field value
func (o *SelectableSkill) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *SelectableSkill) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *SelectableSkill) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value
func (o *SelectableSkill) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SelectableSkill) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SelectableSkill) SetDescription(v string) {
	o.Description = v
}

// GetSource returns the Source field value
func (o *SelectableSkill) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SelectableSkill) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SelectableSkill) SetSource(v string) {
	o.Source = v
}

// GetVisibility returns the Visibility field value
func (o *SelectableSkill) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *SelectableSkill) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *SelectableSkill) SetVisibility(v string) {
	o.Visibility = v
}

func (o SelectableSkill) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectableSkill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["skill_id"] = o.SkillId
	toSerialize["skill_name"] = o.SkillName
	toSerialize["display_name"] = o.DisplayName
	toSerialize["description"] = o.Description
	toSerialize["source"] = o.Source
	toSerialize["visibility"] = o.Visibility
	return toSerialize, nil
}

func (o *SelectableSkill) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"skill_id",
		"skill_name",
		"display_name",
		"description",
		"source",
		"visibility",
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

	varSelectableSkill := _SelectableSkill{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelectableSkill)

	if err != nil {
		return err
	}

	*o = SelectableSkill(varSelectableSkill)

	return err
}

type NullableSelectableSkill struct {
	value *SelectableSkill
	isSet bool
}

func (v NullableSelectableSkill) Get() *SelectableSkill {
	return v.value
}

func (v *NullableSelectableSkill) Set(val *SelectableSkill) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectableSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectableSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectableSkill(val *SelectableSkill) *NullableSelectableSkill {
	return &NullableSelectableSkill{value: val, isSet: true}
}

func (v NullableSelectableSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectableSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

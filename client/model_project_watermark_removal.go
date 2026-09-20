package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectWatermarkRemoval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectWatermarkRemoval{}

// ProjectWatermarkRemoval struct for ProjectWatermarkRemoval
type ProjectWatermarkRemoval struct {
	Enabled bool `json:"enabled"`
	// 当前项目状态是否允许修改去水印设置。
	Editable bool `json:"editable"`
}

type _ProjectWatermarkRemoval ProjectWatermarkRemoval

// NewProjectWatermarkRemoval instantiates a new ProjectWatermarkRemoval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectWatermarkRemoval(enabled bool, editable bool) *ProjectWatermarkRemoval {
	this := ProjectWatermarkRemoval{}
	this.Enabled = enabled
	this.Editable = editable
	return &this
}

// NewProjectWatermarkRemovalWithDefaults instantiates a new ProjectWatermarkRemoval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWatermarkRemovalWithDefaults() *ProjectWatermarkRemoval {
	this := ProjectWatermarkRemoval{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ProjectWatermarkRemoval) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProjectWatermarkRemoval) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProjectWatermarkRemoval) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEditable returns the Editable field value
func (o *ProjectWatermarkRemoval) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *ProjectWatermarkRemoval) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *ProjectWatermarkRemoval) SetEditable(v bool) {
	o.Editable = v
}

func (o ProjectWatermarkRemoval) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectWatermarkRemoval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["editable"] = o.Editable
	return toSerialize, nil
}

func (o *ProjectWatermarkRemoval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"editable",
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

	varProjectWatermarkRemoval := _ProjectWatermarkRemoval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectWatermarkRemoval)

	if err != nil {
		return err
	}

	*o = ProjectWatermarkRemoval(varProjectWatermarkRemoval)

	return err
}

type NullableProjectWatermarkRemoval struct {
	value *ProjectWatermarkRemoval
	isSet bool
}

func (v NullableProjectWatermarkRemoval) Get() *ProjectWatermarkRemoval {
	return v.value
}

func (v *NullableProjectWatermarkRemoval) Set(val *ProjectWatermarkRemoval) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectWatermarkRemoval) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectWatermarkRemoval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectWatermarkRemoval(val *ProjectWatermarkRemoval) *NullableProjectWatermarkRemoval {
	return &NullableProjectWatermarkRemoval{value: val, isSet: true}
}

func (v NullableProjectWatermarkRemoval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectWatermarkRemoval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

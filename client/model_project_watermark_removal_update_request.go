package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectWatermarkRemovalUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectWatermarkRemovalUpdateRequest{}

// ProjectWatermarkRemovalUpdateRequest struct for ProjectWatermarkRemovalUpdateRequest
type ProjectWatermarkRemovalUpdateRequest struct {
	// true 开启去水印，false 关闭去水印。
	Enabled bool `json:"enabled"`
}

type _ProjectWatermarkRemovalUpdateRequest ProjectWatermarkRemovalUpdateRequest

// NewProjectWatermarkRemovalUpdateRequest instantiates a new ProjectWatermarkRemovalUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectWatermarkRemovalUpdateRequest(enabled bool) *ProjectWatermarkRemovalUpdateRequest {
	this := ProjectWatermarkRemovalUpdateRequest{}
	this.Enabled = enabled
	return &this
}

// NewProjectWatermarkRemovalUpdateRequestWithDefaults instantiates a new ProjectWatermarkRemovalUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWatermarkRemovalUpdateRequestWithDefaults() *ProjectWatermarkRemovalUpdateRequest {
	this := ProjectWatermarkRemovalUpdateRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ProjectWatermarkRemovalUpdateRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProjectWatermarkRemovalUpdateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProjectWatermarkRemovalUpdateRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ProjectWatermarkRemovalUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectWatermarkRemovalUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *ProjectWatermarkRemovalUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varProjectWatermarkRemovalUpdateRequest := _ProjectWatermarkRemovalUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectWatermarkRemovalUpdateRequest)

	if err != nil {
		return err
	}

	*o = ProjectWatermarkRemovalUpdateRequest(varProjectWatermarkRemovalUpdateRequest)

	return err
}

type NullableProjectWatermarkRemovalUpdateRequest struct {
	value *ProjectWatermarkRemovalUpdateRequest
	isSet bool
}

func (v NullableProjectWatermarkRemovalUpdateRequest) Get() *ProjectWatermarkRemovalUpdateRequest {
	return v.value
}

func (v *NullableProjectWatermarkRemovalUpdateRequest) Set(val *ProjectWatermarkRemovalUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectWatermarkRemovalUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectWatermarkRemovalUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectWatermarkRemovalUpdateRequest(val *ProjectWatermarkRemovalUpdateRequest) *NullableProjectWatermarkRemovalUpdateRequest {
	return &NullableProjectWatermarkRemovalUpdateRequest{value: val, isSet: true}
}

func (v NullableProjectWatermarkRemovalUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectWatermarkRemovalUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

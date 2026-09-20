package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamDebugContextManagementScopesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamDebugContextManagementScopesInner{}

// TeamDebugContextManagementScopesInner struct for TeamDebugContextManagementScopesInner
type TeamDebugContextManagementScopesInner struct {
	Value           string `json:"value"`
	Label           string `json:"label"`
	Description     string `json:"description"`
	DefaultSelected bool   `json:"default_selected"`
}

type _TeamDebugContextManagementScopesInner TeamDebugContextManagementScopesInner

// NewTeamDebugContextManagementScopesInner instantiates a new TeamDebugContextManagementScopesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamDebugContextManagementScopesInner(value string, label string, description string, defaultSelected bool) *TeamDebugContextManagementScopesInner {
	this := TeamDebugContextManagementScopesInner{}
	this.Value = value
	this.Label = label
	this.Description = description
	this.DefaultSelected = defaultSelected
	return &this
}

// NewTeamDebugContextManagementScopesInnerWithDefaults instantiates a new TeamDebugContextManagementScopesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamDebugContextManagementScopesInnerWithDefaults() *TeamDebugContextManagementScopesInner {
	this := TeamDebugContextManagementScopesInner{}
	return &this
}

// GetValue returns the Value field value
func (o *TeamDebugContextManagementScopesInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContextManagementScopesInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TeamDebugContextManagementScopesInner) SetValue(v string) {
	o.Value = v
}

// GetLabel returns the Label field value
func (o *TeamDebugContextManagementScopesInner) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContextManagementScopesInner) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TeamDebugContextManagementScopesInner) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value
func (o *TeamDebugContextManagementScopesInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContextManagementScopesInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TeamDebugContextManagementScopesInner) SetDescription(v string) {
	o.Description = v
}

// GetDefaultSelected returns the DefaultSelected field value
func (o *TeamDebugContextManagementScopesInner) GetDefaultSelected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultSelected
}

// GetDefaultSelectedOk returns a tuple with the DefaultSelected field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContextManagementScopesInner) GetDefaultSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSelected, true
}

// SetDefaultSelected sets field value
func (o *TeamDebugContextManagementScopesInner) SetDefaultSelected(v bool) {
	o.DefaultSelected = v
}

func (o TeamDebugContextManagementScopesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamDebugContextManagementScopesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["label"] = o.Label
	toSerialize["description"] = o.Description
	toSerialize["default_selected"] = o.DefaultSelected
	return toSerialize, nil
}

func (o *TeamDebugContextManagementScopesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"label",
		"description",
		"default_selected",
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

	varTeamDebugContextManagementScopesInner := _TeamDebugContextManagementScopesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamDebugContextManagementScopesInner)

	if err != nil {
		return err
	}

	*o = TeamDebugContextManagementScopesInner(varTeamDebugContextManagementScopesInner)

	return err
}

type NullableTeamDebugContextManagementScopesInner struct {
	value *TeamDebugContextManagementScopesInner
	isSet bool
}

func (v NullableTeamDebugContextManagementScopesInner) Get() *TeamDebugContextManagementScopesInner {
	return v.value
}

func (v *NullableTeamDebugContextManagementScopesInner) Set(val *TeamDebugContextManagementScopesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamDebugContextManagementScopesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamDebugContextManagementScopesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamDebugContextManagementScopesInner(val *TeamDebugContextManagementScopesInner) *NullableTeamDebugContextManagementScopesInner {
	return &NullableTeamDebugContextManagementScopesInner{value: val, isSet: true}
}

func (v NullableTeamDebugContextManagementScopesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamDebugContextManagementScopesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

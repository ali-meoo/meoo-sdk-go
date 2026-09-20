package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamDebugContextTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamDebugContextTeam{}

// TeamDebugContextTeam struct for TeamDebugContextTeam
type TeamDebugContextTeam struct {
	Name    string `json:"name"`
	Edition string `json:"edition"`
}

type _TeamDebugContextTeam TeamDebugContextTeam

// NewTeamDebugContextTeam instantiates a new TeamDebugContextTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamDebugContextTeam(name string, edition string) *TeamDebugContextTeam {
	this := TeamDebugContextTeam{}
	this.Name = name
	this.Edition = edition
	return &this
}

// NewTeamDebugContextTeamWithDefaults instantiates a new TeamDebugContextTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamDebugContextTeamWithDefaults() *TeamDebugContextTeam {
	this := TeamDebugContextTeam{}
	return &this
}

// GetName returns the Name field value
func (o *TeamDebugContextTeam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContextTeam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamDebugContextTeam) SetName(v string) {
	o.Name = v
}

// GetEdition returns the Edition field value
func (o *TeamDebugContextTeam) GetEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edition
}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContextTeam) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edition, true
}

// SetEdition sets field value
func (o *TeamDebugContextTeam) SetEdition(v string) {
	o.Edition = v
}

func (o TeamDebugContextTeam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamDebugContextTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["edition"] = o.Edition
	return toSerialize, nil
}

func (o *TeamDebugContextTeam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"edition",
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

	varTeamDebugContextTeam := _TeamDebugContextTeam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamDebugContextTeam)

	if err != nil {
		return err
	}

	*o = TeamDebugContextTeam(varTeamDebugContextTeam)

	return err
}

type NullableTeamDebugContextTeam struct {
	value *TeamDebugContextTeam
	isSet bool
}

func (v NullableTeamDebugContextTeam) Get() *TeamDebugContextTeam {
	return v.value
}

func (v *NullableTeamDebugContextTeam) Set(val *TeamDebugContextTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamDebugContextTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamDebugContextTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamDebugContextTeam(val *TeamDebugContextTeam) *NullableTeamDebugContextTeam {
	return &NullableTeamDebugContextTeam{value: val, isSet: true}
}

func (v NullableTeamDebugContextTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamDebugContextTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

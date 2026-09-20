package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamMemberUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMemberUpdateRequest{}

// TeamMemberUpdateRequest struct for TeamMemberUpdateRequest
type TeamMemberUpdateRequest struct {
	Status string `json:"status"`
}

type _TeamMemberUpdateRequest TeamMemberUpdateRequest

// NewTeamMemberUpdateRequest instantiates a new TeamMemberUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberUpdateRequest(status string) *TeamMemberUpdateRequest {
	this := TeamMemberUpdateRequest{}
	this.Status = status
	return &this
}

// NewTeamMemberUpdateRequestWithDefaults instantiates a new TeamMemberUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberUpdateRequestWithDefaults() *TeamMemberUpdateRequest {
	this := TeamMemberUpdateRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *TeamMemberUpdateRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TeamMemberUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TeamMemberUpdateRequest) SetStatus(v string) {
	o.Status = v
}

func (o TeamMemberUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMemberUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *TeamMemberUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varTeamMemberUpdateRequest := _TeamMemberUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamMemberUpdateRequest)

	if err != nil {
		return err
	}

	*o = TeamMemberUpdateRequest(varTeamMemberUpdateRequest)

	return err
}

type NullableTeamMemberUpdateRequest struct {
	value *TeamMemberUpdateRequest
	isSet bool
}

func (v NullableTeamMemberUpdateRequest) Get() *TeamMemberUpdateRequest {
	return v.value
}

func (v *NullableTeamMemberUpdateRequest) Set(val *TeamMemberUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberUpdateRequest(val *TeamMemberUpdateRequest) *NullableTeamMemberUpdateRequest {
	return &NullableTeamMemberUpdateRequest{value: val, isSet: true}
}

func (v NullableTeamMemberUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

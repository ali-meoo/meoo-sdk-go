package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamMemberList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMemberList{}

// TeamMemberList struct for TeamMemberList
type TeamMemberList struct {
	Members       []TeamMember `json:"members"`
	NextPageToken *string      `json:"next_page_token,omitempty"`
}

type _TeamMemberList TeamMemberList

// NewTeamMemberList instantiates a new TeamMemberList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberList(members []TeamMember) *TeamMemberList {
	this := TeamMemberList{}
	this.Members = members
	return &this
}

// NewTeamMemberListWithDefaults instantiates a new TeamMemberList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberListWithDefaults() *TeamMemberList {
	this := TeamMemberList{}
	return &this
}

// GetMembers returns the Members field value
func (o *TeamMemberList) GetMembers() []TeamMember {
	if o == nil {
		var ret []TeamMember
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *TeamMemberList) GetMembersOk() ([]TeamMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *TeamMemberList) SetMembers(v []TeamMember) {
	o.Members = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *TeamMemberList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *TeamMemberList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *TeamMemberList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o TeamMemberList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMemberList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["members"] = o.Members
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return toSerialize, nil
}

func (o *TeamMemberList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"members",
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

	varTeamMemberList := _TeamMemberList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamMemberList)

	if err != nil {
		return err
	}

	*o = TeamMemberList(varTeamMemberList)

	return err
}

type NullableTeamMemberList struct {
	value *TeamMemberList
	isSet bool
}

func (v NullableTeamMemberList) Get() *TeamMemberList {
	return v.value
}

func (v *NullableTeamMemberList) Set(val *TeamMemberList) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberList) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberList(val *TeamMemberList) *NullableTeamMemberList {
	return &NullableTeamMemberList{value: val, isSet: true}
}

func (v NullableTeamMemberList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

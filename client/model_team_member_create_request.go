package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamMemberCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMemberCreateRequest{}

// TeamMemberCreateRequest struct for TeamMemberCreateRequest
type TeamMemberCreateRequest struct {
	// 成员昵称；先去除首尾空白，再校验长度。
	Nickname string         `json:"nickname"`
	Avatar   NullableString `json:"avatar,omitempty"`
}

type _TeamMemberCreateRequest TeamMemberCreateRequest

// NewTeamMemberCreateRequest instantiates a new TeamMemberCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberCreateRequest(nickname string) *TeamMemberCreateRequest {
	this := TeamMemberCreateRequest{}
	this.Nickname = nickname
	return &this
}

// NewTeamMemberCreateRequestWithDefaults instantiates a new TeamMemberCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberCreateRequestWithDefaults() *TeamMemberCreateRequest {
	this := TeamMemberCreateRequest{}
	return &this
}

// GetNickname returns the Nickname field value
func (o *TeamMemberCreateRequest) GetNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value
// and a boolean to check if the value has been set.
func (o *TeamMemberCreateRequest) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nickname, true
}

// SetNickname sets field value
func (o *TeamMemberCreateRequest) SetNickname(v string) {
	o.Nickname = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamMemberCreateRequest) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamMemberCreateRequest) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamMemberCreateRequest) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *TeamMemberCreateRequest) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *TeamMemberCreateRequest) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *TeamMemberCreateRequest) UnsetAvatar() {
	o.Avatar.Unset()
}

func (o TeamMemberCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMemberCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nickname"] = o.Nickname
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	return toSerialize, nil
}

func (o *TeamMemberCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nickname",
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

	varTeamMemberCreateRequest := _TeamMemberCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamMemberCreateRequest)

	if err != nil {
		return err
	}

	*o = TeamMemberCreateRequest(varTeamMemberCreateRequest)

	return err
}

type NullableTeamMemberCreateRequest struct {
	value *TeamMemberCreateRequest
	isSet bool
}

func (v NullableTeamMemberCreateRequest) Get() *TeamMemberCreateRequest {
	return v.value
}

func (v *NullableTeamMemberCreateRequest) Set(val *TeamMemberCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberCreateRequest(val *TeamMemberCreateRequest) *NullableTeamMemberCreateRequest {
	return &NullableTeamMemberCreateRequest{value: val, isSet: true}
}

func (v NullableTeamMemberCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMember{}

// TeamMember struct for TeamMember
type TeamMember struct {
	// 团队成员公开 ID。
	MemberId string         `json:"member_id"`
	Nickname string         `json:"nickname"`
	Avatar   NullableString `json:"avatar"`
	Role     string         `json:"role"`
	Status   string         `json:"status"`
	// 创建时间，Unix epoch 毫秒。
	CreatedAt int64 `json:"created_at"`
}

type _TeamMember TeamMember

// NewTeamMember instantiates a new TeamMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMember(memberId string, nickname string, avatar NullableString, role string, status string, createdAt int64) *TeamMember {
	this := TeamMember{}
	this.MemberId = memberId
	this.Nickname = nickname
	this.Avatar = avatar
	this.Role = role
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewTeamMemberWithDefaults instantiates a new TeamMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberWithDefaults() *TeamMember {
	this := TeamMember{}
	return &this
}

// GetMemberId returns the MemberId field value
func (o *TeamMember) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *TeamMember) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *TeamMember) SetMemberId(v string) {
	o.MemberId = v
}

// GetNickname returns the Nickname field value
func (o *TeamMember) GetNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value
// and a boolean to check if the value has been set.
func (o *TeamMember) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nickname, true
}

// SetNickname sets field value
func (o *TeamMember) SetNickname(v string) {
	o.Nickname = v
}

// GetAvatar returns the Avatar field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TeamMember) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}

	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamMember) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// SetAvatar sets field value
func (o *TeamMember) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// GetRole returns the Role field value
func (o *TeamMember) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *TeamMember) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *TeamMember) SetRole(v string) {
	o.Role = v
}

// GetStatus returns the Status field value
func (o *TeamMember) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TeamMember) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TeamMember) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TeamMember) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamMember) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TeamMember) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o TeamMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["member_id"] = o.MemberId
	toSerialize["nickname"] = o.Nickname
	toSerialize["avatar"] = o.Avatar.Get()
	toSerialize["role"] = o.Role
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *TeamMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"member_id",
		"nickname",
		"avatar",
		"role",
		"status",
		"created_at",
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

	varTeamMember := _TeamMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamMember)

	if err != nil {
		return err
	}

	*o = TeamMember(varTeamMember)

	return err
}

type NullableTeamMember struct {
	value *TeamMember
	isSet bool
}

func (v NullableTeamMember) Get() *TeamMember {
	return v.value
}

func (v *NullableTeamMember) Set(val *TeamMember) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMember) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMember(val *TeamMember) *NullableTeamMember {
	return &NullableTeamMember{value: val, isSet: true}
}

func (v NullableTeamMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

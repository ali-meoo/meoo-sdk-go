package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamMemberTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMemberTokenRequest{}

// TeamMemberTokenRequest struct for TeamMemberTokenRequest
type TeamMemberTokenRequest struct {
	// 必须显式指定且属于当前 AK 可委托范围的 Scope。
	Scopes []string `json:"scopes"`
	// 有效期秒数；默认 1800，实际不超过当前 AK 的 TTL 上限。
	ExpiresIn *int32 `json:"expires_in,omitempty"`
}

type _TeamMemberTokenRequest TeamMemberTokenRequest

// NewTeamMemberTokenRequest instantiates a new TeamMemberTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberTokenRequest(scopes []string) *TeamMemberTokenRequest {
	this := TeamMemberTokenRequest{}
	this.Scopes = scopes
	return &this
}

// NewTeamMemberTokenRequestWithDefaults instantiates a new TeamMemberTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberTokenRequestWithDefaults() *TeamMemberTokenRequest {
	this := TeamMemberTokenRequest{}
	return &this
}

// GetScopes returns the Scopes field value
func (o *TeamMemberTokenRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *TeamMemberTokenRequest) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *TeamMemberTokenRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *TeamMemberTokenRequest) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberTokenRequest) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *TeamMemberTokenRequest) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *TeamMemberTokenRequest) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

func (o TeamMemberTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMemberTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scopes"] = o.Scopes
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	return toSerialize, nil
}

func (o *TeamMemberTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scopes",
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

	varTeamMemberTokenRequest := _TeamMemberTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamMemberTokenRequest)

	if err != nil {
		return err
	}

	*o = TeamMemberTokenRequest(varTeamMemberTokenRequest)

	return err
}

type NullableTeamMemberTokenRequest struct {
	value *TeamMemberTokenRequest
	isSet bool
}

func (v NullableTeamMemberTokenRequest) Get() *TeamMemberTokenRequest {
	return v.value
}

func (v *NullableTeamMemberTokenRequest) Set(val *TeamMemberTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberTokenRequest(val *TeamMemberTokenRequest) *NullableTeamMemberTokenRequest {
	return &NullableTeamMemberTokenRequest{value: val, isSet: true}
}

func (v NullableTeamMemberTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

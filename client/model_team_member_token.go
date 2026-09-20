package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamMemberToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMemberToken{}

// TeamMemberToken struct for TeamMemberToken
type TeamMemberToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	// 以空格分隔的实际授权 Scope。
	Scope string `json:"scope"`
}

type _TeamMemberToken TeamMemberToken

// NewTeamMemberToken instantiates a new TeamMemberToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberToken(accessToken string, tokenType string, expiresIn int32, scope string) *TeamMemberToken {
	this := TeamMemberToken{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	this.ExpiresIn = expiresIn
	this.Scope = scope
	return &this
}

// NewTeamMemberTokenWithDefaults instantiates a new TeamMemberToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberTokenWithDefaults() *TeamMemberToken {
	this := TeamMemberToken{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TeamMemberToken) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TeamMemberToken) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TeamMemberToken) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *TeamMemberToken) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *TeamMemberToken) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *TeamMemberToken) SetTokenType(v string) {
	o.TokenType = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *TeamMemberToken) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *TeamMemberToken) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *TeamMemberToken) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetScope returns the Scope field value
func (o *TeamMemberToken) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TeamMemberToken) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *TeamMemberToken) SetScope(v string) {
	o.Scope = v
}

func (o TeamMemberToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMemberToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["token_type"] = o.TokenType
	toSerialize["expires_in"] = o.ExpiresIn
	toSerialize["scope"] = o.Scope
	return toSerialize, nil
}

func (o *TeamMemberToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"token_type",
		"expires_in",
		"scope",
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

	varTeamMemberToken := _TeamMemberToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamMemberToken)

	if err != nil {
		return err
	}

	*o = TeamMemberToken(varTeamMemberToken)

	return err
}

type NullableTeamMemberToken struct {
	value *TeamMemberToken
	isSet bool
}

func (v NullableTeamMemberToken) Get() *TeamMemberToken {
	return v.value
}

func (v *NullableTeamMemberToken) Set(val *TeamMemberToken) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberToken) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberToken(val *TeamMemberToken) *NullableTeamMemberToken {
	return &NullableTeamMemberToken{value: val, isSet: true}
}

func (v NullableTeamMemberToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectToken{}

// ProjectToken struct for ProjectToken
type ProjectToken struct {
	// 完整 Bearer Token，仅在创建响应中返回一次。
	Token     string `json:"token" validate:"regexp=^meoo_ak_"`
	TokenType string `json:"token_type"`
	ProjectId string `json:"project_id"`
	Scope     string `json:"scope"`
	// Unix 毫秒时间戳；null 表示永不过期。
	ExpiresAt NullableInt64 `json:"expires_at"`
	CreatedAt int64         `json:"created_at"`
}

type _ProjectToken ProjectToken

// NewProjectToken instantiates a new ProjectToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectToken(token string, tokenType string, projectId string, scope string, expiresAt NullableInt64, createdAt int64) *ProjectToken {
	this := ProjectToken{}
	this.Token = token
	this.TokenType = tokenType
	this.ProjectId = projectId
	this.Scope = scope
	this.ExpiresAt = expiresAt
	this.CreatedAt = createdAt
	return &this
}

// NewProjectTokenWithDefaults instantiates a new ProjectToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTokenWithDefaults() *ProjectToken {
	this := ProjectToken{}
	return &this
}

// GetToken returns the Token field value
func (o *ProjectToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ProjectToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ProjectToken) SetToken(v string) {
	o.Token = v
}

// GetTokenType returns the TokenType field value
func (o *ProjectToken) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *ProjectToken) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *ProjectToken) SetTokenType(v string) {
	o.TokenType = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectToken) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectToken) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectToken) SetProjectId(v string) {
	o.ProjectId = v
}

// GetScope returns the Scope field value
func (o *ProjectToken) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ProjectToken) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ProjectToken) SetScope(v string) {
	o.Scope = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ProjectToken) GetExpiresAt() int64 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectToken) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *ProjectToken) SetExpiresAt(v int64) {
	o.ExpiresAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectToken) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectToken) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectToken) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o ProjectToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["token_type"] = o.TokenType
	toSerialize["project_id"] = o.ProjectId
	toSerialize["scope"] = o.Scope
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *ProjectToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"token_type",
		"project_id",
		"scope",
		"expires_at",
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

	varProjectToken := _ProjectToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectToken)

	if err != nil {
		return err
	}

	*o = ProjectToken(varProjectToken)

	return err
}

type NullableProjectToken struct {
	value *ProjectToken
	isSet bool
}

func (v NullableProjectToken) Get() *ProjectToken {
	return v.value
}

func (v *NullableProjectToken) Set(val *ProjectToken) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectToken) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectToken(val *ProjectToken) *NullableProjectToken {
	return &NullableProjectToken{value: val, isSet: true}
}

func (v NullableProjectToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

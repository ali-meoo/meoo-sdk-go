package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectTokenCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectTokenCreateRequest{}

// ProjectTokenCreateRequest struct for ProjectTokenCreateRequest
type ProjectTokenCreateRequest struct {
	Name          string  `json:"name"`
	Description   *string `json:"description,omitempty"`
	ExpiresInDays *int32  `json:"expires_in_days,omitempty"`
	// 项目 API Key 权限；`*` 表示当前及未来所有允许配置给项目 API Key 的权限，`cli.compat` 表示绑定项目内的全部 CLI 兼容操作。
	Scopes []string `json:"scopes"`
}

type _ProjectTokenCreateRequest ProjectTokenCreateRequest

// NewProjectTokenCreateRequest instantiates a new ProjectTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTokenCreateRequest(name string, scopes []string) *ProjectTokenCreateRequest {
	this := ProjectTokenCreateRequest{}
	this.Name = name
	this.Scopes = scopes
	return &this
}

// NewProjectTokenCreateRequestWithDefaults instantiates a new ProjectTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTokenCreateRequestWithDefaults() *ProjectTokenCreateRequest {
	this := ProjectTokenCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectTokenCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectTokenCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectTokenCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectTokenCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTokenCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectTokenCreateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectTokenCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresInDays returns the ExpiresInDays field value if set, zero value otherwise.
func (o *ProjectTokenCreateRequest) GetExpiresInDays() int32 {
	if o == nil || IsNil(o.ExpiresInDays) {
		var ret int32
		return ret
	}
	return *o.ExpiresInDays
}

// GetExpiresInDaysOk returns a tuple with the ExpiresInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTokenCreateRequest) GetExpiresInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresInDays) {
		return nil, false
	}
	return o.ExpiresInDays, true
}

// HasExpiresInDays returns a boolean if a field has been set.
func (o *ProjectTokenCreateRequest) HasExpiresInDays() bool {
	if o != nil && !IsNil(o.ExpiresInDays) {
		return true
	}

	return false
}

// SetExpiresInDays gets a reference to the given int32 and assigns it to the ExpiresInDays field.
func (o *ProjectTokenCreateRequest) SetExpiresInDays(v int32) {
	o.ExpiresInDays = &v
}

// GetScopes returns the Scopes field value
func (o *ProjectTokenCreateRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ProjectTokenCreateRequest) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ProjectTokenCreateRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o ProjectTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectTokenCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiresInDays) {
		toSerialize["expires_in_days"] = o.ExpiresInDays
	}
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

func (o *ProjectTokenCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varProjectTokenCreateRequest := _ProjectTokenCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectTokenCreateRequest)

	if err != nil {
		return err
	}

	*o = ProjectTokenCreateRequest(varProjectTokenCreateRequest)

	return err
}

type NullableProjectTokenCreateRequest struct {
	value *ProjectTokenCreateRequest
	isSet bool
}

func (v NullableProjectTokenCreateRequest) Get() *ProjectTokenCreateRequest {
	return v.value
}

func (v *NullableProjectTokenCreateRequest) Set(val *ProjectTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTokenCreateRequest(val *ProjectTokenCreateRequest) *NullableProjectTokenCreateRequest {
	return &NullableProjectTokenCreateRequest{value: val, isSet: true}
}

func (v NullableProjectTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
)

// checks if the ProjectCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCreateRequest{}

// ProjectCreateRequest struct for ProjectCreateRequest
type ProjectCreateRequest struct {
	// 先去除首尾空白，再校验长度。
	Name *string `json:"name,omitempty" validate:"regexp=\\\\S"`
	// 可创建 web、app 或 miniprogram；非 Web 类型可创建不代表已开放其原生发布渠道。
	Type *string `json:"type,omitempty"`
}

// NewProjectCreateRequest instantiates a new ProjectCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreateRequest() *ProjectCreateRequest {
	this := ProjectCreateRequest{}
	var type_ string = "web"
	this.Type = &type_
	return &this
}

// NewProjectCreateRequestWithDefaults instantiates a new ProjectCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateRequestWithDefaults() *ProjectCreateRequest {
	this := ProjectCreateRequest{}
	var type_ string = "web"
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectCreateRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectCreateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectCreateRequest) SetType(v string) {
	o.Type = &v
}

func (o ProjectCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableProjectCreateRequest struct {
	value *ProjectCreateRequest
	isSet bool
}

func (v NullableProjectCreateRequest) Get() *ProjectCreateRequest {
	return v.value
}

func (v *NullableProjectCreateRequest) Set(val *ProjectCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreateRequest(val *ProjectCreateRequest) *NullableProjectCreateRequest {
	return &NullableProjectCreateRequest{value: val, isSet: true}
}

func (v NullableProjectCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

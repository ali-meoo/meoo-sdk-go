package client

import (
	"encoding/json"
)

// checks if the ReleaseCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseCreateRequest{}

// ReleaseCreateRequest struct for ReleaseCreateRequest
type ReleaseCreateRequest struct {
	// 发布访问过期时间，Unix 毫秒时间戳，范围 1-253402271999999；省略或 null 表示永久有效。
	ExpiresAt NullableInt64 `json:"expires_at,omitempty"`
}

// NewReleaseCreateRequest instantiates a new ReleaseCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseCreateRequest() *ReleaseCreateRequest {
	this := ReleaseCreateRequest{}
	return &this
}

// NewReleaseCreateRequestWithDefaults instantiates a new ReleaseCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseCreateRequestWithDefaults() *ReleaseCreateRequest {
	this := ReleaseCreateRequest{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseCreateRequest) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseCreateRequest) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ReleaseCreateRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableInt64 and assigns it to the ExpiresAt field.
func (o *ReleaseCreateRequest) SetExpiresAt(v int64) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ReleaseCreateRequest) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ReleaseCreateRequest) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o ReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	return toSerialize, nil
}

type NullableReleaseCreateRequest struct {
	value *ReleaseCreateRequest
	isSet bool
}

func (v NullableReleaseCreateRequest) Get() *ReleaseCreateRequest {
	return v.value
}

func (v *NullableReleaseCreateRequest) Set(val *ReleaseCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseCreateRequest(val *ReleaseCreateRequest) *NullableReleaseCreateRequest {
	return &NullableReleaseCreateRequest{value: val, isSet: true}
}

func (v NullableReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

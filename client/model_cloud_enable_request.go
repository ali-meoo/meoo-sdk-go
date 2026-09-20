package client

import (
	"encoding/json"
)

// checks if the CloudEnableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEnableRequest{}

// CloudEnableRequest struct for CloudEnableRequest
type CloudEnableRequest struct {
	Description *string `json:"description,omitempty"`
}

// NewCloudEnableRequest instantiates a new CloudEnableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEnableRequest() *CloudEnableRequest {
	this := CloudEnableRequest{}
	return &this
}

// NewCloudEnableRequestWithDefaults instantiates a new CloudEnableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEnableRequestWithDefaults() *CloudEnableRequest {
	this := CloudEnableRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudEnableRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEnableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudEnableRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudEnableRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CloudEnableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEnableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCloudEnableRequest struct {
	value *CloudEnableRequest
	isSet bool
}

func (v NullableCloudEnableRequest) Get() *CloudEnableRequest {
	return v.value
}

func (v *NullableCloudEnableRequest) Set(val *CloudEnableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEnableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEnableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEnableRequest(val *CloudEnableRequest) *NullableCloudEnableRequest {
	return &NullableCloudEnableRequest{value: val, isSet: true}
}

func (v NullableCloudEnableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEnableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

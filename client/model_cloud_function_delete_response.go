package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudFunctionDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunctionDeleteResponse{}

// CloudFunctionDeleteResponse struct for CloudFunctionDeleteResponse
type CloudFunctionDeleteResponse struct {
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
}

type _CloudFunctionDeleteResponse CloudFunctionDeleteResponse

// NewCloudFunctionDeleteResponse instantiates a new CloudFunctionDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunctionDeleteResponse(name string, deleted bool) *CloudFunctionDeleteResponse {
	this := CloudFunctionDeleteResponse{}
	this.Name = name
	this.Deleted = deleted
	return &this
}

// NewCloudFunctionDeleteResponseWithDefaults instantiates a new CloudFunctionDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionDeleteResponseWithDefaults() *CloudFunctionDeleteResponse {
	this := CloudFunctionDeleteResponse{}
	return &this
}

// GetName returns the Name field value
func (o *CloudFunctionDeleteResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionDeleteResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudFunctionDeleteResponse) SetName(v string) {
	o.Name = v
}

// GetDeleted returns the Deleted field value
func (o *CloudFunctionDeleteResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionDeleteResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *CloudFunctionDeleteResponse) SetDeleted(v bool) {
	o.Deleted = v
}

func (o CloudFunctionDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunctionDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

func (o *CloudFunctionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"deleted",
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

	varCloudFunctionDeleteResponse := _CloudFunctionDeleteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudFunctionDeleteResponse)

	if err != nil {
		return err
	}

	*o = CloudFunctionDeleteResponse(varCloudFunctionDeleteResponse)

	return err
}

type NullableCloudFunctionDeleteResponse struct {
	value *CloudFunctionDeleteResponse
	isSet bool
}

func (v NullableCloudFunctionDeleteResponse) Get() *CloudFunctionDeleteResponse {
	return v.value
}

func (v *NullableCloudFunctionDeleteResponse) Set(val *CloudFunctionDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunctionDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunctionDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunctionDeleteResponse(val *CloudFunctionDeleteResponse) *NullableCloudFunctionDeleteResponse {
	return &NullableCloudFunctionDeleteResponse{value: val, isSet: true}
}

func (v NullableCloudFunctionDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunctionDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

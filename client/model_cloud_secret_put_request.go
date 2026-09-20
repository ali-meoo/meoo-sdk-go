package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudSecretPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudSecretPutRequest{}

// CloudSecretPutRequest struct for CloudSecretPutRequest
type CloudSecretPutRequest struct {
	Value string `json:"value"`
}

type _CloudSecretPutRequest CloudSecretPutRequest

// NewCloudSecretPutRequest instantiates a new CloudSecretPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSecretPutRequest(value string) *CloudSecretPutRequest {
	this := CloudSecretPutRequest{}
	this.Value = value
	return &this
}

// NewCloudSecretPutRequestWithDefaults instantiates a new CloudSecretPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSecretPutRequestWithDefaults() *CloudSecretPutRequest {
	this := CloudSecretPutRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *CloudSecretPutRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CloudSecretPutRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CloudSecretPutRequest) SetValue(v string) {
	o.Value = v
}

func (o CloudSecretPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudSecretPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CloudSecretPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varCloudSecretPutRequest := _CloudSecretPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudSecretPutRequest)

	if err != nil {
		return err
	}

	*o = CloudSecretPutRequest(varCloudSecretPutRequest)

	return err
}

type NullableCloudSecretPutRequest struct {
	value *CloudSecretPutRequest
	isSet bool
}

func (v NullableCloudSecretPutRequest) Get() *CloudSecretPutRequest {
	return v.value
}

func (v *NullableCloudSecretPutRequest) Set(val *CloudSecretPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSecretPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSecretPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSecretPutRequest(val *CloudSecretPutRequest) *NullableCloudSecretPutRequest {
	return &NullableCloudSecretPutRequest{value: val, isSet: true}
}

func (v NullableCloudSecretPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSecretPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

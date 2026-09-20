package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageObjectsDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageObjectsDeleteRequest{}

// CloudStorageObjectsDeleteRequest struct for CloudStorageObjectsDeleteRequest
type CloudStorageObjectsDeleteRequest struct {
	Paths []string `json:"paths"`
}

type _CloudStorageObjectsDeleteRequest CloudStorageObjectsDeleteRequest

// NewCloudStorageObjectsDeleteRequest instantiates a new CloudStorageObjectsDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageObjectsDeleteRequest(paths []string) *CloudStorageObjectsDeleteRequest {
	this := CloudStorageObjectsDeleteRequest{}
	this.Paths = paths
	return &this
}

// NewCloudStorageObjectsDeleteRequestWithDefaults instantiates a new CloudStorageObjectsDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageObjectsDeleteRequestWithDefaults() *CloudStorageObjectsDeleteRequest {
	this := CloudStorageObjectsDeleteRequest{}
	return &this
}

// GetPaths returns the Paths field value
func (o *CloudStorageObjectsDeleteRequest) GetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectsDeleteRequest) GetPathsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *CloudStorageObjectsDeleteRequest) SetPaths(v []string) {
	o.Paths = v
}

func (o CloudStorageObjectsDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageObjectsDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paths"] = o.Paths
	return toSerialize, nil
}

func (o *CloudStorageObjectsDeleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"paths",
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

	varCloudStorageObjectsDeleteRequest := _CloudStorageObjectsDeleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageObjectsDeleteRequest)

	if err != nil {
		return err
	}

	*o = CloudStorageObjectsDeleteRequest(varCloudStorageObjectsDeleteRequest)

	return err
}

type NullableCloudStorageObjectsDeleteRequest struct {
	value *CloudStorageObjectsDeleteRequest
	isSet bool
}

func (v NullableCloudStorageObjectsDeleteRequest) Get() *CloudStorageObjectsDeleteRequest {
	return v.value
}

func (v *NullableCloudStorageObjectsDeleteRequest) Set(val *CloudStorageObjectsDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageObjectsDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageObjectsDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageObjectsDeleteRequest(val *CloudStorageObjectsDeleteRequest) *NullableCloudStorageObjectsDeleteRequest {
	return &NullableCloudStorageObjectsDeleteRequest{value: val, isSet: true}
}

func (v NullableCloudStorageObjectsDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageObjectsDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

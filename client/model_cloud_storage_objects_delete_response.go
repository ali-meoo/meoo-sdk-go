package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageObjectsDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageObjectsDeleteResponse{}

// CloudStorageObjectsDeleteResponse struct for CloudStorageObjectsDeleteResponse
type CloudStorageObjectsDeleteResponse struct {
	BucketName string   `json:"bucket_name"`
	Deleted    []string `json:"deleted"`
}

type _CloudStorageObjectsDeleteResponse CloudStorageObjectsDeleteResponse

// NewCloudStorageObjectsDeleteResponse instantiates a new CloudStorageObjectsDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageObjectsDeleteResponse(bucketName string, deleted []string) *CloudStorageObjectsDeleteResponse {
	this := CloudStorageObjectsDeleteResponse{}
	this.BucketName = bucketName
	this.Deleted = deleted
	return &this
}

// NewCloudStorageObjectsDeleteResponseWithDefaults instantiates a new CloudStorageObjectsDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageObjectsDeleteResponseWithDefaults() *CloudStorageObjectsDeleteResponse {
	this := CloudStorageObjectsDeleteResponse{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *CloudStorageObjectsDeleteResponse) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectsDeleteResponse) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *CloudStorageObjectsDeleteResponse) SetBucketName(v string) {
	o.BucketName = v
}

// GetDeleted returns the Deleted field value
func (o *CloudStorageObjectsDeleteResponse) GetDeleted() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectsDeleteResponse) GetDeletedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted, true
}

// SetDeleted sets field value
func (o *CloudStorageObjectsDeleteResponse) SetDeleted(v []string) {
	o.Deleted = v
}

func (o CloudStorageObjectsDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageObjectsDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

func (o *CloudStorageObjectsDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket_name",
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

	varCloudStorageObjectsDeleteResponse := _CloudStorageObjectsDeleteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageObjectsDeleteResponse)

	if err != nil {
		return err
	}

	*o = CloudStorageObjectsDeleteResponse(varCloudStorageObjectsDeleteResponse)

	return err
}

type NullableCloudStorageObjectsDeleteResponse struct {
	value *CloudStorageObjectsDeleteResponse
	isSet bool
}

func (v NullableCloudStorageObjectsDeleteResponse) Get() *CloudStorageObjectsDeleteResponse {
	return v.value
}

func (v *NullableCloudStorageObjectsDeleteResponse) Set(val *CloudStorageObjectsDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageObjectsDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageObjectsDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageObjectsDeleteResponse(val *CloudStorageObjectsDeleteResponse) *NullableCloudStorageObjectsDeleteResponse {
	return &NullableCloudStorageObjectsDeleteResponse{value: val, isSet: true}
}

func (v NullableCloudStorageObjectsDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageObjectsDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

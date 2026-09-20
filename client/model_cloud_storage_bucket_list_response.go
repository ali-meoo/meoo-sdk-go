package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageBucketListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageBucketListResponse{}

// CloudStorageBucketListResponse struct for CloudStorageBucketListResponse
type CloudStorageBucketListResponse struct {
	Buckets []CloudStorageBucket `json:"buckets"`
}

type _CloudStorageBucketListResponse CloudStorageBucketListResponse

// NewCloudStorageBucketListResponse instantiates a new CloudStorageBucketListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageBucketListResponse(buckets []CloudStorageBucket) *CloudStorageBucketListResponse {
	this := CloudStorageBucketListResponse{}
	this.Buckets = buckets
	return &this
}

// NewCloudStorageBucketListResponseWithDefaults instantiates a new CloudStorageBucketListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageBucketListResponseWithDefaults() *CloudStorageBucketListResponse {
	this := CloudStorageBucketListResponse{}
	return &this
}

// GetBuckets returns the Buckets field value
func (o *CloudStorageBucketListResponse) GetBuckets() []CloudStorageBucket {
	if o == nil {
		var ret []CloudStorageBucket
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *CloudStorageBucketListResponse) GetBucketsOk() ([]CloudStorageBucket, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buckets, true
}

// SetBuckets sets field value
func (o *CloudStorageBucketListResponse) SetBuckets(v []CloudStorageBucket) {
	o.Buckets = v
}

func (o CloudStorageBucketListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageBucketListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buckets"] = o.Buckets
	return toSerialize, nil
}

func (o *CloudStorageBucketListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buckets",
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

	varCloudStorageBucketListResponse := _CloudStorageBucketListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageBucketListResponse)

	if err != nil {
		return err
	}

	*o = CloudStorageBucketListResponse(varCloudStorageBucketListResponse)

	return err
}

type NullableCloudStorageBucketListResponse struct {
	value *CloudStorageBucketListResponse
	isSet bool
}

func (v NullableCloudStorageBucketListResponse) Get() *CloudStorageBucketListResponse {
	return v.value
}

func (v *NullableCloudStorageBucketListResponse) Set(val *CloudStorageBucketListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageBucketListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageBucketListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageBucketListResponse(val *CloudStorageBucketListResponse) *NullableCloudStorageBucketListResponse {
	return &NullableCloudStorageBucketListResponse{value: val, isSet: true}
}

func (v NullableCloudStorageBucketListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageBucketListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

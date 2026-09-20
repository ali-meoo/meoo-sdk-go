package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageObjectPutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageObjectPutResponse{}

// CloudStorageObjectPutResponse struct for CloudStorageObjectPutResponse
type CloudStorageObjectPutResponse struct {
	BucketName  string `json:"bucket_name"`
	Path        string `json:"path"`
	Size        string `json:"size" validate:"regexp=^[0-9]+$"`
	ContentType string `json:"content_type"`
	Uploaded    bool   `json:"uploaded"`
}

type _CloudStorageObjectPutResponse CloudStorageObjectPutResponse

// NewCloudStorageObjectPutResponse instantiates a new CloudStorageObjectPutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageObjectPutResponse(bucketName string, path string, size string, contentType string, uploaded bool) *CloudStorageObjectPutResponse {
	this := CloudStorageObjectPutResponse{}
	this.BucketName = bucketName
	this.Path = path
	this.Size = size
	this.ContentType = contentType
	this.Uploaded = uploaded
	return &this
}

// NewCloudStorageObjectPutResponseWithDefaults instantiates a new CloudStorageObjectPutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageObjectPutResponseWithDefaults() *CloudStorageObjectPutResponse {
	this := CloudStorageObjectPutResponse{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *CloudStorageObjectPutResponse) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutResponse) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *CloudStorageObjectPutResponse) SetBucketName(v string) {
	o.BucketName = v
}

// GetPath returns the Path field value
func (o *CloudStorageObjectPutResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CloudStorageObjectPutResponse) SetPath(v string) {
	o.Path = v
}

// GetSize returns the Size field value
func (o *CloudStorageObjectPutResponse) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutResponse) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CloudStorageObjectPutResponse) SetSize(v string) {
	o.Size = v
}

// GetContentType returns the ContentType field value
func (o *CloudStorageObjectPutResponse) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutResponse) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *CloudStorageObjectPutResponse) SetContentType(v string) {
	o.ContentType = v
}

// GetUploaded returns the Uploaded field value
func (o *CloudStorageObjectPutResponse) GetUploaded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutResponse) GetUploadedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uploaded, true
}

// SetUploaded sets field value
func (o *CloudStorageObjectPutResponse) SetUploaded(v bool) {
	o.Uploaded = v
}

func (o CloudStorageObjectPutResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageObjectPutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["path"] = o.Path
	toSerialize["size"] = o.Size
	toSerialize["content_type"] = o.ContentType
	toSerialize["uploaded"] = o.Uploaded
	return toSerialize, nil
}

func (o *CloudStorageObjectPutResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket_name",
		"path",
		"size",
		"content_type",
		"uploaded",
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

	varCloudStorageObjectPutResponse := _CloudStorageObjectPutResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageObjectPutResponse)

	if err != nil {
		return err
	}

	*o = CloudStorageObjectPutResponse(varCloudStorageObjectPutResponse)

	return err
}

type NullableCloudStorageObjectPutResponse struct {
	value *CloudStorageObjectPutResponse
	isSet bool
}

func (v NullableCloudStorageObjectPutResponse) Get() *CloudStorageObjectPutResponse {
	return v.value
}

func (v *NullableCloudStorageObjectPutResponse) Set(val *CloudStorageObjectPutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageObjectPutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageObjectPutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageObjectPutResponse(val *CloudStorageObjectPutResponse) *NullableCloudStorageObjectPutResponse {
	return &NullableCloudStorageObjectPutResponse{value: val, isSet: true}
}

func (v NullableCloudStorageObjectPutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageObjectPutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

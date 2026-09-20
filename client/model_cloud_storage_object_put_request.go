package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageObjectPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageObjectPutRequest{}

// CloudStorageObjectPutRequest struct for CloudStorageObjectPutRequest
type CloudStorageObjectPutRequest struct {
	Path string `json:"path"`
	// Base64 编码的对象内容，解码后最大 5 MiB。
	ContentBase64 string  `json:"content_base64"`
	ContentType   string  `json:"content_type" validate:"regexp=^[\\\\w.+-]+\\/[\\\\w.+-]+$"`
	CacheControl  *string `json:"cache_control,omitempty"`
	Upsert        *bool   `json:"upsert,omitempty"`
}

type _CloudStorageObjectPutRequest CloudStorageObjectPutRequest

// NewCloudStorageObjectPutRequest instantiates a new CloudStorageObjectPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageObjectPutRequest(path string, contentBase64 string, contentType string) *CloudStorageObjectPutRequest {
	this := CloudStorageObjectPutRequest{}
	this.Path = path
	this.ContentBase64 = contentBase64
	this.ContentType = contentType
	var upsert bool = true
	this.Upsert = &upsert
	return &this
}

// NewCloudStorageObjectPutRequestWithDefaults instantiates a new CloudStorageObjectPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageObjectPutRequestWithDefaults() *CloudStorageObjectPutRequest {
	this := CloudStorageObjectPutRequest{}
	var upsert bool = true
	this.Upsert = &upsert
	return &this
}

// GetPath returns the Path field value
func (o *CloudStorageObjectPutRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CloudStorageObjectPutRequest) SetPath(v string) {
	o.Path = v
}

// GetContentBase64 returns the ContentBase64 field value
func (o *CloudStorageObjectPutRequest) GetContentBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentBase64
}

// GetContentBase64Ok returns a tuple with the ContentBase64 field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutRequest) GetContentBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentBase64, true
}

// SetContentBase64 sets field value
func (o *CloudStorageObjectPutRequest) SetContentBase64(v string) {
	o.ContentBase64 = v
}

// GetContentType returns the ContentType field value
func (o *CloudStorageObjectPutRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *CloudStorageObjectPutRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetCacheControl returns the CacheControl field value if set, zero value otherwise.
func (o *CloudStorageObjectPutRequest) GetCacheControl() string {
	if o == nil || IsNil(o.CacheControl) {
		var ret string
		return ret
	}
	return *o.CacheControl
}

// GetCacheControlOk returns a tuple with the CacheControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutRequest) GetCacheControlOk() (*string, bool) {
	if o == nil || IsNil(o.CacheControl) {
		return nil, false
	}
	return o.CacheControl, true
}

// HasCacheControl returns a boolean if a field has been set.
func (o *CloudStorageObjectPutRequest) HasCacheControl() bool {
	if o != nil && !IsNil(o.CacheControl) {
		return true
	}

	return false
}

// SetCacheControl gets a reference to the given string and assigns it to the CacheControl field.
func (o *CloudStorageObjectPutRequest) SetCacheControl(v string) {
	o.CacheControl = &v
}

// GetUpsert returns the Upsert field value if set, zero value otherwise.
func (o *CloudStorageObjectPutRequest) GetUpsert() bool {
	if o == nil || IsNil(o.Upsert) {
		var ret bool
		return ret
	}
	return *o.Upsert
}

// GetUpsertOk returns a tuple with the Upsert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudStorageObjectPutRequest) GetUpsertOk() (*bool, bool) {
	if o == nil || IsNil(o.Upsert) {
		return nil, false
	}
	return o.Upsert, true
}

// HasUpsert returns a boolean if a field has been set.
func (o *CloudStorageObjectPutRequest) HasUpsert() bool {
	if o != nil && !IsNil(o.Upsert) {
		return true
	}

	return false
}

// SetUpsert gets a reference to the given bool and assigns it to the Upsert field.
func (o *CloudStorageObjectPutRequest) SetUpsert(v bool) {
	o.Upsert = &v
}

func (o CloudStorageObjectPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageObjectPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["content_base64"] = o.ContentBase64
	toSerialize["content_type"] = o.ContentType
	if !IsNil(o.CacheControl) {
		toSerialize["cache_control"] = o.CacheControl
	}
	if !IsNil(o.Upsert) {
		toSerialize["upsert"] = o.Upsert
	}
	return toSerialize, nil
}

func (o *CloudStorageObjectPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"content_base64",
		"content_type",
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

	varCloudStorageObjectPutRequest := _CloudStorageObjectPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageObjectPutRequest)

	if err != nil {
		return err
	}

	*o = CloudStorageObjectPutRequest(varCloudStorageObjectPutRequest)

	return err
}

type NullableCloudStorageObjectPutRequest struct {
	value *CloudStorageObjectPutRequest
	isSet bool
}

func (v NullableCloudStorageObjectPutRequest) Get() *CloudStorageObjectPutRequest {
	return v.value
}

func (v *NullableCloudStorageObjectPutRequest) Set(val *CloudStorageObjectPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageObjectPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageObjectPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageObjectPutRequest(val *CloudStorageObjectPutRequest) *NullableCloudStorageObjectPutRequest {
	return &NullableCloudStorageObjectPutRequest{value: val, isSet: true}
}

func (v NullableCloudStorageObjectPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageObjectPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

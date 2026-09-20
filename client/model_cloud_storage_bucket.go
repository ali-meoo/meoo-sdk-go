package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageBucket{}

// CloudStorageBucket struct for CloudStorageBucket
type CloudStorageBucket struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
	// 字符串形式的字节数，避免 JavaScript 整数精度损失；未设置时为 null。
	FileSizeLimit    NullableString `json:"file_size_limit"`
	AllowedMimeTypes []string       `json:"allowed_mime_types"`
	CreatedAt        NullableInt64  `json:"created_at"`
	UpdatedAt        NullableInt64  `json:"updated_at"`
}

type _CloudStorageBucket CloudStorageBucket

// NewCloudStorageBucket instantiates a new CloudStorageBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageBucket(name string, public bool, fileSizeLimit NullableString, allowedMimeTypes []string, createdAt NullableInt64, updatedAt NullableInt64) *CloudStorageBucket {
	this := CloudStorageBucket{}
	this.Name = name
	this.Public = public
	this.FileSizeLimit = fileSizeLimit
	this.AllowedMimeTypes = allowedMimeTypes
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCloudStorageBucketWithDefaults instantiates a new CloudStorageBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageBucketWithDefaults() *CloudStorageBucket {
	this := CloudStorageBucket{}
	return &this
}

// GetName returns the Name field value
func (o *CloudStorageBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudStorageBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudStorageBucket) SetName(v string) {
	o.Name = v
}

// GetPublic returns the Public field value
func (o *CloudStorageBucket) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *CloudStorageBucket) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *CloudStorageBucket) SetPublic(v bool) {
	o.Public = v
}

// GetFileSizeLimit returns the FileSizeLimit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloudStorageBucket) GetFileSizeLimit() string {
	if o == nil || o.FileSizeLimit.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileSizeLimit.Get()
}

// GetFileSizeLimitOk returns a tuple with the FileSizeLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageBucket) GetFileSizeLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSizeLimit.Get(), o.FileSizeLimit.IsSet()
}

// SetFileSizeLimit sets field value
func (o *CloudStorageBucket) SetFileSizeLimit(v string) {
	o.FileSizeLimit.Set(&v)
}

// GetAllowedMimeTypes returns the AllowedMimeTypes field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CloudStorageBucket) GetAllowedMimeTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedMimeTypes
}

// GetAllowedMimeTypesOk returns a tuple with the AllowedMimeTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageBucket) GetAllowedMimeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMimeTypes) {
		return nil, false
	}
	return o.AllowedMimeTypes, true
}

// SetAllowedMimeTypes sets field value
func (o *CloudStorageBucket) SetAllowedMimeTypes(v []string) {
	o.AllowedMimeTypes = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudStorageBucket) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageBucket) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *CloudStorageBucket) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudStorageBucket) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageBucket) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *CloudStorageBucket) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}

func (o CloudStorageBucket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["public"] = o.Public
	toSerialize["file_size_limit"] = o.FileSizeLimit.Get()
	if o.AllowedMimeTypes != nil {
		toSerialize["allowed_mime_types"] = o.AllowedMimeTypes
	}
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	return toSerialize, nil
}

func (o *CloudStorageBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"public",
		"file_size_limit",
		"allowed_mime_types",
		"created_at",
		"updated_at",
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

	varCloudStorageBucket := _CloudStorageBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageBucket)

	if err != nil {
		return err
	}

	*o = CloudStorageBucket(varCloudStorageBucket)

	return err
}

type NullableCloudStorageBucket struct {
	value *CloudStorageBucket
	isSet bool
}

func (v NullableCloudStorageBucket) Get() *CloudStorageBucket {
	return v.value
}

func (v *NullableCloudStorageBucket) Set(val *CloudStorageBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageBucket(val *CloudStorageBucket) *NullableCloudStorageBucket {
	return &NullableCloudStorageBucket{value: val, isSet: true}
}

func (v NullableCloudStorageBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

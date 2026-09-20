package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudStorageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudStorageObject{}

// CloudStorageObject struct for CloudStorageObject
type CloudStorageObject struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// 字符串形式的字节数；目录前缀为 null。
	Size           NullableString `json:"size" validate:"regexp=^[0-9]+$"`
	MimeType       NullableString `json:"mime_type"`
	CacheControl   NullableString `json:"cache_control"`
	Etag           NullableString `json:"etag"`
	CreatedAt      NullableInt64  `json:"created_at"`
	UpdatedAt      NullableInt64  `json:"updated_at"`
	LastAccessedAt NullableInt64  `json:"last_accessed_at"`
}

type _CloudStorageObject CloudStorageObject

// NewCloudStorageObject instantiates a new CloudStorageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudStorageObject(name string, type_ string, size NullableString, mimeType NullableString, cacheControl NullableString, etag NullableString, createdAt NullableInt64, updatedAt NullableInt64, lastAccessedAt NullableInt64) *CloudStorageObject {
	this := CloudStorageObject{}
	this.Name = name
	this.Type = type_
	this.Size = size
	this.MimeType = mimeType
	this.CacheControl = cacheControl
	this.Etag = etag
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.LastAccessedAt = lastAccessedAt
	return &this
}

// NewCloudStorageObjectWithDefaults instantiates a new CloudStorageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudStorageObjectWithDefaults() *CloudStorageObject {
	this := CloudStorageObject{}
	return &this
}

// GetName returns the Name field value
func (o *CloudStorageObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudStorageObject) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CloudStorageObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudStorageObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudStorageObject) SetType(v string) {
	o.Type = v
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloudStorageObject) GetSize() string {
	if o == nil || o.Size.Get() == nil {
		var ret string
		return ret
	}

	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// SetSize sets field value
func (o *CloudStorageObject) SetSize(v string) {
	o.Size.Set(&v)
}

// GetMimeType returns the MimeType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloudStorageObject) GetMimeType() string {
	if o == nil || o.MimeType.Get() == nil {
		var ret string
		return ret
	}

	return *o.MimeType.Get()
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MimeType.Get(), o.MimeType.IsSet()
}

// SetMimeType sets field value
func (o *CloudStorageObject) SetMimeType(v string) {
	o.MimeType.Set(&v)
}

// GetCacheControl returns the CacheControl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloudStorageObject) GetCacheControl() string {
	if o == nil || o.CacheControl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CacheControl.Get()
}

// GetCacheControlOk returns a tuple with the CacheControl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetCacheControlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CacheControl.Get(), o.CacheControl.IsSet()
}

// SetCacheControl sets field value
func (o *CloudStorageObject) SetCacheControl(v string) {
	o.CacheControl.Set(&v)
}

// GetEtag returns the Etag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloudStorageObject) GetEtag() string {
	if o == nil || o.Etag.Get() == nil {
		var ret string
		return ret
	}

	return *o.Etag.Get()
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Etag.Get(), o.Etag.IsSet()
}

// SetEtag sets field value
func (o *CloudStorageObject) SetEtag(v string) {
	o.Etag.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudStorageObject) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *CloudStorageObject) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudStorageObject) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *CloudStorageObject) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}

// GetLastAccessedAt returns the LastAccessedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudStorageObject) GetLastAccessedAt() int64 {
	if o == nil || o.LastAccessedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.LastAccessedAt.Get()
}

// GetLastAccessedAtOk returns a tuple with the LastAccessedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudStorageObject) GetLastAccessedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAccessedAt.Get(), o.LastAccessedAt.IsSet()
}

// SetLastAccessedAt sets field value
func (o *CloudStorageObject) SetLastAccessedAt(v int64) {
	o.LastAccessedAt.Set(&v)
}

func (o CloudStorageObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudStorageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["size"] = o.Size.Get()
	toSerialize["mime_type"] = o.MimeType.Get()
	toSerialize["cache_control"] = o.CacheControl.Get()
	toSerialize["etag"] = o.Etag.Get()
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["last_accessed_at"] = o.LastAccessedAt.Get()
	return toSerialize, nil
}

func (o *CloudStorageObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"size",
		"mime_type",
		"cache_control",
		"etag",
		"created_at",
		"updated_at",
		"last_accessed_at",
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

	varCloudStorageObject := _CloudStorageObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudStorageObject)

	if err != nil {
		return err
	}

	*o = CloudStorageObject(varCloudStorageObject)

	return err
}

type NullableCloudStorageObject struct {
	value *CloudStorageObject
	isSet bool
}

func (v NullableCloudStorageObject) Get() *CloudStorageObject {
	return v.value
}

func (v *NullableCloudStorageObject) Set(val *CloudStorageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStorageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStorageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStorageObject(val *CloudStorageObject) *NullableCloudStorageObject {
	return &NullableCloudStorageObject{value: val, isSet: true}
}

func (v NullableCloudStorageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStorageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

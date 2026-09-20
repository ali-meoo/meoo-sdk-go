package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StaticReleasePrepareRequestArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleasePrepareRequestArtifact{}

// StaticReleasePrepareRequestArtifact struct for StaticReleasePrepareRequestArtifact
type StaticReleasePrepareRequestArtifact struct {
	Type string `json:"type"`
	// 先去除首尾空白，再校验长度与 .zip 文件名格式（不区分大小写）。
	Filename      string                                       `json:"filename" validate:"regexp=^[^\\/\\\\\\\\]+\\\\.zip$"`
	ContentType   string                                       `json:"content_type"`
	ContentLength int64                                        `json:"content_length"`
	Checksum      *StaticReleasePrepareRequestArtifactChecksum `json:"checksum,omitempty"`
}

type _StaticReleasePrepareRequestArtifact StaticReleasePrepareRequestArtifact

// NewStaticReleasePrepareRequestArtifact instantiates a new StaticReleasePrepareRequestArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleasePrepareRequestArtifact(type_ string, filename string, contentType string, contentLength int64) *StaticReleasePrepareRequestArtifact {
	this := StaticReleasePrepareRequestArtifact{}
	this.Type = type_
	this.Filename = filename
	this.ContentType = contentType
	this.ContentLength = contentLength
	return &this
}

// NewStaticReleasePrepareRequestArtifactWithDefaults instantiates a new StaticReleasePrepareRequestArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleasePrepareRequestArtifactWithDefaults() *StaticReleasePrepareRequestArtifact {
	this := StaticReleasePrepareRequestArtifact{}
	return &this
}

// GetType returns the Type field value
func (o *StaticReleasePrepareRequestArtifact) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifact) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StaticReleasePrepareRequestArtifact) SetType(v string) {
	o.Type = v
}

// GetFilename returns the Filename field value
func (o *StaticReleasePrepareRequestArtifact) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifact) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *StaticReleasePrepareRequestArtifact) SetFilename(v string) {
	o.Filename = v
}

// GetContentType returns the ContentType field value
func (o *StaticReleasePrepareRequestArtifact) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifact) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *StaticReleasePrepareRequestArtifact) SetContentType(v string) {
	o.ContentType = v
}

// GetContentLength returns the ContentLength field value
func (o *StaticReleasePrepareRequestArtifact) GetContentLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ContentLength
}

// GetContentLengthOk returns a tuple with the ContentLength field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifact) GetContentLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentLength, true
}

// SetContentLength sets field value
func (o *StaticReleasePrepareRequestArtifact) SetContentLength(v int64) {
	o.ContentLength = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *StaticReleasePrepareRequestArtifact) GetChecksum() StaticReleasePrepareRequestArtifactChecksum {
	if o == nil || IsNil(o.Checksum) {
		var ret StaticReleasePrepareRequestArtifactChecksum
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifact) GetChecksumOk() (*StaticReleasePrepareRequestArtifactChecksum, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *StaticReleasePrepareRequestArtifact) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given StaticReleasePrepareRequestArtifactChecksum and assigns it to the Checksum field.
func (o *StaticReleasePrepareRequestArtifact) SetChecksum(v StaticReleasePrepareRequestArtifactChecksum) {
	o.Checksum = &v
}

func (o StaticReleasePrepareRequestArtifact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleasePrepareRequestArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["filename"] = o.Filename
	toSerialize["content_type"] = o.ContentType
	toSerialize["content_length"] = o.ContentLength
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	return toSerialize, nil
}

func (o *StaticReleasePrepareRequestArtifact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"filename",
		"content_type",
		"content_length",
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

	varStaticReleasePrepareRequestArtifact := _StaticReleasePrepareRequestArtifact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticReleasePrepareRequestArtifact)

	if err != nil {
		return err
	}

	*o = StaticReleasePrepareRequestArtifact(varStaticReleasePrepareRequestArtifact)

	return err
}

type NullableStaticReleasePrepareRequestArtifact struct {
	value *StaticReleasePrepareRequestArtifact
	isSet bool
}

func (v NullableStaticReleasePrepareRequestArtifact) Get() *StaticReleasePrepareRequestArtifact {
	return v.value
}

func (v *NullableStaticReleasePrepareRequestArtifact) Set(val *StaticReleasePrepareRequestArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleasePrepareRequestArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleasePrepareRequestArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleasePrepareRequestArtifact(val *StaticReleasePrepareRequestArtifact) *NullableStaticReleasePrepareRequestArtifact {
	return &NullableStaticReleasePrepareRequestArtifact{value: val, isSet: true}
}

func (v NullableStaticReleasePrepareRequestArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleasePrepareRequestArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

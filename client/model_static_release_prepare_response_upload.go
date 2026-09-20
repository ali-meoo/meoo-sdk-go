package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StaticReleasePrepareResponseUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleasePrepareResponseUpload{}

// StaticReleasePrepareResponseUpload struct for StaticReleasePrepareResponseUpload
type StaticReleasePrepareResponseUpload struct {
	Method    string            `json:"method"`
	Url       string            `json:"url"`
	Headers   map[string]string `json:"headers"`
	ExpiresAt int64             `json:"expires_at"`
}

type _StaticReleasePrepareResponseUpload StaticReleasePrepareResponseUpload

// NewStaticReleasePrepareResponseUpload instantiates a new StaticReleasePrepareResponseUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleasePrepareResponseUpload(method string, url string, headers map[string]string, expiresAt int64) *StaticReleasePrepareResponseUpload {
	this := StaticReleasePrepareResponseUpload{}
	this.Method = method
	this.Url = url
	this.Headers = headers
	this.ExpiresAt = expiresAt
	return &this
}

// NewStaticReleasePrepareResponseUploadWithDefaults instantiates a new StaticReleasePrepareResponseUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleasePrepareResponseUploadWithDefaults() *StaticReleasePrepareResponseUpload {
	this := StaticReleasePrepareResponseUpload{}
	return &this
}

// GetMethod returns the Method field value
func (o *StaticReleasePrepareResponseUpload) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponseUpload) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *StaticReleasePrepareResponseUpload) SetMethod(v string) {
	o.Method = v
}

// GetUrl returns the Url field value
func (o *StaticReleasePrepareResponseUpload) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponseUpload) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *StaticReleasePrepareResponseUpload) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value
func (o *StaticReleasePrepareResponseUpload) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponseUpload) GetHeadersOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *StaticReleasePrepareResponseUpload) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *StaticReleasePrepareResponseUpload) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponseUpload) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *StaticReleasePrepareResponseUpload) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

func (o StaticReleasePrepareResponseUpload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleasePrepareResponseUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["url"] = o.Url
	toSerialize["headers"] = o.Headers
	toSerialize["expires_at"] = o.ExpiresAt
	return toSerialize, nil
}

func (o *StaticReleasePrepareResponseUpload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"url",
		"headers",
		"expires_at",
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

	varStaticReleasePrepareResponseUpload := _StaticReleasePrepareResponseUpload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticReleasePrepareResponseUpload)

	if err != nil {
		return err
	}

	*o = StaticReleasePrepareResponseUpload(varStaticReleasePrepareResponseUpload)

	return err
}

type NullableStaticReleasePrepareResponseUpload struct {
	value *StaticReleasePrepareResponseUpload
	isSet bool
}

func (v NullableStaticReleasePrepareResponseUpload) Get() *StaticReleasePrepareResponseUpload {
	return v.value
}

func (v *NullableStaticReleasePrepareResponseUpload) Set(val *StaticReleasePrepareResponseUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleasePrepareResponseUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleasePrepareResponseUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleasePrepareResponseUpload(val *StaticReleasePrepareResponseUpload) *NullableStaticReleasePrepareResponseUpload {
	return &NullableStaticReleasePrepareResponseUpload{value: val, isSet: true}
}

func (v NullableStaticReleasePrepareResponseUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleasePrepareResponseUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

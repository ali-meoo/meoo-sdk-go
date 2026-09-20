package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudEnableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEnableResponse{}

// CloudEnableResponse struct for CloudEnableResponse
type CloudEnableResponse struct {
	Status    string         `json:"status"`
	Created   bool           `json:"created"`
	PublicUrl NullableString `json:"public_url,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
}

type _CloudEnableResponse CloudEnableResponse

// NewCloudEnableResponse instantiates a new CloudEnableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEnableResponse(status string, created bool) *CloudEnableResponse {
	this := CloudEnableResponse{}
	this.Status = status
	this.Created = created
	return &this
}

// NewCloudEnableResponseWithDefaults instantiates a new CloudEnableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEnableResponseWithDefaults() *CloudEnableResponse {
	this := CloudEnableResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *CloudEnableResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudEnableResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudEnableResponse) SetStatus(v string) {
	o.Status = v
}

// GetCreated returns the Created field value
func (o *CloudEnableResponse) GetCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CloudEnableResponse) GetCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CloudEnableResponse) SetCreated(v bool) {
	o.Created = v
}

// GetPublicUrl returns the PublicUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudEnableResponse) GetPublicUrl() string {
	if o == nil || IsNil(o.PublicUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PublicUrl.Get()
}

// GetPublicUrlOk returns a tuple with the PublicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudEnableResponse) GetPublicUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicUrl.Get(), o.PublicUrl.IsSet()
}

// HasPublicUrl returns a boolean if a field has been set.
func (o *CloudEnableResponse) HasPublicUrl() bool {
	if o != nil && o.PublicUrl.IsSet() {
		return true
	}

	return false
}

// SetPublicUrl gets a reference to the given NullableString and assigns it to the PublicUrl field.
func (o *CloudEnableResponse) SetPublicUrl(v string) {
	o.PublicUrl.Set(&v)
}

// SetPublicUrlNil sets the value for PublicUrl to be an explicit nil
func (o *CloudEnableResponse) SetPublicUrlNil() {
	o.PublicUrl.Set(nil)
}

// UnsetPublicUrl ensures that no value is present for PublicUrl, not even an explicit nil
func (o *CloudEnableResponse) UnsetPublicUrl() {
	o.PublicUrl.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CloudEnableResponse) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEnableResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CloudEnableResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *CloudEnableResponse) SetWarnings(v []string) {
	o.Warnings = v
}

func (o CloudEnableResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEnableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["created"] = o.Created
	if o.PublicUrl.IsSet() {
		toSerialize["public_url"] = o.PublicUrl.Get()
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CloudEnableResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"created",
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

	varCloudEnableResponse := _CloudEnableResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudEnableResponse)

	if err != nil {
		return err
	}

	*o = CloudEnableResponse(varCloudEnableResponse)

	return err
}

type NullableCloudEnableResponse struct {
	value *CloudEnableResponse
	isSet bool
}

func (v NullableCloudEnableResponse) Get() *CloudEnableResponse {
	return v.value
}

func (v *NullableCloudEnableResponse) Set(val *CloudEnableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEnableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEnableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEnableResponse(val *CloudEnableResponse) *NullableCloudEnableResponse {
	return &NullableCloudEnableResponse{value: val, isSet: true}
}

func (v NullableCloudEnableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEnableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

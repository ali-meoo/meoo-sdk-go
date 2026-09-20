package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudCredentialsResponse{}

// CloudCredentialsResponse struct for CloudCredentialsResponse
type CloudCredentialsResponse struct {
	Url     NullableString `json:"url"`
	AnonKey string         `json:"anon_key"`
	// Service Role Key，仅可在可信服务端使用，可绕过 RLS。
	ServiceKey string `json:"service_key"`
}

type _CloudCredentialsResponse CloudCredentialsResponse

// NewCloudCredentialsResponse instantiates a new CloudCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCredentialsResponse(url NullableString, anonKey string, serviceKey string) *CloudCredentialsResponse {
	this := CloudCredentialsResponse{}
	this.Url = url
	this.AnonKey = anonKey
	this.ServiceKey = serviceKey
	return &this
}

// NewCloudCredentialsResponseWithDefaults instantiates a new CloudCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCredentialsResponseWithDefaults() *CloudCredentialsResponse {
	this := CloudCredentialsResponse{}
	return &this
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CloudCredentialsResponse) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *CloudCredentialsResponse) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetAnonKey returns the AnonKey field value
func (o *CloudCredentialsResponse) GetAnonKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnonKey
}

// GetAnonKeyOk returns a tuple with the AnonKey field value
// and a boolean to check if the value has been set.
func (o *CloudCredentialsResponse) GetAnonKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnonKey, true
}

// SetAnonKey sets field value
func (o *CloudCredentialsResponse) SetAnonKey(v string) {
	o.AnonKey = v
}

// GetServiceKey returns the ServiceKey field value
func (o *CloudCredentialsResponse) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *CloudCredentialsResponse) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *CloudCredentialsResponse) SetServiceKey(v string) {
	o.ServiceKey = v
}

func (o CloudCredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url.Get()
	toSerialize["anon_key"] = o.AnonKey
	toSerialize["service_key"] = o.ServiceKey
	return toSerialize, nil
}

func (o *CloudCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"anon_key",
		"service_key",
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

	varCloudCredentialsResponse := _CloudCredentialsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudCredentialsResponse)

	if err != nil {
		return err
	}

	*o = CloudCredentialsResponse(varCloudCredentialsResponse)

	return err
}

type NullableCloudCredentialsResponse struct {
	value *CloudCredentialsResponse
	isSet bool
}

func (v NullableCloudCredentialsResponse) Get() *CloudCredentialsResponse {
	return v.value
}

func (v *NullableCloudCredentialsResponse) Set(val *CloudCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCredentialsResponse(val *CloudCredentialsResponse) *NullableCloudCredentialsResponse {
	return &NullableCloudCredentialsResponse{value: val, isSet: true}
}

func (v NullableCloudCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

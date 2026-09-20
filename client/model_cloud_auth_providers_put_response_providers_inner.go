package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudAuthProvidersPutResponseProvidersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudAuthProvidersPutResponseProvidersInner{}

// CloudAuthProvidersPutResponseProvidersInner struct for CloudAuthProvidersPutResponseProvidersInner
type CloudAuthProvidersPutResponseProvidersInner struct {
	Provider CloudAuthProvider `json:"provider"`
	Enabled  bool              `json:"enabled"`
}

type _CloudAuthProvidersPutResponseProvidersInner CloudAuthProvidersPutResponseProvidersInner

// NewCloudAuthProvidersPutResponseProvidersInner instantiates a new CloudAuthProvidersPutResponseProvidersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAuthProvidersPutResponseProvidersInner(provider CloudAuthProvider, enabled bool) *CloudAuthProvidersPutResponseProvidersInner {
	this := CloudAuthProvidersPutResponseProvidersInner{}
	this.Provider = provider
	this.Enabled = enabled
	return &this
}

// NewCloudAuthProvidersPutResponseProvidersInnerWithDefaults instantiates a new CloudAuthProvidersPutResponseProvidersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAuthProvidersPutResponseProvidersInnerWithDefaults() *CloudAuthProvidersPutResponseProvidersInner {
	this := CloudAuthProvidersPutResponseProvidersInner{}
	return &this
}

// GetProvider returns the Provider field value
func (o *CloudAuthProvidersPutResponseProvidersInner) GetProvider() CloudAuthProvider {
	if o == nil {
		var ret CloudAuthProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CloudAuthProvidersPutResponseProvidersInner) GetProviderOk() (*CloudAuthProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CloudAuthProvidersPutResponseProvidersInner) SetProvider(v CloudAuthProvider) {
	o.Provider = v
}

// GetEnabled returns the Enabled field value
func (o *CloudAuthProvidersPutResponseProvidersInner) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CloudAuthProvidersPutResponseProvidersInner) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CloudAuthProvidersPutResponseProvidersInner) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CloudAuthProvidersPutResponseProvidersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudAuthProvidersPutResponseProvidersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *CloudAuthProvidersPutResponseProvidersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"enabled",
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

	varCloudAuthProvidersPutResponseProvidersInner := _CloudAuthProvidersPutResponseProvidersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudAuthProvidersPutResponseProvidersInner)

	if err != nil {
		return err
	}

	*o = CloudAuthProvidersPutResponseProvidersInner(varCloudAuthProvidersPutResponseProvidersInner)

	return err
}

type NullableCloudAuthProvidersPutResponseProvidersInner struct {
	value *CloudAuthProvidersPutResponseProvidersInner
	isSet bool
}

func (v NullableCloudAuthProvidersPutResponseProvidersInner) Get() *CloudAuthProvidersPutResponseProvidersInner {
	return v.value
}

func (v *NullableCloudAuthProvidersPutResponseProvidersInner) Set(val *CloudAuthProvidersPutResponseProvidersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAuthProvidersPutResponseProvidersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAuthProvidersPutResponseProvidersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAuthProvidersPutResponseProvidersInner(val *CloudAuthProvidersPutResponseProvidersInner) *NullableCloudAuthProvidersPutResponseProvidersInner {
	return &NullableCloudAuthProvidersPutResponseProvidersInner{value: val, isSet: true}
}

func (v NullableCloudAuthProvidersPutResponseProvidersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAuthProvidersPutResponseProvidersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

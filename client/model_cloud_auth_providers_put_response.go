package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudAuthProvidersPutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudAuthProvidersPutResponse{}

// CloudAuthProvidersPutResponse struct for CloudAuthProvidersPutResponse
type CloudAuthProvidersPutResponse struct {
	Providers            []CloudAuthProvidersPutResponseProvidersInner `json:"providers"`
	AdditionalProperties map[string]interface{}
}

type _CloudAuthProvidersPutResponse CloudAuthProvidersPutResponse

// NewCloudAuthProvidersPutResponse instantiates a new CloudAuthProvidersPutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAuthProvidersPutResponse(providers []CloudAuthProvidersPutResponseProvidersInner) *CloudAuthProvidersPutResponse {
	this := CloudAuthProvidersPutResponse{}
	this.Providers = providers
	return &this
}

// NewCloudAuthProvidersPutResponseWithDefaults instantiates a new CloudAuthProvidersPutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAuthProvidersPutResponseWithDefaults() *CloudAuthProvidersPutResponse {
	this := CloudAuthProvidersPutResponse{}
	return &this
}

// GetProviders returns the Providers field value
func (o *CloudAuthProvidersPutResponse) GetProviders() []CloudAuthProvidersPutResponseProvidersInner {
	if o == nil {
		var ret []CloudAuthProvidersPutResponseProvidersInner
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *CloudAuthProvidersPutResponse) GetProvidersOk() ([]CloudAuthProvidersPutResponseProvidersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *CloudAuthProvidersPutResponse) SetProviders(v []CloudAuthProvidersPutResponseProvidersInner) {
	o.Providers = v
}

func (o CloudAuthProvidersPutResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudAuthProvidersPutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providers"] = o.Providers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudAuthProvidersPutResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"providers",
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

	varCloudAuthProvidersPutResponse := _CloudAuthProvidersPutResponse{}

	err = json.Unmarshal(data, &varCloudAuthProvidersPutResponse)

	if err != nil {
		return err
	}

	*o = CloudAuthProvidersPutResponse(varCloudAuthProvidersPutResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "providers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAuthProvidersPutResponse struct {
	value *CloudAuthProvidersPutResponse
	isSet bool
}

func (v NullableCloudAuthProvidersPutResponse) Get() *CloudAuthProvidersPutResponse {
	return v.value
}

func (v *NullableCloudAuthProvidersPutResponse) Set(val *CloudAuthProvidersPutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAuthProvidersPutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAuthProvidersPutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAuthProvidersPutResponse(val *CloudAuthProvidersPutResponse) *NullableCloudAuthProvidersPutResponse {
	return &NullableCloudAuthProvidersPutResponse{value: val, isSet: true}
}

func (v NullableCloudAuthProvidersPutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAuthProvidersPutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

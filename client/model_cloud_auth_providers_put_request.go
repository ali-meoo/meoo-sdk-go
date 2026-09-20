package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudAuthProvidersPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudAuthProvidersPutRequest{}

// CloudAuthProvidersPutRequest struct for CloudAuthProvidersPutRequest
type CloudAuthProvidersPutRequest struct {
	Providers []CloudAuthProvider `json:"providers"`
}

type _CloudAuthProvidersPutRequest CloudAuthProvidersPutRequest

// NewCloudAuthProvidersPutRequest instantiates a new CloudAuthProvidersPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAuthProvidersPutRequest(providers []CloudAuthProvider) *CloudAuthProvidersPutRequest {
	this := CloudAuthProvidersPutRequest{}
	this.Providers = providers
	return &this
}

// NewCloudAuthProvidersPutRequestWithDefaults instantiates a new CloudAuthProvidersPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAuthProvidersPutRequestWithDefaults() *CloudAuthProvidersPutRequest {
	this := CloudAuthProvidersPutRequest{}
	return &this
}

// GetProviders returns the Providers field value
func (o *CloudAuthProvidersPutRequest) GetProviders() []CloudAuthProvider {
	if o == nil {
		var ret []CloudAuthProvider
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *CloudAuthProvidersPutRequest) GetProvidersOk() ([]CloudAuthProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *CloudAuthProvidersPutRequest) SetProviders(v []CloudAuthProvider) {
	o.Providers = v
}

func (o CloudAuthProvidersPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudAuthProvidersPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providers"] = o.Providers
	return toSerialize, nil
}

func (o *CloudAuthProvidersPutRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCloudAuthProvidersPutRequest := _CloudAuthProvidersPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudAuthProvidersPutRequest)

	if err != nil {
		return err
	}

	*o = CloudAuthProvidersPutRequest(varCloudAuthProvidersPutRequest)

	return err
}

type NullableCloudAuthProvidersPutRequest struct {
	value *CloudAuthProvidersPutRequest
	isSet bool
}

func (v NullableCloudAuthProvidersPutRequest) Get() *CloudAuthProvidersPutRequest {
	return v.value
}

func (v *NullableCloudAuthProvidersPutRequest) Set(val *CloudAuthProvidersPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAuthProvidersPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAuthProvidersPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAuthProvidersPutRequest(val *CloudAuthProvidersPutRequest) *NullableCloudAuthProvidersPutRequest {
	return &NullableCloudAuthProvidersPutRequest{value: val, isSet: true}
}

func (v NullableCloudAuthProvidersPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAuthProvidersPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

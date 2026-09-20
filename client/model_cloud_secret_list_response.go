package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudSecretListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudSecretListResponse{}

// CloudSecretListResponse struct for CloudSecretListResponse
type CloudSecretListResponse struct {
	Secrets              []CloudSecretListResponseSecretsInner `json:"secrets"`
	AdditionalProperties map[string]interface{}
}

type _CloudSecretListResponse CloudSecretListResponse

// NewCloudSecretListResponse instantiates a new CloudSecretListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSecretListResponse(secrets []CloudSecretListResponseSecretsInner) *CloudSecretListResponse {
	this := CloudSecretListResponse{}
	this.Secrets = secrets
	return &this
}

// NewCloudSecretListResponseWithDefaults instantiates a new CloudSecretListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSecretListResponseWithDefaults() *CloudSecretListResponse {
	this := CloudSecretListResponse{}
	return &this
}

// GetSecrets returns the Secrets field value
func (o *CloudSecretListResponse) GetSecrets() []CloudSecretListResponseSecretsInner {
	if o == nil {
		var ret []CloudSecretListResponseSecretsInner
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *CloudSecretListResponse) GetSecretsOk() ([]CloudSecretListResponseSecretsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *CloudSecretListResponse) SetSecrets(v []CloudSecretListResponseSecretsInner) {
	o.Secrets = v
}

func (o CloudSecretListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudSecretListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secrets"] = o.Secrets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secrets",
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

	varCloudSecretListResponse := _CloudSecretListResponse{}

	err = json.Unmarshal(data, &varCloudSecretListResponse)

	if err != nil {
		return err
	}

	*o = CloudSecretListResponse(varCloudSecretListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "secrets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudSecretListResponse struct {
	value *CloudSecretListResponse
	isSet bool
}

func (v NullableCloudSecretListResponse) Get() *CloudSecretListResponse {
	return v.value
}

func (v *NullableCloudSecretListResponse) Set(val *CloudSecretListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSecretListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSecretListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSecretListResponse(val *CloudSecretListResponse) *NullableCloudSecretListResponse {
	return &NullableCloudSecretListResponse{value: val, isSet: true}
}

func (v NullableCloudSecretListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSecretListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

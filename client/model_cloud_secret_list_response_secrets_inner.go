package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudSecretListResponseSecretsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudSecretListResponseSecretsInner{}

// CloudSecretListResponseSecretsInner struct for CloudSecretListResponseSecretsInner
type CloudSecretListResponseSecretsInner struct {
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
}

type _CloudSecretListResponseSecretsInner CloudSecretListResponseSecretsInner

// NewCloudSecretListResponseSecretsInner instantiates a new CloudSecretListResponseSecretsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSecretListResponseSecretsInner(name string, updatedAt int64) *CloudSecretListResponseSecretsInner {
	this := CloudSecretListResponseSecretsInner{}
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewCloudSecretListResponseSecretsInnerWithDefaults instantiates a new CloudSecretListResponseSecretsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSecretListResponseSecretsInnerWithDefaults() *CloudSecretListResponseSecretsInner {
	this := CloudSecretListResponseSecretsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CloudSecretListResponseSecretsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudSecretListResponseSecretsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudSecretListResponseSecretsInner) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CloudSecretListResponseSecretsInner) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CloudSecretListResponseSecretsInner) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CloudSecretListResponseSecretsInner) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o CloudSecretListResponseSecretsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudSecretListResponseSecretsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *CloudSecretListResponseSecretsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCloudSecretListResponseSecretsInner := _CloudSecretListResponseSecretsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudSecretListResponseSecretsInner)

	if err != nil {
		return err
	}

	*o = CloudSecretListResponseSecretsInner(varCloudSecretListResponseSecretsInner)

	return err
}

type NullableCloudSecretListResponseSecretsInner struct {
	value *CloudSecretListResponseSecretsInner
	isSet bool
}

func (v NullableCloudSecretListResponseSecretsInner) Get() *CloudSecretListResponseSecretsInner {
	return v.value
}

func (v *NullableCloudSecretListResponseSecretsInner) Set(val *CloudSecretListResponseSecretsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSecretListResponseSecretsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSecretListResponseSecretsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSecretListResponseSecretsInner(val *CloudSecretListResponseSecretsInner) *NullableCloudSecretListResponseSecretsInner {
	return &NullableCloudSecretListResponseSecretsInner{value: val, isSet: true}
}

func (v NullableCloudSecretListResponseSecretsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSecretListResponseSecretsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudSecretMutationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudSecretMutationResponse{}

// CloudSecretMutationResponse struct for CloudSecretMutationResponse
type CloudSecretMutationResponse struct {
	Name    string
	Updated *bool
	Deleted *bool
}

type _CloudSecretMutationResponse CloudSecretMutationResponse

// NewCloudSecretMutationResponse instantiates a new CloudSecretMutationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSecretMutationResponse(name string) *CloudSecretMutationResponse {
	this := CloudSecretMutationResponse{}
	return &this
}

// NewCloudSecretMutationResponseWithDefaults instantiates a new CloudSecretMutationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSecretMutationResponseWithDefaults() *CloudSecretMutationResponse {
	this := CloudSecretMutationResponse{}
	return &this
}

// GetName returns the Name field value
func (o *CloudSecretMutationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudSecretMutationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudSecretMutationResponse) SetName(v string) {
	o.Name = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *CloudSecretMutationResponse) GetUpdated() bool {
	if o == nil || IsNil(o.Updated) {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecretMutationResponse) GetUpdatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *CloudSecretMutationResponse) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *CloudSecretMutationResponse) SetUpdated(v bool) {
	o.Updated = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CloudSecretMutationResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecretMutationResponse) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CloudSecretMutationResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CloudSecretMutationResponse) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o CloudSecretMutationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudSecretMutationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

func (o *CloudSecretMutationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCloudSecretMutationResponse := _CloudSecretMutationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudSecretMutationResponse)

	if err != nil {
		return err
	}

	*o = CloudSecretMutationResponse(varCloudSecretMutationResponse)

	return err
}

type NullableCloudSecretMutationResponse struct {
	value *CloudSecretMutationResponse
	isSet bool
}

func (v NullableCloudSecretMutationResponse) Get() *CloudSecretMutationResponse {
	return v.value
}

func (v *NullableCloudSecretMutationResponse) Set(val *CloudSecretMutationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSecretMutationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSecretMutationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSecretMutationResponse(val *CloudSecretMutationResponse) *NullableCloudSecretMutationResponse {
	return &NullableCloudSecretMutationResponse{value: val, isSet: true}
}

func (v NullableCloudSecretMutationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSecretMutationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StaticReleaseCompleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleaseCompleteRequest{}

// StaticReleaseCompleteRequest struct for StaticReleaseCompleteRequest
type StaticReleaseCompleteRequest struct {
	ReleaseToken string `json:"release_token"`
}

type _StaticReleaseCompleteRequest StaticReleaseCompleteRequest

// NewStaticReleaseCompleteRequest instantiates a new StaticReleaseCompleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleaseCompleteRequest(releaseToken string) *StaticReleaseCompleteRequest {
	this := StaticReleaseCompleteRequest{}
	this.ReleaseToken = releaseToken
	return &this
}

// NewStaticReleaseCompleteRequestWithDefaults instantiates a new StaticReleaseCompleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleaseCompleteRequestWithDefaults() *StaticReleaseCompleteRequest {
	this := StaticReleaseCompleteRequest{}
	return &this
}

// GetReleaseToken returns the ReleaseToken field value
func (o *StaticReleaseCompleteRequest) GetReleaseToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseToken
}

// GetReleaseTokenOk returns a tuple with the ReleaseToken field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseCompleteRequest) GetReleaseTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseToken, true
}

// SetReleaseToken sets field value
func (o *StaticReleaseCompleteRequest) SetReleaseToken(v string) {
	o.ReleaseToken = v
}

func (o StaticReleaseCompleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleaseCompleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["release_token"] = o.ReleaseToken
	return toSerialize, nil
}

func (o *StaticReleaseCompleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"release_token",
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

	varStaticReleaseCompleteRequest := _StaticReleaseCompleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticReleaseCompleteRequest)

	if err != nil {
		return err
	}

	*o = StaticReleaseCompleteRequest(varStaticReleaseCompleteRequest)

	return err
}

type NullableStaticReleaseCompleteRequest struct {
	value *StaticReleaseCompleteRequest
	isSet bool
}

func (v NullableStaticReleaseCompleteRequest) Get() *StaticReleaseCompleteRequest {
	return v.value
}

func (v *NullableStaticReleaseCompleteRequest) Set(val *StaticReleaseCompleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleaseCompleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleaseCompleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleaseCompleteRequest(val *StaticReleaseCompleteRequest) *NullableStaticReleaseCompleteRequest {
	return &NullableStaticReleaseCompleteRequest{value: val, isSet: true}
}

func (v NullableStaticReleaseCompleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleaseCompleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

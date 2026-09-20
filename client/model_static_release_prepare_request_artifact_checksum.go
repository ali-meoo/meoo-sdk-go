package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StaticReleasePrepareRequestArtifactChecksum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleasePrepareRequestArtifactChecksum{}

// StaticReleasePrepareRequestArtifactChecksum struct for StaticReleasePrepareRequestArtifactChecksum
type StaticReleasePrepareRequestArtifactChecksum struct {
	Algorithm string `json:"algorithm"`
	// MD5 的 Base64 值
	Value string `json:"value"`
}

type _StaticReleasePrepareRequestArtifactChecksum StaticReleasePrepareRequestArtifactChecksum

// NewStaticReleasePrepareRequestArtifactChecksum instantiates a new StaticReleasePrepareRequestArtifactChecksum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleasePrepareRequestArtifactChecksum(algorithm string, value string) *StaticReleasePrepareRequestArtifactChecksum {
	this := StaticReleasePrepareRequestArtifactChecksum{}
	this.Algorithm = algorithm
	this.Value = value
	return &this
}

// NewStaticReleasePrepareRequestArtifactChecksumWithDefaults instantiates a new StaticReleasePrepareRequestArtifactChecksum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleasePrepareRequestArtifactChecksumWithDefaults() *StaticReleasePrepareRequestArtifactChecksum {
	this := StaticReleasePrepareRequestArtifactChecksum{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *StaticReleasePrepareRequestArtifactChecksum) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifactChecksum) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *StaticReleasePrepareRequestArtifactChecksum) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetValue returns the Value field value
func (o *StaticReleasePrepareRequestArtifactChecksum) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequestArtifactChecksum) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StaticReleasePrepareRequestArtifactChecksum) SetValue(v string) {
	o.Value = v
}

func (o StaticReleasePrepareRequestArtifactChecksum) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleasePrepareRequestArtifactChecksum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *StaticReleasePrepareRequestArtifactChecksum) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algorithm",
		"value",
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

	varStaticReleasePrepareRequestArtifactChecksum := _StaticReleasePrepareRequestArtifactChecksum{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticReleasePrepareRequestArtifactChecksum)

	if err != nil {
		return err
	}

	*o = StaticReleasePrepareRequestArtifactChecksum(varStaticReleasePrepareRequestArtifactChecksum)

	return err
}

type NullableStaticReleasePrepareRequestArtifactChecksum struct {
	value *StaticReleasePrepareRequestArtifactChecksum
	isSet bool
}

func (v NullableStaticReleasePrepareRequestArtifactChecksum) Get() *StaticReleasePrepareRequestArtifactChecksum {
	return v.value
}

func (v *NullableStaticReleasePrepareRequestArtifactChecksum) Set(val *StaticReleasePrepareRequestArtifactChecksum) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleasePrepareRequestArtifactChecksum) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleasePrepareRequestArtifactChecksum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleasePrepareRequestArtifactChecksum(val *StaticReleasePrepareRequestArtifactChecksum) *NullableStaticReleasePrepareRequestArtifactChecksum {
	return &NullableStaticReleasePrepareRequestArtifactChecksum{value: val, isSet: true}
}

func (v NullableStaticReleasePrepareRequestArtifactChecksum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleasePrepareRequestArtifactChecksum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the StaticReleasePrepareRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleasePrepareRequest{}

// StaticReleasePrepareRequest struct for StaticReleasePrepareRequest
type StaticReleasePrepareRequest struct {
	Runtime              string                              `json:"runtime"`
	Artifact             StaticReleasePrepareRequestArtifact `json:"artifact"`
	AdditionalProperties map[string]interface{}
}

type _StaticReleasePrepareRequest StaticReleasePrepareRequest

// NewStaticReleasePrepareRequest instantiates a new StaticReleasePrepareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleasePrepareRequest(runtime string, artifact StaticReleasePrepareRequestArtifact) *StaticReleasePrepareRequest {
	this := StaticReleasePrepareRequest{}
	this.Runtime = runtime
	this.Artifact = artifact
	return &this
}

// NewStaticReleasePrepareRequestWithDefaults instantiates a new StaticReleasePrepareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleasePrepareRequestWithDefaults() *StaticReleasePrepareRequest {
	this := StaticReleasePrepareRequest{}
	return &this
}

// GetRuntime returns the Runtime field value
func (o *StaticReleasePrepareRequest) GetRuntime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequest) GetRuntimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *StaticReleasePrepareRequest) SetRuntime(v string) {
	o.Runtime = v
}

// GetArtifact returns the Artifact field value
func (o *StaticReleasePrepareRequest) GetArtifact() StaticReleasePrepareRequestArtifact {
	if o == nil {
		var ret StaticReleasePrepareRequestArtifact
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareRequest) GetArtifactOk() (*StaticReleasePrepareRequestArtifact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *StaticReleasePrepareRequest) SetArtifact(v StaticReleasePrepareRequestArtifact) {
	o.Artifact = v
}

func (o StaticReleasePrepareRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleasePrepareRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["runtime"] = o.Runtime
	toSerialize["artifact"] = o.Artifact

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StaticReleasePrepareRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"runtime",
		"artifact",
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

	varStaticReleasePrepareRequest := _StaticReleasePrepareRequest{}

	err = json.Unmarshal(data, &varStaticReleasePrepareRequest)

	if err != nil {
		return err
	}

	*o = StaticReleasePrepareRequest(varStaticReleasePrepareRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "runtime")
		delete(additionalProperties, "artifact")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStaticReleasePrepareRequest struct {
	value *StaticReleasePrepareRequest
	isSet bool
}

func (v NullableStaticReleasePrepareRequest) Get() *StaticReleasePrepareRequest {
	return v.value
}

func (v *NullableStaticReleasePrepareRequest) Set(val *StaticReleasePrepareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleasePrepareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleasePrepareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleasePrepareRequest(val *StaticReleasePrepareRequest) *NullableStaticReleasePrepareRequest {
	return &NullableStaticReleasePrepareRequest{value: val, isSet: true}
}

func (v NullableStaticReleasePrepareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleasePrepareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

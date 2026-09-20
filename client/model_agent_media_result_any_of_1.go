package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMediaResultAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMediaResultAnyOf1{}

// AgentMediaResultAnyOf1 struct for AgentMediaResultAnyOf1
type AgentMediaResultAnyOf1 struct {
	Type  string `json:"type"`
	Video Items  `json:"video"`
}

type _AgentMediaResultAnyOf1 AgentMediaResultAnyOf1

// NewAgentMediaResultAnyOf1 instantiates a new AgentMediaResultAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMediaResultAnyOf1(type_ string, video Items) *AgentMediaResultAnyOf1 {
	this := AgentMediaResultAnyOf1{}
	this.Type = type_
	this.Video = video
	return &this
}

// NewAgentMediaResultAnyOf1WithDefaults instantiates a new AgentMediaResultAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMediaResultAnyOf1WithDefaults() *AgentMediaResultAnyOf1 {
	this := AgentMediaResultAnyOf1{}
	return &this
}

// GetType returns the Type field value
func (o *AgentMediaResultAnyOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgentMediaResultAnyOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgentMediaResultAnyOf1) SetType(v string) {
	o.Type = v
}

// GetVideo returns the Video field value
func (o *AgentMediaResultAnyOf1) GetVideo() Items {
	if o == nil {
		var ret Items
		return ret
	}

	return o.Video
}

// GetVideoOk returns a tuple with the Video field value
// and a boolean to check if the value has been set.
func (o *AgentMediaResultAnyOf1) GetVideoOk() (*Items, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Video, true
}

// SetVideo sets field value
func (o *AgentMediaResultAnyOf1) SetVideo(v Items) {
	o.Video = v
}

func (o AgentMediaResultAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMediaResultAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["video"] = o.Video
	return toSerialize, nil
}

func (o *AgentMediaResultAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"video",
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

	varAgentMediaResultAnyOf1 := _AgentMediaResultAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentMediaResultAnyOf1)

	if err != nil {
		return err
	}

	*o = AgentMediaResultAnyOf1(varAgentMediaResultAnyOf1)

	return err
}

type NullableAgentMediaResultAnyOf1 struct {
	value *AgentMediaResultAnyOf1
	isSet bool
}

func (v NullableAgentMediaResultAnyOf1) Get() *AgentMediaResultAnyOf1 {
	return v.value
}

func (v *NullableAgentMediaResultAnyOf1) Set(val *AgentMediaResultAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMediaResultAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMediaResultAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMediaResultAnyOf1(val *AgentMediaResultAnyOf1) *NullableAgentMediaResultAnyOf1 {
	return &NullableAgentMediaResultAnyOf1{value: val, isSet: true}
}

func (v NullableAgentMediaResultAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMediaResultAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

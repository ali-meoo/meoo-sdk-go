package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMediaResultAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMediaResultAnyOf{}

// AgentMediaResultAnyOf struct for AgentMediaResultAnyOf
type AgentMediaResultAnyOf struct {
	Type   string   `json:"type"`
	Images []string `json:"images"`
	// 部分图片因 URL 校验或数量/大小限制未返回。
	Truncated *bool `json:"truncated,omitempty"`
}

type _AgentMediaResultAnyOf AgentMediaResultAnyOf

// NewAgentMediaResultAnyOf instantiates a new AgentMediaResultAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMediaResultAnyOf(type_ string, images []string) *AgentMediaResultAnyOf {
	this := AgentMediaResultAnyOf{}
	this.Type = type_
	this.Images = images
	return &this
}

// NewAgentMediaResultAnyOfWithDefaults instantiates a new AgentMediaResultAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMediaResultAnyOfWithDefaults() *AgentMediaResultAnyOf {
	this := AgentMediaResultAnyOf{}
	return &this
}

// GetType returns the Type field value
func (o *AgentMediaResultAnyOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgentMediaResultAnyOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgentMediaResultAnyOf) SetType(v string) {
	o.Type = v
}

// GetImages returns the Images field value
func (o *AgentMediaResultAnyOf) GetImages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *AgentMediaResultAnyOf) GetImagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *AgentMediaResultAnyOf) SetImages(v []string) {
	o.Images = v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *AgentMediaResultAnyOf) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentMediaResultAnyOf) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *AgentMediaResultAnyOf) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *AgentMediaResultAnyOf) SetTruncated(v bool) {
	o.Truncated = &v
}

func (o AgentMediaResultAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMediaResultAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["images"] = o.Images
	if !IsNil(o.Truncated) {
		toSerialize["truncated"] = o.Truncated
	}
	return toSerialize, nil
}

func (o *AgentMediaResultAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"images",
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

	varAgentMediaResultAnyOf := _AgentMediaResultAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentMediaResultAnyOf)

	if err != nil {
		return err
	}

	*o = AgentMediaResultAnyOf(varAgentMediaResultAnyOf)

	return err
}

type NullableAgentMediaResultAnyOf struct {
	value *AgentMediaResultAnyOf
	isSet bool
}

func (v NullableAgentMediaResultAnyOf) Get() *AgentMediaResultAnyOf {
	return v.value
}

func (v *NullableAgentMediaResultAnyOf) Set(val *AgentMediaResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMediaResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMediaResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMediaResultAnyOf(val *AgentMediaResultAnyOf) *NullableAgentMediaResultAnyOf {
	return &NullableAgentMediaResultAnyOf{value: val, isSet: true}
}

func (v NullableAgentMediaResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMediaResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

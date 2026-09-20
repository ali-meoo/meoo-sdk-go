package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunMediaAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunMediaAttachment{}

// AgentRunMediaAttachment struct for AgentRunMediaAttachment
type AgentRunMediaAttachment struct {
	Kind string `json:"kind"`
	// 可公开访问的 HTTP(S) 地址；服务端会拒绝本地、私网及链路本地地址。
	Url string `json:"url"`
}

type _AgentRunMediaAttachment AgentRunMediaAttachment

// NewAgentRunMediaAttachment instantiates a new AgentRunMediaAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunMediaAttachment(kind string, url string) *AgentRunMediaAttachment {
	this := AgentRunMediaAttachment{}
	this.Kind = kind
	this.Url = url
	return &this
}

// NewAgentRunMediaAttachmentWithDefaults instantiates a new AgentRunMediaAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunMediaAttachmentWithDefaults() *AgentRunMediaAttachment {
	this := AgentRunMediaAttachment{}
	return &this
}

// GetKind returns the Kind field value
func (o *AgentRunMediaAttachment) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AgentRunMediaAttachment) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AgentRunMediaAttachment) SetKind(v string) {
	o.Kind = v
}

// GetUrl returns the Url field value
func (o *AgentRunMediaAttachment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AgentRunMediaAttachment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AgentRunMediaAttachment) SetUrl(v string) {
	o.Url = v
}

func (o AgentRunMediaAttachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunMediaAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *AgentRunMediaAttachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"url",
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

	varAgentRunMediaAttachment := _AgentRunMediaAttachment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRunMediaAttachment)

	if err != nil {
		return err
	}

	*o = AgentRunMediaAttachment(varAgentRunMediaAttachment)

	return err
}

type NullableAgentRunMediaAttachment struct {
	value *AgentRunMediaAttachment
	isSet bool
}

func (v NullableAgentRunMediaAttachment) Get() *AgentRunMediaAttachment {
	return v.value
}

func (v *NullableAgentRunMediaAttachment) Set(val *AgentRunMediaAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunMediaAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunMediaAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunMediaAttachment(val *AgentRunMediaAttachment) *NullableAgentRunMediaAttachment {
	return &NullableAgentRunMediaAttachment{value: val, isSet: true}
}

func (v NullableAgentRunMediaAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunMediaAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

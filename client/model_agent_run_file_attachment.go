package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunFileAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunFileAttachment{}

// AgentRunFileAttachment struct for AgentRunFileAttachment
type AgentRunFileAttachment struct {
	Kind string `json:"kind"`
	// 可供 Meoo 后端下载的公开 HTTP(S) URL；服务端会拒绝本地、私网及链路本地地址。
	Url string `json:"url"`
	// 不含目录路径的文件名；不支持 .zip 和 .skill 技能包。
	Filename string `json:"filename"`
	// 可选 MIME type；例如 application/pdf。
	Type *string `json:"type,omitempty" validate:"regexp=^[A-Za-z0-9][A-Za-z0-9!#$&^_.+-]*_\\/[A-Za-z0-9][A-Za-z0-9!#$&^_.+-]*$"`
}

type _AgentRunFileAttachment AgentRunFileAttachment

// NewAgentRunFileAttachment instantiates a new AgentRunFileAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunFileAttachment(kind string, url string, filename string) *AgentRunFileAttachment {
	this := AgentRunFileAttachment{}
	this.Kind = kind
	this.Url = url
	this.Filename = filename
	return &this
}

// NewAgentRunFileAttachmentWithDefaults instantiates a new AgentRunFileAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunFileAttachmentWithDefaults() *AgentRunFileAttachment {
	this := AgentRunFileAttachment{}
	return &this
}

// GetKind returns the Kind field value
func (o *AgentRunFileAttachment) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AgentRunFileAttachment) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AgentRunFileAttachment) SetKind(v string) {
	o.Kind = v
}

// GetUrl returns the Url field value
func (o *AgentRunFileAttachment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AgentRunFileAttachment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AgentRunFileAttachment) SetUrl(v string) {
	o.Url = v
}

// GetFilename returns the Filename field value
func (o *AgentRunFileAttachment) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *AgentRunFileAttachment) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *AgentRunFileAttachment) SetFilename(v string) {
	o.Filename = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AgentRunFileAttachment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunFileAttachment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AgentRunFileAttachment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AgentRunFileAttachment) SetType(v string) {
	o.Type = &v
}

func (o AgentRunFileAttachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunFileAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["url"] = o.Url
	toSerialize["filename"] = o.Filename
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *AgentRunFileAttachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"url",
		"filename",
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

	varAgentRunFileAttachment := _AgentRunFileAttachment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRunFileAttachment)

	if err != nil {
		return err
	}

	*o = AgentRunFileAttachment(varAgentRunFileAttachment)

	return err
}

type NullableAgentRunFileAttachment struct {
	value *AgentRunFileAttachment
	isSet bool
}

func (v NullableAgentRunFileAttachment) Get() *AgentRunFileAttachment {
	return v.value
}

func (v *NullableAgentRunFileAttachment) Set(val *AgentRunFileAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunFileAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunFileAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunFileAttachment(val *AgentRunFileAttachment) *NullableAgentRunFileAttachment {
	return &NullableAgentRunFileAttachment{value: val, isSet: true}
}

func (v NullableAgentRunFileAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunFileAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

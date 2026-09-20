package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMessage{}

// AgentMessage struct for AgentMessage
type AgentMessage struct {
	MessageId string `json:"message_id"`
	Role      string `json:"role"`
	// 已移除内部指令和工具数据的公开文本。
	Content string `json:"content"`
	// 消息创建时间，Unix 毫秒时间戳。
	CreatedAt int64 `json:"created_at"`
}

type _AgentMessage AgentMessage

// NewAgentMessage instantiates a new AgentMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMessage(messageId string, role string, content string, createdAt int64) *AgentMessage {
	this := AgentMessage{}
	this.MessageId = messageId
	this.Role = role
	this.Content = content
	this.CreatedAt = createdAt
	return &this
}

// NewAgentMessageWithDefaults instantiates a new AgentMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMessageWithDefaults() *AgentMessage {
	this := AgentMessage{}
	return &this
}

// GetMessageId returns the MessageId field value
func (o *AgentMessage) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AgentMessage) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AgentMessage) SetMessageId(v string) {
	o.MessageId = v
}

// GetRole returns the Role field value
func (o *AgentMessage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AgentMessage) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *AgentMessage) SetRole(v string) {
	o.Role = v
}

// GetContent returns the Content field value
func (o *AgentMessage) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AgentMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AgentMessage) SetContent(v string) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AgentMessage) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AgentMessage) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AgentMessage) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o AgentMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_id"] = o.MessageId
	toSerialize["role"] = o.Role
	toSerialize["content"] = o.Content
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *AgentMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_id",
		"role",
		"content",
		"created_at",
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

	varAgentMessage := _AgentMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentMessage)

	if err != nil {
		return err
	}

	*o = AgentMessage(varAgentMessage)

	return err
}

type NullableAgentMessage struct {
	value *AgentMessage
	isSet bool
}

func (v NullableAgentMessage) Get() *AgentMessage {
	return v.value
}

func (v *NullableAgentMessage) Set(val *AgentMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMessage(val *AgentMessage) *NullableAgentMessage {
	return &NullableAgentMessage{value: val, isSet: true}
}

func (v NullableAgentMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

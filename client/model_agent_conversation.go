package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentConversation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentConversation{}

// AgentConversation struct for AgentConversation
type AgentConversation struct {
	ConversationId string `json:"conversation_id"`
	// 会话创建时间，Unix 毫秒时间戳。
	CreatedAt int64 `json:"created_at"`
}

type _AgentConversation AgentConversation

// NewAgentConversation instantiates a new AgentConversation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentConversation(conversationId string, createdAt int64) *AgentConversation {
	this := AgentConversation{}
	this.ConversationId = conversationId
	this.CreatedAt = createdAt
	return &this
}

// NewAgentConversationWithDefaults instantiates a new AgentConversation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentConversationWithDefaults() *AgentConversation {
	this := AgentConversation{}
	return &this
}

// GetConversationId returns the ConversationId field value
func (o *AgentConversation) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AgentConversation) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value
func (o *AgentConversation) SetConversationId(v string) {
	o.ConversationId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AgentConversation) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AgentConversation) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AgentConversation) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o AgentConversation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentConversation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conversation_id"] = o.ConversationId
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *AgentConversation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conversation_id",
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

	varAgentConversation := _AgentConversation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentConversation)

	if err != nil {
		return err
	}

	*o = AgentConversation(varAgentConversation)

	return err
}

type NullableAgentConversation struct {
	value *AgentConversation
	isSet bool
}

func (v NullableAgentConversation) Get() *AgentConversation {
	return v.value
}

func (v *NullableAgentConversation) Set(val *AgentConversation) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentConversation) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentConversation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentConversation(val *AgentConversation) *NullableAgentConversation {
	return &NullableAgentConversation{value: val, isSet: true}
}

func (v NullableAgentConversation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentConversation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

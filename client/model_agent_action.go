package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentAction{}

// AgentAction struct for AgentAction
type AgentAction struct {
	// 短期 HMAC 签名 token；通过 POST agent/action-responses 回复。
	ActionId string `json:"action_id"`
	// 回复形状类别；answers 为问答，confirmation 为 approve/reject 确认决策。
	Kind   string `json:"kind"`
	Prompt string `json:"prompt"`
	// 回复 response 字段需满足的 JSON Schema。
	InputSchema map[string]interface{} `json:"input_schema"`
	// action_id 过期时间，Unix 毫秒时间戳。
	ExpiresAt int64 `json:"expires_at"`
}

type _AgentAction AgentAction

// NewAgentAction instantiates a new AgentAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAction(actionId string, kind string, prompt string, inputSchema map[string]interface{}, expiresAt int64) *AgentAction {
	this := AgentAction{}
	this.ActionId = actionId
	this.Kind = kind
	this.Prompt = prompt
	this.InputSchema = inputSchema
	this.ExpiresAt = expiresAt
	return &this
}

// NewAgentActionWithDefaults instantiates a new AgentAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentActionWithDefaults() *AgentAction {
	this := AgentAction{}
	return &this
}

// GetActionId returns the ActionId field value
func (o *AgentAction) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *AgentAction) SetActionId(v string) {
	o.ActionId = v
}

// GetKind returns the Kind field value
func (o *AgentAction) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AgentAction) SetKind(v string) {
	o.Kind = v
}

// GetPrompt returns the Prompt field value
func (o *AgentAction) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *AgentAction) SetPrompt(v string) {
	o.Prompt = v
}

// GetInputSchema returns the InputSchema field value
func (o *AgentAction) GetInputSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetInputSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.InputSchema, true
}

// SetInputSchema sets field value
func (o *AgentAction) SetInputSchema(v map[string]interface{}) {
	o.InputSchema = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *AgentAction) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *AgentAction) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

func (o AgentAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action_id"] = o.ActionId
	toSerialize["kind"] = o.Kind
	toSerialize["prompt"] = o.Prompt
	toSerialize["input_schema"] = o.InputSchema
	toSerialize["expires_at"] = o.ExpiresAt
	return toSerialize, nil
}

func (o *AgentAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_id",
		"kind",
		"prompt",
		"input_schema",
		"expires_at",
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

	varAgentAction := _AgentAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentAction)

	if err != nil {
		return err
	}

	*o = AgentAction(varAgentAction)

	return err
}

type NullableAgentAction struct {
	value *AgentAction
	isSet bool
}

func (v NullableAgentAction) Get() *AgentAction {
	return v.value
}

func (v *NullableAgentAction) Set(val *AgentAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentAction(val *AgentAction) *NullableAgentAction {
	return &NullableAgentAction{value: val, isSet: true}
}

func (v NullableAgentAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

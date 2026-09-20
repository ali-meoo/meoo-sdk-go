package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentActionResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentActionResponseRequest{}

// AgentActionResponseRequest struct for AgentActionResponseRequest
type AgentActionResponseRequest struct {
	// run.input_required 事件返回的短期 HMAC 签名 token；调用方不得解析或修改。
	ActionId string                             `json:"action_id"`
	Response AgentActionResponseRequestResponse `json:"response"`
}

type _AgentActionResponseRequest AgentActionResponseRequest

// NewAgentActionResponseRequest instantiates a new AgentActionResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentActionResponseRequest(actionId string, response AgentActionResponseRequestResponse) *AgentActionResponseRequest {
	this := AgentActionResponseRequest{}
	this.ActionId = actionId
	this.Response = response
	return &this
}

// NewAgentActionResponseRequestWithDefaults instantiates a new AgentActionResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentActionResponseRequestWithDefaults() *AgentActionResponseRequest {
	this := AgentActionResponseRequest{}
	return &this
}

// GetActionId returns the ActionId field value
func (o *AgentActionResponseRequest) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *AgentActionResponseRequest) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *AgentActionResponseRequest) SetActionId(v string) {
	o.ActionId = v
}

// GetResponse returns the Response field value
func (o *AgentActionResponseRequest) GetResponse() AgentActionResponseRequestResponse {
	if o == nil {
		var ret AgentActionResponseRequestResponse
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *AgentActionResponseRequest) GetResponseOk() (*AgentActionResponseRequestResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *AgentActionResponseRequest) SetResponse(v AgentActionResponseRequestResponse) {
	o.Response = v
}

func (o AgentActionResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentActionResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action_id"] = o.ActionId
	toSerialize["response"] = o.Response
	return toSerialize, nil
}

func (o *AgentActionResponseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_id",
		"response",
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

	varAgentActionResponseRequest := _AgentActionResponseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentActionResponseRequest)

	if err != nil {
		return err
	}

	*o = AgentActionResponseRequest(varAgentActionResponseRequest)

	return err
}

type NullableAgentActionResponseRequest struct {
	value *AgentActionResponseRequest
	isSet bool
}

func (v NullableAgentActionResponseRequest) Get() *AgentActionResponseRequest {
	return v.value
}

func (v *NullableAgentActionResponseRequest) Set(val *AgentActionResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentActionResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentActionResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentActionResponseRequest(val *AgentActionResponseRequest) *NullableAgentActionResponseRequest {
	return &NullableAgentActionResponseRequest{value: val, isSet: true}
}

func (v NullableAgentActionResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentActionResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentAnswersActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentAnswersActionResponse{}

// AgentAnswersActionResponse struct for AgentAnswersActionResponse
type AgentAnswersActionResponse struct {
	// key 和 value 均先去除首尾空白，再按长度与数量约束校验并传递；trim 后为空或 key 冲突时拒绝请求。
	Answers map[string]string `json:"answers"`
}

type _AgentAnswersActionResponse AgentAnswersActionResponse

// NewAgentAnswersActionResponse instantiates a new AgentAnswersActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAnswersActionResponse(answers map[string]string) *AgentAnswersActionResponse {
	this := AgentAnswersActionResponse{}
	this.Answers = answers
	return &this
}

// NewAgentAnswersActionResponseWithDefaults instantiates a new AgentAnswersActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentAnswersActionResponseWithDefaults() *AgentAnswersActionResponse {
	this := AgentAnswersActionResponse{}
	return &this
}

// GetAnswers returns the Answers field value
func (o *AgentAnswersActionResponse) GetAnswers() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value
// and a boolean to check if the value has been set.
func (o *AgentAnswersActionResponse) GetAnswersOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Answers, true
}

// SetAnswers sets field value
func (o *AgentAnswersActionResponse) SetAnswers(v map[string]string) {
	o.Answers = v
}

func (o AgentAnswersActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentAnswersActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answers"] = o.Answers
	return toSerialize, nil
}

func (o *AgentAnswersActionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"answers",
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

	varAgentAnswersActionResponse := _AgentAnswersActionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentAnswersActionResponse)

	if err != nil {
		return err
	}

	*o = AgentAnswersActionResponse(varAgentAnswersActionResponse)

	return err
}

type NullableAgentAnswersActionResponse struct {
	value *AgentAnswersActionResponse
	isSet bool
}

func (v NullableAgentAnswersActionResponse) Get() *AgentAnswersActionResponse {
	return v.value
}

func (v *NullableAgentAnswersActionResponse) Set(val *AgentAnswersActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentAnswersActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentAnswersActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentAnswersActionResponse(val *AgentAnswersActionResponse) *NullableAgentAnswersActionResponse {
	return &NullableAgentAnswersActionResponse{value: val, isSet: true}
}

func (v NullableAgentAnswersActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentAnswersActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

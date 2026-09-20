package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentConfirmationActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentConfirmationActionResponse{}

// AgentConfirmationActionResponse struct for AgentConfirmationActionResponse
type AgentConfirmationActionResponse struct {
	// approve 表示同意并继续；reject 表示拒绝。
	Decision string `json:"decision"`
	// 可选补充意见；仅当 action.input_schema 声明 comment 属性时才可传。
	Comment *string `json:"comment,omitempty"`
}

type _AgentConfirmationActionResponse AgentConfirmationActionResponse

// NewAgentConfirmationActionResponse instantiates a new AgentConfirmationActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentConfirmationActionResponse(decision string) *AgentConfirmationActionResponse {
	this := AgentConfirmationActionResponse{}
	this.Decision = decision
	return &this
}

// NewAgentConfirmationActionResponseWithDefaults instantiates a new AgentConfirmationActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentConfirmationActionResponseWithDefaults() *AgentConfirmationActionResponse {
	this := AgentConfirmationActionResponse{}
	return &this
}

// GetDecision returns the Decision field value
func (o *AgentConfirmationActionResponse) GetDecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *AgentConfirmationActionResponse) GetDecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *AgentConfirmationActionResponse) SetDecision(v string) {
	o.Decision = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AgentConfirmationActionResponse) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentConfirmationActionResponse) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AgentConfirmationActionResponse) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AgentConfirmationActionResponse) SetComment(v string) {
	o.Comment = &v
}

func (o AgentConfirmationActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentConfirmationActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["decision"] = o.Decision
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

func (o *AgentConfirmationActionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"decision",
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

	varAgentConfirmationActionResponse := _AgentConfirmationActionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentConfirmationActionResponse)

	if err != nil {
		return err
	}

	*o = AgentConfirmationActionResponse(varAgentConfirmationActionResponse)

	return err
}

type NullableAgentConfirmationActionResponse struct {
	value *AgentConfirmationActionResponse
	isSet bool
}

func (v NullableAgentConfirmationActionResponse) Get() *AgentConfirmationActionResponse {
	return v.value
}

func (v *NullableAgentConfirmationActionResponse) Set(val *AgentConfirmationActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentConfirmationActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentConfirmationActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentConfirmationActionResponse(val *AgentConfirmationActionResponse) *NullableAgentConfirmationActionResponse {
	return &NullableAgentConfirmationActionResponse{value: val, isSet: true}
}

func (v NullableAgentConfirmationActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentConfirmationActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

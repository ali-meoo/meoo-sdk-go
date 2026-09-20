package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMessageListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMessageListResponse{}

// AgentMessageListResponse struct for AgentMessageListResponse
type AgentMessageListResponse struct {
	Messages []AgentMessage `json:"messages"`
	// 仅在还有下一页时返回。
	NextPageToken *string `json:"next_page_token,omitempty"`
}

type _AgentMessageListResponse AgentMessageListResponse

// NewAgentMessageListResponse instantiates a new AgentMessageListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMessageListResponse(messages []AgentMessage) *AgentMessageListResponse {
	this := AgentMessageListResponse{}
	this.Messages = messages
	return &this
}

// NewAgentMessageListResponseWithDefaults instantiates a new AgentMessageListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMessageListResponseWithDefaults() *AgentMessageListResponse {
	this := AgentMessageListResponse{}
	return &this
}

// GetMessages returns the Messages field value
func (o *AgentMessageListResponse) GetMessages() []AgentMessage {
	if o == nil {
		var ret []AgentMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *AgentMessageListResponse) GetMessagesOk() ([]AgentMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *AgentMessageListResponse) SetMessages(v []AgentMessage) {
	o.Messages = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AgentMessageListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentMessageListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AgentMessageListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AgentMessageListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o AgentMessageListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMessageListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return toSerialize, nil
}

func (o *AgentMessageListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
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

	varAgentMessageListResponse := _AgentMessageListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentMessageListResponse)

	if err != nil {
		return err
	}

	*o = AgentMessageListResponse(varAgentMessageListResponse)

	return err
}

type NullableAgentMessageListResponse struct {
	value *AgentMessageListResponse
	isSet bool
}

func (v NullableAgentMessageListResponse) Get() *AgentMessageListResponse {
	return v.value
}

func (v *NullableAgentMessageListResponse) Set(val *AgentMessageListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMessageListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMessageListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMessageListResponse(val *AgentMessageListResponse) *NullableAgentMessageListResponse {
	return &NullableAgentMessageListResponse{value: val, isSet: true}
}

func (v NullableAgentMessageListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMessageListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

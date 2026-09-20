package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentConversationListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentConversationListResponse{}

// AgentConversationListResponse struct for AgentConversationListResponse
type AgentConversationListResponse struct {
	Conversations []AgentConversation `json:"conversations"`
	// 仅在还有下一页时返回。
	NextPageToken *string `json:"next_page_token,omitempty"`
}

type _AgentConversationListResponse AgentConversationListResponse

// NewAgentConversationListResponse instantiates a new AgentConversationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentConversationListResponse(conversations []AgentConversation) *AgentConversationListResponse {
	this := AgentConversationListResponse{}
	this.Conversations = conversations
	return &this
}

// NewAgentConversationListResponseWithDefaults instantiates a new AgentConversationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentConversationListResponseWithDefaults() *AgentConversationListResponse {
	this := AgentConversationListResponse{}
	return &this
}

// GetConversations returns the Conversations field value
func (o *AgentConversationListResponse) GetConversations() []AgentConversation {
	if o == nil {
		var ret []AgentConversation
		return ret
	}

	return o.Conversations
}

// GetConversationsOk returns a tuple with the Conversations field value
// and a boolean to check if the value has been set.
func (o *AgentConversationListResponse) GetConversationsOk() ([]AgentConversation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conversations, true
}

// SetConversations sets field value
func (o *AgentConversationListResponse) SetConversations(v []AgentConversation) {
	o.Conversations = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AgentConversationListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentConversationListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AgentConversationListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AgentConversationListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o AgentConversationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentConversationListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conversations"] = o.Conversations
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return toSerialize, nil
}

func (o *AgentConversationListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conversations",
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

	varAgentConversationListResponse := _AgentConversationListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentConversationListResponse)

	if err != nil {
		return err
	}

	*o = AgentConversationListResponse(varAgentConversationListResponse)

	return err
}

type NullableAgentConversationListResponse struct {
	value *AgentConversationListResponse
	isSet bool
}

func (v NullableAgentConversationListResponse) Get() *AgentConversationListResponse {
	return v.value
}

func (v *NullableAgentConversationListResponse) Set(val *AgentConversationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentConversationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentConversationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentConversationListResponse(val *AgentConversationListResponse) *NullableAgentConversationListResponse {
	return &NullableAgentConversationListResponse{value: val, isSet: true}
}

func (v NullableAgentConversationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentConversationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
)

// checks if the AgentSnapshotToolCallInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentSnapshotToolCallInput{}

// AgentSnapshotToolCallInput Write/Edit 的白名单完整输入；Read 内容与 Skill 原始参数不会出现。
type AgentSnapshotToolCallInput struct {
	Content    *string `json:"content,omitempty"`
	OldString  *string `json:"old_string,omitempty"`
	NewString  *string `json:"new_string,omitempty"`
	ReplaceAll *bool   `json:"replace_all,omitempty"`
}

// NewAgentSnapshotToolCallInput instantiates a new AgentSnapshotToolCallInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentSnapshotToolCallInput() *AgentSnapshotToolCallInput {
	this := AgentSnapshotToolCallInput{}
	return &this
}

// NewAgentSnapshotToolCallInputWithDefaults instantiates a new AgentSnapshotToolCallInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentSnapshotToolCallInputWithDefaults() *AgentSnapshotToolCallInput {
	this := AgentSnapshotToolCallInput{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AgentSnapshotToolCallInput) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCallInput) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AgentSnapshotToolCallInput) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AgentSnapshotToolCallInput) SetContent(v string) {
	o.Content = &v
}

// GetOldString returns the OldString field value if set, zero value otherwise.
func (o *AgentSnapshotToolCallInput) GetOldString() string {
	if o == nil || IsNil(o.OldString) {
		var ret string
		return ret
	}
	return *o.OldString
}

// GetOldStringOk returns a tuple with the OldString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCallInput) GetOldStringOk() (*string, bool) {
	if o == nil || IsNil(o.OldString) {
		return nil, false
	}
	return o.OldString, true
}

// HasOldString returns a boolean if a field has been set.
func (o *AgentSnapshotToolCallInput) HasOldString() bool {
	if o != nil && !IsNil(o.OldString) {
		return true
	}

	return false
}

// SetOldString gets a reference to the given string and assigns it to the OldString field.
func (o *AgentSnapshotToolCallInput) SetOldString(v string) {
	o.OldString = &v
}

// GetNewString returns the NewString field value if set, zero value otherwise.
func (o *AgentSnapshotToolCallInput) GetNewString() string {
	if o == nil || IsNil(o.NewString) {
		var ret string
		return ret
	}
	return *o.NewString
}

// GetNewStringOk returns a tuple with the NewString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCallInput) GetNewStringOk() (*string, bool) {
	if o == nil || IsNil(o.NewString) {
		return nil, false
	}
	return o.NewString, true
}

// HasNewString returns a boolean if a field has been set.
func (o *AgentSnapshotToolCallInput) HasNewString() bool {
	if o != nil && !IsNil(o.NewString) {
		return true
	}

	return false
}

// SetNewString gets a reference to the given string and assigns it to the NewString field.
func (o *AgentSnapshotToolCallInput) SetNewString(v string) {
	o.NewString = &v
}

// GetReplaceAll returns the ReplaceAll field value if set, zero value otherwise.
func (o *AgentSnapshotToolCallInput) GetReplaceAll() bool {
	if o == nil || IsNil(o.ReplaceAll) {
		var ret bool
		return ret
	}
	return *o.ReplaceAll
}

// GetReplaceAllOk returns a tuple with the ReplaceAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSnapshotToolCallInput) GetReplaceAllOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceAll) {
		return nil, false
	}
	return o.ReplaceAll, true
}

// HasReplaceAll returns a boolean if a field has been set.
func (o *AgentSnapshotToolCallInput) HasReplaceAll() bool {
	if o != nil && !IsNil(o.ReplaceAll) {
		return true
	}

	return false
}

// SetReplaceAll gets a reference to the given bool and assigns it to the ReplaceAll field.
func (o *AgentSnapshotToolCallInput) SetReplaceAll(v bool) {
	o.ReplaceAll = &v
}

func (o AgentSnapshotToolCallInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentSnapshotToolCallInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.OldString) {
		toSerialize["old_string"] = o.OldString
	}
	if !IsNil(o.NewString) {
		toSerialize["new_string"] = o.NewString
	}
	if !IsNil(o.ReplaceAll) {
		toSerialize["replace_all"] = o.ReplaceAll
	}
	return toSerialize, nil
}

type NullableAgentSnapshotToolCallInput struct {
	value *AgentSnapshotToolCallInput
	isSet bool
}

func (v NullableAgentSnapshotToolCallInput) Get() *AgentSnapshotToolCallInput {
	return v.value
}

func (v *NullableAgentSnapshotToolCallInput) Set(val *AgentSnapshotToolCallInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSnapshotToolCallInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSnapshotToolCallInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSnapshotToolCallInput(val *AgentSnapshotToolCallInput) *NullableAgentSnapshotToolCallInput {
	return &NullableAgentSnapshotToolCallInput{value: val, isSet: true}
}

func (v NullableAgentSnapshotToolCallInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSnapshotToolCallInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

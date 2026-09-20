package client

import (
	"encoding/json"
)

// checks if the AgentToolInputSnapshotEventDataInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolInputSnapshotEventDataInput{}

// AgentToolInputSnapshotEventDataInput 当前已公开的输入全量状态。Write 使用 content；Edit 使用 old_string、 new_string，并可在最终快照携带 replace_all。
type AgentToolInputSnapshotEventDataInput struct {
	Content    *string `json:"content,omitempty"`
	OldString  *string `json:"old_string,omitempty"`
	NewString  *string `json:"new_string,omitempty"`
	ReplaceAll *bool   `json:"replace_all,omitempty"`
}

// NewAgentToolInputSnapshotEventDataInput instantiates a new AgentToolInputSnapshotEventDataInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolInputSnapshotEventDataInput() *AgentToolInputSnapshotEventDataInput {
	this := AgentToolInputSnapshotEventDataInput{}
	return &this
}

// NewAgentToolInputSnapshotEventDataInputWithDefaults instantiates a new AgentToolInputSnapshotEventDataInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolInputSnapshotEventDataInputWithDefaults() *AgentToolInputSnapshotEventDataInput {
	this := AgentToolInputSnapshotEventDataInput{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AgentToolInputSnapshotEventDataInput) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventDataInput) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AgentToolInputSnapshotEventDataInput) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AgentToolInputSnapshotEventDataInput) SetContent(v string) {
	o.Content = &v
}

// GetOldString returns the OldString field value if set, zero value otherwise.
func (o *AgentToolInputSnapshotEventDataInput) GetOldString() string {
	if o == nil || IsNil(o.OldString) {
		var ret string
		return ret
	}
	return *o.OldString
}

// GetOldStringOk returns a tuple with the OldString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventDataInput) GetOldStringOk() (*string, bool) {
	if o == nil || IsNil(o.OldString) {
		return nil, false
	}
	return o.OldString, true
}

// HasOldString returns a boolean if a field has been set.
func (o *AgentToolInputSnapshotEventDataInput) HasOldString() bool {
	if o != nil && !IsNil(o.OldString) {
		return true
	}

	return false
}

// SetOldString gets a reference to the given string and assigns it to the OldString field.
func (o *AgentToolInputSnapshotEventDataInput) SetOldString(v string) {
	o.OldString = &v
}

// GetNewString returns the NewString field value if set, zero value otherwise.
func (o *AgentToolInputSnapshotEventDataInput) GetNewString() string {
	if o == nil || IsNil(o.NewString) {
		var ret string
		return ret
	}
	return *o.NewString
}

// GetNewStringOk returns a tuple with the NewString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventDataInput) GetNewStringOk() (*string, bool) {
	if o == nil || IsNil(o.NewString) {
		return nil, false
	}
	return o.NewString, true
}

// HasNewString returns a boolean if a field has been set.
func (o *AgentToolInputSnapshotEventDataInput) HasNewString() bool {
	if o != nil && !IsNil(o.NewString) {
		return true
	}

	return false
}

// SetNewString gets a reference to the given string and assigns it to the NewString field.
func (o *AgentToolInputSnapshotEventDataInput) SetNewString(v string) {
	o.NewString = &v
}

// GetReplaceAll returns the ReplaceAll field value if set, zero value otherwise.
func (o *AgentToolInputSnapshotEventDataInput) GetReplaceAll() bool {
	if o == nil || IsNil(o.ReplaceAll) {
		var ret bool
		return ret
	}
	return *o.ReplaceAll
}

// GetReplaceAllOk returns a tuple with the ReplaceAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEventDataInput) GetReplaceAllOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceAll) {
		return nil, false
	}
	return o.ReplaceAll, true
}

// HasReplaceAll returns a boolean if a field has been set.
func (o *AgentToolInputSnapshotEventDataInput) HasReplaceAll() bool {
	if o != nil && !IsNil(o.ReplaceAll) {
		return true
	}

	return false
}

// SetReplaceAll gets a reference to the given bool and assigns it to the ReplaceAll field.
func (o *AgentToolInputSnapshotEventDataInput) SetReplaceAll(v bool) {
	o.ReplaceAll = &v
}

func (o AgentToolInputSnapshotEventDataInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolInputSnapshotEventDataInput) ToMap() (map[string]interface{}, error) {
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

type NullableAgentToolInputSnapshotEventDataInput struct {
	value *AgentToolInputSnapshotEventDataInput
	isSet bool
}

func (v NullableAgentToolInputSnapshotEventDataInput) Get() *AgentToolInputSnapshotEventDataInput {
	return v.value
}

func (v *NullableAgentToolInputSnapshotEventDataInput) Set(val *AgentToolInputSnapshotEventDataInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolInputSnapshotEventDataInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolInputSnapshotEventDataInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolInputSnapshotEventDataInput(val *AgentToolInputSnapshotEventDataInput) *NullableAgentToolInputSnapshotEventDataInput {
	return &NullableAgentToolInputSnapshotEventDataInput{value: val, isSet: true}
}

func (v NullableAgentToolInputSnapshotEventDataInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolInputSnapshotEventDataInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

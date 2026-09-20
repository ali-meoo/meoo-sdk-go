package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRun{}

// AgentRun struct for AgentRun
type AgentRun struct {
	ProjectId      string         `json:"project_id"`
	RunId          string         `json:"run_id"`
	ConversationId string         `json:"conversation_id"`
	Status         AgentRunStatus `json:"status"`
}

type _AgentRun AgentRun

// NewAgentRun instantiates a new AgentRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRun(projectId string, runId string, conversationId string, status AgentRunStatus) *AgentRun {
	this := AgentRun{}
	this.ProjectId = projectId
	this.RunId = runId
	this.ConversationId = conversationId
	this.Status = status
	return &this
}

// NewAgentRunWithDefaults instantiates a new AgentRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunWithDefaults() *AgentRun {
	this := AgentRun{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *AgentRun) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AgentRun) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AgentRun) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRunId returns the RunId field value
func (o *AgentRun) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentRun) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentRun) SetRunId(v string) {
	o.RunId = v
}

// GetConversationId returns the ConversationId field value
func (o *AgentRun) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AgentRun) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value
func (o *AgentRun) SetConversationId(v string) {
	o.ConversationId = v
}

// GetStatus returns the Status field value
func (o *AgentRun) GetStatus() AgentRunStatus {
	if o == nil {
		var ret AgentRunStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentRun) GetStatusOk() (*AgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentRun) SetStatus(v AgentRunStatus) {
	o.Status = v
}

func (o AgentRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["run_id"] = o.RunId
	toSerialize["conversation_id"] = o.ConversationId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *AgentRun) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
		"run_id",
		"conversation_id",
		"status",
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

	varAgentRun := _AgentRun{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRun)

	if err != nil {
		return err
	}

	*o = AgentRun(varAgentRun)

	return err
}

type NullableAgentRun struct {
	value *AgentRun
	isSet bool
}

func (v NullableAgentRun) Get() *AgentRun {
	return v.value
}

func (v *NullableAgentRun) Set(val *AgentRun) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRun) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRun(val *AgentRun) *NullableAgentRun {
	return &NullableAgentRun{value: val, isSet: true}
}

func (v NullableAgentRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

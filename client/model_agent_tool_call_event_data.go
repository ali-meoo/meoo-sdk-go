package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentToolCallEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolCallEventData{}

// AgentToolCallEventData struct for AgentToolCallEventData
type AgentToolCallEventData struct {
	RunId string `json:"run_id"`
	// 工具调用 ID；同一次调用的 started 与 completed 共享该 ID。
	ToolCallId string             `json:"tool_call_id"`
	ToolName   string             `json:"tool_name"`
	Status     string             `json:"status"`
	Phase      AgentProgressPhase `json:"phase"`
	// 可安全展示的进度文案；不含工具参数、路径或原始执行结果。
	Message string `json:"message"`
}

type _AgentToolCallEventData AgentToolCallEventData

// NewAgentToolCallEventData instantiates a new AgentToolCallEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolCallEventData(runId string, toolCallId string, toolName string, status string, phase AgentProgressPhase, message string) *AgentToolCallEventData {
	this := AgentToolCallEventData{}
	this.RunId = runId
	this.ToolCallId = toolCallId
	this.ToolName = toolName
	this.Status = status
	this.Phase = phase
	this.Message = message
	return &this
}

// NewAgentToolCallEventDataWithDefaults instantiates a new AgentToolCallEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolCallEventDataWithDefaults() *AgentToolCallEventData {
	this := AgentToolCallEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentToolCallEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentToolCallEventData) SetRunId(v string) {
	o.RunId = v
}

// GetToolCallId returns the ToolCallId field value
func (o *AgentToolCallEventData) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEventData) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value
func (o *AgentToolCallEventData) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetToolName returns the ToolName field value
func (o *AgentToolCallEventData) GetToolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEventData) GetToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolName, true
}

// SetToolName sets field value
func (o *AgentToolCallEventData) SetToolName(v string) {
	o.ToolName = v
}

// GetStatus returns the Status field value
func (o *AgentToolCallEventData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEventData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentToolCallEventData) SetStatus(v string) {
	o.Status = v
}

// GetPhase returns the Phase field value
func (o *AgentToolCallEventData) GetPhase() AgentProgressPhase {
	if o == nil {
		var ret AgentProgressPhase
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEventData) GetPhaseOk() (*AgentProgressPhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *AgentToolCallEventData) SetPhase(v AgentProgressPhase) {
	o.Phase = v
}

// GetMessage returns the Message field value
func (o *AgentToolCallEventData) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEventData) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AgentToolCallEventData) SetMessage(v string) {
	o.Message = v
}

func (o AgentToolCallEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolCallEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["tool_call_id"] = o.ToolCallId
	toSerialize["tool_name"] = o.ToolName
	toSerialize["status"] = o.Status
	toSerialize["phase"] = o.Phase
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AgentToolCallEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
		"tool_call_id",
		"tool_name",
		"status",
		"phase",
		"message",
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

	varAgentToolCallEventData := _AgentToolCallEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentToolCallEventData)

	if err != nil {
		return err
	}

	*o = AgentToolCallEventData(varAgentToolCallEventData)

	return err
}

type NullableAgentToolCallEventData struct {
	value *AgentToolCallEventData
	isSet bool
}

func (v NullableAgentToolCallEventData) Get() *AgentToolCallEventData {
	return v.value
}

func (v *NullableAgentToolCallEventData) Set(val *AgentToolCallEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolCallEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolCallEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolCallEventData(val *AgentToolCallEventData) *NullableAgentToolCallEventData {
	return &NullableAgentToolCallEventData{value: val, isSet: true}
}

func (v NullableAgentToolCallEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolCallEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

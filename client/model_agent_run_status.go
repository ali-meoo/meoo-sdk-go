package client

import (
	"encoding/json"
	"fmt"
)

// AgentRunStatus the model 'AgentRunStatus'
type AgentRunStatus string

// List of AgentRunStatus
const (
	AGENTRUNSTATUS_SUBMITTED      AgentRunStatus = "submitted"
	AGENTRUNSTATUS_WORKING        AgentRunStatus = "working"
	AGENTRUNSTATUS_INPUT_REQUIRED AgentRunStatus = "input_required"
	AGENTRUNSTATUS_COMPLETED      AgentRunStatus = "completed"
	AGENTRUNSTATUS_FAILED         AgentRunStatus = "failed"
	AGENTRUNSTATUS_CANCELED       AgentRunStatus = "canceled"
	AGENTRUNSTATUS_INTERRUPTED    AgentRunStatus = "interrupted"
)

// All allowed values of AgentRunStatus enum
var AllowedAgentRunStatusEnumValues = []AgentRunStatus{
	"submitted",
	"working",
	"input_required",
	"completed",
	"failed",
	"canceled",
	"interrupted",
}

func (v *AgentRunStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentRunStatus(value)
	for _, existing := range AllowedAgentRunStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	// 前向兼容：未知枚举值原样保留而非报错（契约要求未知取值按非终态/进行中兜底）。
	*v = enumTypeValue
	return nil
}

// NewAgentRunStatusFromValue returns a pointer to a valid AgentRunStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentRunStatusFromValue(v string) (*AgentRunStatus, error) {
	ev := AgentRunStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentRunStatus: valid values are %v", v, AllowedAgentRunStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentRunStatus) IsValid() bool {
	for _, existing := range AllowedAgentRunStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentRunStatus value
func (v AgentRunStatus) Ptr() *AgentRunStatus {
	return &v
}

type NullableAgentRunStatus struct {
	value *AgentRunStatus
	isSet bool
}

func (v NullableAgentRunStatus) Get() *AgentRunStatus {
	return v.value
}

func (v *NullableAgentRunStatus) Set(val *AgentRunStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunStatus(val *AgentRunStatus) *NullableAgentRunStatus {
	return &NullableAgentRunStatus{value: val, isSet: true}
}

func (v NullableAgentRunStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// AgentProgressPhase 公开进度阶段；客户端遇到未知取值时一律视为进行中。
type AgentProgressPhase string

// List of AgentProgressPhase
const (
	AGENTPROGRESSPHASE_ANALYZING    AgentProgressPhase = "analyzing"
	AGENTPROGRESSPHASE_BUILDING     AgentProgressPhase = "building"
	AGENTPROGRESSPHASE_VALIDATING   AgentProgressPhase = "validating"
	AGENTPROGRESSPHASE_COORDINATING AgentProgressPhase = "coordinating"
	AGENTPROGRESSPHASE_FINALIZING   AgentProgressPhase = "finalizing"
)

// All allowed values of AgentProgressPhase enum
var AllowedAgentProgressPhaseEnumValues = []AgentProgressPhase{
	"analyzing",
	"building",
	"validating",
	"coordinating",
	"finalizing",
}

func (v *AgentProgressPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentProgressPhase(value)
	for _, existing := range AllowedAgentProgressPhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	// 前向兼容：未知枚举值原样保留而非报错（契约要求未知取值按非终态/进行中兜底）。
	*v = enumTypeValue
	return nil
}

// NewAgentProgressPhaseFromValue returns a pointer to a valid AgentProgressPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentProgressPhaseFromValue(v string) (*AgentProgressPhase, error) {
	ev := AgentProgressPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentProgressPhase: valid values are %v", v, AllowedAgentProgressPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentProgressPhase) IsValid() bool {
	for _, existing := range AllowedAgentProgressPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentProgressPhase value
func (v AgentProgressPhase) Ptr() *AgentProgressPhase {
	return &v
}

type NullableAgentProgressPhase struct {
	value *AgentProgressPhase
	isSet bool
}

func (v NullableAgentProgressPhase) Get() *AgentProgressPhase {
	return v.value
}

func (v *NullableAgentProgressPhase) Set(val *AgentProgressPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentProgressPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentProgressPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentProgressPhase(val *AgentProgressPhase) *NullableAgentProgressPhase {
	return &NullableAgentProgressPhase{value: val, isSet: true}
}

func (v NullableAgentProgressPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentProgressPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

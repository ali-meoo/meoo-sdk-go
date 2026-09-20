package client

import (
	"encoding/json"
	"fmt"
)

// AgentMediaResult struct for AgentMediaResult
type AgentMediaResult struct {
	AgentMediaResultAnyOf  *AgentMediaResultAnyOf
	AgentMediaResultAnyOf1 *AgentMediaResultAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AgentMediaResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AgentMediaResultAnyOf
	err = json.Unmarshal(data, &dst.AgentMediaResultAnyOf)
	if err == nil {
		jsonAgentMediaResultAnyOf, _ := json.Marshal(dst.AgentMediaResultAnyOf)
		if string(jsonAgentMediaResultAnyOf) == "{}" { // empty struct
			dst.AgentMediaResultAnyOf = nil
		} else {
			return nil // data stored in dst.AgentMediaResultAnyOf, return on the first match
		}
	} else {
		dst.AgentMediaResultAnyOf = nil
	}

	// try to unmarshal JSON data into AgentMediaResultAnyOf1
	err = json.Unmarshal(data, &dst.AgentMediaResultAnyOf1)
	if err == nil {
		jsonAgentMediaResultAnyOf1, _ := json.Marshal(dst.AgentMediaResultAnyOf1)
		if string(jsonAgentMediaResultAnyOf1) == "{}" { // empty struct
			dst.AgentMediaResultAnyOf1 = nil
		} else {
			return nil // data stored in dst.AgentMediaResultAnyOf1, return on the first match
		}
	} else {
		dst.AgentMediaResultAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AgentMediaResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentMediaResult) MarshalJSON() ([]byte, error) {
	if src.AgentMediaResultAnyOf != nil {
		return json.Marshal(&src.AgentMediaResultAnyOf)
	}

	if src.AgentMediaResultAnyOf1 != nil {
		return json.Marshal(&src.AgentMediaResultAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAgentMediaResult struct {
	value *AgentMediaResult
	isSet bool
}

func (v NullableAgentMediaResult) Get() *AgentMediaResult {
	return v.value
}

func (v *NullableAgentMediaResult) Set(val *AgentMediaResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMediaResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMediaResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMediaResult(val *AgentMediaResult) *NullableAgentMediaResult {
	return &NullableAgentMediaResult{value: val, isSet: true}
}

func (v NullableAgentMediaResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMediaResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// AgentActionResponseRequestResponse - 形状由 action.kind 决定，并须满足 action.input_schema： kind=answers 为问答回复；kind=confirmation 为确认决策回复。
type AgentActionResponseRequestResponse struct {
	AgentAnswersActionResponse      *AgentAnswersActionResponse
	AgentConfirmationActionResponse *AgentConfirmationActionResponse
}

// AgentAnswersActionResponseAsAgentActionResponseRequestResponse is a convenience function that returns AgentAnswersActionResponse wrapped in AgentActionResponseRequestResponse
func AgentAnswersActionResponseAsAgentActionResponseRequestResponse(v *AgentAnswersActionResponse) AgentActionResponseRequestResponse {
	return AgentActionResponseRequestResponse{
		AgentAnswersActionResponse: v,
	}
}

// AgentConfirmationActionResponseAsAgentActionResponseRequestResponse is a convenience function that returns AgentConfirmationActionResponse wrapped in AgentActionResponseRequestResponse
func AgentConfirmationActionResponseAsAgentActionResponseRequestResponse(v *AgentConfirmationActionResponse) AgentActionResponseRequestResponse {
	return AgentActionResponseRequestResponse{
		AgentConfirmationActionResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AgentActionResponseRequestResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AgentAnswersActionResponse
	err = newStrictDecoder(data).Decode(&dst.AgentAnswersActionResponse)
	if err == nil {
		jsonAgentAnswersActionResponse, _ := json.Marshal(dst.AgentAnswersActionResponse)
		if string(jsonAgentAnswersActionResponse) == "{}" { // empty struct
			dst.AgentAnswersActionResponse = nil
		} else {
			match++
		}
	} else {
		dst.AgentAnswersActionResponse = nil
	}

	// try to unmarshal data into AgentConfirmationActionResponse
	err = newStrictDecoder(data).Decode(&dst.AgentConfirmationActionResponse)
	if err == nil {
		jsonAgentConfirmationActionResponse, _ := json.Marshal(dst.AgentConfirmationActionResponse)
		if string(jsonAgentConfirmationActionResponse) == "{}" { // empty struct
			dst.AgentConfirmationActionResponse = nil
		} else {
			match++
		}
	} else {
		dst.AgentConfirmationActionResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AgentAnswersActionResponse = nil
		dst.AgentConfirmationActionResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AgentActionResponseRequestResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AgentActionResponseRequestResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentActionResponseRequestResponse) MarshalJSON() ([]byte, error) {
	if src.AgentAnswersActionResponse != nil {
		return json.Marshal(&src.AgentAnswersActionResponse)
	}

	if src.AgentConfirmationActionResponse != nil {
		return json.Marshal(&src.AgentConfirmationActionResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AgentActionResponseRequestResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AgentAnswersActionResponse != nil {
		return obj.AgentAnswersActionResponse
	}

	if obj.AgentConfirmationActionResponse != nil {
		return obj.AgentConfirmationActionResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AgentActionResponseRequestResponse) GetActualInstanceValue() interface{} {
	if obj.AgentAnswersActionResponse != nil {
		return *obj.AgentAnswersActionResponse
	}

	if obj.AgentConfirmationActionResponse != nil {
		return *obj.AgentConfirmationActionResponse
	}

	// all schemas are nil
	return nil
}

type NullableAgentActionResponseRequestResponse struct {
	value *AgentActionResponseRequestResponse
	isSet bool
}

func (v NullableAgentActionResponseRequestResponse) Get() *AgentActionResponseRequestResponse {
	return v.value
}

func (v *NullableAgentActionResponseRequestResponse) Set(val *AgentActionResponseRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentActionResponseRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentActionResponseRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentActionResponseRequestResponse(val *AgentActionResponseRequestResponse) *NullableAgentActionResponseRequestResponse {
	return &NullableAgentActionResponseRequestResponse{value: val, isSet: true}
}

func (v NullableAgentActionResponseRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentActionResponseRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

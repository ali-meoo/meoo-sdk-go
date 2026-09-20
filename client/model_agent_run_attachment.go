package client

import (
	"encoding/json"
	"fmt"
)

// AgentRunAttachment - struct for AgentRunAttachment
type AgentRunAttachment struct {
	AgentRunFileAttachment  *AgentRunFileAttachment
	AgentRunMediaAttachment *AgentRunMediaAttachment
}

// AgentRunFileAttachmentAsAgentRunAttachment is a convenience function that returns AgentRunFileAttachment wrapped in AgentRunAttachment
func AgentRunFileAttachmentAsAgentRunAttachment(v *AgentRunFileAttachment) AgentRunAttachment {
	return AgentRunAttachment{
		AgentRunFileAttachment: v,
	}
}

// AgentRunMediaAttachmentAsAgentRunAttachment is a convenience function that returns AgentRunMediaAttachment wrapped in AgentRunAttachment
func AgentRunMediaAttachmentAsAgentRunAttachment(v *AgentRunMediaAttachment) AgentRunAttachment {
	return AgentRunAttachment{
		AgentRunMediaAttachment: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AgentRunAttachment) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AgentRunFileAttachment
	err = newStrictDecoder(data).Decode(&dst.AgentRunFileAttachment)
	if err == nil {
		jsonAgentRunFileAttachment, _ := json.Marshal(dst.AgentRunFileAttachment)
		if string(jsonAgentRunFileAttachment) == "{}" { // empty struct
			dst.AgentRunFileAttachment = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunFileAttachment = nil
	}

	// try to unmarshal data into AgentRunMediaAttachment
	err = newStrictDecoder(data).Decode(&dst.AgentRunMediaAttachment)
	if err == nil {
		jsonAgentRunMediaAttachment, _ := json.Marshal(dst.AgentRunMediaAttachment)
		if string(jsonAgentRunMediaAttachment) == "{}" { // empty struct
			dst.AgentRunMediaAttachment = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunMediaAttachment = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AgentRunFileAttachment = nil
		dst.AgentRunMediaAttachment = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AgentRunAttachment)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AgentRunAttachment)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentRunAttachment) MarshalJSON() ([]byte, error) {
	if src.AgentRunFileAttachment != nil {
		return json.Marshal(&src.AgentRunFileAttachment)
	}

	if src.AgentRunMediaAttachment != nil {
		return json.Marshal(&src.AgentRunMediaAttachment)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AgentRunAttachment) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AgentRunFileAttachment != nil {
		return obj.AgentRunFileAttachment
	}

	if obj.AgentRunMediaAttachment != nil {
		return obj.AgentRunMediaAttachment
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AgentRunAttachment) GetActualInstanceValue() interface{} {
	if obj.AgentRunFileAttachment != nil {
		return *obj.AgentRunFileAttachment
	}

	if obj.AgentRunMediaAttachment != nil {
		return *obj.AgentRunMediaAttachment
	}

	// all schemas are nil
	return nil
}

type NullableAgentRunAttachment struct {
	value *AgentRunAttachment
	isSet bool
}

func (v NullableAgentRunAttachment) Get() *AgentRunAttachment {
	return v.value
}

func (v *NullableAgentRunAttachment) Set(val *AgentRunAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunAttachment(val *AgentRunAttachment) *NullableAgentRunAttachment {
	return &NullableAgentRunAttachment{value: val, isSet: true}
}

func (v NullableAgentRunAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

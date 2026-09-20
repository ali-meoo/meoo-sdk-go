package client

import (
	"encoding/json"
	"fmt"
)

// AgentRunEvent - struct for AgentRunEvent
type AgentRunEvent struct {
	AgentMessageDeltaEvent      *AgentMessageDeltaEvent
	AgentMessageSnapshotEvent   *AgentMessageSnapshotEvent
	AgentPreviewReadyEvent      *AgentPreviewReadyEvent
	AgentRunInputRequiredEvent  *AgentRunInputRequiredEvent
	AgentRunSupersededEvent     *AgentRunSupersededEvent
	AgentRunTerminalEvent       *AgentRunTerminalEvent
	AgentRunUsageEvent          *AgentRunUsageEvent
	AgentRunWorkingEvent        *AgentRunWorkingEvent
	AgentToolCallEvent          *AgentToolCallEvent
	AgentToolInputDeltaEvent    *AgentToolInputDeltaEvent
	AgentToolInputSnapshotEvent *AgentToolInputSnapshotEvent
}

// AgentMessageDeltaEventAsAgentRunEvent is a convenience function that returns AgentMessageDeltaEvent wrapped in AgentRunEvent
func AgentMessageDeltaEventAsAgentRunEvent(v *AgentMessageDeltaEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentMessageDeltaEvent: v,
	}
}

// AgentMessageSnapshotEventAsAgentRunEvent is a convenience function that returns AgentMessageSnapshotEvent wrapped in AgentRunEvent
func AgentMessageSnapshotEventAsAgentRunEvent(v *AgentMessageSnapshotEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentMessageSnapshotEvent: v,
	}
}

// AgentPreviewReadyEventAsAgentRunEvent is a convenience function that returns AgentPreviewReadyEvent wrapped in AgentRunEvent
func AgentPreviewReadyEventAsAgentRunEvent(v *AgentPreviewReadyEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentPreviewReadyEvent: v,
	}
}

// AgentRunInputRequiredEventAsAgentRunEvent is a convenience function that returns AgentRunInputRequiredEvent wrapped in AgentRunEvent
func AgentRunInputRequiredEventAsAgentRunEvent(v *AgentRunInputRequiredEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentRunInputRequiredEvent: v,
	}
}

// AgentRunSupersededEventAsAgentRunEvent is a convenience function that returns AgentRunSupersededEvent wrapped in AgentRunEvent
func AgentRunSupersededEventAsAgentRunEvent(v *AgentRunSupersededEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentRunSupersededEvent: v,
	}
}

// AgentRunTerminalEventAsAgentRunEvent is a convenience function that returns AgentRunTerminalEvent wrapped in AgentRunEvent
func AgentRunTerminalEventAsAgentRunEvent(v *AgentRunTerminalEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentRunTerminalEvent: v,
	}
}

// AgentRunUsageEventAsAgentRunEvent is a convenience function that returns AgentRunUsageEvent wrapped in AgentRunEvent
func AgentRunUsageEventAsAgentRunEvent(v *AgentRunUsageEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentRunUsageEvent: v,
	}
}

// AgentRunWorkingEventAsAgentRunEvent is a convenience function that returns AgentRunWorkingEvent wrapped in AgentRunEvent
func AgentRunWorkingEventAsAgentRunEvent(v *AgentRunWorkingEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentRunWorkingEvent: v,
	}
}

// AgentToolCallEventAsAgentRunEvent is a convenience function that returns AgentToolCallEvent wrapped in AgentRunEvent
func AgentToolCallEventAsAgentRunEvent(v *AgentToolCallEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentToolCallEvent: v,
	}
}

// AgentToolInputDeltaEventAsAgentRunEvent is a convenience function that returns AgentToolInputDeltaEvent wrapped in AgentRunEvent
func AgentToolInputDeltaEventAsAgentRunEvent(v *AgentToolInputDeltaEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentToolInputDeltaEvent: v,
	}
}

// AgentToolInputSnapshotEventAsAgentRunEvent is a convenience function that returns AgentToolInputSnapshotEvent wrapped in AgentRunEvent
func AgentToolInputSnapshotEventAsAgentRunEvent(v *AgentToolInputSnapshotEvent) AgentRunEvent {
	return AgentRunEvent{
		AgentToolInputSnapshotEvent: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AgentRunEvent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AgentMessageDeltaEvent
	err = newStrictDecoder(data).Decode(&dst.AgentMessageDeltaEvent)
	if err == nil {
		jsonAgentMessageDeltaEvent, _ := json.Marshal(dst.AgentMessageDeltaEvent)
		if string(jsonAgentMessageDeltaEvent) == "{}" { // empty struct
			dst.AgentMessageDeltaEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentMessageDeltaEvent = nil
	}

	// try to unmarshal data into AgentMessageSnapshotEvent
	err = newStrictDecoder(data).Decode(&dst.AgentMessageSnapshotEvent)
	if err == nil {
		jsonAgentMessageSnapshotEvent, _ := json.Marshal(dst.AgentMessageSnapshotEvent)
		if string(jsonAgentMessageSnapshotEvent) == "{}" { // empty struct
			dst.AgentMessageSnapshotEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentMessageSnapshotEvent = nil
	}

	// try to unmarshal data into AgentPreviewReadyEvent
	err = newStrictDecoder(data).Decode(&dst.AgentPreviewReadyEvent)
	if err == nil {
		jsonAgentPreviewReadyEvent, _ := json.Marshal(dst.AgentPreviewReadyEvent)
		if string(jsonAgentPreviewReadyEvent) == "{}" { // empty struct
			dst.AgentPreviewReadyEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentPreviewReadyEvent = nil
	}

	// try to unmarshal data into AgentRunInputRequiredEvent
	err = newStrictDecoder(data).Decode(&dst.AgentRunInputRequiredEvent)
	if err == nil {
		jsonAgentRunInputRequiredEvent, _ := json.Marshal(dst.AgentRunInputRequiredEvent)
		if string(jsonAgentRunInputRequiredEvent) == "{}" { // empty struct
			dst.AgentRunInputRequiredEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunInputRequiredEvent = nil
	}

	// try to unmarshal data into AgentRunSupersededEvent
	err = newStrictDecoder(data).Decode(&dst.AgentRunSupersededEvent)
	if err == nil {
		jsonAgentRunSupersededEvent, _ := json.Marshal(dst.AgentRunSupersededEvent)
		if string(jsonAgentRunSupersededEvent) == "{}" { // empty struct
			dst.AgentRunSupersededEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunSupersededEvent = nil
	}

	// try to unmarshal data into AgentRunTerminalEvent
	err = newStrictDecoder(data).Decode(&dst.AgentRunTerminalEvent)
	if err == nil {
		jsonAgentRunTerminalEvent, _ := json.Marshal(dst.AgentRunTerminalEvent)
		if string(jsonAgentRunTerminalEvent) == "{}" { // empty struct
			dst.AgentRunTerminalEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunTerminalEvent = nil
	}

	// try to unmarshal data into AgentRunUsageEvent
	err = newStrictDecoder(data).Decode(&dst.AgentRunUsageEvent)
	if err == nil {
		jsonAgentRunUsageEvent, _ := json.Marshal(dst.AgentRunUsageEvent)
		if string(jsonAgentRunUsageEvent) == "{}" { // empty struct
			dst.AgentRunUsageEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunUsageEvent = nil
	}

	// try to unmarshal data into AgentRunWorkingEvent
	err = newStrictDecoder(data).Decode(&dst.AgentRunWorkingEvent)
	if err == nil {
		jsonAgentRunWorkingEvent, _ := json.Marshal(dst.AgentRunWorkingEvent)
		if string(jsonAgentRunWorkingEvent) == "{}" { // empty struct
			dst.AgentRunWorkingEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentRunWorkingEvent = nil
	}

	// try to unmarshal data into AgentToolCallEvent
	err = newStrictDecoder(data).Decode(&dst.AgentToolCallEvent)
	if err == nil {
		jsonAgentToolCallEvent, _ := json.Marshal(dst.AgentToolCallEvent)
		if string(jsonAgentToolCallEvent) == "{}" { // empty struct
			dst.AgentToolCallEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentToolCallEvent = nil
	}

	// try to unmarshal data into AgentToolInputDeltaEvent
	err = newStrictDecoder(data).Decode(&dst.AgentToolInputDeltaEvent)
	if err == nil {
		jsonAgentToolInputDeltaEvent, _ := json.Marshal(dst.AgentToolInputDeltaEvent)
		if string(jsonAgentToolInputDeltaEvent) == "{}" { // empty struct
			dst.AgentToolInputDeltaEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentToolInputDeltaEvent = nil
	}

	// try to unmarshal data into AgentToolInputSnapshotEvent
	err = newStrictDecoder(data).Decode(&dst.AgentToolInputSnapshotEvent)
	if err == nil {
		jsonAgentToolInputSnapshotEvent, _ := json.Marshal(dst.AgentToolInputSnapshotEvent)
		if string(jsonAgentToolInputSnapshotEvent) == "{}" { // empty struct
			dst.AgentToolInputSnapshotEvent = nil
		} else {
			match++
		}
	} else {
		dst.AgentToolInputSnapshotEvent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AgentMessageDeltaEvent = nil
		dst.AgentMessageSnapshotEvent = nil
		dst.AgentPreviewReadyEvent = nil
		dst.AgentRunInputRequiredEvent = nil
		dst.AgentRunSupersededEvent = nil
		dst.AgentRunTerminalEvent = nil
		dst.AgentRunUsageEvent = nil
		dst.AgentRunWorkingEvent = nil
		dst.AgentToolCallEvent = nil
		dst.AgentToolInputDeltaEvent = nil
		dst.AgentToolInputSnapshotEvent = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AgentRunEvent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AgentRunEvent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentRunEvent) MarshalJSON() ([]byte, error) {
	if src.AgentMessageDeltaEvent != nil {
		return json.Marshal(&src.AgentMessageDeltaEvent)
	}

	if src.AgentMessageSnapshotEvent != nil {
		return json.Marshal(&src.AgentMessageSnapshotEvent)
	}

	if src.AgentPreviewReadyEvent != nil {
		return json.Marshal(&src.AgentPreviewReadyEvent)
	}

	if src.AgentRunInputRequiredEvent != nil {
		return json.Marshal(&src.AgentRunInputRequiredEvent)
	}

	if src.AgentRunSupersededEvent != nil {
		return json.Marshal(&src.AgentRunSupersededEvent)
	}

	if src.AgentRunTerminalEvent != nil {
		return json.Marshal(&src.AgentRunTerminalEvent)
	}

	if src.AgentRunUsageEvent != nil {
		return json.Marshal(&src.AgentRunUsageEvent)
	}

	if src.AgentRunWorkingEvent != nil {
		return json.Marshal(&src.AgentRunWorkingEvent)
	}

	if src.AgentToolCallEvent != nil {
		return json.Marshal(&src.AgentToolCallEvent)
	}

	if src.AgentToolInputDeltaEvent != nil {
		return json.Marshal(&src.AgentToolInputDeltaEvent)
	}

	if src.AgentToolInputSnapshotEvent != nil {
		return json.Marshal(&src.AgentToolInputSnapshotEvent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AgentRunEvent) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AgentMessageDeltaEvent != nil {
		return obj.AgentMessageDeltaEvent
	}

	if obj.AgentMessageSnapshotEvent != nil {
		return obj.AgentMessageSnapshotEvent
	}

	if obj.AgentPreviewReadyEvent != nil {
		return obj.AgentPreviewReadyEvent
	}

	if obj.AgentRunInputRequiredEvent != nil {
		return obj.AgentRunInputRequiredEvent
	}

	if obj.AgentRunSupersededEvent != nil {
		return obj.AgentRunSupersededEvent
	}

	if obj.AgentRunTerminalEvent != nil {
		return obj.AgentRunTerminalEvent
	}

	if obj.AgentRunUsageEvent != nil {
		return obj.AgentRunUsageEvent
	}

	if obj.AgentRunWorkingEvent != nil {
		return obj.AgentRunWorkingEvent
	}

	if obj.AgentToolCallEvent != nil {
		return obj.AgentToolCallEvent
	}

	if obj.AgentToolInputDeltaEvent != nil {
		return obj.AgentToolInputDeltaEvent
	}

	if obj.AgentToolInputSnapshotEvent != nil {
		return obj.AgentToolInputSnapshotEvent
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AgentRunEvent) GetActualInstanceValue() interface{} {
	if obj.AgentMessageDeltaEvent != nil {
		return *obj.AgentMessageDeltaEvent
	}

	if obj.AgentMessageSnapshotEvent != nil {
		return *obj.AgentMessageSnapshotEvent
	}

	if obj.AgentPreviewReadyEvent != nil {
		return *obj.AgentPreviewReadyEvent
	}

	if obj.AgentRunInputRequiredEvent != nil {
		return *obj.AgentRunInputRequiredEvent
	}

	if obj.AgentRunSupersededEvent != nil {
		return *obj.AgentRunSupersededEvent
	}

	if obj.AgentRunTerminalEvent != nil {
		return *obj.AgentRunTerminalEvent
	}

	if obj.AgentRunUsageEvent != nil {
		return *obj.AgentRunUsageEvent
	}

	if obj.AgentRunWorkingEvent != nil {
		return *obj.AgentRunWorkingEvent
	}

	if obj.AgentToolCallEvent != nil {
		return *obj.AgentToolCallEvent
	}

	if obj.AgentToolInputDeltaEvent != nil {
		return *obj.AgentToolInputDeltaEvent
	}

	if obj.AgentToolInputSnapshotEvent != nil {
		return *obj.AgentToolInputSnapshotEvent
	}

	// all schemas are nil
	return nil
}

type NullableAgentRunEvent struct {
	value *AgentRunEvent
	isSet bool
}

func (v NullableAgentRunEvent) Get() *AgentRunEvent {
	return v.value
}

func (v *NullableAgentRunEvent) Set(val *AgentRunEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunEvent(val *AgentRunEvent) *NullableAgentRunEvent {
	return &NullableAgentRunEvent{value: val, isSet: true}
}

func (v NullableAgentRunEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

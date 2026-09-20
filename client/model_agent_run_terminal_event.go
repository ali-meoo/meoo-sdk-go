package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentRunTerminalEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunTerminalEvent{}

// AgentRunTerminalEvent struct for AgentRunTerminalEvent
type AgentRunTerminalEvent struct {
	Event                string                    `json:"event"`
	Data                 AgentRunTerminalEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentRunTerminalEvent AgentRunTerminalEvent

// NewAgentRunTerminalEvent instantiates a new AgentRunTerminalEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunTerminalEvent(event string, data AgentRunTerminalEventData) *AgentRunTerminalEvent {
	this := AgentRunTerminalEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentRunTerminalEventWithDefaults instantiates a new AgentRunTerminalEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunTerminalEventWithDefaults() *AgentRunTerminalEvent {
	this := AgentRunTerminalEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentRunTerminalEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentRunTerminalEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentRunTerminalEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentRunTerminalEvent) GetData() AgentRunTerminalEventData {
	if o == nil {
		var ret AgentRunTerminalEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentRunTerminalEvent) GetDataOk() (*AgentRunTerminalEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentRunTerminalEvent) SetData(v AgentRunTerminalEventData) {
	o.Data = v
}

func (o AgentRunTerminalEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunTerminalEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentRunTerminalEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"data",
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

	varAgentRunTerminalEvent := _AgentRunTerminalEvent{}

	err = json.Unmarshal(data, &varAgentRunTerminalEvent)

	if err != nil {
		return err
	}

	*o = AgentRunTerminalEvent(varAgentRunTerminalEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentRunTerminalEvent struct {
	value *AgentRunTerminalEvent
	isSet bool
}

func (v NullableAgentRunTerminalEvent) Get() *AgentRunTerminalEvent {
	return v.value
}

func (v *NullableAgentRunTerminalEvent) Set(val *AgentRunTerminalEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunTerminalEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunTerminalEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunTerminalEvent(val *AgentRunTerminalEvent) *NullableAgentRunTerminalEvent {
	return &NullableAgentRunTerminalEvent{value: val, isSet: true}
}

func (v NullableAgentRunTerminalEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunTerminalEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

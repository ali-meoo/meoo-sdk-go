package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunWorkingEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunWorkingEvent{}

// AgentRunWorkingEvent struct for AgentRunWorkingEvent
type AgentRunWorkingEvent struct {
	Event string                   `json:"event"`
	Data  AgentRunWorkingEventData `json:"data"`
}

type _AgentRunWorkingEvent AgentRunWorkingEvent

// NewAgentRunWorkingEvent instantiates a new AgentRunWorkingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunWorkingEvent(event string, data AgentRunWorkingEventData) *AgentRunWorkingEvent {
	this := AgentRunWorkingEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentRunWorkingEventWithDefaults instantiates a new AgentRunWorkingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunWorkingEventWithDefaults() *AgentRunWorkingEvent {
	this := AgentRunWorkingEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentRunWorkingEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentRunWorkingEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentRunWorkingEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentRunWorkingEvent) GetData() AgentRunWorkingEventData {
	if o == nil {
		var ret AgentRunWorkingEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentRunWorkingEvent) GetDataOk() (*AgentRunWorkingEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentRunWorkingEvent) SetData(v AgentRunWorkingEventData) {
	o.Data = v
}

func (o AgentRunWorkingEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunWorkingEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AgentRunWorkingEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunWorkingEvent := _AgentRunWorkingEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentRunWorkingEvent)

	if err != nil {
		return err
	}

	*o = AgentRunWorkingEvent(varAgentRunWorkingEvent)

	return err
}

type NullableAgentRunWorkingEvent struct {
	value *AgentRunWorkingEvent
	isSet bool
}

func (v NullableAgentRunWorkingEvent) Get() *AgentRunWorkingEvent {
	return v.value
}

func (v *NullableAgentRunWorkingEvent) Set(val *AgentRunWorkingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunWorkingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunWorkingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunWorkingEvent(val *AgentRunWorkingEvent) *NullableAgentRunWorkingEvent {
	return &NullableAgentRunWorkingEvent{value: val, isSet: true}
}

func (v NullableAgentRunWorkingEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunWorkingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

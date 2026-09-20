package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentToolCallEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolCallEvent{}

// AgentToolCallEvent struct for AgentToolCallEvent
type AgentToolCallEvent struct {
	Event string                 `json:"event"`
	Data  AgentToolCallEventData `json:"data"`
}

type _AgentToolCallEvent AgentToolCallEvent

// NewAgentToolCallEvent instantiates a new AgentToolCallEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolCallEvent(event string, data AgentToolCallEventData) *AgentToolCallEvent {
	this := AgentToolCallEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentToolCallEventWithDefaults instantiates a new AgentToolCallEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolCallEventWithDefaults() *AgentToolCallEvent {
	this := AgentToolCallEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentToolCallEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentToolCallEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentToolCallEvent) GetData() AgentToolCallEventData {
	if o == nil {
		var ret AgentToolCallEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentToolCallEvent) GetDataOk() (*AgentToolCallEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentToolCallEvent) SetData(v AgentToolCallEventData) {
	o.Data = v
}

func (o AgentToolCallEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolCallEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AgentToolCallEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentToolCallEvent := _AgentToolCallEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentToolCallEvent)

	if err != nil {
		return err
	}

	*o = AgentToolCallEvent(varAgentToolCallEvent)

	return err
}

type NullableAgentToolCallEvent struct {
	value *AgentToolCallEvent
	isSet bool
}

func (v NullableAgentToolCallEvent) Get() *AgentToolCallEvent {
	return v.value
}

func (v *NullableAgentToolCallEvent) Set(val *AgentToolCallEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolCallEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolCallEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolCallEvent(val *AgentToolCallEvent) *NullableAgentToolCallEvent {
	return &NullableAgentToolCallEvent{value: val, isSet: true}
}

func (v NullableAgentToolCallEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolCallEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

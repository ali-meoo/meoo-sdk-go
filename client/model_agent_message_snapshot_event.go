package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentMessageSnapshotEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMessageSnapshotEvent{}

// AgentMessageSnapshotEvent struct for AgentMessageSnapshotEvent
type AgentMessageSnapshotEvent struct {
	Event string                        `json:"event"`
	Data  AgentMessageSnapshotEventData `json:"data"`
}

type _AgentMessageSnapshotEvent AgentMessageSnapshotEvent

// NewAgentMessageSnapshotEvent instantiates a new AgentMessageSnapshotEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMessageSnapshotEvent(event string, data AgentMessageSnapshotEventData) *AgentMessageSnapshotEvent {
	this := AgentMessageSnapshotEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentMessageSnapshotEventWithDefaults instantiates a new AgentMessageSnapshotEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMessageSnapshotEventWithDefaults() *AgentMessageSnapshotEvent {
	this := AgentMessageSnapshotEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentMessageSnapshotEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentMessageSnapshotEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentMessageSnapshotEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentMessageSnapshotEvent) GetData() AgentMessageSnapshotEventData {
	if o == nil {
		var ret AgentMessageSnapshotEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentMessageSnapshotEvent) GetDataOk() (*AgentMessageSnapshotEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentMessageSnapshotEvent) SetData(v AgentMessageSnapshotEventData) {
	o.Data = v
}

func (o AgentMessageSnapshotEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMessageSnapshotEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AgentMessageSnapshotEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentMessageSnapshotEvent := _AgentMessageSnapshotEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentMessageSnapshotEvent)

	if err != nil {
		return err
	}

	*o = AgentMessageSnapshotEvent(varAgentMessageSnapshotEvent)

	return err
}

type NullableAgentMessageSnapshotEvent struct {
	value *AgentMessageSnapshotEvent
	isSet bool
}

func (v NullableAgentMessageSnapshotEvent) Get() *AgentMessageSnapshotEvent {
	return v.value
}

func (v *NullableAgentMessageSnapshotEvent) Set(val *AgentMessageSnapshotEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMessageSnapshotEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMessageSnapshotEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMessageSnapshotEvent(val *AgentMessageSnapshotEvent) *NullableAgentMessageSnapshotEvent {
	return &NullableAgentMessageSnapshotEvent{value: val, isSet: true}
}

func (v NullableAgentMessageSnapshotEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMessageSnapshotEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

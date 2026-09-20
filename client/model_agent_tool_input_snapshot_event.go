package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentToolInputSnapshotEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolInputSnapshotEvent{}

// AgentToolInputSnapshotEvent struct for AgentToolInputSnapshotEvent
type AgentToolInputSnapshotEvent struct {
	Event                string                          `json:"event"`
	Data                 AgentToolInputSnapshotEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentToolInputSnapshotEvent AgentToolInputSnapshotEvent

// NewAgentToolInputSnapshotEvent instantiates a new AgentToolInputSnapshotEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolInputSnapshotEvent(event string, data AgentToolInputSnapshotEventData) *AgentToolInputSnapshotEvent {
	this := AgentToolInputSnapshotEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentToolInputSnapshotEventWithDefaults instantiates a new AgentToolInputSnapshotEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolInputSnapshotEventWithDefaults() *AgentToolInputSnapshotEvent {
	this := AgentToolInputSnapshotEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentToolInputSnapshotEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentToolInputSnapshotEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentToolInputSnapshotEvent) GetData() AgentToolInputSnapshotEventData {
	if o == nil {
		var ret AgentToolInputSnapshotEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputSnapshotEvent) GetDataOk() (*AgentToolInputSnapshotEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentToolInputSnapshotEvent) SetData(v AgentToolInputSnapshotEventData) {
	o.Data = v
}

func (o AgentToolInputSnapshotEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolInputSnapshotEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentToolInputSnapshotEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentToolInputSnapshotEvent := _AgentToolInputSnapshotEvent{}

	err = json.Unmarshal(data, &varAgentToolInputSnapshotEvent)

	if err != nil {
		return err
	}

	*o = AgentToolInputSnapshotEvent(varAgentToolInputSnapshotEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentToolInputSnapshotEvent struct {
	value *AgentToolInputSnapshotEvent
	isSet bool
}

func (v NullableAgentToolInputSnapshotEvent) Get() *AgentToolInputSnapshotEvent {
	return v.value
}

func (v *NullableAgentToolInputSnapshotEvent) Set(val *AgentToolInputSnapshotEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolInputSnapshotEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolInputSnapshotEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolInputSnapshotEvent(val *AgentToolInputSnapshotEvent) *NullableAgentToolInputSnapshotEvent {
	return &NullableAgentToolInputSnapshotEvent{value: val, isSet: true}
}

func (v NullableAgentToolInputSnapshotEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolInputSnapshotEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentMessageDeltaEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentMessageDeltaEvent{}

// AgentMessageDeltaEvent struct for AgentMessageDeltaEvent
type AgentMessageDeltaEvent struct {
	Event                string                     `json:"event"`
	Data                 AgentMessageDeltaEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentMessageDeltaEvent AgentMessageDeltaEvent

// NewAgentMessageDeltaEvent instantiates a new AgentMessageDeltaEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMessageDeltaEvent(event string, data AgentMessageDeltaEventData) *AgentMessageDeltaEvent {
	this := AgentMessageDeltaEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentMessageDeltaEventWithDefaults instantiates a new AgentMessageDeltaEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMessageDeltaEventWithDefaults() *AgentMessageDeltaEvent {
	this := AgentMessageDeltaEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentMessageDeltaEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentMessageDeltaEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentMessageDeltaEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentMessageDeltaEvent) GetData() AgentMessageDeltaEventData {
	if o == nil {
		var ret AgentMessageDeltaEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentMessageDeltaEvent) GetDataOk() (*AgentMessageDeltaEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentMessageDeltaEvent) SetData(v AgentMessageDeltaEventData) {
	o.Data = v
}

func (o AgentMessageDeltaEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMessageDeltaEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentMessageDeltaEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentMessageDeltaEvent := _AgentMessageDeltaEvent{}

	err = json.Unmarshal(data, &varAgentMessageDeltaEvent)

	if err != nil {
		return err
	}

	*o = AgentMessageDeltaEvent(varAgentMessageDeltaEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentMessageDeltaEvent struct {
	value *AgentMessageDeltaEvent
	isSet bool
}

func (v NullableAgentMessageDeltaEvent) Get() *AgentMessageDeltaEvent {
	return v.value
}

func (v *NullableAgentMessageDeltaEvent) Set(val *AgentMessageDeltaEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMessageDeltaEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMessageDeltaEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMessageDeltaEvent(val *AgentMessageDeltaEvent) *NullableAgentMessageDeltaEvent {
	return &NullableAgentMessageDeltaEvent{value: val, isSet: true}
}

func (v NullableAgentMessageDeltaEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMessageDeltaEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentToolInputDeltaEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentToolInputDeltaEvent{}

// AgentToolInputDeltaEvent struct for AgentToolInputDeltaEvent
type AgentToolInputDeltaEvent struct {
	Event                string                       `json:"event"`
	Data                 AgentToolInputDeltaEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentToolInputDeltaEvent AgentToolInputDeltaEvent

// NewAgentToolInputDeltaEvent instantiates a new AgentToolInputDeltaEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToolInputDeltaEvent(event string, data AgentToolInputDeltaEventData) *AgentToolInputDeltaEvent {
	this := AgentToolInputDeltaEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentToolInputDeltaEventWithDefaults instantiates a new AgentToolInputDeltaEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToolInputDeltaEventWithDefaults() *AgentToolInputDeltaEvent {
	this := AgentToolInputDeltaEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentToolInputDeltaEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentToolInputDeltaEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentToolInputDeltaEvent) GetData() AgentToolInputDeltaEventData {
	if o == nil {
		var ret AgentToolInputDeltaEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentToolInputDeltaEvent) GetDataOk() (*AgentToolInputDeltaEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentToolInputDeltaEvent) SetData(v AgentToolInputDeltaEventData) {
	o.Data = v
}

func (o AgentToolInputDeltaEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToolInputDeltaEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentToolInputDeltaEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentToolInputDeltaEvent := _AgentToolInputDeltaEvent{}

	err = json.Unmarshal(data, &varAgentToolInputDeltaEvent)

	if err != nil {
		return err
	}

	*o = AgentToolInputDeltaEvent(varAgentToolInputDeltaEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentToolInputDeltaEvent struct {
	value *AgentToolInputDeltaEvent
	isSet bool
}

func (v NullableAgentToolInputDeltaEvent) Get() *AgentToolInputDeltaEvent {
	return v.value
}

func (v *NullableAgentToolInputDeltaEvent) Set(val *AgentToolInputDeltaEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToolInputDeltaEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToolInputDeltaEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToolInputDeltaEvent(val *AgentToolInputDeltaEvent) *NullableAgentToolInputDeltaEvent {
	return &NullableAgentToolInputDeltaEvent{value: val, isSet: true}
}

func (v NullableAgentToolInputDeltaEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToolInputDeltaEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

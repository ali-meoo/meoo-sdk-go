package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentRunSupersededEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunSupersededEvent{}

// AgentRunSupersededEvent struct for AgentRunSupersededEvent
type AgentRunSupersededEvent struct {
	Event                string                      `json:"event"`
	Data                 AgentRunSupersededEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentRunSupersededEvent AgentRunSupersededEvent

// NewAgentRunSupersededEvent instantiates a new AgentRunSupersededEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunSupersededEvent(event string, data AgentRunSupersededEventData) *AgentRunSupersededEvent {
	this := AgentRunSupersededEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentRunSupersededEventWithDefaults instantiates a new AgentRunSupersededEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunSupersededEventWithDefaults() *AgentRunSupersededEvent {
	this := AgentRunSupersededEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentRunSupersededEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentRunSupersededEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentRunSupersededEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentRunSupersededEvent) GetData() AgentRunSupersededEventData {
	if o == nil {
		var ret AgentRunSupersededEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentRunSupersededEvent) GetDataOk() (*AgentRunSupersededEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentRunSupersededEvent) SetData(v AgentRunSupersededEventData) {
	o.Data = v
}

func (o AgentRunSupersededEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunSupersededEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentRunSupersededEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunSupersededEvent := _AgentRunSupersededEvent{}

	err = json.Unmarshal(data, &varAgentRunSupersededEvent)

	if err != nil {
		return err
	}

	*o = AgentRunSupersededEvent(varAgentRunSupersededEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentRunSupersededEvent struct {
	value *AgentRunSupersededEvent
	isSet bool
}

func (v NullableAgentRunSupersededEvent) Get() *AgentRunSupersededEvent {
	return v.value
}

func (v *NullableAgentRunSupersededEvent) Set(val *AgentRunSupersededEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunSupersededEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunSupersededEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunSupersededEvent(val *AgentRunSupersededEvent) *NullableAgentRunSupersededEvent {
	return &NullableAgentRunSupersededEvent{value: val, isSet: true}
}

func (v NullableAgentRunSupersededEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunSupersededEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

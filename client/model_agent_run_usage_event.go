package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentRunUsageEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunUsageEvent{}

// AgentRunUsageEvent struct for AgentRunUsageEvent
type AgentRunUsageEvent struct {
	Event                string                 `json:"event"`
	Data                 AgentRunUsageEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentRunUsageEvent AgentRunUsageEvent

// NewAgentRunUsageEvent instantiates a new AgentRunUsageEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunUsageEvent(event string, data AgentRunUsageEventData) *AgentRunUsageEvent {
	this := AgentRunUsageEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentRunUsageEventWithDefaults instantiates a new AgentRunUsageEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunUsageEventWithDefaults() *AgentRunUsageEvent {
	this := AgentRunUsageEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentRunUsageEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentRunUsageEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentRunUsageEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentRunUsageEvent) GetData() AgentRunUsageEventData {
	if o == nil {
		var ret AgentRunUsageEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentRunUsageEvent) GetDataOk() (*AgentRunUsageEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentRunUsageEvent) SetData(v AgentRunUsageEventData) {
	o.Data = v
}

func (o AgentRunUsageEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunUsageEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentRunUsageEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunUsageEvent := _AgentRunUsageEvent{}

	err = json.Unmarshal(data, &varAgentRunUsageEvent)

	if err != nil {
		return err
	}

	*o = AgentRunUsageEvent(varAgentRunUsageEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentRunUsageEvent struct {
	value *AgentRunUsageEvent
	isSet bool
}

func (v NullableAgentRunUsageEvent) Get() *AgentRunUsageEvent {
	return v.value
}

func (v *NullableAgentRunUsageEvent) Set(val *AgentRunUsageEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunUsageEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunUsageEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunUsageEvent(val *AgentRunUsageEvent) *NullableAgentRunUsageEvent {
	return &NullableAgentRunUsageEvent{value: val, isSet: true}
}

func (v NullableAgentRunUsageEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunUsageEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

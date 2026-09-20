package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunInputRequiredEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunInputRequiredEvent{}

// AgentRunInputRequiredEvent struct for AgentRunInputRequiredEvent
type AgentRunInputRequiredEvent struct {
	Event string                         `json:"event"`
	Data  AgentRunInputRequiredEventData `json:"data"`
}

type _AgentRunInputRequiredEvent AgentRunInputRequiredEvent

// NewAgentRunInputRequiredEvent instantiates a new AgentRunInputRequiredEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunInputRequiredEvent(event string, data AgentRunInputRequiredEventData) *AgentRunInputRequiredEvent {
	this := AgentRunInputRequiredEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentRunInputRequiredEventWithDefaults instantiates a new AgentRunInputRequiredEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunInputRequiredEventWithDefaults() *AgentRunInputRequiredEvent {
	this := AgentRunInputRequiredEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentRunInputRequiredEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentRunInputRequiredEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentRunInputRequiredEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentRunInputRequiredEvent) GetData() AgentRunInputRequiredEventData {
	if o == nil {
		var ret AgentRunInputRequiredEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentRunInputRequiredEvent) GetDataOk() (*AgentRunInputRequiredEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentRunInputRequiredEvent) SetData(v AgentRunInputRequiredEventData) {
	o.Data = v
}

func (o AgentRunInputRequiredEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunInputRequiredEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AgentRunInputRequiredEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentRunInputRequiredEvent := _AgentRunInputRequiredEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentRunInputRequiredEvent)

	if err != nil {
		return err
	}

	*o = AgentRunInputRequiredEvent(varAgentRunInputRequiredEvent)

	return err
}

type NullableAgentRunInputRequiredEvent struct {
	value *AgentRunInputRequiredEvent
	isSet bool
}

func (v NullableAgentRunInputRequiredEvent) Get() *AgentRunInputRequiredEvent {
	return v.value
}

func (v *NullableAgentRunInputRequiredEvent) Set(val *AgentRunInputRequiredEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunInputRequiredEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunInputRequiredEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunInputRequiredEvent(val *AgentRunInputRequiredEvent) *NullableAgentRunInputRequiredEvent {
	return &NullableAgentRunInputRequiredEvent{value: val, isSet: true}
}

func (v NullableAgentRunInputRequiredEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunInputRequiredEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentPreviewReadyEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPreviewReadyEvent{}

// AgentPreviewReadyEvent struct for AgentPreviewReadyEvent
type AgentPreviewReadyEvent struct {
	Event                string                     `json:"event"`
	Data                 AgentPreviewReadyEventData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AgentPreviewReadyEvent AgentPreviewReadyEvent

// NewAgentPreviewReadyEvent instantiates a new AgentPreviewReadyEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPreviewReadyEvent(event string, data AgentPreviewReadyEventData) *AgentPreviewReadyEvent {
	this := AgentPreviewReadyEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewAgentPreviewReadyEventWithDefaults instantiates a new AgentPreviewReadyEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPreviewReadyEventWithDefaults() *AgentPreviewReadyEvent {
	this := AgentPreviewReadyEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *AgentPreviewReadyEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AgentPreviewReadyEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AgentPreviewReadyEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *AgentPreviewReadyEvent) GetData() AgentPreviewReadyEventData {
	if o == nil {
		var ret AgentPreviewReadyEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgentPreviewReadyEvent) GetDataOk() (*AgentPreviewReadyEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgentPreviewReadyEvent) SetData(v AgentPreviewReadyEventData) {
	o.Data = v
}

func (o AgentPreviewReadyEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPreviewReadyEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentPreviewReadyEvent) UnmarshalJSON(data []byte) (err error) {
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

	varAgentPreviewReadyEvent := _AgentPreviewReadyEvent{}

	err = json.Unmarshal(data, &varAgentPreviewReadyEvent)

	if err != nil {
		return err
	}

	*o = AgentPreviewReadyEvent(varAgentPreviewReadyEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentPreviewReadyEvent struct {
	value *AgentPreviewReadyEvent
	isSet bool
}

func (v NullableAgentPreviewReadyEvent) Get() *AgentPreviewReadyEvent {
	return v.value
}

func (v *NullableAgentPreviewReadyEvent) Set(val *AgentPreviewReadyEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPreviewReadyEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPreviewReadyEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPreviewReadyEvent(val *AgentPreviewReadyEvent) *NullableAgentPreviewReadyEvent {
	return &NullableAgentPreviewReadyEvent{value: val, isSet: true}
}

func (v NullableAgentPreviewReadyEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPreviewReadyEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

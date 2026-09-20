package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentCapabilityModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentCapabilityModel{}

// AgentCapabilityModel struct for AgentCapabilityModel
type AgentCapabilityModel struct {
	// 启动 Run 时可原样传入 model 的 canonical ID。
	Id string `json:"id"`
	// 面向用户的模型展示名称。
	Name string `json:"name"`
}

type _AgentCapabilityModel AgentCapabilityModel

// NewAgentCapabilityModel instantiates a new AgentCapabilityModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentCapabilityModel(id string, name string) *AgentCapabilityModel {
	this := AgentCapabilityModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAgentCapabilityModelWithDefaults instantiates a new AgentCapabilityModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentCapabilityModelWithDefaults() *AgentCapabilityModel {
	this := AgentCapabilityModel{}
	return &this
}

// GetId returns the Id field value
func (o *AgentCapabilityModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilityModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentCapabilityModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AgentCapabilityModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilityModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AgentCapabilityModel) SetName(v string) {
	o.Name = v
}

func (o AgentCapabilityModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentCapabilityModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *AgentCapabilityModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varAgentCapabilityModel := _AgentCapabilityModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentCapabilityModel)

	if err != nil {
		return err
	}

	*o = AgentCapabilityModel(varAgentCapabilityModel)

	return err
}

type NullableAgentCapabilityModel struct {
	value *AgentCapabilityModel
	isSet bool
}

func (v NullableAgentCapabilityModel) Get() *AgentCapabilityModel {
	return v.value
}

func (v *NullableAgentCapabilityModel) Set(val *AgentCapabilityModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentCapabilityModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentCapabilityModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentCapabilityModel(val *AgentCapabilityModel) *NullableAgentCapabilityModel {
	return &NullableAgentCapabilityModel{value: val, isSet: true}
}

func (v NullableAgentCapabilityModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentCapabilityModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

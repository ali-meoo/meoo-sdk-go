package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AgentCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentCapabilities{}

// AgentCapabilities struct for AgentCapabilities
type AgentCapabilities struct {
	Object    string `json:"object"`
	ProjectId string `json:"project_id"`
	// 当前可选模型目录；调用方不要缓存为永久静态列表。
	Models               []AgentCapabilityModel     `json:"models"`
	SpeedTiers           []AgentCapabilitySpeedTier `json:"speed_tiers"`
	Defaults             AgentCapabilitiesDefaults  `json:"defaults"`
	AdditionalProperties map[string]interface{}
}

type _AgentCapabilities AgentCapabilities

// NewAgentCapabilities instantiates a new AgentCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentCapabilities(object string, projectId string, models []AgentCapabilityModel, speedTiers []AgentCapabilitySpeedTier, defaults AgentCapabilitiesDefaults) *AgentCapabilities {
	this := AgentCapabilities{}
	this.Object = object
	this.ProjectId = projectId
	this.Models = models
	this.SpeedTiers = speedTiers
	this.Defaults = defaults
	return &this
}

// NewAgentCapabilitiesWithDefaults instantiates a new AgentCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentCapabilitiesWithDefaults() *AgentCapabilities {
	this := AgentCapabilities{}
	return &this
}

// GetObject returns the Object field value
func (o *AgentCapabilities) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilities) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *AgentCapabilities) SetObject(v string) {
	o.Object = v
}

// GetProjectId returns the ProjectId field value
func (o *AgentCapabilities) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilities) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AgentCapabilities) SetProjectId(v string) {
	o.ProjectId = v
}

// GetModels returns the Models field value
func (o *AgentCapabilities) GetModels() []AgentCapabilityModel {
	if o == nil {
		var ret []AgentCapabilityModel
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilities) GetModelsOk() ([]AgentCapabilityModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *AgentCapabilities) SetModels(v []AgentCapabilityModel) {
	o.Models = v
}

// GetSpeedTiers returns the SpeedTiers field value
func (o *AgentCapabilities) GetSpeedTiers() []AgentCapabilitySpeedTier {
	if o == nil {
		var ret []AgentCapabilitySpeedTier
		return ret
	}

	return o.SpeedTiers
}

// GetSpeedTiersOk returns a tuple with the SpeedTiers field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilities) GetSpeedTiersOk() ([]AgentCapabilitySpeedTier, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpeedTiers, true
}

// SetSpeedTiers sets field value
func (o *AgentCapabilities) SetSpeedTiers(v []AgentCapabilitySpeedTier) {
	o.SpeedTiers = v
}

// GetDefaults returns the Defaults field value
func (o *AgentCapabilities) GetDefaults() AgentCapabilitiesDefaults {
	if o == nil {
		var ret AgentCapabilitiesDefaults
		return ret
	}

	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value
// and a boolean to check if the value has been set.
func (o *AgentCapabilities) GetDefaultsOk() (*AgentCapabilitiesDefaults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Defaults, true
}

// SetDefaults sets field value
func (o *AgentCapabilities) SetDefaults(v AgentCapabilitiesDefaults) {
	o.Defaults = v
}

func (o AgentCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["project_id"] = o.ProjectId
	toSerialize["models"] = o.Models
	toSerialize["speed_tiers"] = o.SpeedTiers
	toSerialize["defaults"] = o.Defaults

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentCapabilities) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"project_id",
		"models",
		"speed_tiers",
		"defaults",
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

	varAgentCapabilities := _AgentCapabilities{}

	err = json.Unmarshal(data, &varAgentCapabilities)

	if err != nil {
		return err
	}

	*o = AgentCapabilities(varAgentCapabilities)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "models")
		delete(additionalProperties, "speed_tiers")
		delete(additionalProperties, "defaults")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentCapabilities struct {
	value *AgentCapabilities
	isSet bool
}

func (v NullableAgentCapabilities) Get() *AgentCapabilities {
	return v.value
}

func (v *NullableAgentCapabilities) Set(val *AgentCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentCapabilities(val *AgentCapabilities) *NullableAgentCapabilities {
	return &NullableAgentCapabilities{value: val, isSet: true}
}

func (v NullableAgentCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentPreviewReadyEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPreviewReadyEventData{}

// AgentPreviewReadyEventData struct for AgentPreviewReadyEventData
type AgentPreviewReadyEventData struct {
	RunId string `json:"run_id"`
	// 短时 opaque Preview Link；不得解析、记录或长期存储。
	Url string `json:"url"`
	// Preview Link 过期时间，Unix 毫秒时间戳。
	ExpiresAt int64 `json:"expires_at"`
}

type _AgentPreviewReadyEventData AgentPreviewReadyEventData

// NewAgentPreviewReadyEventData instantiates a new AgentPreviewReadyEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPreviewReadyEventData(runId string, url string, expiresAt int64) *AgentPreviewReadyEventData {
	this := AgentPreviewReadyEventData{}
	this.RunId = runId
	this.Url = url
	this.ExpiresAt = expiresAt
	return &this
}

// NewAgentPreviewReadyEventDataWithDefaults instantiates a new AgentPreviewReadyEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPreviewReadyEventDataWithDefaults() *AgentPreviewReadyEventData {
	this := AgentPreviewReadyEventData{}
	return &this
}

// GetRunId returns the RunId field value
func (o *AgentPreviewReadyEventData) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AgentPreviewReadyEventData) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *AgentPreviewReadyEventData) SetRunId(v string) {
	o.RunId = v
}

// GetUrl returns the Url field value
func (o *AgentPreviewReadyEventData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AgentPreviewReadyEventData) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AgentPreviewReadyEventData) SetUrl(v string) {
	o.Url = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *AgentPreviewReadyEventData) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *AgentPreviewReadyEventData) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *AgentPreviewReadyEventData) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

func (o AgentPreviewReadyEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPreviewReadyEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["run_id"] = o.RunId
	toSerialize["url"] = o.Url
	toSerialize["expires_at"] = o.ExpiresAt
	return toSerialize, nil
}

func (o *AgentPreviewReadyEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"run_id",
		"url",
		"expires_at",
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

	varAgentPreviewReadyEventData := _AgentPreviewReadyEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentPreviewReadyEventData)

	if err != nil {
		return err
	}

	*o = AgentPreviewReadyEventData(varAgentPreviewReadyEventData)

	return err
}

type NullableAgentPreviewReadyEventData struct {
	value *AgentPreviewReadyEventData
	isSet bool
}

func (v NullableAgentPreviewReadyEventData) Get() *AgentPreviewReadyEventData {
	return v.value
}

func (v *NullableAgentPreviewReadyEventData) Set(val *AgentPreviewReadyEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPreviewReadyEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPreviewReadyEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPreviewReadyEventData(val *AgentPreviewReadyEventData) *NullableAgentPreviewReadyEventData {
	return &NullableAgentPreviewReadyEventData{value: val, isSet: true}
}

func (v NullableAgentPreviewReadyEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPreviewReadyEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

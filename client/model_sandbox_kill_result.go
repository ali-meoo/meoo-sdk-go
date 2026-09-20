package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SandboxKillResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxKillResult{}

// SandboxKillResult struct for SandboxKillResult
type SandboxKillResult struct {
	// 项目的公开 URL ID。
	ProjectId string `json:"project_id"`
	// 项目没有沙箱映射时为空字符串。
	SandboxId string `json:"sandbox_id"`
	// 销毁成功或平台确认沙箱不存在。
	Outcome string `json:"outcome"`
	// Unix 毫秒时间戳。
	ConfirmedAt int64 `json:"confirmed_at"`
}

type _SandboxKillResult SandboxKillResult

// NewSandboxKillResult instantiates a new SandboxKillResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxKillResult(projectId string, sandboxId string, outcome string, confirmedAt int64) *SandboxKillResult {
	this := SandboxKillResult{}
	this.ProjectId = projectId
	this.SandboxId = sandboxId
	this.Outcome = outcome
	this.ConfirmedAt = confirmedAt
	return &this
}

// NewSandboxKillResultWithDefaults instantiates a new SandboxKillResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxKillResultWithDefaults() *SandboxKillResult {
	this := SandboxKillResult{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *SandboxKillResult) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SandboxKillResult) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SandboxKillResult) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSandboxId returns the SandboxId field value
func (o *SandboxKillResult) GetSandboxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SandboxId
}

// GetSandboxIdOk returns a tuple with the SandboxId field value
// and a boolean to check if the value has been set.
func (o *SandboxKillResult) GetSandboxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SandboxId, true
}

// SetSandboxId sets field value
func (o *SandboxKillResult) SetSandboxId(v string) {
	o.SandboxId = v
}

// GetOutcome returns the Outcome field value
func (o *SandboxKillResult) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *SandboxKillResult) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *SandboxKillResult) SetOutcome(v string) {
	o.Outcome = v
}

// GetConfirmedAt returns the ConfirmedAt field value
func (o *SandboxKillResult) GetConfirmedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConfirmedAt
}

// GetConfirmedAtOk returns a tuple with the ConfirmedAt field value
// and a boolean to check if the value has been set.
func (o *SandboxKillResult) GetConfirmedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmedAt, true
}

// SetConfirmedAt sets field value
func (o *SandboxKillResult) SetConfirmedAt(v int64) {
	o.ConfirmedAt = v
}

func (o SandboxKillResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxKillResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["sandbox_id"] = o.SandboxId
	toSerialize["outcome"] = o.Outcome
	toSerialize["confirmed_at"] = o.ConfirmedAt
	return toSerialize, nil
}

func (o *SandboxKillResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
		"sandbox_id",
		"outcome",
		"confirmed_at",
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

	varSandboxKillResult := _SandboxKillResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxKillResult)

	if err != nil {
		return err
	}

	*o = SandboxKillResult(varSandboxKillResult)

	return err
}

type NullableSandboxKillResult struct {
	value *SandboxKillResult
	isSet bool
}

func (v NullableSandboxKillResult) Get() *SandboxKillResult {
	return v.value
}

func (v *NullableSandboxKillResult) Set(val *SandboxKillResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxKillResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxKillResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxKillResult(val *SandboxKillResult) *NullableSandboxKillResult {
	return &NullableSandboxKillResult{value: val, isSet: true}
}

func (v NullableSandboxKillResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxKillResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

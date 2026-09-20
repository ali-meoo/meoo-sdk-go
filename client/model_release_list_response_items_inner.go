package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ReleaseListResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseListResponseItemsInner{}

// ReleaseListResponseItemsInner struct for ReleaseListResponseItemsInner
type ReleaseListResponseItemsInner struct {
	ReleaseId  string `json:"release_id"`
	Version    int32  `json:"version"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"created_at"`
	FinishedAt int64  `json:"finished_at"`
	Current    bool   `json:"current"`
}

type _ReleaseListResponseItemsInner ReleaseListResponseItemsInner

// NewReleaseListResponseItemsInner instantiates a new ReleaseListResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseListResponseItemsInner(releaseId string, version int32, status string, createdAt int64, finishedAt int64, current bool) *ReleaseListResponseItemsInner {
	this := ReleaseListResponseItemsInner{}
	this.ReleaseId = releaseId
	this.Version = version
	this.Status = status
	this.CreatedAt = createdAt
	this.FinishedAt = finishedAt
	this.Current = current
	return &this
}

// NewReleaseListResponseItemsInnerWithDefaults instantiates a new ReleaseListResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseListResponseItemsInnerWithDefaults() *ReleaseListResponseItemsInner {
	this := ReleaseListResponseItemsInner{}
	return &this
}

// GetReleaseId returns the ReleaseId field value
func (o *ReleaseListResponseItemsInner) GetReleaseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponseItemsInner) GetReleaseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseId, true
}

// SetReleaseId sets field value
func (o *ReleaseListResponseItemsInner) SetReleaseId(v string) {
	o.ReleaseId = v
}

// GetVersion returns the Version field value
func (o *ReleaseListResponseItemsInner) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponseItemsInner) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ReleaseListResponseItemsInner) SetVersion(v int32) {
	o.Version = v
}

// GetStatus returns the Status field value
func (o *ReleaseListResponseItemsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponseItemsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReleaseListResponseItemsInner) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ReleaseListResponseItemsInner) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponseItemsInner) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ReleaseListResponseItemsInner) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *ReleaseListResponseItemsInner) GetFinishedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponseItemsInner) GetFinishedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *ReleaseListResponseItemsInner) SetFinishedAt(v int64) {
	o.FinishedAt = v
}

// GetCurrent returns the Current field value
func (o *ReleaseListResponseItemsInner) GetCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponseItemsInner) GetCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *ReleaseListResponseItemsInner) SetCurrent(v bool) {
	o.Current = v
}

func (o ReleaseListResponseItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseListResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["release_id"] = o.ReleaseId
	toSerialize["version"] = o.Version
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["finished_at"] = o.FinishedAt
	toSerialize["current"] = o.Current
	return toSerialize, nil
}

func (o *ReleaseListResponseItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"release_id",
		"version",
		"status",
		"created_at",
		"finished_at",
		"current",
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

	varReleaseListResponseItemsInner := _ReleaseListResponseItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReleaseListResponseItemsInner)

	if err != nil {
		return err
	}

	*o = ReleaseListResponseItemsInner(varReleaseListResponseItemsInner)

	return err
}

type NullableReleaseListResponseItemsInner struct {
	value *ReleaseListResponseItemsInner
	isSet bool
}

func (v NullableReleaseListResponseItemsInner) Get() *ReleaseListResponseItemsInner {
	return v.value
}

func (v *NullableReleaseListResponseItemsInner) Set(val *ReleaseListResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseListResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseListResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseListResponseItemsInner(val *ReleaseListResponseItemsInner) *NullableReleaseListResponseItemsInner {
	return &NullableReleaseListResponseItemsInner{value: val, isSet: true}
}

func (v NullableReleaseListResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseListResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

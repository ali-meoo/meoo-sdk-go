package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingPointsSummaryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPointsSummaryRequest{}

// BillingPointsSummaryRequest struct for BillingPointsSummaryRequest
type BillingPointsSummaryRequest struct {
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	StartTime string `json:"start_time"`
	// 主账号当前账期结束时间；NONE 计量项不受该账期范围限制。
	EndTime  string  `json:"end_time"`
	MemberId *string `json:"member_id,omitempty" validate:"regexp=^tm_[A-Za-z0-9_-]{20,32}$"`
}

type _BillingPointsSummaryRequest BillingPointsSummaryRequest

// NewBillingPointsSummaryRequest instantiates a new BillingPointsSummaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPointsSummaryRequest(startTime string, endTime string) *BillingPointsSummaryRequest {
	this := BillingPointsSummaryRequest{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewBillingPointsSummaryRequestWithDefaults instantiates a new BillingPointsSummaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPointsSummaryRequestWithDefaults() *BillingPointsSummaryRequest {
	this := BillingPointsSummaryRequest{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *BillingPointsSummaryRequest) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *BillingPointsSummaryRequest) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *BillingPointsSummaryRequest) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *BillingPointsSummaryRequest) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *BillingPointsSummaryRequest) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *BillingPointsSummaryRequest) SetEndTime(v string) {
	o.EndTime = v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *BillingPointsSummaryRequest) GetMemberId() string {
	if o == nil || IsNil(o.MemberId) {
		var ret string
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPointsSummaryRequest) GetMemberIdOk() (*string, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *BillingPointsSummaryRequest) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given string and assigns it to the MemberId field.
func (o *BillingPointsSummaryRequest) SetMemberId(v string) {
	o.MemberId = &v
}

func (o BillingPointsSummaryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPointsSummaryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	return toSerialize, nil
}

func (o *BillingPointsSummaryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start_time",
		"end_time",
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

	varBillingPointsSummaryRequest := _BillingPointsSummaryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingPointsSummaryRequest)

	if err != nil {
		return err
	}

	*o = BillingPointsSummaryRequest(varBillingPointsSummaryRequest)

	return err
}

type NullableBillingPointsSummaryRequest struct {
	value *BillingPointsSummaryRequest
	isSet bool
}

func (v NullableBillingPointsSummaryRequest) Get() *BillingPointsSummaryRequest {
	return v.value
}

func (v *NullableBillingPointsSummaryRequest) Set(val *BillingPointsSummaryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPointsSummaryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPointsSummaryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPointsSummaryRequest(val *BillingPointsSummaryRequest) *NullableBillingPointsSummaryRequest {
	return &NullableBillingPointsSummaryRequest{value: val, isSet: true}
}

func (v NullableBillingPointsSummaryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPointsSummaryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

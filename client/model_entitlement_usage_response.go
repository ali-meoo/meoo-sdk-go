package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EntitlementUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementUsageResponse{}

// EntitlementUsageResponse struct for EntitlementUsageResponse
type EntitlementUsageResponse struct {
	// 本次查询的公开成员 ID。
	MemberId string `json:"member_id"`
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	StartTime string `json:"start_time"`
	// 主账号当前账期结束时间；NONE 计量项不受该账期范围限制。
	EndTime string                 `json:"end_time"`
	Items   []EntitlementUsageItem `json:"items"`
}

type _EntitlementUsageResponse EntitlementUsageResponse

// NewEntitlementUsageResponse instantiates a new EntitlementUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementUsageResponse(memberId string, startTime string, endTime string, items []EntitlementUsageItem) *EntitlementUsageResponse {
	this := EntitlementUsageResponse{}
	this.MemberId = memberId
	this.StartTime = startTime
	this.EndTime = endTime
	this.Items = items
	return &this
}

// NewEntitlementUsageResponseWithDefaults instantiates a new EntitlementUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementUsageResponseWithDefaults() *EntitlementUsageResponse {
	this := EntitlementUsageResponse{}
	return &this
}

// GetMemberId returns the MemberId field value
func (o *EntitlementUsageResponse) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageResponse) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *EntitlementUsageResponse) SetMemberId(v string) {
	o.MemberId = v
}

// GetStartTime returns the StartTime field value
func (o *EntitlementUsageResponse) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageResponse) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *EntitlementUsageResponse) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *EntitlementUsageResponse) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageResponse) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *EntitlementUsageResponse) SetEndTime(v string) {
	o.EndTime = v
}

// GetItems returns the Items field value
func (o *EntitlementUsageResponse) GetItems() []EntitlementUsageItem {
	if o == nil {
		var ret []EntitlementUsageItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageResponse) GetItemsOk() ([]EntitlementUsageItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *EntitlementUsageResponse) SetItems(v []EntitlementUsageItem) {
	o.Items = v
}

func (o EntitlementUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["member_id"] = o.MemberId
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *EntitlementUsageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"member_id",
		"start_time",
		"end_time",
		"items",
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

	varEntitlementUsageResponse := _EntitlementUsageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntitlementUsageResponse)

	if err != nil {
		return err
	}

	*o = EntitlementUsageResponse(varEntitlementUsageResponse)

	return err
}

type NullableEntitlementUsageResponse struct {
	value *EntitlementUsageResponse
	isSet bool
}

func (v NullableEntitlementUsageResponse) Get() *EntitlementUsageResponse {
	return v.value
}

func (v *NullableEntitlementUsageResponse) Set(val *EntitlementUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementUsageResponse(val *EntitlementUsageResponse) *NullableEntitlementUsageResponse {
	return &NullableEntitlementUsageResponse{value: val, isSet: true}
}

func (v NullableEntitlementUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

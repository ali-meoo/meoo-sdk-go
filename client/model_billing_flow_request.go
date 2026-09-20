package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingFlowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingFlowRequest{}

// BillingFlowRequest struct for BillingFlowRequest
type BillingFlowRequest struct {
	// 按账单 period_start 筛选，含开始时间。
	StartTime string `json:"start_time"`
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	EndTime  string  `json:"end_time"`
	MemberId *string `json:"member_id,omitempty" validate:"regexp=^tm_[A-Za-z0-9_-]{20,32}$"`
	// 权益编码；按其配置的计量项过滤。
	EntitlementCode *string `json:"entitlement_code,omitempty"`
	// 页码从 0 开始。
	Page     *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
}

type _BillingFlowRequest BillingFlowRequest

// NewBillingFlowRequest instantiates a new BillingFlowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingFlowRequest(startTime string, endTime string) *BillingFlowRequest {
	this := BillingFlowRequest{}
	this.StartTime = startTime
	this.EndTime = endTime
	var page int32 = 0
	this.Page = &page
	var pageSize int32 = 20
	this.PageSize = &pageSize
	return &this
}

// NewBillingFlowRequestWithDefaults instantiates a new BillingFlowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingFlowRequestWithDefaults() *BillingFlowRequest {
	this := BillingFlowRequest{}
	var page int32 = 0
	this.Page = &page
	var pageSize int32 = 20
	this.PageSize = &pageSize
	return &this
}

// GetStartTime returns the StartTime field value
func (o *BillingFlowRequest) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *BillingFlowRequest) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *BillingFlowRequest) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *BillingFlowRequest) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *BillingFlowRequest) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *BillingFlowRequest) SetEndTime(v string) {
	o.EndTime = v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *BillingFlowRequest) GetMemberId() string {
	if o == nil || IsNil(o.MemberId) {
		var ret string
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFlowRequest) GetMemberIdOk() (*string, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *BillingFlowRequest) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given string and assigns it to the MemberId field.
func (o *BillingFlowRequest) SetMemberId(v string) {
	o.MemberId = &v
}

// GetEntitlementCode returns the EntitlementCode field value if set, zero value otherwise.
func (o *BillingFlowRequest) GetEntitlementCode() string {
	if o == nil || IsNil(o.EntitlementCode) {
		var ret string
		return ret
	}
	return *o.EntitlementCode
}

// GetEntitlementCodeOk returns a tuple with the EntitlementCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFlowRequest) GetEntitlementCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementCode) {
		return nil, false
	}
	return o.EntitlementCode, true
}

// HasEntitlementCode returns a boolean if a field has been set.
func (o *BillingFlowRequest) HasEntitlementCode() bool {
	if o != nil && !IsNil(o.EntitlementCode) {
		return true
	}

	return false
}

// SetEntitlementCode gets a reference to the given string and assigns it to the EntitlementCode field.
func (o *BillingFlowRequest) SetEntitlementCode(v string) {
	o.EntitlementCode = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *BillingFlowRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFlowRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *BillingFlowRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *BillingFlowRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *BillingFlowRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFlowRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *BillingFlowRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *BillingFlowRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o BillingFlowRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingFlowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	if !IsNil(o.EntitlementCode) {
		toSerialize["entitlement_code"] = o.EntitlementCode
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

func (o *BillingFlowRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBillingFlowRequest := _BillingFlowRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingFlowRequest)

	if err != nil {
		return err
	}

	*o = BillingFlowRequest(varBillingFlowRequest)

	return err
}

type NullableBillingFlowRequest struct {
	value *BillingFlowRequest
	isSet bool
}

func (v NullableBillingFlowRequest) Get() *BillingFlowRequest {
	return v.value
}

func (v *NullableBillingFlowRequest) Set(val *BillingFlowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingFlowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingFlowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingFlowRequest(val *BillingFlowRequest) *NullableBillingFlowRequest {
	return &NullableBillingFlowRequest{value: val, isSet: true}
}

func (v NullableBillingFlowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingFlowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

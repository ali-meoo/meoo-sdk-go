package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingPointsSummaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPointsSummaryResponse{}

// BillingPointsSummaryResponse struct for BillingPointsSummaryResponse
type BillingPointsSummaryResponse struct {
	// 查询范围内已出账积分合计。
	TotalPoints int32 `json:"total_points"`
	// 已出账账单中已完成扣减的积分合计。
	ActualDeducted int32 `json:"actual_deducted"`
	// 已出账未扣减积分，可为负数，不做钳制。
	PendingPoints int32 `json:"pending_points"`
}

type _BillingPointsSummaryResponse BillingPointsSummaryResponse

// NewBillingPointsSummaryResponse instantiates a new BillingPointsSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPointsSummaryResponse(totalPoints int32, actualDeducted int32, pendingPoints int32) *BillingPointsSummaryResponse {
	this := BillingPointsSummaryResponse{}
	this.TotalPoints = totalPoints
	this.ActualDeducted = actualDeducted
	this.PendingPoints = pendingPoints
	return &this
}

// NewBillingPointsSummaryResponseWithDefaults instantiates a new BillingPointsSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPointsSummaryResponseWithDefaults() *BillingPointsSummaryResponse {
	this := BillingPointsSummaryResponse{}
	return &this
}

// GetTotalPoints returns the TotalPoints field value
func (o *BillingPointsSummaryResponse) GetTotalPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPoints
}

// GetTotalPointsOk returns a tuple with the TotalPoints field value
// and a boolean to check if the value has been set.
func (o *BillingPointsSummaryResponse) GetTotalPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPoints, true
}

// SetTotalPoints sets field value
func (o *BillingPointsSummaryResponse) SetTotalPoints(v int32) {
	o.TotalPoints = v
}

// GetActualDeducted returns the ActualDeducted field value
func (o *BillingPointsSummaryResponse) GetActualDeducted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActualDeducted
}

// GetActualDeductedOk returns a tuple with the ActualDeducted field value
// and a boolean to check if the value has been set.
func (o *BillingPointsSummaryResponse) GetActualDeductedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualDeducted, true
}

// SetActualDeducted sets field value
func (o *BillingPointsSummaryResponse) SetActualDeducted(v int32) {
	o.ActualDeducted = v
}

// GetPendingPoints returns the PendingPoints field value
func (o *BillingPointsSummaryResponse) GetPendingPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PendingPoints
}

// GetPendingPointsOk returns a tuple with the PendingPoints field value
// and a boolean to check if the value has been set.
func (o *BillingPointsSummaryResponse) GetPendingPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingPoints, true
}

// SetPendingPoints sets field value
func (o *BillingPointsSummaryResponse) SetPendingPoints(v int32) {
	o.PendingPoints = v
}

func (o BillingPointsSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPointsSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_points"] = o.TotalPoints
	toSerialize["actual_deducted"] = o.ActualDeducted
	toSerialize["pending_points"] = o.PendingPoints
	return toSerialize, nil
}

func (o *BillingPointsSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_points",
		"actual_deducted",
		"pending_points",
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

	varBillingPointsSummaryResponse := _BillingPointsSummaryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingPointsSummaryResponse)

	if err != nil {
		return err
	}

	*o = BillingPointsSummaryResponse(varBillingPointsSummaryResponse)

	return err
}

type NullableBillingPointsSummaryResponse struct {
	value *BillingPointsSummaryResponse
	isSet bool
}

func (v NullableBillingPointsSummaryResponse) Get() *BillingPointsSummaryResponse {
	return v.value
}

func (v *NullableBillingPointsSummaryResponse) Set(val *BillingPointsSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPointsSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPointsSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPointsSummaryResponse(val *BillingPointsSummaryResponse) *NullableBillingPointsSummaryResponse {
	return &NullableBillingPointsSummaryResponse{value: val, isSet: true}
}

func (v NullableBillingPointsSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPointsSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

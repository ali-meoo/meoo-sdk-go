package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingFlowItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingFlowItem{}

// BillingFlowItem struct for BillingFlowItem
type BillingFlowItem struct {
	BillNo    string `json:"bill_no"`
	MemberId  string `json:"member_id" validate:"regexp=^tm_[A-Za-z0-9_-]{20,32}$"`
	ProjectId string `json:"project_id" validate:"regexp=^[A-Za-z0-9_-]+$"`
	BillType  string `json:"bill_type"`
	// 计量项编码，兼容保留；展示使用权益编码和名称。
	ItemCode string `json:"item_code"`
	// 权益编码；按其配置的计量项过滤。
	EntitlementCode string `json:"entitlement_code"`
	// 当前权益名称，非出账时快照；未配置映射时为 null。
	EntitlementName NullableString `json:"entitlement_name"`
	MeteringType    NullableString `json:"metering_type"`
	Source          NullableString `json:"source"`
	// 账单明细用量；无明细时为 null，不补零。
	Quantity       NullableFloat32 `json:"quantity"`
	Unit           NullableString  `json:"unit"`
	TotalPoints    int32           `json:"total_points"`
	ActualDeducted int32           `json:"actual_deducted"`
	// 包含 PENDING 等尚未完成扣减的账单，不按结算状态过滤。
	SettleStatus string `json:"settle_status"`
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	PeriodStart string `json:"period_start"`
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	PeriodEnd string `json:"period_end"`
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	SettleTime string `json:"settle_time"`
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	CreatedAt string `json:"created_at"`
}

type _BillingFlowItem BillingFlowItem

// NewBillingFlowItem instantiates a new BillingFlowItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingFlowItem(billNo string, memberId string, projectId string, billType string, itemCode string, entitlementCode string, entitlementName NullableString, meteringType NullableString, source NullableString, quantity NullableFloat32, unit NullableString, totalPoints int32, actualDeducted int32, settleStatus string, periodStart string, periodEnd string, settleTime string, createdAt string) *BillingFlowItem {
	this := BillingFlowItem{}
	this.BillNo = billNo
	this.MemberId = memberId
	this.ProjectId = projectId
	this.BillType = billType
	this.ItemCode = itemCode
	this.EntitlementCode = entitlementCode
	this.EntitlementName = entitlementName
	this.MeteringType = meteringType
	this.Source = source
	this.Quantity = quantity
	this.Unit = unit
	this.TotalPoints = totalPoints
	this.ActualDeducted = actualDeducted
	this.SettleStatus = settleStatus
	this.PeriodStart = periodStart
	this.PeriodEnd = periodEnd
	this.SettleTime = settleTime
	this.CreatedAt = createdAt
	return &this
}

// NewBillingFlowItemWithDefaults instantiates a new BillingFlowItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingFlowItemWithDefaults() *BillingFlowItem {
	this := BillingFlowItem{}
	return &this
}

// GetBillNo returns the BillNo field value
func (o *BillingFlowItem) GetBillNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillNo
}

// GetBillNoOk returns a tuple with the BillNo field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetBillNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillNo, true
}

// SetBillNo sets field value
func (o *BillingFlowItem) SetBillNo(v string) {
	o.BillNo = v
}

// GetMemberId returns the MemberId field value
func (o *BillingFlowItem) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *BillingFlowItem) SetMemberId(v string) {
	o.MemberId = v
}

// GetProjectId returns the ProjectId field value
func (o *BillingFlowItem) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *BillingFlowItem) SetProjectId(v string) {
	o.ProjectId = v
}

// GetBillType returns the BillType field value
func (o *BillingFlowItem) GetBillType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillType
}

// GetBillTypeOk returns a tuple with the BillType field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetBillTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillType, true
}

// SetBillType sets field value
func (o *BillingFlowItem) SetBillType(v string) {
	o.BillType = v
}

// GetItemCode returns the ItemCode field value
func (o *BillingFlowItem) GetItemCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemCode
}

// GetItemCodeOk returns a tuple with the ItemCode field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetItemCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCode, true
}

// SetItemCode sets field value
func (o *BillingFlowItem) SetItemCode(v string) {
	o.ItemCode = v
}

// GetEntitlementCode returns the EntitlementCode field value
func (o *BillingFlowItem) GetEntitlementCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementCode
}

// GetEntitlementCodeOk returns a tuple with the EntitlementCode field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetEntitlementCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementCode, true
}

// SetEntitlementCode sets field value
func (o *BillingFlowItem) SetEntitlementCode(v string) {
	o.EntitlementCode = v
}

// GetEntitlementName returns the EntitlementName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingFlowItem) GetEntitlementName() string {
	if o == nil || o.EntitlementName.Get() == nil {
		var ret string
		return ret
	}

	return *o.EntitlementName.Get()
}

// GetEntitlementNameOk returns a tuple with the EntitlementName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingFlowItem) GetEntitlementNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementName.Get(), o.EntitlementName.IsSet()
}

// SetEntitlementName sets field value
func (o *BillingFlowItem) SetEntitlementName(v string) {
	o.EntitlementName.Set(&v)
}

// GetMeteringType returns the MeteringType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingFlowItem) GetMeteringType() string {
	if o == nil || o.MeteringType.Get() == nil {
		var ret string
		return ret
	}

	return *o.MeteringType.Get()
}

// GetMeteringTypeOk returns a tuple with the MeteringType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingFlowItem) GetMeteringTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeteringType.Get(), o.MeteringType.IsSet()
}

// SetMeteringType sets field value
func (o *BillingFlowItem) SetMeteringType(v string) {
	o.MeteringType.Set(&v)
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingFlowItem) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingFlowItem) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *BillingFlowItem) SetSource(v string) {
	o.Source.Set(&v)
}

// GetQuantity returns the Quantity field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *BillingFlowItem) GetQuantity() float32 {
	if o == nil || o.Quantity.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingFlowItem) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// SetQuantity sets field value
func (o *BillingFlowItem) SetQuantity(v float32) {
	o.Quantity.Set(&v)
}

// GetUnit returns the Unit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingFlowItem) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}

	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingFlowItem) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// SetUnit sets field value
func (o *BillingFlowItem) SetUnit(v string) {
	o.Unit.Set(&v)
}

// GetTotalPoints returns the TotalPoints field value
func (o *BillingFlowItem) GetTotalPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPoints
}

// GetTotalPointsOk returns a tuple with the TotalPoints field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetTotalPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPoints, true
}

// SetTotalPoints sets field value
func (o *BillingFlowItem) SetTotalPoints(v int32) {
	o.TotalPoints = v
}

// GetActualDeducted returns the ActualDeducted field value
func (o *BillingFlowItem) GetActualDeducted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActualDeducted
}

// GetActualDeductedOk returns a tuple with the ActualDeducted field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetActualDeductedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualDeducted, true
}

// SetActualDeducted sets field value
func (o *BillingFlowItem) SetActualDeducted(v int32) {
	o.ActualDeducted = v
}

// GetSettleStatus returns the SettleStatus field value
func (o *BillingFlowItem) GetSettleStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettleStatus
}

// GetSettleStatusOk returns a tuple with the SettleStatus field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetSettleStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettleStatus, true
}

// SetSettleStatus sets field value
func (o *BillingFlowItem) SetSettleStatus(v string) {
	o.SettleStatus = v
}

// GetPeriodStart returns the PeriodStart field value
func (o *BillingFlowItem) GetPeriodStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetPeriodStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodStart, true
}

// SetPeriodStart sets field value
func (o *BillingFlowItem) SetPeriodStart(v string) {
	o.PeriodStart = v
}

// GetPeriodEnd returns the PeriodEnd field value
func (o *BillingFlowItem) GetPeriodEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodEnd, true
}

// SetPeriodEnd sets field value
func (o *BillingFlowItem) SetPeriodEnd(v string) {
	o.PeriodEnd = v
}

// GetSettleTime returns the SettleTime field value
func (o *BillingFlowItem) GetSettleTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettleTime
}

// GetSettleTimeOk returns a tuple with the SettleTime field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetSettleTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettleTime, true
}

// SetSettleTime sets field value
func (o *BillingFlowItem) SetSettleTime(v string) {
	o.SettleTime = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BillingFlowItem) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingFlowItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BillingFlowItem) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o BillingFlowItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingFlowItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bill_no"] = o.BillNo
	toSerialize["member_id"] = o.MemberId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["bill_type"] = o.BillType
	toSerialize["item_code"] = o.ItemCode
	toSerialize["entitlement_code"] = o.EntitlementCode
	toSerialize["entitlement_name"] = o.EntitlementName.Get()
	toSerialize["metering_type"] = o.MeteringType.Get()
	toSerialize["source"] = o.Source.Get()
	toSerialize["quantity"] = o.Quantity.Get()
	toSerialize["unit"] = o.Unit.Get()
	toSerialize["total_points"] = o.TotalPoints
	toSerialize["actual_deducted"] = o.ActualDeducted
	toSerialize["settle_status"] = o.SettleStatus
	toSerialize["period_start"] = o.PeriodStart
	toSerialize["period_end"] = o.PeriodEnd
	toSerialize["settle_time"] = o.SettleTime
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *BillingFlowItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bill_no",
		"member_id",
		"project_id",
		"bill_type",
		"item_code",
		"entitlement_code",
		"entitlement_name",
		"metering_type",
		"source",
		"quantity",
		"unit",
		"total_points",
		"actual_deducted",
		"settle_status",
		"period_start",
		"period_end",
		"settle_time",
		"created_at",
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

	varBillingFlowItem := _BillingFlowItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingFlowItem)

	if err != nil {
		return err
	}

	*o = BillingFlowItem(varBillingFlowItem)

	return err
}

type NullableBillingFlowItem struct {
	value *BillingFlowItem
	isSet bool
}

func (v NullableBillingFlowItem) Get() *BillingFlowItem {
	return v.value
}

func (v *NullableBillingFlowItem) Set(val *BillingFlowItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingFlowItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingFlowItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingFlowItem(val *BillingFlowItem) *NullableBillingFlowItem {
	return &NullableBillingFlowItem{value: val, isSet: true}
}

func (v NullableBillingFlowItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingFlowItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

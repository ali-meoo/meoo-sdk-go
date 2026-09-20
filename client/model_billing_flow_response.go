package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingFlowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingFlowResponse{}

// BillingFlowResponse struct for BillingFlowResponse
type BillingFlowResponse struct {
	Items []BillingFlowItem `json:"items"`
	Total int32             `json:"total"`
	// 页码从 0 开始。
	Page     *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
}

type _BillingFlowResponse BillingFlowResponse

// NewBillingFlowResponse instantiates a new BillingFlowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingFlowResponse(items []BillingFlowItem, total int32) *BillingFlowResponse {
	this := BillingFlowResponse{}
	this.Items = items
	this.Total = total
	var page int32 = 0
	this.Page = &page
	var pageSize int32 = 20
	this.PageSize = &pageSize
	return &this
}

// NewBillingFlowResponseWithDefaults instantiates a new BillingFlowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingFlowResponseWithDefaults() *BillingFlowResponse {
	this := BillingFlowResponse{}
	var page int32 = 0
	this.Page = &page
	var pageSize int32 = 20
	this.PageSize = &pageSize
	return &this
}

// GetItems returns the Items field value
func (o *BillingFlowResponse) GetItems() []BillingFlowItem {
	if o == nil {
		var ret []BillingFlowItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *BillingFlowResponse) GetItemsOk() ([]BillingFlowItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *BillingFlowResponse) SetItems(v []BillingFlowItem) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *BillingFlowResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *BillingFlowResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *BillingFlowResponse) SetTotal(v int32) {
	o.Total = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *BillingFlowResponse) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFlowResponse) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *BillingFlowResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *BillingFlowResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *BillingFlowResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFlowResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *BillingFlowResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *BillingFlowResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o BillingFlowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingFlowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["total"] = o.Total
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

func (o *BillingFlowResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"total",
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

	varBillingFlowResponse := _BillingFlowResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingFlowResponse)

	if err != nil {
		return err
	}

	*o = BillingFlowResponse(varBillingFlowResponse)

	return err
}

type NullableBillingFlowResponse struct {
	value *BillingFlowResponse
	isSet bool
}

func (v NullableBillingFlowResponse) Get() *BillingFlowResponse {
	return v.value
}

func (v *NullableBillingFlowResponse) Set(val *BillingFlowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingFlowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingFlowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingFlowResponse(val *BillingFlowResponse) *NullableBillingFlowResponse {
	return &NullableBillingFlowResponse{value: val, isSet: true}
}

func (v NullableBillingFlowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingFlowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EntitlementUsageItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementUsageItem{}

// EntitlementUsageItem struct for EntitlementUsageItem
type EntitlementUsageItem struct {
	Code string `json:"code"`
	// 权益展示名。
	Name string `json:"name"`
	// 真实用量；MONTHLY 使用主账号当前账期，其余按 Metering 原有语义。
	Used float32 `json:"used"`
	Unit string  `json:"unit"`
}

type _EntitlementUsageItem EntitlementUsageItem

// NewEntitlementUsageItem instantiates a new EntitlementUsageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementUsageItem(code string, name string, used float32, unit string) *EntitlementUsageItem {
	this := EntitlementUsageItem{}
	this.Code = code
	this.Name = name
	this.Used = used
	this.Unit = unit
	return &this
}

// NewEntitlementUsageItemWithDefaults instantiates a new EntitlementUsageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementUsageItemWithDefaults() *EntitlementUsageItem {
	this := EntitlementUsageItem{}
	return &this
}

// GetCode returns the Code field value
func (o *EntitlementUsageItem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageItem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EntitlementUsageItem) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *EntitlementUsageItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementUsageItem) SetName(v string) {
	o.Name = v
}

// GetUsed returns the Used field value
func (o *EntitlementUsageItem) GetUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageItem) GetUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *EntitlementUsageItem) SetUsed(v float32) {
	o.Used = v
}

// GetUnit returns the Unit field value
func (o *EntitlementUsageItem) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageItem) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *EntitlementUsageItem) SetUnit(v string) {
	o.Unit = v
}

func (o EntitlementUsageItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementUsageItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	toSerialize["used"] = o.Used
	toSerialize["unit"] = o.Unit
	return toSerialize, nil
}

func (o *EntitlementUsageItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
		"used",
		"unit",
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

	varEntitlementUsageItem := _EntitlementUsageItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntitlementUsageItem)

	if err != nil {
		return err
	}

	*o = EntitlementUsageItem(varEntitlementUsageItem)

	return err
}

type NullableEntitlementUsageItem struct {
	value *EntitlementUsageItem
	isSet bool
}

func (v NullableEntitlementUsageItem) Get() *EntitlementUsageItem {
	return v.value
}

func (v *NullableEntitlementUsageItem) Set(val *EntitlementUsageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementUsageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementUsageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementUsageItem(val *EntitlementUsageItem) *NullableEntitlementUsageItem {
	return &NullableEntitlementUsageItem{value: val, isSet: true}
}

func (v NullableEntitlementUsageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementUsageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

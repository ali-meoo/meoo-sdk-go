package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EntitlementUsageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementUsageRequest{}

// EntitlementUsageRequest struct for EntitlementUsageRequest
type EntitlementUsageRequest struct {
	// 凭证所属团队的公开成员 ID，冻结或已删除成员亦可查询。
	MemberId string `json:"member_id" validate:"regexp=^tm_[A-Za-z0-9_-]{20,32}$"`
	// 省略或空数组查询全部白名单；重复项去重，按请求顺序返回。
	Codes []string `json:"codes,omitempty"`
}

type _EntitlementUsageRequest EntitlementUsageRequest

// NewEntitlementUsageRequest instantiates a new EntitlementUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementUsageRequest(memberId string) *EntitlementUsageRequest {
	this := EntitlementUsageRequest{}
	this.MemberId = memberId
	return &this
}

// NewEntitlementUsageRequestWithDefaults instantiates a new EntitlementUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementUsageRequestWithDefaults() *EntitlementUsageRequest {
	this := EntitlementUsageRequest{}
	return &this
}

// GetMemberId returns the MemberId field value
func (o *EntitlementUsageRequest) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsageRequest) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *EntitlementUsageRequest) SetMemberId(v string) {
	o.MemberId = v
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *EntitlementUsageRequest) GetCodes() []string {
	if o == nil || IsNil(o.Codes) {
		var ret []string
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementUsageRequest) GetCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *EntitlementUsageRequest) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []string and assigns it to the Codes field.
func (o *EntitlementUsageRequest) SetCodes(v []string) {
	o.Codes = v
}

func (o EntitlementUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementUsageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["member_id"] = o.MemberId
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	return toSerialize, nil
}

func (o *EntitlementUsageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"member_id",
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

	varEntitlementUsageRequest := _EntitlementUsageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntitlementUsageRequest)

	if err != nil {
		return err
	}

	*o = EntitlementUsageRequest(varEntitlementUsageRequest)

	return err
}

type NullableEntitlementUsageRequest struct {
	value *EntitlementUsageRequest
	isSet bool
}

func (v NullableEntitlementUsageRequest) Get() *EntitlementUsageRequest {
	return v.value
}

func (v *NullableEntitlementUsageRequest) Set(val *EntitlementUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementUsageRequest(val *EntitlementUsageRequest) *NullableEntitlementUsageRequest {
	return &NullableEntitlementUsageRequest{value: val, isSet: true}
}

func (v NullableEntitlementUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

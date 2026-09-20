package client

import (
	"encoding/json"
)

// checks if the TeamEntitlementUsageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamEntitlementUsageRequest{}

// TeamEntitlementUsageRequest struct for TeamEntitlementUsageRequest
type TeamEntitlementUsageRequest struct {
	// 省略或空数组查询全部白名单；重复项去重，按请求顺序返回。
	Codes []string `json:"codes,omitempty"`
}

// NewTeamEntitlementUsageRequest instantiates a new TeamEntitlementUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamEntitlementUsageRequest() *TeamEntitlementUsageRequest {
	this := TeamEntitlementUsageRequest{}
	return &this
}

// NewTeamEntitlementUsageRequestWithDefaults instantiates a new TeamEntitlementUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamEntitlementUsageRequestWithDefaults() *TeamEntitlementUsageRequest {
	this := TeamEntitlementUsageRequest{}
	return &this
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *TeamEntitlementUsageRequest) GetCodes() []string {
	if o == nil || IsNil(o.Codes) {
		var ret []string
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamEntitlementUsageRequest) GetCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *TeamEntitlementUsageRequest) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []string and assigns it to the Codes field.
func (o *TeamEntitlementUsageRequest) SetCodes(v []string) {
	o.Codes = v
}

func (o TeamEntitlementUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamEntitlementUsageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	return toSerialize, nil
}

type NullableTeamEntitlementUsageRequest struct {
	value *TeamEntitlementUsageRequest
	isSet bool
}

func (v NullableTeamEntitlementUsageRequest) Get() *TeamEntitlementUsageRequest {
	return v.value
}

func (v *NullableTeamEntitlementUsageRequest) Set(val *TeamEntitlementUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamEntitlementUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamEntitlementUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamEntitlementUsageRequest(val *TeamEntitlementUsageRequest) *NullableTeamEntitlementUsageRequest {
	return &NullableTeamEntitlementUsageRequest{value: val, isSet: true}
}

func (v NullableTeamEntitlementUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamEntitlementUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

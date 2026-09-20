package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TeamEntitlementUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamEntitlementUsageResponse{}

// TeamEntitlementUsageResponse struct for TeamEntitlementUsageResponse
type TeamEntitlementUsageResponse struct {
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	StartTime string `json:"start_time"`
	// 主账号当前账期结束时间；NONE 计量项不受该账期范围限制。
	EndTime string                 `json:"end_time"`
	Items   []EntitlementUsageItem `json:"items"`
}

type _TeamEntitlementUsageResponse TeamEntitlementUsageResponse

// NewTeamEntitlementUsageResponse instantiates a new TeamEntitlementUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamEntitlementUsageResponse(startTime string, endTime string, items []EntitlementUsageItem) *TeamEntitlementUsageResponse {
	this := TeamEntitlementUsageResponse{}
	this.StartTime = startTime
	this.EndTime = endTime
	this.Items = items
	return &this
}

// NewTeamEntitlementUsageResponseWithDefaults instantiates a new TeamEntitlementUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamEntitlementUsageResponseWithDefaults() *TeamEntitlementUsageResponse {
	this := TeamEntitlementUsageResponse{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *TeamEntitlementUsageResponse) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *TeamEntitlementUsageResponse) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *TeamEntitlementUsageResponse) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *TeamEntitlementUsageResponse) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *TeamEntitlementUsageResponse) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *TeamEntitlementUsageResponse) SetEndTime(v string) {
	o.EndTime = v
}

// GetItems returns the Items field value
func (o *TeamEntitlementUsageResponse) GetItems() []EntitlementUsageItem {
	if o == nil {
		var ret []EntitlementUsageItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TeamEntitlementUsageResponse) GetItemsOk() ([]EntitlementUsageItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *TeamEntitlementUsageResponse) SetItems(v []EntitlementUsageItem) {
	o.Items = v
}

func (o TeamEntitlementUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamEntitlementUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *TeamEntitlementUsageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varTeamEntitlementUsageResponse := _TeamEntitlementUsageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamEntitlementUsageResponse)

	if err != nil {
		return err
	}

	*o = TeamEntitlementUsageResponse(varTeamEntitlementUsageResponse)

	return err
}

type NullableTeamEntitlementUsageResponse struct {
	value *TeamEntitlementUsageResponse
	isSet bool
}

func (v NullableTeamEntitlementUsageResponse) Get() *TeamEntitlementUsageResponse {
	return v.value
}

func (v *NullableTeamEntitlementUsageResponse) Set(val *TeamEntitlementUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamEntitlementUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamEntitlementUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamEntitlementUsageResponse(val *TeamEntitlementUsageResponse) *NullableTeamEntitlementUsageResponse {
	return &NullableTeamEntitlementUsageResponse{value: val, isSet: true}
}

func (v NullableTeamEntitlementUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamEntitlementUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

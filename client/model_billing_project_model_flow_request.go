package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingProjectModelFlowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingProjectModelFlowRequest{}

// BillingProjectModelFlowRequest struct for BillingProjectModelFlowRequest
type BillingProjectModelFlowRequest struct {
	// 主账号当前账期开始时间，保留 Console 返回的本地时间原值。
	StartTime string `json:"start_time"`
	// 主账号当前账期结束时间；NONE 计量项不受该账期范围限制。
	EndTime   string `json:"end_time"`
	ProjectId string `json:"project_id" validate:"regexp=^[A-Za-z0-9_-]+$"`
	// 页码从 0 开始。
	Page     *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
}

type _BillingProjectModelFlowRequest BillingProjectModelFlowRequest

// NewBillingProjectModelFlowRequest instantiates a new BillingProjectModelFlowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingProjectModelFlowRequest(startTime string, endTime string, projectId string) *BillingProjectModelFlowRequest {
	this := BillingProjectModelFlowRequest{}
	this.StartTime = startTime
	this.EndTime = endTime
	this.ProjectId = projectId
	var page int32 = 0
	this.Page = &page
	var pageSize int32 = 20
	this.PageSize = &pageSize
	return &this
}

// NewBillingProjectModelFlowRequestWithDefaults instantiates a new BillingProjectModelFlowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingProjectModelFlowRequestWithDefaults() *BillingProjectModelFlowRequest {
	this := BillingProjectModelFlowRequest{}
	var page int32 = 0
	this.Page = &page
	var pageSize int32 = 20
	this.PageSize = &pageSize
	return &this
}

// GetStartTime returns the StartTime field value
func (o *BillingProjectModelFlowRequest) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *BillingProjectModelFlowRequest) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *BillingProjectModelFlowRequest) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *BillingProjectModelFlowRequest) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *BillingProjectModelFlowRequest) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *BillingProjectModelFlowRequest) SetEndTime(v string) {
	o.EndTime = v
}

// GetProjectId returns the ProjectId field value
func (o *BillingProjectModelFlowRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *BillingProjectModelFlowRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *BillingProjectModelFlowRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *BillingProjectModelFlowRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingProjectModelFlowRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *BillingProjectModelFlowRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *BillingProjectModelFlowRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *BillingProjectModelFlowRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingProjectModelFlowRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *BillingProjectModelFlowRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *BillingProjectModelFlowRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o BillingProjectModelFlowRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingProjectModelFlowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	toSerialize["project_id"] = o.ProjectId
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

func (o *BillingProjectModelFlowRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start_time",
		"end_time",
		"project_id",
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

	varBillingProjectModelFlowRequest := _BillingProjectModelFlowRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingProjectModelFlowRequest)

	if err != nil {
		return err
	}

	*o = BillingProjectModelFlowRequest(varBillingProjectModelFlowRequest)

	return err
}

type NullableBillingProjectModelFlowRequest struct {
	value *BillingProjectModelFlowRequest
	isSet bool
}

func (v NullableBillingProjectModelFlowRequest) Get() *BillingProjectModelFlowRequest {
	return v.value
}

func (v *NullableBillingProjectModelFlowRequest) Set(val *BillingProjectModelFlowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingProjectModelFlowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingProjectModelFlowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingProjectModelFlowRequest(val *BillingProjectModelFlowRequest) *NullableBillingProjectModelFlowRequest {
	return &NullableBillingProjectModelFlowRequest{value: val, isSet: true}
}

func (v NullableBillingProjectModelFlowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingProjectModelFlowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

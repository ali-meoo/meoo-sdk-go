package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudFunctionLogsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunctionLogsResponse{}

// CloudFunctionLogsResponse struct for CloudFunctionLogsResponse
type CloudFunctionLogsResponse struct {
	FunctionName         string                                  `json:"function_name"`
	StartTime            int64                                   `json:"start_time"`
	EndTime              int64                                   `json:"end_time"`
	Total                int32                                   `json:"total"`
	Returned             int32                                   `json:"returned"`
	Entries              []CloudFunctionLogsResponseEntriesInner `json:"entries"`
	AdditionalProperties map[string]interface{}
}

type _CloudFunctionLogsResponse CloudFunctionLogsResponse

// NewCloudFunctionLogsResponse instantiates a new CloudFunctionLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunctionLogsResponse(functionName string, startTime int64, endTime int64, total int32, returned int32, entries []CloudFunctionLogsResponseEntriesInner) *CloudFunctionLogsResponse {
	this := CloudFunctionLogsResponse{}
	this.FunctionName = functionName
	this.StartTime = startTime
	this.EndTime = endTime
	this.Total = total
	this.Returned = returned
	this.Entries = entries
	return &this
}

// NewCloudFunctionLogsResponseWithDefaults instantiates a new CloudFunctionLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionLogsResponseWithDefaults() *CloudFunctionLogsResponse {
	this := CloudFunctionLogsResponse{}
	return &this
}

// GetFunctionName returns the FunctionName field value
func (o *CloudFunctionLogsResponse) GetFunctionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponse) GetFunctionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionName, true
}

// SetFunctionName sets field value
func (o *CloudFunctionLogsResponse) SetFunctionName(v string) {
	o.FunctionName = v
}

// GetStartTime returns the StartTime field value
func (o *CloudFunctionLogsResponse) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponse) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *CloudFunctionLogsResponse) SetStartTime(v int64) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *CloudFunctionLogsResponse) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponse) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *CloudFunctionLogsResponse) SetEndTime(v int64) {
	o.EndTime = v
}

// GetTotal returns the Total field value
func (o *CloudFunctionLogsResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CloudFunctionLogsResponse) SetTotal(v int32) {
	o.Total = v
}

// GetReturned returns the Returned field value
func (o *CloudFunctionLogsResponse) GetReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Returned
}

// GetReturnedOk returns a tuple with the Returned field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponse) GetReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Returned, true
}

// SetReturned sets field value
func (o *CloudFunctionLogsResponse) SetReturned(v int32) {
	o.Returned = v
}

// GetEntries returns the Entries field value
func (o *CloudFunctionLogsResponse) GetEntries() []CloudFunctionLogsResponseEntriesInner {
	if o == nil {
		var ret []CloudFunctionLogsResponseEntriesInner
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponse) GetEntriesOk() ([]CloudFunctionLogsResponseEntriesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *CloudFunctionLogsResponse) SetEntries(v []CloudFunctionLogsResponseEntriesInner) {
	o.Entries = v
}

func (o CloudFunctionLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunctionLogsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["function_name"] = o.FunctionName
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	toSerialize["total"] = o.Total
	toSerialize["returned"] = o.Returned
	toSerialize["entries"] = o.Entries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudFunctionLogsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"function_name",
		"start_time",
		"end_time",
		"total",
		"returned",
		"entries",
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

	varCloudFunctionLogsResponse := _CloudFunctionLogsResponse{}

	err = json.Unmarshal(data, &varCloudFunctionLogsResponse)

	if err != nil {
		return err
	}

	*o = CloudFunctionLogsResponse(varCloudFunctionLogsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "function_name")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "total")
		delete(additionalProperties, "returned")
		delete(additionalProperties, "entries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudFunctionLogsResponse struct {
	value *CloudFunctionLogsResponse
	isSet bool
}

func (v NullableCloudFunctionLogsResponse) Get() *CloudFunctionLogsResponse {
	return v.value
}

func (v *NullableCloudFunctionLogsResponse) Set(val *CloudFunctionLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunctionLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunctionLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunctionLogsResponse(val *CloudFunctionLogsResponse) *NullableCloudFunctionLogsResponse {
	return &NullableCloudFunctionLogsResponse{value: val, isSet: true}
}

func (v NullableCloudFunctionLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunctionLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

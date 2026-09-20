package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Problem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Problem{}

// Problem struct for Problem
type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int32  `json:"status"`
	Detail string `json:"detail"`
	// 公开错误码；全集与 src/open/contract/errorCodes.ts 注册表同步，客户端遇到未知取值时按 HTTP status 处理。
	Code    string `json:"code"`
	TraceId string `json:"trace_id"`
}

type _Problem Problem

// NewProblem instantiates a new Problem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblem(type_ string, title string, status int32, detail string, code string, traceId string) *Problem {
	this := Problem{}
	this.Type = type_
	this.Title = title
	this.Status = status
	this.Detail = detail
	this.Code = code
	this.TraceId = traceId
	return &this
}

// NewProblemWithDefaults instantiates a new Problem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemWithDefaults() *Problem {
	this := Problem{}
	return &this
}

// GetType returns the Type field value
func (o *Problem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Problem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Problem) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *Problem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Problem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Problem) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *Problem) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Problem) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Problem) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *Problem) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *Problem) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *Problem) SetDetail(v string) {
	o.Detail = v
}

// GetCode returns the Code field value
func (o *Problem) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Problem) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Problem) SetCode(v string) {
	o.Code = v
}

// GetTraceId returns the TraceId field value
func (o *Problem) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *Problem) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *Problem) SetTraceId(v string) {
	o.TraceId = v
}

func (o Problem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Problem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["status"] = o.Status
	toSerialize["detail"] = o.Detail
	toSerialize["code"] = o.Code
	toSerialize["trace_id"] = o.TraceId
	return toSerialize, nil
}

func (o *Problem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"title",
		"status",
		"detail",
		"code",
		"trace_id",
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

	varProblem := _Problem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varProblem)

	if err != nil {
		return err
	}

	*o = Problem(varProblem)

	return err
}

type NullableProblem struct {
	value *Problem
	isSet bool
}

func (v NullableProblem) Get() *Problem {
	return v.value
}

func (v *NullableProblem) Set(val *Problem) {
	v.value = val
	v.isSet = true
}

func (v NullableProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblem(val *Problem) *NullableProblem {
	return &NullableProblem{value: val, isSet: true}
}

func (v NullableProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

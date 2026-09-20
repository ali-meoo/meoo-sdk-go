package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudQueryResponse{}

// CloudQueryResponse struct for CloudQueryResponse
type CloudQueryResponse struct {
	Rows    []map[string]interface{} `json:"rows"`
	Columns []string                 `json:"columns"`
	// SQL 返回的总结果行数；上游提供受影响行数时也用于 DML，可能大于 rows 数量。
	RowCount int32 `json:"row_count"`
	// rows 是否因 100 行上限被截断。
	Truncated            bool `json:"truncated"`
	AdditionalProperties map[string]interface{}
}

type _CloudQueryResponse CloudQueryResponse

// NewCloudQueryResponse instantiates a new CloudQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudQueryResponse(rows []map[string]interface{}, columns []string, rowCount int32, truncated bool) *CloudQueryResponse {
	this := CloudQueryResponse{}
	this.Rows = rows
	this.Columns = columns
	this.RowCount = rowCount
	this.Truncated = truncated
	return &this
}

// NewCloudQueryResponseWithDefaults instantiates a new CloudQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudQueryResponseWithDefaults() *CloudQueryResponse {
	this := CloudQueryResponse{}
	return &this
}

// GetRows returns the Rows field value
func (o *CloudQueryResponse) GetRows() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *CloudQueryResponse) GetRowsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rows, true
}

// SetRows sets field value
func (o *CloudQueryResponse) SetRows(v []map[string]interface{}) {
	o.Rows = v
}

// GetColumns returns the Columns field value
func (o *CloudQueryResponse) GetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CloudQueryResponse) GetColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *CloudQueryResponse) SetColumns(v []string) {
	o.Columns = v
}

// GetRowCount returns the RowCount field value
func (o *CloudQueryResponse) GetRowCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value
// and a boolean to check if the value has been set.
func (o *CloudQueryResponse) GetRowCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowCount, true
}

// SetRowCount sets field value
func (o *CloudQueryResponse) SetRowCount(v int32) {
	o.RowCount = v
}

// GetTruncated returns the Truncated field value
func (o *CloudQueryResponse) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *CloudQueryResponse) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value
func (o *CloudQueryResponse) SetTruncated(v bool) {
	o.Truncated = v
}

func (o CloudQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rows"] = o.Rows
	toSerialize["columns"] = o.Columns
	toSerialize["row_count"] = o.RowCount
	toSerialize["truncated"] = o.Truncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudQueryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rows",
		"columns",
		"row_count",
		"truncated",
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

	varCloudQueryResponse := _CloudQueryResponse{}

	err = json.Unmarshal(data, &varCloudQueryResponse)

	if err != nil {
		return err
	}

	*o = CloudQueryResponse(varCloudQueryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rows")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "row_count")
		delete(additionalProperties, "truncated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudQueryResponse struct {
	value *CloudQueryResponse
	isSet bool
}

func (v NullableCloudQueryResponse) Get() *CloudQueryResponse {
	return v.value
}

func (v *NullableCloudQueryResponse) Set(val *CloudQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudQueryResponse(val *CloudQueryResponse) *NullableCloudQueryResponse {
	return &NullableCloudQueryResponse{value: val, isSet: true}
}

func (v NullableCloudQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

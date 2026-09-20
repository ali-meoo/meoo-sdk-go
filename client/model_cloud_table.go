package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudTable{}

// CloudTable struct for CloudTable
type CloudTable struct {
	Schema     string `json:"schema"`
	Name       string `json:"name"`
	RlsEnabled bool   `json:"rls_enabled"`
	RlsForced  bool   `json:"rls_forced"`
	// 字符串形式的估算行数，避免 JavaScript 整数精度损失。
	EstimatedRows        string                   `json:"estimated_rows" validate:"regexp=^[0-9]+$"`
	Comment              string                   `json:"comment"`
	Columns              []CloudTableColumnsInner `json:"columns"`
	AdditionalProperties map[string]interface{}
}

type _CloudTable CloudTable

// NewCloudTable instantiates a new CloudTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTable(schema string, name string, rlsEnabled bool, rlsForced bool, estimatedRows string, comment string, columns []CloudTableColumnsInner) *CloudTable {
	this := CloudTable{}
	this.Schema = schema
	this.Name = name
	this.RlsEnabled = rlsEnabled
	this.RlsForced = rlsForced
	this.EstimatedRows = estimatedRows
	this.Comment = comment
	this.Columns = columns
	return &this
}

// NewCloudTableWithDefaults instantiates a new CloudTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTableWithDefaults() *CloudTable {
	this := CloudTable{}
	return &this
}

// GetSchema returns the Schema field value
func (o *CloudTable) GetSchema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *CloudTable) SetSchema(v string) {
	o.Schema = v
}

// GetName returns the Name field value
func (o *CloudTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudTable) SetName(v string) {
	o.Name = v
}

// GetRlsEnabled returns the RlsEnabled field value
func (o *CloudTable) GetRlsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RlsEnabled
}

// GetRlsEnabledOk returns a tuple with the RlsEnabled field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetRlsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RlsEnabled, true
}

// SetRlsEnabled sets field value
func (o *CloudTable) SetRlsEnabled(v bool) {
	o.RlsEnabled = v
}

// GetRlsForced returns the RlsForced field value
func (o *CloudTable) GetRlsForced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RlsForced
}

// GetRlsForcedOk returns a tuple with the RlsForced field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetRlsForcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RlsForced, true
}

// SetRlsForced sets field value
func (o *CloudTable) SetRlsForced(v bool) {
	o.RlsForced = v
}

// GetEstimatedRows returns the EstimatedRows field value
func (o *CloudTable) GetEstimatedRows() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EstimatedRows
}

// GetEstimatedRowsOk returns a tuple with the EstimatedRows field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetEstimatedRowsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedRows, true
}

// SetEstimatedRows sets field value
func (o *CloudTable) SetEstimatedRows(v string) {
	o.EstimatedRows = v
}

// GetComment returns the Comment field value
func (o *CloudTable) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *CloudTable) SetComment(v string) {
	o.Comment = v
}

// GetColumns returns the Columns field value
func (o *CloudTable) GetColumns() []CloudTableColumnsInner {
	if o == nil {
		var ret []CloudTableColumnsInner
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CloudTable) GetColumnsOk() ([]CloudTableColumnsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *CloudTable) SetColumns(v []CloudTableColumnsInner) {
	o.Columns = v
}

func (o CloudTable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	toSerialize["name"] = o.Name
	toSerialize["rls_enabled"] = o.RlsEnabled
	toSerialize["rls_forced"] = o.RlsForced
	toSerialize["estimated_rows"] = o.EstimatedRows
	toSerialize["comment"] = o.Comment
	toSerialize["columns"] = o.Columns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudTable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schema",
		"name",
		"rls_enabled",
		"rls_forced",
		"estimated_rows",
		"comment",
		"columns",
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

	varCloudTable := _CloudTable{}

	err = json.Unmarshal(data, &varCloudTable)

	if err != nil {
		return err
	}

	*o = CloudTable(varCloudTable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "schema")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rls_enabled")
		delete(additionalProperties, "rls_forced")
		delete(additionalProperties, "estimated_rows")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "columns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudTable struct {
	value *CloudTable
	isSet bool
}

func (v NullableCloudTable) Get() *CloudTable {
	return v.value
}

func (v *NullableCloudTable) Set(val *CloudTable) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTable) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTable(val *CloudTable) *NullableCloudTable {
	return &NullableCloudTable{value: val, isSet: true}
}

func (v NullableCloudTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

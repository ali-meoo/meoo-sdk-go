package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudTableColumnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudTableColumnsInner{}

// CloudTableColumnsInner struct for CloudTableColumnsInner
type CloudTableColumnsInner struct {
	Name         string   `json:"name"`
	DataType     string   `json:"data_type"`
	Format       string   `json:"format"`
	Nullable     bool     `json:"nullable"`
	Updatable    bool     `json:"updatable"`
	Unique       bool     `json:"unique"`
	DefaultValue string   `json:"default_value"`
	Enums        []string `json:"enums"`
	Comment      string   `json:"comment"`
}

type _CloudTableColumnsInner CloudTableColumnsInner

// NewCloudTableColumnsInner instantiates a new CloudTableColumnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTableColumnsInner(name string, dataType string, format string, nullable bool, updatable bool, unique bool, defaultValue string, enums []string, comment string) *CloudTableColumnsInner {
	this := CloudTableColumnsInner{}
	this.Name = name
	this.DataType = dataType
	this.Format = format
	this.Nullable = nullable
	this.Updatable = updatable
	this.Unique = unique
	this.DefaultValue = defaultValue
	this.Enums = enums
	this.Comment = comment
	return &this
}

// NewCloudTableColumnsInnerWithDefaults instantiates a new CloudTableColumnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTableColumnsInnerWithDefaults() *CloudTableColumnsInner {
	this := CloudTableColumnsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CloudTableColumnsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudTableColumnsInner) SetName(v string) {
	o.Name = v
}

// GetDataType returns the DataType field value
func (o *CloudTableColumnsInner) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *CloudTableColumnsInner) SetDataType(v string) {
	o.DataType = v
}

// GetFormat returns the Format field value
func (o *CloudTableColumnsInner) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CloudTableColumnsInner) SetFormat(v string) {
	o.Format = v
}

// GetNullable returns the Nullable field value
func (o *CloudTableColumnsInner) GetNullable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Nullable
}

// GetNullableOk returns a tuple with the Nullable field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetNullableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nullable, true
}

// SetNullable sets field value
func (o *CloudTableColumnsInner) SetNullable(v bool) {
	o.Nullable = v
}

// GetUpdatable returns the Updatable field value
func (o *CloudTableColumnsInner) GetUpdatable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Updatable
}

// GetUpdatableOk returns a tuple with the Updatable field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetUpdatableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updatable, true
}

// SetUpdatable sets field value
func (o *CloudTableColumnsInner) SetUpdatable(v bool) {
	o.Updatable = v
}

// GetUnique returns the Unique field value
func (o *CloudTableColumnsInner) GetUnique() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unique, true
}

// SetUnique sets field value
func (o *CloudTableColumnsInner) SetUnique(v bool) {
	o.Unique = v
}

// GetDefaultValue returns the DefaultValue field value
func (o *CloudTableColumnsInner) GetDefaultValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *CloudTableColumnsInner) SetDefaultValue(v string) {
	o.DefaultValue = v
}

// GetEnums returns the Enums field value
func (o *CloudTableColumnsInner) GetEnums() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Enums
}

// GetEnumsOk returns a tuple with the Enums field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetEnumsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enums, true
}

// SetEnums sets field value
func (o *CloudTableColumnsInner) SetEnums(v []string) {
	o.Enums = v
}

// GetComment returns the Comment field value
func (o *CloudTableColumnsInner) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *CloudTableColumnsInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *CloudTableColumnsInner) SetComment(v string) {
	o.Comment = v
}

func (o CloudTableColumnsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudTableColumnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["data_type"] = o.DataType
	toSerialize["format"] = o.Format
	toSerialize["nullable"] = o.Nullable
	toSerialize["updatable"] = o.Updatable
	toSerialize["unique"] = o.Unique
	toSerialize["default_value"] = o.DefaultValue
	toSerialize["enums"] = o.Enums
	toSerialize["comment"] = o.Comment
	return toSerialize, nil
}

func (o *CloudTableColumnsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"data_type",
		"format",
		"nullable",
		"updatable",
		"unique",
		"default_value",
		"enums",
		"comment",
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

	varCloudTableColumnsInner := _CloudTableColumnsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudTableColumnsInner)

	if err != nil {
		return err
	}

	*o = CloudTableColumnsInner(varCloudTableColumnsInner)

	return err
}

type NullableCloudTableColumnsInner struct {
	value *CloudTableColumnsInner
	isSet bool
}

func (v NullableCloudTableColumnsInner) Get() *CloudTableColumnsInner {
	return v.value
}

func (v *NullableCloudTableColumnsInner) Set(val *CloudTableColumnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTableColumnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTableColumnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTableColumnsInner(val *CloudTableColumnsInner) *NullableCloudTableColumnsInner {
	return &NullableCloudTableColumnsInner{value: val, isSet: true}
}

func (v NullableCloudTableColumnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTableColumnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

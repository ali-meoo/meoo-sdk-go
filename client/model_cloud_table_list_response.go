package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudTableListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudTableListResponse{}

// CloudTableListResponse struct for CloudTableListResponse
type CloudTableListResponse struct {
	Tables []CloudTable `json:"tables"`
}

type _CloudTableListResponse CloudTableListResponse

// NewCloudTableListResponse instantiates a new CloudTableListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTableListResponse(tables []CloudTable) *CloudTableListResponse {
	this := CloudTableListResponse{}
	this.Tables = tables
	return &this
}

// NewCloudTableListResponseWithDefaults instantiates a new CloudTableListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTableListResponseWithDefaults() *CloudTableListResponse {
	this := CloudTableListResponse{}
	return &this
}

// GetTables returns the Tables field value
func (o *CloudTableListResponse) GetTables() []CloudTable {
	if o == nil {
		var ret []CloudTable
		return ret
	}

	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *CloudTableListResponse) GetTablesOk() ([]CloudTable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tables, true
}

// SetTables sets field value
func (o *CloudTableListResponse) SetTables(v []CloudTable) {
	o.Tables = v
}

func (o CloudTableListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudTableListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tables"] = o.Tables
	return toSerialize, nil
}

func (o *CloudTableListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tables",
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

	varCloudTableListResponse := _CloudTableListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudTableListResponse)

	if err != nil {
		return err
	}

	*o = CloudTableListResponse(varCloudTableListResponse)

	return err
}

type NullableCloudTableListResponse struct {
	value *CloudTableListResponse
	isSet bool
}

func (v NullableCloudTableListResponse) Get() *CloudTableListResponse {
	return v.value
}

func (v *NullableCloudTableListResponse) Set(val *CloudTableListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTableListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTableListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTableListResponse(val *CloudTableListResponse) *NullableCloudTableListResponse {
	return &NullableCloudTableListResponse{value: val, isSet: true}
}

func (v NullableCloudTableListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTableListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

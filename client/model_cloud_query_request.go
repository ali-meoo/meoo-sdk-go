package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudQueryRequest{}

// CloudQueryRequest struct for CloudQueryRequest
type CloudQueryRequest struct {
	// 先去除首尾空白，再按 UTF-8 字节数校验 65536 字节上限。
	Query string `json:"query"`
}

type _CloudQueryRequest CloudQueryRequest

// NewCloudQueryRequest instantiates a new CloudQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudQueryRequest(query string) *CloudQueryRequest {
	this := CloudQueryRequest{}
	this.Query = query
	return &this
}

// NewCloudQueryRequestWithDefaults instantiates a new CloudQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudQueryRequestWithDefaults() *CloudQueryRequest {
	this := CloudQueryRequest{}
	return &this
}

// GetQuery returns the Query field value
func (o *CloudQueryRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CloudQueryRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CloudQueryRequest) SetQuery(v string) {
	o.Query = v
}

func (o CloudQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

func (o *CloudQueryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
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

	varCloudQueryRequest := _CloudQueryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudQueryRequest)

	if err != nil {
		return err
	}

	*o = CloudQueryRequest(varCloudQueryRequest)

	return err
}

type NullableCloudQueryRequest struct {
	value *CloudQueryRequest
	isSet bool
}

func (v NullableCloudQueryRequest) Get() *CloudQueryRequest {
	return v.value
}

func (v *NullableCloudQueryRequest) Set(val *CloudQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudQueryRequest(val *CloudQueryRequest) *NullableCloudQueryRequest {
	return &NullableCloudQueryRequest{value: val, isSet: true}
}

func (v NullableCloudQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

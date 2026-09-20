package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudFunctionListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunctionListResponse{}

// CloudFunctionListResponse struct for CloudFunctionListResponse
type CloudFunctionListResponse struct {
	Functions []CloudFunction `json:"functions"`
}

type _CloudFunctionListResponse CloudFunctionListResponse

// NewCloudFunctionListResponse instantiates a new CloudFunctionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunctionListResponse(functions []CloudFunction) *CloudFunctionListResponse {
	this := CloudFunctionListResponse{}
	this.Functions = functions
	return &this
}

// NewCloudFunctionListResponseWithDefaults instantiates a new CloudFunctionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionListResponseWithDefaults() *CloudFunctionListResponse {
	this := CloudFunctionListResponse{}
	return &this
}

// GetFunctions returns the Functions field value
func (o *CloudFunctionListResponse) GetFunctions() []CloudFunction {
	if o == nil {
		var ret []CloudFunction
		return ret
	}

	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionListResponse) GetFunctionsOk() ([]CloudFunction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functions, true
}

// SetFunctions sets field value
func (o *CloudFunctionListResponse) SetFunctions(v []CloudFunction) {
	o.Functions = v
}

func (o CloudFunctionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunctionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["functions"] = o.Functions
	return toSerialize, nil
}

func (o *CloudFunctionListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"functions",
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

	varCloudFunctionListResponse := _CloudFunctionListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudFunctionListResponse)

	if err != nil {
		return err
	}

	*o = CloudFunctionListResponse(varCloudFunctionListResponse)

	return err
}

type NullableCloudFunctionListResponse struct {
	value *CloudFunctionListResponse
	isSet bool
}

func (v NullableCloudFunctionListResponse) Get() *CloudFunctionListResponse {
	return v.value
}

func (v *NullableCloudFunctionListResponse) Set(val *CloudFunctionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunctionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunctionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunctionListResponse(val *CloudFunctionListResponse) *NullableCloudFunctionListResponse {
	return &NullableCloudFunctionListResponse{value: val, isSet: true}
}

func (v NullableCloudFunctionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunctionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

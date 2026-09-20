package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudFunctionPutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunctionPutResponse{}

// CloudFunctionPutResponse struct for CloudFunctionPutResponse
type CloudFunctionPutResponse struct {
	Function CloudFunction `json:"function"`
}

type _CloudFunctionPutResponse CloudFunctionPutResponse

// NewCloudFunctionPutResponse instantiates a new CloudFunctionPutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunctionPutResponse(function CloudFunction) *CloudFunctionPutResponse {
	this := CloudFunctionPutResponse{}
	this.Function = function
	return &this
}

// NewCloudFunctionPutResponseWithDefaults instantiates a new CloudFunctionPutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionPutResponseWithDefaults() *CloudFunctionPutResponse {
	this := CloudFunctionPutResponse{}
	return &this
}

// GetFunction returns the Function field value
func (o *CloudFunctionPutResponse) GetFunction() CloudFunction {
	if o == nil {
		var ret CloudFunction
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionPutResponse) GetFunctionOk() (*CloudFunction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *CloudFunctionPutResponse) SetFunction(v CloudFunction) {
	o.Function = v
}

func (o CloudFunctionPutResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunctionPutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["function"] = o.Function
	return toSerialize, nil
}

func (o *CloudFunctionPutResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"function",
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

	varCloudFunctionPutResponse := _CloudFunctionPutResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudFunctionPutResponse)

	if err != nil {
		return err
	}

	*o = CloudFunctionPutResponse(varCloudFunctionPutResponse)

	return err
}

type NullableCloudFunctionPutResponse struct {
	value *CloudFunctionPutResponse
	isSet bool
}

func (v NullableCloudFunctionPutResponse) Get() *CloudFunctionPutResponse {
	return v.value
}

func (v *NullableCloudFunctionPutResponse) Set(val *CloudFunctionPutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunctionPutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunctionPutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunctionPutResponse(val *CloudFunctionPutResponse) *NullableCloudFunctionPutResponse {
	return &NullableCloudFunctionPutResponse{value: val, isSet: true}
}

func (v NullableCloudFunctionPutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunctionPutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudDatabaseStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudDatabaseStatusResponse{}

// CloudDatabaseStatusResponse struct for CloudDatabaseStatusResponse
type CloudDatabaseStatusResponse struct {
	// not_enabled 表示尚未开通；ready 表示已存在当前有效实例。
	Status string `json:"status"`
}

type _CloudDatabaseStatusResponse CloudDatabaseStatusResponse

// NewCloudDatabaseStatusResponse instantiates a new CloudDatabaseStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudDatabaseStatusResponse(status string) *CloudDatabaseStatusResponse {
	this := CloudDatabaseStatusResponse{}
	this.Status = status
	return &this
}

// NewCloudDatabaseStatusResponseWithDefaults instantiates a new CloudDatabaseStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudDatabaseStatusResponseWithDefaults() *CloudDatabaseStatusResponse {
	this := CloudDatabaseStatusResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *CloudDatabaseStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudDatabaseStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudDatabaseStatusResponse) SetStatus(v string) {
	o.Status = v
}

func (o CloudDatabaseStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudDatabaseStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *CloudDatabaseStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varCloudDatabaseStatusResponse := _CloudDatabaseStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudDatabaseStatusResponse)

	if err != nil {
		return err
	}

	*o = CloudDatabaseStatusResponse(varCloudDatabaseStatusResponse)

	return err
}

type NullableCloudDatabaseStatusResponse struct {
	value *CloudDatabaseStatusResponse
	isSet bool
}

func (v NullableCloudDatabaseStatusResponse) Get() *CloudDatabaseStatusResponse {
	return v.value
}

func (v *NullableCloudDatabaseStatusResponse) Set(val *CloudDatabaseStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudDatabaseStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudDatabaseStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudDatabaseStatusResponse(val *CloudDatabaseStatusResponse) *NullableCloudDatabaseStatusResponse {
	return &NullableCloudDatabaseStatusResponse{value: val, isSet: true}
}

func (v NullableCloudDatabaseStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudDatabaseStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

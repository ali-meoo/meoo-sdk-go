package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudFunctionLogsResponseEntriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunctionLogsResponseEntriesInner{}

// CloudFunctionLogsResponseEntriesInner struct for CloudFunctionLogsResponseEntriesInner
type CloudFunctionLogsResponseEntriesInner struct {
	Timestamp int64   `json:"timestamp"`
	Level     string  `json:"level"`
	Message   string  `json:"message"`
	Stack     *string `json:"stack,omitempty"`
}

type _CloudFunctionLogsResponseEntriesInner CloudFunctionLogsResponseEntriesInner

// NewCloudFunctionLogsResponseEntriesInner instantiates a new CloudFunctionLogsResponseEntriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunctionLogsResponseEntriesInner(timestamp int64, level string, message string) *CloudFunctionLogsResponseEntriesInner {
	this := CloudFunctionLogsResponseEntriesInner{}
	this.Timestamp = timestamp
	this.Level = level
	this.Message = message
	return &this
}

// NewCloudFunctionLogsResponseEntriesInnerWithDefaults instantiates a new CloudFunctionLogsResponseEntriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionLogsResponseEntriesInnerWithDefaults() *CloudFunctionLogsResponseEntriesInner {
	this := CloudFunctionLogsResponseEntriesInner{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *CloudFunctionLogsResponseEntriesInner) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponseEntriesInner) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CloudFunctionLogsResponseEntriesInner) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetLevel returns the Level field value
func (o *CloudFunctionLogsResponseEntriesInner) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponseEntriesInner) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CloudFunctionLogsResponseEntriesInner) SetLevel(v string) {
	o.Level = v
}

// GetMessage returns the Message field value
func (o *CloudFunctionLogsResponseEntriesInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponseEntriesInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CloudFunctionLogsResponseEntriesInner) SetMessage(v string) {
	o.Message = v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *CloudFunctionLogsResponseEntriesInner) GetStack() string {
	if o == nil || IsNil(o.Stack) {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudFunctionLogsResponseEntriesInner) GetStackOk() (*string, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *CloudFunctionLogsResponseEntriesInner) HasStack() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *CloudFunctionLogsResponseEntriesInner) SetStack(v string) {
	o.Stack = &v
}

func (o CloudFunctionLogsResponseEntriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunctionLogsResponseEntriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["level"] = o.Level
	toSerialize["message"] = o.Message
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	return toSerialize, nil
}

func (o *CloudFunctionLogsResponseEntriesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timestamp",
		"level",
		"message",
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

	varCloudFunctionLogsResponseEntriesInner := _CloudFunctionLogsResponseEntriesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudFunctionLogsResponseEntriesInner)

	if err != nil {
		return err
	}

	*o = CloudFunctionLogsResponseEntriesInner(varCloudFunctionLogsResponseEntriesInner)

	return err
}

type NullableCloudFunctionLogsResponseEntriesInner struct {
	value *CloudFunctionLogsResponseEntriesInner
	isSet bool
}

func (v NullableCloudFunctionLogsResponseEntriesInner) Get() *CloudFunctionLogsResponseEntriesInner {
	return v.value
}

func (v *NullableCloudFunctionLogsResponseEntriesInner) Set(val *CloudFunctionLogsResponseEntriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunctionLogsResponseEntriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunctionLogsResponseEntriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunctionLogsResponseEntriesInner(val *CloudFunctionLogsResponseEntriesInner) *NullableCloudFunctionLogsResponseEntriesInner {
	return &NullableCloudFunctionLogsResponseEntriesInner{value: val, isSet: true}
}

func (v NullableCloudFunctionLogsResponseEntriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunctionLogsResponseEntriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

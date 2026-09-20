package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunction{}

// CloudFunction struct for CloudFunction
type CloudFunction struct {
	Name      string        `json:"name"`
	Status    string        `json:"status"`
	Version   int32         `json:"version"`
	VerifyJwt bool          `json:"verify_jwt"`
	CreatedAt NullableInt64 `json:"created_at"`
	UpdatedAt NullableInt64 `json:"updated_at"`
}

type _CloudFunction CloudFunction

// NewCloudFunction instantiates a new CloudFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunction(name string, status string, version int32, verifyJwt bool, createdAt NullableInt64, updatedAt NullableInt64) *CloudFunction {
	this := CloudFunction{}
	this.Name = name
	this.Status = status
	this.Version = version
	this.VerifyJwt = verifyJwt
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCloudFunctionWithDefaults instantiates a new CloudFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionWithDefaults() *CloudFunction {
	this := CloudFunction{}
	return &this
}

// GetName returns the Name field value
func (o *CloudFunction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudFunction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudFunction) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *CloudFunction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudFunction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudFunction) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *CloudFunction) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CloudFunction) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CloudFunction) SetVersion(v int32) {
	o.Version = v
}

// GetVerifyJwt returns the VerifyJwt field value
func (o *CloudFunction) GetVerifyJwt() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VerifyJwt
}

// GetVerifyJwtOk returns a tuple with the VerifyJwt field value
// and a boolean to check if the value has been set.
func (o *CloudFunction) GetVerifyJwtOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyJwt, true
}

// SetVerifyJwt sets field value
func (o *CloudFunction) SetVerifyJwt(v bool) {
	o.VerifyJwt = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudFunction) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudFunction) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *CloudFunction) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CloudFunction) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudFunction) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *CloudFunction) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}

func (o CloudFunction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version
	toSerialize["verify_jwt"] = o.VerifyJwt
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	return toSerialize, nil
}

func (o *CloudFunction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"status",
		"version",
		"verify_jwt",
		"created_at",
		"updated_at",
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

	varCloudFunction := _CloudFunction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudFunction)

	if err != nil {
		return err
	}

	*o = CloudFunction(varCloudFunction)

	return err
}

type NullableCloudFunction struct {
	value *CloudFunction
	isSet bool
}

func (v NullableCloudFunction) Get() *CloudFunction {
	return v.value
}

func (v *NullableCloudFunction) Set(val *CloudFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunction(val *CloudFunction) *NullableCloudFunction {
	return &NullableCloudFunction{value: val, isSet: true}
}

func (v NullableCloudFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

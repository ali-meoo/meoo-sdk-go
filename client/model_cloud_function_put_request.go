package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CloudFunctionPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFunctionPutRequest{}

// CloudFunctionPutRequest struct for CloudFunctionPutRequest
type CloudFunctionPutRequest struct {
	// Base64 编码的 ZIP 文件，解码后最大 10 MiB。
	ArchiveBase64 string `json:"archive_base64"`
	VerifyJwt     *bool  `json:"verify_jwt,omitempty"`
}

type _CloudFunctionPutRequest CloudFunctionPutRequest

// NewCloudFunctionPutRequest instantiates a new CloudFunctionPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFunctionPutRequest(archiveBase64 string) *CloudFunctionPutRequest {
	this := CloudFunctionPutRequest{}
	this.ArchiveBase64 = archiveBase64
	var verifyJwt bool = true
	this.VerifyJwt = &verifyJwt
	return &this
}

// NewCloudFunctionPutRequestWithDefaults instantiates a new CloudFunctionPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFunctionPutRequestWithDefaults() *CloudFunctionPutRequest {
	this := CloudFunctionPutRequest{}
	var verifyJwt bool = true
	this.VerifyJwt = &verifyJwt
	return &this
}

// GetArchiveBase64 returns the ArchiveBase64 field value
func (o *CloudFunctionPutRequest) GetArchiveBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArchiveBase64
}

// GetArchiveBase64Ok returns a tuple with the ArchiveBase64 field value
// and a boolean to check if the value has been set.
func (o *CloudFunctionPutRequest) GetArchiveBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArchiveBase64, true
}

// SetArchiveBase64 sets field value
func (o *CloudFunctionPutRequest) SetArchiveBase64(v string) {
	o.ArchiveBase64 = v
}

// GetVerifyJwt returns the VerifyJwt field value if set, zero value otherwise.
func (o *CloudFunctionPutRequest) GetVerifyJwt() bool {
	if o == nil || IsNil(o.VerifyJwt) {
		var ret bool
		return ret
	}
	return *o.VerifyJwt
}

// GetVerifyJwtOk returns a tuple with the VerifyJwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudFunctionPutRequest) GetVerifyJwtOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyJwt) {
		return nil, false
	}
	return o.VerifyJwt, true
}

// HasVerifyJwt returns a boolean if a field has been set.
func (o *CloudFunctionPutRequest) HasVerifyJwt() bool {
	if o != nil && !IsNil(o.VerifyJwt) {
		return true
	}

	return false
}

// SetVerifyJwt gets a reference to the given bool and assigns it to the VerifyJwt field.
func (o *CloudFunctionPutRequest) SetVerifyJwt(v bool) {
	o.VerifyJwt = &v
}

func (o CloudFunctionPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFunctionPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archive_base64"] = o.ArchiveBase64
	if !IsNil(o.VerifyJwt) {
		toSerialize["verify_jwt"] = o.VerifyJwt
	}
	return toSerialize, nil
}

func (o *CloudFunctionPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archive_base64",
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

	varCloudFunctionPutRequest := _CloudFunctionPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudFunctionPutRequest)

	if err != nil {
		return err
	}

	*o = CloudFunctionPutRequest(varCloudFunctionPutRequest)

	return err
}

type NullableCloudFunctionPutRequest struct {
	value *CloudFunctionPutRequest
	isSet bool
}

func (v NullableCloudFunctionPutRequest) Get() *CloudFunctionPutRequest {
	return v.value
}

func (v *NullableCloudFunctionPutRequest) Set(val *CloudFunctionPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFunctionPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFunctionPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFunctionPutRequest(val *CloudFunctionPutRequest) *NullableCloudFunctionPutRequest {
	return &NullableCloudFunctionPutRequest{value: val, isSet: true}
}

func (v NullableCloudFunctionPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFunctionPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

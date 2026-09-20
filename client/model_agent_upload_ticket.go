package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentUploadTicket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentUploadTicket{}

// AgentUploadTicket struct for AgentUploadTicket
type AgentUploadTicket struct {
	// 本次直传票据的唯一标识，用于问题排查。
	UploadId string `json:"upload_id"`
	// 预签名 PUT 地址；PUT 请求须携带与签发时 type 一致的 Content-Type 头。
	UploadUrl string `json:"upload_url"`
	// 签名 GET 地址；可直接作为 startAgentRun attachments 的 url。
	FileUrl string `json:"file_url"`
	// upload_url（PUT）过期时间，Unix 毫秒。
	UploadExpiresAt int64 `json:"upload_expires_at"`
	// file_url（GET）过期时间，Unix 毫秒；过期后需重新上传。
	FileExpiresAt int64 `json:"file_expires_at"`
}

type _AgentUploadTicket AgentUploadTicket

// NewAgentUploadTicket instantiates a new AgentUploadTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUploadTicket(uploadId string, uploadUrl string, fileUrl string, uploadExpiresAt int64, fileExpiresAt int64) *AgentUploadTicket {
	this := AgentUploadTicket{}
	this.UploadId = uploadId
	this.UploadUrl = uploadUrl
	this.FileUrl = fileUrl
	this.UploadExpiresAt = uploadExpiresAt
	this.FileExpiresAt = fileExpiresAt
	return &this
}

// NewAgentUploadTicketWithDefaults instantiates a new AgentUploadTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUploadTicketWithDefaults() *AgentUploadTicket {
	this := AgentUploadTicket{}
	return &this
}

// GetUploadId returns the UploadId field value
func (o *AgentUploadTicket) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *AgentUploadTicket) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *AgentUploadTicket) SetUploadId(v string) {
	o.UploadId = v
}

// GetUploadUrl returns the UploadUrl field value
func (o *AgentUploadTicket) GetUploadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value
// and a boolean to check if the value has been set.
func (o *AgentUploadTicket) GetUploadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadUrl, true
}

// SetUploadUrl sets field value
func (o *AgentUploadTicket) SetUploadUrl(v string) {
	o.UploadUrl = v
}

// GetFileUrl returns the FileUrl field value
func (o *AgentUploadTicket) GetFileUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value
// and a boolean to check if the value has been set.
func (o *AgentUploadTicket) GetFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUrl, true
}

// SetFileUrl sets field value
func (o *AgentUploadTicket) SetFileUrl(v string) {
	o.FileUrl = v
}

// GetUploadExpiresAt returns the UploadExpiresAt field value
func (o *AgentUploadTicket) GetUploadExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UploadExpiresAt
}

// GetUploadExpiresAtOk returns a tuple with the UploadExpiresAt field value
// and a boolean to check if the value has been set.
func (o *AgentUploadTicket) GetUploadExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadExpiresAt, true
}

// SetUploadExpiresAt sets field value
func (o *AgentUploadTicket) SetUploadExpiresAt(v int64) {
	o.UploadExpiresAt = v
}

// GetFileExpiresAt returns the FileExpiresAt field value
func (o *AgentUploadTicket) GetFileExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileExpiresAt
}

// GetFileExpiresAtOk returns a tuple with the FileExpiresAt field value
// and a boolean to check if the value has been set.
func (o *AgentUploadTicket) GetFileExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileExpiresAt, true
}

// SetFileExpiresAt sets field value
func (o *AgentUploadTicket) SetFileExpiresAt(v int64) {
	o.FileExpiresAt = v
}

func (o AgentUploadTicket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentUploadTicket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upload_id"] = o.UploadId
	toSerialize["upload_url"] = o.UploadUrl
	toSerialize["file_url"] = o.FileUrl
	toSerialize["upload_expires_at"] = o.UploadExpiresAt
	toSerialize["file_expires_at"] = o.FileExpiresAt
	return toSerialize, nil
}

func (o *AgentUploadTicket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"upload_id",
		"upload_url",
		"file_url",
		"upload_expires_at",
		"file_expires_at",
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

	varAgentUploadTicket := _AgentUploadTicket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentUploadTicket)

	if err != nil {
		return err
	}

	*o = AgentUploadTicket(varAgentUploadTicket)

	return err
}

type NullableAgentUploadTicket struct {
	value *AgentUploadTicket
	isSet bool
}

func (v NullableAgentUploadTicket) Get() *AgentUploadTicket {
	return v.value
}

func (v *NullableAgentUploadTicket) Set(val *AgentUploadTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUploadTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUploadTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUploadTicket(val *AgentUploadTicket) *NullableAgentUploadTicket {
	return &NullableAgentUploadTicket{value: val, isSet: true}
}

func (v NullableAgentUploadTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUploadTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

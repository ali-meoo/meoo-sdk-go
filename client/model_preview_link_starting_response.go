package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PreviewLinkStartingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewLinkStartingResponse{}

// PreviewLinkStartingResponse struct for PreviewLinkStartingResponse
type PreviewLinkStartingResponse struct {
	Object    string `json:"object"`
	ProjectId string `json:"project_id"`
	Status    string `json:"status"`
	Code      string `json:"code"`
	// 建议重试等待时间，毫秒。
	RetryAfterMs int32 `json:"retry_after_ms"`
}

type _PreviewLinkStartingResponse PreviewLinkStartingResponse

// NewPreviewLinkStartingResponse instantiates a new PreviewLinkStartingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewLinkStartingResponse(object string, projectId string, status string, code string, retryAfterMs int32) *PreviewLinkStartingResponse {
	this := PreviewLinkStartingResponse{}
	this.Object = object
	this.ProjectId = projectId
	this.Status = status
	this.Code = code
	this.RetryAfterMs = retryAfterMs
	return &this
}

// NewPreviewLinkStartingResponseWithDefaults instantiates a new PreviewLinkStartingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewLinkStartingResponseWithDefaults() *PreviewLinkStartingResponse {
	this := PreviewLinkStartingResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PreviewLinkStartingResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkStartingResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PreviewLinkStartingResponse) SetObject(v string) {
	o.Object = v
}

// GetProjectId returns the ProjectId field value
func (o *PreviewLinkStartingResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkStartingResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PreviewLinkStartingResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStatus returns the Status field value
func (o *PreviewLinkStartingResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkStartingResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PreviewLinkStartingResponse) SetStatus(v string) {
	o.Status = v
}

// GetCode returns the Code field value
func (o *PreviewLinkStartingResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkStartingResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PreviewLinkStartingResponse) SetCode(v string) {
	o.Code = v
}

// GetRetryAfterMs returns the RetryAfterMs field value
func (o *PreviewLinkStartingResponse) GetRetryAfterMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetryAfterMs
}

// GetRetryAfterMsOk returns a tuple with the RetryAfterMs field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkStartingResponse) GetRetryAfterMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryAfterMs, true
}

// SetRetryAfterMs sets field value
func (o *PreviewLinkStartingResponse) SetRetryAfterMs(v int32) {
	o.RetryAfterMs = v
}

func (o PreviewLinkStartingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewLinkStartingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["project_id"] = o.ProjectId
	toSerialize["status"] = o.Status
	toSerialize["code"] = o.Code
	toSerialize["retry_after_ms"] = o.RetryAfterMs
	return toSerialize, nil
}

func (o *PreviewLinkStartingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"project_id",
		"status",
		"code",
		"retry_after_ms",
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

	varPreviewLinkStartingResponse := _PreviewLinkStartingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreviewLinkStartingResponse)

	if err != nil {
		return err
	}

	*o = PreviewLinkStartingResponse(varPreviewLinkStartingResponse)

	return err
}

type NullablePreviewLinkStartingResponse struct {
	value *PreviewLinkStartingResponse
	isSet bool
}

func (v NullablePreviewLinkStartingResponse) Get() *PreviewLinkStartingResponse {
	return v.value
}

func (v *NullablePreviewLinkStartingResponse) Set(val *PreviewLinkStartingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewLinkStartingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewLinkStartingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewLinkStartingResponse(val *PreviewLinkStartingResponse) *NullablePreviewLinkStartingResponse {
	return &NullablePreviewLinkStartingResponse{value: val, isSet: true}
}

func (v NullablePreviewLinkStartingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewLinkStartingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

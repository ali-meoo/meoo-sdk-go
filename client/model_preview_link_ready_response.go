package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PreviewLinkReadyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewLinkReadyResponse{}

// PreviewLinkReadyResponse struct for PreviewLinkReadyResponse
type PreviewLinkReadyResponse struct {
	Object    string `json:"object"`
	ProjectId string `json:"project_id"`
	Status    string `json:"status"`
	// 短时 opaque 预览壳 URL；不得解析、记录或长期存储。
	Url string `json:"url"`
	// URL 过期时间，Unix 毫秒。
	ExpiresAt int64 `json:"expires_at"`
}

type _PreviewLinkReadyResponse PreviewLinkReadyResponse

// NewPreviewLinkReadyResponse instantiates a new PreviewLinkReadyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewLinkReadyResponse(object string, projectId string, status string, url string, expiresAt int64) *PreviewLinkReadyResponse {
	this := PreviewLinkReadyResponse{}
	this.Object = object
	this.ProjectId = projectId
	this.Status = status
	this.Url = url
	this.ExpiresAt = expiresAt
	return &this
}

// NewPreviewLinkReadyResponseWithDefaults instantiates a new PreviewLinkReadyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewLinkReadyResponseWithDefaults() *PreviewLinkReadyResponse {
	this := PreviewLinkReadyResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PreviewLinkReadyResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkReadyResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PreviewLinkReadyResponse) SetObject(v string) {
	o.Object = v
}

// GetProjectId returns the ProjectId field value
func (o *PreviewLinkReadyResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkReadyResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PreviewLinkReadyResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStatus returns the Status field value
func (o *PreviewLinkReadyResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkReadyResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PreviewLinkReadyResponse) SetStatus(v string) {
	o.Status = v
}

// GetUrl returns the Url field value
func (o *PreviewLinkReadyResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkReadyResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PreviewLinkReadyResponse) SetUrl(v string) {
	o.Url = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *PreviewLinkReadyResponse) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PreviewLinkReadyResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *PreviewLinkReadyResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

func (o PreviewLinkReadyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewLinkReadyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["project_id"] = o.ProjectId
	toSerialize["status"] = o.Status
	toSerialize["url"] = o.Url
	toSerialize["expires_at"] = o.ExpiresAt
	return toSerialize, nil
}

func (o *PreviewLinkReadyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"project_id",
		"status",
		"url",
		"expires_at",
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

	varPreviewLinkReadyResponse := _PreviewLinkReadyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreviewLinkReadyResponse)

	if err != nil {
		return err
	}

	*o = PreviewLinkReadyResponse(varPreviewLinkReadyResponse)

	return err
}

type NullablePreviewLinkReadyResponse struct {
	value *PreviewLinkReadyResponse
	isSet bool
}

func (v NullablePreviewLinkReadyResponse) Get() *PreviewLinkReadyResponse {
	return v.value
}

func (v *NullablePreviewLinkReadyResponse) Set(val *PreviewLinkReadyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewLinkReadyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewLinkReadyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewLinkReadyResponse(val *PreviewLinkReadyResponse) *NullablePreviewLinkReadyResponse {
	return &NullablePreviewLinkReadyResponse{value: val, isSet: true}
}

func (v NullablePreviewLinkReadyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewLinkReadyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

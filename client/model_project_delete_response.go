package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDeleteResponse{}

// ProjectDeleteResponse struct for ProjectDeleteResponse
type ProjectDeleteResponse struct {
	// 已软删除项目的公开 URL ID。
	ProjectId string `json:"project_id"`
	// 项目已软删除；不代表异步云资源清理已完成。
	Status string `json:"status"`
}

type _ProjectDeleteResponse ProjectDeleteResponse

// NewProjectDeleteResponse instantiates a new ProjectDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDeleteResponse(projectId string, status string) *ProjectDeleteResponse {
	this := ProjectDeleteResponse{}
	this.ProjectId = projectId
	this.Status = status
	return &this
}

// NewProjectDeleteResponseWithDefaults instantiates a new ProjectDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDeleteResponseWithDefaults() *ProjectDeleteResponse {
	this := ProjectDeleteResponse{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ProjectDeleteResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectDeleteResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectDeleteResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStatus returns the Status field value
func (o *ProjectDeleteResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProjectDeleteResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProjectDeleteResponse) SetStatus(v string) {
	o.Status = v
}

func (o ProjectDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ProjectDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
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

	varProjectDeleteResponse := _ProjectDeleteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectDeleteResponse)

	if err != nil {
		return err
	}

	*o = ProjectDeleteResponse(varProjectDeleteResponse)

	return err
}

type NullableProjectDeleteResponse struct {
	value *ProjectDeleteResponse
	isSet bool
}

func (v NullableProjectDeleteResponse) Get() *ProjectDeleteResponse {
	return v.value
}

func (v *NullableProjectDeleteResponse) Set(val *ProjectDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDeleteResponse(val *ProjectDeleteResponse) *NullableProjectDeleteResponse {
	return &NullableProjectDeleteResponse{value: val, isSet: true}
}

func (v NullableProjectDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

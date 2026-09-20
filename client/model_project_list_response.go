package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListResponse{}

// ProjectListResponse struct for ProjectListResponse
type ProjectListResponse struct {
	Projects []Project `json:"projects"`
	// 仅在还有下一页时返回。
	NextPageToken *string `json:"next_page_token,omitempty"`
}

type _ProjectListResponse ProjectListResponse

// NewProjectListResponse instantiates a new ProjectListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListResponse(projects []Project) *ProjectListResponse {
	this := ProjectListResponse{}
	this.Projects = projects
	return &this
}

// NewProjectListResponseWithDefaults instantiates a new ProjectListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListResponseWithDefaults() *ProjectListResponse {
	this := ProjectListResponse{}
	return &this
}

// GetProjects returns the Projects field value
func (o *ProjectListResponse) GetProjects() []Project {
	if o == nil {
		var ret []Project
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetProjectsOk() ([]Project, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *ProjectListResponse) SetProjects(v []Project) {
	o.Projects = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ProjectListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ProjectListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ProjectListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ProjectListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projects"] = o.Projects
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return toSerialize, nil
}

func (o *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projects",
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

	varProjectListResponse := _ProjectListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectListResponse)

	if err != nil {
		return err
	}

	*o = ProjectListResponse(varProjectListResponse)

	return err
}

type NullableProjectListResponse struct {
	value *ProjectListResponse
	isSet bool
}

func (v NullableProjectListResponse) Get() *ProjectListResponse {
	return v.value
}

func (v *NullableProjectListResponse) Set(val *ProjectListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListResponse(val *ProjectListResponse) *NullableProjectListResponse {
	return &NullableProjectListResponse{value: val, isSet: true}
}

func (v NullableProjectListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

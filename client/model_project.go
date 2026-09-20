package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project struct for Project
type Project struct {
	// 项目的公开 URL ID，不是数据库 BIGINT 主键。
	UrlId string      `json:"url_id"`
	Name  string      `json:"name"`
	Type  ProjectType `json:"type"`
	// Unix 毫秒时间戳。
	CreatedAt int64 `json:"created_at"`
}

type _Project Project

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(urlId string, name string, type_ ProjectType, createdAt int64) *Project {
	this := Project{}
	this.UrlId = urlId
	this.Name = name
	this.Type = type_
	this.CreatedAt = createdAt
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetUrlId returns the UrlId field value
func (o *Project) GetUrlId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlId
}

// GetUrlIdOk returns a tuple with the UrlId field value
// and a boolean to check if the value has been set.
func (o *Project) GetUrlIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlId, true
}

// SetUrlId sets field value
func (o *Project) SetUrlId(v string) {
	o.UrlId = v
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Project) GetType() ProjectType {
	if o == nil {
		var ret ProjectType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Project) GetTypeOk() (*ProjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Project) SetType(v ProjectType) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Project) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Project) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url_id"] = o.UrlId
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *Project) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url_id",
		"name",
		"type",
		"created_at",
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

	varProject := _Project{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProject)

	if err != nil {
		return err
	}

	*o = Project(varProject)

	return err
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

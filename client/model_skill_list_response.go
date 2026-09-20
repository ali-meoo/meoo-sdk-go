package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SkillListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillListResponse{}

// SkillListResponse struct for SkillListResponse
type SkillListResponse struct {
	Skills        []SelectableSkill `json:"skills"`
	NextPageToken *string           `json:"next_page_token,omitempty"`
}

type _SkillListResponse SkillListResponse

// NewSkillListResponse instantiates a new SkillListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillListResponse(skills []SelectableSkill) *SkillListResponse {
	this := SkillListResponse{}
	this.Skills = skills
	return &this
}

// NewSkillListResponseWithDefaults instantiates a new SkillListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillListResponseWithDefaults() *SkillListResponse {
	this := SkillListResponse{}
	return &this
}

// GetSkills returns the Skills field value
func (o *SkillListResponse) GetSkills() []SelectableSkill {
	if o == nil {
		var ret []SelectableSkill
		return ret
	}

	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value
// and a boolean to check if the value has been set.
func (o *SkillListResponse) GetSkillsOk() ([]SelectableSkill, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skills, true
}

// SetSkills sets field value
func (o *SkillListResponse) SetSkills(v []SelectableSkill) {
	o.Skills = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *SkillListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *SkillListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *SkillListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o SkillListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["skills"] = o.Skills
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return toSerialize, nil
}

func (o *SkillListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"skills",
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

	varSkillListResponse := _SkillListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkillListResponse)

	if err != nil {
		return err
	}

	*o = SkillListResponse(varSkillListResponse)

	return err
}

type NullableSkillListResponse struct {
	value *SkillListResponse
	isSet bool
}

func (v NullableSkillListResponse) Get() *SkillListResponse {
	return v.value
}

func (v *NullableSkillListResponse) Set(val *SkillListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillListResponse(val *SkillListResponse) *NullableSkillListResponse {
	return &NullableSkillListResponse{value: val, isSet: true}
}

func (v NullableSkillListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

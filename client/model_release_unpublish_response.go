package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ReleaseUnpublishResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseUnpublishResponse{}

// ReleaseUnpublishResponse struct for ReleaseUnpublishResponse
type ReleaseUnpublishResponse struct {
	Status              string         `json:"status"`
	PreviousVersion     NullableString `json:"previous_version"`
	ShowcaseRemoved     bool           `json:"showcase_removed"`
	TemplateUnpublished bool           `json:"template_unpublished"`
	UnpublishedAt       int64          `json:"unpublished_at"`
}

type _ReleaseUnpublishResponse ReleaseUnpublishResponse

// NewReleaseUnpublishResponse instantiates a new ReleaseUnpublishResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseUnpublishResponse(status string, previousVersion NullableString, showcaseRemoved bool, templateUnpublished bool, unpublishedAt int64) *ReleaseUnpublishResponse {
	this := ReleaseUnpublishResponse{}
	this.Status = status
	this.PreviousVersion = previousVersion
	this.ShowcaseRemoved = showcaseRemoved
	this.TemplateUnpublished = templateUnpublished
	this.UnpublishedAt = unpublishedAt
	return &this
}

// NewReleaseUnpublishResponseWithDefaults instantiates a new ReleaseUnpublishResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseUnpublishResponseWithDefaults() *ReleaseUnpublishResponse {
	this := ReleaseUnpublishResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *ReleaseUnpublishResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReleaseUnpublishResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReleaseUnpublishResponse) SetStatus(v string) {
	o.Status = v
}

// GetPreviousVersion returns the PreviousVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReleaseUnpublishResponse) GetPreviousVersion() string {
	if o == nil || o.PreviousVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.PreviousVersion.Get()
}

// GetPreviousVersionOk returns a tuple with the PreviousVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseUnpublishResponse) GetPreviousVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousVersion.Get(), o.PreviousVersion.IsSet()
}

// SetPreviousVersion sets field value
func (o *ReleaseUnpublishResponse) SetPreviousVersion(v string) {
	o.PreviousVersion.Set(&v)
}

// GetShowcaseRemoved returns the ShowcaseRemoved field value
func (o *ReleaseUnpublishResponse) GetShowcaseRemoved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowcaseRemoved
}

// GetShowcaseRemovedOk returns a tuple with the ShowcaseRemoved field value
// and a boolean to check if the value has been set.
func (o *ReleaseUnpublishResponse) GetShowcaseRemovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowcaseRemoved, true
}

// SetShowcaseRemoved sets field value
func (o *ReleaseUnpublishResponse) SetShowcaseRemoved(v bool) {
	o.ShowcaseRemoved = v
}

// GetTemplateUnpublished returns the TemplateUnpublished field value
func (o *ReleaseUnpublishResponse) GetTemplateUnpublished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TemplateUnpublished
}

// GetTemplateUnpublishedOk returns a tuple with the TemplateUnpublished field value
// and a boolean to check if the value has been set.
func (o *ReleaseUnpublishResponse) GetTemplateUnpublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateUnpublished, true
}

// SetTemplateUnpublished sets field value
func (o *ReleaseUnpublishResponse) SetTemplateUnpublished(v bool) {
	o.TemplateUnpublished = v
}

// GetUnpublishedAt returns the UnpublishedAt field value
func (o *ReleaseUnpublishResponse) GetUnpublishedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnpublishedAt
}

// GetUnpublishedAtOk returns a tuple with the UnpublishedAt field value
// and a boolean to check if the value has been set.
func (o *ReleaseUnpublishResponse) GetUnpublishedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnpublishedAt, true
}

// SetUnpublishedAt sets field value
func (o *ReleaseUnpublishResponse) SetUnpublishedAt(v int64) {
	o.UnpublishedAt = v
}

func (o ReleaseUnpublishResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseUnpublishResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["previous_version"] = o.PreviousVersion.Get()
	toSerialize["showcase_removed"] = o.ShowcaseRemoved
	toSerialize["template_unpublished"] = o.TemplateUnpublished
	toSerialize["unpublished_at"] = o.UnpublishedAt
	return toSerialize, nil
}

func (o *ReleaseUnpublishResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"previous_version",
		"showcase_removed",
		"template_unpublished",
		"unpublished_at",
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

	varReleaseUnpublishResponse := _ReleaseUnpublishResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReleaseUnpublishResponse)

	if err != nil {
		return err
	}

	*o = ReleaseUnpublishResponse(varReleaseUnpublishResponse)

	return err
}

type NullableReleaseUnpublishResponse struct {
	value *ReleaseUnpublishResponse
	isSet bool
}

func (v NullableReleaseUnpublishResponse) Get() *ReleaseUnpublishResponse {
	return v.value
}

func (v *NullableReleaseUnpublishResponse) Set(val *ReleaseUnpublishResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseUnpublishResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseUnpublishResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseUnpublishResponse(val *ReleaseUnpublishResponse) *NullableReleaseUnpublishResponse {
	return &NullableReleaseUnpublishResponse{value: val, isSet: true}
}

func (v NullableReleaseUnpublishResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseUnpublishResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

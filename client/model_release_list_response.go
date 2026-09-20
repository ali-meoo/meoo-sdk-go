package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ReleaseListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseListResponse{}

// ReleaseListResponse struct for ReleaseListResponse
type ReleaseListResponse struct {
	Items []ReleaseListResponseItemsInner `json:"items"`
	// 仅在还有下一页时返回。
	NextPageToken        *string `json:"next_page_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReleaseListResponse ReleaseListResponse

// NewReleaseListResponse instantiates a new ReleaseListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseListResponse(items []ReleaseListResponseItemsInner) *ReleaseListResponse {
	this := ReleaseListResponse{}
	this.Items = items
	return &this
}

// NewReleaseListResponseWithDefaults instantiates a new ReleaseListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseListResponseWithDefaults() *ReleaseListResponse {
	this := ReleaseListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ReleaseListResponse) GetItems() []ReleaseListResponseItemsInner {
	if o == nil {
		var ret []ReleaseListResponseItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ReleaseListResponse) GetItemsOk() ([]ReleaseListResponseItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ReleaseListResponse) SetItems(v []ReleaseListResponseItemsInner) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ReleaseListResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ReleaseListResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ReleaseListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ReleaseListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReleaseListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varReleaseListResponse := _ReleaseListResponse{}

	err = json.Unmarshal(data, &varReleaseListResponse)

	if err != nil {
		return err
	}

	*o = ReleaseListResponse(varReleaseListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "next_page_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReleaseListResponse struct {
	value *ReleaseListResponse
	isSet bool
}

func (v NullableReleaseListResponse) Get() *ReleaseListResponse {
	return v.value
}

func (v *NullableReleaseListResponse) Set(val *ReleaseListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseListResponse(val *ReleaseListResponse) *NullableReleaseListResponse {
	return &NullableReleaseListResponse{value: val, isSet: true}
}

func (v NullableReleaseListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

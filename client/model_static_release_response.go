package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StaticReleaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleaseResponse{}

// StaticReleaseResponse struct for StaticReleaseResponse
type StaticReleaseResponse struct {
	ReleaseId    string         `json:"release_id"`
	Version      int32          `json:"version"`
	ArtifactId   string         `json:"artifact_id"`
	Runtime      string         `json:"runtime"`
	ArtifactType string         `json:"artifact_type"`
	Status       string         `json:"status"`
	AccessUrl    NullableString `json:"access_url"`
	PublishedAt  int64          `json:"published_at"`
	ExpiresAt    NullableInt64  `json:"expires_at"`
}

type _StaticReleaseResponse StaticReleaseResponse

// NewStaticReleaseResponse instantiates a new StaticReleaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleaseResponse(releaseId string, version int32, artifactId string, runtime string, artifactType string, status string, accessUrl NullableString, publishedAt int64, expiresAt NullableInt64) *StaticReleaseResponse {
	this := StaticReleaseResponse{}
	this.ReleaseId = releaseId
	this.Version = version
	this.ArtifactId = artifactId
	this.Runtime = runtime
	this.ArtifactType = artifactType
	this.Status = status
	this.AccessUrl = accessUrl
	this.PublishedAt = publishedAt
	this.ExpiresAt = expiresAt
	return &this
}

// NewStaticReleaseResponseWithDefaults instantiates a new StaticReleaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleaseResponseWithDefaults() *StaticReleaseResponse {
	this := StaticReleaseResponse{}
	return &this
}

// GetReleaseId returns the ReleaseId field value
func (o *StaticReleaseResponse) GetReleaseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetReleaseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseId, true
}

// SetReleaseId sets field value
func (o *StaticReleaseResponse) SetReleaseId(v string) {
	o.ReleaseId = v
}

// GetVersion returns the Version field value
func (o *StaticReleaseResponse) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *StaticReleaseResponse) SetVersion(v int32) {
	o.Version = v
}

// GetArtifactId returns the ArtifactId field value
func (o *StaticReleaseResponse) GetArtifactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetArtifactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactId, true
}

// SetArtifactId sets field value
func (o *StaticReleaseResponse) SetArtifactId(v string) {
	o.ArtifactId = v
}

// GetRuntime returns the Runtime field value
func (o *StaticReleaseResponse) GetRuntime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetRuntimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *StaticReleaseResponse) SetRuntime(v string) {
	o.Runtime = v
}

// GetArtifactType returns the ArtifactType field value
func (o *StaticReleaseResponse) GetArtifactType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetArtifactTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *StaticReleaseResponse) SetArtifactType(v string) {
	o.ArtifactType = v
}

// GetStatus returns the Status field value
func (o *StaticReleaseResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StaticReleaseResponse) SetStatus(v string) {
	o.Status = v
}

// GetAccessUrl returns the AccessUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StaticReleaseResponse) GetAccessUrl() string {
	if o == nil || o.AccessUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessUrl.Get()
}

// GetAccessUrlOk returns a tuple with the AccessUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticReleaseResponse) GetAccessUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessUrl.Get(), o.AccessUrl.IsSet()
}

// SetAccessUrl sets field value
func (o *StaticReleaseResponse) SetAccessUrl(v string) {
	o.AccessUrl.Set(&v)
}

// GetPublishedAt returns the PublishedAt field value
func (o *StaticReleaseResponse) GetPublishedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value
// and a boolean to check if the value has been set.
func (o *StaticReleaseResponse) GetPublishedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishedAt, true
}

// SetPublishedAt sets field value
func (o *StaticReleaseResponse) SetPublishedAt(v int64) {
	o.PublishedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *StaticReleaseResponse) GetExpiresAt() int64 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticReleaseResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *StaticReleaseResponse) SetExpiresAt(v int64) {
	o.ExpiresAt.Set(&v)
}

func (o StaticReleaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["release_id"] = o.ReleaseId
	toSerialize["version"] = o.Version
	toSerialize["artifact_id"] = o.ArtifactId
	toSerialize["runtime"] = o.Runtime
	toSerialize["artifact_type"] = o.ArtifactType
	toSerialize["status"] = o.Status
	toSerialize["access_url"] = o.AccessUrl.Get()
	toSerialize["published_at"] = o.PublishedAt
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	return toSerialize, nil
}

func (o *StaticReleaseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"release_id",
		"version",
		"artifact_id",
		"runtime",
		"artifact_type",
		"status",
		"access_url",
		"published_at",
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

	varStaticReleaseResponse := _StaticReleaseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticReleaseResponse)

	if err != nil {
		return err
	}

	*o = StaticReleaseResponse(varStaticReleaseResponse)

	return err
}

type NullableStaticReleaseResponse struct {
	value *StaticReleaseResponse
	isSet bool
}

func (v NullableStaticReleaseResponse) Get() *StaticReleaseResponse {
	return v.value
}

func (v *NullableStaticReleaseResponse) Set(val *StaticReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleaseResponse(val *StaticReleaseResponse) *NullableStaticReleaseResponse {
	return &NullableStaticReleaseResponse{value: val, isSet: true}
}

func (v NullableStaticReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the StaticReleasePrepareResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticReleasePrepareResponse{}

// StaticReleasePrepareResponse struct for StaticReleasePrepareResponse
type StaticReleasePrepareResponse struct {
	// complete 所需的不透明短期 Token
	ReleaseToken         string                             `json:"release_token"`
	ArtifactId           string                             `json:"artifact_id"`
	Runtime              string                             `json:"runtime"`
	ArtifactType         string                             `json:"artifact_type"`
	Status               string                             `json:"status"`
	Upload               StaticReleasePrepareResponseUpload `json:"upload"`
	AdditionalProperties map[string]interface{}
}

type _StaticReleasePrepareResponse StaticReleasePrepareResponse

// NewStaticReleasePrepareResponse instantiates a new StaticReleasePrepareResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticReleasePrepareResponse(releaseToken string, artifactId string, runtime string, artifactType string, status string, upload StaticReleasePrepareResponseUpload) *StaticReleasePrepareResponse {
	this := StaticReleasePrepareResponse{}
	this.ReleaseToken = releaseToken
	this.ArtifactId = artifactId
	this.Runtime = runtime
	this.ArtifactType = artifactType
	this.Status = status
	this.Upload = upload
	return &this
}

// NewStaticReleasePrepareResponseWithDefaults instantiates a new StaticReleasePrepareResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticReleasePrepareResponseWithDefaults() *StaticReleasePrepareResponse {
	this := StaticReleasePrepareResponse{}
	return &this
}

// GetReleaseToken returns the ReleaseToken field value
func (o *StaticReleasePrepareResponse) GetReleaseToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseToken
}

// GetReleaseTokenOk returns a tuple with the ReleaseToken field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponse) GetReleaseTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseToken, true
}

// SetReleaseToken sets field value
func (o *StaticReleasePrepareResponse) SetReleaseToken(v string) {
	o.ReleaseToken = v
}

// GetArtifactId returns the ArtifactId field value
func (o *StaticReleasePrepareResponse) GetArtifactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponse) GetArtifactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactId, true
}

// SetArtifactId sets field value
func (o *StaticReleasePrepareResponse) SetArtifactId(v string) {
	o.ArtifactId = v
}

// GetRuntime returns the Runtime field value
func (o *StaticReleasePrepareResponse) GetRuntime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponse) GetRuntimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *StaticReleasePrepareResponse) SetRuntime(v string) {
	o.Runtime = v
}

// GetArtifactType returns the ArtifactType field value
func (o *StaticReleasePrepareResponse) GetArtifactType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponse) GetArtifactTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *StaticReleasePrepareResponse) SetArtifactType(v string) {
	o.ArtifactType = v
}

// GetStatus returns the Status field value
func (o *StaticReleasePrepareResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StaticReleasePrepareResponse) SetStatus(v string) {
	o.Status = v
}

// GetUpload returns the Upload field value
func (o *StaticReleasePrepareResponse) GetUpload() StaticReleasePrepareResponseUpload {
	if o == nil {
		var ret StaticReleasePrepareResponseUpload
		return ret
	}

	return o.Upload
}

// GetUploadOk returns a tuple with the Upload field value
// and a boolean to check if the value has been set.
func (o *StaticReleasePrepareResponse) GetUploadOk() (*StaticReleasePrepareResponseUpload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upload, true
}

// SetUpload sets field value
func (o *StaticReleasePrepareResponse) SetUpload(v StaticReleasePrepareResponseUpload) {
	o.Upload = v
}

func (o StaticReleasePrepareResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticReleasePrepareResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["release_token"] = o.ReleaseToken
	toSerialize["artifact_id"] = o.ArtifactId
	toSerialize["runtime"] = o.Runtime
	toSerialize["artifact_type"] = o.ArtifactType
	toSerialize["status"] = o.Status
	toSerialize["upload"] = o.Upload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StaticReleasePrepareResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"release_token",
		"artifact_id",
		"runtime",
		"artifact_type",
		"status",
		"upload",
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

	varStaticReleasePrepareResponse := _StaticReleasePrepareResponse{}

	err = json.Unmarshal(data, &varStaticReleasePrepareResponse)

	if err != nil {
		return err
	}

	*o = StaticReleasePrepareResponse(varStaticReleasePrepareResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "release_token")
		delete(additionalProperties, "artifact_id")
		delete(additionalProperties, "runtime")
		delete(additionalProperties, "artifact_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "upload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStaticReleasePrepareResponse struct {
	value *StaticReleasePrepareResponse
	isSet bool
}

func (v NullableStaticReleasePrepareResponse) Get() *StaticReleasePrepareResponse {
	return v.value
}

func (v *NullableStaticReleasePrepareResponse) Set(val *StaticReleasePrepareResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticReleasePrepareResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticReleasePrepareResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticReleasePrepareResponse(val *StaticReleasePrepareResponse) *NullableStaticReleasePrepareResponse {
	return &NullableStaticReleasePrepareResponse{value: val, isSet: true}
}

func (v NullableStaticReleasePrepareResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticReleasePrepareResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

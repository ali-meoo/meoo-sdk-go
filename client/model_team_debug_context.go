package client

import (
	"encoding/json"
	"fmt"
)

// checks if the TeamDebugContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamDebugContext{}

// TeamDebugContext struct for TeamDebugContext
type TeamDebugContext struct {
	Team                 TeamDebugContextTeam                    `json:"team"`
	CredentialId         string                                  `json:"credential_id"`
	ManagementScopes     []TeamDebugContextManagementScopesInner `json:"management_scopes"`
	DelegableScopes      []Items                                 `json:"delegable_scopes"`
	TokenTtlMaxSeconds   int32                                   `json:"token_ttl_max_seconds"`
	AdditionalProperties map[string]interface{}
}

type _TeamDebugContext TeamDebugContext

// NewTeamDebugContext instantiates a new TeamDebugContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamDebugContext(team TeamDebugContextTeam, credentialId string, managementScopes []TeamDebugContextManagementScopesInner, delegableScopes []Items, tokenTtlMaxSeconds int32) *TeamDebugContext {
	this := TeamDebugContext{}
	this.Team = team
	this.CredentialId = credentialId
	this.ManagementScopes = managementScopes
	this.DelegableScopes = delegableScopes
	this.TokenTtlMaxSeconds = tokenTtlMaxSeconds
	return &this
}

// NewTeamDebugContextWithDefaults instantiates a new TeamDebugContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamDebugContextWithDefaults() *TeamDebugContext {
	this := TeamDebugContext{}
	return &this
}

// GetTeam returns the Team field value
func (o *TeamDebugContext) GetTeam() TeamDebugContextTeam {
	if o == nil {
		var ret TeamDebugContextTeam
		return ret
	}

	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContext) GetTeamOk() (*TeamDebugContextTeam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value
func (o *TeamDebugContext) SetTeam(v TeamDebugContextTeam) {
	o.Team = v
}

// GetCredentialId returns the CredentialId field value
func (o *TeamDebugContext) GetCredentialId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContext) GetCredentialIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialId, true
}

// SetCredentialId sets field value
func (o *TeamDebugContext) SetCredentialId(v string) {
	o.CredentialId = v
}

// GetManagementScopes returns the ManagementScopes field value
func (o *TeamDebugContext) GetManagementScopes() []TeamDebugContextManagementScopesInner {
	if o == nil {
		var ret []TeamDebugContextManagementScopesInner
		return ret
	}

	return o.ManagementScopes
}

// GetManagementScopesOk returns a tuple with the ManagementScopes field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContext) GetManagementScopesOk() ([]TeamDebugContextManagementScopesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementScopes, true
}

// SetManagementScopes sets field value
func (o *TeamDebugContext) SetManagementScopes(v []TeamDebugContextManagementScopesInner) {
	o.ManagementScopes = v
}

// GetDelegableScopes returns the DelegableScopes field value
func (o *TeamDebugContext) GetDelegableScopes() []Items {
	if o == nil {
		var ret []Items
		return ret
	}

	return o.DelegableScopes
}

// GetDelegableScopesOk returns a tuple with the DelegableScopes field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContext) GetDelegableScopesOk() ([]Items, bool) {
	if o == nil {
		return nil, false
	}
	return o.DelegableScopes, true
}

// SetDelegableScopes sets field value
func (o *TeamDebugContext) SetDelegableScopes(v []Items) {
	o.DelegableScopes = v
}

// GetTokenTtlMaxSeconds returns the TokenTtlMaxSeconds field value
func (o *TeamDebugContext) GetTokenTtlMaxSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenTtlMaxSeconds
}

// GetTokenTtlMaxSecondsOk returns a tuple with the TokenTtlMaxSeconds field value
// and a boolean to check if the value has been set.
func (o *TeamDebugContext) GetTokenTtlMaxSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTtlMaxSeconds, true
}

// SetTokenTtlMaxSeconds sets field value
func (o *TeamDebugContext) SetTokenTtlMaxSeconds(v int32) {
	o.TokenTtlMaxSeconds = v
}

func (o TeamDebugContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamDebugContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["team"] = o.Team
	toSerialize["credential_id"] = o.CredentialId
	toSerialize["management_scopes"] = o.ManagementScopes
	toSerialize["delegable_scopes"] = o.DelegableScopes
	toSerialize["token_ttl_max_seconds"] = o.TokenTtlMaxSeconds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TeamDebugContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"team",
		"credential_id",
		"management_scopes",
		"delegable_scopes",
		"token_ttl_max_seconds",
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

	varTeamDebugContext := _TeamDebugContext{}

	err = json.Unmarshal(data, &varTeamDebugContext)

	if err != nil {
		return err
	}

	*o = TeamDebugContext(varTeamDebugContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "team")
		delete(additionalProperties, "credential_id")
		delete(additionalProperties, "management_scopes")
		delete(additionalProperties, "delegable_scopes")
		delete(additionalProperties, "token_ttl_max_seconds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTeamDebugContext struct {
	value *TeamDebugContext
	isSet bool
}

func (v NullableTeamDebugContext) Get() *TeamDebugContext {
	return v.value
}

func (v *NullableTeamDebugContext) Set(val *TeamDebugContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamDebugContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamDebugContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamDebugContext(val *TeamDebugContext) *NullableTeamDebugContext {
	return &NullableTeamDebugContext{value: val, isSet: true}
}

func (v NullableTeamDebugContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamDebugContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

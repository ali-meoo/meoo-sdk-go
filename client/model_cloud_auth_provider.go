package client

import (
	"encoding/json"
	"fmt"
)

// CloudAuthProvider the model 'CloudAuthProvider'
type CloudAuthProvider string

// List of CloudAuthProvider
const (
	CLOUDAUTHPROVIDER_PASSWORD CloudAuthProvider = "password"
	CLOUDAUTHPROVIDER_EMAIL    CloudAuthProvider = "email"
	CLOUDAUTHPROVIDER_SMS      CloudAuthProvider = "sms"
)

// All allowed values of CloudAuthProvider enum
var AllowedCloudAuthProviderEnumValues = []CloudAuthProvider{
	"password",
	"email",
	"sms",
}

func (v *CloudAuthProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudAuthProvider(value)
	for _, existing := range AllowedCloudAuthProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	// 前向兼容：未知枚举值原样保留而非报错（契约要求未知取值按非终态/进行中兜底）。
	*v = enumTypeValue
	return nil
}

// NewCloudAuthProviderFromValue returns a pointer to a valid CloudAuthProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudAuthProviderFromValue(v string) (*CloudAuthProvider, error) {
	ev := CloudAuthProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudAuthProvider: valid values are %v", v, AllowedCloudAuthProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudAuthProvider) IsValid() bool {
	for _, existing := range AllowedCloudAuthProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudAuthProvider value
func (v CloudAuthProvider) Ptr() *CloudAuthProvider {
	return &v
}

type NullableCloudAuthProvider struct {
	value *CloudAuthProvider
	isSet bool
}

func (v NullableCloudAuthProvider) Get() *CloudAuthProvider {
	return v.value
}

func (v *NullableCloudAuthProvider) Set(val *CloudAuthProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAuthProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAuthProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAuthProvider(val *CloudAuthProvider) *NullableCloudAuthProvider {
	return &NullableCloudAuthProvider{value: val, isSet: true}
}

func (v NullableCloudAuthProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAuthProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

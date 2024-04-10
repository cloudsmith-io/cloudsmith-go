/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.392.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryGeoIpCountryCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGeoIpCountryCode{}

// RepositoryGeoIpCountryCode struct for RepositoryGeoIpCountryCode
type RepositoryGeoIpCountryCode struct {
	// The allowed country codes for this repository
	Allow []string `json:"allow"`
	// The denied country codes for this repository
	Deny []string `json:"deny"`
}

// NewRepositoryGeoIpCountryCode instantiates a new RepositoryGeoIpCountryCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpCountryCode(allow []string, deny []string) *RepositoryGeoIpCountryCode {
	this := RepositoryGeoIpCountryCode{}
	this.Allow = allow
	this.Deny = deny
	return &this
}

// NewRepositoryGeoIpCountryCodeWithDefaults instantiates a new RepositoryGeoIpCountryCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpCountryCodeWithDefaults() *RepositoryGeoIpCountryCode {
	this := RepositoryGeoIpCountryCode{}
	return &this
}

// GetAllow returns the Allow field value
func (o *RepositoryGeoIpCountryCode) GetAllow() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpCountryCode) GetAllowOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Allow, true
}

// SetAllow sets field value
func (o *RepositoryGeoIpCountryCode) SetAllow(v []string) {
	o.Allow = v
}

// GetDeny returns the Deny field value
func (o *RepositoryGeoIpCountryCode) GetDeny() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpCountryCode) GetDenyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deny, true
}

// SetDeny sets field value
func (o *RepositoryGeoIpCountryCode) SetDeny(v []string) {
	o.Deny = v
}

func (o RepositoryGeoIpCountryCode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGeoIpCountryCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allow"] = o.Allow
	toSerialize["deny"] = o.Deny
	return toSerialize, nil
}

type NullableRepositoryGeoIpCountryCode struct {
	value *RepositoryGeoIpCountryCode
	isSet bool
}

func (v NullableRepositoryGeoIpCountryCode) Get() *RepositoryGeoIpCountryCode {
	return v.value
}

func (v *NullableRepositoryGeoIpCountryCode) Set(val *RepositoryGeoIpCountryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpCountryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpCountryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpCountryCode(val *RepositoryGeoIpCountryCode) *NullableRepositoryGeoIpCountryCode {
	return &NullableRepositoryGeoIpCountryCode{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpCountryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpCountryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

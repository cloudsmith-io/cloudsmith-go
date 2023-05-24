/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.273.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIpCidr struct for RepositoryGeoIpCidr
type RepositoryGeoIpCidr struct {
	// The allowed CIDRs for this repository
	Allow []string `json:"allow"`
	// The denied CIDRs for this repository
	Deny []string `json:"deny"`
}

// NewRepositoryGeoIpCidr instantiates a new RepositoryGeoIpCidr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpCidr(allow []string, deny []string) *RepositoryGeoIpCidr {
	this := RepositoryGeoIpCidr{}
	this.Allow = allow
	this.Deny = deny
	return &this
}

// NewRepositoryGeoIpCidrWithDefaults instantiates a new RepositoryGeoIpCidr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpCidrWithDefaults() *RepositoryGeoIpCidr {
	this := RepositoryGeoIpCidr{}
	return &this
}

// GetAllow returns the Allow field value
func (o *RepositoryGeoIpCidr) GetAllow() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpCidr) GetAllowOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Allow, true
}

// SetAllow sets field value
func (o *RepositoryGeoIpCidr) SetAllow(v []string) {
	o.Allow = v
}

// GetDeny returns the Deny field value
func (o *RepositoryGeoIpCidr) GetDeny() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpCidr) GetDenyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deny, true
}

// SetDeny sets field value
func (o *RepositoryGeoIpCidr) SetDeny(v []string) {
	o.Deny = v
}

func (o RepositoryGeoIpCidr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allow"] = o.Allow
	}
	if true {
		toSerialize["deny"] = o.Deny
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIpCidr struct {
	value *RepositoryGeoIpCidr
	isSet bool
}

func (v NullableRepositoryGeoIpCidr) Get() *RepositoryGeoIpCidr {
	return v.value
}

func (v *NullableRepositoryGeoIpCidr) Set(val *RepositoryGeoIpCidr) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpCidr) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpCidr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpCidr(val *RepositoryGeoIpCidr) *NullableRepositoryGeoIpCidr {
	return &NullableRepositoryGeoIpCidr{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpCidr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpCidr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

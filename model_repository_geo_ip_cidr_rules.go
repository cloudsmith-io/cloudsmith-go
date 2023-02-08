/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.206.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIpCidrRules struct for RepositoryGeoIpCidrRules
type RepositoryGeoIpCidrRules struct {
	Allow []string `json:"allow,omitempty"`
	Deny  []string `json:"deny,omitempty"`
}

// NewRepositoryGeoIpCidrRules instantiates a new RepositoryGeoIpCidrRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpCidrRules() *RepositoryGeoIpCidrRules {
	this := RepositoryGeoIpCidrRules{}
	return &this
}

// NewRepositoryGeoIpCidrRulesWithDefaults instantiates a new RepositoryGeoIpCidrRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpCidrRulesWithDefaults() *RepositoryGeoIpCidrRules {
	this := RepositoryGeoIpCidrRules{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *RepositoryGeoIpCidrRules) GetAllow() []string {
	if o == nil || isNil(o.Allow) {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpCidrRules) GetAllowOk() ([]string, bool) {
	if o == nil || isNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *RepositoryGeoIpCidrRules) HasAllow() bool {
	if o != nil && !isNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *RepositoryGeoIpCidrRules) SetAllow(v []string) {
	o.Allow = v
}

// GetDeny returns the Deny field value if set, zero value otherwise.
func (o *RepositoryGeoIpCidrRules) GetDeny() []string {
	if o == nil || isNil(o.Deny) {
		var ret []string
		return ret
	}
	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpCidrRules) GetDenyOk() ([]string, bool) {
	if o == nil || isNil(o.Deny) {
		return nil, false
	}
	return o.Deny, true
}

// HasDeny returns a boolean if a field has been set.
func (o *RepositoryGeoIpCidrRules) HasDeny() bool {
	if o != nil && !isNil(o.Deny) {
		return true
	}

	return false
}

// SetDeny gets a reference to the given []string and assigns it to the Deny field.
func (o *RepositoryGeoIpCidrRules) SetDeny(v []string) {
	o.Deny = v
}

func (o RepositoryGeoIpCidrRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Allow) {
		toSerialize["allow"] = o.Allow
	}
	if !isNil(o.Deny) {
		toSerialize["deny"] = o.Deny
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIpCidrRules struct {
	value *RepositoryGeoIpCidrRules
	isSet bool
}

func (v NullableRepositoryGeoIpCidrRules) Get() *RepositoryGeoIpCidrRules {
	return v.value
}

func (v *NullableRepositoryGeoIpCidrRules) Set(val *RepositoryGeoIpCidrRules) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpCidrRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpCidrRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpCidrRules(val *RepositoryGeoIpCidrRules) *NullableRepositoryGeoIpCidrRules {
	return &NullableRepositoryGeoIpCidrRules{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpCidrRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpCidrRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

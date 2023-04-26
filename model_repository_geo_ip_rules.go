/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.249.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIPRules struct for RepositoryGeoIPRules
type RepositoryGeoIPRules struct {
	Cidr        RepositoryGeoIPCidr        `json:"cidr"`
	CountryCode RepositoryGeoIPCountryCode `json:"country_code"`
}

// NewRepositoryGeoIPRules instantiates a new RepositoryGeoIPRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIPRules(cidr RepositoryGeoIPCidr, countryCode RepositoryGeoIPCountryCode) *RepositoryGeoIPRules {
	this := RepositoryGeoIPRules{}
	this.Cidr = cidr
	this.CountryCode = countryCode
	return &this
}

// NewRepositoryGeoIPRulesWithDefaults instantiates a new RepositoryGeoIPRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIPRulesWithDefaults() *RepositoryGeoIPRules {
	this := RepositoryGeoIPRules{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *RepositoryGeoIPRules) GetCidr() RepositoryGeoIPCidr {
	if o == nil {
		var ret RepositoryGeoIPCidr
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPRules) GetCidrOk() (*RepositoryGeoIPCidr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *RepositoryGeoIPRules) SetCidr(v RepositoryGeoIPCidr) {
	o.Cidr = v
}

// GetCountryCode returns the CountryCode field value
func (o *RepositoryGeoIPRules) GetCountryCode() RepositoryGeoIPCountryCode {
	if o == nil {
		var ret RepositoryGeoIPCountryCode
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPRules) GetCountryCodeOk() (*RepositoryGeoIPCountryCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *RepositoryGeoIPRules) SetCountryCode(v RepositoryGeoIPCountryCode) {
	o.CountryCode = v
}

func (o RepositoryGeoIPRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIPRules struct {
	value *RepositoryGeoIPRules
	isSet bool
}

func (v NullableRepositoryGeoIPRules) Get() *RepositoryGeoIPRules {
	return v.value
}

func (v *NullableRepositoryGeoIPRules) Set(val *RepositoryGeoIPRules) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIPRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIPRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIPRules(val *RepositoryGeoIPRules) *NullableRepositoryGeoIPRules {
	return &NullableRepositoryGeoIPRules{value: val, isSet: true}
}

func (v NullableRepositoryGeoIPRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIPRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

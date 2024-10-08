/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.533.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryGeoIpRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGeoIpRules{}

// RepositoryGeoIpRules struct for RepositoryGeoIpRules
type RepositoryGeoIpRules struct {
	Cidr        RepositoryGeoIpCidr        `json:"cidr"`
	CountryCode RepositoryGeoIpCountryCode `json:"country_code"`
}

// NewRepositoryGeoIpRules instantiates a new RepositoryGeoIpRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpRules(cidr RepositoryGeoIpCidr, countryCode RepositoryGeoIpCountryCode) *RepositoryGeoIpRules {
	this := RepositoryGeoIpRules{}
	this.Cidr = cidr
	this.CountryCode = countryCode
	return &this
}

// NewRepositoryGeoIpRulesWithDefaults instantiates a new RepositoryGeoIpRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpRulesWithDefaults() *RepositoryGeoIpRules {
	this := RepositoryGeoIpRules{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *RepositoryGeoIpRules) GetCidr() RepositoryGeoIpCidr {
	if o == nil {
		var ret RepositoryGeoIpCidr
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpRules) GetCidrOk() (*RepositoryGeoIpCidr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *RepositoryGeoIpRules) SetCidr(v RepositoryGeoIpCidr) {
	o.Cidr = v
}

// GetCountryCode returns the CountryCode field value
func (o *RepositoryGeoIpRules) GetCountryCode() RepositoryGeoIpCountryCode {
	if o == nil {
		var ret RepositoryGeoIpCountryCode
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpRules) GetCountryCodeOk() (*RepositoryGeoIpCountryCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *RepositoryGeoIpRules) SetCountryCode(v RepositoryGeoIpCountryCode) {
	o.CountryCode = v
}

func (o RepositoryGeoIpRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGeoIpRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	toSerialize["country_code"] = o.CountryCode
	return toSerialize, nil
}

type NullableRepositoryGeoIpRules struct {
	value *RepositoryGeoIpRules
	isSet bool
}

func (v NullableRepositoryGeoIpRules) Get() *RepositoryGeoIpRules {
	return v.value
}

func (v *NullableRepositoryGeoIpRules) Set(val *RepositoryGeoIpRules) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpRules(val *RepositoryGeoIpRules) *NullableRepositoryGeoIpRules {
	return &NullableRepositoryGeoIpRules{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

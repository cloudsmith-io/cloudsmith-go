/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.247.7
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIPTestAddress struct for RepositoryGeoIPTestAddress
type RepositoryGeoIPTestAddress struct {
	// The IP addresses to test against this repository
	Addresses []string `json:"addresses"`
}

// NewRepositoryGeoIPTestAddress instantiates a new RepositoryGeoIPTestAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIPTestAddress(addresses []string) *RepositoryGeoIPTestAddress {
	this := RepositoryGeoIPTestAddress{}
	this.Addresses = addresses
	return &this
}

// NewRepositoryGeoIPTestAddressWithDefaults instantiates a new RepositoryGeoIPTestAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIPTestAddressWithDefaults() *RepositoryGeoIPTestAddress {
	this := RepositoryGeoIPTestAddress{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *RepositoryGeoIPTestAddress) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPTestAddress) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *RepositoryGeoIPTestAddress) SetAddresses(v []string) {
	o.Addresses = v
}

func (o RepositoryGeoIPTestAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIPTestAddress struct {
	value *RepositoryGeoIPTestAddress
	isSet bool
}

func (v NullableRepositoryGeoIPTestAddress) Get() *RepositoryGeoIPTestAddress {
	return v.value
}

func (v *NullableRepositoryGeoIPTestAddress) Set(val *RepositoryGeoIPTestAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIPTestAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIPTestAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIPTestAddress(val *RepositoryGeoIPTestAddress) *NullableRepositoryGeoIPTestAddress {
	return &NullableRepositoryGeoIPTestAddress{value: val, isSet: true}
}

func (v NullableRepositoryGeoIPTestAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIPTestAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

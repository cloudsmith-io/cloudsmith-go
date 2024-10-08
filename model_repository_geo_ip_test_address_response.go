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

// checks if the RepositoryGeoIpTestAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryGeoIpTestAddressResponse{}

// RepositoryGeoIpTestAddressResponse struct for RepositoryGeoIpTestAddressResponse
type RepositoryGeoIpTestAddressResponse struct {
	// The IP address test results ordered by allowed
	Addresses []RepositoryGeoIpTestAddressResponseDict `json:"addresses"`
}

// NewRepositoryGeoIpTestAddressResponse instantiates a new RepositoryGeoIpTestAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIpTestAddressResponse(addresses []RepositoryGeoIpTestAddressResponseDict) *RepositoryGeoIpTestAddressResponse {
	this := RepositoryGeoIpTestAddressResponse{}
	this.Addresses = addresses
	return &this
}

// NewRepositoryGeoIpTestAddressResponseWithDefaults instantiates a new RepositoryGeoIpTestAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIpTestAddressResponseWithDefaults() *RepositoryGeoIpTestAddressResponse {
	this := RepositoryGeoIpTestAddressResponse{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *RepositoryGeoIpTestAddressResponse) GetAddresses() []RepositoryGeoIpTestAddressResponseDict {
	if o == nil {
		var ret []RepositoryGeoIpTestAddressResponseDict
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIpTestAddressResponse) GetAddressesOk() ([]RepositoryGeoIpTestAddressResponseDict, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *RepositoryGeoIpTestAddressResponse) SetAddresses(v []RepositoryGeoIpTestAddressResponseDict) {
	o.Addresses = v
}

func (o RepositoryGeoIpTestAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryGeoIpTestAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

type NullableRepositoryGeoIpTestAddressResponse struct {
	value *RepositoryGeoIpTestAddressResponse
	isSet bool
}

func (v NullableRepositoryGeoIpTestAddressResponse) Get() *RepositoryGeoIpTestAddressResponse {
	return v.value
}

func (v *NullableRepositoryGeoIpTestAddressResponse) Set(val *RepositoryGeoIpTestAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIpTestAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIpTestAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIpTestAddressResponse(val *RepositoryGeoIpTestAddressResponse) *NullableRepositoryGeoIpTestAddressResponse {
	return &NullableRepositoryGeoIpTestAddressResponse{value: val, isSet: true}
}

func (v NullableRepositoryGeoIpTestAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIpTestAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

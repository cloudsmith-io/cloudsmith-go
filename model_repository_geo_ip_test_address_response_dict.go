/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.202.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RepositoryGeoIPTestAddressResponseDict struct for RepositoryGeoIPTestAddressResponseDict
type RepositoryGeoIPTestAddressResponseDict struct {
	// The result of the IP test
	Allowed bool `json:"allowed"`
	// The country code of the tested IP address
	CountryCode NullableString `json:"country_code"`
	// The IP address that was tested
	IpAddress string `json:"ip_address"`
	// The reason for the result
	Reason string `json:"reason"`
}

// NewRepositoryGeoIPTestAddressResponseDict instantiates a new RepositoryGeoIPTestAddressResponseDict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryGeoIPTestAddressResponseDict(allowed bool, countryCode NullableString, ipAddress string, reason string) *RepositoryGeoIPTestAddressResponseDict {
	this := RepositoryGeoIPTestAddressResponseDict{}
	this.Allowed = allowed
	this.CountryCode = countryCode
	this.IpAddress = ipAddress
	this.Reason = reason
	return &this
}

// NewRepositoryGeoIPTestAddressResponseDictWithDefaults instantiates a new RepositoryGeoIPTestAddressResponseDict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryGeoIPTestAddressResponseDictWithDefaults() *RepositoryGeoIPTestAddressResponseDict {
	this := RepositoryGeoIPTestAddressResponseDict{}
	return &this
}

// GetAllowed returns the Allowed field value
func (o *RepositoryGeoIPTestAddressResponseDict) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPTestAddressResponseDict) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *RepositoryGeoIPTestAddressResponseDict) SetAllowed(v bool) {
	o.Allowed = v
}

// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RepositoryGeoIPTestAddressResponseDict) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryGeoIPTestAddressResponseDict) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// SetCountryCode sets field value
func (o *RepositoryGeoIPTestAddressResponseDict) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// GetIpAddress returns the IpAddress field value
func (o *RepositoryGeoIPTestAddressResponseDict) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPTestAddressResponseDict) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *RepositoryGeoIPTestAddressResponseDict) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetReason returns the Reason field value
func (o *RepositoryGeoIPTestAddressResponseDict) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RepositoryGeoIPTestAddressResponseDict) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RepositoryGeoIPTestAddressResponseDict) SetReason(v string) {
	o.Reason = v
}

func (o RepositoryGeoIPTestAddressResponseDict) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowed"] = o.Allowed
	}
	if true {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if true {
		toSerialize["ip_address"] = o.IpAddress
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryGeoIPTestAddressResponseDict struct {
	value *RepositoryGeoIPTestAddressResponseDict
	isSet bool
}

func (v NullableRepositoryGeoIPTestAddressResponseDict) Get() *RepositoryGeoIPTestAddressResponseDict {
	return v.value
}

func (v *NullableRepositoryGeoIPTestAddressResponseDict) Set(val *RepositoryGeoIPTestAddressResponseDict) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryGeoIPTestAddressResponseDict) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryGeoIPTestAddressResponseDict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryGeoIPTestAddressResponseDict(val *RepositoryGeoIPTestAddressResponseDict) *NullableRepositoryGeoIPTestAddressResponseDict {
	return &NullableRepositoryGeoIPTestAddressResponseDict{value: val, isSet: true}
}

func (v NullableRepositoryGeoIPTestAddressResponseDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryGeoIPTestAddressResponseDict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// UsageLimits struct for UsageLimits
type UsageLimits struct {
	Bandwidth AllocatedLimit `json:"bandwidth"`
	Storage   AllocatedLimit `json:"storage"`
}

// NewUsageLimits instantiates a new UsageLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimits(bandwidth AllocatedLimit, storage AllocatedLimit) *UsageLimits {
	this := UsageLimits{}
	this.Bandwidth = bandwidth
	this.Storage = storage
	return &this
}

// NewUsageLimitsWithDefaults instantiates a new UsageLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLimitsWithDefaults() *UsageLimits {
	this := UsageLimits{}
	return &this
}

// GetBandwidth returns the Bandwidth field value
func (o *UsageLimits) GetBandwidth() AllocatedLimit {
	if o == nil {
		var ret AllocatedLimit
		return ret
	}

	return o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value
// and a boolean to check if the value has been set.
func (o *UsageLimits) GetBandwidthOk() (*AllocatedLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bandwidth, true
}

// SetBandwidth sets field value
func (o *UsageLimits) SetBandwidth(v AllocatedLimit) {
	o.Bandwidth = v
}

// GetStorage returns the Storage field value
func (o *UsageLimits) GetStorage() AllocatedLimit {
	if o == nil {
		var ret AllocatedLimit
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *UsageLimits) GetStorageOk() (*AllocatedLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *UsageLimits) SetStorage(v AllocatedLimit) {
	o.Storage = v
}

func (o UsageLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if true {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageLimits struct {
	value *UsageLimits
	isSet bool
}

func (v NullableUsageLimits) Get() *UsageLimits {
	return v.value
}

func (v *NullableUsageLimits) Set(val *UsageLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLimits(val *UsageLimits) *NullableUsageLimits {
	return &NullableUsageLimits{value: val, isSet: true}
}

func (v NullableUsageLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

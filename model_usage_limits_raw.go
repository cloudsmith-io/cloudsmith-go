/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.498.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the UsageLimitsRaw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageLimitsRaw{}

// UsageLimitsRaw struct for UsageLimitsRaw
type UsageLimitsRaw struct {
	Bandwidth AllocatedLimitRaw        `json:"bandwidth"`
	Storage   StorageAllocatedLimitRaw `json:"storage"`
}

// NewUsageLimitsRaw instantiates a new UsageLimitsRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimitsRaw(bandwidth AllocatedLimitRaw, storage StorageAllocatedLimitRaw) *UsageLimitsRaw {
	this := UsageLimitsRaw{}
	this.Bandwidth = bandwidth
	this.Storage = storage
	return &this
}

// NewUsageLimitsRawWithDefaults instantiates a new UsageLimitsRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLimitsRawWithDefaults() *UsageLimitsRaw {
	this := UsageLimitsRaw{}
	return &this
}

// GetBandwidth returns the Bandwidth field value
func (o *UsageLimitsRaw) GetBandwidth() AllocatedLimitRaw {
	if o == nil {
		var ret AllocatedLimitRaw
		return ret
	}

	return o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value
// and a boolean to check if the value has been set.
func (o *UsageLimitsRaw) GetBandwidthOk() (*AllocatedLimitRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bandwidth, true
}

// SetBandwidth sets field value
func (o *UsageLimitsRaw) SetBandwidth(v AllocatedLimitRaw) {
	o.Bandwidth = v
}

// GetStorage returns the Storage field value
func (o *UsageLimitsRaw) GetStorage() StorageAllocatedLimitRaw {
	if o == nil {
		var ret StorageAllocatedLimitRaw
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *UsageLimitsRaw) GetStorageOk() (*StorageAllocatedLimitRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *UsageLimitsRaw) SetStorage(v StorageAllocatedLimitRaw) {
	o.Storage = v
}

func (o UsageLimitsRaw) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageLimitsRaw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bandwidth"] = o.Bandwidth
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

type NullableUsageLimitsRaw struct {
	value *UsageLimitsRaw
	isSet bool
}

func (v NullableUsageLimitsRaw) Get() *UsageLimitsRaw {
	return v.value
}

func (v *NullableUsageLimitsRaw) Set(val *UsageLimitsRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLimitsRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLimitsRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLimitsRaw(val *UsageLimitsRaw) *NullableUsageLimitsRaw {
	return &NullableUsageLimitsRaw{value: val, isSet: true}
}

func (v NullableUsageLimitsRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLimitsRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

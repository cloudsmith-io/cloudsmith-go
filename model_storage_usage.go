/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.390.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the StorageUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageUsage{}

// StorageUsage struct for StorageUsage
type StorageUsage struct {
	Limit      *string `json:"limit,omitempty"`
	Peak       *string `json:"peak,omitempty"`
	Percentage *string `json:"percentage,omitempty"`
	Used       *string `json:"used,omitempty"`
}

// NewStorageUsage instantiates a new StorageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageUsage() *StorageUsage {
	this := StorageUsage{}
	return &this
}

// NewStorageUsageWithDefaults instantiates a new StorageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageUsageWithDefaults() *StorageUsage {
	this := StorageUsage{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *StorageUsage) GetLimit() string {
	if o == nil || IsNil(o.Limit) {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsage) GetLimitOk() (*string, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *StorageUsage) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *StorageUsage) SetLimit(v string) {
	o.Limit = &v
}

// GetPeak returns the Peak field value if set, zero value otherwise.
func (o *StorageUsage) GetPeak() string {
	if o == nil || IsNil(o.Peak) {
		var ret string
		return ret
	}
	return *o.Peak
}

// GetPeakOk returns a tuple with the Peak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsage) GetPeakOk() (*string, bool) {
	if o == nil || IsNil(o.Peak) {
		return nil, false
	}
	return o.Peak, true
}

// HasPeak returns a boolean if a field has been set.
func (o *StorageUsage) HasPeak() bool {
	if o != nil && !IsNil(o.Peak) {
		return true
	}

	return false
}

// SetPeak gets a reference to the given string and assigns it to the Peak field.
func (o *StorageUsage) SetPeak(v string) {
	o.Peak = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *StorageUsage) GetPercentage() string {
	if o == nil || IsNil(o.Percentage) {
		var ret string
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsage) GetPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *StorageUsage) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given string and assigns it to the Percentage field.
func (o *StorageUsage) SetPercentage(v string) {
	o.Percentage = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageUsage) GetUsed() string {
	if o == nil || IsNil(o.Used) {
		var ret string
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsage) GetUsedOk() (*string, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageUsage) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given string and assigns it to the Used field.
func (o *StorageUsage) SetUsed(v string) {
	o.Used = &v
}

func (o StorageUsage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Peak) {
		toSerialize["peak"] = o.Peak
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableStorageUsage struct {
	value *StorageUsage
	isSet bool
}

func (v NullableStorageUsage) Get() *StorageUsage {
	return v.value
}

func (v *NullableStorageUsage) Set(val *StorageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageUsage(val *StorageUsage) *NullableStorageUsage {
	return &NullableStorageUsage{value: val, isSet: true}
}

func (v NullableStorageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

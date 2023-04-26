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

// UsageRaw struct for UsageRaw
type UsageRaw struct {
	Limit      *int64   `json:"limit,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
	Used       *int64   `json:"used,omitempty"`
}

// NewUsageRaw instantiates a new UsageRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageRaw() *UsageRaw {
	this := UsageRaw{}
	return &this
}

// NewUsageRawWithDefaults instantiates a new UsageRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageRawWithDefaults() *UsageRaw {
	this := UsageRaw{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *UsageRaw) GetLimit() int64 {
	if o == nil || isNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageRaw) GetLimitOk() (*int64, bool) {
	if o == nil || isNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *UsageRaw) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *UsageRaw) SetLimit(v int64) {
	o.Limit = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *UsageRaw) GetPercentage() float64 {
	if o == nil || isNil(o.Percentage) {
		var ret float64
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageRaw) GetPercentageOk() (*float64, bool) {
	if o == nil || isNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *UsageRaw) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float64 and assigns it to the Percentage field.
func (o *UsageRaw) SetPercentage(v float64) {
	o.Percentage = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *UsageRaw) GetUsed() int64 {
	if o == nil || isNil(o.Used) {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageRaw) GetUsedOk() (*int64, bool) {
	if o == nil || isNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *UsageRaw) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *UsageRaw) SetUsed(v int64) {
	o.Used = &v
}

func (o UsageRaw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableUsageRaw struct {
	value *UsageRaw
	isSet bool
}

func (v NullableUsageRaw) Get() *UsageRaw {
	return v.value
}

func (v *NullableUsageRaw) Set(val *UsageRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageRaw(val *UsageRaw) *NullableUsageRaw {
	return &NullableUsageRaw{value: val, isSet: true}
}

func (v NullableUsageRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

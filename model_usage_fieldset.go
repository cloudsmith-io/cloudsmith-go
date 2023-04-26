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

// UsageFieldset struct for UsageFieldset
type UsageFieldset struct {
	Display UsageLimits    `json:"display"`
	Raw     UsageLimitsRaw `json:"raw"`
}

// NewUsageFieldset instantiates a new UsageFieldset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageFieldset(display UsageLimits, raw UsageLimitsRaw) *UsageFieldset {
	this := UsageFieldset{}
	this.Display = display
	this.Raw = raw
	return &this
}

// NewUsageFieldsetWithDefaults instantiates a new UsageFieldset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageFieldsetWithDefaults() *UsageFieldset {
	this := UsageFieldset{}
	return &this
}

// GetDisplay returns the Display field value
func (o *UsageFieldset) GetDisplay() UsageLimits {
	if o == nil {
		var ret UsageLimits
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *UsageFieldset) GetDisplayOk() (*UsageLimits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *UsageFieldset) SetDisplay(v UsageLimits) {
	o.Display = v
}

// GetRaw returns the Raw field value
func (o *UsageFieldset) GetRaw() UsageLimitsRaw {
	if o == nil {
		var ret UsageLimitsRaw
		return ret
	}

	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value
// and a boolean to check if the value has been set.
func (o *UsageFieldset) GetRawOk() (*UsageLimitsRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raw, true
}

// SetRaw sets field value
func (o *UsageFieldset) SetRaw(v UsageLimitsRaw) {
	o.Raw = v
}

func (o UsageFieldset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["display"] = o.Display
	}
	if true {
		toSerialize["raw"] = o.Raw
	}
	return json.Marshal(toSerialize)
}

type NullableUsageFieldset struct {
	value *UsageFieldset
	isSet bool
}

func (v NullableUsageFieldset) Get() *UsageFieldset {
	return v.value
}

func (v *NullableUsageFieldset) Set(val *UsageFieldset) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageFieldset) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageFieldset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageFieldset(val *UsageFieldset) *NullableUsageFieldset {
	return &NullableUsageFieldset{value: val, isSet: true}
}

func (v NullableUsageFieldset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageFieldset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

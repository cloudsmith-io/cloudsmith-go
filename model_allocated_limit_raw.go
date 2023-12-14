/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.327.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the AllocatedLimitRaw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllocatedLimitRaw{}

// AllocatedLimitRaw struct for AllocatedLimitRaw
type AllocatedLimitRaw struct {
	Configured     *int64   `json:"configured,omitempty"`
	PercentageUsed *float64 `json:"percentage_used,omitempty"`
	PlanLimit      *int64   `json:"plan_limit,omitempty"`
	Used           *int64   `json:"used,omitempty"`
}

// NewAllocatedLimitRaw instantiates a new AllocatedLimitRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocatedLimitRaw() *AllocatedLimitRaw {
	this := AllocatedLimitRaw{}
	return &this
}

// NewAllocatedLimitRawWithDefaults instantiates a new AllocatedLimitRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocatedLimitRawWithDefaults() *AllocatedLimitRaw {
	this := AllocatedLimitRaw{}
	return &this
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *AllocatedLimitRaw) GetConfigured() int64 {
	if o == nil || IsNil(o.Configured) {
		var ret int64
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimitRaw) GetConfiguredOk() (*int64, bool) {
	if o == nil || IsNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *AllocatedLimitRaw) HasConfigured() bool {
	if o != nil && !IsNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given int64 and assigns it to the Configured field.
func (o *AllocatedLimitRaw) SetConfigured(v int64) {
	o.Configured = &v
}

// GetPercentageUsed returns the PercentageUsed field value if set, zero value otherwise.
func (o *AllocatedLimitRaw) GetPercentageUsed() float64 {
	if o == nil || IsNil(o.PercentageUsed) {
		var ret float64
		return ret
	}
	return *o.PercentageUsed
}

// GetPercentageUsedOk returns a tuple with the PercentageUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimitRaw) GetPercentageUsedOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentageUsed) {
		return nil, false
	}
	return o.PercentageUsed, true
}

// HasPercentageUsed returns a boolean if a field has been set.
func (o *AllocatedLimitRaw) HasPercentageUsed() bool {
	if o != nil && !IsNil(o.PercentageUsed) {
		return true
	}

	return false
}

// SetPercentageUsed gets a reference to the given float64 and assigns it to the PercentageUsed field.
func (o *AllocatedLimitRaw) SetPercentageUsed(v float64) {
	o.PercentageUsed = &v
}

// GetPlanLimit returns the PlanLimit field value if set, zero value otherwise.
func (o *AllocatedLimitRaw) GetPlanLimit() int64 {
	if o == nil || IsNil(o.PlanLimit) {
		var ret int64
		return ret
	}
	return *o.PlanLimit
}

// GetPlanLimitOk returns a tuple with the PlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimitRaw) GetPlanLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanLimit) {
		return nil, false
	}
	return o.PlanLimit, true
}

// HasPlanLimit returns a boolean if a field has been set.
func (o *AllocatedLimitRaw) HasPlanLimit() bool {
	if o != nil && !IsNil(o.PlanLimit) {
		return true
	}

	return false
}

// SetPlanLimit gets a reference to the given int64 and assigns it to the PlanLimit field.
func (o *AllocatedLimitRaw) SetPlanLimit(v int64) {
	o.PlanLimit = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *AllocatedLimitRaw) GetUsed() int64 {
	if o == nil || IsNil(o.Used) {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimitRaw) GetUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *AllocatedLimitRaw) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *AllocatedLimitRaw) SetUsed(v int64) {
	o.Used = &v
}

func (o AllocatedLimitRaw) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllocatedLimitRaw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	if !IsNil(o.PercentageUsed) {
		toSerialize["percentage_used"] = o.PercentageUsed
	}
	if !IsNil(o.PlanLimit) {
		toSerialize["plan_limit"] = o.PlanLimit
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableAllocatedLimitRaw struct {
	value *AllocatedLimitRaw
	isSet bool
}

func (v NullableAllocatedLimitRaw) Get() *AllocatedLimitRaw {
	return v.value
}

func (v *NullableAllocatedLimitRaw) Set(val *AllocatedLimitRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocatedLimitRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocatedLimitRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocatedLimitRaw(val *AllocatedLimitRaw) *NullableAllocatedLimitRaw {
	return &NullableAllocatedLimitRaw{value: val, isSet: true}
}

func (v NullableAllocatedLimitRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocatedLimitRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

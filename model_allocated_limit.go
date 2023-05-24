/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.273.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// AllocatedLimit struct for AllocatedLimit
type AllocatedLimit struct {
	Configured     *string `json:"configured,omitempty"`
	PercentageUsed *string `json:"percentage_used,omitempty"`
	PlanLimit      *string `json:"plan_limit,omitempty"`
	Used           *string `json:"used,omitempty"`
}

// NewAllocatedLimit instantiates a new AllocatedLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocatedLimit() *AllocatedLimit {
	this := AllocatedLimit{}
	return &this
}

// NewAllocatedLimitWithDefaults instantiates a new AllocatedLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocatedLimitWithDefaults() *AllocatedLimit {
	this := AllocatedLimit{}
	return &this
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *AllocatedLimit) GetConfigured() string {
	if o == nil || isNil(o.Configured) {
		var ret string
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimit) GetConfiguredOk() (*string, bool) {
	if o == nil || isNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *AllocatedLimit) HasConfigured() bool {
	if o != nil && !isNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given string and assigns it to the Configured field.
func (o *AllocatedLimit) SetConfigured(v string) {
	o.Configured = &v
}

// GetPercentageUsed returns the PercentageUsed field value if set, zero value otherwise.
func (o *AllocatedLimit) GetPercentageUsed() string {
	if o == nil || isNil(o.PercentageUsed) {
		var ret string
		return ret
	}
	return *o.PercentageUsed
}

// GetPercentageUsedOk returns a tuple with the PercentageUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimit) GetPercentageUsedOk() (*string, bool) {
	if o == nil || isNil(o.PercentageUsed) {
		return nil, false
	}
	return o.PercentageUsed, true
}

// HasPercentageUsed returns a boolean if a field has been set.
func (o *AllocatedLimit) HasPercentageUsed() bool {
	if o != nil && !isNil(o.PercentageUsed) {
		return true
	}

	return false
}

// SetPercentageUsed gets a reference to the given string and assigns it to the PercentageUsed field.
func (o *AllocatedLimit) SetPercentageUsed(v string) {
	o.PercentageUsed = &v
}

// GetPlanLimit returns the PlanLimit field value if set, zero value otherwise.
func (o *AllocatedLimit) GetPlanLimit() string {
	if o == nil || isNil(o.PlanLimit) {
		var ret string
		return ret
	}
	return *o.PlanLimit
}

// GetPlanLimitOk returns a tuple with the PlanLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimit) GetPlanLimitOk() (*string, bool) {
	if o == nil || isNil(o.PlanLimit) {
		return nil, false
	}
	return o.PlanLimit, true
}

// HasPlanLimit returns a boolean if a field has been set.
func (o *AllocatedLimit) HasPlanLimit() bool {
	if o != nil && !isNil(o.PlanLimit) {
		return true
	}

	return false
}

// SetPlanLimit gets a reference to the given string and assigns it to the PlanLimit field.
func (o *AllocatedLimit) SetPlanLimit(v string) {
	o.PlanLimit = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *AllocatedLimit) GetUsed() string {
	if o == nil || isNil(o.Used) {
		var ret string
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedLimit) GetUsedOk() (*string, bool) {
	if o == nil || isNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *AllocatedLimit) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given string and assigns it to the Used field.
func (o *AllocatedLimit) SetUsed(v string) {
	o.Used = &v
}

func (o AllocatedLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	if !isNil(o.PercentageUsed) {
		toSerialize["percentage_used"] = o.PercentageUsed
	}
	if !isNil(o.PlanLimit) {
		toSerialize["plan_limit"] = o.PlanLimit
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableAllocatedLimit struct {
	value *AllocatedLimit
	isSet bool
}

func (v NullableAllocatedLimit) Get() *AllocatedLimit {
	return v.value
}

func (v *NullableAllocatedLimit) Set(val *AllocatedLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocatedLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocatedLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocatedLimit(val *AllocatedLimit) *NullableAllocatedLimit {
	return &NullableAllocatedLimit{value: val, isSet: true}
}

func (v NullableAllocatedLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocatedLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

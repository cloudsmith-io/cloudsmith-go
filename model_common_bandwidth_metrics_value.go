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

// CommonBandwidthMetricsValue struct for CommonBandwidthMetricsValue
type CommonBandwidthMetricsValue struct {
	Display string  `json:"display"`
	Units   *string `json:"units,omitempty"`
	Value   int64   `json:"value"`
}

// NewCommonBandwidthMetricsValue instantiates a new CommonBandwidthMetricsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBandwidthMetricsValue(display string, value int64) *CommonBandwidthMetricsValue {
	this := CommonBandwidthMetricsValue{}
	this.Display = display
	var units string = "bytes"
	this.Units = &units
	this.Value = value
	return &this
}

// NewCommonBandwidthMetricsValueWithDefaults instantiates a new CommonBandwidthMetricsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBandwidthMetricsValueWithDefaults() *CommonBandwidthMetricsValue {
	this := CommonBandwidthMetricsValue{}
	var units string = "bytes"
	this.Units = &units
	return &this
}

// GetDisplay returns the Display field value
func (o *CommonBandwidthMetricsValue) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetricsValue) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CommonBandwidthMetricsValue) SetDisplay(v string) {
	o.Display = v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *CommonBandwidthMetricsValue) GetUnits() string {
	if o == nil || isNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetricsValue) GetUnitsOk() (*string, bool) {
	if o == nil || isNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *CommonBandwidthMetricsValue) HasUnits() bool {
	if o != nil && !isNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *CommonBandwidthMetricsValue) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value
func (o *CommonBandwidthMetricsValue) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CommonBandwidthMetricsValue) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CommonBandwidthMetricsValue) SetValue(v int64) {
	o.Value = v
}

func (o CommonBandwidthMetricsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["display"] = o.Display
	}
	if !isNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCommonBandwidthMetricsValue struct {
	value *CommonBandwidthMetricsValue
	isSet bool
}

func (v NullableCommonBandwidthMetricsValue) Get() *CommonBandwidthMetricsValue {
	return v.value
}

func (v *NullableCommonBandwidthMetricsValue) Set(val *CommonBandwidthMetricsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBandwidthMetricsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBandwidthMetricsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBandwidthMetricsValue(val *CommonBandwidthMetricsValue) *NullableCommonBandwidthMetricsValue {
	return &NullableCommonBandwidthMetricsValue{value: val, isSet: true}
}

func (v NullableCommonBandwidthMetricsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBandwidthMetricsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

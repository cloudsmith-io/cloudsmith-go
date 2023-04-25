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

// CommonDownloadsMetricsValue struct for CommonDownloadsMetricsValue
type CommonDownloadsMetricsValue struct {
	Value int64 `json:"value"`
}

// NewCommonDownloadsMetricsValue instantiates a new CommonDownloadsMetricsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonDownloadsMetricsValue(value int64) *CommonDownloadsMetricsValue {
	this := CommonDownloadsMetricsValue{}
	this.Value = value
	return &this
}

// NewCommonDownloadsMetricsValueWithDefaults instantiates a new CommonDownloadsMetricsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonDownloadsMetricsValueWithDefaults() *CommonDownloadsMetricsValue {
	this := CommonDownloadsMetricsValue{}
	return &this
}

// GetValue returns the Value field value
func (o *CommonDownloadsMetricsValue) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CommonDownloadsMetricsValue) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CommonDownloadsMetricsValue) SetValue(v int64) {
	o.Value = v
}

func (o CommonDownloadsMetricsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCommonDownloadsMetricsValue struct {
	value *CommonDownloadsMetricsValue
	isSet bool
}

func (v NullableCommonDownloadsMetricsValue) Get() *CommonDownloadsMetricsValue {
	return v.value
}

func (v *NullableCommonDownloadsMetricsValue) Set(val *CommonDownloadsMetricsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonDownloadsMetricsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonDownloadsMetricsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonDownloadsMetricsValue(val *CommonDownloadsMetricsValue) *NullableCommonDownloadsMetricsValue {
	return &NullableCommonDownloadsMetricsValue{value: val, isSet: true}
}

func (v NullableCommonDownloadsMetricsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonDownloadsMetricsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

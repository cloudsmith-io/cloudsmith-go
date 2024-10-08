/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.533.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the CommonDownloadsMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonDownloadsMetrics{}

// CommonDownloadsMetrics struct for CommonDownloadsMetrics
type CommonDownloadsMetrics struct {
	Average CommonDownloadsMetricsValue `json:"average"`
	Highest CommonDownloadsMetricsValue `json:"highest"`
	Lowest  CommonDownloadsMetricsValue `json:"lowest"`
	Total   CommonDownloadsMetricsValue `json:"total"`
}

// NewCommonDownloadsMetrics instantiates a new CommonDownloadsMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonDownloadsMetrics(average CommonDownloadsMetricsValue, highest CommonDownloadsMetricsValue, lowest CommonDownloadsMetricsValue, total CommonDownloadsMetricsValue) *CommonDownloadsMetrics {
	this := CommonDownloadsMetrics{}
	this.Average = average
	this.Highest = highest
	this.Lowest = lowest
	this.Total = total
	return &this
}

// NewCommonDownloadsMetricsWithDefaults instantiates a new CommonDownloadsMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonDownloadsMetricsWithDefaults() *CommonDownloadsMetrics {
	this := CommonDownloadsMetrics{}
	return &this
}

// GetAverage returns the Average field value
func (o *CommonDownloadsMetrics) GetAverage() CommonDownloadsMetricsValue {
	if o == nil {
		var ret CommonDownloadsMetricsValue
		return ret
	}

	return o.Average
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
func (o *CommonDownloadsMetrics) GetAverageOk() (*CommonDownloadsMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Average, true
}

// SetAverage sets field value
func (o *CommonDownloadsMetrics) SetAverage(v CommonDownloadsMetricsValue) {
	o.Average = v
}

// GetHighest returns the Highest field value
func (o *CommonDownloadsMetrics) GetHighest() CommonDownloadsMetricsValue {
	if o == nil {
		var ret CommonDownloadsMetricsValue
		return ret
	}

	return o.Highest
}

// GetHighestOk returns a tuple with the Highest field value
// and a boolean to check if the value has been set.
func (o *CommonDownloadsMetrics) GetHighestOk() (*CommonDownloadsMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Highest, true
}

// SetHighest sets field value
func (o *CommonDownloadsMetrics) SetHighest(v CommonDownloadsMetricsValue) {
	o.Highest = v
}

// GetLowest returns the Lowest field value
func (o *CommonDownloadsMetrics) GetLowest() CommonDownloadsMetricsValue {
	if o == nil {
		var ret CommonDownloadsMetricsValue
		return ret
	}

	return o.Lowest
}

// GetLowestOk returns a tuple with the Lowest field value
// and a boolean to check if the value has been set.
func (o *CommonDownloadsMetrics) GetLowestOk() (*CommonDownloadsMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lowest, true
}

// SetLowest sets field value
func (o *CommonDownloadsMetrics) SetLowest(v CommonDownloadsMetricsValue) {
	o.Lowest = v
}

// GetTotal returns the Total field value
func (o *CommonDownloadsMetrics) GetTotal() CommonDownloadsMetricsValue {
	if o == nil {
		var ret CommonDownloadsMetricsValue
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CommonDownloadsMetrics) GetTotalOk() (*CommonDownloadsMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CommonDownloadsMetrics) SetTotal(v CommonDownloadsMetricsValue) {
	o.Total = v
}

func (o CommonDownloadsMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonDownloadsMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["average"] = o.Average
	toSerialize["highest"] = o.Highest
	toSerialize["lowest"] = o.Lowest
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableCommonDownloadsMetrics struct {
	value *CommonDownloadsMetrics
	isSet bool
}

func (v NullableCommonDownloadsMetrics) Get() *CommonDownloadsMetrics {
	return v.value
}

func (v *NullableCommonDownloadsMetrics) Set(val *CommonDownloadsMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonDownloadsMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonDownloadsMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonDownloadsMetrics(val *CommonDownloadsMetrics) *NullableCommonDownloadsMetrics {
	return &NullableCommonDownloadsMetrics{value: val, isSet: true}
}

func (v NullableCommonDownloadsMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonDownloadsMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

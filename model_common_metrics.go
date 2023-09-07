/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.297.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the CommonMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonMetrics{}

// CommonMetrics struct for CommonMetrics
type CommonMetrics struct {
	Active    *int64                 `json:"active,omitempty"`
	Bandwidth CommonBandwidthMetrics `json:"bandwidth"`
	Downloads CommonDownloadsMetrics `json:"downloads"`
	Inactive  *int64                 `json:"inactive,omitempty"`
	Total     *int64                 `json:"total,omitempty"`
}

// NewCommonMetrics instantiates a new CommonMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonMetrics(bandwidth CommonBandwidthMetrics, downloads CommonDownloadsMetrics) *CommonMetrics {
	this := CommonMetrics{}
	this.Bandwidth = bandwidth
	this.Downloads = downloads
	return &this
}

// NewCommonMetricsWithDefaults instantiates a new CommonMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonMetricsWithDefaults() *CommonMetrics {
	this := CommonMetrics{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CommonMetrics) GetActive() int64 {
	if o == nil || IsNil(o.Active) {
		var ret int64
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonMetrics) GetActiveOk() (*int64, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CommonMetrics) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int64 and assigns it to the Active field.
func (o *CommonMetrics) SetActive(v int64) {
	o.Active = &v
}

// GetBandwidth returns the Bandwidth field value
func (o *CommonMetrics) GetBandwidth() CommonBandwidthMetrics {
	if o == nil {
		var ret CommonBandwidthMetrics
		return ret
	}

	return o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value
// and a boolean to check if the value has been set.
func (o *CommonMetrics) GetBandwidthOk() (*CommonBandwidthMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bandwidth, true
}

// SetBandwidth sets field value
func (o *CommonMetrics) SetBandwidth(v CommonBandwidthMetrics) {
	o.Bandwidth = v
}

// GetDownloads returns the Downloads field value
func (o *CommonMetrics) GetDownloads() CommonDownloadsMetrics {
	if o == nil {
		var ret CommonDownloadsMetrics
		return ret
	}

	return o.Downloads
}

// GetDownloadsOk returns a tuple with the Downloads field value
// and a boolean to check if the value has been set.
func (o *CommonMetrics) GetDownloadsOk() (*CommonDownloadsMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Downloads, true
}

// SetDownloads sets field value
func (o *CommonMetrics) SetDownloads(v CommonDownloadsMetrics) {
	o.Downloads = v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *CommonMetrics) GetInactive() int64 {
	if o == nil || IsNil(o.Inactive) {
		var ret int64
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonMetrics) GetInactiveOk() (*int64, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *CommonMetrics) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given int64 and assigns it to the Inactive field.
func (o *CommonMetrics) SetInactive(v int64) {
	o.Inactive = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CommonMetrics) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonMetrics) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CommonMetrics) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *CommonMetrics) SetTotal(v int64) {
	o.Total = &v
}

func (o CommonMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["bandwidth"] = o.Bandwidth
	toSerialize["downloads"] = o.Downloads
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableCommonMetrics struct {
	value *CommonMetrics
	isSet bool
}

func (v NullableCommonMetrics) Get() *CommonMetrics {
	return v.value
}

func (v *NullableCommonMetrics) Set(val *CommonMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonMetrics(val *CommonMetrics) *NullableCommonMetrics {
	return &NullableCommonMetrics{value: val, isSet: true}
}

func (v NullableCommonMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

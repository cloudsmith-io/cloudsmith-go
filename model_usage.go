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

// checks if the Usage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Usage{}

// Usage struct for Usage
type Usage struct {
	Limit      *string `json:"limit,omitempty"`
	Percentage *string `json:"percentage,omitempty"`
	Used       *string `json:"used,omitempty"`
}

// NewUsage instantiates a new Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsage() *Usage {
	this := Usage{}
	return &this
}

// NewUsageWithDefaults instantiates a new Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageWithDefaults() *Usage {
	this := Usage{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Usage) GetLimit() string {
	if o == nil || IsNil(o.Limit) {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetLimitOk() (*string, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Usage) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *Usage) SetLimit(v string) {
	o.Limit = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *Usage) GetPercentage() string {
	if o == nil || IsNil(o.Percentage) {
		var ret string
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *Usage) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given string and assigns it to the Percentage field.
func (o *Usage) SetPercentage(v string) {
	o.Percentage = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *Usage) GetUsed() string {
	if o == nil || IsNil(o.Used) {
		var ret string
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetUsedOk() (*string, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *Usage) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given string and assigns it to the Used field.
func (o *Usage) SetUsed(v string) {
	o.Used = &v
}

func (o Usage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Usage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableUsage struct {
	value *Usage
	isSet bool
}

func (v NullableUsage) Get() *Usage {
	return v.value
}

func (v *NullableUsage) Set(val *Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsage(val *Usage) *NullableUsage {
	return &NullableUsage{value: val, isSet: true}
}

func (v NullableUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.121.3
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// Quota struct for Quota
type Quota struct {
	//
	Usage map[string]interface{} `json:"usage"`
}

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota(usage map[string]interface{}) *Quota {
	this := Quota{}
	this.Usage = usage
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetUsage returns the Usage field value
func (o *Quota) GetUsage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *Quota) GetUsageOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *Quota) SetUsage(v map[string]interface{}) {
	o.Usage = v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

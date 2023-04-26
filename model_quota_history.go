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

// QuotaHistory struct for QuotaHistory
type QuotaHistory struct {
	History []History `json:"history"`
}

// NewQuotaHistory instantiates a new QuotaHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaHistory(history []History) *QuotaHistory {
	this := QuotaHistory{}
	this.History = history
	return &this
}

// NewQuotaHistoryWithDefaults instantiates a new QuotaHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaHistoryWithDefaults() *QuotaHistory {
	this := QuotaHistory{}
	return &this
}

// GetHistory returns the History field value
func (o *QuotaHistory) GetHistory() []History {
	if o == nil {
		var ret []History
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *QuotaHistory) GetHistoryOk() ([]History, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *QuotaHistory) SetHistory(v []History) {
	o.History = v
}

func (o QuotaHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["history"] = o.History
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaHistory struct {
	value *QuotaHistory
	isSet bool
}

func (v NullableQuotaHistory) Get() *QuotaHistory {
	return v.value
}

func (v *NullableQuotaHistory) Set(val *QuotaHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaHistory(val *QuotaHistory) *NullableQuotaHistory {
	return &NullableQuotaHistory{value: val, isSet: true}
}

func (v NullableQuotaHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.42.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// EntitlementsSync struct for EntitlementsSync
type EntitlementsSync struct {
	// The source repository slug (in the same owner namespace).
	Source string `json:"source"`
}

// NewEntitlementsSync instantiates a new EntitlementsSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsSync(source string) *EntitlementsSync {
	this := EntitlementsSync{}
	this.Source = source
	return &this
}

// NewEntitlementsSyncWithDefaults instantiates a new EntitlementsSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsSyncWithDefaults() *EntitlementsSync {
	this := EntitlementsSync{}
	return &this
}

// GetSource returns the Source field value
func (o *EntitlementsSync) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EntitlementsSync) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EntitlementsSync) SetSource(v string) {
	o.Source = v
}

func (o EntitlementsSync) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementsSync struct {
	value *EntitlementsSync
	isSet bool
}

func (v NullableEntitlementsSync) Get() *EntitlementsSync {
	return v.value
}

func (v *NullableEntitlementsSync) Set(val *EntitlementsSync) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsSync) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsSync(val *EntitlementsSync) *NullableEntitlementsSync {
	return &NullableEntitlementsSync{value: val, isSet: true}
}

func (v NullableEntitlementsSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

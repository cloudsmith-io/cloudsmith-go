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

// checks if the ResourcesRateCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcesRateCheck{}

// ResourcesRateCheck struct for ResourcesRateCheck
type ResourcesRateCheck struct {
	// Rate limit values per resource
	Resources *map[string]RateCheck `json:"resources,omitempty"`
}

// NewResourcesRateCheck instantiates a new ResourcesRateCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesRateCheck() *ResourcesRateCheck {
	this := ResourcesRateCheck{}
	return &this
}

// NewResourcesRateCheckWithDefaults instantiates a new ResourcesRateCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesRateCheckWithDefaults() *ResourcesRateCheck {
	this := ResourcesRateCheck{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourcesRateCheck) GetResources() map[string]RateCheck {
	if o == nil || IsNil(o.Resources) {
		var ret map[string]RateCheck
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcesRateCheck) GetResourcesOk() (*map[string]RateCheck, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourcesRateCheck) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given map[string]RateCheck and assigns it to the Resources field.
func (o *ResourcesRateCheck) SetResources(v map[string]RateCheck) {
	o.Resources = &v
}

func (o ResourcesRateCheck) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcesRateCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableResourcesRateCheck struct {
	value *ResourcesRateCheck
	isSet bool
}

func (v NullableResourcesRateCheck) Get() *ResourcesRateCheck {
	return v.value
}

func (v *NullableResourcesRateCheck) Set(val *ResourcesRateCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesRateCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesRateCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesRateCheck(val *ResourcesRateCheck) *NullableResourcesRateCheck {
	return &NullableResourcesRateCheck{value: val, isSet: true}
}

func (v NullableResourcesRateCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesRateCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.206.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// PackageDependencies struct for PackageDependencies
type PackageDependencies struct {
	Dependencies []PackageDependency `json:"dependencies"`
}

// NewPackageDependencies instantiates a new PackageDependencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDependencies(dependencies []PackageDependency) *PackageDependencies {
	this := PackageDependencies{}
	this.Dependencies = dependencies
	return &this
}

// NewPackageDependenciesWithDefaults instantiates a new PackageDependencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDependenciesWithDefaults() *PackageDependencies {
	this := PackageDependencies{}
	return &this
}

// GetDependencies returns the Dependencies field value
func (o *PackageDependencies) GetDependencies() []PackageDependency {
	if o == nil {
		var ret []PackageDependency
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
func (o *PackageDependencies) GetDependenciesOk() ([]PackageDependency, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// SetDependencies sets field value
func (o *PackageDependencies) SetDependencies(v []PackageDependency) {
	o.Dependencies = v
}

func (o PackageDependencies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dependencies"] = o.Dependencies
	}
	return json.Marshal(toSerialize)
}

type NullablePackageDependencies struct {
	value *PackageDependencies
	isSet bool
}

func (v NullablePackageDependencies) Get() *PackageDependencies {
	return v.value
}

func (v *NullablePackageDependencies) Set(val *PackageDependencies) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDependencies) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDependencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDependencies(val *PackageDependencies) *NullablePackageDependencies {
	return &NullablePackageDependencies{value: val, isSet: true}
}

func (v NullablePackageDependencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDependencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

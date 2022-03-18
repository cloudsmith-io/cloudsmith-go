/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.42.3
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// PackageDependenciesDependencies struct for PackageDependenciesDependencies
type PackageDependenciesDependencies struct {
	//
	DepType *string `json:"dep_type,omitempty"`
	// The name of the package dependency.
	Name *string `json:"name,omitempty"`
	//
	Operator *string `json:"operator,omitempty"`
	// The raw dependency version (if any).
	Version *string `json:"version,omitempty"`
}

// NewPackageDependenciesDependencies instantiates a new PackageDependenciesDependencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDependenciesDependencies() *PackageDependenciesDependencies {
	this := PackageDependenciesDependencies{}
	return &this
}

// NewPackageDependenciesDependenciesWithDefaults instantiates a new PackageDependenciesDependencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDependenciesDependenciesWithDefaults() *PackageDependenciesDependencies {
	this := PackageDependenciesDependencies{}
	return &this
}

// GetDepType returns the DepType field value if set, zero value otherwise.
func (o *PackageDependenciesDependencies) GetDepType() string {
	if o == nil || o.DepType == nil {
		var ret string
		return ret
	}
	return *o.DepType
}

// GetDepTypeOk returns a tuple with the DepType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDependenciesDependencies) GetDepTypeOk() (*string, bool) {
	if o == nil || o.DepType == nil {
		return nil, false
	}
	return o.DepType, true
}

// HasDepType returns a boolean if a field has been set.
func (o *PackageDependenciesDependencies) HasDepType() bool {
	if o != nil && o.DepType != nil {
		return true
	}

	return false
}

// SetDepType gets a reference to the given string and assigns it to the DepType field.
func (o *PackageDependenciesDependencies) SetDepType(v string) {
	o.DepType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackageDependenciesDependencies) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDependenciesDependencies) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackageDependenciesDependencies) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackageDependenciesDependencies) SetName(v string) {
	o.Name = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *PackageDependenciesDependencies) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDependenciesDependencies) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *PackageDependenciesDependencies) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *PackageDependenciesDependencies) SetOperator(v string) {
	o.Operator = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PackageDependenciesDependencies) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDependenciesDependencies) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PackageDependenciesDependencies) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PackageDependenciesDependencies) SetVersion(v string) {
	o.Version = &v
}

func (o PackageDependenciesDependencies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DepType != nil {
		toSerialize["dep_type"] = o.DepType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePackageDependenciesDependencies struct {
	value *PackageDependenciesDependencies
	isSet bool
}

func (v NullablePackageDependenciesDependencies) Get() *PackageDependenciesDependencies {
	return v.value
}

func (v *NullablePackageDependenciesDependencies) Set(val *PackageDependenciesDependencies) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDependenciesDependencies) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDependenciesDependencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDependenciesDependencies(val *PackageDependenciesDependencies) *NullablePackageDependenciesDependencies {
	return &NullablePackageDependenciesDependencies{value: val, isSet: true}
}

func (v NullablePackageDependenciesDependencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDependenciesDependencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

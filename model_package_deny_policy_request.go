/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.392.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PackageDenyPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageDenyPolicyRequest{}

// PackageDenyPolicyRequest struct for PackageDenyPolicyRequest
type PackageDenyPolicyRequest struct {
	Description NullableString `json:"description,omitempty"`
	// Whether this rule is enabled or disabled.
	Enabled *bool          `json:"enabled,omitempty"`
	Name    NullableString `json:"name,omitempty"`
	// Packages that match this query will trigger this deny rule.
	PackageQueryString string `json:"package_query_string"`
}

// NewPackageDenyPolicyRequest instantiates a new PackageDenyPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDenyPolicyRequest(packageQueryString string) *PackageDenyPolicyRequest {
	this := PackageDenyPolicyRequest{}
	this.PackageQueryString = packageQueryString
	return &this
}

// NewPackageDenyPolicyRequestWithDefaults instantiates a new PackageDenyPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDenyPolicyRequestWithDefaults() *PackageDenyPolicyRequest {
	this := PackageDenyPolicyRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDenyPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDenyPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PackageDenyPolicyRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PackageDenyPolicyRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PackageDenyPolicyRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PackageDenyPolicyRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicyRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PackageDenyPolicyRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageDenyPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageDenyPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PackageDenyPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PackageDenyPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PackageDenyPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PackageDenyPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetPackageQueryString returns the PackageQueryString field value
func (o *PackageDenyPolicyRequest) GetPackageQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageQueryString
}

// GetPackageQueryStringOk returns a tuple with the PackageQueryString field value
// and a boolean to check if the value has been set.
func (o *PackageDenyPolicyRequest) GetPackageQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageQueryString, true
}

// SetPackageQueryString sets field value
func (o *PackageDenyPolicyRequest) SetPackageQueryString(v string) {
	o.PackageQueryString = v
}

func (o PackageDenyPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageDenyPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["package_query_string"] = o.PackageQueryString
	return toSerialize, nil
}

type NullablePackageDenyPolicyRequest struct {
	value *PackageDenyPolicyRequest
	isSet bool
}

func (v NullablePackageDenyPolicyRequest) Get() *PackageDenyPolicyRequest {
	return v.value
}

func (v *NullablePackageDenyPolicyRequest) Set(val *PackageDenyPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDenyPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDenyPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDenyPolicyRequest(val *PackageDenyPolicyRequest) *NullablePackageDenyPolicyRequest {
	return &NullablePackageDenyPolicyRequest{value: val, isSet: true}
}

func (v NullablePackageDenyPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDenyPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

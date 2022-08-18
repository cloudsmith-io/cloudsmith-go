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

// PackagesValidateUploadRpm struct for PackagesValidateUploadRpm
type PackagesValidateUploadRpm struct {
	// The distribution to store the package for.
	Distribution string `json:"distribution"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags *string `json:"tags,omitempty"`
}

// NewPackagesValidateUploadRpm instantiates a new PackagesValidateUploadRpm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesValidateUploadRpm(distribution string, packageFile string) *PackagesValidateUploadRpm {
	this := PackagesValidateUploadRpm{}
	this.Distribution = distribution
	this.PackageFile = packageFile
	return &this
}

// NewPackagesValidateUploadRpmWithDefaults instantiates a new PackagesValidateUploadRpm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesValidateUploadRpmWithDefaults() *PackagesValidateUploadRpm {
	this := PackagesValidateUploadRpm{}
	return &this
}

// GetDistribution returns the Distribution field value
func (o *PackagesValidateUploadRpm) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRpm) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *PackagesValidateUploadRpm) SetDistribution(v string) {
	o.Distribution = v
}

// GetPackageFile returns the PackageFile field value
func (o *PackagesValidateUploadRpm) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRpm) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *PackagesValidateUploadRpm) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *PackagesValidateUploadRpm) GetRepublish() bool {
	if o == nil || o.Republish == nil {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRpm) GetRepublishOk() (*bool, bool) {
	if o == nil || o.Republish == nil {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *PackagesValidateUploadRpm) HasRepublish() bool {
	if o != nil && o.Republish != nil {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *PackagesValidateUploadRpm) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PackagesValidateUploadRpm) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadRpm) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackagesValidateUploadRpm) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PackagesValidateUploadRpm) SetTags(v string) {
	o.Tags = &v
}

func (o PackagesValidateUploadRpm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distribution"] = o.Distribution
	}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if o.Republish != nil {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesValidateUploadRpm struct {
	value *PackagesValidateUploadRpm
	isSet bool
}

func (v NullablePackagesValidateUploadRpm) Get() *PackagesValidateUploadRpm {
	return v.value
}

func (v *NullablePackagesValidateUploadRpm) Set(val *PackagesValidateUploadRpm) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesValidateUploadRpm) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesValidateUploadRpm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesValidateUploadRpm(val *PackagesValidateUploadRpm) *NullablePackagesValidateUploadRpm {
	return &NullablePackagesValidateUploadRpm{value: val, isSet: true}
}

func (v NullablePackagesValidateUploadRpm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesValidateUploadRpm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

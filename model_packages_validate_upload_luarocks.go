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

// PackagesValidateUploadLuarocks struct for PackagesValidateUploadLuarocks
type PackagesValidateUploadLuarocks struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags *string `json:"tags,omitempty"`
}

// NewPackagesValidateUploadLuarocks instantiates a new PackagesValidateUploadLuarocks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesValidateUploadLuarocks(packageFile string) *PackagesValidateUploadLuarocks {
	this := PackagesValidateUploadLuarocks{}
	this.PackageFile = packageFile
	return &this
}

// NewPackagesValidateUploadLuarocksWithDefaults instantiates a new PackagesValidateUploadLuarocks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesValidateUploadLuarocksWithDefaults() *PackagesValidateUploadLuarocks {
	this := PackagesValidateUploadLuarocks{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *PackagesValidateUploadLuarocks) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadLuarocks) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *PackagesValidateUploadLuarocks) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *PackagesValidateUploadLuarocks) GetRepublish() bool {
	if o == nil || o.Republish == nil {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadLuarocks) GetRepublishOk() (*bool, bool) {
	if o == nil || o.Republish == nil {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *PackagesValidateUploadLuarocks) HasRepublish() bool {
	if o != nil && o.Republish != nil {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *PackagesValidateUploadLuarocks) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PackagesValidateUploadLuarocks) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesValidateUploadLuarocks) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackagesValidateUploadLuarocks) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PackagesValidateUploadLuarocks) SetTags(v string) {
	o.Tags = &v
}

func (o PackagesValidateUploadLuarocks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullablePackagesValidateUploadLuarocks struct {
	value *PackagesValidateUploadLuarocks
	isSet bool
}

func (v NullablePackagesValidateUploadLuarocks) Get() *PackagesValidateUploadLuarocks {
	return v.value
}

func (v *NullablePackagesValidateUploadLuarocks) Set(val *PackagesValidateUploadLuarocks) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesValidateUploadLuarocks) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesValidateUploadLuarocks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesValidateUploadLuarocks(val *PackagesValidateUploadLuarocks) *NullablePackagesValidateUploadLuarocks {
	return &NullablePackagesValidateUploadLuarocks{value: val, isSet: true}
}

func (v NullablePackagesValidateUploadLuarocks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesValidateUploadLuarocks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

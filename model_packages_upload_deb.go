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

// PackagesUploadDeb struct for PackagesUploadDeb
type PackagesUploadDeb struct {
	// The changes archive containing the changes made to the source and debian packaging files
	ChangesFile *string `json:"changes_file,omitempty"`
	// The distribution to store the package for.
	Distribution string `json:"distribution"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// The sources archive containing the source code for the binary
	SourcesFile *string `json:"sources_file,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags *string `json:"tags,omitempty"`
}

// NewPackagesUploadDeb instantiates a new PackagesUploadDeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesUploadDeb(distribution string, packageFile string) *PackagesUploadDeb {
	this := PackagesUploadDeb{}
	this.Distribution = distribution
	this.PackageFile = packageFile
	return &this
}

// NewPackagesUploadDebWithDefaults instantiates a new PackagesUploadDeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesUploadDebWithDefaults() *PackagesUploadDeb {
	this := PackagesUploadDeb{}
	return &this
}

// GetChangesFile returns the ChangesFile field value if set, zero value otherwise.
func (o *PackagesUploadDeb) GetChangesFile() string {
	if o == nil || o.ChangesFile == nil {
		var ret string
		return ret
	}
	return *o.ChangesFile
}

// GetChangesFileOk returns a tuple with the ChangesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesUploadDeb) GetChangesFileOk() (*string, bool) {
	if o == nil || o.ChangesFile == nil {
		return nil, false
	}
	return o.ChangesFile, true
}

// HasChangesFile returns a boolean if a field has been set.
func (o *PackagesUploadDeb) HasChangesFile() bool {
	if o != nil && o.ChangesFile != nil {
		return true
	}

	return false
}

// SetChangesFile gets a reference to the given string and assigns it to the ChangesFile field.
func (o *PackagesUploadDeb) SetChangesFile(v string) {
	o.ChangesFile = &v
}

// GetDistribution returns the Distribution field value
func (o *PackagesUploadDeb) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *PackagesUploadDeb) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *PackagesUploadDeb) SetDistribution(v string) {
	o.Distribution = v
}

// GetPackageFile returns the PackageFile field value
func (o *PackagesUploadDeb) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *PackagesUploadDeb) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *PackagesUploadDeb) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *PackagesUploadDeb) GetRepublish() bool {
	if o == nil || o.Republish == nil {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesUploadDeb) GetRepublishOk() (*bool, bool) {
	if o == nil || o.Republish == nil {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *PackagesUploadDeb) HasRepublish() bool {
	if o != nil && o.Republish != nil {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *PackagesUploadDeb) SetRepublish(v bool) {
	o.Republish = &v
}

// GetSourcesFile returns the SourcesFile field value if set, zero value otherwise.
func (o *PackagesUploadDeb) GetSourcesFile() string {
	if o == nil || o.SourcesFile == nil {
		var ret string
		return ret
	}
	return *o.SourcesFile
}

// GetSourcesFileOk returns a tuple with the SourcesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesUploadDeb) GetSourcesFileOk() (*string, bool) {
	if o == nil || o.SourcesFile == nil {
		return nil, false
	}
	return o.SourcesFile, true
}

// HasSourcesFile returns a boolean if a field has been set.
func (o *PackagesUploadDeb) HasSourcesFile() bool {
	if o != nil && o.SourcesFile != nil {
		return true
	}

	return false
}

// SetSourcesFile gets a reference to the given string and assigns it to the SourcesFile field.
func (o *PackagesUploadDeb) SetSourcesFile(v string) {
	o.SourcesFile = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PackagesUploadDeb) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesUploadDeb) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PackagesUploadDeb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PackagesUploadDeb) SetTags(v string) {
	o.Tags = &v
}

func (o PackagesUploadDeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangesFile != nil {
		toSerialize["changes_file"] = o.ChangesFile
	}
	if true {
		toSerialize["distribution"] = o.Distribution
	}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if o.Republish != nil {
		toSerialize["republish"] = o.Republish
	}
	if o.SourcesFile != nil {
		toSerialize["sources_file"] = o.SourcesFile
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesUploadDeb struct {
	value *PackagesUploadDeb
	isSet bool
}

func (v NullablePackagesUploadDeb) Get() *PackagesUploadDeb {
	return v.value
}

func (v *NullablePackagesUploadDeb) Set(val *PackagesUploadDeb) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesUploadDeb) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesUploadDeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesUploadDeb(val *PackagesUploadDeb) *NullablePackagesUploadDeb {
	return &NullablePackagesUploadDeb{value: val, isSet: true}
}

func (v NullablePackagesUploadDeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesUploadDeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

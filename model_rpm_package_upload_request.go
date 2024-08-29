/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.498.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RpmPackageUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmPackageUploadRequest{}

// RpmPackageUploadRequest struct for RpmPackageUploadRequest
type RpmPackageUploadRequest struct {
	// The distribution to store the package for.
	Distribution string `json:"distribution"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

// NewRpmPackageUploadRequest instantiates a new RpmPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmPackageUploadRequest(distribution string, packageFile string) *RpmPackageUploadRequest {
	this := RpmPackageUploadRequest{}
	this.Distribution = distribution
	this.PackageFile = packageFile
	return &this
}

// NewRpmPackageUploadRequestWithDefaults instantiates a new RpmPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmPackageUploadRequestWithDefaults() *RpmPackageUploadRequest {
	this := RpmPackageUploadRequest{}
	return &this
}

// GetDistribution returns the Distribution field value
func (o *RpmPackageUploadRequest) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *RpmPackageUploadRequest) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *RpmPackageUploadRequest) SetDistribution(v string) {
	o.Distribution = v
}

// GetPackageFile returns the PackageFile field value
func (o *RpmPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *RpmPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *RpmPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *RpmPackageUploadRequest) GetRepublish() bool {
	if o == nil || IsNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *RpmPackageUploadRequest) HasRepublish() bool {
	if o != nil && !IsNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *RpmPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmPackageUploadRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *RpmPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *RpmPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *RpmPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *RpmPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o RpmPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmPackageUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["distribution"] = o.Distribution
	toSerialize["package_file"] = o.PackageFile
	if !IsNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return toSerialize, nil
}

type NullableRpmPackageUploadRequest struct {
	value *RpmPackageUploadRequest
	isSet bool
}

func (v NullableRpmPackageUploadRequest) Get() *RpmPackageUploadRequest {
	return v.value
}

func (v *NullableRpmPackageUploadRequest) Set(val *RpmPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmPackageUploadRequest(val *RpmPackageUploadRequest) *NullableRpmPackageUploadRequest {
	return &NullableRpmPackageUploadRequest{value: val, isSet: true}
}

func (v NullableRpmPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

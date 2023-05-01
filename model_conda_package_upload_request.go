/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.250.8
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// CondaPackageUploadRequest struct for CondaPackageUploadRequest
type CondaPackageUploadRequest struct {
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
}

// NewCondaPackageUploadRequest instantiates a new CondaPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondaPackageUploadRequest(packageFile string) *CondaPackageUploadRequest {
	this := CondaPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewCondaPackageUploadRequestWithDefaults instantiates a new CondaPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondaPackageUploadRequestWithDefaults() *CondaPackageUploadRequest {
	this := CondaPackageUploadRequest{}
	return &this
}

// GetPackageFile returns the PackageFile field value
func (o *CondaPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *CondaPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *CondaPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *CondaPackageUploadRequest) GetRepublish() bool {
	if o == nil || isNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondaPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || isNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *CondaPackageUploadRequest) HasRepublish() bool {
	if o != nil && !isNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *CondaPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondaPackageUploadRequest) GetTags() string {
	if o == nil || isNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondaPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *CondaPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *CondaPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *CondaPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *CondaPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o CondaPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if !isNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCondaPackageUploadRequest struct {
	value *CondaPackageUploadRequest
	isSet bool
}

func (v NullableCondaPackageUploadRequest) Get() *CondaPackageUploadRequest {
	return v.value
}

func (v *NullableCondaPackageUploadRequest) Set(val *CondaPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCondaPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCondaPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondaPackageUploadRequest(val *CondaPackageUploadRequest) *NullableCondaPackageUploadRequest {
	return &NullableCondaPackageUploadRequest{value: val, isSet: true}
}

func (v NullableCondaPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondaPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

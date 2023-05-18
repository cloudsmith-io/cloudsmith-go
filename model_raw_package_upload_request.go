/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.266.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// RawPackageUploadRequest struct for RawPackageUploadRequest
type RawPackageUploadRequest struct {
	// A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here.
	ContentType NullableString `json:"content_type,omitempty"`
	// A textual description of this package.
	Description NullableString `json:"description,omitempty"`
	// The name of this package.
	Name NullableString `json:"name,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// A one-liner synopsis of this package.
	Summary NullableString `json:"summary,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
	// The raw version for this package.
	Version NullableString `json:"version,omitempty"`
}

// NewRawPackageUploadRequest instantiates a new RawPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawPackageUploadRequest(packageFile string) *RawPackageUploadRequest {
	this := RawPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewRawPackageUploadRequestWithDefaults instantiates a new RawPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawPackageUploadRequestWithDefaults() *RawPackageUploadRequest {
	this := RawPackageUploadRequest{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawPackageUploadRequest) GetContentType() string {
	if o == nil || isNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawPackageUploadRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *RawPackageUploadRequest) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *RawPackageUploadRequest) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *RawPackageUploadRequest) UnsetContentType() {
	o.ContentType.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawPackageUploadRequest) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawPackageUploadRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RawPackageUploadRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RawPackageUploadRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RawPackageUploadRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawPackageUploadRequest) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawPackageUploadRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RawPackageUploadRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *RawPackageUploadRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RawPackageUploadRequest) UnsetName() {
	o.Name.Unset()
}

// GetPackageFile returns the PackageFile field value
func (o *RawPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *RawPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *RawPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *RawPackageUploadRequest) GetRepublish() bool {
	if o == nil || isNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || isNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasRepublish() bool {
	if o != nil && !isNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *RawPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawPackageUploadRequest) GetSummary() string {
	if o == nil || isNil(o.Summary.Get()) {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawPackageUploadRequest) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasSummary() bool {
	if o != nil && o.Summary.IsSet() {
		return true
	}

	return false
}

// SetSummary gets a reference to the given NullableString and assigns it to the Summary field.
func (o *RawPackageUploadRequest) SetSummary(v string) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil
func (o *RawPackageUploadRequest) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil
func (o *RawPackageUploadRequest) UnsetSummary() {
	o.Summary.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawPackageUploadRequest) GetTags() string {
	if o == nil || isNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *RawPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *RawPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *RawPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawPackageUploadRequest) GetVersion() string {
	if o == nil || isNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawPackageUploadRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *RawPackageUploadRequest) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *RawPackageUploadRequest) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *RawPackageUploadRequest) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *RawPackageUploadRequest) UnsetVersion() {
	o.Version.Unset()
}

func (o RawPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if !isNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRawPackageUploadRequest struct {
	value *RawPackageUploadRequest
	isSet bool
}

func (v NullableRawPackageUploadRequest) Get() *RawPackageUploadRequest {
	return v.value
}

func (v *NullableRawPackageUploadRequest) Set(val *RawPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRawPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRawPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawPackageUploadRequest(val *RawPackageUploadRequest) *NullableRawPackageUploadRequest {
	return &NullableRawPackageUploadRequest{value: val, isSet: true}
}

func (v NullableRawPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

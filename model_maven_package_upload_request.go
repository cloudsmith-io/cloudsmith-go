/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.202.5
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// MavenPackageUploadRequest struct for MavenPackageUploadRequest
type MavenPackageUploadRequest struct {
	// The ID of the artifact.
	ArtifactId NullableString `json:"artifact_id,omitempty"`
	// Artifact's group ID.
	GroupId NullableString `json:"group_id,omitempty"`
	// Adds bundled Java documentation to the Maven package
	JavadocFile NullableString `json:"javadoc_file,omitempty"`
	// The primary file for the package.
	PackageFile string `json:"package_file"`
	// Artifact's Maven packaging type.
	Packaging NullableString `json:"packaging,omitempty"`
	// The POM file is an XML file containing the Maven coordinates.
	PomFile NullableString `json:"pom_file,omitempty"`
	// If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate.
	Republish *bool `json:"republish,omitempty"`
	// Adds bundled Java source code to the Maven package.
	SourcesFile NullableString `json:"sources_file,omitempty"`
	// A comma-separated values list of tags to add to the package.
	Tags NullableString `json:"tags,omitempty"`
	// Adds bundled Java tests to the Maven package.
	TestsFile NullableString `json:"tests_file,omitempty"`
	// The raw version for this package.
	Version NullableString `json:"version,omitempty"`
}

// NewMavenPackageUploadRequest instantiates a new MavenPackageUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMavenPackageUploadRequest(packageFile string) *MavenPackageUploadRequest {
	this := MavenPackageUploadRequest{}
	this.PackageFile = packageFile
	return &this
}

// NewMavenPackageUploadRequestWithDefaults instantiates a new MavenPackageUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenPackageUploadRequestWithDefaults() *MavenPackageUploadRequest {
	this := MavenPackageUploadRequest{}
	return &this
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetArtifactId() string {
	if o == nil || isNil(o.ArtifactId.Get()) {
		var ret string
		return ret
	}
	return *o.ArtifactId.Get()
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetArtifactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArtifactId.Get(), o.ArtifactId.IsSet()
}

// HasArtifactId returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasArtifactId() bool {
	if o != nil && o.ArtifactId.IsSet() {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given NullableString and assigns it to the ArtifactId field.
func (o *MavenPackageUploadRequest) SetArtifactId(v string) {
	o.ArtifactId.Set(&v)
}

// SetArtifactIdNil sets the value for ArtifactId to be an explicit nil
func (o *MavenPackageUploadRequest) SetArtifactIdNil() {
	o.ArtifactId.Set(nil)
}

// UnsetArtifactId ensures that no value is present for ArtifactId, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetArtifactId() {
	o.ArtifactId.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetGroupId() string {
	if o == nil || isNil(o.GroupId.Get()) {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *MavenPackageUploadRequest) SetGroupId(v string) {
	o.GroupId.Set(&v)
}

// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *MavenPackageUploadRequest) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetJavadocFile returns the JavadocFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetJavadocFile() string {
	if o == nil || isNil(o.JavadocFile.Get()) {
		var ret string
		return ret
	}
	return *o.JavadocFile.Get()
}

// GetJavadocFileOk returns a tuple with the JavadocFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetJavadocFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JavadocFile.Get(), o.JavadocFile.IsSet()
}

// HasJavadocFile returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasJavadocFile() bool {
	if o != nil && o.JavadocFile.IsSet() {
		return true
	}

	return false
}

// SetJavadocFile gets a reference to the given NullableString and assigns it to the JavadocFile field.
func (o *MavenPackageUploadRequest) SetJavadocFile(v string) {
	o.JavadocFile.Set(&v)
}

// SetJavadocFileNil sets the value for JavadocFile to be an explicit nil
func (o *MavenPackageUploadRequest) SetJavadocFileNil() {
	o.JavadocFile.Set(nil)
}

// UnsetJavadocFile ensures that no value is present for JavadocFile, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetJavadocFile() {
	o.JavadocFile.Unset()
}

// GetPackageFile returns the PackageFile field value
func (o *MavenPackageUploadRequest) GetPackageFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageFile
}

// GetPackageFileOk returns a tuple with the PackageFile field value
// and a boolean to check if the value has been set.
func (o *MavenPackageUploadRequest) GetPackageFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageFile, true
}

// SetPackageFile sets field value
func (o *MavenPackageUploadRequest) SetPackageFile(v string) {
	o.PackageFile = v
}

// GetPackaging returns the Packaging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetPackaging() string {
	if o == nil || isNil(o.Packaging.Get()) {
		var ret string
		return ret
	}
	return *o.Packaging.Get()
}

// GetPackagingOk returns a tuple with the Packaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetPackagingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Packaging.Get(), o.Packaging.IsSet()
}

// HasPackaging returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasPackaging() bool {
	if o != nil && o.Packaging.IsSet() {
		return true
	}

	return false
}

// SetPackaging gets a reference to the given NullableString and assigns it to the Packaging field.
func (o *MavenPackageUploadRequest) SetPackaging(v string) {
	o.Packaging.Set(&v)
}

// SetPackagingNil sets the value for Packaging to be an explicit nil
func (o *MavenPackageUploadRequest) SetPackagingNil() {
	o.Packaging.Set(nil)
}

// UnsetPackaging ensures that no value is present for Packaging, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetPackaging() {
	o.Packaging.Unset()
}

// GetPomFile returns the PomFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetPomFile() string {
	if o == nil || isNil(o.PomFile.Get()) {
		var ret string
		return ret
	}
	return *o.PomFile.Get()
}

// GetPomFileOk returns a tuple with the PomFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetPomFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PomFile.Get(), o.PomFile.IsSet()
}

// HasPomFile returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasPomFile() bool {
	if o != nil && o.PomFile.IsSet() {
		return true
	}

	return false
}

// SetPomFile gets a reference to the given NullableString and assigns it to the PomFile field.
func (o *MavenPackageUploadRequest) SetPomFile(v string) {
	o.PomFile.Set(&v)
}

// SetPomFileNil sets the value for PomFile to be an explicit nil
func (o *MavenPackageUploadRequest) SetPomFileNil() {
	o.PomFile.Set(nil)
}

// UnsetPomFile ensures that no value is present for PomFile, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetPomFile() {
	o.PomFile.Unset()
}

// GetRepublish returns the Republish field value if set, zero value otherwise.
func (o *MavenPackageUploadRequest) GetRepublish() bool {
	if o == nil || isNil(o.Republish) {
		var ret bool
		return ret
	}
	return *o.Republish
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenPackageUploadRequest) GetRepublishOk() (*bool, bool) {
	if o == nil || isNil(o.Republish) {
		return nil, false
	}
	return o.Republish, true
}

// HasRepublish returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasRepublish() bool {
	if o != nil && !isNil(o.Republish) {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given bool and assigns it to the Republish field.
func (o *MavenPackageUploadRequest) SetRepublish(v bool) {
	o.Republish = &v
}

// GetSourcesFile returns the SourcesFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetSourcesFile() string {
	if o == nil || isNil(o.SourcesFile.Get()) {
		var ret string
		return ret
	}
	return *o.SourcesFile.Get()
}

// GetSourcesFileOk returns a tuple with the SourcesFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetSourcesFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourcesFile.Get(), o.SourcesFile.IsSet()
}

// HasSourcesFile returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasSourcesFile() bool {
	if o != nil && o.SourcesFile.IsSet() {
		return true
	}

	return false
}

// SetSourcesFile gets a reference to the given NullableString and assigns it to the SourcesFile field.
func (o *MavenPackageUploadRequest) SetSourcesFile(v string) {
	o.SourcesFile.Set(&v)
}

// SetSourcesFileNil sets the value for SourcesFile to be an explicit nil
func (o *MavenPackageUploadRequest) SetSourcesFileNil() {
	o.SourcesFile.Set(nil)
}

// UnsetSourcesFile ensures that no value is present for SourcesFile, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetSourcesFile() {
	o.SourcesFile.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetTags() string {
	if o == nil || isNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *MavenPackageUploadRequest) SetTags(v string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil
func (o *MavenPackageUploadRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetTags() {
	o.Tags.Unset()
}

// GetTestsFile returns the TestsFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetTestsFile() string {
	if o == nil || isNil(o.TestsFile.Get()) {
		var ret string
		return ret
	}
	return *o.TestsFile.Get()
}

// GetTestsFileOk returns a tuple with the TestsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetTestsFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestsFile.Get(), o.TestsFile.IsSet()
}

// HasTestsFile returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasTestsFile() bool {
	if o != nil && o.TestsFile.IsSet() {
		return true
	}

	return false
}

// SetTestsFile gets a reference to the given NullableString and assigns it to the TestsFile field.
func (o *MavenPackageUploadRequest) SetTestsFile(v string) {
	o.TestsFile.Set(&v)
}

// SetTestsFileNil sets the value for TestsFile to be an explicit nil
func (o *MavenPackageUploadRequest) SetTestsFileNil() {
	o.TestsFile.Set(nil)
}

// UnsetTestsFile ensures that no value is present for TestsFile, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetTestsFile() {
	o.TestsFile.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MavenPackageUploadRequest) GetVersion() string {
	if o == nil || isNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MavenPackageUploadRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *MavenPackageUploadRequest) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *MavenPackageUploadRequest) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *MavenPackageUploadRequest) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *MavenPackageUploadRequest) UnsetVersion() {
	o.Version.Unset()
}

func (o MavenPackageUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactId.IsSet() {
		toSerialize["artifact_id"] = o.ArtifactId.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if o.JavadocFile.IsSet() {
		toSerialize["javadoc_file"] = o.JavadocFile.Get()
	}
	if true {
		toSerialize["package_file"] = o.PackageFile
	}
	if o.Packaging.IsSet() {
		toSerialize["packaging"] = o.Packaging.Get()
	}
	if o.PomFile.IsSet() {
		toSerialize["pom_file"] = o.PomFile.Get()
	}
	if !isNil(o.Republish) {
		toSerialize["republish"] = o.Republish
	}
	if o.SourcesFile.IsSet() {
		toSerialize["sources_file"] = o.SourcesFile.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.TestsFile.IsSet() {
		toSerialize["tests_file"] = o.TestsFile.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMavenPackageUploadRequest struct {
	value *MavenPackageUploadRequest
	isSet bool
}

func (v NullableMavenPackageUploadRequest) Get() *MavenPackageUploadRequest {
	return v.value
}

func (v *NullableMavenPackageUploadRequest) Set(val *MavenPackageUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMavenPackageUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMavenPackageUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMavenPackageUploadRequest(val *MavenPackageUploadRequest) *NullableMavenPackageUploadRequest {
	return &NullableMavenPackageUploadRequest{value: val, isSet: true}
}

func (v NullableMavenPackageUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMavenPackageUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

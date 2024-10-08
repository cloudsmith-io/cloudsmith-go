/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.533.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PackageFilePartsUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageFilePartsUpload{}

// PackageFilePartsUpload struct for PackageFilePartsUpload
type PackageFilePartsUpload struct {
	// The identifier for the file to use uploading parts.
	Identifier *string `json:"identifier,omitempty"`
	// The querystring to use for the next-step PUT upload.
	UploadQuerystring NullableString `json:"upload_querystring,omitempty"`
	// The URL to use for the next-step PUT upload
	UploadUrl *string `json:"upload_url,omitempty"`
}

// NewPackageFilePartsUpload instantiates a new PackageFilePartsUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageFilePartsUpload() *PackageFilePartsUpload {
	this := PackageFilePartsUpload{}
	return &this
}

// NewPackageFilePartsUploadWithDefaults instantiates a new PackageFilePartsUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageFilePartsUploadWithDefaults() *PackageFilePartsUpload {
	this := PackageFilePartsUpload{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PackageFilePartsUpload) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFilePartsUpload) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PackageFilePartsUpload) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PackageFilePartsUpload) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetUploadQuerystring returns the UploadQuerystring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageFilePartsUpload) GetUploadQuerystring() string {
	if o == nil || IsNil(o.UploadQuerystring.Get()) {
		var ret string
		return ret
	}
	return *o.UploadQuerystring.Get()
}

// GetUploadQuerystringOk returns a tuple with the UploadQuerystring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageFilePartsUpload) GetUploadQuerystringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadQuerystring.Get(), o.UploadQuerystring.IsSet()
}

// HasUploadQuerystring returns a boolean if a field has been set.
func (o *PackageFilePartsUpload) HasUploadQuerystring() bool {
	if o != nil && o.UploadQuerystring.IsSet() {
		return true
	}

	return false
}

// SetUploadQuerystring gets a reference to the given NullableString and assigns it to the UploadQuerystring field.
func (o *PackageFilePartsUpload) SetUploadQuerystring(v string) {
	o.UploadQuerystring.Set(&v)
}

// SetUploadQuerystringNil sets the value for UploadQuerystring to be an explicit nil
func (o *PackageFilePartsUpload) SetUploadQuerystringNil() {
	o.UploadQuerystring.Set(nil)
}

// UnsetUploadQuerystring ensures that no value is present for UploadQuerystring, not even an explicit nil
func (o *PackageFilePartsUpload) UnsetUploadQuerystring() {
	o.UploadQuerystring.Unset()
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *PackageFilePartsUpload) GetUploadUrl() string {
	if o == nil || IsNil(o.UploadUrl) {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageFilePartsUpload) GetUploadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UploadUrl) {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *PackageFilePartsUpload) HasUploadUrl() bool {
	if o != nil && !IsNil(o.UploadUrl) {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *PackageFilePartsUpload) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

func (o PackageFilePartsUpload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageFilePartsUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if o.UploadQuerystring.IsSet() {
		toSerialize["upload_querystring"] = o.UploadQuerystring.Get()
	}
	if !IsNil(o.UploadUrl) {
		toSerialize["upload_url"] = o.UploadUrl
	}
	return toSerialize, nil
}

type NullablePackageFilePartsUpload struct {
	value *PackageFilePartsUpload
	isSet bool
}

func (v NullablePackageFilePartsUpload) Get() *PackageFilePartsUpload {
	return v.value
}

func (v *NullablePackageFilePartsUpload) Set(val *PackageFilePartsUpload) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageFilePartsUpload) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageFilePartsUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageFilePartsUpload(val *PackageFilePartsUpload) *NullablePackageFilePartsUpload {
	return &NullablePackageFilePartsUpload{value: val, isSet: true}
}

func (v NullablePackageFilePartsUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageFilePartsUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

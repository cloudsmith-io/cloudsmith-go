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

// FilesValidate struct for FilesValidate
type FilesValidate struct {
	// Filename for the package file upload.
	Filename string `json:"filename"`
	// MD5 checksum for a POST-based package file upload.
	Md5Checksum *string `json:"md5_checksum,omitempty"`
	// The method to use for package file upload.
	Method *string `json:"method,omitempty"`
	// SHA256 checksum for a PUT-based package file upload.
	Sha256Checksum *string `json:"sha256_checksum,omitempty"`
}

// NewFilesValidate instantiates a new FilesValidate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesValidate(filename string) *FilesValidate {
	this := FilesValidate{}
	this.Filename = filename
	return &this
}

// NewFilesValidateWithDefaults instantiates a new FilesValidate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesValidateWithDefaults() *FilesValidate {
	this := FilesValidate{}
	return &this
}

// GetFilename returns the Filename field value
func (o *FilesValidate) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *FilesValidate) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *FilesValidate) SetFilename(v string) {
	o.Filename = v
}

// GetMd5Checksum returns the Md5Checksum field value if set, zero value otherwise.
func (o *FilesValidate) GetMd5Checksum() string {
	if o == nil || o.Md5Checksum == nil {
		var ret string
		return ret
	}
	return *o.Md5Checksum
}

// GetMd5ChecksumOk returns a tuple with the Md5Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesValidate) GetMd5ChecksumOk() (*string, bool) {
	if o == nil || o.Md5Checksum == nil {
		return nil, false
	}
	return o.Md5Checksum, true
}

// HasMd5Checksum returns a boolean if a field has been set.
func (o *FilesValidate) HasMd5Checksum() bool {
	if o != nil && o.Md5Checksum != nil {
		return true
	}

	return false
}

// SetMd5Checksum gets a reference to the given string and assigns it to the Md5Checksum field.
func (o *FilesValidate) SetMd5Checksum(v string) {
	o.Md5Checksum = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *FilesValidate) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesValidate) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *FilesValidate) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *FilesValidate) SetMethod(v string) {
	o.Method = &v
}

// GetSha256Checksum returns the Sha256Checksum field value if set, zero value otherwise.
func (o *FilesValidate) GetSha256Checksum() string {
	if o == nil || o.Sha256Checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha256Checksum
}

// GetSha256ChecksumOk returns a tuple with the Sha256Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesValidate) GetSha256ChecksumOk() (*string, bool) {
	if o == nil || o.Sha256Checksum == nil {
		return nil, false
	}
	return o.Sha256Checksum, true
}

// HasSha256Checksum returns a boolean if a field has been set.
func (o *FilesValidate) HasSha256Checksum() bool {
	if o != nil && o.Sha256Checksum != nil {
		return true
	}

	return false
}

// SetSha256Checksum gets a reference to the given string and assigns it to the Sha256Checksum field.
func (o *FilesValidate) SetSha256Checksum(v string) {
	o.Sha256Checksum = &v
}

func (o FilesValidate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filename"] = o.Filename
	}
	if o.Md5Checksum != nil {
		toSerialize["md5_checksum"] = o.Md5Checksum
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Sha256Checksum != nil {
		toSerialize["sha256_checksum"] = o.Sha256Checksum
	}
	return json.Marshal(toSerialize)
}

type NullableFilesValidate struct {
	value *FilesValidate
	isSet bool
}

func (v NullableFilesValidate) Get() *FilesValidate {
	return v.value
}

func (v *NullableFilesValidate) Set(val *FilesValidate) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesValidate) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesValidate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesValidate(val *FilesValidate) *NullableFilesValidate {
	return &NullableFilesValidate{value: val, isSet: true}
}

func (v NullableFilesValidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesValidate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

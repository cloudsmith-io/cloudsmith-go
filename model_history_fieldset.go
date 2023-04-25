/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.247.7
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// HistoryFieldset struct for HistoryFieldset
type HistoryFieldset struct {
	Downloaded  Usage        `json:"downloaded"`
	StorageUsed StorageUsage `json:"storage_used"`
	Uploaded    Usage        `json:"uploaded"`
}

// NewHistoryFieldset instantiates a new HistoryFieldset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryFieldset(downloaded Usage, storageUsed StorageUsage, uploaded Usage) *HistoryFieldset {
	this := HistoryFieldset{}
	this.Downloaded = downloaded
	this.StorageUsed = storageUsed
	this.Uploaded = uploaded
	return &this
}

// NewHistoryFieldsetWithDefaults instantiates a new HistoryFieldset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryFieldsetWithDefaults() *HistoryFieldset {
	this := HistoryFieldset{}
	return &this
}

// GetDownloaded returns the Downloaded field value
func (o *HistoryFieldset) GetDownloaded() Usage {
	if o == nil {
		var ret Usage
		return ret
	}

	return o.Downloaded
}

// GetDownloadedOk returns a tuple with the Downloaded field value
// and a boolean to check if the value has been set.
func (o *HistoryFieldset) GetDownloadedOk() (*Usage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Downloaded, true
}

// SetDownloaded sets field value
func (o *HistoryFieldset) SetDownloaded(v Usage) {
	o.Downloaded = v
}

// GetStorageUsed returns the StorageUsed field value
func (o *HistoryFieldset) GetStorageUsed() StorageUsage {
	if o == nil {
		var ret StorageUsage
		return ret
	}

	return o.StorageUsed
}

// GetStorageUsedOk returns a tuple with the StorageUsed field value
// and a boolean to check if the value has been set.
func (o *HistoryFieldset) GetStorageUsedOk() (*StorageUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageUsed, true
}

// SetStorageUsed sets field value
func (o *HistoryFieldset) SetStorageUsed(v StorageUsage) {
	o.StorageUsed = v
}

// GetUploaded returns the Uploaded field value
func (o *HistoryFieldset) GetUploaded() Usage {
	if o == nil {
		var ret Usage
		return ret
	}

	return o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value
// and a boolean to check if the value has been set.
func (o *HistoryFieldset) GetUploadedOk() (*Usage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uploaded, true
}

// SetUploaded sets field value
func (o *HistoryFieldset) SetUploaded(v Usage) {
	o.Uploaded = v
}

func (o HistoryFieldset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["downloaded"] = o.Downloaded
	}
	if true {
		toSerialize["storage_used"] = o.StorageUsed
	}
	if true {
		toSerialize["uploaded"] = o.Uploaded
	}
	return json.Marshal(toSerialize)
}

type NullableHistoryFieldset struct {
	value *HistoryFieldset
	isSet bool
}

func (v NullableHistoryFieldset) Get() *HistoryFieldset {
	return v.value
}

func (v *NullableHistoryFieldset) Set(val *HistoryFieldset) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryFieldset) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryFieldset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryFieldset(val *HistoryFieldset) *NullableHistoryFieldset {
	return &NullableHistoryFieldset{value: val, isSet: true}
}

func (v NullableHistoryFieldset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryFieldset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

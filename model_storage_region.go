/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.249.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// StorageRegion struct for StorageRegion
type StorageRegion struct {
	// Name of the storage region
	Label string `json:"label"`
	// Slug for the storage region
	Slug string `json:"slug"`
}

// NewStorageRegion instantiates a new StorageRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRegion(label string, slug string) *StorageRegion {
	this := StorageRegion{}
	this.Label = label
	this.Slug = slug
	return &this
}

// NewStorageRegionWithDefaults instantiates a new StorageRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRegionWithDefaults() *StorageRegion {
	this := StorageRegion{}
	return &this
}

// GetLabel returns the Label field value
func (o *StorageRegion) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *StorageRegion) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *StorageRegion) SetLabel(v string) {
	o.Label = v
}

// GetSlug returns the Slug field value
func (o *StorageRegion) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *StorageRegion) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *StorageRegion) SetSlug(v string) {
	o.Slug = v
}

func (o StorageRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullableStorageRegion struct {
	value *StorageRegion
	isSet bool
}

func (v NullableStorageRegion) Get() *StorageRegion {
	return v.value
}

func (v *NullableStorageRegion) Set(val *StorageRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRegion(val *StorageRegion) *NullableStorageRegion {
	return &NullableStorageRegion{value: val, isSet: true}
}

func (v NullableStorageRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

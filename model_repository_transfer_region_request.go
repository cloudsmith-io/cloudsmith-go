/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.536.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryTransferRegionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryTransferRegionRequest{}

// RepositoryTransferRegionRequest struct for RepositoryTransferRegionRequest
type RepositoryTransferRegionRequest struct {
	// The Cloudsmith region in which package files are stored.
	StorageRegion *string `json:"storage_region,omitempty"`
}

// NewRepositoryTransferRegionRequest instantiates a new RepositoryTransferRegionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryTransferRegionRequest() *RepositoryTransferRegionRequest {
	this := RepositoryTransferRegionRequest{}
	var storageRegion string = "default"
	this.StorageRegion = &storageRegion
	return &this
}

// NewRepositoryTransferRegionRequestWithDefaults instantiates a new RepositoryTransferRegionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryTransferRegionRequestWithDefaults() *RepositoryTransferRegionRequest {
	this := RepositoryTransferRegionRequest{}
	var storageRegion string = "default"
	this.StorageRegion = &storageRegion
	return &this
}

// GetStorageRegion returns the StorageRegion field value if set, zero value otherwise.
func (o *RepositoryTransferRegionRequest) GetStorageRegion() string {
	if o == nil || IsNil(o.StorageRegion) {
		var ret string
		return ret
	}
	return *o.StorageRegion
}

// GetStorageRegionOk returns a tuple with the StorageRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTransferRegionRequest) GetStorageRegionOk() (*string, bool) {
	if o == nil || IsNil(o.StorageRegion) {
		return nil, false
	}
	return o.StorageRegion, true
}

// HasStorageRegion returns a boolean if a field has been set.
func (o *RepositoryTransferRegionRequest) HasStorageRegion() bool {
	if o != nil && !IsNil(o.StorageRegion) {
		return true
	}

	return false
}

// SetStorageRegion gets a reference to the given string and assigns it to the StorageRegion field.
func (o *RepositoryTransferRegionRequest) SetStorageRegion(v string) {
	o.StorageRegion = &v
}

func (o RepositoryTransferRegionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryTransferRegionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageRegion) {
		toSerialize["storage_region"] = o.StorageRegion
	}
	return toSerialize, nil
}

type NullableRepositoryTransferRegionRequest struct {
	value *RepositoryTransferRegionRequest
	isSet bool
}

func (v NullableRepositoryTransferRegionRequest) Get() *RepositoryTransferRegionRequest {
	return v.value
}

func (v *NullableRepositoryTransferRegionRequest) Set(val *RepositoryTransferRegionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryTransferRegionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryTransferRegionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryTransferRegionRequest(val *RepositoryTransferRegionRequest) *NullableRepositoryTransferRegionRequest {
	return &NullableRepositoryTransferRegionRequest{value: val, isSet: true}
}

func (v NullableRepositoryTransferRegionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryTransferRegionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

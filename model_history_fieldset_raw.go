/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.536.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HistoryFieldsetRaw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryFieldsetRaw{}

// HistoryFieldsetRaw struct for HistoryFieldsetRaw
type HistoryFieldsetRaw struct {
	Downloaded  UsageRaw        `json:"downloaded"`
	StorageUsed StorageUsageRaw `json:"storage_used"`
	Uploaded    UsageRaw        `json:"uploaded"`
}

type _HistoryFieldsetRaw HistoryFieldsetRaw

// NewHistoryFieldsetRaw instantiates a new HistoryFieldsetRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryFieldsetRaw(downloaded UsageRaw, storageUsed StorageUsageRaw, uploaded UsageRaw) *HistoryFieldsetRaw {
	this := HistoryFieldsetRaw{}
	this.Downloaded = downloaded
	this.StorageUsed = storageUsed
	this.Uploaded = uploaded
	return &this
}

// NewHistoryFieldsetRawWithDefaults instantiates a new HistoryFieldsetRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryFieldsetRawWithDefaults() *HistoryFieldsetRaw {
	this := HistoryFieldsetRaw{}
	return &this
}

// GetDownloaded returns the Downloaded field value
func (o *HistoryFieldsetRaw) GetDownloaded() UsageRaw {
	if o == nil {
		var ret UsageRaw
		return ret
	}

	return o.Downloaded
}

// GetDownloadedOk returns a tuple with the Downloaded field value
// and a boolean to check if the value has been set.
func (o *HistoryFieldsetRaw) GetDownloadedOk() (*UsageRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Downloaded, true
}

// SetDownloaded sets field value
func (o *HistoryFieldsetRaw) SetDownloaded(v UsageRaw) {
	o.Downloaded = v
}

// GetStorageUsed returns the StorageUsed field value
func (o *HistoryFieldsetRaw) GetStorageUsed() StorageUsageRaw {
	if o == nil {
		var ret StorageUsageRaw
		return ret
	}

	return o.StorageUsed
}

// GetStorageUsedOk returns a tuple with the StorageUsed field value
// and a boolean to check if the value has been set.
func (o *HistoryFieldsetRaw) GetStorageUsedOk() (*StorageUsageRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageUsed, true
}

// SetStorageUsed sets field value
func (o *HistoryFieldsetRaw) SetStorageUsed(v StorageUsageRaw) {
	o.StorageUsed = v
}

// GetUploaded returns the Uploaded field value
func (o *HistoryFieldsetRaw) GetUploaded() UsageRaw {
	if o == nil {
		var ret UsageRaw
		return ret
	}

	return o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value
// and a boolean to check if the value has been set.
func (o *HistoryFieldsetRaw) GetUploadedOk() (*UsageRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uploaded, true
}

// SetUploaded sets field value
func (o *HistoryFieldsetRaw) SetUploaded(v UsageRaw) {
	o.Uploaded = v
}

func (o HistoryFieldsetRaw) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryFieldsetRaw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["downloaded"] = o.Downloaded
	toSerialize["storage_used"] = o.StorageUsed
	toSerialize["uploaded"] = o.Uploaded
	return toSerialize, nil
}

func (o *HistoryFieldsetRaw) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"downloaded",
		"storage_used",
		"uploaded",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHistoryFieldsetRaw := _HistoryFieldsetRaw{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHistoryFieldsetRaw)

	if err != nil {
		return err
	}

	*o = HistoryFieldsetRaw(varHistoryFieldsetRaw)

	return err
}

type NullableHistoryFieldsetRaw struct {
	value *HistoryFieldsetRaw
	isSet bool
}

func (v NullableHistoryFieldsetRaw) Get() *HistoryFieldsetRaw {
	return v.value
}

func (v *NullableHistoryFieldsetRaw) Set(val *HistoryFieldsetRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryFieldsetRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryFieldsetRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryFieldsetRaw(val *HistoryFieldsetRaw) *NullableHistoryFieldsetRaw {
	return &NullableHistoryFieldsetRaw{value: val, isSet: true}
}

func (v NullableHistoryFieldsetRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryFieldsetRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

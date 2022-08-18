/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.121.3
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// ReposGpgKeys struct for ReposGpgKeys
type ReposGpgKeys struct {
	// If selected this is the active key for this repository.
	Active *bool `json:"active,omitempty"`
	//
	Comment *string `json:"comment,omitempty"`
	//
	CreatedAt *string `json:"created_at,omitempty"`
	// If selected this is the default key for this repository.
	Default *bool `json:"default,omitempty"`
	// The long identifier used by GPG for this key.
	Fingerprint *string `json:"fingerprint,omitempty"`
	//
	FingerprintShort *string `json:"fingerprint_short,omitempty"`
	// The public key given to repository users.
	PublicKey *string `json:"public_key,omitempty"`
}

// NewReposGpgKeys instantiates a new ReposGpgKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposGpgKeys() *ReposGpgKeys {
	this := ReposGpgKeys{}
	return &this
}

// NewReposGpgKeysWithDefaults instantiates a new ReposGpgKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposGpgKeysWithDefaults() *ReposGpgKeys {
	this := ReposGpgKeys{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ReposGpgKeys) SetActive(v bool) {
	o.Active = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ReposGpgKeys) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ReposGpgKeys) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *ReposGpgKeys) SetDefault(v bool) {
	o.Default = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *ReposGpgKeys) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFingerprintShort returns the FingerprintShort field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetFingerprintShort() string {
	if o == nil || o.FingerprintShort == nil {
		var ret string
		return ret
	}
	return *o.FingerprintShort
}

// GetFingerprintShortOk returns a tuple with the FingerprintShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetFingerprintShortOk() (*string, bool) {
	if o == nil || o.FingerprintShort == nil {
		return nil, false
	}
	return o.FingerprintShort, true
}

// HasFingerprintShort returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasFingerprintShort() bool {
	if o != nil && o.FingerprintShort != nil {
		return true
	}

	return false
}

// SetFingerprintShort gets a reference to the given string and assigns it to the FingerprintShort field.
func (o *ReposGpgKeys) SetFingerprintShort(v string) {
	o.FingerprintShort = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ReposGpgKeys) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGpgKeys) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ReposGpgKeys) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ReposGpgKeys) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o ReposGpgKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.FingerprintShort != nil {
		toSerialize["fingerprint_short"] = o.FingerprintShort
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableReposGpgKeys struct {
	value *ReposGpgKeys
	isSet bool
}

func (v NullableReposGpgKeys) Get() *ReposGpgKeys {
	return v.value
}

func (v *NullableReposGpgKeys) Set(val *ReposGpgKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableReposGpgKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableReposGpgKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposGpgKeys(val *ReposGpgKeys) *NullableReposGpgKeys {
	return &NullableReposGpgKeys{value: val, isSet: true}
}

func (v NullableReposGpgKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposGpgKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

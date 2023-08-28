/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.297.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// checks if the PackageLicensePolicyEvaluationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageLicensePolicyEvaluationRequest{}

// PackageLicensePolicyEvaluationRequest struct for PackageLicensePolicyEvaluationRequest
type PackageLicensePolicyEvaluationRequest struct {
	CreatedAt       *time.Time           `json:"created_at,omitempty"`
	EvaluationCount *int64               `json:"evaluation_count,omitempty"`
	Policy          *NestedLicensePolicy `json:"policy,omitempty"`
	SlugPerm        *string              `json:"slug_perm,omitempty"`
	Status          *string              `json:"status,omitempty"`
	UpdatedAt       *time.Time           `json:"updated_at,omitempty"`
	ViolationCount  *int64               `json:"violation_count,omitempty"`
}

// NewPackageLicensePolicyEvaluationRequest instantiates a new PackageLicensePolicyEvaluationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageLicensePolicyEvaluationRequest() *PackageLicensePolicyEvaluationRequest {
	this := PackageLicensePolicyEvaluationRequest{}
	return &this
}

// NewPackageLicensePolicyEvaluationRequestWithDefaults instantiates a new PackageLicensePolicyEvaluationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageLicensePolicyEvaluationRequestWithDefaults() *PackageLicensePolicyEvaluationRequest {
	this := PackageLicensePolicyEvaluationRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PackageLicensePolicyEvaluationRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEvaluationCount returns the EvaluationCount field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetEvaluationCount() int64 {
	if o == nil || IsNil(o.EvaluationCount) {
		var ret int64
		return ret
	}
	return *o.EvaluationCount
}

// GetEvaluationCountOk returns a tuple with the EvaluationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetEvaluationCountOk() (*int64, bool) {
	if o == nil || IsNil(o.EvaluationCount) {
		return nil, false
	}
	return o.EvaluationCount, true
}

// HasEvaluationCount returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasEvaluationCount() bool {
	if o != nil && !IsNil(o.EvaluationCount) {
		return true
	}

	return false
}

// SetEvaluationCount gets a reference to the given int64 and assigns it to the EvaluationCount field.
func (o *PackageLicensePolicyEvaluationRequest) SetEvaluationCount(v int64) {
	o.EvaluationCount = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetPolicy() NestedLicensePolicy {
	if o == nil || IsNil(o.Policy) {
		var ret NestedLicensePolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetPolicyOk() (*NestedLicensePolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NestedLicensePolicy and assigns it to the Policy field.
func (o *PackageLicensePolicyEvaluationRequest) SetPolicy(v NestedLicensePolicy) {
	o.Policy = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *PackageLicensePolicyEvaluationRequest) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PackageLicensePolicyEvaluationRequest) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PackageLicensePolicyEvaluationRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetViolationCount returns the ViolationCount field value if set, zero value otherwise.
func (o *PackageLicensePolicyEvaluationRequest) GetViolationCount() int64 {
	if o == nil || IsNil(o.ViolationCount) {
		var ret int64
		return ret
	}
	return *o.ViolationCount
}

// GetViolationCountOk returns a tuple with the ViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyEvaluationRequest) GetViolationCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ViolationCount) {
		return nil, false
	}
	return o.ViolationCount, true
}

// HasViolationCount returns a boolean if a field has been set.
func (o *PackageLicensePolicyEvaluationRequest) HasViolationCount() bool {
	if o != nil && !IsNil(o.ViolationCount) {
		return true
	}

	return false
}

// SetViolationCount gets a reference to the given int64 and assigns it to the ViolationCount field.
func (o *PackageLicensePolicyEvaluationRequest) SetViolationCount(v int64) {
	o.ViolationCount = &v
}

func (o PackageLicensePolicyEvaluationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageLicensePolicyEvaluationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.EvaluationCount) {
		toSerialize["evaluation_count"] = o.EvaluationCount
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ViolationCount) {
		toSerialize["violation_count"] = o.ViolationCount
	}
	return toSerialize, nil
}

type NullablePackageLicensePolicyEvaluationRequest struct {
	value *PackageLicensePolicyEvaluationRequest
	isSet bool
}

func (v NullablePackageLicensePolicyEvaluationRequest) Get() *PackageLicensePolicyEvaluationRequest {
	return v.value
}

func (v *NullablePackageLicensePolicyEvaluationRequest) Set(val *PackageLicensePolicyEvaluationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageLicensePolicyEvaluationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageLicensePolicyEvaluationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageLicensePolicyEvaluationRequest(val *PackageLicensePolicyEvaluationRequest) *NullablePackageLicensePolicyEvaluationRequest {
	return &NullablePackageLicensePolicyEvaluationRequest{value: val, isSet: true}
}

func (v NullablePackageLicensePolicyEvaluationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageLicensePolicyEvaluationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

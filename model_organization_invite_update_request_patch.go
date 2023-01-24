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

// OrganizationInviteUpdateRequestPatch struct for OrganizationInviteUpdateRequestPatch
type OrganizationInviteUpdateRequestPatch struct {
	// The role to be assigned to the invited user.
	Role *string `json:"role,omitempty"`
}

// NewOrganizationInviteUpdateRequestPatch instantiates a new OrganizationInviteUpdateRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInviteUpdateRequestPatch() *OrganizationInviteUpdateRequestPatch {
	this := OrganizationInviteUpdateRequestPatch{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewOrganizationInviteUpdateRequestPatchWithDefaults instantiates a new OrganizationInviteUpdateRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInviteUpdateRequestPatchWithDefaults() *OrganizationInviteUpdateRequestPatch {
	this := OrganizationInviteUpdateRequestPatch{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationInviteUpdateRequestPatch) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteUpdateRequestPatch) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationInviteUpdateRequestPatch) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationInviteUpdateRequestPatch) SetRole(v string) {
	o.Role = &v
}

func (o OrganizationInviteUpdateRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationInviteUpdateRequestPatch struct {
	value *OrganizationInviteUpdateRequestPatch
	isSet bool
}

func (v NullableOrganizationInviteUpdateRequestPatch) Get() *OrganizationInviteUpdateRequestPatch {
	return v.value
}

func (v *NullableOrganizationInviteUpdateRequestPatch) Set(val *OrganizationInviteUpdateRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInviteUpdateRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInviteUpdateRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInviteUpdateRequestPatch(val *OrganizationInviteUpdateRequestPatch) *NullableOrganizationInviteUpdateRequestPatch {
	return &NullableOrganizationInviteUpdateRequestPatch{value: val, isSet: true}
}

func (v NullableOrganizationInviteUpdateRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInviteUpdateRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

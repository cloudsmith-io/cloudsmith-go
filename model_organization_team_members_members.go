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

// OrganizationTeamMembersMembers struct for OrganizationTeamMembersMembers
type OrganizationTeamMembersMembers struct {
	//
	Role *string `json:"role,omitempty"`
	//
	User *string `json:"user,omitempty"`
}

// NewOrganizationTeamMembersMembers instantiates a new OrganizationTeamMembersMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamMembersMembers() *OrganizationTeamMembersMembers {
	this := OrganizationTeamMembersMembers{}
	return &this
}

// NewOrganizationTeamMembersMembersWithDefaults instantiates a new OrganizationTeamMembersMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamMembersMembersWithDefaults() *OrganizationTeamMembersMembers {
	this := OrganizationTeamMembersMembers{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationTeamMembersMembers) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamMembersMembers) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationTeamMembersMembers) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationTeamMembersMembers) SetRole(v string) {
	o.Role = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrganizationTeamMembersMembers) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamMembersMembers) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrganizationTeamMembersMembers) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *OrganizationTeamMembersMembers) SetUser(v string) {
	o.User = &v
}

func (o OrganizationTeamMembersMembers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationTeamMembersMembers struct {
	value *OrganizationTeamMembersMembers
	isSet bool
}

func (v NullableOrganizationTeamMembersMembers) Get() *OrganizationTeamMembersMembers {
	return v.value
}

func (v *NullableOrganizationTeamMembersMembers) Set(val *OrganizationTeamMembersMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamMembersMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamMembersMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamMembersMembers(val *OrganizationTeamMembersMembers) *NullableOrganizationTeamMembersMembers {
	return &NullableOrganizationTeamMembersMembers{value: val, isSet: true}
}

func (v NullableOrganizationTeamMembersMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamMembersMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

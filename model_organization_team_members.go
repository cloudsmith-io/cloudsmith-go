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

// OrganizationTeamMembers struct for OrganizationTeamMembers
type OrganizationTeamMembers struct {
	// The team members
	Members []OrganizationTeamMembership `json:"members"`
}

// NewOrganizationTeamMembers instantiates a new OrganizationTeamMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamMembers(members []OrganizationTeamMembership) *OrganizationTeamMembers {
	this := OrganizationTeamMembers{}
	this.Members = members
	return &this
}

// NewOrganizationTeamMembersWithDefaults instantiates a new OrganizationTeamMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamMembersWithDefaults() *OrganizationTeamMembers {
	this := OrganizationTeamMembers{}
	return &this
}

// GetMembers returns the Members field value
func (o *OrganizationTeamMembers) GetMembers() []OrganizationTeamMembership {
	if o == nil {
		var ret []OrganizationTeamMembership
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamMembers) GetMembersOk() ([]OrganizationTeamMembership, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *OrganizationTeamMembers) SetMembers(v []OrganizationTeamMembership) {
	o.Members = v
}

func (o OrganizationTeamMembers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationTeamMembers struct {
	value *OrganizationTeamMembers
	isSet bool
}

func (v NullableOrganizationTeamMembers) Get() *OrganizationTeamMembers {
	return v.value
}

func (v *NullableOrganizationTeamMembers) Set(val *OrganizationTeamMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamMembers(val *OrganizationTeamMembers) *NullableOrganizationTeamMembers {
	return &NullableOrganizationTeamMembers{value: val, isSet: true}
}

func (v NullableOrganizationTeamMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

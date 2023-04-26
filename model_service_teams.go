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

// ServiceTeams struct for ServiceTeams
type ServiceTeams struct {
	// The team role associated with the service
	Role *string `json:"role,omitempty"`
	// The teams associated with the service
	Slug string `json:"slug"`
}

// NewServiceTeams instantiates a new ServiceTeams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTeams(slug string) *ServiceTeams {
	this := ServiceTeams{}
	var role string = "Manager"
	this.Role = &role
	this.Slug = slug
	return &this
}

// NewServiceTeamsWithDefaults instantiates a new ServiceTeams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTeamsWithDefaults() *ServiceTeams {
	this := ServiceTeams{}
	var role string = "Manager"
	this.Role = &role
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ServiceTeams) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTeams) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ServiceTeams) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ServiceTeams) SetRole(v string) {
	o.Role = &v
}

// GetSlug returns the Slug field value
func (o *ServiceTeams) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *ServiceTeams) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *ServiceTeams) SetSlug(v string) {
	o.Slug = v
}

func (o ServiceTeams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullableServiceTeams struct {
	value *ServiceTeams
	isSet bool
}

func (v NullableServiceTeams) Get() *ServiceTeams {
	return v.value
}

func (v *NullableServiceTeams) Set(val *ServiceTeams) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTeams) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTeams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTeams(val *ServiceTeams) *NullableServiceTeams {
	return &NullableServiceTeams{value: val, isSet: true}
}

func (v NullableServiceTeams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTeams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

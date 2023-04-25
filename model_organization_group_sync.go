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

// OrganizationGroupSync struct for OrganizationGroupSync
type OrganizationGroupSync struct {
	IdpKey   string  `json:"idp_key"`
	IdpValue string  `json:"idp_value"`
	Role     *string `json:"role,omitempty"`
	SlugPerm *string `json:"slug_perm,omitempty"`
	Team     string  `json:"team"`
}

// NewOrganizationGroupSync instantiates a new OrganizationGroupSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationGroupSync(idpKey string, idpValue string, team string) *OrganizationGroupSync {
	this := OrganizationGroupSync{}
	this.IdpKey = idpKey
	this.IdpValue = idpValue
	var role string = "Member"
	this.Role = &role
	this.Team = team
	return &this
}

// NewOrganizationGroupSyncWithDefaults instantiates a new OrganizationGroupSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationGroupSyncWithDefaults() *OrganizationGroupSync {
	this := OrganizationGroupSync{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetIdpKey returns the IdpKey field value
func (o *OrganizationGroupSync) GetIdpKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpKey
}

// GetIdpKeyOk returns a tuple with the IdpKey field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSync) GetIdpKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpKey, true
}

// SetIdpKey sets field value
func (o *OrganizationGroupSync) SetIdpKey(v string) {
	o.IdpKey = v
}

// GetIdpValue returns the IdpValue field value
func (o *OrganizationGroupSync) GetIdpValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpValue
}

// GetIdpValueOk returns a tuple with the IdpValue field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSync) GetIdpValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpValue, true
}

// SetIdpValue sets field value
func (o *OrganizationGroupSync) SetIdpValue(v string) {
	o.IdpValue = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationGroupSync) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSync) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationGroupSync) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationGroupSync) SetRole(v string) {
	o.Role = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *OrganizationGroupSync) GetSlugPerm() string {
	if o == nil || isNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSync) GetSlugPermOk() (*string, bool) {
	if o == nil || isNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *OrganizationGroupSync) HasSlugPerm() bool {
	if o != nil && !isNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *OrganizationGroupSync) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetTeam returns the Team field value
func (o *OrganizationGroupSync) GetTeam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *OrganizationGroupSync) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value
func (o *OrganizationGroupSync) SetTeam(v string) {
	o.Team = v
}

func (o OrganizationGroupSync) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["idp_key"] = o.IdpKey
	}
	if true {
		toSerialize["idp_value"] = o.IdpValue
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if true {
		toSerialize["team"] = o.Team
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationGroupSync struct {
	value *OrganizationGroupSync
	isSet bool
}

func (v NullableOrganizationGroupSync) Get() *OrganizationGroupSync {
	return v.value
}

func (v *NullableOrganizationGroupSync) Set(val *OrganizationGroupSync) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationGroupSync) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationGroupSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationGroupSync(val *OrganizationGroupSync) *NullableOrganizationGroupSync {
	return &NullableOrganizationGroupSync{value: val, isSet: true}
}

func (v NullableOrganizationGroupSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationGroupSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

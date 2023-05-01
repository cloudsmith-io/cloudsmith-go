/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.250.8
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// OrganizationInviteExtend struct for OrganizationInviteExtend
type OrganizationInviteExtend struct {
	// The email of the user to be invited.
	Email      *string        `json:"email,omitempty"`
	ExpiresAt  *time.Time     `json:"expires_at,omitempty"`
	Inviter    *string        `json:"inviter,omitempty"`
	InviterUrl NullableString `json:"inviter_url,omitempty"`
	Org        *string        `json:"org,omitempty"`
	// The role to be assigned to the invited user.
	Role *string `json:"role,omitempty"`
	// The slug_perm of the invite to be extended.
	SlugPerm *string `json:"slug_perm,omitempty"`
	// The slug of the user to be invited.
	User    *string        `json:"user,omitempty"`
	UserUrl NullableString `json:"user_url,omitempty"`
}

// NewOrganizationInviteExtend instantiates a new OrganizationInviteExtend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInviteExtend() *OrganizationInviteExtend {
	this := OrganizationInviteExtend{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewOrganizationInviteExtendWithDefaults instantiates a new OrganizationInviteExtend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInviteExtendWithDefaults() *OrganizationInviteExtend {
	this := OrganizationInviteExtend{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganizationInviteExtend) SetEmail(v string) {
	o.Email = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OrganizationInviteExtend) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetInviter() string {
	if o == nil || isNil(o.Inviter) {
		var ret string
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetInviterOk() (*string, bool) {
	if o == nil || isNil(o.Inviter) {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasInviter() bool {
	if o != nil && !isNil(o.Inviter) {
		return true
	}

	return false
}

// SetInviter gets a reference to the given string and assigns it to the Inviter field.
func (o *OrganizationInviteExtend) SetInviter(v string) {
	o.Inviter = &v
}

// GetInviterUrl returns the InviterUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationInviteExtend) GetInviterUrl() string {
	if o == nil || isNil(o.InviterUrl.Get()) {
		var ret string
		return ret
	}
	return *o.InviterUrl.Get()
}

// GetInviterUrlOk returns a tuple with the InviterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationInviteExtend) GetInviterUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InviterUrl.Get(), o.InviterUrl.IsSet()
}

// HasInviterUrl returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasInviterUrl() bool {
	if o != nil && o.InviterUrl.IsSet() {
		return true
	}

	return false
}

// SetInviterUrl gets a reference to the given NullableString and assigns it to the InviterUrl field.
func (o *OrganizationInviteExtend) SetInviterUrl(v string) {
	o.InviterUrl.Set(&v)
}

// SetInviterUrlNil sets the value for InviterUrl to be an explicit nil
func (o *OrganizationInviteExtend) SetInviterUrlNil() {
	o.InviterUrl.Set(nil)
}

// UnsetInviterUrl ensures that no value is present for InviterUrl, not even an explicit nil
func (o *OrganizationInviteExtend) UnsetInviterUrl() {
	o.InviterUrl.Unset()
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetOrg() string {
	if o == nil || isNil(o.Org) {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetOrgOk() (*string, bool) {
	if o == nil || isNil(o.Org) {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasOrg() bool {
	if o != nil && !isNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *OrganizationInviteExtend) SetOrg(v string) {
	o.Org = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationInviteExtend) SetRole(v string) {
	o.Role = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetSlugPerm() string {
	if o == nil || isNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetSlugPermOk() (*string, bool) {
	if o == nil || isNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasSlugPerm() bool {
	if o != nil && !isNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *OrganizationInviteExtend) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrganizationInviteExtend) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInviteExtend) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *OrganizationInviteExtend) SetUser(v string) {
	o.User = &v
}

// GetUserUrl returns the UserUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationInviteExtend) GetUserUrl() string {
	if o == nil || isNil(o.UserUrl.Get()) {
		var ret string
		return ret
	}
	return *o.UserUrl.Get()
}

// GetUserUrlOk returns a tuple with the UserUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationInviteExtend) GetUserUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserUrl.Get(), o.UserUrl.IsSet()
}

// HasUserUrl returns a boolean if a field has been set.
func (o *OrganizationInviteExtend) HasUserUrl() bool {
	if o != nil && o.UserUrl.IsSet() {
		return true
	}

	return false
}

// SetUserUrl gets a reference to the given NullableString and assigns it to the UserUrl field.
func (o *OrganizationInviteExtend) SetUserUrl(v string) {
	o.UserUrl.Set(&v)
}

// SetUserUrlNil sets the value for UserUrl to be an explicit nil
func (o *OrganizationInviteExtend) SetUserUrlNil() {
	o.UserUrl.Set(nil)
}

// UnsetUserUrl ensures that no value is present for UserUrl, not even an explicit nil
func (o *OrganizationInviteExtend) UnsetUserUrl() {
	o.UserUrl.Unset()
}

func (o OrganizationInviteExtend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !isNil(o.Inviter) {
		toSerialize["inviter"] = o.Inviter
	}
	if o.InviterUrl.IsSet() {
		toSerialize["inviter_url"] = o.InviterUrl.Get()
	}
	if !isNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if o.UserUrl.IsSet() {
		toSerialize["user_url"] = o.UserUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationInviteExtend struct {
	value *OrganizationInviteExtend
	isSet bool
}

func (v NullableOrganizationInviteExtend) Get() *OrganizationInviteExtend {
	return v.value
}

func (v *NullableOrganizationInviteExtend) Set(val *OrganizationInviteExtend) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInviteExtend) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInviteExtend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInviteExtend(val *OrganizationInviteExtend) *NullableOrganizationInviteExtend {
	return &NullableOrganizationInviteExtend{value: val, isSet: true}
}

func (v NullableOrganizationInviteExtend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInviteExtend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

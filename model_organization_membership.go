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

// OrganizationMembership struct for OrganizationMembership
type OrganizationMembership struct {
	Email           *string      `json:"email,omitempty"`
	HasTwoFactor    *bool        `json:"has_two_factor,omitempty"`
	JoinedAt        *time.Time   `json:"joined_at,omitempty"`
	LastLoginAt     NullableTime `json:"last_login_at,omitempty"`
	LastLoginMethod *string      `json:"last_login_method,omitempty"`
	Role            *string      `json:"role,omitempty"`
	User            *string      `json:"user,omitempty"`
	UserId          *string      `json:"user_id,omitempty"`
	UserName        *string      `json:"user_name,omitempty"`
	UserUrl         *string      `json:"user_url,omitempty"`
	Visibility      *string      `json:"visibility,omitempty"`
}

// NewOrganizationMembership instantiates a new OrganizationMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationMembership() *OrganizationMembership {
	this := OrganizationMembership{}
	return &this
}

// NewOrganizationMembershipWithDefaults instantiates a new OrganizationMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationMembershipWithDefaults() *OrganizationMembership {
	this := OrganizationMembership{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganizationMembership) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationMembership) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganizationMembership) SetEmail(v string) {
	o.Email = &v
}

// GetHasTwoFactor returns the HasTwoFactor field value if set, zero value otherwise.
func (o *OrganizationMembership) GetHasTwoFactor() bool {
	if o == nil || isNil(o.HasTwoFactor) {
		var ret bool
		return ret
	}
	return *o.HasTwoFactor
}

// GetHasTwoFactorOk returns a tuple with the HasTwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetHasTwoFactorOk() (*bool, bool) {
	if o == nil || isNil(o.HasTwoFactor) {
		return nil, false
	}
	return o.HasTwoFactor, true
}

// HasHasTwoFactor returns a boolean if a field has been set.
func (o *OrganizationMembership) HasHasTwoFactor() bool {
	if o != nil && !isNil(o.HasTwoFactor) {
		return true
	}

	return false
}

// SetHasTwoFactor gets a reference to the given bool and assigns it to the HasTwoFactor field.
func (o *OrganizationMembership) SetHasTwoFactor(v bool) {
	o.HasTwoFactor = &v
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *OrganizationMembership) GetJoinedAt() time.Time {
	if o == nil || isNil(o.JoinedAt) {
		var ret time.Time
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.JoinedAt) {
		return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *OrganizationMembership) HasJoinedAt() bool {
	if o != nil && !isNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given time.Time and assigns it to the JoinedAt field.
func (o *OrganizationMembership) SetJoinedAt(v time.Time) {
	o.JoinedAt = &v
}

// GetLastLoginAt returns the LastLoginAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationMembership) GetLastLoginAt() time.Time {
	if o == nil || isNil(o.LastLoginAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginAt.Get()
}

// GetLastLoginAtOk returns a tuple with the LastLoginAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationMembership) GetLastLoginAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoginAt.Get(), o.LastLoginAt.IsSet()
}

// HasLastLoginAt returns a boolean if a field has been set.
func (o *OrganizationMembership) HasLastLoginAt() bool {
	if o != nil && o.LastLoginAt.IsSet() {
		return true
	}

	return false
}

// SetLastLoginAt gets a reference to the given NullableTime and assigns it to the LastLoginAt field.
func (o *OrganizationMembership) SetLastLoginAt(v time.Time) {
	o.LastLoginAt.Set(&v)
}

// SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil
func (o *OrganizationMembership) SetLastLoginAtNil() {
	o.LastLoginAt.Set(nil)
}

// UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
func (o *OrganizationMembership) UnsetLastLoginAt() {
	o.LastLoginAt.Unset()
}

// GetLastLoginMethod returns the LastLoginMethod field value if set, zero value otherwise.
func (o *OrganizationMembership) GetLastLoginMethod() string {
	if o == nil || isNil(o.LastLoginMethod) {
		var ret string
		return ret
	}
	return *o.LastLoginMethod
}

// GetLastLoginMethodOk returns a tuple with the LastLoginMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetLastLoginMethodOk() (*string, bool) {
	if o == nil || isNil(o.LastLoginMethod) {
		return nil, false
	}
	return o.LastLoginMethod, true
}

// HasLastLoginMethod returns a boolean if a field has been set.
func (o *OrganizationMembership) HasLastLoginMethod() bool {
	if o != nil && !isNil(o.LastLoginMethod) {
		return true
	}

	return false
}

// SetLastLoginMethod gets a reference to the given string and assigns it to the LastLoginMethod field.
func (o *OrganizationMembership) SetLastLoginMethod(v string) {
	o.LastLoginMethod = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationMembership) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationMembership) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrganizationMembership) SetRole(v string) {
	o.Role = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrganizationMembership) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrganizationMembership) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *OrganizationMembership) SetUser(v string) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OrganizationMembership) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OrganizationMembership) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OrganizationMembership) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *OrganizationMembership) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *OrganizationMembership) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *OrganizationMembership) SetUserName(v string) {
	o.UserName = &v
}

// GetUserUrl returns the UserUrl field value if set, zero value otherwise.
func (o *OrganizationMembership) GetUserUrl() string {
	if o == nil || isNil(o.UserUrl) {
		var ret string
		return ret
	}
	return *o.UserUrl
}

// GetUserUrlOk returns a tuple with the UserUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetUserUrlOk() (*string, bool) {
	if o == nil || isNil(o.UserUrl) {
		return nil, false
	}
	return o.UserUrl, true
}

// HasUserUrl returns a boolean if a field has been set.
func (o *OrganizationMembership) HasUserUrl() bool {
	if o != nil && !isNil(o.UserUrl) {
		return true
	}

	return false
}

// SetUserUrl gets a reference to the given string and assigns it to the UserUrl field.
func (o *OrganizationMembership) SetUserUrl(v string) {
	o.UserUrl = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OrganizationMembership) GetVisibility() string {
	if o == nil || isNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMembership) GetVisibilityOk() (*string, bool) {
	if o == nil || isNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OrganizationMembership) HasVisibility() bool {
	if o != nil && !isNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OrganizationMembership) SetVisibility(v string) {
	o.Visibility = &v
}

func (o OrganizationMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.HasTwoFactor) {
		toSerialize["has_two_factor"] = o.HasTwoFactor
	}
	if !isNil(o.JoinedAt) {
		toSerialize["joined_at"] = o.JoinedAt
	}
	if o.LastLoginAt.IsSet() {
		toSerialize["last_login_at"] = o.LastLoginAt.Get()
	}
	if !isNil(o.LastLoginMethod) {
		toSerialize["last_login_method"] = o.LastLoginMethod
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !isNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !isNil(o.UserUrl) {
		toSerialize["user_url"] = o.UserUrl
	}
	if !isNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationMembership struct {
	value *OrganizationMembership
	isSet bool
}

func (v NullableOrganizationMembership) Get() *OrganizationMembership {
	return v.value
}

func (v *NullableOrganizationMembership) Set(val *OrganizationMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationMembership(val *OrganizationMembership) *NullableOrganizationMembership {
	return &NullableOrganizationMembership{value: val, isSet: true}
}

func (v NullableOrganizationMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

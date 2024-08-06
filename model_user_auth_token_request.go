/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.478.2
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the UserAuthTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAuthTokenRequest{}

// UserAuthTokenRequest struct for UserAuthTokenRequest
type UserAuthTokenRequest struct {
	// Email address to authenticate with
	Email *string `json:"email,omitempty"`
	// Password to authenticate with
	Password *string `json:"password,omitempty"`
}

// NewUserAuthTokenRequest instantiates a new UserAuthTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuthTokenRequest() *UserAuthTokenRequest {
	this := UserAuthTokenRequest{}
	return &this
}

// NewUserAuthTokenRequestWithDefaults instantiates a new UserAuthTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthTokenRequestWithDefaults() *UserAuthTokenRequest {
	this := UserAuthTokenRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserAuthTokenRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthTokenRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserAuthTokenRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserAuthTokenRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserAuthTokenRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthTokenRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserAuthTokenRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserAuthTokenRequest) SetPassword(v string) {
	o.Password = &v
}

func (o UserAuthTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAuthTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUserAuthTokenRequest struct {
	value *UserAuthTokenRequest
	isSet bool
}

func (v NullableUserAuthTokenRequest) Get() *UserAuthTokenRequest {
	return v.value
}

func (v *NullableUserAuthTokenRequest) Set(val *UserAuthTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthTokenRequest(val *UserAuthTokenRequest) *NullableUserAuthTokenRequest {
	return &NullableUserAuthTokenRequest{value: val, isSet: true}
}

func (v NullableUserAuthTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.273.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// UserProfile struct for UserProfile
type UserProfile struct {
	Company   NullableString `json:"company,omitempty"`
	FirstName string         `json:"first_name"`
	JobTitle  NullableString `json:"job_title,omitempty"`
	JoinedAt  *time.Time     `json:"joined_at,omitempty"`
	LastName  string         `json:"last_name"`
	Name      *string        `json:"name,omitempty"`
	Slug      *string        `json:"slug,omitempty"`
	SlugPerm  *string        `json:"slug_perm,omitempty"`
	// Your tagline is a sentence about you. Make it funny. Make it professional. Either way, it's public and it represents who you are.
	Tagline NullableString `json:"tagline,omitempty"`
	Url     *string        `json:"url,omitempty"`
}

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile(firstName string, lastName string) *UserProfile {
	this := UserProfile{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetCompany() string {
	if o == nil || isNil(o.Company.Get()) {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *UserProfile) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *UserProfile) SetCompany(v string) {
	o.Company.Set(&v)
}

// SetCompanyNil sets the value for Company to be an explicit nil
func (o *UserProfile) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *UserProfile) UnsetCompany() {
	o.Company.Unset()
}

// GetFirstName returns the FirstName field value
func (o *UserProfile) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserProfile) SetFirstName(v string) {
	o.FirstName = v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetJobTitle() string {
	if o == nil || isNil(o.JobTitle.Get()) {
		var ret string
		return ret
	}
	return *o.JobTitle.Get()
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetJobTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobTitle.Get(), o.JobTitle.IsSet()
}

// HasJobTitle returns a boolean if a field has been set.
func (o *UserProfile) HasJobTitle() bool {
	if o != nil && o.JobTitle.IsSet() {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given NullableString and assigns it to the JobTitle field.
func (o *UserProfile) SetJobTitle(v string) {
	o.JobTitle.Set(&v)
}

// SetJobTitleNil sets the value for JobTitle to be an explicit nil
func (o *UserProfile) SetJobTitleNil() {
	o.JobTitle.Set(nil)
}

// UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
func (o *UserProfile) UnsetJobTitle() {
	o.JobTitle.Unset()
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *UserProfile) GetJoinedAt() time.Time {
	if o == nil || isNil(o.JoinedAt) {
		var ret time.Time
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.JoinedAt) {
		return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *UserProfile) HasJoinedAt() bool {
	if o != nil && !isNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given time.Time and assigns it to the JoinedAt field.
func (o *UserProfile) SetJoinedAt(v time.Time) {
	o.JoinedAt = &v
}

// GetLastName returns the LastName field value
func (o *UserProfile) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserProfile) SetLastName(v string) {
	o.LastName = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserProfile) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserProfile) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserProfile) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserProfile) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserProfile) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserProfile) SetSlug(v string) {
	o.Slug = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *UserProfile) GetSlugPerm() string {
	if o == nil || isNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetSlugPermOk() (*string, bool) {
	if o == nil || isNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *UserProfile) HasSlugPerm() bool {
	if o != nil && !isNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *UserProfile) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetTagline returns the Tagline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetTagline() string {
	if o == nil || isNil(o.Tagline.Get()) {
		var ret string
		return ret
	}
	return *o.Tagline.Get()
}

// GetTaglineOk returns a tuple with the Tagline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetTaglineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tagline.Get(), o.Tagline.IsSet()
}

// HasTagline returns a boolean if a field has been set.
func (o *UserProfile) HasTagline() bool {
	if o != nil && o.Tagline.IsSet() {
		return true
	}

	return false
}

// SetTagline gets a reference to the given NullableString and assigns it to the Tagline field.
func (o *UserProfile) SetTagline(v string) {
	o.Tagline.Set(&v)
}

// SetTaglineNil sets the value for Tagline to be an explicit nil
func (o *UserProfile) SetTaglineNil() {
	o.Tagline.Set(nil)
}

// UnsetTagline ensures that no value is present for Tagline, not even an explicit nil
func (o *UserProfile) UnsetTagline() {
	o.Tagline.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserProfile) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserProfile) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserProfile) SetUrl(v string) {
	o.Url = &v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if o.JobTitle.IsSet() {
		toSerialize["job_title"] = o.JobTitle.Get()
	}
	if !isNil(o.JoinedAt) {
		toSerialize["joined_at"] = o.JoinedAt
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if o.Tagline.IsSet() {
		toSerialize["tagline"] = o.Tagline.Get()
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.275.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
	"time"
)

// Organization struct for Organization
type Organization struct {
	Country   NullableString `json:"country,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	// The city/town/area your organization is based in.
	Location NullableString `json:"location,omitempty"`
	// A descriptive name for your organization.
	Name     *string `json:"name,omitempty"`
	Slug     *string `json:"slug,omitempty"`
	SlugPerm *string `json:"slug_perm,omitempty"`
	// A short public descriptive for your organization.
	Tagline NullableString `json:"tagline,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Organization) GetCountry() string {
	if o == nil || isNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *Organization) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *Organization) SetCountry(v string) {
	o.Country.Set(&v)
}

// SetCountryNil sets the value for Country to be an explicit nil
func (o *Organization) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *Organization) UnsetCountry() {
	o.Country.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Organization) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Organization) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Organization) GetLocation() string {
	if o == nil || isNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *Organization) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *Organization) SetLocation(v string) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *Organization) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *Organization) UnsetLocation() {
	o.Location.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organization) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organization) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organization) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Organization) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Organization) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Organization) SetSlug(v string) {
	o.Slug = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *Organization) GetSlugPerm() string {
	if o == nil || isNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSlugPermOk() (*string, bool) {
	if o == nil || isNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *Organization) HasSlugPerm() bool {
	if o != nil && !isNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *Organization) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetTagline returns the Tagline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Organization) GetTagline() string {
	if o == nil || isNil(o.Tagline.Get()) {
		var ret string
		return ret
	}
	return *o.Tagline.Get()
}

// GetTaglineOk returns a tuple with the Tagline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetTaglineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tagline.Get(), o.Tagline.IsSet()
}

// HasTagline returns a boolean if a field has been set.
func (o *Organization) HasTagline() bool {
	if o != nil && o.Tagline.IsSet() {
		return true
	}

	return false
}

// SetTagline gets a reference to the given NullableString and assigns it to the Tagline field.
func (o *Organization) SetTagline(v string) {
	o.Tagline.Set(&v)
}

// SetTaglineNil sets the value for Tagline to be an explicit nil
func (o *Organization) SetTaglineNil() {
	o.Tagline.Set(nil)
}

// UnsetTagline ensures that no value is present for Tagline, not even an explicit nil
func (o *Organization) UnsetTagline() {
	o.Tagline.Unset()
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
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
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

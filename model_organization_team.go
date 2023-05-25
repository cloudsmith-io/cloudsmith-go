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
)

// OrganizationTeam struct for OrganizationTeam
type OrganizationTeam struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	Slug        *string `json:"slug,omitempty"`
	SlugPerm    *string `json:"slug_perm,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
}

// NewOrganizationTeam instantiates a new OrganizationTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeam(name string) *OrganizationTeam {
	this := OrganizationTeam{}
	this.Name = name
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// NewOrganizationTeamWithDefaults instantiates a new OrganizationTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamWithDefaults() *OrganizationTeam {
	this := OrganizationTeam{}
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationTeam) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationTeam) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationTeam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *OrganizationTeam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationTeam) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OrganizationTeam) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OrganizationTeam) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OrganizationTeam) SetSlug(v string) {
	o.Slug = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *OrganizationTeam) GetSlugPerm() string {
	if o == nil || isNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetSlugPermOk() (*string, bool) {
	if o == nil || isNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *OrganizationTeam) HasSlugPerm() bool {
	if o != nil && !isNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *OrganizationTeam) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OrganizationTeam) GetVisibility() string {
	if o == nil || isNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetVisibilityOk() (*string, bool) {
	if o == nil || isNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OrganizationTeam) HasVisibility() bool {
	if o != nil && !isNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OrganizationTeam) SetVisibility(v string) {
	o.Visibility = &v
}

func (o OrganizationTeam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !isNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationTeam struct {
	value *OrganizationTeam
	isSet bool
}

func (v NullableOrganizationTeam) Get() *OrganizationTeam {
	return v.value
}

func (v *NullableOrganizationTeam) Set(val *OrganizationTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeam(val *OrganizationTeam) *NullableOrganizationTeam {
	return &NullableOrganizationTeam{value: val, isSet: true}
}

func (v NullableOrganizationTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

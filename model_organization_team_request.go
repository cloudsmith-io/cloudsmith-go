/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.533.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the OrganizationTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTeamRequest{}

// OrganizationTeamRequest struct for OrganizationTeamRequest
type OrganizationTeamRequest struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	Slug        *string `json:"slug,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
}

// NewOrganizationTeamRequest instantiates a new OrganizationTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamRequest(name string) *OrganizationTeamRequest {
	this := OrganizationTeamRequest{}
	this.Name = name
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// NewOrganizationTeamRequestWithDefaults instantiates a new OrganizationTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamRequestWithDefaults() *OrganizationTeamRequest {
	this := OrganizationTeamRequest{}
	var visibility string = "Visible"
	this.Visibility = &visibility
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationTeamRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationTeamRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationTeamRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *OrganizationTeamRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationTeamRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OrganizationTeamRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OrganizationTeamRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OrganizationTeamRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OrganizationTeamRequest) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTeamRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OrganizationTeamRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OrganizationTeamRequest) SetVisibility(v string) {
	o.Visibility = &v
}

func (o OrganizationTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableOrganizationTeamRequest struct {
	value *OrganizationTeamRequest
	isSet bool
}

func (v NullableOrganizationTeamRequest) Get() *OrganizationTeamRequest {
	return v.value
}

func (v *NullableOrganizationTeamRequest) Set(val *OrganizationTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamRequest(val *OrganizationTeamRequest) *NullableOrganizationTeamRequest {
	return &NullableOrganizationTeamRequest{value: val, isSet: true}
}

func (v NullableOrganizationTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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
)

// Namespace struct for Namespace
type Namespace struct {
	Name     *string `json:"name,omitempty"`
	Slug     *string `json:"slug,omitempty"`
	SlugPerm *string `json:"slug_perm,omitempty"`
	TypeName *string `json:"type_name,omitempty"`
}

// NewNamespace instantiates a new Namespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespace() *Namespace {
	this := Namespace{}
	return &this
}

// NewNamespaceWithDefaults instantiates a new Namespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceWithDefaults() *Namespace {
	this := Namespace{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Namespace) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namespace) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Namespace) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Namespace) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Namespace) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namespace) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Namespace) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Namespace) SetSlug(v string) {
	o.Slug = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *Namespace) GetSlugPerm() string {
	if o == nil || isNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namespace) GetSlugPermOk() (*string, bool) {
	if o == nil || isNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *Namespace) HasSlugPerm() bool {
	if o != nil && !isNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *Namespace) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *Namespace) GetTypeName() string {
	if o == nil || isNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namespace) GetTypeNameOk() (*string, bool) {
	if o == nil || isNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *Namespace) HasTypeName() bool {
	if o != nil && !isNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *Namespace) SetTypeName(v string) {
	o.TypeName = &v
}

func (o Namespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !isNil(o.TypeName) {
		toSerialize["type_name"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableNamespace struct {
	value *Namespace
	isSet bool
}

func (v NullableNamespace) Get() *Namespace {
	return v.value
}

func (v *NullableNamespace) Set(val *Namespace) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespace(val *Namespace) *NullableNamespace {
	return &NullableNamespace{value: val, isSet: true}
}

func (v NullableNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

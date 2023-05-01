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

// Distribution The distributions supported by this package format
type Distribution struct {
	Name    string  `json:"name"`
	SelfUrl *string `json:"self_url,omitempty"`
	// The slug identifier for this distribution
	Slug     *string        `json:"slug,omitempty"`
	Variants NullableString `json:"variants,omitempty"`
}

// NewDistribution instantiates a new Distribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistribution(name string) *Distribution {
	this := Distribution{}
	this.Name = name
	return &this
}

// NewDistributionWithDefaults instantiates a new Distribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionWithDefaults() *Distribution {
	this := Distribution{}
	return &this
}

// GetName returns the Name field value
func (o *Distribution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Distribution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Distribution) SetName(v string) {
	o.Name = v
}

// GetSelfUrl returns the SelfUrl field value if set, zero value otherwise.
func (o *Distribution) GetSelfUrl() string {
	if o == nil || isNil(o.SelfUrl) {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetSelfUrlOk() (*string, bool) {
	if o == nil || isNil(o.SelfUrl) {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *Distribution) HasSelfUrl() bool {
	if o != nil && !isNil(o.SelfUrl) {
		return true
	}

	return false
}

// SetSelfUrl gets a reference to the given string and assigns it to the SelfUrl field.
func (o *Distribution) SetSelfUrl(v string) {
	o.SelfUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Distribution) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Distribution) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Distribution) SetSlug(v string) {
	o.Slug = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Distribution) GetVariants() string {
	if o == nil || isNil(o.Variants.Get()) {
		var ret string
		return ret
	}
	return *o.Variants.Get()
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Distribution) GetVariantsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants.Get(), o.Variants.IsSet()
}

// HasVariants returns a boolean if a field has been set.
func (o *Distribution) HasVariants() bool {
	if o != nil && o.Variants.IsSet() {
		return true
	}

	return false
}

// SetVariants gets a reference to the given NullableString and assigns it to the Variants field.
func (o *Distribution) SetVariants(v string) {
	o.Variants.Set(&v)
}

// SetVariantsNil sets the value for Variants to be an explicit nil
func (o *Distribution) SetVariantsNil() {
	o.Variants.Set(nil)
}

// UnsetVariants ensures that no value is present for Variants, not even an explicit nil
func (o *Distribution) UnsetVariants() {
	o.Variants.Unset()
}

func (o Distribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SelfUrl) {
		toSerialize["self_url"] = o.SelfUrl
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if o.Variants.IsSet() {
		toSerialize["variants"] = o.Variants.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDistribution struct {
	value *Distribution
	isSet bool
}

func (v NullableDistribution) Get() *Distribution {
	return v.value
}

func (v *NullableDistribution) Set(val *Distribution) {
	v.value = val
	v.isSet = true
}

func (v NullableDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistribution(val *Distribution) *NullableDistribution {
	return &NullableDistribution{value: val, isSet: true}
}

func (v NullableDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

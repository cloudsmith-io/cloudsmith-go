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
)

// DistributionFull struct for DistributionFull
type DistributionFull struct {
	Format    *string `json:"format,omitempty"`
	FormatUrl *string `json:"format_url,omitempty"`
	Name      string  `json:"name"`
	SelfUrl   *string `json:"self_url,omitempty"`
	// The slug identifier for this distribution
	Slug     *string        `json:"slug,omitempty"`
	Variants NullableString `json:"variants,omitempty"`
	// A list of the versions for this distribution
	Versions []DistributionVersion `json:"versions,omitempty"`
}

// NewDistributionFull instantiates a new DistributionFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionFull(name string) *DistributionFull {
	this := DistributionFull{}
	this.Name = name
	return &this
}

// NewDistributionFullWithDefaults instantiates a new DistributionFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionFullWithDefaults() *DistributionFull {
	this := DistributionFull{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DistributionFull) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionFull) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DistributionFull) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DistributionFull) SetFormat(v string) {
	o.Format = &v
}

// GetFormatUrl returns the FormatUrl field value if set, zero value otherwise.
func (o *DistributionFull) GetFormatUrl() string {
	if o == nil || isNil(o.FormatUrl) {
		var ret string
		return ret
	}
	return *o.FormatUrl
}

// GetFormatUrlOk returns a tuple with the FormatUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionFull) GetFormatUrlOk() (*string, bool) {
	if o == nil || isNil(o.FormatUrl) {
		return nil, false
	}
	return o.FormatUrl, true
}

// HasFormatUrl returns a boolean if a field has been set.
func (o *DistributionFull) HasFormatUrl() bool {
	if o != nil && !isNil(o.FormatUrl) {
		return true
	}

	return false
}

// SetFormatUrl gets a reference to the given string and assigns it to the FormatUrl field.
func (o *DistributionFull) SetFormatUrl(v string) {
	o.FormatUrl = &v
}

// GetName returns the Name field value
func (o *DistributionFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DistributionFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DistributionFull) SetName(v string) {
	o.Name = v
}

// GetSelfUrl returns the SelfUrl field value if set, zero value otherwise.
func (o *DistributionFull) GetSelfUrl() string {
	if o == nil || isNil(o.SelfUrl) {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionFull) GetSelfUrlOk() (*string, bool) {
	if o == nil || isNil(o.SelfUrl) {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *DistributionFull) HasSelfUrl() bool {
	if o != nil && !isNil(o.SelfUrl) {
		return true
	}

	return false
}

// SetSelfUrl gets a reference to the given string and assigns it to the SelfUrl field.
func (o *DistributionFull) SetSelfUrl(v string) {
	o.SelfUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *DistributionFull) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionFull) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *DistributionFull) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *DistributionFull) SetSlug(v string) {
	o.Slug = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DistributionFull) GetVariants() string {
	if o == nil || isNil(o.Variants.Get()) {
		var ret string
		return ret
	}
	return *o.Variants.Get()
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionFull) GetVariantsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants.Get(), o.Variants.IsSet()
}

// HasVariants returns a boolean if a field has been set.
func (o *DistributionFull) HasVariants() bool {
	if o != nil && o.Variants.IsSet() {
		return true
	}

	return false
}

// SetVariants gets a reference to the given NullableString and assigns it to the Variants field.
func (o *DistributionFull) SetVariants(v string) {
	o.Variants.Set(&v)
}

// SetVariantsNil sets the value for Variants to be an explicit nil
func (o *DistributionFull) SetVariantsNil() {
	o.Variants.Set(nil)
}

// UnsetVariants ensures that no value is present for Variants, not even an explicit nil
func (o *DistributionFull) UnsetVariants() {
	o.Variants.Unset()
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *DistributionFull) GetVersions() []DistributionVersion {
	if o == nil || isNil(o.Versions) {
		var ret []DistributionVersion
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionFull) GetVersionsOk() ([]DistributionVersion, bool) {
	if o == nil || isNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *DistributionFull) HasVersions() bool {
	if o != nil && !isNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []DistributionVersion and assigns it to the Versions field.
func (o *DistributionFull) SetVersions(v []DistributionVersion) {
	o.Versions = v
}

func (o DistributionFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.FormatUrl) {
		toSerialize["format_url"] = o.FormatUrl
	}
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
	if !isNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableDistributionFull struct {
	value *DistributionFull
	isSet bool
}

func (v NullableDistributionFull) Get() *DistributionFull {
	return v.value
}

func (v *NullableDistributionFull) Set(val *DistributionFull) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionFull) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionFull(val *DistributionFull) *NullableDistributionFull {
	return &NullableDistributionFull{value: val, isSet: true}
}

func (v NullableDistributionFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

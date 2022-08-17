/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.121.3
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// Distribution struct for Distribution
type Distribution struct {
	//
	Format *string `json:"format,omitempty"`
	//
	FormatUrl *string `json:"format_url,omitempty"`
	//
	Name string `json:"name"`
	//
	SelfUrl *string `json:"self_url,omitempty"`
	// The slug identifier for this distribution
	Slug *string `json:"slug,omitempty"`
	//
	Variants *string `json:"variants,omitempty"`
	// A list of the versions for this distribution
	Versions []DistrosVersions `json:"versions,omitempty"`
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

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *Distribution) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *Distribution) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *Distribution) SetFormat(v string) {
	o.Format = &v
}

// GetFormatUrl returns the FormatUrl field value if set, zero value otherwise.
func (o *Distribution) GetFormatUrl() string {
	if o == nil || o.FormatUrl == nil {
		var ret string
		return ret
	}
	return *o.FormatUrl
}

// GetFormatUrlOk returns a tuple with the FormatUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetFormatUrlOk() (*string, bool) {
	if o == nil || o.FormatUrl == nil {
		return nil, false
	}
	return o.FormatUrl, true
}

// HasFormatUrl returns a boolean if a field has been set.
func (o *Distribution) HasFormatUrl() bool {
	if o != nil && o.FormatUrl != nil {
		return true
	}

	return false
}

// SetFormatUrl gets a reference to the given string and assigns it to the FormatUrl field.
func (o *Distribution) SetFormatUrl(v string) {
	o.FormatUrl = &v
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
	if o == nil || o.SelfUrl == nil {
		var ret string
		return ret
	}
	return *o.SelfUrl
}

// GetSelfUrlOk returns a tuple with the SelfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetSelfUrlOk() (*string, bool) {
	if o == nil || o.SelfUrl == nil {
		return nil, false
	}
	return o.SelfUrl, true
}

// HasSelfUrl returns a boolean if a field has been set.
func (o *Distribution) HasSelfUrl() bool {
	if o != nil && o.SelfUrl != nil {
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
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Distribution) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Distribution) SetSlug(v string) {
	o.Slug = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *Distribution) GetVariants() string {
	if o == nil || o.Variants == nil {
		var ret string
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetVariantsOk() (*string, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *Distribution) HasVariants() bool {
	if o != nil && o.Variants != nil {
		return true
	}

	return false
}

// SetVariants gets a reference to the given string and assigns it to the Variants field.
func (o *Distribution) SetVariants(v string) {
	o.Variants = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *Distribution) GetVersions() []DistrosVersions {
	if o == nil || o.Versions == nil {
		var ret []DistrosVersions
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Distribution) GetVersionsOk() ([]DistrosVersions, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *Distribution) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []DistrosVersions and assigns it to the Versions field.
func (o *Distribution) SetVersions(v []DistrosVersions) {
	o.Versions = v
}

func (o Distribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.FormatUrl != nil {
		toSerialize["format_url"] = o.FormatUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SelfUrl != nil {
		toSerialize["self_url"] = o.SelfUrl
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
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

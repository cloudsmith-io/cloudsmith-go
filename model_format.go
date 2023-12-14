/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.327.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the Format type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Format{}

// Format struct for Format
type Format struct {
	// Description of the package format
	Description string `json:"description"`
	// The distributions supported by this package format
	Distributions []Distribution `json:"distributions,omitempty"`
	// A non-exhaustive list of extensions supported
	Extensions []string `json:"extensions"`
	// Name for the package format
	Name string `json:"name"`
	// If true the package format is a premium-only feature
	Premium bool `json:"premium"`
	// The minimum plan id required for this package format
	PremiumPlanId NullableString `json:"premium_plan_id,omitempty"`
	// The minimum plan name required for this package format
	PremiumPlanName NullableString `json:"premium_plan_name,omitempty"`
	// Slug for the package format
	Slug     string        `json:"slug"`
	Supports FormatSupport `json:"supports"`
}

// NewFormat instantiates a new Format object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormat(description string, extensions []string, name string, premium bool, slug string, supports FormatSupport) *Format {
	this := Format{}
	this.Description = description
	this.Extensions = extensions
	this.Name = name
	this.Premium = premium
	this.Slug = slug
	this.Supports = supports
	return &this
}

// NewFormatWithDefaults instantiates a new Format object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatWithDefaults() *Format {
	this := Format{}
	return &this
}

// GetDescription returns the Description field value
func (o *Format) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Format) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Format) SetDescription(v string) {
	o.Description = v
}

// GetDistributions returns the Distributions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Format) GetDistributions() []Distribution {
	if o == nil {
		var ret []Distribution
		return ret
	}
	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Format) GetDistributionsOk() ([]Distribution, bool) {
	if o == nil || IsNil(o.Distributions) {
		return nil, false
	}
	return o.Distributions, true
}

// HasDistributions returns a boolean if a field has been set.
func (o *Format) HasDistributions() bool {
	if o != nil && IsNil(o.Distributions) {
		return true
	}

	return false
}

// SetDistributions gets a reference to the given []Distribution and assigns it to the Distributions field.
func (o *Format) SetDistributions(v []Distribution) {
	o.Distributions = v
}

// GetExtensions returns the Extensions field value
func (o *Format) GetExtensions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value
// and a boolean to check if the value has been set.
func (o *Format) GetExtensionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extensions, true
}

// SetExtensions sets field value
func (o *Format) SetExtensions(v []string) {
	o.Extensions = v
}

// GetName returns the Name field value
func (o *Format) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Format) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Format) SetName(v string) {
	o.Name = v
}

// GetPremium returns the Premium field value
func (o *Format) GetPremium() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value
// and a boolean to check if the value has been set.
func (o *Format) GetPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Premium, true
}

// SetPremium sets field value
func (o *Format) SetPremium(v bool) {
	o.Premium = v
}

// GetPremiumPlanId returns the PremiumPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Format) GetPremiumPlanId() string {
	if o == nil || IsNil(o.PremiumPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PremiumPlanId.Get()
}

// GetPremiumPlanIdOk returns a tuple with the PremiumPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Format) GetPremiumPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiumPlanId.Get(), o.PremiumPlanId.IsSet()
}

// HasPremiumPlanId returns a boolean if a field has been set.
func (o *Format) HasPremiumPlanId() bool {
	if o != nil && o.PremiumPlanId.IsSet() {
		return true
	}

	return false
}

// SetPremiumPlanId gets a reference to the given NullableString and assigns it to the PremiumPlanId field.
func (o *Format) SetPremiumPlanId(v string) {
	o.PremiumPlanId.Set(&v)
}

// SetPremiumPlanIdNil sets the value for PremiumPlanId to be an explicit nil
func (o *Format) SetPremiumPlanIdNil() {
	o.PremiumPlanId.Set(nil)
}

// UnsetPremiumPlanId ensures that no value is present for PremiumPlanId, not even an explicit nil
func (o *Format) UnsetPremiumPlanId() {
	o.PremiumPlanId.Unset()
}

// GetPremiumPlanName returns the PremiumPlanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Format) GetPremiumPlanName() string {
	if o == nil || IsNil(o.PremiumPlanName.Get()) {
		var ret string
		return ret
	}
	return *o.PremiumPlanName.Get()
}

// GetPremiumPlanNameOk returns a tuple with the PremiumPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Format) GetPremiumPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiumPlanName.Get(), o.PremiumPlanName.IsSet()
}

// HasPremiumPlanName returns a boolean if a field has been set.
func (o *Format) HasPremiumPlanName() bool {
	if o != nil && o.PremiumPlanName.IsSet() {
		return true
	}

	return false
}

// SetPremiumPlanName gets a reference to the given NullableString and assigns it to the PremiumPlanName field.
func (o *Format) SetPremiumPlanName(v string) {
	o.PremiumPlanName.Set(&v)
}

// SetPremiumPlanNameNil sets the value for PremiumPlanName to be an explicit nil
func (o *Format) SetPremiumPlanNameNil() {
	o.PremiumPlanName.Set(nil)
}

// UnsetPremiumPlanName ensures that no value is present for PremiumPlanName, not even an explicit nil
func (o *Format) UnsetPremiumPlanName() {
	o.PremiumPlanName.Unset()
}

// GetSlug returns the Slug field value
func (o *Format) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Format) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Format) SetSlug(v string) {
	o.Slug = v
}

// GetSupports returns the Supports field value
func (o *Format) GetSupports() FormatSupport {
	if o == nil {
		var ret FormatSupport
		return ret
	}

	return o.Supports
}

// GetSupportsOk returns a tuple with the Supports field value
// and a boolean to check if the value has been set.
func (o *Format) GetSupportsOk() (*FormatSupport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supports, true
}

// SetSupports sets field value
func (o *Format) SetSupports(v FormatSupport) {
	o.Supports = v
}

func (o Format) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Format) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if o.Distributions != nil {
		toSerialize["distributions"] = o.Distributions
	}
	toSerialize["extensions"] = o.Extensions
	toSerialize["name"] = o.Name
	toSerialize["premium"] = o.Premium
	if o.PremiumPlanId.IsSet() {
		toSerialize["premium_plan_id"] = o.PremiumPlanId.Get()
	}
	if o.PremiumPlanName.IsSet() {
		toSerialize["premium_plan_name"] = o.PremiumPlanName.Get()
	}
	toSerialize["slug"] = o.Slug
	toSerialize["supports"] = o.Supports
	return toSerialize, nil
}

type NullableFormat struct {
	value *Format
	isSet bool
}

func (v NullableFormat) Get() *Format {
	return v.value
}

func (v *NullableFormat) Set(val *Format) {
	v.value = val
	v.isSet = true
}

func (v NullableFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormat(val *Format) *NullableFormat {
	return &NullableFormat{value: val, isSet: true}
}

func (v NullableFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the ProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderSettings{}

// ProviderSettings struct for ProviderSettings
type ProviderSettings struct {
	// The set of claims that any received tokens from the provider must contain to authenticate as the configured service account.
	Claims map[string]interface{} `json:"claims"`
	// Whether the provider settings should be used for incoming OIDC requests.
	Enabled bool `json:"enabled"`
	// The name of the provider settings are being configured for
	Name string `json:"name"`
	// The URL from the provider that serves as the base for the OpenID configuration. For example, if the OpenID configuration is available at https://token.actions.githubusercontent.com/.well-known/openid-configuration, the provider URL would be https://token.actions.githubusercontent.com/
	ProviderUrl string `json:"provider_url"`
	// The service accounts associated with these provider settings
	ServiceAccounts []string `json:"service_accounts"`
	// The slug of the provider settings
	Slug *string `json:"slug,omitempty"`
	// The unique, immutable identifier of the provider settings.
	SlugPerm *string `json:"slug_perm,omitempty"`
}

// NewProviderSettings instantiates a new ProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderSettings(claims map[string]interface{}, enabled bool, name string, providerUrl string, serviceAccounts []string) *ProviderSettings {
	this := ProviderSettings{}
	this.Claims = claims
	this.Enabled = enabled
	this.Name = name
	this.ProviderUrl = providerUrl
	this.ServiceAccounts = serviceAccounts
	return &this
}

// NewProviderSettingsWithDefaults instantiates a new ProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderSettingsWithDefaults() *ProviderSettings {
	this := ProviderSettings{}
	return &this
}

// GetClaims returns the Claims field value
func (o *ProviderSettings) GetClaims() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetClaimsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Claims, true
}

// SetClaims sets field value
func (o *ProviderSettings) SetClaims(v map[string]interface{}) {
	o.Claims = v
}

// GetEnabled returns the Enabled field value
func (o *ProviderSettings) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProviderSettings) SetEnabled(v bool) {
	o.Enabled = v
}

// GetName returns the Name field value
func (o *ProviderSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProviderSettings) SetName(v string) {
	o.Name = v
}

// GetProviderUrl returns the ProviderUrl field value
func (o *ProviderSettings) GetProviderUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderUrl
}

// GetProviderUrlOk returns a tuple with the ProviderUrl field value
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetProviderUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderUrl, true
}

// SetProviderUrl sets field value
func (o *ProviderSettings) SetProviderUrl(v string) {
	o.ProviderUrl = v
}

// GetServiceAccounts returns the ServiceAccounts field value
func (o *ProviderSettings) GetServiceAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServiceAccounts
}

// GetServiceAccountsOk returns a tuple with the ServiceAccounts field value
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetServiceAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceAccounts, true
}

// SetServiceAccounts sets field value
func (o *ProviderSettings) SetServiceAccounts(v []string) {
	o.ServiceAccounts = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ProviderSettings) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ProviderSettings) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ProviderSettings) SetSlug(v string) {
	o.Slug = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *ProviderSettings) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSettings) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *ProviderSettings) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *ProviderSettings) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

func (o ProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["claims"] = o.Claims
	toSerialize["enabled"] = o.Enabled
	toSerialize["name"] = o.Name
	toSerialize["provider_url"] = o.ProviderUrl
	toSerialize["service_accounts"] = o.ServiceAccounts
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	return toSerialize, nil
}

type NullableProviderSettings struct {
	value *ProviderSettings
	isSet bool
}

func (v NullableProviderSettings) Get() *ProviderSettings {
	return v.value
}

func (v *NullableProviderSettings) Set(val *ProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderSettings(val *ProviderSettings) *NullableProviderSettings {
	return &NullableProviderSettings{value: val, isSet: true}
}

func (v NullableProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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
	"time"
)

// checks if the DebUpstream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebUpstream{}

// DebUpstream struct for DebUpstream
type DebUpstream struct {
	// The authentication mode to use when accessing this upstream.
	AuthMode *string `json:"auth_mode,omitempty"`
	// Secret to provide with requests to upstream.
	AuthSecret NullableString `json:"auth_secret,omitempty"`
	// Username to provide with requests to upstream.
	AuthUsername NullableString `json:"auth_username,omitempty"`
	// The component to fetch from the upstream
	Component *string `json:"component,omitempty"`
	// The datetime the upstream source was created.
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	DisableReason *string    `json:"disable_reason,omitempty"`
	// The distribution version that packages found on this upstream could be associated with.
	DistroVersions []string `json:"distro_versions"`
	// The key for extra header #1 to send to upstream.
	ExtraHeader1 NullableString `json:"extra_header_1,omitempty"`
	// The key for extra header #2 to send to upstream.
	ExtraHeader2 NullableString `json:"extra_header_2,omitempty"`
	// The value for extra header #1 to send to upstream. This is stored as plaintext, and is NOT encrypted.
	ExtraValue1 NullableString `json:"extra_value_1,omitempty"`
	// The value for extra header #2 to send to upstream. This is stored as plaintext, and is NOT encrypted.
	ExtraValue2 NullableString `json:"extra_value_2,omitempty"`
	// A public GPG key to associate with packages found on this upstream. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install.
	GpgKeyInline NullableString `json:"gpg_key_inline,omitempty"`
	// When provided, Cloudsmith will fetch, validate, and associate a public GPG key found at the provided URL. When using the Cloudsmith setup script, this GPG key will be automatically imported on your deployment machines to allow upstream packages to validate and install.
	GpgKeyUrl NullableString `json:"gpg_key_url,omitempty"`
	// The GPG signature verification mode for this upstream.
	GpgVerification *string `json:"gpg_verification,omitempty"`
	// When true, source packages will be available from this upstream.
	IncludeSources *bool `json:"include_sources,omitempty"`
	// Whether or not this upstream is active and ready for requests.
	IsActive *bool `json:"is_active,omitempty"`
	// The mode that this upstream should operate in. Upstream sources can be used to proxy resolved packages, as well as operate in a proxy/cache or cache only mode.
	Mode *string `json:"mode,omitempty"`
	// A descriptive name for this upstream source. A shortened version of this name will be used for tagging cached packages retrieved from this upstream.
	Name string `json:"name"`
	// When true, this upstream source is pending validation.
	PendingValidation *bool `json:"pending_validation,omitempty"`
	// Upstream sources are selected for resolving requests by sequential order (1..n), followed by creation date.
	Priority  *int64     `json:"priority,omitempty"`
	SlugPerm  *string    `json:"slug_perm,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The distribution to fetch from the upstream
	UpstreamDistribution NullableString `json:"upstream_distribution,omitempty"`
	// The URL for this upstream source. This must be a fully qualified URL including any path elements required to reach the root of the repository.
	UpstreamUrl string `json:"upstream_url"`
	// The signature verification status for this upstream.
	VerificationStatus *string `json:"verification_status,omitempty"`
	// If enabled, SSL certificates are verified when requests are made to this upstream. It's recommended to leave this enabled for all public sources to help mitigate Man-In-The-Middle (MITM) attacks. Please note this only applies to HTTPS upstreams.
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

// NewDebUpstream instantiates a new DebUpstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebUpstream(distroVersions []string, name string, upstreamUrl string) *DebUpstream {
	this := DebUpstream{}
	var authMode string = "None"
	this.AuthMode = &authMode
	this.DistroVersions = distroVersions
	var gpgVerification string = "Allow All"
	this.GpgVerification = &gpgVerification
	var mode string = "Proxy Only"
	this.Mode = &mode
	this.Name = name
	this.UpstreamUrl = upstreamUrl
	return &this
}

// NewDebUpstreamWithDefaults instantiates a new DebUpstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebUpstreamWithDefaults() *DebUpstream {
	this := DebUpstream{}
	var authMode string = "None"
	this.AuthMode = &authMode
	var gpgVerification string = "Allow All"
	this.GpgVerification = &gpgVerification
	var mode string = "Proxy Only"
	this.Mode = &mode
	return &this
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *DebUpstream) GetAuthMode() string {
	if o == nil || IsNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetAuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthMode) {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *DebUpstream) HasAuthMode() bool {
	if o != nil && !IsNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *DebUpstream) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetAuthSecret returns the AuthSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetAuthSecret() string {
	if o == nil || IsNil(o.AuthSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AuthSecret.Get()
}

// GetAuthSecretOk returns a tuple with the AuthSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetAuthSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthSecret.Get(), o.AuthSecret.IsSet()
}

// HasAuthSecret returns a boolean if a field has been set.
func (o *DebUpstream) HasAuthSecret() bool {
	if o != nil && o.AuthSecret.IsSet() {
		return true
	}

	return false
}

// SetAuthSecret gets a reference to the given NullableString and assigns it to the AuthSecret field.
func (o *DebUpstream) SetAuthSecret(v string) {
	o.AuthSecret.Set(&v)
}

// SetAuthSecretNil sets the value for AuthSecret to be an explicit nil
func (o *DebUpstream) SetAuthSecretNil() {
	o.AuthSecret.Set(nil)
}

// UnsetAuthSecret ensures that no value is present for AuthSecret, not even an explicit nil
func (o *DebUpstream) UnsetAuthSecret() {
	o.AuthSecret.Unset()
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetAuthUsername() string {
	if o == nil || IsNil(o.AuthUsername.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUsername.Get()
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUsername.Get(), o.AuthUsername.IsSet()
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *DebUpstream) HasAuthUsername() bool {
	if o != nil && o.AuthUsername.IsSet() {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given NullableString and assigns it to the AuthUsername field.
func (o *DebUpstream) SetAuthUsername(v string) {
	o.AuthUsername.Set(&v)
}

// SetAuthUsernameNil sets the value for AuthUsername to be an explicit nil
func (o *DebUpstream) SetAuthUsernameNil() {
	o.AuthUsername.Set(nil)
}

// UnsetAuthUsername ensures that no value is present for AuthUsername, not even an explicit nil
func (o *DebUpstream) UnsetAuthUsername() {
	o.AuthUsername.Unset()
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *DebUpstream) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *DebUpstream) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *DebUpstream) SetComponent(v string) {
	o.Component = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DebUpstream) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DebUpstream) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DebUpstream) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisableReason returns the DisableReason field value if set, zero value otherwise.
func (o *DebUpstream) GetDisableReason() string {
	if o == nil || IsNil(o.DisableReason) {
		var ret string
		return ret
	}
	return *o.DisableReason
}

// GetDisableReasonOk returns a tuple with the DisableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetDisableReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DisableReason) {
		return nil, false
	}
	return o.DisableReason, true
}

// HasDisableReason returns a boolean if a field has been set.
func (o *DebUpstream) HasDisableReason() bool {
	if o != nil && !IsNil(o.DisableReason) {
		return true
	}

	return false
}

// SetDisableReason gets a reference to the given string and assigns it to the DisableReason field.
func (o *DebUpstream) SetDisableReason(v string) {
	o.DisableReason = &v
}

// GetDistroVersions returns the DistroVersions field value
func (o *DebUpstream) GetDistroVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DistroVersions
}

// GetDistroVersionsOk returns a tuple with the DistroVersions field value
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetDistroVersionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistroVersions, true
}

// SetDistroVersions sets field value
func (o *DebUpstream) SetDistroVersions(v []string) {
	o.DistroVersions = v
}

// GetExtraHeader1 returns the ExtraHeader1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetExtraHeader1() string {
	if o == nil || IsNil(o.ExtraHeader1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader1.Get()
}

// GetExtraHeader1Ok returns a tuple with the ExtraHeader1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetExtraHeader1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader1.Get(), o.ExtraHeader1.IsSet()
}

// HasExtraHeader1 returns a boolean if a field has been set.
func (o *DebUpstream) HasExtraHeader1() bool {
	if o != nil && o.ExtraHeader1.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader1 gets a reference to the given NullableString and assigns it to the ExtraHeader1 field.
func (o *DebUpstream) SetExtraHeader1(v string) {
	o.ExtraHeader1.Set(&v)
}

// SetExtraHeader1Nil sets the value for ExtraHeader1 to be an explicit nil
func (o *DebUpstream) SetExtraHeader1Nil() {
	o.ExtraHeader1.Set(nil)
}

// UnsetExtraHeader1 ensures that no value is present for ExtraHeader1, not even an explicit nil
func (o *DebUpstream) UnsetExtraHeader1() {
	o.ExtraHeader1.Unset()
}

// GetExtraHeader2 returns the ExtraHeader2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetExtraHeader2() string {
	if o == nil || IsNil(o.ExtraHeader2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraHeader2.Get()
}

// GetExtraHeader2Ok returns a tuple with the ExtraHeader2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetExtraHeader2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraHeader2.Get(), o.ExtraHeader2.IsSet()
}

// HasExtraHeader2 returns a boolean if a field has been set.
func (o *DebUpstream) HasExtraHeader2() bool {
	if o != nil && o.ExtraHeader2.IsSet() {
		return true
	}

	return false
}

// SetExtraHeader2 gets a reference to the given NullableString and assigns it to the ExtraHeader2 field.
func (o *DebUpstream) SetExtraHeader2(v string) {
	o.ExtraHeader2.Set(&v)
}

// SetExtraHeader2Nil sets the value for ExtraHeader2 to be an explicit nil
func (o *DebUpstream) SetExtraHeader2Nil() {
	o.ExtraHeader2.Set(nil)
}

// UnsetExtraHeader2 ensures that no value is present for ExtraHeader2, not even an explicit nil
func (o *DebUpstream) UnsetExtraHeader2() {
	o.ExtraHeader2.Unset()
}

// GetExtraValue1 returns the ExtraValue1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetExtraValue1() string {
	if o == nil || IsNil(o.ExtraValue1.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue1.Get()
}

// GetExtraValue1Ok returns a tuple with the ExtraValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetExtraValue1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue1.Get(), o.ExtraValue1.IsSet()
}

// HasExtraValue1 returns a boolean if a field has been set.
func (o *DebUpstream) HasExtraValue1() bool {
	if o != nil && o.ExtraValue1.IsSet() {
		return true
	}

	return false
}

// SetExtraValue1 gets a reference to the given NullableString and assigns it to the ExtraValue1 field.
func (o *DebUpstream) SetExtraValue1(v string) {
	o.ExtraValue1.Set(&v)
}

// SetExtraValue1Nil sets the value for ExtraValue1 to be an explicit nil
func (o *DebUpstream) SetExtraValue1Nil() {
	o.ExtraValue1.Set(nil)
}

// UnsetExtraValue1 ensures that no value is present for ExtraValue1, not even an explicit nil
func (o *DebUpstream) UnsetExtraValue1() {
	o.ExtraValue1.Unset()
}

// GetExtraValue2 returns the ExtraValue2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetExtraValue2() string {
	if o == nil || IsNil(o.ExtraValue2.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraValue2.Get()
}

// GetExtraValue2Ok returns a tuple with the ExtraValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetExtraValue2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraValue2.Get(), o.ExtraValue2.IsSet()
}

// HasExtraValue2 returns a boolean if a field has been set.
func (o *DebUpstream) HasExtraValue2() bool {
	if o != nil && o.ExtraValue2.IsSet() {
		return true
	}

	return false
}

// SetExtraValue2 gets a reference to the given NullableString and assigns it to the ExtraValue2 field.
func (o *DebUpstream) SetExtraValue2(v string) {
	o.ExtraValue2.Set(&v)
}

// SetExtraValue2Nil sets the value for ExtraValue2 to be an explicit nil
func (o *DebUpstream) SetExtraValue2Nil() {
	o.ExtraValue2.Set(nil)
}

// UnsetExtraValue2 ensures that no value is present for ExtraValue2, not even an explicit nil
func (o *DebUpstream) UnsetExtraValue2() {
	o.ExtraValue2.Unset()
}

// GetGpgKeyInline returns the GpgKeyInline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetGpgKeyInline() string {
	if o == nil || IsNil(o.GpgKeyInline.Get()) {
		var ret string
		return ret
	}
	return *o.GpgKeyInline.Get()
}

// GetGpgKeyInlineOk returns a tuple with the GpgKeyInline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetGpgKeyInlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpgKeyInline.Get(), o.GpgKeyInline.IsSet()
}

// HasGpgKeyInline returns a boolean if a field has been set.
func (o *DebUpstream) HasGpgKeyInline() bool {
	if o != nil && o.GpgKeyInline.IsSet() {
		return true
	}

	return false
}

// SetGpgKeyInline gets a reference to the given NullableString and assigns it to the GpgKeyInline field.
func (o *DebUpstream) SetGpgKeyInline(v string) {
	o.GpgKeyInline.Set(&v)
}

// SetGpgKeyInlineNil sets the value for GpgKeyInline to be an explicit nil
func (o *DebUpstream) SetGpgKeyInlineNil() {
	o.GpgKeyInline.Set(nil)
}

// UnsetGpgKeyInline ensures that no value is present for GpgKeyInline, not even an explicit nil
func (o *DebUpstream) UnsetGpgKeyInline() {
	o.GpgKeyInline.Unset()
}

// GetGpgKeyUrl returns the GpgKeyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetGpgKeyUrl() string {
	if o == nil || IsNil(o.GpgKeyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.GpgKeyUrl.Get()
}

// GetGpgKeyUrlOk returns a tuple with the GpgKeyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetGpgKeyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpgKeyUrl.Get(), o.GpgKeyUrl.IsSet()
}

// HasGpgKeyUrl returns a boolean if a field has been set.
func (o *DebUpstream) HasGpgKeyUrl() bool {
	if o != nil && o.GpgKeyUrl.IsSet() {
		return true
	}

	return false
}

// SetGpgKeyUrl gets a reference to the given NullableString and assigns it to the GpgKeyUrl field.
func (o *DebUpstream) SetGpgKeyUrl(v string) {
	o.GpgKeyUrl.Set(&v)
}

// SetGpgKeyUrlNil sets the value for GpgKeyUrl to be an explicit nil
func (o *DebUpstream) SetGpgKeyUrlNil() {
	o.GpgKeyUrl.Set(nil)
}

// UnsetGpgKeyUrl ensures that no value is present for GpgKeyUrl, not even an explicit nil
func (o *DebUpstream) UnsetGpgKeyUrl() {
	o.GpgKeyUrl.Unset()
}

// GetGpgVerification returns the GpgVerification field value if set, zero value otherwise.
func (o *DebUpstream) GetGpgVerification() string {
	if o == nil || IsNil(o.GpgVerification) {
		var ret string
		return ret
	}
	return *o.GpgVerification
}

// GetGpgVerificationOk returns a tuple with the GpgVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetGpgVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.GpgVerification) {
		return nil, false
	}
	return o.GpgVerification, true
}

// HasGpgVerification returns a boolean if a field has been set.
func (o *DebUpstream) HasGpgVerification() bool {
	if o != nil && !IsNil(o.GpgVerification) {
		return true
	}

	return false
}

// SetGpgVerification gets a reference to the given string and assigns it to the GpgVerification field.
func (o *DebUpstream) SetGpgVerification(v string) {
	o.GpgVerification = &v
}

// GetIncludeSources returns the IncludeSources field value if set, zero value otherwise.
func (o *DebUpstream) GetIncludeSources() bool {
	if o == nil || IsNil(o.IncludeSources) {
		var ret bool
		return ret
	}
	return *o.IncludeSources
}

// GetIncludeSourcesOk returns a tuple with the IncludeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetIncludeSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSources) {
		return nil, false
	}
	return o.IncludeSources, true
}

// HasIncludeSources returns a boolean if a field has been set.
func (o *DebUpstream) HasIncludeSources() bool {
	if o != nil && !IsNil(o.IncludeSources) {
		return true
	}

	return false
}

// SetIncludeSources gets a reference to the given bool and assigns it to the IncludeSources field.
func (o *DebUpstream) SetIncludeSources(v bool) {
	o.IncludeSources = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *DebUpstream) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *DebUpstream) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *DebUpstream) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DebUpstream) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DebUpstream) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DebUpstream) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value
func (o *DebUpstream) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DebUpstream) SetName(v string) {
	o.Name = v
}

// GetPendingValidation returns the PendingValidation field value if set, zero value otherwise.
func (o *DebUpstream) GetPendingValidation() bool {
	if o == nil || IsNil(o.PendingValidation) {
		var ret bool
		return ret
	}
	return *o.PendingValidation
}

// GetPendingValidationOk returns a tuple with the PendingValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetPendingValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.PendingValidation) {
		return nil, false
	}
	return o.PendingValidation, true
}

// HasPendingValidation returns a boolean if a field has been set.
func (o *DebUpstream) HasPendingValidation() bool {
	if o != nil && !IsNil(o.PendingValidation) {
		return true
	}

	return false
}

// SetPendingValidation gets a reference to the given bool and assigns it to the PendingValidation field.
func (o *DebUpstream) SetPendingValidation(v bool) {
	o.PendingValidation = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DebUpstream) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DebUpstream) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *DebUpstream) SetPriority(v int64) {
	o.Priority = &v
}

// GetSlugPerm returns the SlugPerm field value if set, zero value otherwise.
func (o *DebUpstream) GetSlugPerm() string {
	if o == nil || IsNil(o.SlugPerm) {
		var ret string
		return ret
	}
	return *o.SlugPerm
}

// GetSlugPermOk returns a tuple with the SlugPerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetSlugPermOk() (*string, bool) {
	if o == nil || IsNil(o.SlugPerm) {
		return nil, false
	}
	return o.SlugPerm, true
}

// HasSlugPerm returns a boolean if a field has been set.
func (o *DebUpstream) HasSlugPerm() bool {
	if o != nil && !IsNil(o.SlugPerm) {
		return true
	}

	return false
}

// SetSlugPerm gets a reference to the given string and assigns it to the SlugPerm field.
func (o *DebUpstream) SetSlugPerm(v string) {
	o.SlugPerm = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DebUpstream) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DebUpstream) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DebUpstream) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpstreamDistribution returns the UpstreamDistribution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DebUpstream) GetUpstreamDistribution() string {
	if o == nil || IsNil(o.UpstreamDistribution.Get()) {
		var ret string
		return ret
	}
	return *o.UpstreamDistribution.Get()
}

// GetUpstreamDistributionOk returns a tuple with the UpstreamDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DebUpstream) GetUpstreamDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamDistribution.Get(), o.UpstreamDistribution.IsSet()
}

// HasUpstreamDistribution returns a boolean if a field has been set.
func (o *DebUpstream) HasUpstreamDistribution() bool {
	if o != nil && o.UpstreamDistribution.IsSet() {
		return true
	}

	return false
}

// SetUpstreamDistribution gets a reference to the given NullableString and assigns it to the UpstreamDistribution field.
func (o *DebUpstream) SetUpstreamDistribution(v string) {
	o.UpstreamDistribution.Set(&v)
}

// SetUpstreamDistributionNil sets the value for UpstreamDistribution to be an explicit nil
func (o *DebUpstream) SetUpstreamDistributionNil() {
	o.UpstreamDistribution.Set(nil)
}

// UnsetUpstreamDistribution ensures that no value is present for UpstreamDistribution, not even an explicit nil
func (o *DebUpstream) UnsetUpstreamDistribution() {
	o.UpstreamDistribution.Unset()
}

// GetUpstreamUrl returns the UpstreamUrl field value
func (o *DebUpstream) GetUpstreamUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpstreamUrl
}

// GetUpstreamUrlOk returns a tuple with the UpstreamUrl field value
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetUpstreamUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpstreamUrl, true
}

// SetUpstreamUrl sets field value
func (o *DebUpstream) SetUpstreamUrl(v string) {
	o.UpstreamUrl = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *DebUpstream) GetVerificationStatus() string {
	if o == nil || IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetVerificationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *DebUpstream) HasVerificationStatus() bool {
	if o != nil && !IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *DebUpstream) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *DebUpstream) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebUpstream) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *DebUpstream) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *DebUpstream) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o DebUpstream) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebUpstream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthMode) {
		toSerialize["auth_mode"] = o.AuthMode
	}
	if o.AuthSecret.IsSet() {
		toSerialize["auth_secret"] = o.AuthSecret.Get()
	}
	if o.AuthUsername.IsSet() {
		toSerialize["auth_username"] = o.AuthUsername.Get()
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DisableReason) {
		toSerialize["disable_reason"] = o.DisableReason
	}
	toSerialize["distro_versions"] = o.DistroVersions
	if o.ExtraHeader1.IsSet() {
		toSerialize["extra_header_1"] = o.ExtraHeader1.Get()
	}
	if o.ExtraHeader2.IsSet() {
		toSerialize["extra_header_2"] = o.ExtraHeader2.Get()
	}
	if o.ExtraValue1.IsSet() {
		toSerialize["extra_value_1"] = o.ExtraValue1.Get()
	}
	if o.ExtraValue2.IsSet() {
		toSerialize["extra_value_2"] = o.ExtraValue2.Get()
	}
	if o.GpgKeyInline.IsSet() {
		toSerialize["gpg_key_inline"] = o.GpgKeyInline.Get()
	}
	if o.GpgKeyUrl.IsSet() {
		toSerialize["gpg_key_url"] = o.GpgKeyUrl.Get()
	}
	if !IsNil(o.GpgVerification) {
		toSerialize["gpg_verification"] = o.GpgVerification
	}
	if !IsNil(o.IncludeSources) {
		toSerialize["include_sources"] = o.IncludeSources
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PendingValidation) {
		toSerialize["pending_validation"] = o.PendingValidation
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.SlugPerm) {
		toSerialize["slug_perm"] = o.SlugPerm
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpstreamDistribution.IsSet() {
		toSerialize["upstream_distribution"] = o.UpstreamDistribution.Get()
	}
	toSerialize["upstream_url"] = o.UpstreamUrl
	if !IsNil(o.VerificationStatus) {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	return toSerialize, nil
}

type NullableDebUpstream struct {
	value *DebUpstream
	isSet bool
}

func (v NullableDebUpstream) Get() *DebUpstream {
	return v.value
}

func (v *NullableDebUpstream) Set(val *DebUpstream) {
	v.value = val
	v.isSet = true
}

func (v NullableDebUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableDebUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebUpstream(val *DebUpstream) *NullableDebUpstream {
	return &NullableDebUpstream{value: val, isSet: true}
}

func (v NullableDebUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.390.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the RepositoryWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryWebhookRequest{}

// RepositoryWebhookRequest struct for RepositoryWebhookRequest
type RepositoryWebhookRequest struct {
	Events []string `json:"events"`
	// If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL.
	IsActive *bool `json:"is_active,omitempty"`
	// The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire.
	PackageQuery NullableString `json:"package_query,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyFormat *int64 `json:"request_body_format,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyTemplateFormat *int64 `json:"request_body_template_format,omitempty"`
	// The value that will be sent for the 'Content Type' header.
	RequestContentType NullableString `json:"request_content_type,omitempty"`
	// The header to send the predefined secret in. This must be unique from existing headers or it won't be sent. You can use this as a form of authentication on the endpoint side.
	SecretHeader NullableString `json:"secret_header,omitempty"`
	// The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side.
	SecretValue NullableString `json:"secret_value,omitempty"`
	// The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn't modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it).
	SignatureKey *string `json:"signature_key,omitempty"`
	// The destination URL that webhook payloads will be POST'ed to.
	TargetUrl string            `json:"target_url"`
	Templates []WebhookTemplate `json:"templates"`
	// If enabled, SSL certificates is verified when webhooks are sent. It's recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks.
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

// NewRepositoryWebhookRequest instantiates a new RepositoryWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryWebhookRequest(events []string, targetUrl string, templates []WebhookTemplate) *RepositoryWebhookRequest {
	this := RepositoryWebhookRequest{}
	this.Events = events
	this.TargetUrl = targetUrl
	this.Templates = templates
	return &this
}

// NewRepositoryWebhookRequestWithDefaults instantiates a new RepositoryWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWebhookRequestWithDefaults() *RepositoryWebhookRequest {
	this := RepositoryWebhookRequest{}
	return &this
}

// GetEvents returns the Events field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *RepositoryWebhookRequest) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *RepositoryWebhookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RepositoryWebhookRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RepositoryWebhookRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetPackageQuery returns the PackageQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequest) GetPackageQuery() string {
	if o == nil || IsNil(o.PackageQuery.Get()) {
		var ret string
		return ret
	}
	return *o.PackageQuery.Get()
}

// GetPackageQueryOk returns a tuple with the PackageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequest) GetPackageQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageQuery.Get(), o.PackageQuery.IsSet()
}

// HasPackageQuery returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasPackageQuery() bool {
	if o != nil && o.PackageQuery.IsSet() {
		return true
	}

	return false
}

// SetPackageQuery gets a reference to the given NullableString and assigns it to the PackageQuery field.
func (o *RepositoryWebhookRequest) SetPackageQuery(v string) {
	o.PackageQuery.Set(&v)
}

// SetPackageQueryNil sets the value for PackageQuery to be an explicit nil
func (o *RepositoryWebhookRequest) SetPackageQueryNil() {
	o.PackageQuery.Set(nil)
}

// UnsetPackageQuery ensures that no value is present for PackageQuery, not even an explicit nil
func (o *RepositoryWebhookRequest) UnsetPackageQuery() {
	o.PackageQuery.Unset()
}

// GetRequestBodyFormat returns the RequestBodyFormat field value if set, zero value otherwise.
func (o *RepositoryWebhookRequest) GetRequestBodyFormat() int64 {
	if o == nil || IsNil(o.RequestBodyFormat) {
		var ret int64
		return ret
	}
	return *o.RequestBodyFormat
}

// GetRequestBodyFormatOk returns a tuple with the RequestBodyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequest) GetRequestBodyFormatOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestBodyFormat) {
		return nil, false
	}
	return o.RequestBodyFormat, true
}

// HasRequestBodyFormat returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasRequestBodyFormat() bool {
	if o != nil && !IsNil(o.RequestBodyFormat) {
		return true
	}

	return false
}

// SetRequestBodyFormat gets a reference to the given int64 and assigns it to the RequestBodyFormat field.
func (o *RepositoryWebhookRequest) SetRequestBodyFormat(v int64) {
	o.RequestBodyFormat = &v
}

// GetRequestBodyTemplateFormat returns the RequestBodyTemplateFormat field value if set, zero value otherwise.
func (o *RepositoryWebhookRequest) GetRequestBodyTemplateFormat() int64 {
	if o == nil || IsNil(o.RequestBodyTemplateFormat) {
		var ret int64
		return ret
	}
	return *o.RequestBodyTemplateFormat
}

// GetRequestBodyTemplateFormatOk returns a tuple with the RequestBodyTemplateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequest) GetRequestBodyTemplateFormatOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestBodyTemplateFormat) {
		return nil, false
	}
	return o.RequestBodyTemplateFormat, true
}

// HasRequestBodyTemplateFormat returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasRequestBodyTemplateFormat() bool {
	if o != nil && !IsNil(o.RequestBodyTemplateFormat) {
		return true
	}

	return false
}

// SetRequestBodyTemplateFormat gets a reference to the given int64 and assigns it to the RequestBodyTemplateFormat field.
func (o *RepositoryWebhookRequest) SetRequestBodyTemplateFormat(v int64) {
	o.RequestBodyTemplateFormat = &v
}

// GetRequestContentType returns the RequestContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequest) GetRequestContentType() string {
	if o == nil || IsNil(o.RequestContentType.Get()) {
		var ret string
		return ret
	}
	return *o.RequestContentType.Get()
}

// GetRequestContentTypeOk returns a tuple with the RequestContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequest) GetRequestContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestContentType.Get(), o.RequestContentType.IsSet()
}

// HasRequestContentType returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasRequestContentType() bool {
	if o != nil && o.RequestContentType.IsSet() {
		return true
	}

	return false
}

// SetRequestContentType gets a reference to the given NullableString and assigns it to the RequestContentType field.
func (o *RepositoryWebhookRequest) SetRequestContentType(v string) {
	o.RequestContentType.Set(&v)
}

// SetRequestContentTypeNil sets the value for RequestContentType to be an explicit nil
func (o *RepositoryWebhookRequest) SetRequestContentTypeNil() {
	o.RequestContentType.Set(nil)
}

// UnsetRequestContentType ensures that no value is present for RequestContentType, not even an explicit nil
func (o *RepositoryWebhookRequest) UnsetRequestContentType() {
	o.RequestContentType.Unset()
}

// GetSecretHeader returns the SecretHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequest) GetSecretHeader() string {
	if o == nil || IsNil(o.SecretHeader.Get()) {
		var ret string
		return ret
	}
	return *o.SecretHeader.Get()
}

// GetSecretHeaderOk returns a tuple with the SecretHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequest) GetSecretHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretHeader.Get(), o.SecretHeader.IsSet()
}

// HasSecretHeader returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasSecretHeader() bool {
	if o != nil && o.SecretHeader.IsSet() {
		return true
	}

	return false
}

// SetSecretHeader gets a reference to the given NullableString and assigns it to the SecretHeader field.
func (o *RepositoryWebhookRequest) SetSecretHeader(v string) {
	o.SecretHeader.Set(&v)
}

// SetSecretHeaderNil sets the value for SecretHeader to be an explicit nil
func (o *RepositoryWebhookRequest) SetSecretHeaderNil() {
	o.SecretHeader.Set(nil)
}

// UnsetSecretHeader ensures that no value is present for SecretHeader, not even an explicit nil
func (o *RepositoryWebhookRequest) UnsetSecretHeader() {
	o.SecretHeader.Unset()
}

// GetSecretValue returns the SecretValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryWebhookRequest) GetSecretValue() string {
	if o == nil || IsNil(o.SecretValue.Get()) {
		var ret string
		return ret
	}
	return *o.SecretValue.Get()
}

// GetSecretValueOk returns a tuple with the SecretValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequest) GetSecretValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretValue.Get(), o.SecretValue.IsSet()
}

// HasSecretValue returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasSecretValue() bool {
	if o != nil && o.SecretValue.IsSet() {
		return true
	}

	return false
}

// SetSecretValue gets a reference to the given NullableString and assigns it to the SecretValue field.
func (o *RepositoryWebhookRequest) SetSecretValue(v string) {
	o.SecretValue.Set(&v)
}

// SetSecretValueNil sets the value for SecretValue to be an explicit nil
func (o *RepositoryWebhookRequest) SetSecretValueNil() {
	o.SecretValue.Set(nil)
}

// UnsetSecretValue ensures that no value is present for SecretValue, not even an explicit nil
func (o *RepositoryWebhookRequest) UnsetSecretValue() {
	o.SecretValue.Unset()
}

// GetSignatureKey returns the SignatureKey field value if set, zero value otherwise.
func (o *RepositoryWebhookRequest) GetSignatureKey() string {
	if o == nil || IsNil(o.SignatureKey) {
		var ret string
		return ret
	}
	return *o.SignatureKey
}

// GetSignatureKeyOk returns a tuple with the SignatureKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequest) GetSignatureKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureKey) {
		return nil, false
	}
	return o.SignatureKey, true
}

// HasSignatureKey returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasSignatureKey() bool {
	if o != nil && !IsNil(o.SignatureKey) {
		return true
	}

	return false
}

// SetSignatureKey gets a reference to the given string and assigns it to the SignatureKey field.
func (o *RepositoryWebhookRequest) SetSignatureKey(v string) {
	o.SignatureKey = &v
}

// GetTargetUrl returns the TargetUrl field value
func (o *RepositoryWebhookRequest) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequest) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *RepositoryWebhookRequest) SetTargetUrl(v string) {
	o.TargetUrl = v
}

// GetTemplates returns the Templates field value
// If the value is explicit nil, the zero value for []WebhookTemplate will be returned
func (o *RepositoryWebhookRequest) GetTemplates() []WebhookTemplate {
	if o == nil {
		var ret []WebhookTemplate
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryWebhookRequest) GetTemplatesOk() ([]WebhookTemplate, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *RepositoryWebhookRequest) SetTemplates(v []WebhookTemplate) {
	o.Templates = v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *RepositoryWebhookRequest) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryWebhookRequest) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *RepositoryWebhookRequest) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *RepositoryWebhookRequest) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o RepositoryWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if o.PackageQuery.IsSet() {
		toSerialize["package_query"] = o.PackageQuery.Get()
	}
	if !IsNil(o.RequestBodyFormat) {
		toSerialize["request_body_format"] = o.RequestBodyFormat
	}
	if !IsNil(o.RequestBodyTemplateFormat) {
		toSerialize["request_body_template_format"] = o.RequestBodyTemplateFormat
	}
	if o.RequestContentType.IsSet() {
		toSerialize["request_content_type"] = o.RequestContentType.Get()
	}
	if o.SecretHeader.IsSet() {
		toSerialize["secret_header"] = o.SecretHeader.Get()
	}
	if o.SecretValue.IsSet() {
		toSerialize["secret_value"] = o.SecretValue.Get()
	}
	if !IsNil(o.SignatureKey) {
		toSerialize["signature_key"] = o.SignatureKey
	}
	toSerialize["target_url"] = o.TargetUrl
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	return toSerialize, nil
}

type NullableRepositoryWebhookRequest struct {
	value *RepositoryWebhookRequest
	isSet bool
}

func (v NullableRepositoryWebhookRequest) Get() *RepositoryWebhookRequest {
	return v.value
}

func (v *NullableRepositoryWebhookRequest) Set(val *RepositoryWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryWebhookRequest(val *RepositoryWebhookRequest) *NullableRepositoryWebhookRequest {
	return &NullableRepositoryWebhookRequest{value: val, isSet: true}
}

func (v NullableRepositoryWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

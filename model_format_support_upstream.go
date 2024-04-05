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

// checks if the FormatSupportUpstream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormatSupportUpstream{}

// FormatSupportUpstream The upstream support for the package format
type FormatSupportUpstream struct {
	// The authentication modes supported by the upstream format
	AuthModes []string `json:"auth_modes"`
	// If true the upstream format supports caching
	Caching bool `json:"caching"`
	// If true the upstream format supports indexing
	Indexing bool `json:"indexing"`
	// The behavior of the upstream when indexing
	IndexingBehavior *string `json:"indexing_behavior,omitempty"`
	// If true the upstream format supports proxying
	Proxying bool `json:"proxying"`
	// The signature verification supported by the upstream format
	SignatureVerification *string `json:"signature_verification,omitempty"`
}

// NewFormatSupportUpstream instantiates a new FormatSupportUpstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatSupportUpstream(authModes []string, caching bool, indexing bool, proxying bool) *FormatSupportUpstream {
	this := FormatSupportUpstream{}
	this.AuthModes = authModes
	this.Caching = caching
	this.Indexing = indexing
	var indexingBehavior string = "Unsupported"
	this.IndexingBehavior = &indexingBehavior
	this.Proxying = proxying
	var signatureVerification string = "Unsupported"
	this.SignatureVerification = &signatureVerification
	return &this
}

// NewFormatSupportUpstreamWithDefaults instantiates a new FormatSupportUpstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatSupportUpstreamWithDefaults() *FormatSupportUpstream {
	this := FormatSupportUpstream{}
	var indexingBehavior string = "Unsupported"
	this.IndexingBehavior = &indexingBehavior
	var signatureVerification string = "Unsupported"
	this.SignatureVerification = &signatureVerification
	return &this
}

// GetAuthModes returns the AuthModes field value
func (o *FormatSupportUpstream) GetAuthModes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthModes
}

// GetAuthModesOk returns a tuple with the AuthModes field value
// and a boolean to check if the value has been set.
func (o *FormatSupportUpstream) GetAuthModesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthModes, true
}

// SetAuthModes sets field value
func (o *FormatSupportUpstream) SetAuthModes(v []string) {
	o.AuthModes = v
}

// GetCaching returns the Caching field value
func (o *FormatSupportUpstream) GetCaching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Caching
}

// GetCachingOk returns a tuple with the Caching field value
// and a boolean to check if the value has been set.
func (o *FormatSupportUpstream) GetCachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Caching, true
}

// SetCaching sets field value
func (o *FormatSupportUpstream) SetCaching(v bool) {
	o.Caching = v
}

// GetIndexing returns the Indexing field value
func (o *FormatSupportUpstream) GetIndexing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value
// and a boolean to check if the value has been set.
func (o *FormatSupportUpstream) GetIndexingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexing, true
}

// SetIndexing sets field value
func (o *FormatSupportUpstream) SetIndexing(v bool) {
	o.Indexing = v
}

// GetIndexingBehavior returns the IndexingBehavior field value if set, zero value otherwise.
func (o *FormatSupportUpstream) GetIndexingBehavior() string {
	if o == nil || IsNil(o.IndexingBehavior) {
		var ret string
		return ret
	}
	return *o.IndexingBehavior
}

// GetIndexingBehaviorOk returns a tuple with the IndexingBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatSupportUpstream) GetIndexingBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.IndexingBehavior) {
		return nil, false
	}
	return o.IndexingBehavior, true
}

// HasIndexingBehavior returns a boolean if a field has been set.
func (o *FormatSupportUpstream) HasIndexingBehavior() bool {
	if o != nil && !IsNil(o.IndexingBehavior) {
		return true
	}

	return false
}

// SetIndexingBehavior gets a reference to the given string and assigns it to the IndexingBehavior field.
func (o *FormatSupportUpstream) SetIndexingBehavior(v string) {
	o.IndexingBehavior = &v
}

// GetProxying returns the Proxying field value
func (o *FormatSupportUpstream) GetProxying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Proxying
}

// GetProxyingOk returns a tuple with the Proxying field value
// and a boolean to check if the value has been set.
func (o *FormatSupportUpstream) GetProxyingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxying, true
}

// SetProxying sets field value
func (o *FormatSupportUpstream) SetProxying(v bool) {
	o.Proxying = v
}

// GetSignatureVerification returns the SignatureVerification field value if set, zero value otherwise.
func (o *FormatSupportUpstream) GetSignatureVerification() string {
	if o == nil || IsNil(o.SignatureVerification) {
		var ret string
		return ret
	}
	return *o.SignatureVerification
}

// GetSignatureVerificationOk returns a tuple with the SignatureVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatSupportUpstream) GetSignatureVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureVerification) {
		return nil, false
	}
	return o.SignatureVerification, true
}

// HasSignatureVerification returns a boolean if a field has been set.
func (o *FormatSupportUpstream) HasSignatureVerification() bool {
	if o != nil && !IsNil(o.SignatureVerification) {
		return true
	}

	return false
}

// SetSignatureVerification gets a reference to the given string and assigns it to the SignatureVerification field.
func (o *FormatSupportUpstream) SetSignatureVerification(v string) {
	o.SignatureVerification = &v
}

func (o FormatSupportUpstream) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormatSupportUpstream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth_modes"] = o.AuthModes
	toSerialize["caching"] = o.Caching
	toSerialize["indexing"] = o.Indexing
	if !IsNil(o.IndexingBehavior) {
		toSerialize["indexing_behavior"] = o.IndexingBehavior
	}
	toSerialize["proxying"] = o.Proxying
	if !IsNil(o.SignatureVerification) {
		toSerialize["signature_verification"] = o.SignatureVerification
	}
	return toSerialize, nil
}

type NullableFormatSupportUpstream struct {
	value *FormatSupportUpstream
	isSet bool
}

func (v NullableFormatSupportUpstream) Get() *FormatSupportUpstream {
	return v.value
}

func (v *NullableFormatSupportUpstream) Set(val *FormatSupportUpstream) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatSupportUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatSupportUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatSupportUpstream(val *FormatSupportUpstream) *NullableFormatSupportUpstream {
	return &NullableFormatSupportUpstream{value: val, isSet: true}
}

func (v NullableFormatSupportUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatSupportUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

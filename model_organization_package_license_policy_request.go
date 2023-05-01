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

// OrganizationPackageLicensePolicyRequest struct for OrganizationPackageLicensePolicyRequest
type OrganizationPackageLicensePolicyRequest struct {
	AllowUnknownLicenses  *bool          `json:"allow_unknown_licenses,omitempty"`
	Description           NullableString `json:"description,omitempty"`
	Name                  string         `json:"name"`
	OnViolationQuarantine *bool          `json:"on_violation_quarantine,omitempty"`
	SpdxIdentifiers       []string       `json:"spdx_identifiers"`
}

// NewOrganizationPackageLicensePolicyRequest instantiates a new OrganizationPackageLicensePolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationPackageLicensePolicyRequest(name string, spdxIdentifiers []string) *OrganizationPackageLicensePolicyRequest {
	this := OrganizationPackageLicensePolicyRequest{}
	this.Name = name
	this.SpdxIdentifiers = spdxIdentifiers
	return &this
}

// NewOrganizationPackageLicensePolicyRequestWithDefaults instantiates a new OrganizationPackageLicensePolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationPackageLicensePolicyRequestWithDefaults() *OrganizationPackageLicensePolicyRequest {
	this := OrganizationPackageLicensePolicyRequest{}
	return &this
}

// GetAllowUnknownLicenses returns the AllowUnknownLicenses field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicyRequest) GetAllowUnknownLicenses() bool {
	if o == nil || isNil(o.AllowUnknownLicenses) {
		var ret bool
		return ret
	}
	return *o.AllowUnknownLicenses
}

// GetAllowUnknownLicensesOk returns a tuple with the AllowUnknownLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequest) GetAllowUnknownLicensesOk() (*bool, bool) {
	if o == nil || isNil(o.AllowUnknownLicenses) {
		return nil, false
	}
	return o.AllowUnknownLicenses, true
}

// HasAllowUnknownLicenses returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequest) HasAllowUnknownLicenses() bool {
	if o != nil && !isNil(o.AllowUnknownLicenses) {
		return true
	}

	return false
}

// SetAllowUnknownLicenses gets a reference to the given bool and assigns it to the AllowUnknownLicenses field.
func (o *OrganizationPackageLicensePolicyRequest) SetAllowUnknownLicenses(v bool) {
	o.AllowUnknownLicenses = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationPackageLicensePolicyRequest) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationPackageLicensePolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OrganizationPackageLicensePolicyRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OrganizationPackageLicensePolicyRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OrganizationPackageLicensePolicyRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *OrganizationPackageLicensePolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationPackageLicensePolicyRequest) SetName(v string) {
	o.Name = v
}

// GetOnViolationQuarantine returns the OnViolationQuarantine field value if set, zero value otherwise.
func (o *OrganizationPackageLicensePolicyRequest) GetOnViolationQuarantine() bool {
	if o == nil || isNil(o.OnViolationQuarantine) {
		var ret bool
		return ret
	}
	return *o.OnViolationQuarantine
}

// GetOnViolationQuarantineOk returns a tuple with the OnViolationQuarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequest) GetOnViolationQuarantineOk() (*bool, bool) {
	if o == nil || isNil(o.OnViolationQuarantine) {
		return nil, false
	}
	return o.OnViolationQuarantine, true
}

// HasOnViolationQuarantine returns a boolean if a field has been set.
func (o *OrganizationPackageLicensePolicyRequest) HasOnViolationQuarantine() bool {
	if o != nil && !isNil(o.OnViolationQuarantine) {
		return true
	}

	return false
}

// SetOnViolationQuarantine gets a reference to the given bool and assigns it to the OnViolationQuarantine field.
func (o *OrganizationPackageLicensePolicyRequest) SetOnViolationQuarantine(v bool) {
	o.OnViolationQuarantine = &v
}

// GetSpdxIdentifiers returns the SpdxIdentifiers field value
func (o *OrganizationPackageLicensePolicyRequest) GetSpdxIdentifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SpdxIdentifiers
}

// GetSpdxIdentifiersOk returns a tuple with the SpdxIdentifiers field value
// and a boolean to check if the value has been set.
func (o *OrganizationPackageLicensePolicyRequest) GetSpdxIdentifiersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpdxIdentifiers, true
}

// SetSpdxIdentifiers sets field value
func (o *OrganizationPackageLicensePolicyRequest) SetSpdxIdentifiers(v []string) {
	o.SpdxIdentifiers = v
}

func (o OrganizationPackageLicensePolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowUnknownLicenses) {
		toSerialize["allow_unknown_licenses"] = o.AllowUnknownLicenses
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OnViolationQuarantine) {
		toSerialize["on_violation_quarantine"] = o.OnViolationQuarantine
	}
	if true {
		toSerialize["spdx_identifiers"] = o.SpdxIdentifiers
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationPackageLicensePolicyRequest struct {
	value *OrganizationPackageLicensePolicyRequest
	isSet bool
}

func (v NullableOrganizationPackageLicensePolicyRequest) Get() *OrganizationPackageLicensePolicyRequest {
	return v.value
}

func (v *NullableOrganizationPackageLicensePolicyRequest) Set(val *OrganizationPackageLicensePolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationPackageLicensePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationPackageLicensePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationPackageLicensePolicyRequest(val *OrganizationPackageLicensePolicyRequest) *NullableOrganizationPackageLicensePolicyRequest {
	return &NullableOrganizationPackageLicensePolicyRequest{value: val, isSet: true}
}

func (v NullableOrganizationPackageLicensePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationPackageLicensePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

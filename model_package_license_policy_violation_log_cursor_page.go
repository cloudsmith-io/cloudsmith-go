/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.297.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the PackageLicensePolicyViolationLogCursorPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageLicensePolicyViolationLogCursorPage{}

// PackageLicensePolicyViolationLogCursorPage struct for PackageLicensePolicyViolationLogCursorPage
type PackageLicensePolicyViolationLogCursorPage struct {
	Next     NullableString                     `json:"next,omitempty"`
	Previous NullableString                     `json:"previous,omitempty"`
	Results  []PackageLicensePolicyViolationLog `json:"results"`
}

// NewPackageLicensePolicyViolationLogCursorPage instantiates a new PackageLicensePolicyViolationLogCursorPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageLicensePolicyViolationLogCursorPage(results []PackageLicensePolicyViolationLog) *PackageLicensePolicyViolationLogCursorPage {
	this := PackageLicensePolicyViolationLogCursorPage{}
	this.Results = results
	return &this
}

// NewPackageLicensePolicyViolationLogCursorPageWithDefaults instantiates a new PackageLicensePolicyViolationLogCursorPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageLicensePolicyViolationLogCursorPageWithDefaults() *PackageLicensePolicyViolationLogCursorPage {
	this := PackageLicensePolicyViolationLogCursorPage{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageLicensePolicyViolationLogCursorPage) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageLicensePolicyViolationLogCursorPage) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PackageLicensePolicyViolationLogCursorPage) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PackageLicensePolicyViolationLogCursorPage) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PackageLicensePolicyViolationLogCursorPage) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PackageLicensePolicyViolationLogCursorPage) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageLicensePolicyViolationLogCursorPage) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageLicensePolicyViolationLogCursorPage) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PackageLicensePolicyViolationLogCursorPage) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PackageLicensePolicyViolationLogCursorPage) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PackageLicensePolicyViolationLogCursorPage) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PackageLicensePolicyViolationLogCursorPage) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PackageLicensePolicyViolationLogCursorPage) GetResults() []PackageLicensePolicyViolationLog {
	if o == nil {
		var ret []PackageLicensePolicyViolationLog
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PackageLicensePolicyViolationLogCursorPage) GetResultsOk() ([]PackageLicensePolicyViolationLog, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PackageLicensePolicyViolationLogCursorPage) SetResults(v []PackageLicensePolicyViolationLog) {
	o.Results = v
}

func (o PackageLicensePolicyViolationLogCursorPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageLicensePolicyViolationLogCursorPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePackageLicensePolicyViolationLogCursorPage struct {
	value *PackageLicensePolicyViolationLogCursorPage
	isSet bool
}

func (v NullablePackageLicensePolicyViolationLogCursorPage) Get() *PackageLicensePolicyViolationLogCursorPage {
	return v.value
}

func (v *NullablePackageLicensePolicyViolationLogCursorPage) Set(val *PackageLicensePolicyViolationLogCursorPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageLicensePolicyViolationLogCursorPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageLicensePolicyViolationLogCursorPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageLicensePolicyViolationLogCursorPage(val *PackageLicensePolicyViolationLogCursorPage) *NullablePackageLicensePolicyViolationLogCursorPage {
	return &NullablePackageLicensePolicyViolationLogCursorPage{value: val, isSet: true}
}

func (v NullablePackageLicensePolicyViolationLogCursorPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageLicensePolicyViolationLogCursorPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

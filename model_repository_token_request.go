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
	"time"
)

// checks if the RepositoryTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryTokenRequest{}

// RepositoryTokenRequest struct for RepositoryTokenRequest
type RepositoryTokenRequest struct {
	// If checked, a EULA acceptance is required for this token.
	EulaRequired *bool `json:"eula_required,omitempty"`
	// If enabled, the token will allow downloads based on configured restrictions (if any).
	IsActive *bool `json:"is_active,omitempty"`
	// The maximum download bandwidth allowed for the token. Values are expressed as the selected unit of bandwidth. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitBandwidth     NullableInt64  `json:"limit_bandwidth,omitempty"`
	LimitBandwidthUnit NullableString `json:"limit_bandwidth_unit,omitempty"`
	// The starting date/time the token is allowed to be used from.
	LimitDateRangeFrom NullableTime `json:"limit_date_range_from,omitempty"`
	// The ending date/time the token is allowed to be used until.
	LimitDateRangeTo NullableTime `json:"limit_date_range_to,omitempty"`
	// The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumClients NullableInt64 `json:"limit_num_clients,omitempty"`
	// The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.
	LimitNumDownloads NullableInt64 `json:"limit_num_downloads,omitempty"`
	// The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata.
	LimitPackageQuery NullableString `json:"limit_package_query,omitempty"`
	// THIS WILL SOON BE DEPRECATED, please use limit_package_query instead. The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.
	LimitPathQuery NullableString         `json:"limit_path_query,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	Name           string                 `json:"name"`
	// The time at which the scheduled reset period has elapsed and the token limits were automatically reset to zero.
	ScheduledResetAt     NullableTime   `json:"scheduled_reset_at,omitempty"`
	ScheduledResetPeriod NullableString `json:"scheduled_reset_period,omitempty"`
	Token                *string        `json:"token,omitempty"`
}

// NewRepositoryTokenRequest instantiates a new RepositoryTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryTokenRequest(name string) *RepositoryTokenRequest {
	this := RepositoryTokenRequest{}
	var limitBandwidthUnit string = "Byte"
	this.LimitBandwidthUnit = *NewNullableString(&limitBandwidthUnit)
	this.Name = name
	var scheduledResetPeriod string = "Never Reset"
	this.ScheduledResetPeriod = *NewNullableString(&scheduledResetPeriod)
	return &this
}

// NewRepositoryTokenRequestWithDefaults instantiates a new RepositoryTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryTokenRequestWithDefaults() *RepositoryTokenRequest {
	this := RepositoryTokenRequest{}
	var limitBandwidthUnit string = "Byte"
	this.LimitBandwidthUnit = *NewNullableString(&limitBandwidthUnit)
	var scheduledResetPeriod string = "Never Reset"
	this.ScheduledResetPeriod = *NewNullableString(&scheduledResetPeriod)
	return &this
}

// GetEulaRequired returns the EulaRequired field value if set, zero value otherwise.
func (o *RepositoryTokenRequest) GetEulaRequired() bool {
	if o == nil || IsNil(o.EulaRequired) {
		var ret bool
		return ret
	}
	return *o.EulaRequired
}

// GetEulaRequiredOk returns a tuple with the EulaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequest) GetEulaRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.EulaRequired) {
		return nil, false
	}
	return o.EulaRequired, true
}

// HasEulaRequired returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasEulaRequired() bool {
	if o != nil && !IsNil(o.EulaRequired) {
		return true
	}

	return false
}

// SetEulaRequired gets a reference to the given bool and assigns it to the EulaRequired field.
func (o *RepositoryTokenRequest) SetEulaRequired(v bool) {
	o.EulaRequired = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RepositoryTokenRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RepositoryTokenRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLimitBandwidth returns the LimitBandwidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitBandwidth() int64 {
	if o == nil || IsNil(o.LimitBandwidth.Get()) {
		var ret int64
		return ret
	}
	return *o.LimitBandwidth.Get()
}

// GetLimitBandwidthOk returns a tuple with the LimitBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitBandwidthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitBandwidth.Get(), o.LimitBandwidth.IsSet()
}

// HasLimitBandwidth returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitBandwidth() bool {
	if o != nil && o.LimitBandwidth.IsSet() {
		return true
	}

	return false
}

// SetLimitBandwidth gets a reference to the given NullableInt64 and assigns it to the LimitBandwidth field.
func (o *RepositoryTokenRequest) SetLimitBandwidth(v int64) {
	o.LimitBandwidth.Set(&v)
}

// SetLimitBandwidthNil sets the value for LimitBandwidth to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitBandwidthNil() {
	o.LimitBandwidth.Set(nil)
}

// UnsetLimitBandwidth ensures that no value is present for LimitBandwidth, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitBandwidth() {
	o.LimitBandwidth.Unset()
}

// GetLimitBandwidthUnit returns the LimitBandwidthUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitBandwidthUnit() string {
	if o == nil || IsNil(o.LimitBandwidthUnit.Get()) {
		var ret string
		return ret
	}
	return *o.LimitBandwidthUnit.Get()
}

// GetLimitBandwidthUnitOk returns a tuple with the LimitBandwidthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitBandwidthUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitBandwidthUnit.Get(), o.LimitBandwidthUnit.IsSet()
}

// HasLimitBandwidthUnit returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitBandwidthUnit() bool {
	if o != nil && o.LimitBandwidthUnit.IsSet() {
		return true
	}

	return false
}

// SetLimitBandwidthUnit gets a reference to the given NullableString and assigns it to the LimitBandwidthUnit field.
func (o *RepositoryTokenRequest) SetLimitBandwidthUnit(v string) {
	o.LimitBandwidthUnit.Set(&v)
}

// SetLimitBandwidthUnitNil sets the value for LimitBandwidthUnit to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitBandwidthUnitNil() {
	o.LimitBandwidthUnit.Set(nil)
}

// UnsetLimitBandwidthUnit ensures that no value is present for LimitBandwidthUnit, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitBandwidthUnit() {
	o.LimitBandwidthUnit.Unset()
}

// GetLimitDateRangeFrom returns the LimitDateRangeFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitDateRangeFrom() time.Time {
	if o == nil || IsNil(o.LimitDateRangeFrom.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LimitDateRangeFrom.Get()
}

// GetLimitDateRangeFromOk returns a tuple with the LimitDateRangeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitDateRangeFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitDateRangeFrom.Get(), o.LimitDateRangeFrom.IsSet()
}

// HasLimitDateRangeFrom returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitDateRangeFrom() bool {
	if o != nil && o.LimitDateRangeFrom.IsSet() {
		return true
	}

	return false
}

// SetLimitDateRangeFrom gets a reference to the given NullableTime and assigns it to the LimitDateRangeFrom field.
func (o *RepositoryTokenRequest) SetLimitDateRangeFrom(v time.Time) {
	o.LimitDateRangeFrom.Set(&v)
}

// SetLimitDateRangeFromNil sets the value for LimitDateRangeFrom to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitDateRangeFromNil() {
	o.LimitDateRangeFrom.Set(nil)
}

// UnsetLimitDateRangeFrom ensures that no value is present for LimitDateRangeFrom, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitDateRangeFrom() {
	o.LimitDateRangeFrom.Unset()
}

// GetLimitDateRangeTo returns the LimitDateRangeTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitDateRangeTo() time.Time {
	if o == nil || IsNil(o.LimitDateRangeTo.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LimitDateRangeTo.Get()
}

// GetLimitDateRangeToOk returns a tuple with the LimitDateRangeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitDateRangeToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitDateRangeTo.Get(), o.LimitDateRangeTo.IsSet()
}

// HasLimitDateRangeTo returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitDateRangeTo() bool {
	if o != nil && o.LimitDateRangeTo.IsSet() {
		return true
	}

	return false
}

// SetLimitDateRangeTo gets a reference to the given NullableTime and assigns it to the LimitDateRangeTo field.
func (o *RepositoryTokenRequest) SetLimitDateRangeTo(v time.Time) {
	o.LimitDateRangeTo.Set(&v)
}

// SetLimitDateRangeToNil sets the value for LimitDateRangeTo to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitDateRangeToNil() {
	o.LimitDateRangeTo.Set(nil)
}

// UnsetLimitDateRangeTo ensures that no value is present for LimitDateRangeTo, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitDateRangeTo() {
	o.LimitDateRangeTo.Unset()
}

// GetLimitNumClients returns the LimitNumClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitNumClients() int64 {
	if o == nil || IsNil(o.LimitNumClients.Get()) {
		var ret int64
		return ret
	}
	return *o.LimitNumClients.Get()
}

// GetLimitNumClientsOk returns a tuple with the LimitNumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitNumClientsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitNumClients.Get(), o.LimitNumClients.IsSet()
}

// HasLimitNumClients returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitNumClients() bool {
	if o != nil && o.LimitNumClients.IsSet() {
		return true
	}

	return false
}

// SetLimitNumClients gets a reference to the given NullableInt64 and assigns it to the LimitNumClients field.
func (o *RepositoryTokenRequest) SetLimitNumClients(v int64) {
	o.LimitNumClients.Set(&v)
}

// SetLimitNumClientsNil sets the value for LimitNumClients to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitNumClientsNil() {
	o.LimitNumClients.Set(nil)
}

// UnsetLimitNumClients ensures that no value is present for LimitNumClients, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitNumClients() {
	o.LimitNumClients.Unset()
}

// GetLimitNumDownloads returns the LimitNumDownloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitNumDownloads() int64 {
	if o == nil || IsNil(o.LimitNumDownloads.Get()) {
		var ret int64
		return ret
	}
	return *o.LimitNumDownloads.Get()
}

// GetLimitNumDownloadsOk returns a tuple with the LimitNumDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitNumDownloadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitNumDownloads.Get(), o.LimitNumDownloads.IsSet()
}

// HasLimitNumDownloads returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitNumDownloads() bool {
	if o != nil && o.LimitNumDownloads.IsSet() {
		return true
	}

	return false
}

// SetLimitNumDownloads gets a reference to the given NullableInt64 and assigns it to the LimitNumDownloads field.
func (o *RepositoryTokenRequest) SetLimitNumDownloads(v int64) {
	o.LimitNumDownloads.Set(&v)
}

// SetLimitNumDownloadsNil sets the value for LimitNumDownloads to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitNumDownloadsNil() {
	o.LimitNumDownloads.Set(nil)
}

// UnsetLimitNumDownloads ensures that no value is present for LimitNumDownloads, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitNumDownloads() {
	o.LimitNumDownloads.Unset()
}

// GetLimitPackageQuery returns the LimitPackageQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitPackageQuery() string {
	if o == nil || IsNil(o.LimitPackageQuery.Get()) {
		var ret string
		return ret
	}
	return *o.LimitPackageQuery.Get()
}

// GetLimitPackageQueryOk returns a tuple with the LimitPackageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitPackageQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitPackageQuery.Get(), o.LimitPackageQuery.IsSet()
}

// HasLimitPackageQuery returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitPackageQuery() bool {
	if o != nil && o.LimitPackageQuery.IsSet() {
		return true
	}

	return false
}

// SetLimitPackageQuery gets a reference to the given NullableString and assigns it to the LimitPackageQuery field.
func (o *RepositoryTokenRequest) SetLimitPackageQuery(v string) {
	o.LimitPackageQuery.Set(&v)
}

// SetLimitPackageQueryNil sets the value for LimitPackageQuery to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitPackageQueryNil() {
	o.LimitPackageQuery.Set(nil)
}

// UnsetLimitPackageQuery ensures that no value is present for LimitPackageQuery, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitPackageQuery() {
	o.LimitPackageQuery.Unset()
}

// GetLimitPathQuery returns the LimitPathQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetLimitPathQuery() string {
	if o == nil || IsNil(o.LimitPathQuery.Get()) {
		var ret string
		return ret
	}
	return *o.LimitPathQuery.Get()
}

// GetLimitPathQueryOk returns a tuple with the LimitPathQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetLimitPathQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitPathQuery.Get(), o.LimitPathQuery.IsSet()
}

// HasLimitPathQuery returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasLimitPathQuery() bool {
	if o != nil && o.LimitPathQuery.IsSet() {
		return true
	}

	return false
}

// SetLimitPathQuery gets a reference to the given NullableString and assigns it to the LimitPathQuery field.
func (o *RepositoryTokenRequest) SetLimitPathQuery(v string) {
	o.LimitPathQuery.Set(&v)
}

// SetLimitPathQueryNil sets the value for LimitPathQuery to be an explicit nil
func (o *RepositoryTokenRequest) SetLimitPathQueryNil() {
	o.LimitPathQuery.Set(nil)
}

// UnsetLimitPathQuery ensures that no value is present for LimitPathQuery, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetLimitPathQuery() {
	o.LimitPathQuery.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RepositoryTokenRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *RepositoryTokenRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RepositoryTokenRequest) SetName(v string) {
	o.Name = v
}

// GetScheduledResetAt returns the ScheduledResetAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetScheduledResetAt() time.Time {
	if o == nil || IsNil(o.ScheduledResetAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledResetAt.Get()
}

// GetScheduledResetAtOk returns a tuple with the ScheduledResetAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetScheduledResetAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledResetAt.Get(), o.ScheduledResetAt.IsSet()
}

// HasScheduledResetAt returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasScheduledResetAt() bool {
	if o != nil && o.ScheduledResetAt.IsSet() {
		return true
	}

	return false
}

// SetScheduledResetAt gets a reference to the given NullableTime and assigns it to the ScheduledResetAt field.
func (o *RepositoryTokenRequest) SetScheduledResetAt(v time.Time) {
	o.ScheduledResetAt.Set(&v)
}

// SetScheduledResetAtNil sets the value for ScheduledResetAt to be an explicit nil
func (o *RepositoryTokenRequest) SetScheduledResetAtNil() {
	o.ScheduledResetAt.Set(nil)
}

// UnsetScheduledResetAt ensures that no value is present for ScheduledResetAt, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetScheduledResetAt() {
	o.ScheduledResetAt.Unset()
}

// GetScheduledResetPeriod returns the ScheduledResetPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RepositoryTokenRequest) GetScheduledResetPeriod() string {
	if o == nil || IsNil(o.ScheduledResetPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduledResetPeriod.Get()
}

// GetScheduledResetPeriodOk returns a tuple with the ScheduledResetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryTokenRequest) GetScheduledResetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledResetPeriod.Get(), o.ScheduledResetPeriod.IsSet()
}

// HasScheduledResetPeriod returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasScheduledResetPeriod() bool {
	if o != nil && o.ScheduledResetPeriod.IsSet() {
		return true
	}

	return false
}

// SetScheduledResetPeriod gets a reference to the given NullableString and assigns it to the ScheduledResetPeriod field.
func (o *RepositoryTokenRequest) SetScheduledResetPeriod(v string) {
	o.ScheduledResetPeriod.Set(&v)
}

// SetScheduledResetPeriodNil sets the value for ScheduledResetPeriod to be an explicit nil
func (o *RepositoryTokenRequest) SetScheduledResetPeriodNil() {
	o.ScheduledResetPeriod.Set(nil)
}

// UnsetScheduledResetPeriod ensures that no value is present for ScheduledResetPeriod, not even an explicit nil
func (o *RepositoryTokenRequest) UnsetScheduledResetPeriod() {
	o.ScheduledResetPeriod.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RepositoryTokenRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RepositoryTokenRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RepositoryTokenRequest) SetToken(v string) {
	o.Token = &v
}

func (o RepositoryTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EulaRequired) {
		toSerialize["eula_required"] = o.EulaRequired
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LimitBandwidth.IsSet() {
		toSerialize["limit_bandwidth"] = o.LimitBandwidth.Get()
	}
	if o.LimitBandwidthUnit.IsSet() {
		toSerialize["limit_bandwidth_unit"] = o.LimitBandwidthUnit.Get()
	}
	if o.LimitDateRangeFrom.IsSet() {
		toSerialize["limit_date_range_from"] = o.LimitDateRangeFrom.Get()
	}
	if o.LimitDateRangeTo.IsSet() {
		toSerialize["limit_date_range_to"] = o.LimitDateRangeTo.Get()
	}
	if o.LimitNumClients.IsSet() {
		toSerialize["limit_num_clients"] = o.LimitNumClients.Get()
	}
	if o.LimitNumDownloads.IsSet() {
		toSerialize["limit_num_downloads"] = o.LimitNumDownloads.Get()
	}
	if o.LimitPackageQuery.IsSet() {
		toSerialize["limit_package_query"] = o.LimitPackageQuery.Get()
	}
	if o.LimitPathQuery.IsSet() {
		toSerialize["limit_path_query"] = o.LimitPathQuery.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if o.ScheduledResetAt.IsSet() {
		toSerialize["scheduled_reset_at"] = o.ScheduledResetAt.Get()
	}
	if o.ScheduledResetPeriod.IsSet() {
		toSerialize["scheduled_reset_period"] = o.ScheduledResetPeriod.Get()
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableRepositoryTokenRequest struct {
	value *RepositoryTokenRequest
	isSet bool
}

func (v NullableRepositoryTokenRequest) Get() *RepositoryTokenRequest {
	return v.value
}

func (v *NullableRepositoryTokenRequest) Set(val *RepositoryTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryTokenRequest(val *RepositoryTokenRequest) *NullableRepositoryTokenRequest {
	return &NullableRepositoryTokenRequest{value: val, isSet: true}
}

func (v NullableRepositoryTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

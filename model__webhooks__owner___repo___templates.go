/*
Cloudsmith API

The API to the Cloudsmith Service

API version: 1.42.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// WebhooksOwnerRepoTemplates struct for WebhooksOwnerRepoTemplates
type WebhooksOwnerRepoTemplates struct {
	//
	Event *string `json:"event,omitempty"`
	//
	Template *string `json:"template,omitempty"`
}

// NewWebhooksOwnerRepoTemplates instantiates a new WebhooksOwnerRepoTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksOwnerRepoTemplates() *WebhooksOwnerRepoTemplates {
	this := WebhooksOwnerRepoTemplates{}
	return &this
}

// NewWebhooksOwnerRepoTemplatesWithDefaults instantiates a new WebhooksOwnerRepoTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksOwnerRepoTemplatesWithDefaults() *WebhooksOwnerRepoTemplates {
	this := WebhooksOwnerRepoTemplates{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *WebhooksOwnerRepoTemplates) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOwnerRepoTemplates) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *WebhooksOwnerRepoTemplates) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *WebhooksOwnerRepoTemplates) SetEvent(v string) {
	o.Event = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *WebhooksOwnerRepoTemplates) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOwnerRepoTemplates) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *WebhooksOwnerRepoTemplates) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *WebhooksOwnerRepoTemplates) SetTemplate(v string) {
	o.Template = &v
}

func (o WebhooksOwnerRepoTemplates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableWebhooksOwnerRepoTemplates struct {
	value *WebhooksOwnerRepoTemplates
	isSet bool
}

func (v NullableWebhooksOwnerRepoTemplates) Get() *WebhooksOwnerRepoTemplates {
	return v.value
}

func (v *NullableWebhooksOwnerRepoTemplates) Set(val *WebhooksOwnerRepoTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksOwnerRepoTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksOwnerRepoTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksOwnerRepoTemplates(val *WebhooksOwnerRepoTemplates) *NullableWebhooksOwnerRepoTemplates {
	return &NullableWebhooksOwnerRepoTemplates{value: val, isSet: true}
}

func (v NullableWebhooksOwnerRepoTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksOwnerRepoTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

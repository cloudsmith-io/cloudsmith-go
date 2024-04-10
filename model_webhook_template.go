/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.392.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"encoding/json"
)

// checks if the WebhookTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookTemplate{}

// WebhookTemplate struct for WebhookTemplate
type WebhookTemplate struct {
	Event    string         `json:"event"`
	Template NullableString `json:"template,omitempty"`
}

// NewWebhookTemplate instantiates a new WebhookTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookTemplate(event string) *WebhookTemplate {
	this := WebhookTemplate{}
	this.Event = event
	return &this
}

// NewWebhookTemplateWithDefaults instantiates a new WebhookTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookTemplateWithDefaults() *WebhookTemplate {
	this := WebhookTemplate{}
	return &this
}

// GetEvent returns the Event field value
func (o *WebhookTemplate) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *WebhookTemplate) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *WebhookTemplate) SetEvent(v string) {
	o.Event = v
}

// GetTemplate returns the Template field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookTemplate) GetTemplate() string {
	if o == nil || IsNil(o.Template.Get()) {
		var ret string
		return ret
	}
	return *o.Template.Get()
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookTemplate) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Template.Get(), o.Template.IsSet()
}

// HasTemplate returns a boolean if a field has been set.
func (o *WebhookTemplate) HasTemplate() bool {
	if o != nil && o.Template.IsSet() {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given NullableString and assigns it to the Template field.
func (o *WebhookTemplate) SetTemplate(v string) {
	o.Template.Set(&v)
}

// SetTemplateNil sets the value for Template to be an explicit nil
func (o *WebhookTemplate) SetTemplateNil() {
	o.Template.Set(nil)
}

// UnsetTemplate ensures that no value is present for Template, not even an explicit nil
func (o *WebhookTemplate) UnsetTemplate() {
	o.Template.Unset()
}

func (o WebhookTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if o.Template.IsSet() {
		toSerialize["template"] = o.Template.Get()
	}
	return toSerialize, nil
}

type NullableWebhookTemplate struct {
	value *WebhookTemplate
	isSet bool
}

func (v NullableWebhookTemplate) Get() *WebhookTemplate {
	return v.value
}

func (v *NullableWebhookTemplate) Set(val *WebhookTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTemplate(val *WebhookTemplate) *NullableWebhookTemplate {
	return &NullableWebhookTemplate{value: val, isSet: true}
}

func (v NullableWebhookTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the ServiceRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceRequestPatch{}

// ServiceRequestPatch struct for ServiceRequestPatch
type ServiceRequestPatch struct {
	// The description of the service
	Description *string `json:"description,omitempty"`
	// The name of the service
	Name *string `json:"name,omitempty"`
	// The role of the service.
	Role  *string        `json:"role,omitempty"`
	Teams []ServiceTeams `json:"teams,omitempty"`
}

// NewServiceRequestPatch instantiates a new ServiceRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRequestPatch() *ServiceRequestPatch {
	this := ServiceRequestPatch{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// NewServiceRequestPatchWithDefaults instantiates a new ServiceRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRequestPatchWithDefaults() *ServiceRequestPatch {
	this := ServiceRequestPatch{}
	var role string = "Member"
	this.Role = &role
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceRequestPatch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequestPatch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceRequestPatch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceRequestPatch) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceRequestPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequestPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceRequestPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceRequestPatch) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ServiceRequestPatch) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequestPatch) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ServiceRequestPatch) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ServiceRequestPatch) SetRole(v string) {
	o.Role = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ServiceRequestPatch) GetTeams() []ServiceTeams {
	if o == nil || IsNil(o.Teams) {
		var ret []ServiceTeams
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequestPatch) GetTeamsOk() ([]ServiceTeams, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ServiceRequestPatch) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []ServiceTeams and assigns it to the Teams field.
func (o *ServiceRequestPatch) SetTeams(v []ServiceTeams) {
	o.Teams = v
}

func (o ServiceRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return toSerialize, nil
}

type NullableServiceRequestPatch struct {
	value *ServiceRequestPatch
	isSet bool
}

func (v NullableServiceRequestPatch) Get() *ServiceRequestPatch {
	return v.value
}

func (v *NullableServiceRequestPatch) Set(val *ServiceRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRequestPatch(val *ServiceRequestPatch) *NullableServiceRequestPatch {
	return &NullableServiceRequestPatch{value: val, isSet: true}
}

func (v NullableServiceRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

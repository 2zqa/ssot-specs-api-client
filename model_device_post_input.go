/*
SSOT API

This page describes the api endpoints for single source of truth infra project

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DevicePostInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostInput{}

// DevicePostInput struct for DevicePostInput
type DevicePostInput struct {
	// Version 4 UUID
	Uuid *string `json:"uuid,omitempty"`
	Specs *Specs `json:"specs,omitempty"`
}

// NewDevicePostInput instantiates a new DevicePostInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostInput() *DevicePostInput {
	this := DevicePostInput{}
	return &this
}

// NewDevicePostInputWithDefaults instantiates a new DevicePostInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostInputWithDefaults() *DevicePostInput {
	this := DevicePostInput{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DevicePostInput) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostInput) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DevicePostInput) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DevicePostInput) SetUuid(v string) {
	o.Uuid = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *DevicePostInput) GetSpecs() Specs {
	if o == nil || IsNil(o.Specs) {
		var ret Specs
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostInput) GetSpecsOk() (*Specs, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *DevicePostInput) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given Specs and assigns it to the Specs field.
func (o *DevicePostInput) SetSpecs(v Specs) {
	o.Specs = &v
}

func (o DevicePostInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	return toSerialize, nil
}

type NullableDevicePostInput struct {
	value *DevicePostInput
	isSet bool
}

func (v NullableDevicePostInput) Get() *DevicePostInput {
	return v.value
}

func (v *NullableDevicePostInput) Set(val *DevicePostInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostInput(val *DevicePostInput) *NullableDevicePostInput {
	return &NullableDevicePostInput{value: val, isSet: true}
}

func (v NullableDevicePostInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


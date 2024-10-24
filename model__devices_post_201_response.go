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

// checks if the DevicesPost201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesPost201Response{}

// DevicesPost201Response struct for DevicesPost201Response
type DevicesPost201Response struct {
	Device *Device `json:"device,omitempty"`
}

// NewDevicesPost201Response instantiates a new DevicesPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesPost201Response() *DevicesPost201Response {
	this := DevicesPost201Response{}
	return &this
}

// NewDevicesPost201ResponseWithDefaults instantiates a new DevicesPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesPost201ResponseWithDefaults() *DevicesPost201Response {
	this := DevicesPost201Response{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DevicesPost201Response) GetDevice() Device {
	if o == nil || IsNil(o.Device) {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesPost201Response) GetDeviceOk() (*Device, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DevicesPost201Response) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *DevicesPost201Response) SetDevice(v Device) {
	o.Device = &v
}

func (o DevicesPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesPost201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableDevicesPost201Response struct {
	value *DevicesPost201Response
	isSet bool
}

func (v NullableDevicesPost201Response) Get() *DevicesPost201Response {
	return v.value
}

func (v *NullableDevicesPost201Response) Set(val *DevicesPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesPost201Response(val *DevicesPost201Response) *NullableDevicesPost201Response {
	return &NullableDevicesPost201Response{value: val, isSet: true}
}

func (v NullableDevicesPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



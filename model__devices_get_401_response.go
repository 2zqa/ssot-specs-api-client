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

// checks if the DevicesGet401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesGet401Response{}

// DevicesGet401Response struct for DevicesGet401Response
type DevicesGet401Response struct {
	Error *string `json:"error,omitempty"`
}

// NewDevicesGet401Response instantiates a new DevicesGet401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesGet401Response() *DevicesGet401Response {
	this := DevicesGet401Response{}
	return &this
}

// NewDevicesGet401ResponseWithDefaults instantiates a new DevicesGet401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesGet401ResponseWithDefaults() *DevicesGet401Response {
	this := DevicesGet401Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DevicesGet401Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesGet401Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DevicesGet401Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DevicesGet401Response) SetError(v string) {
	o.Error = &v
}

func (o DevicesGet401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesGet401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableDevicesGet401Response struct {
	value *DevicesGet401Response
	isSet bool
}

func (v NullableDevicesGet401Response) Get() *DevicesGet401Response {
	return v.value
}

func (v *NullableDevicesGet401Response) Set(val *DevicesGet401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesGet401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesGet401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesGet401Response(val *DevicesGet401Response) *NullableDevicesGet401Response {
	return &NullableDevicesGet401Response{value: val, isSet: true}
}

func (v NullableDevicesGet401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesGet401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



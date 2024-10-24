# BIOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | Version as reported by the BIOS | [optional] 
**Date** | Pointer to **string** | Date the BIOS was released in ISO 8601 format | [optional] 

## Methods

### NewBIOS

`func NewBIOS() *BIOS`

NewBIOS instantiates a new BIOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBIOSWithDefaults

`func NewBIOSWithDefaults() *BIOS`

NewBIOSWithDefaults instantiates a new BIOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *BIOS) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BIOS) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BIOS) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *BIOS) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *BIOS) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BIOS) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BIOS) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BIOS) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDate

`func (o *BIOS) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BIOS) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BIOS) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BIOS) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



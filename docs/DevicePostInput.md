# DevicePostInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Version 4 UUID | [optional] 
**Specs** | Pointer to [**Specs**](Specs.md) |  | [optional] 

## Methods

### NewDevicePostInput

`func NewDevicePostInput() *DevicePostInput`

NewDevicePostInput instantiates a new DevicePostInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePostInputWithDefaults

`func NewDevicePostInputWithDefaults() *DevicePostInput`

NewDevicePostInputWithDefaults instantiates a new DevicePostInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DevicePostInput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DevicePostInput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DevicePostInput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DevicePostInput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSpecs

`func (o *DevicePostInput) GetSpecs() Specs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *DevicePostInput) GetSpecsOk() (*Specs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *DevicePostInput) SetSpecs(v Specs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *DevicePostInput) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



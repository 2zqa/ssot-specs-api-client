# Memory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int64** | Total memory installed on device in megabytes | [optional] 
**Swap** | Pointer to **int64** | Total swap memory installed on device in megabytes | [optional] 
**SwapDevices** | Pointer to [**[]SwapDevice**](SwapDevice.md) |  | [optional] 

## Methods

### NewMemory

`func NewMemory() *Memory`

NewMemory instantiates a new Memory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryWithDefaults

`func NewMemoryWithDefaults() *Memory`

NewMemoryWithDefaults instantiates a new Memory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *Memory) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Memory) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Memory) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Memory) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSwap

`func (o *Memory) GetSwap() int64`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *Memory) GetSwapOk() (*int64, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *Memory) SetSwap(v int64)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *Memory) HasSwap() bool`

HasSwap returns a boolean if a field has been set.

### GetSwapDevices

`func (o *Memory) GetSwapDevices() []SwapDevice`

GetSwapDevices returns the SwapDevices field if non-nil, zero value otherwise.

### GetSwapDevicesOk

`func (o *Memory) GetSwapDevicesOk() (*[]SwapDevice, bool)`

GetSwapDevicesOk returns a tuple with the SwapDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapDevices

`func (o *Memory) SetSwapDevices(v []SwapDevice)`

SetSwapDevices sets SwapDevices field to given value.

### HasSwapDevices

`func (o *Memory) HasSwapDevices() bool`

HasSwapDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



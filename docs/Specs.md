# Specs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Motherboard** | Pointer to [**Motherboard**](Motherboard.md) |  | [optional] 
**Cpu** | Pointer to [**CPU**](CPU.md) |  | [optional] 
**Disks** | Pointer to [**[]DisksInner**](DisksInner.md) |  | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**Bios** | Pointer to [**BIOS**](BIOS.md) |  | [optional] 
**Memory** | Pointer to [**Memory**](Memory.md) |  | [optional] 
**Dimms** | Pointer to [**[]DIMMsInner**](DIMMsInner.md) |  | [optional] 
**BootTime** | Pointer to **time.Time** | ISO 8601 formatted date-time string of the time the device booted | [optional] 
**Kernel** | Pointer to [**Kernel**](Kernel.md) |  | [optional] 
**Release** | Pointer to [**Release**](Release.md) |  | [optional] 
**Oem** | Pointer to [**OEM**](OEM.md) |  | [optional] 
**Virtualization** | Pointer to [**Virtualization**](Virtualization.md) |  | [optional] 

## Methods

### NewSpecs

`func NewSpecs() *Specs`

NewSpecs instantiates a new Specs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecsWithDefaults

`func NewSpecsWithDefaults() *Specs`

NewSpecsWithDefaults instantiates a new Specs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMotherboard

`func (o *Specs) GetMotherboard() Motherboard`

GetMotherboard returns the Motherboard field if non-nil, zero value otherwise.

### GetMotherboardOk

`func (o *Specs) GetMotherboardOk() (*Motherboard, bool)`

GetMotherboardOk returns a tuple with the Motherboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherboard

`func (o *Specs) SetMotherboard(v Motherboard)`

SetMotherboard sets Motherboard field to given value.

### HasMotherboard

`func (o *Specs) HasMotherboard() bool`

HasMotherboard returns a boolean if a field has been set.

### GetCpu

`func (o *Specs) GetCpu() CPU`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Specs) GetCpuOk() (*CPU, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Specs) SetCpu(v CPU)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Specs) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisks

`func (o *Specs) GetDisks() []DisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Specs) GetDisksOk() (*[]DisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Specs) SetDisks(v []DisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Specs) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetNetwork

`func (o *Specs) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Specs) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Specs) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Specs) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetBios

`func (o *Specs) GetBios() BIOS`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *Specs) GetBiosOk() (*BIOS, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *Specs) SetBios(v BIOS)`

SetBios sets Bios field to given value.

### HasBios

`func (o *Specs) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetMemory

`func (o *Specs) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Specs) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Specs) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Specs) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDimms

`func (o *Specs) GetDimms() []DIMMsInner`

GetDimms returns the Dimms field if non-nil, zero value otherwise.

### GetDimmsOk

`func (o *Specs) GetDimmsOk() (*[]DIMMsInner, bool)`

GetDimmsOk returns a tuple with the Dimms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimms

`func (o *Specs) SetDimms(v []DIMMsInner)`

SetDimms sets Dimms field to given value.

### HasDimms

`func (o *Specs) HasDimms() bool`

HasDimms returns a boolean if a field has been set.

### GetBootTime

`func (o *Specs) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *Specs) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *Specs) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *Specs) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetKernel

`func (o *Specs) GetKernel() Kernel`

GetKernel returns the Kernel field if non-nil, zero value otherwise.

### GetKernelOk

`func (o *Specs) GetKernelOk() (*Kernel, bool)`

GetKernelOk returns a tuple with the Kernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernel

`func (o *Specs) SetKernel(v Kernel)`

SetKernel sets Kernel field to given value.

### HasKernel

`func (o *Specs) HasKernel() bool`

HasKernel returns a boolean if a field has been set.

### GetRelease

`func (o *Specs) GetRelease() Release`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *Specs) GetReleaseOk() (*Release, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *Specs) SetRelease(v Release)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *Specs) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetOem

`func (o *Specs) GetOem() OEM`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *Specs) GetOemOk() (*OEM, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *Specs) SetOem(v OEM)`

SetOem sets Oem field to given value.

### HasOem

`func (o *Specs) HasOem() bool`

HasOem returns a boolean if a field has been set.

### GetVirtualization

`func (o *Specs) GetVirtualization() Virtualization`

GetVirtualization returns the Virtualization field if non-nil, zero value otherwise.

### GetVirtualizationOk

`func (o *Specs) GetVirtualizationOk() (*Virtualization, bool)`

GetVirtualizationOk returns a tuple with the Virtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualization

`func (o *Specs) SetVirtualization(v Virtualization)`

SetVirtualization sets Virtualization field to given value.

### HasVirtualization

`func (o *Specs) HasVirtualization() bool`

HasVirtualization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Model name | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**CoreCount** | Pointer to **int64** |  | [optional] 
**CpuCount** | Pointer to **int64** |  | [optional] 
**MaxFrequencyMegahertz** | Pointer to **int64** | Max clock frequency in MHz | [optional] 
**Mitigations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCPU

`func NewCPU() *CPU`

NewCPU instantiates a new CPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCPUWithDefaults

`func NewCPUWithDefaults() *CPU`

NewCPUWithDefaults instantiates a new CPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CPU) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CPU) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CPU) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CPU) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchitecture

`func (o *CPU) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *CPU) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *CPU) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *CPU) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetCoreCount

`func (o *CPU) GetCoreCount() int64`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *CPU) GetCoreCountOk() (*int64, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *CPU) SetCoreCount(v int64)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *CPU) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetCpuCount

`func (o *CPU) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *CPU) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *CPU) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *CPU) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMaxFrequencyMegahertz

`func (o *CPU) GetMaxFrequencyMegahertz() int64`

GetMaxFrequencyMegahertz returns the MaxFrequencyMegahertz field if non-nil, zero value otherwise.

### GetMaxFrequencyMegahertzOk

`func (o *CPU) GetMaxFrequencyMegahertzOk() (*int64, bool)`

GetMaxFrequencyMegahertzOk returns a tuple with the MaxFrequencyMegahertz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFrequencyMegahertz

`func (o *CPU) SetMaxFrequencyMegahertz(v int64)`

SetMaxFrequencyMegahertz sets MaxFrequencyMegahertz field to given value.

### HasMaxFrequencyMegahertz

`func (o *CPU) HasMaxFrequencyMegahertz() bool`

HasMaxFrequencyMegahertz returns a boolean if a field has been set.

### GetMitigations

`func (o *CPU) GetMitigations() []string`

GetMitigations returns the Mitigations field if non-nil, zero value otherwise.

### GetMitigationsOk

`func (o *CPU) GetMitigationsOk() (*[]string, bool)`

GetMitigationsOk returns a tuple with the Mitigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigations

`func (o *CPU) SetMitigations(v []string)`

SetMitigations sets Mitigations field to given value.

### HasMitigations

`func (o *CPU) HasMitigations() bool`

HasMitigations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



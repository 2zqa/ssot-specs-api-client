# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filesystem** | Pointer to **string** |  | [optional] 
**CapacityMegabytes** | Pointer to **int64** | Amount of megabytes the partition can store | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewPartition

`func NewPartition() *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystem

`func (o *Partition) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *Partition) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *Partition) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *Partition) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetCapacityMegabytes

`func (o *Partition) GetCapacityMegabytes() int64`

GetCapacityMegabytes returns the CapacityMegabytes field if non-nil, zero value otherwise.

### GetCapacityMegabytesOk

`func (o *Partition) GetCapacityMegabytesOk() (*int64, bool)`

GetCapacityMegabytesOk returns a tuple with the CapacityMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMegabytes

`func (o *Partition) SetCapacityMegabytes(v int64)`

SetCapacityMegabytes sets CapacityMegabytes field to given value.

### HasCapacityMegabytes

`func (o *Partition) HasCapacityMegabytes() bool`

HasCapacityMegabytes returns a boolean if a field has been set.

### GetSource

`func (o *Partition) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Partition) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Partition) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Partition) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *Partition) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Partition) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Partition) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Partition) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



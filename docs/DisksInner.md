# DisksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Model name | [optional] 
**SizeMegabytes** | Pointer to **int64** | Size in megabytes | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) |  | [optional] 

## Methods

### NewDisksInner

`func NewDisksInner() *DisksInner`

NewDisksInner instantiates a new DisksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisksInnerWithDefaults

`func NewDisksInnerWithDefaults() *DisksInner`

NewDisksInnerWithDefaults instantiates a new DisksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DisksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DisksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DisksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DisksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSizeMegabytes

`func (o *DisksInner) GetSizeMegabytes() int64`

GetSizeMegabytes returns the SizeMegabytes field if non-nil, zero value otherwise.

### GetSizeMegabytesOk

`func (o *DisksInner) GetSizeMegabytesOk() (*int64, bool)`

GetSizeMegabytesOk returns a tuple with the SizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMegabytes

`func (o *DisksInner) SetSizeMegabytes(v int64)`

SetSizeMegabytes sets SizeMegabytes field to given value.

### HasSizeMegabytes

`func (o *DisksInner) HasSizeMegabytes() bool`

HasSizeMegabytes returns a boolean if a field has been set.

### GetPartitions

`func (o *DisksInner) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *DisksInner) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *DisksInner) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *DisksInner) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



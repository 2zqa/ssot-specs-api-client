# DIMMsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGigabytes** | Pointer to **int64** | Size of the DIMM in gigabytes | [optional] 
**SpeedMtS** | Pointer to **int64** | Speed of the DIMM in megatransfers per second | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the DIMM | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the DIMM | [optional] 
**Type** | Pointer to **string** | Type of the DIMM | [optional] 
**PartNumber** | Pointer to **string** | Part number of the DIMM | [optional] 
**FormFactor** | Pointer to **string** | Form factor of the DIMM | [optional] 
**Locator** | Pointer to **string** | Locator of the DIMM | [optional] 
**BankLocator** | Pointer to **string** | Bank locator of the DIMM | [optional] 

## Methods

### NewDIMMsInner

`func NewDIMMsInner() *DIMMsInner`

NewDIMMsInner instantiates a new DIMMsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDIMMsInnerWithDefaults

`func NewDIMMsInnerWithDefaults() *DIMMsInner`

NewDIMMsInnerWithDefaults instantiates a new DIMMsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGigabytes

`func (o *DIMMsInner) GetSizeGigabytes() int64`

GetSizeGigabytes returns the SizeGigabytes field if non-nil, zero value otherwise.

### GetSizeGigabytesOk

`func (o *DIMMsInner) GetSizeGigabytesOk() (*int64, bool)`

GetSizeGigabytesOk returns a tuple with the SizeGigabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGigabytes

`func (o *DIMMsInner) SetSizeGigabytes(v int64)`

SetSizeGigabytes sets SizeGigabytes field to given value.

### HasSizeGigabytes

`func (o *DIMMsInner) HasSizeGigabytes() bool`

HasSizeGigabytes returns a boolean if a field has been set.

### GetSpeedMtS

`func (o *DIMMsInner) GetSpeedMtS() int64`

GetSpeedMtS returns the SpeedMtS field if non-nil, zero value otherwise.

### GetSpeedMtSOk

`func (o *DIMMsInner) GetSpeedMtSOk() (*int64, bool)`

GetSpeedMtSOk returns a tuple with the SpeedMtS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMtS

`func (o *DIMMsInner) SetSpeedMtS(v int64)`

SetSpeedMtS sets SpeedMtS field to given value.

### HasSpeedMtS

`func (o *DIMMsInner) HasSpeedMtS() bool`

HasSpeedMtS returns a boolean if a field has been set.

### GetManufacturer

`func (o *DIMMsInner) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DIMMsInner) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DIMMsInner) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DIMMsInner) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DIMMsInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DIMMsInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DIMMsInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DIMMsInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetType

`func (o *DIMMsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DIMMsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DIMMsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DIMMsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPartNumber

`func (o *DIMMsInner) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *DIMMsInner) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *DIMMsInner) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *DIMMsInner) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetFormFactor

`func (o *DIMMsInner) GetFormFactor() string`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *DIMMsInner) GetFormFactorOk() (*string, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *DIMMsInner) SetFormFactor(v string)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *DIMMsInner) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetLocator

`func (o *DIMMsInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *DIMMsInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *DIMMsInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *DIMMsInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetBankLocator

`func (o *DIMMsInner) GetBankLocator() string`

GetBankLocator returns the BankLocator field if non-nil, zero value otherwise.

### GetBankLocatorOk

`func (o *DIMMsInner) GetBankLocatorOk() (*string, bool)`

GetBankLocatorOk returns a tuple with the BankLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankLocator

`func (o *DIMMsInner) SetBankLocator(v string)`

SetBankLocator sets BankLocator field to given value.

### HasBankLocator

`func (o *DIMMsInner) HasBankLocator() bool`

HasBankLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



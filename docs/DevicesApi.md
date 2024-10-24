# \DevicesApi

All URIs are relative to *http://localhost:4000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesGet**](DevicesApi.md#DevicesGet) | **Get** /devices | Get all devices
[**DevicesPost**](DevicesApi.md#DevicesPost) | **Post** /devices | Create a device
[**DevicesUuidDelete**](DevicesApi.md#DevicesUuidDelete) | **Delete** /devices/{uuid} | Delete a device by uuid
[**DevicesUuidGet**](DevicesApi.md#DevicesUuidGet) | **Get** /devices/{uuid} | Find device by uuid
[**DevicesUuidPut**](DevicesApi.md#DevicesUuidPut) | **Put** /devices/{uuid} | Update or create a device by uuid



## DevicesGet

> DevicesGet200Response DevicesGet(ctx).Q(q).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get all devices

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/2zqa/ssot-specs-api-client"
)

func main() {
    q := "q_example" // string | Search query (optional)
    page := int32(56) // int32 | Page number (optional) (default to 1)
    pageSize := int32(56) // int32 | Page size (optional) (default to 10)
    sort := "sort_example" // string | Sort by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesGet(context.Background()).Q(q).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGet`: DevicesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 
 **page** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Page size | [default to 10]
 **sort** | **string** | Sort by | 

### Return type

[**DevicesGet200Response**](DevicesGet200Response.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesPost

> DevicesPost201Response DevicesPost(ctx).DevicePostInput(devicePostInput).Execute()

Create a device

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/2zqa/ssot-specs-api-client"
)

func main() {
    devicePostInput := *openapiclient.NewDevicePostInput() // DevicePostInput | The device to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesPost(context.Background()).DevicePostInput(devicePostInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesPost`: DevicesPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devicePostInput** | [**DevicePostInput**](DevicePostInput.md) | The device to create | 

### Return type

[**DevicesPost201Response**](DevicesPost201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUuidDelete

> DevicesUuidDelete(ctx, uuid).Execute()

Delete a device by uuid



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/2zqa/ssot-specs-api-client"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of device to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesApi.DevicesUuidDelete(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesUuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of device to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUuidGet

> DeviceResponse DevicesUuidGet(ctx, uuid).Execute()

Find device by uuid



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/2zqa/ssot-specs-api-client"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of device to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesUuidGet(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesUuidGet`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of device to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUuidPut

> DeviceResponse DevicesUuidPut(ctx, uuid).DevicePutInput(devicePutInput).Execute()

Update or create a device by uuid

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/2zqa/ssot-specs-api-client"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of device to update or create
    devicePutInput := *openapiclient.NewDevicePutInput() // DevicePutInput | The device to create or to update the specified device with

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesUuidPut(context.Background(), uuid).DevicePutInput(devicePutInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesUuidPut`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of device to update or create | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicePutInput** | [**DevicePutInput**](DevicePutInput.md) | The device to create or to update the specified device with | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


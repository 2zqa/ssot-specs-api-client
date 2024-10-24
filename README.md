# Go API client for openapi

This page describes the api endpoints for single source of truth infra project

> [!NOTE]
> This project is part of a suite of projects that work together. For all other related projects, see this search query: [`owner:2zqa topic:ssot`](https://github.com/search?q=owner%3A2zqa+topic%3Assot&type=repositories)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/2zqa/ssot-specs-api-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4000/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DevicesApi* | [**DevicesGet**](docs/DevicesApi.md#devicesget) | **Get** /devices | Get all devices
*DevicesApi* | [**DevicesPost**](docs/DevicesApi.md#devicespost) | **Post** /devices | Create a device
*DevicesApi* | [**DevicesUuidDelete**](docs/DevicesApi.md#devicesuuiddelete) | **Delete** /devices/{uuid} | Delete a device by uuid
*DevicesApi* | [**DevicesUuidGet**](docs/DevicesApi.md#devicesuuidget) | **Get** /devices/{uuid} | Find device by uuid
*DevicesApi* | [**DevicesUuidPut**](docs/DevicesApi.md#devicesuuidput) | **Put** /devices/{uuid} | Update or create a device by uuid
*ServerApi* | [**HealthcheckGet**](docs/ServerApi.md#healthcheckget) | **Get** /healthcheck | Check the health of the API


## Documentation For Models

 - [BIOS](docs/BIOS.md)
 - [CPU](docs/CPU.md)
 - [DIMMsInner](docs/DIMMsInner.md)
 - [Device](docs/Device.md)
 - [DevicePostInput](docs/DevicePostInput.md)
 - [DevicePutInput](docs/DevicePutInput.md)
 - [DeviceResponse](docs/DeviceResponse.md)
 - [DevicesGet200Response](docs/DevicesGet200Response.md)
 - [DevicesGet401Response](docs/DevicesGet401Response.md)
 - [DevicesPost201Response](docs/DevicesPost201Response.md)
 - [DisksInner](docs/DisksInner.md)
 - [Driver](docs/Driver.md)
 - [Kernel](docs/Kernel.md)
 - [Memory](docs/Memory.md)
 - [Metadata](docs/Metadata.md)
 - [Motherboard](docs/Motherboard.md)
 - [Network](docs/Network.md)
 - [NetworkInterface](docs/NetworkInterface.md)
 - [OEM](docs/OEM.md)
 - [Partition](docs/Partition.md)
 - [Release](docs/Release.md)
 - [ServerInfo](docs/ServerInfo.md)
 - [ServerInfoSystemInfo](docs/ServerInfoSystemInfo.md)
 - [Specs](docs/Specs.md)
 - [SwapDevice](docs/SwapDevice.md)
 - [Virtualization](docs/Virtualization.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Key and passed in as the auth context for each request.


### JwtAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
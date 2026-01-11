# \SandboxInformationAPIAPI

All URIs are relative to *http://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSandboxUsingPOST**](SandboxInformationAPIAPI.md#CreateSandboxUsingPOST) | **Post** /appsec/v1/applications/{applicationGuid}/sandboxes | createSandbox
[**GetSandboxUsingGET**](SandboxInformationAPIAPI.md#GetSandboxUsingGET) | **Get** /appsec/v1/applications/{applicationGuid}/sandboxes/{sandboxGuid} | getSandbox
[**GetSandboxesUsingGET**](SandboxInformationAPIAPI.md#GetSandboxesUsingGET) | **Get** /appsec/v1/applications/{applicationGuid}/sandboxes | getSandboxes



## CreateSandboxUsingPOST

> Sandbox CreateSandboxUsingPOST(ctx, applicationGuid).Sandbox(sandbox).Execute()

createSandbox



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationGuid := "applicationGuid_example" // string | The application GUID.
	sandbox := *openapiclient.NewSandbox() // Sandbox | The sandbox object to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxInformationAPIAPI.CreateSandboxUsingPOST(context.Background(), applicationGuid).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxInformationAPIAPI.CreateSandboxUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSandboxUsingPOST`: Sandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxInformationAPIAPI.CreateSandboxUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandbox** | [**Sandbox**](Sandbox.md) | The sandbox object to create. | 

### Return type

[**Sandbox**](Sandbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxUsingGET

> Sandbox GetSandboxUsingGET(ctx, applicationGuid, sandboxGuid).Execute()

getSandbox



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationGuid := "applicationGuid_example" // string | The application GUID.
	sandboxGuid := "sandboxGuid_example" // string | The sandbox GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxInformationAPIAPI.GetSandboxUsingGET(context.Background(), applicationGuid, sandboxGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxInformationAPIAPI.GetSandboxUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandboxUsingGET`: Sandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxInformationAPIAPI.GetSandboxUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application GUID. | 
**sandboxGuid** | **string** | The sandbox GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sandbox**](Sandbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxesUsingGET

> PagedResourceOfSandbox GetSandboxesUsingGET(ctx, applicationGuid).Page(page).Size(size).Execute()

getSandboxes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationGuid := "applicationGuid_example" // string | The application GUID.
	page := int32(56) // int32 | Page number. Defaults to 0. (optional)
	size := int32(56) // int32 | Page size, up to 500. The default is 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxInformationAPIAPI.GetSandboxesUsingGET(context.Background(), applicationGuid).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxInformationAPIAPI.GetSandboxesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandboxesUsingGET`: PagedResourceOfSandbox
	fmt.Fprintf(os.Stdout, "Response from `SandboxInformationAPIAPI.GetSandboxesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Defaults to 0. | 
 **size** | **int32** | Page size, up to 500. The default is 50. | 

### Return type

[**PagedResourceOfSandbox**](PagedResourceOfSandbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


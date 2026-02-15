# \ManualScansInformationAPI

All URIs are relative to *https://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetManualFinding**](ManualScansInformationAPI.md#GetManualFinding) | **Get** /mpt/v1/findings/{finding} | getFinding
[**GetManualScan**](ManualScansInformationAPI.md#GetManualScan) | **Get** /mpt/v1/scans/{scan} | get a Manual Scan
[**GetManualScans**](ManualScansInformationAPI.md#GetManualScans) | **Get** /mpt/v1/scans | Get Manual Scans
[**GetScanFindings**](ManualScansInformationAPI.md#GetScanFindings) | **Get** /mpt/v1/scans/{scan_id}/findings | getFindings



## GetManualFinding

> ManualFinding GetManualFinding(ctx, finding).IncludeArtifacts(includeArtifacts).Execute()

getFinding



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
	finding := int64(789) // int64 | The manual finding identifier.
	includeArtifacts := true // bool | Include the artifacts, such as code samples and screenshots, in the manual finding response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManualScansInformationAPI.GetManualFinding(context.Background(), finding).IncludeArtifacts(includeArtifacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualScansInformationAPI.GetManualFinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualFinding`: ManualFinding
	fmt.Fprintf(os.Stdout, "Response from `ManualScansInformationAPI.GetManualFinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**finding** | **int64** | The manual finding identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualFindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeArtifacts** | **bool** | Include the artifacts, such as code samples and screenshots, in the manual finding response. | 

### Return type

[**ManualFinding**](ManualFinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualScan

> Scan GetManualScan(ctx, scan).Execute()

get a Manual Scan



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
	scan := int32(56) // int32 | The manual scan ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManualScansInformationAPI.GetManualScan(context.Background(), scan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualScansInformationAPI.GetManualScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualScan`: Scan
	fmt.Fprintf(os.Stdout, "Response from `ManualScansInformationAPI.GetManualScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scan** | **int32** | The manual scan ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Scan**](Scan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualScans

> PagedModelOfScan GetManualScans(ctx).Application(application).Page(page).Size(size).Execute()

Get Manual Scans



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
	application := "application_example" // string | The application GUID.
	page := int32(56) // int32 | The page number, which defaults to zero. (optional)
	size := int32(56) // int32 | The page size (0-500). Defaults to 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManualScansInformationAPI.GetManualScans(context.Background()).Application(application).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualScansInformationAPI.GetManualScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualScans`: PagedModelOfScan
	fmt.Fprintf(os.Stdout, "Response from `ManualScansInformationAPI.GetManualScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManualScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | The application GUID. | 
 **page** | **int32** | The page number, which defaults to zero. | 
 **size** | **int32** | The page size (0-500). Defaults to 100. | 

### Return type

[**PagedModelOfScan**](PagedModelOfScan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanFindings

> PagedModelOfManualFinding GetScanFindings(ctx, scanId).IncludeArtifacts(includeArtifacts).Execute()

getFindings



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
	scanId := int32(56) // int32 | The manual scan identifier.
	includeArtifacts := true // bool | Include the artifacts, such as code samples and screenshots, in the manual finding response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManualScansInformationAPI.GetScanFindings(context.Background(), scanId).IncludeArtifacts(includeArtifacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualScansInformationAPI.GetScanFindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScanFindings`: PagedModelOfManualFinding
	fmt.Fprintf(os.Stdout, "Response from `ManualScansInformationAPI.GetScanFindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanId** | **int32** | The manual scan identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScanFindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeArtifacts** | **bool** | Include the artifacts, such as code samples and screenshots, in the manual finding response. | 

### Return type

[**PagedModelOfManualFinding**](PagedModelOfManualFinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


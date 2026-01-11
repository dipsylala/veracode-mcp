# \CWEAPIInformationAPI

All URIs are relative to *https://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCategoriesUsingGET**](CWEAPIInformationAPI.md#GetCategoriesUsingGET) | **Get** /appsec/v1/categories | getCategories
[**GetCategoryUsingGET**](CWEAPIInformationAPI.md#GetCategoryUsingGET) | **Get** /appsec/v1/categories/{category} | getCategory
[**GetCweUsingGET**](CWEAPIInformationAPI.md#GetCweUsingGET) | **Get** /appsec/v1/cwes/{cwe} | getCwe
[**GetCwesUsingGET**](CWEAPIInformationAPI.md#GetCwesUsingGET) | **Get** /appsec/v1/cwes | getCwes



## GetCategoriesUsingGET

> PagedResourceOfCategory GetCategoriesUsingGET(ctx).Page(page).Size(size).Execute()

getCategories



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
	page := int32(56) // int32 | Page number. The default is 0. (optional)
	size := int32(56) // int32 | Page size (0-500). The default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CWEAPIInformationAPI.GetCategoriesUsingGET(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CWEAPIInformationAPI.GetCategoriesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoriesUsingGET`: PagedResourceOfCategory
	fmt.Fprintf(os.Stdout, "Response from `CWEAPIInformationAPI.GetCategoriesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default is 0. | 
 **size** | **int32** | Page size (0-500). The default is 100. | 

### Return type

[**PagedResourceOfCategory**](PagedResourceOfCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryUsingGET

> Category GetCategoryUsingGET(ctx, category).Execute()

getCategory



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
	category := int32(56) // int32 | The CWE category identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CWEAPIInformationAPI.GetCategoryUsingGET(context.Background(), category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CWEAPIInformationAPI.GetCategoryUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryUsingGET`: Category
	fmt.Fprintf(os.Stdout, "Response from `CWEAPIInformationAPI.GetCategoryUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **int32** | The CWE category identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCweUsingGET

> CweDetail GetCweUsingGET(ctx, cwe).Execute()

getCwe



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
	cwe := int32(56) // int32 | The CWE ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CWEAPIInformationAPI.GetCweUsingGET(context.Background(), cwe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CWEAPIInformationAPI.GetCweUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCweUsingGET`: CweDetail
	fmt.Fprintf(os.Stdout, "Response from `CWEAPIInformationAPI.GetCweUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cwe** | **int32** | The CWE ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCweUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CweDetail**](CweDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCwesUsingGET

> PagedResourceOfCwe GetCwesUsingGET(ctx).Page(page).Size(size).Execute()

getCwes



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
	page := int32(56) // int32 | Page number. The default is 0. (optional)
	size := int32(56) // int32 | Page size (0-500). The default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CWEAPIInformationAPI.GetCwesUsingGET(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CWEAPIInformationAPI.GetCwesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCwesUsingGET`: PagedResourceOfCwe
	fmt.Fprintf(os.Stdout, "Response from `CWEAPIInformationAPI.GetCwesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCwesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default is 0. | 
 **size** | **int32** | Page size (0-500). The default is 100. | 

### Return type

[**PagedResourceOfCwe**](PagedResourceOfCwe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SCAComponentLicenseInformationAPIAPI

All URIs are relative to *http://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScaLicensesUsingGET**](SCAComponentLicenseInformationAPIAPI.md#GetScaLicensesUsingGET) | **Get** /appsec/v1/policy_licenselist | Get a list of component licenses for SCA.



## GetScaLicensesUsingGET

> PagedResourceOfScaLicenseSummary GetScaLicensesUsingGET(ctx).Page(page).Size(size).Sort(sort).Execute()

Get a list of component licenses for SCA.



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
	page := int32(56) // int32 |  (optional)
	size := int32(56) // int32 |  (optional)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SCAComponentLicenseInformationAPIAPI.GetScaLicensesUsingGET(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCAComponentLicenseInformationAPIAPI.GetScaLicensesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScaLicensesUsingGET`: PagedResourceOfScaLicenseSummary
	fmt.Fprintf(os.Stdout, "Response from `SCAComponentLicenseInformationAPIAPI.GetScaLicensesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScaLicensesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **size** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**PagedResourceOfScaLicenseSummary**](PagedResourceOfScaLicenseSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


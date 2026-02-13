# \ApplicationInformationAPIAPI

All URIs are relative to *http://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationUsingPOST**](ApplicationInformationAPIAPI.md#CreateApplicationUsingPOST) | **Post** /appsec/v1/applications | createApplication
[**DeleteApplicationUsingDELETE**](ApplicationInformationAPIAPI.md#DeleteApplicationUsingDELETE) | **Delete** /appsec/v1/applications/{applicationGuid} | deleteApplication
[**GetApplicationUsingGET**](ApplicationInformationAPIAPI.md#GetApplicationUsingGET) | **Get** /appsec/v1/applications/{applicationGuid} | getApplication
[**GetApplicationsUsingGET**](ApplicationInformationAPIAPI.md#GetApplicationsUsingGET) | **Get** /appsec/v1/applications | getApplications
[**UpdateApplicationUsingPUT**](ApplicationInformationAPIAPI.md#UpdateApplicationUsingPUT) | **Put** /appsec/v1/applications/{applicationGuid} | Updates an application



## CreateApplicationUsingPOST

> Application CreateApplicationUsingPOST(ctx).Application(application).Execute()

createApplication



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
	application := *openapiclient.NewApplication() // Application | The application object to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInformationAPIAPI.CreateApplicationUsingPOST(context.Background()).Application(application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInformationAPIAPI.CreateApplicationUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationUsingPOST`: Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInformationAPIAPI.CreateApplicationUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) | The application object to be created. | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationUsingDELETE

> map[string]interface{} DeleteApplicationUsingDELETE(ctx, applicationGuid).Execute()

deleteApplication

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInformationAPIAPI.DeleteApplicationUsingDELETE(context.Background(), applicationGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInformationAPIAPI.DeleteApplicationUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationUsingDELETE`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInformationAPIAPI.DeleteApplicationUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationUsingGET

> Application GetApplicationUsingGET(ctx, applicationGuid).Execute()

getApplication



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInformationAPIAPI.GetApplicationUsingGET(context.Background(), applicationGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInformationAPIAPI.GetApplicationUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationUsingGET`: Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInformationAPIAPI.GetApplicationUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsUsingGET

> PagedResourceOfApplication GetApplicationsUsingGET(ctx).BusinessUnit(businessUnit).CustomFieldNames(customFieldNames).CustomFieldValues(customFieldValues).LegacyId(legacyId).ModifiedAfter(modifiedAfter).Name(name).Page(page).Policy(policy).PolicyCompliance(policyCompliance).PolicyComplianceCheckedAfter(policyComplianceCheckedAfter).PolicyGuid(policyGuid).ScanStatus(scanStatus).ScanType(scanType).Size(size).SortByCustomFieldName(sortByCustomFieldName).Tag(tag).Team(team).Execute()

getApplications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	businessUnit := "businessUnit_example" // string | Business unit. (optional)
	customFieldNames := []string{"Inner_example"} // []string | Custom field names to search. (optional)
	customFieldValues := []string{"Inner_example"} // []string | Custom field values to search. (optional)
	legacyId := int32(56) // int32 | The unique identifier of the Veracode Platform application. (optional)
	modifiedAfter := time.Now() // time.Time | Filter the results to return only those modified after this date. If there are multiple results for the same application, only the last modified result is returned. Format: yyyy-MM-dd (optional)
	name := "name_example" // string | Application name. URL-encode any special characters. (optional)
	page := int32(56) // int32 | Page number. Defaults to 0. (optional)
	policy := "policy_example" // string | policy (optional)
	policyCompliance := "policyCompliance_example" // string | The policy compliance status. (optional)
	policyComplianceCheckedAfter := time.Now() // time.Time | Filter the results to return only those with policy compliance checked after this date. Format: yyyy-MM-dd (optional)
	policyGuid := "policyGuid_example" // string | Policy GUID of the policy to change. (optional)
	scanStatus := []string{"ScanStatus_example"} // []string | The scan status of the application. (optional)
	scanType := "scanType_example" // string | The scan type of the application scans. (optional)
	size := int32(56) // int32 | Page size, up to 500. The default is 50. (optional)
	sortByCustomFieldName := "sortByCustomFieldName_example" // string | Custom field name on which to sort. (optional)
	tag := "tag_example" // string | tag (optional)
	team := "team_example" // string | Filter the results by team name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInformationAPIAPI.GetApplicationsUsingGET(context.Background()).BusinessUnit(businessUnit).CustomFieldNames(customFieldNames).CustomFieldValues(customFieldValues).LegacyId(legacyId).ModifiedAfter(modifiedAfter).Name(name).Page(page).Policy(policy).PolicyCompliance(policyCompliance).PolicyComplianceCheckedAfter(policyComplianceCheckedAfter).PolicyGuid(policyGuid).ScanStatus(scanStatus).ScanType(scanType).Size(size).SortByCustomFieldName(sortByCustomFieldName).Tag(tag).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInformationAPIAPI.GetApplicationsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationsUsingGET`: PagedResourceOfApplication
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInformationAPIAPI.GetApplicationsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessUnit** | **string** | Business unit. | 
 **customFieldNames** | **[]string** | Custom field names to search. | 
 **customFieldValues** | **[]string** | Custom field values to search. | 
 **legacyId** | **int32** | The unique identifier of the Veracode Platform application. | 
 **modifiedAfter** | **time.Time** | Filter the results to return only those modified after this date. If there are multiple results for the same application, only the last modified result is returned. Format: yyyy-MM-dd | 
 **name** | **string** | Application name. URL-encode any special characters. | 
 **page** | **int32** | Page number. Defaults to 0. | 
 **policy** | **string** | policy | 
 **policyCompliance** | **string** | The policy compliance status. | 
 **policyComplianceCheckedAfter** | **time.Time** | Filter the results to return only those with policy compliance checked after this date. Format: yyyy-MM-dd | 
 **policyGuid** | **string** | Policy GUID of the policy to change. | 
 **scanStatus** | **[]string** | The scan status of the application. | 
 **scanType** | **string** | The scan type of the application scans. | 
 **size** | **int32** | Page size, up to 500. The default is 50. | 
 **sortByCustomFieldName** | **string** | Custom field name on which to sort. | 
 **tag** | **string** | tag | 
 **team** | **string** | Filter the results by team name. | 

### Return type

[**PagedResourceOfApplication**](PagedResourceOfApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationUsingPUT

> Application UpdateApplicationUsingPUT(ctx, applicationGuid).Application(application).Method(method).PolicyGuid(policyGuid).Execute()

Updates an application

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
	application := *openapiclient.NewApplication() // Application | The Application object being updated.
	method := "method_example" // string | This method performs a partial update of any custom policy data. (optional)
	policyGuid := "policyGuid_example" // string | Policy GUID of the policy to be changed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInformationAPIAPI.UpdateApplicationUsingPUT(context.Background(), applicationGuid).Application(application).Method(method).PolicyGuid(policyGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInformationAPIAPI.UpdateApplicationUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationUsingPUT`: Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInformationAPIAPI.UpdateApplicationUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | [**Application**](Application.md) | The Application object being updated. | 
 **method** | **string** | This method performs a partial update of any custom policy data. | 
 **policyGuid** | **string** | Policy GUID of the policy to be changed. | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


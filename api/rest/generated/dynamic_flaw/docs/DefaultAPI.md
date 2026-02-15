# \DefaultAPI

All URIs are relative to *https://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet**](DefaultAPI.md#AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet) | **Get** /appsec/v2/applications/{app_guid}/findings/{issue_id}/dynamic_flaw_info | Returns information on a specific dynamic flaw.



## AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet

> DynamicFlaw AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet(ctx, appGuid, issueId).Execute()

Returns information on a specific dynamic flaw.



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
	appGuid := "appGuid_example" // string | GUID for the application.
	issueId := "issueId_example" // string | Unique issue ID for the scanned application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet(context.Background(), appGuid, issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet`: DynamicFlaw
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appGuid** | **string** | GUID for the application. | 
**issueId** | **string** | Unique issue ID for the scanned application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DynamicFlaw**](DynamicFlaw.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


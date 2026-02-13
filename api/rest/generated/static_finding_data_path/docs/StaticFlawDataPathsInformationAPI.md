# \StaticFlawDataPathsInformationAPI

All URIs are relative to *https://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet**](StaticFlawDataPathsInformationAPI.md#AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet) | **Get** /appsec/v2/applications/{app_guid}/findings/{issue_id}/static_flaw_info | Returns information on the data path for a static analysis finding.



## AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet

> StaticFlaws AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet(ctx, appGuid, issueId).Context(context).Execute()

Returns information on the data path for a static analysis finding.



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
	issueId := "issueId_example" // string | Flaw ID of the finding.
	context := "context_example" // string | If specified, the ID of the development sandbox to which this finding belongs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticFlawDataPathsInformationAPI.AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet(context.Background(), appGuid, issueId).Context(context).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticFlawDataPathsInformationAPI.AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet`: StaticFlaws
	fmt.Fprintf(os.Stdout, "Response from `StaticFlawDataPathsInformationAPI.AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appGuid** | **string** | GUID for the application. | 
**issueId** | **string** | Flaw ID of the finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **context** | **string** | If specified, the ID of the development sandbox to which this finding belongs. | 

### Return type

[**StaticFlaws**](StaticFlaws.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


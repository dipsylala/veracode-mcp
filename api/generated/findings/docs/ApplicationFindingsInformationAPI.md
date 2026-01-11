# \ApplicationFindingsInformationAPI

All URIs are relative to *https://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFindingsUsingGET**](ApplicationFindingsInformationAPI.md#GetFindingsUsingGET) | **Get** /appsec/v2/applications/{application_guid}/findings | getFindings



## GetFindingsUsingGET

> PagedResourceOfFinding GetFindingsUsingGET(ctx, applicationGuid).Context(context).Cve(cve).Cvss(cvss).CvssGte(cvssGte).Cwe(cwe).FindingCategory(findingCategory).IncludeAnnot(includeAnnot).IncludeExpDate(includeExpDate).MitigatedAfter(mitigatedAfter).New(new).ScaDepMode(scaDepMode).ScaScanMode(scaScanMode).ScanType(scanType).Severity(severity).SeverityGte(severityGte).ViolatesPolicy(violatesPolicy).Execute()

getFindings



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
	applicationGuid := "applicationGuid_example" // string | The application identifier.
	context := "context_example" // string | Context type, which filters results to show only the findings of this specific context type. The findings status parameter is relative to this context parameter. (optional)
	cve := "cve_example" // string | CVE ID. (optional)
	cvss := float64(1.2) // float64 | CVSS ID (0-10), which can be double (decimal) values. This filter returns all findings with a CVSS value equal to the provided value. (optional)
	cvssGte := float64(1.2) // float64 | CVSS ID (0-10), which can be double (decimal) values. This filter returns all findings with a CVSS value greater than or equal to the provided value. (optional)
	cwe := []int32{int32(123)} // []int32 | CWE ID (single or comma-delimited). (optional)
	findingCategory := []int32{int32(123)} // []int32 | Category of finding (single or comma-delimited). Not valid for the SCA scan type. (optional)
	includeAnnot := true // bool | Use this flag to include the annotations in the response. Not valid for the SCA scan type. (optional)
	includeExpDate := true // bool | Use this flag to include findings grace period expiry date in the response. (optional)
	mitigatedAfter := time.Now() // time.Time | Return all findings with annotations mitigated after the specified date. Does not apply to the SCA scan type. Use date format: yyyy-MM-dd (optional)
	new := true // bool | Use this flag to include only new findings in the current context (policy or sandbox) in the response. (optional)
	scaDepMode := "scaDepMode_example" // string | Return all findings with the specified SCA dependency mode. Only valid for the SCA scan type. (optional)
	scaScanMode := "scaScanMode_example" // string | Return all findings from SCA scans of the specified scan mode. Only valid for the SCA scan type. (optional)
	scanType := []string{"ScanType_example"} // []string | Type of scan in which the finding was found (single or comma-delimited). (optional)
	severity := int32(56) // int32 | This filter returns all findings with this severity value (0-5). (optional)
	severityGte := int32(56) // int32 | This filter returns all scan findings with a severity value greater than or equal to the value of the filter (0-5). (optional)
	violatesPolicy := true // bool | Use this flag to filter the results based on whether the results violate the policy associated with the application. True means the results negatively impact the policy and should be fixed. Not valid for the SCA scan type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationFindingsInformationAPI.GetFindingsUsingGET(context.Background(), applicationGuid).Context(context).Cve(cve).Cvss(cvss).CvssGte(cvssGte).Cwe(cwe).FindingCategory(findingCategory).IncludeAnnot(includeAnnot).IncludeExpDate(includeExpDate).MitigatedAfter(mitigatedAfter).New(new).ScaDepMode(scaDepMode).ScaScanMode(scaScanMode).ScanType(scanType).Severity(severity).SeverityGte(severityGte).ViolatesPolicy(violatesPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFindingsInformationAPI.GetFindingsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFindingsUsingGET`: PagedResourceOfFinding
	fmt.Fprintf(os.Stdout, "Response from `ApplicationFindingsInformationAPI.GetFindingsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationGuid** | **string** | The application identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFindingsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | Context type, which filters results to show only the findings of this specific context type. The findings status parameter is relative to this context parameter. | 
 **cve** | **string** | CVE ID. | 
 **cvss** | **float64** | CVSS ID (0-10), which can be double (decimal) values. This filter returns all findings with a CVSS value equal to the provided value. | 
 **cvssGte** | **float64** | CVSS ID (0-10), which can be double (decimal) values. This filter returns all findings with a CVSS value greater than or equal to the provided value. | 
 **cwe** | **[]int32** | CWE ID (single or comma-delimited). | 
 **findingCategory** | **[]int32** | Category of finding (single or comma-delimited). Not valid for the SCA scan type. | 
 **includeAnnot** | **bool** | Use this flag to include the annotations in the response. Not valid for the SCA scan type. | 
 **includeExpDate** | **bool** | Use this flag to include findings grace period expiry date in the response. | 
 **mitigatedAfter** | **time.Time** | Return all findings with annotations mitigated after the specified date. Does not apply to the SCA scan type. Use date format: yyyy-MM-dd | 
 **new** | **bool** | Use this flag to include only new findings in the current context (policy or sandbox) in the response. | 
 **scaDepMode** | **string** | Return all findings with the specified SCA dependency mode. Only valid for the SCA scan type. | 
 **scaScanMode** | **string** | Return all findings from SCA scans of the specified scan mode. Only valid for the SCA scan type. | 
 **scanType** | **[]string** | Type of scan in which the finding was found (single or comma-delimited). | 
 **severity** | **int32** | This filter returns all findings with this severity value (0-5). | 
 **severityGte** | **int32** | This filter returns all scan findings with a severity value greater than or equal to the value of the filter (0-5). | 
 **violatesPolicy** | **bool** | Use this flag to filter the results based on whether the results violate the policy associated with the application. True means the results negatively impact the policy and should be fixed. Not valid for the SCA scan type. | 

### Return type

[**PagedResourceOfFinding**](PagedResourceOfFinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


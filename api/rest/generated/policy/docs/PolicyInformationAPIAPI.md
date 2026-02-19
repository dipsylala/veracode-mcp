# \PolicyInformationAPIAPI

All URIs are relative to *http://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyUsingPOST**](PolicyInformationAPIAPI.md#CreatePolicyUsingPOST) | **Post** /appsec/v1/policies | Create a new policy.
[**DeletePolicyUsingDELETE**](PolicyInformationAPIAPI.md#DeletePolicyUsingDELETE) | **Delete** /appsec/v1/policies/{policyGuid} | Delete all versions of the policy.
[**GetPoliciesUsingGET**](PolicyInformationAPIAPI.md#GetPoliciesUsingGET) | **Get** /appsec/v1/policies | getPolicies
[**GetPolicyUsingGET**](PolicyInformationAPIAPI.md#GetPolicyUsingGET) | **Get** /appsec/v1/policies/{policyGuid} | getPolicy
[**GetPolicyVersionUsingGET**](PolicyInformationAPIAPI.md#GetPolicyVersionUsingGET) | **Get** /appsec/v1/policies/{policyGuid}/versions/{version} | getPolicyVersion
[**GetPolicyVersionsUsingGET**](PolicyInformationAPIAPI.md#GetPolicyVersionsUsingGET) | **Get** /appsec/v1/policies/{policyGuid}/versions | getPolicyVersions
[**UpdatePolicyUsingPUT**](PolicyInformationAPIAPI.md#UpdatePolicyUsingPUT) | **Put** /appsec/v1/policies/{policyGuid} | Update the policy.



## CreatePolicyUsingPOST

> PolicyVersion CreatePolicyUsingPOST(ctx).Policy(policy).Execute()

Create a new policy.



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
	policy := *openapiclient.NewPolicyVersion() // PolicyVersion | The policy object to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.CreatePolicyUsingPOST(context.Background()).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.CreatePolicyUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicyUsingPOST`: PolicyVersion
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.CreatePolicyUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**PolicyVersion**](PolicyVersion.md) | The policy object to be created. | 

### Return type

[**PolicyVersion**](PolicyVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyUsingDELETE

> map[string]interface{} DeletePolicyUsingDELETE(ctx, policyGuid).ReplaceWithDefaultPolicy(replaceWithDefaultPolicy).ReplacementGUID(replacementGUID).Execute()

Delete all versions of the policy.



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
	policyGuid := "policyGuid_example" // string | The unique identifier of the policy (GUID).
	replaceWithDefaultPolicy := true // bool | Replace with the default policy, based on the business criticality of the application. (optional) (default to false)
	replacementGUID := "replacementGUID_example" // string | The unique identifier of the replacement policy (GUID). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.DeletePolicyUsingDELETE(context.Background(), policyGuid).ReplaceWithDefaultPolicy(replaceWithDefaultPolicy).ReplacementGUID(replacementGUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.DeletePolicyUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicyUsingDELETE`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.DeletePolicyUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | The unique identifier of the policy (GUID). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceWithDefaultPolicy** | **bool** | Replace with the default policy, based on the business criticality of the application. | [default to false]
 **replacementGUID** | **string** | The unique identifier of the replacement policy (GUID). | 

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


## GetPoliciesUsingGET

> PagedResourceOfPolicyVersion GetPoliciesUsingGET(ctx).Category(category).LegacyPolicyId(legacyPolicyId).Name(name).NameExact(nameExact).Page(page).PublicPolicy(publicPolicy).Size(size).VendorPolicy(vendorPolicy).Execute()

getPolicies



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
	category := "category_example" // string | The category of the policy. (optional)
	legacyPolicyId := int32(56) // int32 | Filters results based on the ID of the custom policy created in the Veracode Platform. (optional)
	name := "name_example" // string | Filter on the policy name. (optional)
	nameExact := true // bool | Use this flag to enforce exact name-matching when filtering on the policy name. (optional)
	page := int32(56) // int32 | Page number. Defaults to 0. (optional)
	publicPolicy := true // bool | Filters results to include or exclude a public Veracode policy. (optional) (default to true)
	size := int32(56) // int32 | Page size (1-500, defaults to 50). (optional)
	vendorPolicy := true // bool | Filters results to those with or without a vendor policy flag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.GetPoliciesUsingGET(context.Background()).Category(category).LegacyPolicyId(legacyPolicyId).Name(name).NameExact(nameExact).Page(page).PublicPolicy(publicPolicy).Size(size).VendorPolicy(vendorPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.GetPoliciesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoliciesUsingGET`: PagedResourceOfPolicyVersion
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.GetPoliciesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | The category of the policy. | 
 **legacyPolicyId** | **int32** | Filters results based on the ID of the custom policy created in the Veracode Platform. | 
 **name** | **string** | Filter on the policy name. | 
 **nameExact** | **bool** | Use this flag to enforce exact name-matching when filtering on the policy name. | 
 **page** | **int32** | Page number. Defaults to 0. | 
 **publicPolicy** | **bool** | Filters results to include or exclude a public Veracode policy. | [default to true]
 **size** | **int32** | Page size (1-500, defaults to 50). | 
 **vendorPolicy** | **bool** | Filters results to those with or without a vendor policy flag. | 

### Return type

[**PagedResourceOfPolicyVersion**](PagedResourceOfPolicyVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyUsingGET

> PolicyVersion GetPolicyUsingGET(ctx, policyGuid).Execute()

getPolicy



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
	policyGuid := "policyGuid_example" // string | The unique identifier of the policy (GUID).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.GetPolicyUsingGET(context.Background(), policyGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.GetPolicyUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyUsingGET`: PolicyVersion
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.GetPolicyUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | The unique identifier of the policy (GUID). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyVersion**](PolicyVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyVersionUsingGET

> PolicyVersion GetPolicyVersionUsingGET(ctx, policyGuid, version).Execute()

getPolicyVersion



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
	policyGuid := "policyGuid_example" // string | The unique identifier of the policy (GUID).
	version := int32(56) // int32 | The specific version of this policy. The default is the last version provided.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.GetPolicyVersionUsingGET(context.Background(), policyGuid, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.GetPolicyVersionUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyVersionUsingGET`: PolicyVersion
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.GetPolicyVersionUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | The unique identifier of the policy (GUID). | 
**version** | **int32** | The specific version of this policy. The default is the last version provided. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyVersionUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyVersion**](PolicyVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyVersionsUsingGET

> PagedResourceOfPolicyVersion GetPolicyVersionsUsingGET(ctx, policyGuid).Page(page).Size(size).Execute()

getPolicyVersions



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
	policyGuid := "policyGuid_example" // string | The unique identifier of the policy (GUID).
	page := int32(56) // int32 | Page number. Defaults to 0. (optional)
	size := int32(56) // int32 | Page size (1-500). Defaults to 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.GetPolicyVersionsUsingGET(context.Background(), policyGuid).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.GetPolicyVersionsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyVersionsUsingGET`: PagedResourceOfPolicyVersion
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.GetPolicyVersionsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | The unique identifier of the policy (GUID). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyVersionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Defaults to 0. | 
 **size** | **int32** | Page size (1-500). Defaults to 50. | 

### Return type

[**PagedResourceOfPolicyVersion**](PagedResourceOfPolicyVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyUsingPUT

> PolicyVersion UpdatePolicyUsingPUT(ctx, policyGuid).Policy(policy).Execute()

Update the policy.



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
	policyGuid := "policyGuid_example" // string | The unique identifier of the policy (GUID).
	policy := *openapiclient.NewPolicyVersion() // PolicyVersion | The new policy version to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyInformationAPIAPI.UpdatePolicyUsingPUT(context.Background(), policyGuid).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyInformationAPIAPI.UpdatePolicyUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicyUsingPUT`: PolicyVersion
	fmt.Fprintf(os.Stdout, "Response from `PolicyInformationAPIAPI.UpdatePolicyUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyGuid** | **string** | The unique identifier of the policy (GUID). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**PolicyVersion**](PolicyVersion.md) | The new policy version to be created. | 

### Return type

[**PolicyVersion**](PolicyVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


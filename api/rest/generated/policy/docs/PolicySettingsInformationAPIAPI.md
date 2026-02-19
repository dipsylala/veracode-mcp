# \PolicySettingsInformationAPIAPI

All URIs are relative to *http://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentsUsingGET**](PolicySettingsInformationAPIAPI.md#GetComponentsUsingGET) | **Get** /appsec/v1/policy_settings/components | Retrieve the default pre-build component policies.
[**GetPolicySettingsUsingGET**](PolicySettingsInformationAPIAPI.md#GetPolicySettingsUsingGET) | **Get** /appsec/v1/policy_settings | Retrieve the application policy Settings
[**GetThirdpartyLibrariesUsingGET**](PolicySettingsInformationAPIAPI.md#GetThirdpartyLibrariesUsingGET) | **Get** /appsec/v1/policy_settings/thirdparty_libraries | Retrieve the default pre-build component policies.
[**UpdateComponentsUsingPUT**](PolicySettingsInformationAPIAPI.md#UpdateComponentsUsingPUT) | **Put** /appsec/v1/policy_settings/components | Update the default pre-build component policies.
[**UpdatePolicySettingsUsingPUT**](PolicySettingsInformationAPIAPI.md#UpdatePolicySettingsUsingPUT) | **Put** /appsec/v1/policy_settings | Update the application policy settings.
[**UpdateThirdpartyLibrariesUsingPUT**](PolicySettingsInformationAPIAPI.md#UpdateThirdpartyLibrariesUsingPUT) | **Put** /appsec/v1/policy_settings/thirdparty_libraries | Update the default pre-build component policies.



## GetComponentsUsingGET

> PagedResourceOfComponentPolicySetting GetComponentsUsingGET(ctx).Execute()

Retrieve the default pre-build component policies.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySettingsInformationAPIAPI.GetComponentsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySettingsInformationAPIAPI.GetComponentsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentsUsingGET`: PagedResourceOfComponentPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `PolicySettingsInformationAPIAPI.GetComponentsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsUsingGETRequest struct via the builder pattern


### Return type

[**PagedResourceOfComponentPolicySetting**](PagedResourceOfComponentPolicySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicySettingsUsingGET

> PagedResourceOfPolicySetting GetPolicySettingsUsingGET(ctx).Execute()

Retrieve the application policy Settings

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySettingsInformationAPIAPI.GetPolicySettingsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySettingsInformationAPIAPI.GetPolicySettingsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicySettingsUsingGET`: PagedResourceOfPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `PolicySettingsInformationAPIAPI.GetPolicySettingsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicySettingsUsingGETRequest struct via the builder pattern


### Return type

[**PagedResourceOfPolicySetting**](PagedResourceOfPolicySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThirdpartyLibrariesUsingGET

> PagedResourceOfComponentPolicySetting GetThirdpartyLibrariesUsingGET(ctx).Execute()

Retrieve the default pre-build component policies.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySettingsInformationAPIAPI.GetThirdpartyLibrariesUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySettingsInformationAPIAPI.GetThirdpartyLibrariesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThirdpartyLibrariesUsingGET`: PagedResourceOfComponentPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `PolicySettingsInformationAPIAPI.GetThirdpartyLibrariesUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdpartyLibrariesUsingGETRequest struct via the builder pattern


### Return type

[**PagedResourceOfComponentPolicySetting**](PagedResourceOfComponentPolicySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponentsUsingPUT

> []ComponentPolicySetting UpdateComponentsUsingPUT(ctx).ComponentPolicySettings(componentPolicySettings).Execute()

Update the default pre-build component policies.



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
	componentPolicySettings := []openapiclient.ComponentPolicySetting{*openapiclient.NewComponentPolicySetting()} // []ComponentPolicySetting | Pre-build component policy settings that you are updating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySettingsInformationAPIAPI.UpdateComponentsUsingPUT(context.Background()).ComponentPolicySettings(componentPolicySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySettingsInformationAPIAPI.UpdateComponentsUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComponentsUsingPUT`: []ComponentPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `PolicySettingsInformationAPIAPI.UpdateComponentsUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComponentsUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentPolicySettings** | [**[]ComponentPolicySetting**](ComponentPolicySetting.md) | Pre-build component policy settings that you are updating. | 

### Return type

[**[]ComponentPolicySetting**](ComponentPolicySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicySettingsUsingPUT

> []PolicySetting UpdatePolicySettingsUsingPUT(ctx).PolicySettings(policySettings).Execute()

Update the application policy settings.



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
	policySettings := []openapiclient.PolicySetting{*openapiclient.NewPolicySetting()} // []PolicySetting | Application policy settings that you are updating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySettingsInformationAPIAPI.UpdatePolicySettingsUsingPUT(context.Background()).PolicySettings(policySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySettingsInformationAPIAPI.UpdatePolicySettingsUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicySettingsUsingPUT`: []PolicySetting
	fmt.Fprintf(os.Stdout, "Response from `PolicySettingsInformationAPIAPI.UpdatePolicySettingsUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicySettingsUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policySettings** | [**[]PolicySetting**](PolicySetting.md) | Application policy settings that you are updating. | 

### Return type

[**[]PolicySetting**](PolicySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThirdpartyLibrariesUsingPUT

> []ComponentPolicySetting UpdateThirdpartyLibrariesUsingPUT(ctx).ComponentPolicySettings(componentPolicySettings).Execute()

Update the default pre-build component policies.



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
	componentPolicySettings := []openapiclient.ComponentPolicySetting{*openapiclient.NewComponentPolicySetting()} // []ComponentPolicySetting | Pre-build component policy settings that you are updating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicySettingsInformationAPIAPI.UpdateThirdpartyLibrariesUsingPUT(context.Background()).ComponentPolicySettings(componentPolicySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicySettingsInformationAPIAPI.UpdateThirdpartyLibrariesUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateThirdpartyLibrariesUsingPUT`: []ComponentPolicySetting
	fmt.Fprintf(os.Stdout, "Response from `PolicySettingsInformationAPIAPI.UpdateThirdpartyLibrariesUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThirdpartyLibrariesUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentPolicySettings** | [**[]ComponentPolicySetting**](ComponentPolicySetting.md) | Pre-build component policy settings that you are updating. | 

### Return type

[**[]ComponentPolicySetting**](ComponentPolicySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


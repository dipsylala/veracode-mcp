# \HealthcheckAPIsAPI

All URIs are relative to *http://api.veracode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthcheckStatusGet**](HealthcheckAPIsAPI.md#HealthcheckStatusGet) | **Get** /healthcheck/status | checkStatus



## HealthcheckStatusGet

> HealthcheckStatusGet(ctx).Execute()

checkStatus



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
	r, err := apiClient.HealthcheckAPIsAPI.HealthcheckStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPIsAPI.HealthcheckStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckStatusGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# IssueSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppGuid** | Pointer to **string** | GUID for the application. | [optional] [readonly] 
**Name** | Pointer to **int32** | ID of the application. | [optional] [readonly] 
**BuildId** | Pointer to **int32** | ID of the build. | [optional] [readonly] 
**IssueId** | Pointer to **int32** | Flaw or issues ID of the finding. | [optional] [readonly] 
**Context** | Pointer to **string** | GUID of the specified sandbox. | [optional] [readonly] 

## Methods

### NewIssueSummary

`func NewIssueSummary() *IssueSummary`

NewIssueSummary instantiates a new IssueSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueSummaryWithDefaults

`func NewIssueSummaryWithDefaults() *IssueSummary`

NewIssueSummaryWithDefaults instantiates a new IssueSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppGuid

`func (o *IssueSummary) GetAppGuid() string`

GetAppGuid returns the AppGuid field if non-nil, zero value otherwise.

### GetAppGuidOk

`func (o *IssueSummary) GetAppGuidOk() (*string, bool)`

GetAppGuidOk returns a tuple with the AppGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGuid

`func (o *IssueSummary) SetAppGuid(v string)`

SetAppGuid sets AppGuid field to given value.

### HasAppGuid

`func (o *IssueSummary) HasAppGuid() bool`

HasAppGuid returns a boolean if a field has been set.

### GetName

`func (o *IssueSummary) GetName() int32`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueSummary) GetNameOk() (*int32, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueSummary) SetName(v int32)`

SetName sets Name field to given value.

### HasName

`func (o *IssueSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuildId

`func (o *IssueSummary) GetBuildId() int32`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *IssueSummary) GetBuildIdOk() (*int32, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *IssueSummary) SetBuildId(v int32)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *IssueSummary) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetIssueId

`func (o *IssueSummary) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *IssueSummary) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *IssueSummary) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *IssueSummary) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetContext

`func (o *IssueSummary) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *IssueSummary) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *IssueSummary) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *IssueSummary) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



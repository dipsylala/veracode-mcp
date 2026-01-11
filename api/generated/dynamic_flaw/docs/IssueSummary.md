# IssueSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppGuid** | Pointer to **string** | GUID for the application. | [optional] [readonly] 
**AppId** | Pointer to **int32** | Unique identifier for the application. | [optional] [readonly] 
**BuildId** | Pointer to **int32** | ID of the build. | [optional] [readonly] 
**IssueId** | Pointer to **int32** | Unique identifier for the flaw. | [optional] [readonly] 
**EngineVersion** | Pointer to **string** | Version of the scanner engine that discovered the flaw. Optionally, the value can be an empty string. | [optional] [readonly] 
**CweId** | Pointer to **int32** | Unique CWE ID for the flaw. | [optional] 
**Description** | Pointer to **string** | Description of the flaw. | [optional] [readonly] 
**Recommendation** | Pointer to **string** | Recommended process for fixing the flaw. | [optional] [readonly] 

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

### GetAppId

`func (o *IssueSummary) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *IssueSummary) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *IssueSummary) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *IssueSummary) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

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

### GetEngineVersion

`func (o *IssueSummary) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *IssueSummary) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *IssueSummary) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *IssueSummary) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### GetCweId

`func (o *IssueSummary) GetCweId() int32`

GetCweId returns the CweId field if non-nil, zero value otherwise.

### GetCweIdOk

`func (o *IssueSummary) GetCweIdOk() (*int32, bool)`

GetCweIdOk returns a tuple with the CweId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCweId

`func (o *IssueSummary) SetCweId(v int32)`

SetCweId sets CweId field to given value.

### HasCweId

`func (o *IssueSummary) HasCweId() bool`

HasCweId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecommendation

`func (o *IssueSummary) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *IssueSummary) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *IssueSummary) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *IssueSummary) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



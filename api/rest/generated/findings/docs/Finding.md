# Finding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to [**[]Annotation**](Annotation.md) | All comments and explanatory notes related to this application finding. | [optional] 
**BuildId** | Pointer to **int64** | The build ID. | [optional] 
**ContextGuid** | Pointer to **string** | The context ID. | [optional] 
**ContextType** | Pointer to **string** | Context type, which filters results to show only the findings of this specific context type. The findings status parameter is relative to this context parameter. | [optional] 
**Count** | Pointer to **int32** | Number of times a finding occurs in an application, often referred to as prevalence. | [optional] 
**Description** | Pointer to **string** | The detailed description of the finding. | [optional] 
**FindingDetails** | Pointer to [**FindingFindingDetails**](FindingFindingDetails.md) |  | [optional] 
**FindingStatus** | Pointer to [**FindingStatus**](FindingStatus.md) |  | [optional] 
**GracePeriodExpiresDate** | Pointer to **time.Time** | The date on which a grace period expires for the finding. Veracode calculates this date based on the last date a finding was opened (First Found or Last Reopened date), and based on the grace period provided in the security policy assigned to the application. This date only applies to findings that impact policy compliance. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] 
**IssueId** | Pointer to **int64** | Unique ID within the context of this application. | [optional] 
**ScanType** | Pointer to **string** | The type of scan that Veracode used to discover this finding: static, dynamic, manual, SCA. | [optional] 
**ViolatesPolicy** | Pointer to **bool** | Policy is violated or not. | [optional] 

## Methods

### NewFinding

`func NewFinding() *Finding`

NewFinding instantiates a new Finding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingWithDefaults

`func NewFindingWithDefaults() *Finding`

NewFindingWithDefaults instantiates a new Finding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *Finding) GetAnnotations() []Annotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Finding) GetAnnotationsOk() (*[]Annotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Finding) SetAnnotations(v []Annotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Finding) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetBuildId

`func (o *Finding) GetBuildId() int64`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *Finding) GetBuildIdOk() (*int64, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *Finding) SetBuildId(v int64)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *Finding) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetContextGuid

`func (o *Finding) GetContextGuid() string`

GetContextGuid returns the ContextGuid field if non-nil, zero value otherwise.

### GetContextGuidOk

`func (o *Finding) GetContextGuidOk() (*string, bool)`

GetContextGuidOk returns a tuple with the ContextGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextGuid

`func (o *Finding) SetContextGuid(v string)`

SetContextGuid sets ContextGuid field to given value.

### HasContextGuid

`func (o *Finding) HasContextGuid() bool`

HasContextGuid returns a boolean if a field has been set.

### GetContextType

`func (o *Finding) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *Finding) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *Finding) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *Finding) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetCount

`func (o *Finding) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Finding) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Finding) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Finding) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDescription

`func (o *Finding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Finding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Finding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Finding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFindingDetails

`func (o *Finding) GetFindingDetails() FindingFindingDetails`

GetFindingDetails returns the FindingDetails field if non-nil, zero value otherwise.

### GetFindingDetailsOk

`func (o *Finding) GetFindingDetailsOk() (*FindingFindingDetails, bool)`

GetFindingDetailsOk returns a tuple with the FindingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingDetails

`func (o *Finding) SetFindingDetails(v FindingFindingDetails)`

SetFindingDetails sets FindingDetails field to given value.

### HasFindingDetails

`func (o *Finding) HasFindingDetails() bool`

HasFindingDetails returns a boolean if a field has been set.

### GetFindingStatus

`func (o *Finding) GetFindingStatus() FindingStatus`

GetFindingStatus returns the FindingStatus field if non-nil, zero value otherwise.

### GetFindingStatusOk

`func (o *Finding) GetFindingStatusOk() (*FindingStatus, bool)`

GetFindingStatusOk returns a tuple with the FindingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingStatus

`func (o *Finding) SetFindingStatus(v FindingStatus)`

SetFindingStatus sets FindingStatus field to given value.

### HasFindingStatus

`func (o *Finding) HasFindingStatus() bool`

HasFindingStatus returns a boolean if a field has been set.

### GetGracePeriodExpiresDate

`func (o *Finding) GetGracePeriodExpiresDate() time.Time`

GetGracePeriodExpiresDate returns the GracePeriodExpiresDate field if non-nil, zero value otherwise.

### GetGracePeriodExpiresDateOk

`func (o *Finding) GetGracePeriodExpiresDateOk() (*time.Time, bool)`

GetGracePeriodExpiresDateOk returns a tuple with the GracePeriodExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodExpiresDate

`func (o *Finding) SetGracePeriodExpiresDate(v time.Time)`

SetGracePeriodExpiresDate sets GracePeriodExpiresDate field to given value.

### HasGracePeriodExpiresDate

`func (o *Finding) HasGracePeriodExpiresDate() bool`

HasGracePeriodExpiresDate returns a boolean if a field has been set.

### GetIssueId

`func (o *Finding) GetIssueId() int64`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Finding) GetIssueIdOk() (*int64, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Finding) SetIssueId(v int64)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *Finding) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetScanType

`func (o *Finding) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *Finding) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *Finding) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *Finding) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetViolatesPolicy

`func (o *Finding) GetViolatesPolicy() bool`

GetViolatesPolicy returns the ViolatesPolicy field if non-nil, zero value otherwise.

### GetViolatesPolicyOk

`func (o *Finding) GetViolatesPolicyOk() (*bool, bool)`

GetViolatesPolicyOk returns a tuple with the ViolatesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatesPolicy

`func (o *Finding) SetViolatesPolicy(v bool)`

SetViolatesPolicy sets ViolatesPolicy field to given value.

### HasViolatesPolicy

`func (o *Finding) HasViolatesPolicy() bool`

HasViolatesPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



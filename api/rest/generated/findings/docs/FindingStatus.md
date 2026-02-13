# FindingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstFoundDate** | Pointer to **time.Time** | Date when the finding was first found. For SCA findings, this date may reference the latest of either the date the vulnerability was published to the Veracode vulnerability database or the date the library was found in a scan. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] 
**LastSeenDate** | Pointer to **time.Time** | The date and time when the finding was last seen. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] 
**MitigationReviewStatus** | Pointer to **string** | Field indicating if the mitigation applied to the finding conforms to or deviates from industry standards. | [optional] 
**New** | Pointer to **bool** | Use this flag to indicate if this is the first time this finding appeared in any context of the latest scan. | [optional] 
**Resolution** | Pointer to **string** | Resolution of the finding. | [optional] 
**ResolutionStatus** | Pointer to **string** | The resolution status of the finding. | [optional] 
**Status** | Pointer to **string** | Status of the finding: open or closed. | [optional] 

## Methods

### NewFindingStatus

`func NewFindingStatus() *FindingStatus`

NewFindingStatus instantiates a new FindingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingStatusWithDefaults

`func NewFindingStatusWithDefaults() *FindingStatus`

NewFindingStatusWithDefaults instantiates a new FindingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstFoundDate

`func (o *FindingStatus) GetFirstFoundDate() time.Time`

GetFirstFoundDate returns the FirstFoundDate field if non-nil, zero value otherwise.

### GetFirstFoundDateOk

`func (o *FindingStatus) GetFirstFoundDateOk() (*time.Time, bool)`

GetFirstFoundDateOk returns a tuple with the FirstFoundDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFoundDate

`func (o *FindingStatus) SetFirstFoundDate(v time.Time)`

SetFirstFoundDate sets FirstFoundDate field to given value.

### HasFirstFoundDate

`func (o *FindingStatus) HasFirstFoundDate() bool`

HasFirstFoundDate returns a boolean if a field has been set.

### GetLastSeenDate

`func (o *FindingStatus) GetLastSeenDate() time.Time`

GetLastSeenDate returns the LastSeenDate field if non-nil, zero value otherwise.

### GetLastSeenDateOk

`func (o *FindingStatus) GetLastSeenDateOk() (*time.Time, bool)`

GetLastSeenDateOk returns a tuple with the LastSeenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenDate

`func (o *FindingStatus) SetLastSeenDate(v time.Time)`

SetLastSeenDate sets LastSeenDate field to given value.

### HasLastSeenDate

`func (o *FindingStatus) HasLastSeenDate() bool`

HasLastSeenDate returns a boolean if a field has been set.

### GetMitigationReviewStatus

`func (o *FindingStatus) GetMitigationReviewStatus() string`

GetMitigationReviewStatus returns the MitigationReviewStatus field if non-nil, zero value otherwise.

### GetMitigationReviewStatusOk

`func (o *FindingStatus) GetMitigationReviewStatusOk() (*string, bool)`

GetMitigationReviewStatusOk returns a tuple with the MitigationReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationReviewStatus

`func (o *FindingStatus) SetMitigationReviewStatus(v string)`

SetMitigationReviewStatus sets MitigationReviewStatus field to given value.

### HasMitigationReviewStatus

`func (o *FindingStatus) HasMitigationReviewStatus() bool`

HasMitigationReviewStatus returns a boolean if a field has been set.

### GetNew

`func (o *FindingStatus) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *FindingStatus) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *FindingStatus) SetNew(v bool)`

SetNew sets New field to given value.

### HasNew

`func (o *FindingStatus) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetResolution

`func (o *FindingStatus) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *FindingStatus) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *FindingStatus) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *FindingStatus) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *FindingStatus) GetResolutionStatus() string`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *FindingStatus) GetResolutionStatusOk() (*string, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *FindingStatus) SetResolutionStatus(v string)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *FindingStatus) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetStatus

`func (o *FindingStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FindingStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FindingStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FindingStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



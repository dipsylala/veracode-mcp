# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProfileUrl** | Pointer to **string** | Unique path to the application profile. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date and time when the application was created. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**Guid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastCompletedScanDate** | Pointer to **time.Time** | Date and time of the last completed scan for this application. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | The date and time when the application was modified. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**Oid** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Profile** | Pointer to [**ApplicationProfile**](ApplicationProfile.md) |  | [optional] 
**ResultsUrl** | Pointer to **string** | Unique path to the latest results. | [optional] [readonly] 
**Scans** | Pointer to [**[]ApplicationScan**](ApplicationScan.md) |  | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProfileUrl

`func (o *Application) GetAppProfileUrl() string`

GetAppProfileUrl returns the AppProfileUrl field if non-nil, zero value otherwise.

### GetAppProfileUrlOk

`func (o *Application) GetAppProfileUrlOk() (*string, bool)`

GetAppProfileUrlOk returns a tuple with the AppProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileUrl

`func (o *Application) SetAppProfileUrl(v string)`

SetAppProfileUrl sets AppProfileUrl field to given value.

### HasAppProfileUrl

`func (o *Application) HasAppProfileUrl() bool`

HasAppProfileUrl returns a boolean if a field has been set.

### GetCreated

`func (o *Application) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Application) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Application) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Application) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGuid

`func (o *Application) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Application) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Application) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Application) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *Application) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastCompletedScanDate

`func (o *Application) GetLastCompletedScanDate() time.Time`

GetLastCompletedScanDate returns the LastCompletedScanDate field if non-nil, zero value otherwise.

### GetLastCompletedScanDateOk

`func (o *Application) GetLastCompletedScanDateOk() (*time.Time, bool)`

GetLastCompletedScanDateOk returns a tuple with the LastCompletedScanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedScanDate

`func (o *Application) SetLastCompletedScanDate(v time.Time)`

SetLastCompletedScanDate sets LastCompletedScanDate field to given value.

### HasLastCompletedScanDate

`func (o *Application) HasLastCompletedScanDate() bool`

HasLastCompletedScanDate returns a boolean if a field has been set.

### GetModified

`func (o *Application) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Application) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Application) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Application) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetOid

`func (o *Application) GetOid() int32`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *Application) GetOidOk() (*int32, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *Application) SetOid(v int32)`

SetOid sets Oid field to given value.

### HasOid

`func (o *Application) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Application) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Application) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Application) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Application) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProfile

`func (o *Application) GetProfile() ApplicationProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Application) GetProfileOk() (*ApplicationProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Application) SetProfile(v ApplicationProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Application) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultsUrl

`func (o *Application) GetResultsUrl() string`

GetResultsUrl returns the ResultsUrl field if non-nil, zero value otherwise.

### GetResultsUrlOk

`func (o *Application) GetResultsUrlOk() (*string, bool)`

GetResultsUrlOk returns a tuple with the ResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsUrl

`func (o *Application) SetResultsUrl(v string)`

SetResultsUrl sets ResultsUrl field to given value.

### HasResultsUrl

`func (o *Application) HasResultsUrl() bool`

HasResultsUrl returns a boolean if a field has been set.

### GetScans

`func (o *Application) GetScans() []ApplicationScan`

GetScans returns the Scans field if non-nil, zero value otherwise.

### GetScansOk

`func (o *Application) GetScansOk() (*[]ApplicationScan, bool)`

GetScansOk returns a tuple with the Scans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScans

`func (o *Application) SetScans(v []ApplicationScan)`

SetScans sets Scans field to given value.

### HasScans

`func (o *Application) HasScans() bool`

HasScans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



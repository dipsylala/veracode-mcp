# EntityModelScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier of the scan. | [optional] 
**ApplicationId** | Pointer to **int32** | Unique identifier of the application. | [optional] 
**ApplicationGuid** | Pointer to **string** | Application identifier (UUID). | [optional] 
**Name** | Pointer to **string** | Name of the scan. | [optional] 
**Created** | Pointer to **time.Time** | Date and time of scan creation. | [optional] 
**Submitted** | Pointer to **time.Time** | Date and time of scan submission. | [optional] 
**Completed** | Pointer to **time.Time** | Date and time the scan completed. | [optional] 
**Published** | Pointer to **time.Time** | Date and time of scan publication. | [optional] 
**Status** | Pointer to **string** | Current scan status. | [optional] 
**ScanType** | Pointer to **string** | Scan type. | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewEntityModelScan

`func NewEntityModelScan() *EntityModelScan`

NewEntityModelScan instantiates a new EntityModelScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityModelScanWithDefaults

`func NewEntityModelScanWithDefaults() *EntityModelScan`

NewEntityModelScanWithDefaults instantiates a new EntityModelScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityModelScan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityModelScan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityModelScan) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntityModelScan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplicationId

`func (o *EntityModelScan) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *EntityModelScan) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *EntityModelScan) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *EntityModelScan) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationGuid

`func (o *EntityModelScan) GetApplicationGuid() string`

GetApplicationGuid returns the ApplicationGuid field if non-nil, zero value otherwise.

### GetApplicationGuidOk

`func (o *EntityModelScan) GetApplicationGuidOk() (*string, bool)`

GetApplicationGuidOk returns a tuple with the ApplicationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGuid

`func (o *EntityModelScan) SetApplicationGuid(v string)`

SetApplicationGuid sets ApplicationGuid field to given value.

### HasApplicationGuid

`func (o *EntityModelScan) HasApplicationGuid() bool`

HasApplicationGuid returns a boolean if a field has been set.

### GetName

`func (o *EntityModelScan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityModelScan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityModelScan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityModelScan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *EntityModelScan) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntityModelScan) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntityModelScan) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntityModelScan) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSubmitted

`func (o *EntityModelScan) GetSubmitted() time.Time`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *EntityModelScan) GetSubmittedOk() (*time.Time, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *EntityModelScan) SetSubmitted(v time.Time)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *EntityModelScan) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetCompleted

`func (o *EntityModelScan) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *EntityModelScan) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *EntityModelScan) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *EntityModelScan) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetPublished

`func (o *EntityModelScan) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *EntityModelScan) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *EntityModelScan) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *EntityModelScan) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetStatus

`func (o *EntityModelScan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntityModelScan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntityModelScan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntityModelScan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScanType

`func (o *EntityModelScan) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *EntityModelScan) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *EntityModelScan) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *EntityModelScan) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetLinks

`func (o *EntityModelScan) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntityModelScan) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntityModelScan) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntityModelScan) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



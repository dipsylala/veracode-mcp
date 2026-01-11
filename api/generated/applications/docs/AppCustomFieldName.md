# AppCustomFieldName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date and time when the application was created. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**Id** | Pointer to **int32** | Unique identifier of the category. | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | The date and time when the application was modified. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 

## Methods

### NewAppCustomFieldName

`func NewAppCustomFieldName() *AppCustomFieldName`

NewAppCustomFieldName instantiates a new AppCustomFieldName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCustomFieldNameWithDefaults

`func NewAppCustomFieldNameWithDefaults() *AppCustomFieldName`

NewAppCustomFieldNameWithDefaults instantiates a new AppCustomFieldName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AppCustomFieldName) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppCustomFieldName) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppCustomFieldName) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppCustomFieldName) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AppCustomFieldName) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppCustomFieldName) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppCustomFieldName) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AppCustomFieldName) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *AppCustomFieldName) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AppCustomFieldName) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AppCustomFieldName) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AppCustomFieldName) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *AppCustomFieldName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppCustomFieldName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppCustomFieldName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppCustomFieldName) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AppCustomFieldName) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AppCustomFieldName) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AppCustomFieldName) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AppCustomFieldName) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSortOrder

`func (o *AppCustomFieldName) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AppCustomFieldName) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AppCustomFieldName) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AppCustomFieldName) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



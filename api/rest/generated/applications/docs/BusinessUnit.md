# BusinessUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | The business unit GUID. | [optional] 
**Id** | Pointer to **int32** | The business unit ID in the Veracode database. | [optional] [readonly] 
**Name** | Pointer to **string** | The business unit name. | [optional] [readonly] 

## Methods

### NewBusinessUnit

`func NewBusinessUnit() *BusinessUnit`

NewBusinessUnit instantiates a new BusinessUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessUnitWithDefaults

`func NewBusinessUnitWithDefaults() *BusinessUnit`

NewBusinessUnitWithDefaults instantiates a new BusinessUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *BusinessUnit) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *BusinessUnit) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *BusinessUnit) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *BusinessUnit) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *BusinessUnit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessUnit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessUnit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessUnit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BusinessUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessUnit) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



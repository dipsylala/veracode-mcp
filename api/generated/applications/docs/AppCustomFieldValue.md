# AppCustomFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppCustomFieldName** | Pointer to [**AppCustomFieldName**](AppCustomFieldName.md) |  | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the application was created. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**FieldNameId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** | Unique identifier of the category. | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewAppCustomFieldValue

`func NewAppCustomFieldValue() *AppCustomFieldValue`

NewAppCustomFieldValue instantiates a new AppCustomFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCustomFieldValueWithDefaults

`func NewAppCustomFieldValueWithDefaults() *AppCustomFieldValue`

NewAppCustomFieldValueWithDefaults instantiates a new AppCustomFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppCustomFieldName

`func (o *AppCustomFieldValue) GetAppCustomFieldName() AppCustomFieldName`

GetAppCustomFieldName returns the AppCustomFieldName field if non-nil, zero value otherwise.

### GetAppCustomFieldNameOk

`func (o *AppCustomFieldValue) GetAppCustomFieldNameOk() (*AppCustomFieldName, bool)`

GetAppCustomFieldNameOk returns a tuple with the AppCustomFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCustomFieldName

`func (o *AppCustomFieldValue) SetAppCustomFieldName(v AppCustomFieldName)`

SetAppCustomFieldName sets AppCustomFieldName field to given value.

### HasAppCustomFieldName

`func (o *AppCustomFieldValue) HasAppCustomFieldName() bool`

HasAppCustomFieldName returns a boolean if a field has been set.

### GetCreated

`func (o *AppCustomFieldValue) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppCustomFieldValue) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppCustomFieldValue) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppCustomFieldValue) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFieldNameId

`func (o *AppCustomFieldValue) GetFieldNameId() int32`

GetFieldNameId returns the FieldNameId field if non-nil, zero value otherwise.

### GetFieldNameIdOk

`func (o *AppCustomFieldValue) GetFieldNameIdOk() (*int32, bool)`

GetFieldNameIdOk returns a tuple with the FieldNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNameId

`func (o *AppCustomFieldValue) SetFieldNameId(v int32)`

SetFieldNameId sets FieldNameId field to given value.

### HasFieldNameId

`func (o *AppCustomFieldValue) HasFieldNameId() bool`

HasFieldNameId returns a boolean if a field has been set.

### GetId

`func (o *AppCustomFieldValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppCustomFieldValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppCustomFieldValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AppCustomFieldValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *AppCustomFieldValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppCustomFieldValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppCustomFieldValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppCustomFieldValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



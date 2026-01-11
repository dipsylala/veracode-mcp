# Sandbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationGuid** | Pointer to **string** |  | [optional] 
**AutoRecreate** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the sandbox was created. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**CustomFields** | Pointer to [**[]CustomNameValue**](CustomNameValue.md) |  | [optional] 
**Guid** | Pointer to **string** | Unique identifier (UUID). | [optional] [readonly] 
**Id** | Pointer to **int32** | Internal ID. | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | The date and time when the sandbox was modified. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**Name** | Pointer to **string** | The sandbox name | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**OwnerUsername** | Pointer to **string** |  | [optional] 

## Methods

### NewSandbox

`func NewSandbox() *Sandbox`

NewSandbox instantiates a new Sandbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxWithDefaults

`func NewSandboxWithDefaults() *Sandbox`

NewSandboxWithDefaults instantiates a new Sandbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationGuid

`func (o *Sandbox) GetApplicationGuid() string`

GetApplicationGuid returns the ApplicationGuid field if non-nil, zero value otherwise.

### GetApplicationGuidOk

`func (o *Sandbox) GetApplicationGuidOk() (*string, bool)`

GetApplicationGuidOk returns a tuple with the ApplicationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGuid

`func (o *Sandbox) SetApplicationGuid(v string)`

SetApplicationGuid sets ApplicationGuid field to given value.

### HasApplicationGuid

`func (o *Sandbox) HasApplicationGuid() bool`

HasApplicationGuid returns a boolean if a field has been set.

### GetAutoRecreate

`func (o *Sandbox) GetAutoRecreate() bool`

GetAutoRecreate returns the AutoRecreate field if non-nil, zero value otherwise.

### GetAutoRecreateOk

`func (o *Sandbox) GetAutoRecreateOk() (*bool, bool)`

GetAutoRecreateOk returns a tuple with the AutoRecreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecreate

`func (o *Sandbox) SetAutoRecreate(v bool)`

SetAutoRecreate sets AutoRecreate field to given value.

### HasAutoRecreate

`func (o *Sandbox) HasAutoRecreate() bool`

HasAutoRecreate returns a boolean if a field has been set.

### GetCreated

`func (o *Sandbox) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Sandbox) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Sandbox) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Sandbox) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomFields

`func (o *Sandbox) GetCustomFields() []CustomNameValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Sandbox) GetCustomFieldsOk() (*[]CustomNameValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Sandbox) SetCustomFields(v []CustomNameValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Sandbox) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetGuid

`func (o *Sandbox) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Sandbox) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Sandbox) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Sandbox) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *Sandbox) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sandbox) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sandbox) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Sandbox) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *Sandbox) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Sandbox) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Sandbox) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Sandbox) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Sandbox) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sandbox) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sandbox) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sandbox) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Sandbox) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Sandbox) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Sandbox) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Sandbox) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOwnerUsername

`func (o *Sandbox) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *Sandbox) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *Sandbox) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.

### HasOwnerUsername

`func (o *Sandbox) HasOwnerUsername() bool`

HasOwnerUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



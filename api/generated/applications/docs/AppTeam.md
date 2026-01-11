# AppTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | The team ID in the Veracode Identity API. | [optional] 
**TeamId** | Pointer to **int32** | The legacy team_id. | [optional] [readonly] 
**TeamName** | Pointer to **string** | Team name. | [optional] [readonly] 

## Methods

### NewAppTeam

`func NewAppTeam() *AppTeam`

NewAppTeam instantiates a new AppTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTeamWithDefaults

`func NewAppTeamWithDefaults() *AppTeam`

NewAppTeamWithDefaults instantiates a new AppTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AppTeam) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AppTeam) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AppTeam) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AppTeam) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetTeamId

`func (o *AppTeam) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AppTeam) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AppTeam) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *AppTeam) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *AppTeam) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *AppTeam) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *AppTeam) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *AppTeam) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



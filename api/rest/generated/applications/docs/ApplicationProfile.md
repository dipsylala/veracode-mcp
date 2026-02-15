# ApplicationProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcherAppName** | Pointer to **string** |  | [optional] 
**BusinessCriticality** | Pointer to **string** |  | [optional] 
**BusinessOwners** | Pointer to [**[]BusinessOwner**](BusinessOwner.md) |  | [optional] 
**BusinessUnit** | Pointer to [**BusinessUnit**](BusinessUnit.md) |  | [optional] 
**CustomFieldValues** | Pointer to [**[]AppCustomFieldValue**](AppCustomFieldValue.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomNameValue**](CustomNameValue.md) |  | [optional] 
**CustomKmsAlias** | Pointer to **string** | Alias for the Customer Managed Encryption Key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GitRepoUrl** | Pointer to **string** | The URL of the Git repository associated with the application. Veracode includes findings from the Git repository in the reporting for this application. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to [**[]AppPolicy**](AppPolicy.md) |  | [optional] 
**Settings** | Pointer to [**ApplicationSettings**](ApplicationSettings.md) |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**[]AppTeam**](AppTeam.md) |  | [optional] 

## Methods

### NewApplicationProfile

`func NewApplicationProfile() *ApplicationProfile`

NewApplicationProfile instantiates a new ApplicationProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationProfileWithDefaults

`func NewApplicationProfileWithDefaults() *ApplicationProfile`

NewApplicationProfileWithDefaults instantiates a new ApplicationProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcherAppName

`func (o *ApplicationProfile) GetArcherAppName() string`

GetArcherAppName returns the ArcherAppName field if non-nil, zero value otherwise.

### GetArcherAppNameOk

`func (o *ApplicationProfile) GetArcherAppNameOk() (*string, bool)`

GetArcherAppNameOk returns a tuple with the ArcherAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcherAppName

`func (o *ApplicationProfile) SetArcherAppName(v string)`

SetArcherAppName sets ArcherAppName field to given value.

### HasArcherAppName

`func (o *ApplicationProfile) HasArcherAppName() bool`

HasArcherAppName returns a boolean if a field has been set.

### GetBusinessCriticality

`func (o *ApplicationProfile) GetBusinessCriticality() string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *ApplicationProfile) GetBusinessCriticalityOk() (*string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *ApplicationProfile) SetBusinessCriticality(v string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *ApplicationProfile) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### GetBusinessOwners

`func (o *ApplicationProfile) GetBusinessOwners() []BusinessOwner`

GetBusinessOwners returns the BusinessOwners field if non-nil, zero value otherwise.

### GetBusinessOwnersOk

`func (o *ApplicationProfile) GetBusinessOwnersOk() (*[]BusinessOwner, bool)`

GetBusinessOwnersOk returns a tuple with the BusinessOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessOwners

`func (o *ApplicationProfile) SetBusinessOwners(v []BusinessOwner)`

SetBusinessOwners sets BusinessOwners field to given value.

### HasBusinessOwners

`func (o *ApplicationProfile) HasBusinessOwners() bool`

HasBusinessOwners returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *ApplicationProfile) GetBusinessUnit() BusinessUnit`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *ApplicationProfile) GetBusinessUnitOk() (*BusinessUnit, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *ApplicationProfile) SetBusinessUnit(v BusinessUnit)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *ApplicationProfile) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetCustomFieldValues

`func (o *ApplicationProfile) GetCustomFieldValues() []AppCustomFieldValue`

GetCustomFieldValues returns the CustomFieldValues field if non-nil, zero value otherwise.

### GetCustomFieldValuesOk

`func (o *ApplicationProfile) GetCustomFieldValuesOk() (*[]AppCustomFieldValue, bool)`

GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValues

`func (o *ApplicationProfile) SetCustomFieldValues(v []AppCustomFieldValue)`

SetCustomFieldValues sets CustomFieldValues field to given value.

### HasCustomFieldValues

`func (o *ApplicationProfile) HasCustomFieldValues() bool`

HasCustomFieldValues returns a boolean if a field has been set.

### GetCustomFields

`func (o *ApplicationProfile) GetCustomFields() []CustomNameValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApplicationProfile) GetCustomFieldsOk() (*[]CustomNameValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApplicationProfile) SetCustomFields(v []CustomNameValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApplicationProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomKmsAlias

`func (o *ApplicationProfile) GetCustomKmsAlias() string`

GetCustomKmsAlias returns the CustomKmsAlias field if non-nil, zero value otherwise.

### GetCustomKmsAliasOk

`func (o *ApplicationProfile) GetCustomKmsAliasOk() (*string, bool)`

GetCustomKmsAliasOk returns a tuple with the CustomKmsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKmsAlias

`func (o *ApplicationProfile) SetCustomKmsAlias(v string)`

SetCustomKmsAlias sets CustomKmsAlias field to given value.

### HasCustomKmsAlias

`func (o *ApplicationProfile) HasCustomKmsAlias() bool`

HasCustomKmsAlias returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitRepoUrl

`func (o *ApplicationProfile) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *ApplicationProfile) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *ApplicationProfile) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *ApplicationProfile) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### GetName

`func (o *ApplicationProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *ApplicationProfile) GetPolicies() []AppPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApplicationProfile) GetPoliciesOk() (*[]AppPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApplicationProfile) SetPolicies(v []AppPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ApplicationProfile) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetSettings

`func (o *ApplicationProfile) GetSettings() ApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ApplicationProfile) GetSettingsOk() (*ApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ApplicationProfile) SetSettings(v ApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ApplicationProfile) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTags

`func (o *ApplicationProfile) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationProfile) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationProfile) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeams

`func (o *ApplicationProfile) GetTeams() []AppTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ApplicationProfile) GetTeamsOk() (*[]AppTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ApplicationProfile) SetTeams(v []AppTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ApplicationProfile) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



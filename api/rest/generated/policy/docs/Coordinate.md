# Coordinate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinate1** | Pointer to **string** | The name of the first coordinate. | [optional] 
**Coordinate2** | Pointer to **string** | The name of the second coordinate. | [optional] 
**CreatedBy** | Pointer to **string** | The name of the user who created this coordinate. | [optional] 
**CreatedDate** | Pointer to **time.Time** | The date when the user created the coordinate. | [optional] 
**FindingRule** | Pointer to [**FindingRule**](FindingRule.md) |  | [optional] 
**RepoType** | Pointer to **string** | The repository type of the coordinate; for example, nexus, or maven. | [optional] 
**Version** | Pointer to **string** | The version of the coordinate. | [optional] 

## Methods

### NewCoordinate

`func NewCoordinate() *Coordinate`

NewCoordinate instantiates a new Coordinate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoordinateWithDefaults

`func NewCoordinateWithDefaults() *Coordinate`

NewCoordinateWithDefaults instantiates a new Coordinate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinate1

`func (o *Coordinate) GetCoordinate1() string`

GetCoordinate1 returns the Coordinate1 field if non-nil, zero value otherwise.

### GetCoordinate1Ok

`func (o *Coordinate) GetCoordinate1Ok() (*string, bool)`

GetCoordinate1Ok returns a tuple with the Coordinate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate1

`func (o *Coordinate) SetCoordinate1(v string)`

SetCoordinate1 sets Coordinate1 field to given value.

### HasCoordinate1

`func (o *Coordinate) HasCoordinate1() bool`

HasCoordinate1 returns a boolean if a field has been set.

### GetCoordinate2

`func (o *Coordinate) GetCoordinate2() string`

GetCoordinate2 returns the Coordinate2 field if non-nil, zero value otherwise.

### GetCoordinate2Ok

`func (o *Coordinate) GetCoordinate2Ok() (*string, bool)`

GetCoordinate2Ok returns a tuple with the Coordinate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate2

`func (o *Coordinate) SetCoordinate2(v string)`

SetCoordinate2 sets Coordinate2 field to given value.

### HasCoordinate2

`func (o *Coordinate) HasCoordinate2() bool`

HasCoordinate2 returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Coordinate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Coordinate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Coordinate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Coordinate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *Coordinate) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Coordinate) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Coordinate) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Coordinate) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetFindingRule

`func (o *Coordinate) GetFindingRule() FindingRule`

GetFindingRule returns the FindingRule field if non-nil, zero value otherwise.

### GetFindingRuleOk

`func (o *Coordinate) GetFindingRuleOk() (*FindingRule, bool)`

GetFindingRuleOk returns a tuple with the FindingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingRule

`func (o *Coordinate) SetFindingRule(v FindingRule)`

SetFindingRule sets FindingRule field to given value.

### HasFindingRule

`func (o *Coordinate) HasFindingRule() bool`

HasFindingRule returns a boolean if a field has been set.

### GetRepoType

`func (o *Coordinate) GetRepoType() string`

GetRepoType returns the RepoType field if non-nil, zero value otherwise.

### GetRepoTypeOk

`func (o *Coordinate) GetRepoTypeOk() (*string, bool)`

GetRepoTypeOk returns a tuple with the RepoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoType

`func (o *Coordinate) SetRepoType(v string)`

SetRepoType sets RepoType field to given value.

### HasRepoType

`func (o *Coordinate) HasRepoType() bool`

HasRepoType returns a boolean if a field has been set.

### GetVersion

`func (o *Coordinate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Coordinate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Coordinate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Coordinate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



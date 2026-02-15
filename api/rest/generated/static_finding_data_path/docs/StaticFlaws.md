# StaticFlaws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueSummary** | Pointer to [**IssueSummary**](IssueSummary.md) |  | [optional] 
**DataPaths** | Pointer to [**[]DataPaths**](DataPaths.md) | Call stacks. | [optional] 

## Methods

### NewStaticFlaws

`func NewStaticFlaws() *StaticFlaws`

NewStaticFlaws instantiates a new StaticFlaws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticFlawsWithDefaults

`func NewStaticFlawsWithDefaults() *StaticFlaws`

NewStaticFlawsWithDefaults instantiates a new StaticFlaws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueSummary

`func (o *StaticFlaws) GetIssueSummary() IssueSummary`

GetIssueSummary returns the IssueSummary field if non-nil, zero value otherwise.

### GetIssueSummaryOk

`func (o *StaticFlaws) GetIssueSummaryOk() (*IssueSummary, bool)`

GetIssueSummaryOk returns a tuple with the IssueSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSummary

`func (o *StaticFlaws) SetIssueSummary(v IssueSummary)`

SetIssueSummary sets IssueSummary field to given value.

### HasIssueSummary

`func (o *StaticFlaws) HasIssueSummary() bool`

HasIssueSummary returns a boolean if a field has been set.

### GetDataPaths

`func (o *StaticFlaws) GetDataPaths() []DataPaths`

GetDataPaths returns the DataPaths field if non-nil, zero value otherwise.

### GetDataPathsOk

`func (o *StaticFlaws) GetDataPathsOk() (*[]DataPaths, bool)`

GetDataPathsOk returns a tuple with the DataPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPaths

`func (o *StaticFlaws) SetDataPaths(v []DataPaths)`

SetDataPaths sets DataPaths field to given value.

### HasDataPaths

`func (o *StaticFlaws) HasDataPaths() bool`

HasDataPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



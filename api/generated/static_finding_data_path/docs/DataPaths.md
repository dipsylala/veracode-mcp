# DataPaths

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleName** | Pointer to **string** | Name of the module that contains the finding. | [optional] [readonly] 
**Steps** | Pointer to **int32** | Call steps. | [optional] [readonly] 
**LocalPath** | Pointer to **string** | Local filepath to the file that contains the finding. | [optional] [readonly] 
**FunctionName** | Pointer to **string** | Name of the function that contains the finding. | [optional] [readonly] 
**LineNumber** | Pointer to **int32** | Code line number where the finding exists. | [optional] [readonly] 
**Calls** | Pointer to [**[]Call**](Call.md) | Attack vector parameters associated with this request. | [optional] 

## Methods

### NewDataPaths

`func NewDataPaths() *DataPaths`

NewDataPaths instantiates a new DataPaths object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPathsWithDefaults

`func NewDataPathsWithDefaults() *DataPaths`

NewDataPathsWithDefaults instantiates a new DataPaths object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleName

`func (o *DataPaths) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *DataPaths) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *DataPaths) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *DataPaths) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetSteps

`func (o *DataPaths) GetSteps() int32`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DataPaths) GetStepsOk() (*int32, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DataPaths) SetSteps(v int32)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *DataPaths) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetLocalPath

`func (o *DataPaths) GetLocalPath() string`

GetLocalPath returns the LocalPath field if non-nil, zero value otherwise.

### GetLocalPathOk

`func (o *DataPaths) GetLocalPathOk() (*string, bool)`

GetLocalPathOk returns a tuple with the LocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPath

`func (o *DataPaths) SetLocalPath(v string)`

SetLocalPath sets LocalPath field to given value.

### HasLocalPath

`func (o *DataPaths) HasLocalPath() bool`

HasLocalPath returns a boolean if a field has been set.

### GetFunctionName

`func (o *DataPaths) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *DataPaths) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *DataPaths) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *DataPaths) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetLineNumber

`func (o *DataPaths) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *DataPaths) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *DataPaths) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *DataPaths) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetCalls

`func (o *DataPaths) GetCalls() []Call`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *DataPaths) GetCallsOk() (*[]Call, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *DataPaths) SetCalls(v []Call)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *DataPaths) HasCalls() bool`

HasCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



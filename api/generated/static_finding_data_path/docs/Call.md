# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPath** | Pointer to **int32** | Call sequence in the data path. | [optional] 
**FileName** | Pointer to **string** | Filename. | [optional] [readonly] 
**FilePath** | Pointer to **string** | Filepath. | [optional] [readonly] 
**FunctionName** | Pointer to **string** | Function name. | [optional] [readonly] 
**LineNumber** | Pointer to **int32** | Code line number within the file. | [optional] [readonly] 

## Methods

### NewCall

`func NewCall() *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPath

`func (o *Call) GetDataPath() int32`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *Call) GetDataPathOk() (*int32, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *Call) SetDataPath(v int32)`

SetDataPath sets DataPath field to given value.

### HasDataPath

`func (o *Call) HasDataPath() bool`

HasDataPath returns a boolean if a field has been set.

### GetFileName

`func (o *Call) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Call) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Call) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Call) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *Call) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *Call) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *Call) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *Call) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFunctionName

`func (o *Call) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *Call) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *Call) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *Call) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetLineNumber

`func (o *Call) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *Call) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *Call) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *Call) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StaticFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwe** | Pointer to [**StaticFindingCwe**](StaticFindingCwe.md) |  | [optional] 
**Cvss** | Pointer to **string** | The CVSS associated with the finding. | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Exploitability** | Pointer to **int32** | The likelihood that this finding could be exploited by an attacker. Values: -2: Very Unlikely, -1: Unlikely, 0: Neutral, 1: Likely, 2: Very Likely. | [optional] 
**AttackVector** | Pointer to **string** | The function or class where the finding exists. | [optional] 
**FileLineNumber** | Pointer to **int32** | The line number where the finding exists in the file. | [optional] 
**FileName** | Pointer to **string** | The name of the file where the finding exists. | [optional] 
**FilePath** | Pointer to **string** | The path to the file where the finding exists. | [optional] 
**FindingCategory** | Pointer to [**StaticFindingFindingCategory**](StaticFindingFindingCategory.md) |  | [optional] 
**Module** | Pointer to **string** | The name of the module where the finding exists. | [optional] 
**Procedure** | Pointer to **string** | The name of the procedure where the finding exists. | [optional] 
**RelativeLocation** | Pointer to **int32** | The relative location of the finding in the procedure. | [optional] 

## Methods

### NewStaticFinding

`func NewStaticFinding() *StaticFinding`

NewStaticFinding instantiates a new StaticFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticFindingWithDefaults

`func NewStaticFindingWithDefaults() *StaticFinding`

NewStaticFindingWithDefaults instantiates a new StaticFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwe

`func (o *StaticFinding) GetCwe() StaticFindingCwe`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *StaticFinding) GetCweOk() (*StaticFindingCwe, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *StaticFinding) SetCwe(v StaticFindingCwe)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *StaticFinding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetCvss

`func (o *StaticFinding) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *StaticFinding) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *StaticFinding) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *StaticFinding) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetSeverity

`func (o *StaticFinding) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StaticFinding) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StaticFinding) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *StaticFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetExploitability

`func (o *StaticFinding) GetExploitability() int32`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *StaticFinding) GetExploitabilityOk() (*int32, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *StaticFinding) SetExploitability(v int32)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *StaticFinding) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetAttackVector

`func (o *StaticFinding) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *StaticFinding) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *StaticFinding) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *StaticFinding) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetFileLineNumber

`func (o *StaticFinding) GetFileLineNumber() int32`

GetFileLineNumber returns the FileLineNumber field if non-nil, zero value otherwise.

### GetFileLineNumberOk

`func (o *StaticFinding) GetFileLineNumberOk() (*int32, bool)`

GetFileLineNumberOk returns a tuple with the FileLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLineNumber

`func (o *StaticFinding) SetFileLineNumber(v int32)`

SetFileLineNumber sets FileLineNumber field to given value.

### HasFileLineNumber

`func (o *StaticFinding) HasFileLineNumber() bool`

HasFileLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *StaticFinding) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *StaticFinding) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *StaticFinding) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *StaticFinding) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *StaticFinding) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *StaticFinding) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *StaticFinding) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *StaticFinding) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFindingCategory

`func (o *StaticFinding) GetFindingCategory() StaticFindingFindingCategory`

GetFindingCategory returns the FindingCategory field if non-nil, zero value otherwise.

### GetFindingCategoryOk

`func (o *StaticFinding) GetFindingCategoryOk() (*StaticFindingFindingCategory, bool)`

GetFindingCategoryOk returns a tuple with the FindingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingCategory

`func (o *StaticFinding) SetFindingCategory(v StaticFindingFindingCategory)`

SetFindingCategory sets FindingCategory field to given value.

### HasFindingCategory

`func (o *StaticFinding) HasFindingCategory() bool`

HasFindingCategory returns a boolean if a field has been set.

### GetModule

`func (o *StaticFinding) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *StaticFinding) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *StaticFinding) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *StaticFinding) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetProcedure

`func (o *StaticFinding) GetProcedure() string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *StaticFinding) GetProcedureOk() (*string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *StaticFinding) SetProcedure(v string)`

SetProcedure sets Procedure field to given value.

### HasProcedure

`func (o *StaticFinding) HasProcedure() bool`

HasProcedure returns a boolean if a field has been set.

### GetRelativeLocation

`func (o *StaticFinding) GetRelativeLocation() int32`

GetRelativeLocation returns the RelativeLocation field if non-nil, zero value otherwise.

### GetRelativeLocationOk

`func (o *StaticFinding) GetRelativeLocationOk() (*int32, bool)`

GetRelativeLocationOk returns a tuple with the RelativeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeLocation

`func (o *StaticFinding) SetRelativeLocation(v int32)`

SetRelativeLocation sets RelativeLocation field to given value.

### HasRelativeLocation

`func (o *StaticFinding) HasRelativeLocation() bool`

HasRelativeLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



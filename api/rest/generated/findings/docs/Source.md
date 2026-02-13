# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ScanId** | Pointer to **int32** |  | [optional] 
**CapecId** | Pointer to **int32** | Unique indentifier for the attack category which contains subcategories as (abuse of functionality, spoofing, probabilistic techniques, exploitation of authentication, resource depletion, exploitation of privilege/trust, injection, data structure attacks, data leakage attacks, resource manipulation, and time and state attacks). | [optional] 
**ExploitDesc** | Pointer to **string** | Description of the cause of the manual finding. | [optional] 
**ExploitDifficulty** | Pointer to **int32** | Category of level of effort needed to fix the manual finding. | [optional] 
**InputVector** | Pointer to **string** | URL for the attack vector. | [optional] 
**Location** | Pointer to **string** | Relative location of the manual finding. | [optional] 
**Module** | Pointer to **string** | Module where the manual finding exists. | [optional] 
**RemediationDesc** | Pointer to **string** | Description of remediation needed for the attack vector. | [optional] 
**SeverityDesc** | Pointer to **string** | Description of the severity of the manual finding. | [optional] 
**Note** | Pointer to **string** | Review note of the manual finding. | [optional] 
**AppendixViews** | Pointer to [**[]Appendix**](Appendix.md) |  | [optional] 

## Methods

### NewSource

`func NewSource() *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Source) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Source) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Source) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Source) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScanId

`func (o *Source) GetScanId() int32`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *Source) GetScanIdOk() (*int32, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *Source) SetScanId(v int32)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *Source) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetCapecId

`func (o *Source) GetCapecId() int32`

GetCapecId returns the CapecId field if non-nil, zero value otherwise.

### GetCapecIdOk

`func (o *Source) GetCapecIdOk() (*int32, bool)`

GetCapecIdOk returns a tuple with the CapecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapecId

`func (o *Source) SetCapecId(v int32)`

SetCapecId sets CapecId field to given value.

### HasCapecId

`func (o *Source) HasCapecId() bool`

HasCapecId returns a boolean if a field has been set.

### GetExploitDesc

`func (o *Source) GetExploitDesc() string`

GetExploitDesc returns the ExploitDesc field if non-nil, zero value otherwise.

### GetExploitDescOk

`func (o *Source) GetExploitDescOk() (*string, bool)`

GetExploitDescOk returns a tuple with the ExploitDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitDesc

`func (o *Source) SetExploitDesc(v string)`

SetExploitDesc sets ExploitDesc field to given value.

### HasExploitDesc

`func (o *Source) HasExploitDesc() bool`

HasExploitDesc returns a boolean if a field has been set.

### GetExploitDifficulty

`func (o *Source) GetExploitDifficulty() int32`

GetExploitDifficulty returns the ExploitDifficulty field if non-nil, zero value otherwise.

### GetExploitDifficultyOk

`func (o *Source) GetExploitDifficultyOk() (*int32, bool)`

GetExploitDifficultyOk returns a tuple with the ExploitDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitDifficulty

`func (o *Source) SetExploitDifficulty(v int32)`

SetExploitDifficulty sets ExploitDifficulty field to given value.

### HasExploitDifficulty

`func (o *Source) HasExploitDifficulty() bool`

HasExploitDifficulty returns a boolean if a field has been set.

### GetInputVector

`func (o *Source) GetInputVector() string`

GetInputVector returns the InputVector field if non-nil, zero value otherwise.

### GetInputVectorOk

`func (o *Source) GetInputVectorOk() (*string, bool)`

GetInputVectorOk returns a tuple with the InputVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVector

`func (o *Source) SetInputVector(v string)`

SetInputVector sets InputVector field to given value.

### HasInputVector

`func (o *Source) HasInputVector() bool`

HasInputVector returns a boolean if a field has been set.

### GetLocation

`func (o *Source) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Source) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Source) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Source) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModule

`func (o *Source) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Source) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Source) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *Source) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetRemediationDesc

`func (o *Source) GetRemediationDesc() string`

GetRemediationDesc returns the RemediationDesc field if non-nil, zero value otherwise.

### GetRemediationDescOk

`func (o *Source) GetRemediationDescOk() (*string, bool)`

GetRemediationDescOk returns a tuple with the RemediationDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationDesc

`func (o *Source) SetRemediationDesc(v string)`

SetRemediationDesc sets RemediationDesc field to given value.

### HasRemediationDesc

`func (o *Source) HasRemediationDesc() bool`

HasRemediationDesc returns a boolean if a field has been set.

### GetSeverityDesc

`func (o *Source) GetSeverityDesc() string`

GetSeverityDesc returns the SeverityDesc field if non-nil, zero value otherwise.

### GetSeverityDescOk

`func (o *Source) GetSeverityDescOk() (*string, bool)`

GetSeverityDescOk returns a tuple with the SeverityDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityDesc

`func (o *Source) SetSeverityDesc(v string)`

SetSeverityDesc sets SeverityDesc field to given value.

### HasSeverityDesc

`func (o *Source) HasSeverityDesc() bool`

HasSeverityDesc returns a boolean if a field has been set.

### GetNote

`func (o *Source) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Source) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Source) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Source) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAppendixViews

`func (o *Source) GetAppendixViews() []Appendix`

GetAppendixViews returns the AppendixViews field if non-nil, zero value otherwise.

### GetAppendixViewsOk

`func (o *Source) GetAppendixViewsOk() (*[]Appendix, bool)`

GetAppendixViewsOk returns a tuple with the AppendixViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendixViews

`func (o *Source) SetAppendixViews(v []Appendix)`

SetAppendixViews sets AppendixViews field to given value.

### HasAppendixViews

`func (o *Source) HasAppendixViews() bool`

HasAppendixViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



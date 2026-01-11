# ManualFindingV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwe** | Pointer to [**StaticFindingCwe**](StaticFindingCwe.md) |  | [optional] 
**Cvss** | Pointer to **string** | The CVSS associated with the finding. | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**CapecId** | Pointer to **string** | Attack category, which contains subcategories, such as abuse of functionality, spoofing, probabilistic techniques, exploitation of authentication, resource depletion, exploitation of privilege/trust, injection, data structure attacks, data leakage attacks, resource manipulation, time and state attacks. | [optional] 
**ExploitDesc** | Pointer to **string** | Description of the cause of the finding. | [optional] 
**ExploitDifficulty** | Pointer to **string** | Category of the level of effort needed to fix the finding. | [optional] 
**InputVector** | Pointer to **string** | URL for the attack vector. | [optional] 
**Location** | Pointer to **string** | The relative location of finding. | [optional] 
**Module** | Pointer to **string** | The module where the finding exists. | [optional] 
**RemediationDesc** | Pointer to **string** | Description of the remediation needed for the attack vector. | [optional] 
**SeverityDesc** | Pointer to **string** | Description of the severity of finding. | [optional] 

## Methods

### NewManualFindingV2

`func NewManualFindingV2() *ManualFindingV2`

NewManualFindingV2 instantiates a new ManualFindingV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualFindingV2WithDefaults

`func NewManualFindingV2WithDefaults() *ManualFindingV2`

NewManualFindingV2WithDefaults instantiates a new ManualFindingV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwe

`func (o *ManualFindingV2) GetCwe() StaticFindingCwe`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *ManualFindingV2) GetCweOk() (*StaticFindingCwe, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *ManualFindingV2) SetCwe(v StaticFindingCwe)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *ManualFindingV2) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetCvss

`func (o *ManualFindingV2) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *ManualFindingV2) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *ManualFindingV2) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *ManualFindingV2) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetSeverity

`func (o *ManualFindingV2) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ManualFindingV2) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ManualFindingV2) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ManualFindingV2) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCapecId

`func (o *ManualFindingV2) GetCapecId() string`

GetCapecId returns the CapecId field if non-nil, zero value otherwise.

### GetCapecIdOk

`func (o *ManualFindingV2) GetCapecIdOk() (*string, bool)`

GetCapecIdOk returns a tuple with the CapecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapecId

`func (o *ManualFindingV2) SetCapecId(v string)`

SetCapecId sets CapecId field to given value.

### HasCapecId

`func (o *ManualFindingV2) HasCapecId() bool`

HasCapecId returns a boolean if a field has been set.

### GetExploitDesc

`func (o *ManualFindingV2) GetExploitDesc() string`

GetExploitDesc returns the ExploitDesc field if non-nil, zero value otherwise.

### GetExploitDescOk

`func (o *ManualFindingV2) GetExploitDescOk() (*string, bool)`

GetExploitDescOk returns a tuple with the ExploitDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitDesc

`func (o *ManualFindingV2) SetExploitDesc(v string)`

SetExploitDesc sets ExploitDesc field to given value.

### HasExploitDesc

`func (o *ManualFindingV2) HasExploitDesc() bool`

HasExploitDesc returns a boolean if a field has been set.

### GetExploitDifficulty

`func (o *ManualFindingV2) GetExploitDifficulty() string`

GetExploitDifficulty returns the ExploitDifficulty field if non-nil, zero value otherwise.

### GetExploitDifficultyOk

`func (o *ManualFindingV2) GetExploitDifficultyOk() (*string, bool)`

GetExploitDifficultyOk returns a tuple with the ExploitDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitDifficulty

`func (o *ManualFindingV2) SetExploitDifficulty(v string)`

SetExploitDifficulty sets ExploitDifficulty field to given value.

### HasExploitDifficulty

`func (o *ManualFindingV2) HasExploitDifficulty() bool`

HasExploitDifficulty returns a boolean if a field has been set.

### GetInputVector

`func (o *ManualFindingV2) GetInputVector() string`

GetInputVector returns the InputVector field if non-nil, zero value otherwise.

### GetInputVectorOk

`func (o *ManualFindingV2) GetInputVectorOk() (*string, bool)`

GetInputVectorOk returns a tuple with the InputVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVector

`func (o *ManualFindingV2) SetInputVector(v string)`

SetInputVector sets InputVector field to given value.

### HasInputVector

`func (o *ManualFindingV2) HasInputVector() bool`

HasInputVector returns a boolean if a field has been set.

### GetLocation

`func (o *ManualFindingV2) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ManualFindingV2) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ManualFindingV2) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ManualFindingV2) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModule

`func (o *ManualFindingV2) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ManualFindingV2) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ManualFindingV2) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ManualFindingV2) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetRemediationDesc

`func (o *ManualFindingV2) GetRemediationDesc() string`

GetRemediationDesc returns the RemediationDesc field if non-nil, zero value otherwise.

### GetRemediationDescOk

`func (o *ManualFindingV2) GetRemediationDescOk() (*string, bool)`

GetRemediationDescOk returns a tuple with the RemediationDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationDesc

`func (o *ManualFindingV2) SetRemediationDesc(v string)`

SetRemediationDesc sets RemediationDesc field to given value.

### HasRemediationDesc

`func (o *ManualFindingV2) HasRemediationDesc() bool`

HasRemediationDesc returns a boolean if a field has been set.

### GetSeverityDesc

`func (o *ManualFindingV2) GetSeverityDesc() string`

GetSeverityDesc returns the SeverityDesc field if non-nil, zero value otherwise.

### GetSeverityDescOk

`func (o *ManualFindingV2) GetSeverityDescOk() (*string, bool)`

GetSeverityDescOk returns a tuple with the SeverityDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityDesc

`func (o *ManualFindingV2) SetSeverityDesc(v string)`

SetSeverityDesc sets SeverityDesc field to given value.

### HasSeverityDesc

`func (o *ManualFindingV2) HasSeverityDesc() bool`

HasSeverityDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



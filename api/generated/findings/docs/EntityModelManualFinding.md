# EntityModelManualFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier (Long). | [optional] 
**ExternalId** | Pointer to **int64** | Alternative identifier of application finding that is unique to this application. | [optional] 
**ScanId** | Pointer to **int32** | Scan identifier of this finding. | [optional] 
**Type** | Pointer to **string** | Internal classification of the finding. | [optional] 
**Description** | Pointer to **string** | Detailed description of the finding. | [optional] 
**Count** | Pointer to **int32** | Number of duplicate findings found in all modules. | [optional] 
**Severity** | Pointer to **int32** | Severity of the finding. | [optional] 
**Exploitability** | Pointer to **int32** | Exploitability of the finding. | [optional] 
**Cwe** | Pointer to **int32** | The CWE identifier. | [optional] 
**Cvss** | Pointer to **float64** | The CVSS score. | [optional] 
**Resolution** | Pointer to **string** | Resolution of the finding. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** | Date the scan finding was found. | [optional] 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**MatchedId** | Pointer to **int64** | Identifier that matches this finding. | [optional] 
**Appendix** | Pointer to [**[]Appendix**](Appendix.md) |  | [optional] 

## Methods

### NewEntityModelManualFinding

`func NewEntityModelManualFinding() *EntityModelManualFinding`

NewEntityModelManualFinding instantiates a new EntityModelManualFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityModelManualFindingWithDefaults

`func NewEntityModelManualFindingWithDefaults() *EntityModelManualFinding`

NewEntityModelManualFindingWithDefaults instantiates a new EntityModelManualFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityModelManualFinding) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityModelManualFinding) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityModelManualFinding) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EntityModelManualFinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *EntityModelManualFinding) GetExternalId() int64`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntityModelManualFinding) GetExternalIdOk() (*int64, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntityModelManualFinding) SetExternalId(v int64)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntityModelManualFinding) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetScanId

`func (o *EntityModelManualFinding) GetScanId() int32`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *EntityModelManualFinding) GetScanIdOk() (*int32, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *EntityModelManualFinding) SetScanId(v int32)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *EntityModelManualFinding) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetType

`func (o *EntityModelManualFinding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityModelManualFinding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityModelManualFinding) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityModelManualFinding) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *EntityModelManualFinding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityModelManualFinding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityModelManualFinding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityModelManualFinding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCount

`func (o *EntityModelManualFinding) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EntityModelManualFinding) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EntityModelManualFinding) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EntityModelManualFinding) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSeverity

`func (o *EntityModelManualFinding) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EntityModelManualFinding) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EntityModelManualFinding) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EntityModelManualFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetExploitability

`func (o *EntityModelManualFinding) GetExploitability() int32`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *EntityModelManualFinding) GetExploitabilityOk() (*int32, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *EntityModelManualFinding) SetExploitability(v int32)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *EntityModelManualFinding) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetCwe

`func (o *EntityModelManualFinding) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *EntityModelManualFinding) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *EntityModelManualFinding) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *EntityModelManualFinding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetCvss

`func (o *EntityModelManualFinding) GetCvss() float64`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *EntityModelManualFinding) GetCvssOk() (*float64, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *EntityModelManualFinding) SetCvss(v float64)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *EntityModelManualFinding) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetResolution

`func (o *EntityModelManualFinding) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *EntityModelManualFinding) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *EntityModelManualFinding) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *EntityModelManualFinding) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetState

`func (o *EntityModelManualFinding) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EntityModelManualFinding) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EntityModelManualFinding) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EntityModelManualFinding) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDate

`func (o *EntityModelManualFinding) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EntityModelManualFinding) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EntityModelManualFinding) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *EntityModelManualFinding) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSource

`func (o *EntityModelManualFinding) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EntityModelManualFinding) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EntityModelManualFinding) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *EntityModelManualFinding) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMatchedId

`func (o *EntityModelManualFinding) GetMatchedId() int64`

GetMatchedId returns the MatchedId field if non-nil, zero value otherwise.

### GetMatchedIdOk

`func (o *EntityModelManualFinding) GetMatchedIdOk() (*int64, bool)`

GetMatchedIdOk returns a tuple with the MatchedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedId

`func (o *EntityModelManualFinding) SetMatchedId(v int64)`

SetMatchedId sets MatchedId field to given value.

### HasMatchedId

`func (o *EntityModelManualFinding) HasMatchedId() bool`

HasMatchedId returns a boolean if a field has been set.

### GetAppendix

`func (o *EntityModelManualFinding) GetAppendix() []Appendix`

GetAppendix returns the Appendix field if non-nil, zero value otherwise.

### GetAppendixOk

`func (o *EntityModelManualFinding) GetAppendixOk() (*[]Appendix, bool)`

GetAppendixOk returns a tuple with the Appendix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendix

`func (o *EntityModelManualFinding) SetAppendix(v []Appendix)`

SetAppendix sets Appendix field to given value.

### HasAppendix

`func (o *EntityModelManualFinding) HasAppendix() bool`

HasAppendix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FindingFindingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwe** | Pointer to [**ScaFindingCwe**](ScaFindingCwe.md) |  | [optional] 
**Cvss** | Pointer to **float64** | The CVSS score. | [optional] 
**Severity** | Pointer to **int32** | An assigned severity for the vulnerability. | [optional] 
**Exploitability** | Pointer to **int32** | Exploitability of the finding. | [optional] 
**AttackVector** | Pointer to **string** | URL or some input parameter. | [optional] 
**FileLineNumber** | Pointer to **int32** | The line number where the finding exists in the file. | [optional] 
**FileName** | Pointer to **string** | The name of the file where the finding exists. | [optional] 
**FilePath** | Pointer to **string** | The path to the file where the finding exists. | [optional] 
**FindingCategory** | Pointer to [**StaticFindingFindingCategory**](StaticFindingFindingCategory.md) |  | [optional] 
**Module** | Pointer to **string** | The name of the module where the finding exists. | [optional] 
**Procedure** | Pointer to **string** | The name of the procedure where the finding exists. | [optional] 
**RelativeLocation** | Pointer to **int32** | The relative location of the finding in the procedure. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the URL that contains the vulnerability. | [optional] 
**Port** | Pointer to **string** | The port of the hostname that was attacked. | [optional] 
**Path** | Pointer to **string** | The URI path. | [optional] 
**Plugin** | Pointer to **string** | The type of attack sent. | [optional] 
**URL** | Pointer to **string** | The URL of the location where the finding exists. | [optional] 
**VulnerableParameter** | Pointer to **string** | The parameter that contains a vulnerability. | [optional] 
**DiscoveredByVsa** | Pointer to **int32** | Whether the finding was discovered by Virtual Scan Appliance. | [optional] 
**Id** | Pointer to **int64** | Unique identifier (Long). | [optional] 
**ExternalId** | Pointer to **int64** | Alternative identifier of application finding that is unique to this application. | [optional] 
**ScanId** | Pointer to **int32** | Scan identifier of this finding. | [optional] 
**Type** | Pointer to **string** | Internal classification of the finding. | [optional] 
**Description** | Pointer to **string** | Detailed description of the finding. | [optional] 
**Count** | Pointer to **int32** | Number of duplicate findings found in all modules. | [optional] 
**Severity** | Pointer to **int32** | Severity of the finding. | [optional] 
**Resolution** | Pointer to **string** | Resolution of the finding. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** | Date the scan finding was found. | [optional] 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**MatchedId** | Pointer to **int64** | Identifier that matches this finding. | [optional] 
**Appendix** | Pointer to [**[]Appendix**](Appendix.md) |  | [optional] 
**ComponentId** | Pointer to **string** | Unique ID for the component. | [optional] 
**Licenses** | Pointer to [**[]ScaFindingLicensesInner**](ScaFindingLicensesInner.md) | Displays all licenses found for a component with the specified risk rating. | [optional] 
**Cve** | Pointer to [**ScaFindingCve**](ScaFindingCve.md) |  | [optional] 
**Version** | Pointer to **string** | The version of the third-party component. | [optional] 
**ProductId** | Pointer to **string** | The product ID containing the vulnerability. | [optional] 
**ComponentFilename** | Pointer to **string** | The component filename. | [optional] 
**Language** | Pointer to **string** | The coding language. | [optional] 
**ComponentPathS** | Pointer to [**[]ScaFindingComponentPathSInner**](ScaFindingComponentPathSInner.md) | The list of component paths containing this vulnerability. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Displays metadata values, such as the SCA scan mode and dependency mode. | [optional] 

## Methods

### NewFindingFindingDetails

`func NewFindingFindingDetails() *FindingFindingDetails`

NewFindingFindingDetails instantiates a new FindingFindingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingFindingDetailsWithDefaults

`func NewFindingFindingDetailsWithDefaults() *FindingFindingDetails`

NewFindingFindingDetailsWithDefaults instantiates a new FindingFindingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwe

`func (o *FindingFindingDetails) GetCwe() ScaFindingCwe`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *FindingFindingDetails) GetCweOk() (*ScaFindingCwe, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *FindingFindingDetails) SetCwe(v ScaFindingCwe)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *FindingFindingDetails) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetCvss

`func (o *FindingFindingDetails) GetCvss() float64`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *FindingFindingDetails) GetCvssOk() (*float64, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *FindingFindingDetails) SetCvss(v float64)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *FindingFindingDetails) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetSeverity

`func (o *FindingFindingDetails) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingFindingDetails) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingFindingDetails) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FindingFindingDetails) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetExploitability

`func (o *FindingFindingDetails) GetExploitability() int32`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *FindingFindingDetails) GetExploitabilityOk() (*int32, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *FindingFindingDetails) SetExploitability(v int32)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *FindingFindingDetails) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetAttackVector

`func (o *FindingFindingDetails) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *FindingFindingDetails) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *FindingFindingDetails) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *FindingFindingDetails) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetFileLineNumber

`func (o *FindingFindingDetails) GetFileLineNumber() int32`

GetFileLineNumber returns the FileLineNumber field if non-nil, zero value otherwise.

### GetFileLineNumberOk

`func (o *FindingFindingDetails) GetFileLineNumberOk() (*int32, bool)`

GetFileLineNumberOk returns a tuple with the FileLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLineNumber

`func (o *FindingFindingDetails) SetFileLineNumber(v int32)`

SetFileLineNumber sets FileLineNumber field to given value.

### HasFileLineNumber

`func (o *FindingFindingDetails) HasFileLineNumber() bool`

HasFileLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *FindingFindingDetails) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FindingFindingDetails) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FindingFindingDetails) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FindingFindingDetails) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *FindingFindingDetails) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FindingFindingDetails) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FindingFindingDetails) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FindingFindingDetails) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFindingCategory

`func (o *FindingFindingDetails) GetFindingCategory() StaticFindingFindingCategory`

GetFindingCategory returns the FindingCategory field if non-nil, zero value otherwise.

### GetFindingCategoryOk

`func (o *FindingFindingDetails) GetFindingCategoryOk() (*StaticFindingFindingCategory, bool)`

GetFindingCategoryOk returns a tuple with the FindingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingCategory

`func (o *FindingFindingDetails) SetFindingCategory(v StaticFindingFindingCategory)`

SetFindingCategory sets FindingCategory field to given value.

### HasFindingCategory

`func (o *FindingFindingDetails) HasFindingCategory() bool`

HasFindingCategory returns a boolean if a field has been set.

### GetModule

`func (o *FindingFindingDetails) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *FindingFindingDetails) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *FindingFindingDetails) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *FindingFindingDetails) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetProcedure

`func (o *FindingFindingDetails) GetProcedure() string`

GetProcedure returns the Procedure field if non-nil, zero value otherwise.

### GetProcedureOk

`func (o *FindingFindingDetails) GetProcedureOk() (*string, bool)`

GetProcedureOk returns a tuple with the Procedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedure

`func (o *FindingFindingDetails) SetProcedure(v string)`

SetProcedure sets Procedure field to given value.

### HasProcedure

`func (o *FindingFindingDetails) HasProcedure() bool`

HasProcedure returns a boolean if a field has been set.

### GetRelativeLocation

`func (o *FindingFindingDetails) GetRelativeLocation() int32`

GetRelativeLocation returns the RelativeLocation field if non-nil, zero value otherwise.

### GetRelativeLocationOk

`func (o *FindingFindingDetails) GetRelativeLocationOk() (*int32, bool)`

GetRelativeLocationOk returns a tuple with the RelativeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeLocation

`func (o *FindingFindingDetails) SetRelativeLocation(v int32)`

SetRelativeLocation sets RelativeLocation field to given value.

### HasRelativeLocation

`func (o *FindingFindingDetails) HasRelativeLocation() bool`

HasRelativeLocation returns a boolean if a field has been set.

### GetHostname

`func (o *FindingFindingDetails) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FindingFindingDetails) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FindingFindingDetails) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FindingFindingDetails) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *FindingFindingDetails) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FindingFindingDetails) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FindingFindingDetails) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *FindingFindingDetails) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *FindingFindingDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FindingFindingDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FindingFindingDetails) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FindingFindingDetails) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlugin

`func (o *FindingFindingDetails) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *FindingFindingDetails) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *FindingFindingDetails) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *FindingFindingDetails) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetURL

`func (o *FindingFindingDetails) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *FindingFindingDetails) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *FindingFindingDetails) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *FindingFindingDetails) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetVulnerableParameter

`func (o *FindingFindingDetails) GetVulnerableParameter() string`

GetVulnerableParameter returns the VulnerableParameter field if non-nil, zero value otherwise.

### GetVulnerableParameterOk

`func (o *FindingFindingDetails) GetVulnerableParameterOk() (*string, bool)`

GetVulnerableParameterOk returns a tuple with the VulnerableParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableParameter

`func (o *FindingFindingDetails) SetVulnerableParameter(v string)`

SetVulnerableParameter sets VulnerableParameter field to given value.

### HasVulnerableParameter

`func (o *FindingFindingDetails) HasVulnerableParameter() bool`

HasVulnerableParameter returns a boolean if a field has been set.

### GetDiscoveredByVsa

`func (o *FindingFindingDetails) GetDiscoveredByVsa() int32`

GetDiscoveredByVsa returns the DiscoveredByVsa field if non-nil, zero value otherwise.

### GetDiscoveredByVsaOk

`func (o *FindingFindingDetails) GetDiscoveredByVsaOk() (*int32, bool)`

GetDiscoveredByVsaOk returns a tuple with the DiscoveredByVsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredByVsa

`func (o *FindingFindingDetails) SetDiscoveredByVsa(v int32)`

SetDiscoveredByVsa sets DiscoveredByVsa field to given value.

### HasDiscoveredByVsa

`func (o *FindingFindingDetails) HasDiscoveredByVsa() bool`

HasDiscoveredByVsa returns a boolean if a field has been set.

### GetId

`func (o *FindingFindingDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingFindingDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingFindingDetails) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FindingFindingDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *FindingFindingDetails) GetExternalId() int64`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FindingFindingDetails) GetExternalIdOk() (*int64, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FindingFindingDetails) SetExternalId(v int64)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FindingFindingDetails) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetScanId

`func (o *FindingFindingDetails) GetScanId() int32`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *FindingFindingDetails) GetScanIdOk() (*int32, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *FindingFindingDetails) SetScanId(v int32)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *FindingFindingDetails) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetType

`func (o *FindingFindingDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindingFindingDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindingFindingDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindingFindingDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *FindingFindingDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingFindingDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingFindingDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindingFindingDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCount

`func (o *FindingFindingDetails) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FindingFindingDetails) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FindingFindingDetails) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FindingFindingDetails) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSeverity

`func (o *FindingFindingDetails) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingFindingDetails) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingFindingDetails) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FindingFindingDetails) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolution

`func (o *FindingFindingDetails) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *FindingFindingDetails) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *FindingFindingDetails) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *FindingFindingDetails) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetState

`func (o *FindingFindingDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindingFindingDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindingFindingDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FindingFindingDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDate

`func (o *FindingFindingDetails) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FindingFindingDetails) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FindingFindingDetails) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *FindingFindingDetails) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSource

`func (o *FindingFindingDetails) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FindingFindingDetails) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FindingFindingDetails) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *FindingFindingDetails) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMatchedId

`func (o *FindingFindingDetails) GetMatchedId() int64`

GetMatchedId returns the MatchedId field if non-nil, zero value otherwise.

### GetMatchedIdOk

`func (o *FindingFindingDetails) GetMatchedIdOk() (*int64, bool)`

GetMatchedIdOk returns a tuple with the MatchedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedId

`func (o *FindingFindingDetails) SetMatchedId(v int64)`

SetMatchedId sets MatchedId field to given value.

### HasMatchedId

`func (o *FindingFindingDetails) HasMatchedId() bool`

HasMatchedId returns a boolean if a field has been set.

### GetAppendix

`func (o *FindingFindingDetails) GetAppendix() []Appendix`

GetAppendix returns the Appendix field if non-nil, zero value otherwise.

### GetAppendixOk

`func (o *FindingFindingDetails) GetAppendixOk() (*[]Appendix, bool)`

GetAppendixOk returns a tuple with the Appendix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendix

`func (o *FindingFindingDetails) SetAppendix(v []Appendix)`

SetAppendix sets Appendix field to given value.

### HasAppendix

`func (o *FindingFindingDetails) HasAppendix() bool`

HasAppendix returns a boolean if a field has been set.

### GetComponentId

`func (o *FindingFindingDetails) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *FindingFindingDetails) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *FindingFindingDetails) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *FindingFindingDetails) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetLicenses

`func (o *FindingFindingDetails) GetLicenses() []ScaFindingLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *FindingFindingDetails) GetLicensesOk() (*[]ScaFindingLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *FindingFindingDetails) SetLicenses(v []ScaFindingLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *FindingFindingDetails) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetCve

`func (o *FindingFindingDetails) GetCve() ScaFindingCve`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *FindingFindingDetails) GetCveOk() (*ScaFindingCve, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *FindingFindingDetails) SetCve(v ScaFindingCve)`

SetCve sets Cve field to given value.

### HasCve

`func (o *FindingFindingDetails) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetVersion

`func (o *FindingFindingDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FindingFindingDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FindingFindingDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FindingFindingDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProductId

`func (o *FindingFindingDetails) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *FindingFindingDetails) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *FindingFindingDetails) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *FindingFindingDetails) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetComponentFilename

`func (o *FindingFindingDetails) GetComponentFilename() string`

GetComponentFilename returns the ComponentFilename field if non-nil, zero value otherwise.

### GetComponentFilenameOk

`func (o *FindingFindingDetails) GetComponentFilenameOk() (*string, bool)`

GetComponentFilenameOk returns a tuple with the ComponentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentFilename

`func (o *FindingFindingDetails) SetComponentFilename(v string)`

SetComponentFilename sets ComponentFilename field to given value.

### HasComponentFilename

`func (o *FindingFindingDetails) HasComponentFilename() bool`

HasComponentFilename returns a boolean if a field has been set.

### GetLanguage

`func (o *FindingFindingDetails) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FindingFindingDetails) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FindingFindingDetails) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FindingFindingDetails) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetComponentPathS

`func (o *FindingFindingDetails) GetComponentPathS() []ScaFindingComponentPathSInner`

GetComponentPathS returns the ComponentPathS field if non-nil, zero value otherwise.

### GetComponentPathSOk

`func (o *FindingFindingDetails) GetComponentPathSOk() (*[]ScaFindingComponentPathSInner, bool)`

GetComponentPathSOk returns a tuple with the ComponentPathS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPathS

`func (o *FindingFindingDetails) SetComponentPathS(v []ScaFindingComponentPathSInner)`

SetComponentPathS sets ComponentPathS field to given value.

### HasComponentPathS

`func (o *FindingFindingDetails) HasComponentPathS() bool`

HasComponentPathS returns a boolean if a field has been set.

### GetMetadata

`func (o *FindingFindingDetails) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FindingFindingDetails) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FindingFindingDetails) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FindingFindingDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



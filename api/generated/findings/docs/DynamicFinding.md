# DynamicFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwe** | Pointer to [**StaticFindingCwe**](StaticFindingCwe.md) |  | [optional] 
**Cvss** | Pointer to **string** | The CVSS associated with the finding. | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**AttackVector** | Pointer to **string** | URL or some input parameter. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the URL that contains the vulnerability. | [optional] 
**Port** | Pointer to **string** | The port of the hostname that was attacked. | [optional] 
**Path** | Pointer to **string** | The URI path. | [optional] 
**Plugin** | Pointer to **string** | The type of attack sent. | [optional] 
**FindingCategory** | Pointer to [**StaticFindingFindingCategory**](StaticFindingFindingCategory.md) |  | [optional] 
**URL** | Pointer to **string** | The URL of the location where the finding exists. | [optional] 
**VulnerableParameter** | Pointer to **string** | The parameter that contains a vulnerability. | [optional] 
**DiscoveredByVsa** | Pointer to **string** | Whether the finding was discovered by Virtual Scan Appliance. | [optional] 

## Methods

### NewDynamicFinding

`func NewDynamicFinding() *DynamicFinding`

NewDynamicFinding instantiates a new DynamicFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicFindingWithDefaults

`func NewDynamicFindingWithDefaults() *DynamicFinding`

NewDynamicFindingWithDefaults instantiates a new DynamicFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwe

`func (o *DynamicFinding) GetCwe() StaticFindingCwe`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *DynamicFinding) GetCweOk() (*StaticFindingCwe, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *DynamicFinding) SetCwe(v StaticFindingCwe)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *DynamicFinding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetCvss

`func (o *DynamicFinding) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *DynamicFinding) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *DynamicFinding) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *DynamicFinding) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetSeverity

`func (o *DynamicFinding) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DynamicFinding) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DynamicFinding) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *DynamicFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAttackVector

`func (o *DynamicFinding) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *DynamicFinding) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *DynamicFinding) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *DynamicFinding) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetHostname

`func (o *DynamicFinding) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DynamicFinding) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DynamicFinding) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DynamicFinding) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *DynamicFinding) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DynamicFinding) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DynamicFinding) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DynamicFinding) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *DynamicFinding) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DynamicFinding) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DynamicFinding) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DynamicFinding) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlugin

`func (o *DynamicFinding) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *DynamicFinding) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *DynamicFinding) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *DynamicFinding) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetFindingCategory

`func (o *DynamicFinding) GetFindingCategory() StaticFindingFindingCategory`

GetFindingCategory returns the FindingCategory field if non-nil, zero value otherwise.

### GetFindingCategoryOk

`func (o *DynamicFinding) GetFindingCategoryOk() (*StaticFindingFindingCategory, bool)`

GetFindingCategoryOk returns a tuple with the FindingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingCategory

`func (o *DynamicFinding) SetFindingCategory(v StaticFindingFindingCategory)`

SetFindingCategory sets FindingCategory field to given value.

### HasFindingCategory

`func (o *DynamicFinding) HasFindingCategory() bool`

HasFindingCategory returns a boolean if a field has been set.

### GetURL

`func (o *DynamicFinding) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *DynamicFinding) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *DynamicFinding) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *DynamicFinding) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetVulnerableParameter

`func (o *DynamicFinding) GetVulnerableParameter() string`

GetVulnerableParameter returns the VulnerableParameter field if non-nil, zero value otherwise.

### GetVulnerableParameterOk

`func (o *DynamicFinding) GetVulnerableParameterOk() (*string, bool)`

GetVulnerableParameterOk returns a tuple with the VulnerableParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableParameter

`func (o *DynamicFinding) SetVulnerableParameter(v string)`

SetVulnerableParameter sets VulnerableParameter field to given value.

### HasVulnerableParameter

`func (o *DynamicFinding) HasVulnerableParameter() bool`

HasVulnerableParameter returns a boolean if a field has been set.

### GetDiscoveredByVsa

`func (o *DynamicFinding) GetDiscoveredByVsa() string`

GetDiscoveredByVsa returns the DiscoveredByVsa field if non-nil, zero value otherwise.

### GetDiscoveredByVsaOk

`func (o *DynamicFinding) GetDiscoveredByVsaOk() (*string, bool)`

GetDiscoveredByVsaOk returns a tuple with the DiscoveredByVsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredByVsa

`func (o *DynamicFinding) SetDiscoveredByVsa(v string)`

SetDiscoveredByVsa sets DiscoveredByVsa field to given value.

### HasDiscoveredByVsa

`func (o *DynamicFinding) HasDiscoveredByVsa() bool`

HasDiscoveredByVsa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



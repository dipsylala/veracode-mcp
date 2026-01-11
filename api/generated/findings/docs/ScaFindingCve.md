# ScaFindingCve

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The identifier for this vulnerability. While typically a CVE, this field can also identify non-CVE assigned vulnerabilities, such as SRCCLR-SID-23134. | [optional] 
**Cvss** | Pointer to **float32** | The Common Vulnerability Scoring System score for this CVE. | [optional] 
**Href** | Pointer to **string** | A link to the CVE in the NVD or other database. | [optional] 
**Severity** | Pointer to **int32** | The assigned severity of this vulnerability. | [optional] 
**Vector** | Pointer to **string** | The assigned vector for this vulnerability. | [optional] 
**Cvss3** | Pointer to [**ScaFindingCveCvss3**](ScaFindingCveCvss3.md) |  | [optional] 
**Exploitability** | Pointer to [**ScaFindingExploitability**](ScaFindingExploitability.md) |  | [optional] 

## Methods

### NewScaFindingCve

`func NewScaFindingCve() *ScaFindingCve`

NewScaFindingCve instantiates a new ScaFindingCve object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaFindingCveWithDefaults

`func NewScaFindingCveWithDefaults() *ScaFindingCve`

NewScaFindingCveWithDefaults instantiates a new ScaFindingCve object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScaFindingCve) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScaFindingCve) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScaFindingCve) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScaFindingCve) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCvss

`func (o *ScaFindingCve) GetCvss() float32`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *ScaFindingCve) GetCvssOk() (*float32, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *ScaFindingCve) SetCvss(v float32)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *ScaFindingCve) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetHref

`func (o *ScaFindingCve) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ScaFindingCve) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ScaFindingCve) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ScaFindingCve) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSeverity

`func (o *ScaFindingCve) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ScaFindingCve) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ScaFindingCve) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ScaFindingCve) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetVector

`func (o *ScaFindingCve) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *ScaFindingCve) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *ScaFindingCve) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *ScaFindingCve) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetCvss3

`func (o *ScaFindingCve) GetCvss3() ScaFindingCveCvss3`

GetCvss3 returns the Cvss3 field if non-nil, zero value otherwise.

### GetCvss3Ok

`func (o *ScaFindingCve) GetCvss3Ok() (*ScaFindingCveCvss3, bool)`

GetCvss3Ok returns a tuple with the Cvss3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3

`func (o *ScaFindingCve) SetCvss3(v ScaFindingCveCvss3)`

SetCvss3 sets Cvss3 field to given value.

### HasCvss3

`func (o *ScaFindingCve) HasCvss3() bool`

HasCvss3 returns a boolean if a field has been set.

### GetExploitability

`func (o *ScaFindingCve) GetExploitability() ScaFindingExploitability`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *ScaFindingCve) GetExploitabilityOk() (*ScaFindingExploitability, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *ScaFindingCve) SetExploitability(v ScaFindingExploitability)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *ScaFindingCve) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



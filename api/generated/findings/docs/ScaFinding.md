# ScaFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwe** | Pointer to [**ScaFindingCwe**](ScaFindingCwe.md) |  | [optional] 
**Severity** | Pointer to **float32** | An assigned severity for the vulnerability. | [optional] 
**ComponentId** | Pointer to **string** | Unique ID for the component. | [optional] 
**Licenses** | Pointer to [**[]ScaFindingLicensesInner**](ScaFindingLicensesInner.md) | Displays all licenses found for a component with the specified risk rating. | [optional] 
**Cve** | Pointer to [**ScaFindingCve**](ScaFindingCve.md) |  | [optional] 
**Version** | Pointer to **string** | The version of the third-party component. | [optional] 
**ProductId** | Pointer to **string** | The product ID containing the vulnerability. | [optional] 
**ComponentFilename** | Pointer to **string** | The component filename. | [optional] 
**Language** | Pointer to **string** | The coding language. | [optional] 
**ComponentPathS** | Pointer to [**[]ScaFindingComponentPathSInner**](ScaFindingComponentPathSInner.md) | The list of component paths containing this vulnerability. | [optional] 
**Metadata** | Pointer to **string** | Displays metadata values, such as the SCA scan mode and dependency mode. | [optional] 

## Methods

### NewScaFinding

`func NewScaFinding() *ScaFinding`

NewScaFinding instantiates a new ScaFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaFindingWithDefaults

`func NewScaFindingWithDefaults() *ScaFinding`

NewScaFindingWithDefaults instantiates a new ScaFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwe

`func (o *ScaFinding) GetCwe() ScaFindingCwe`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *ScaFinding) GetCweOk() (*ScaFindingCwe, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *ScaFinding) SetCwe(v ScaFindingCwe)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *ScaFinding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetSeverity

`func (o *ScaFinding) GetSeverity() float32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ScaFinding) GetSeverityOk() (*float32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ScaFinding) SetSeverity(v float32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ScaFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetComponentId

`func (o *ScaFinding) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ScaFinding) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ScaFinding) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ScaFinding) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetLicenses

`func (o *ScaFinding) GetLicenses() []ScaFindingLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ScaFinding) GetLicensesOk() (*[]ScaFindingLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ScaFinding) SetLicenses(v []ScaFindingLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ScaFinding) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetCve

`func (o *ScaFinding) GetCve() ScaFindingCve`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ScaFinding) GetCveOk() (*ScaFindingCve, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ScaFinding) SetCve(v ScaFindingCve)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ScaFinding) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetVersion

`func (o *ScaFinding) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScaFinding) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScaFinding) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ScaFinding) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProductId

`func (o *ScaFinding) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ScaFinding) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ScaFinding) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ScaFinding) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetComponentFilename

`func (o *ScaFinding) GetComponentFilename() string`

GetComponentFilename returns the ComponentFilename field if non-nil, zero value otherwise.

### GetComponentFilenameOk

`func (o *ScaFinding) GetComponentFilenameOk() (*string, bool)`

GetComponentFilenameOk returns a tuple with the ComponentFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentFilename

`func (o *ScaFinding) SetComponentFilename(v string)`

SetComponentFilename sets ComponentFilename field to given value.

### HasComponentFilename

`func (o *ScaFinding) HasComponentFilename() bool`

HasComponentFilename returns a boolean if a field has been set.

### GetLanguage

`func (o *ScaFinding) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ScaFinding) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ScaFinding) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ScaFinding) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetComponentPathS

`func (o *ScaFinding) GetComponentPathS() []ScaFindingComponentPathSInner`

GetComponentPathS returns the ComponentPathS field if non-nil, zero value otherwise.

### GetComponentPathSOk

`func (o *ScaFinding) GetComponentPathSOk() (*[]ScaFindingComponentPathSInner, bool)`

GetComponentPathSOk returns a tuple with the ComponentPathS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPathS

`func (o *ScaFinding) SetComponentPathS(v []ScaFindingComponentPathSInner)`

SetComponentPathS sets ComponentPathS field to given value.

### HasComponentPathS

`func (o *ScaFinding) HasComponentPathS() bool`

HasComponentPathS returns a boolean if a field has been set.

### GetMetadata

`func (o *ScaFinding) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ScaFinding) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ScaFinding) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ScaFinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



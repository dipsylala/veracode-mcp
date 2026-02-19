# ScaLicenseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **string** | Full name of the license. | [optional] 
**Name** | Pointer to **string** | Short name of the license. | [optional] 
**Risk** | Pointer to **string** | Risk rating of the license. Values are Low, Medium, High, or Unknown. | [optional] 
**SpdxId** | **string** | SPDX identifier for the license. | 
**Url** | Pointer to **string** | URL to the license on the spdx.org website. | [optional] 

## Methods

### NewScaLicenseSummary

`func NewScaLicenseSummary(spdxId string, ) *ScaLicenseSummary`

NewScaLicenseSummary instantiates a new ScaLicenseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaLicenseSummaryWithDefaults

`func NewScaLicenseSummaryWithDefaults() *ScaLicenseSummary`

NewScaLicenseSummaryWithDefaults instantiates a new ScaLicenseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ScaLicenseSummary) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ScaLicenseSummary) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ScaLicenseSummary) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ScaLicenseSummary) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetName

`func (o *ScaLicenseSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScaLicenseSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScaLicenseSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScaLicenseSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRisk

`func (o *ScaLicenseSummary) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *ScaLicenseSummary) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *ScaLicenseSummary) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *ScaLicenseSummary) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetSpdxId

`func (o *ScaLicenseSummary) GetSpdxId() string`

GetSpdxId returns the SpdxId field if non-nil, zero value otherwise.

### GetSpdxIdOk

`func (o *ScaLicenseSummary) GetSpdxIdOk() (*string, bool)`

GetSpdxIdOk returns a tuple with the SpdxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxId

`func (o *ScaLicenseSummary) SetSpdxId(v string)`

SetSpdxId sets SpdxId field to given value.


### GetUrl

`func (o *ScaLicenseSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScaLicenseSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScaLicenseSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ScaLicenseSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



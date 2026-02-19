# ScaGracePeriods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScaBlacklistGracePeriod** | Pointer to **int32** | The grace period in number of days permitted for the component blocklist enforcement rule type. | [optional] 
**LicenseRiskGracePeriod** | Pointer to **int32** | The grace period in number of days permitted for the component license risk rule type. | [optional] 
**SeverityGracePeriod** | Pointer to [**SeverityGracePeriods**](SeverityGracePeriods.md) |  | [optional] 
**CvssScoreGracePeriod** | Pointer to [**[]CvssScoreGracePeriod**](CvssScoreGracePeriod.md) | The grace period in number of days permitted for the vulnerability CVSS score rule type. | [optional] 

## Methods

### NewScaGracePeriods

`func NewScaGracePeriods() *ScaGracePeriods`

NewScaGracePeriods instantiates a new ScaGracePeriods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaGracePeriodsWithDefaults

`func NewScaGracePeriodsWithDefaults() *ScaGracePeriods`

NewScaGracePeriodsWithDefaults instantiates a new ScaGracePeriods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScaBlacklistGracePeriod

`func (o *ScaGracePeriods) GetScaBlacklistGracePeriod() int32`

GetScaBlacklistGracePeriod returns the ScaBlacklistGracePeriod field if non-nil, zero value otherwise.

### GetScaBlacklistGracePeriodOk

`func (o *ScaGracePeriods) GetScaBlacklistGracePeriodOk() (*int32, bool)`

GetScaBlacklistGracePeriodOk returns a tuple with the ScaBlacklistGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaBlacklistGracePeriod

`func (o *ScaGracePeriods) SetScaBlacklistGracePeriod(v int32)`

SetScaBlacklistGracePeriod sets ScaBlacklistGracePeriod field to given value.

### HasScaBlacklistGracePeriod

`func (o *ScaGracePeriods) HasScaBlacklistGracePeriod() bool`

HasScaBlacklistGracePeriod returns a boolean if a field has been set.

### GetLicenseRiskGracePeriod

`func (o *ScaGracePeriods) GetLicenseRiskGracePeriod() int32`

GetLicenseRiskGracePeriod returns the LicenseRiskGracePeriod field if non-nil, zero value otherwise.

### GetLicenseRiskGracePeriodOk

`func (o *ScaGracePeriods) GetLicenseRiskGracePeriodOk() (*int32, bool)`

GetLicenseRiskGracePeriodOk returns a tuple with the LicenseRiskGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRiskGracePeriod

`func (o *ScaGracePeriods) SetLicenseRiskGracePeriod(v int32)`

SetLicenseRiskGracePeriod sets LicenseRiskGracePeriod field to given value.

### HasLicenseRiskGracePeriod

`func (o *ScaGracePeriods) HasLicenseRiskGracePeriod() bool`

HasLicenseRiskGracePeriod returns a boolean if a field has been set.

### GetSeverityGracePeriod

`func (o *ScaGracePeriods) GetSeverityGracePeriod() SeverityGracePeriods`

GetSeverityGracePeriod returns the SeverityGracePeriod field if non-nil, zero value otherwise.

### GetSeverityGracePeriodOk

`func (o *ScaGracePeriods) GetSeverityGracePeriodOk() (*SeverityGracePeriods, bool)`

GetSeverityGracePeriodOk returns a tuple with the SeverityGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityGracePeriod

`func (o *ScaGracePeriods) SetSeverityGracePeriod(v SeverityGracePeriods)`

SetSeverityGracePeriod sets SeverityGracePeriod field to given value.

### HasSeverityGracePeriod

`func (o *ScaGracePeriods) HasSeverityGracePeriod() bool`

HasSeverityGracePeriod returns a boolean if a field has been set.

### GetCvssScoreGracePeriod

`func (o *ScaGracePeriods) GetCvssScoreGracePeriod() []CvssScoreGracePeriod`

GetCvssScoreGracePeriod returns the CvssScoreGracePeriod field if non-nil, zero value otherwise.

### GetCvssScoreGracePeriodOk

`func (o *ScaGracePeriods) GetCvssScoreGracePeriodOk() (*[]CvssScoreGracePeriod, bool)`

GetCvssScoreGracePeriodOk returns a tuple with the CvssScoreGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScoreGracePeriod

`func (o *ScaGracePeriods) SetCvssScoreGracePeriod(v []CvssScoreGracePeriod)`

SetCvssScoreGracePeriod sets CvssScoreGracePeriod field to given value.

### HasCvssScoreGracePeriod

`func (o *ScaGracePeriods) HasCvssScoreGracePeriod() bool`

HasCvssScoreGracePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CvssScoreGracePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upper** | Pointer to **float32** | The upper CVSS score limit for this grace period. Value must be between 0.0 and 10.0. | [optional] 
**Lower** | Pointer to **float32** | The lower CVSS score limit for this grace period. Value must be between 0.0 and 10.0. | [optional] 
**Days** | Pointer to **int32** | The grace period in number of days permitted for findings with a CVSS score within the range between the upper and lower CVSS score values. | [optional] 

## Methods

### NewCvssScoreGracePeriod

`func NewCvssScoreGracePeriod() *CvssScoreGracePeriod`

NewCvssScoreGracePeriod instantiates a new CvssScoreGracePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvssScoreGracePeriodWithDefaults

`func NewCvssScoreGracePeriodWithDefaults() *CvssScoreGracePeriod`

NewCvssScoreGracePeriodWithDefaults instantiates a new CvssScoreGracePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpper

`func (o *CvssScoreGracePeriod) GetUpper() float32`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *CvssScoreGracePeriod) GetUpperOk() (*float32, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *CvssScoreGracePeriod) SetUpper(v float32)`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *CvssScoreGracePeriod) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### GetLower

`func (o *CvssScoreGracePeriod) GetLower() float32`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *CvssScoreGracePeriod) GetLowerOk() (*float32, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *CvssScoreGracePeriod) SetLower(v float32)`

SetLower sets Lower field to given value.

### HasLower

`func (o *CvssScoreGracePeriod) HasLower() bool`

HasLower returns a boolean if a field has been set.

### GetDays

`func (o *CvssScoreGracePeriod) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *CvssScoreGracePeriod) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *CvssScoreGracePeriod) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *CvssScoreGracePeriod) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



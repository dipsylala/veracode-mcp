# ScaFindingCveCvss3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | The assigned CVSS v3 score for this vulnerability. | [optional] 
**Severity** | Pointer to **string** | The assigned CVSS v3 severity for this vulnerability. | [optional] 
**Vector** | Pointer to **string** | The assigned CVSS v3 vector for this vulnerability. | [optional] 

## Methods

### NewScaFindingCveCvss3

`func NewScaFindingCveCvss3() *ScaFindingCveCvss3`

NewScaFindingCveCvss3 instantiates a new ScaFindingCveCvss3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaFindingCveCvss3WithDefaults

`func NewScaFindingCveCvss3WithDefaults() *ScaFindingCveCvss3`

NewScaFindingCveCvss3WithDefaults instantiates a new ScaFindingCveCvss3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ScaFindingCveCvss3) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScaFindingCveCvss3) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScaFindingCveCvss3) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ScaFindingCveCvss3) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSeverity

`func (o *ScaFindingCveCvss3) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ScaFindingCveCvss3) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ScaFindingCveCvss3) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ScaFindingCveCvss3) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetVector

`func (o *ScaFindingCveCvss3) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *ScaFindingCveCvss3) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *ScaFindingCveCvss3) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *ScaFindingCveCvss3) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



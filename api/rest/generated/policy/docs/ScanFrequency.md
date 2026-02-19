# ScanFrequency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **string** | The frequency by which the organization is required to scan. | [optional] 
**PolicyVersion** | Pointer to [**PolicyVersion**](PolicyVersion.md) |  | [optional] 
**ScanType** | Pointer to **string** | The type of scan on which to enforce the rule. | [optional] 

## Methods

### NewScanFrequency

`func NewScanFrequency() *ScanFrequency`

NewScanFrequency instantiates a new ScanFrequency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanFrequencyWithDefaults

`func NewScanFrequencyWithDefaults() *ScanFrequency`

NewScanFrequencyWithDefaults instantiates a new ScanFrequency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *ScanFrequency) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ScanFrequency) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ScanFrequency) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ScanFrequency) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetPolicyVersion

`func (o *ScanFrequency) GetPolicyVersion() PolicyVersion`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *ScanFrequency) GetPolicyVersionOk() (*PolicyVersion, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *ScanFrequency) SetPolicyVersion(v PolicyVersion)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *ScanFrequency) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### GetScanType

`func (o *ScanFrequency) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ScanFrequency) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ScanFrequency) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *ScanFrequency) HasScanType() bool`

HasScanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



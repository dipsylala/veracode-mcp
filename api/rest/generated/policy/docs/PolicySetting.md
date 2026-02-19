# PolicySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessCriticality** | Pointer to **string** | The business criticality for which this policy is the default policy. | [optional] 
**Modified** | Pointer to **time.Time** | The date and time when the application default policy setting was modified. The date and time format is per RFC3339 and ISO-8601. Timezone is UTC. | [optional] 
**PolicyGuid** | Pointer to **string** | Unique identifier for the application policy. | [optional] 

## Methods

### NewPolicySetting

`func NewPolicySetting() *PolicySetting`

NewPolicySetting instantiates a new PolicySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySettingWithDefaults

`func NewPolicySettingWithDefaults() *PolicySetting`

NewPolicySettingWithDefaults instantiates a new PolicySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessCriticality

`func (o *PolicySetting) GetBusinessCriticality() string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *PolicySetting) GetBusinessCriticalityOk() (*string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *PolicySetting) SetBusinessCriticality(v string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *PolicySetting) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### GetModified

`func (o *PolicySetting) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PolicySetting) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PolicySetting) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *PolicySetting) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPolicyGuid

`func (o *PolicySetting) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *PolicySetting) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *PolicySetting) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *PolicySetting) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



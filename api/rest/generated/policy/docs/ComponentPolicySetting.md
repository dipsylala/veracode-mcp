# ComponentPolicySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modified** | Pointer to **time.Time** | The date and time when the pre-build component default policy setting was modified. The date and time format is per RFC3339 and ISO-8601. Timezone is UTC. | [optional] 
**ModifiedBy** | Pointer to **string** | Name of the user who most recently modified the pre-build component default policy setting. | [optional] 
**Plugin** | Pointer to **string** | Plugin | [optional] 
**PolicyGuid** | Pointer to **string** | Unique identifier for the pre-build component policy. | [optional] 

## Methods

### NewComponentPolicySetting

`func NewComponentPolicySetting() *ComponentPolicySetting`

NewComponentPolicySetting instantiates a new ComponentPolicySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPolicySettingWithDefaults

`func NewComponentPolicySettingWithDefaults() *ComponentPolicySetting`

NewComponentPolicySettingWithDefaults instantiates a new ComponentPolicySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModified

`func (o *ComponentPolicySetting) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ComponentPolicySetting) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ComponentPolicySetting) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ComponentPolicySetting) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetModifiedBy

`func (o *ComponentPolicySetting) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ComponentPolicySetting) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ComponentPolicySetting) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ComponentPolicySetting) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetPlugin

`func (o *ComponentPolicySetting) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *ComponentPolicySetting) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *ComponentPolicySetting) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *ComponentPolicySetting) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetPolicyGuid

`func (o *ComponentPolicySetting) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *ComponentPolicySetting) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *ComponentPolicySetting) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *ComponentPolicySetting) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



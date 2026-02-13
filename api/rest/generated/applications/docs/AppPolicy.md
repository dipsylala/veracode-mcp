# AppPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The policy name. | [optional] [readonly] 
**PolicyComplianceStatus** | Pointer to **string** | The policy compliance status. | [optional] [readonly] 

## Methods

### NewAppPolicy

`func NewAppPolicy() *AppPolicy`

NewAppPolicy instantiates a new AppPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPolicyWithDefaults

`func NewAppPolicyWithDefaults() *AppPolicy`

NewAppPolicyWithDefaults instantiates a new AppPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *AppPolicy) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AppPolicy) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AppPolicy) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AppPolicy) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetIsDefault

`func (o *AppPolicy) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AppPolicy) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AppPolicy) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AppPolicy) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *AppPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyComplianceStatus

`func (o *AppPolicy) GetPolicyComplianceStatus() string`

GetPolicyComplianceStatus returns the PolicyComplianceStatus field if non-nil, zero value otherwise.

### GetPolicyComplianceStatusOk

`func (o *AppPolicy) GetPolicyComplianceStatusOk() (*string, bool)`

GetPolicyComplianceStatusOk returns a tuple with the PolicyComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyComplianceStatus

`func (o *AppPolicy) SetPolicyComplianceStatus(v string)`

SetPolicyComplianceStatus sets PolicyComplianceStatus field to given value.

### HasPolicyComplianceStatus

`func (o *AppPolicy) HasPolicyComplianceStatus() bool`

HasPolicyComplianceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



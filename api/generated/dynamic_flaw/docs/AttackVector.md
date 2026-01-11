# AttackVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attack vector parameter. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of attack vector parameter. | [optional] [readonly] 
**OriginalValue** | Pointer to **string** | Original value of the attack vector. | [optional] [readonly] 
**InjectedValue** | Pointer to **string** | Injected value of the attack vector, after the scan engine has modified it to detect a flaw. | [optional] [readonly] 

## Methods

### NewAttackVector

`func NewAttackVector() *AttackVector`

NewAttackVector instantiates a new AttackVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackVectorWithDefaults

`func NewAttackVectorWithDefaults() *AttackVector`

NewAttackVectorWithDefaults instantiates a new AttackVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttackVector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttackVector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttackVector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttackVector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AttackVector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttackVector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttackVector) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttackVector) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOriginalValue

`func (o *AttackVector) GetOriginalValue() string`

GetOriginalValue returns the OriginalValue field if non-nil, zero value otherwise.

### GetOriginalValueOk

`func (o *AttackVector) GetOriginalValueOk() (*string, bool)`

GetOriginalValueOk returns a tuple with the OriginalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalValue

`func (o *AttackVector) SetOriginalValue(v string)`

SetOriginalValue sets OriginalValue field to given value.

### HasOriginalValue

`func (o *AttackVector) HasOriginalValue() bool`

HasOriginalValue returns a boolean if a field has been set.

### GetInjectedValue

`func (o *AttackVector) GetInjectedValue() string`

GetInjectedValue returns the InjectedValue field if non-nil, zero value otherwise.

### GetInjectedValueOk

`func (o *AttackVector) GetInjectedValueOk() (*string, bool)`

GetInjectedValueOk returns a tuple with the InjectedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectedValue

`func (o *AttackVector) SetInjectedValue(v string)`

SetInjectedValue sets InjectedValue field to given value.

### HasInjectedValue

`func (o *AttackVector) HasInjectedValue() bool`

HasInjectedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



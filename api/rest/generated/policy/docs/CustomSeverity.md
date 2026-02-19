# CustomSeverity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cwe** | Pointer to **int32** | The CWE associated with the custom severity. | [optional] 
**Severity** | Pointer to **int32** | The severity to be applied to findings of the specified CWE. | [optional] 

## Methods

### NewCustomSeverity

`func NewCustomSeverity() *CustomSeverity`

NewCustomSeverity instantiates a new CustomSeverity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomSeverityWithDefaults

`func NewCustomSeverityWithDefaults() *CustomSeverity`

NewCustomSeverityWithDefaults instantiates a new CustomSeverity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCwe

`func (o *CustomSeverity) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *CustomSeverity) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *CustomSeverity) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *CustomSeverity) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetSeverity

`func (o *CustomSeverity) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CustomSeverity) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CustomSeverity) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CustomSeverity) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



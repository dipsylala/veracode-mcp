# CweDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | CWE description. | [optional] 
**Id** | Pointer to **int32** | The unique CWE ID. | [optional] 
**Name** | Pointer to **string** | CWE name. | [optional] 
**Recommendation** | Pointer to **string** | CWE recommendation. | [optional] 
**References** | Pointer to [**[]CWEReference**](CWEReference.md) | CWE reference name and URL. | [optional] 
**RemediationEffort** | Pointer to **int32** | The level of effort it will take to fix this finding. Values: 1&#x3D;Trivial, 2&#x3D;Implementation error, 3&#x3D;Complex implementation error, 4&#x3D;Simple design error, 5&#x3D;Complex design error. | [optional] 
**Severity** | Pointer to **int32** | CWE severity. | [optional] 

## Methods

### NewCweDetail

`func NewCweDetail() *CweDetail`

NewCweDetail instantiates a new CweDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCweDetailWithDefaults

`func NewCweDetailWithDefaults() *CweDetail`

NewCweDetailWithDefaults instantiates a new CweDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CweDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CweDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CweDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CweDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CweDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CweDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CweDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CweDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CweDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CweDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CweDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CweDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecommendation

`func (o *CweDetail) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *CweDetail) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *CweDetail) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *CweDetail) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetReferences

`func (o *CweDetail) GetReferences() []CWEReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CweDetail) GetReferencesOk() (*[]CWEReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CweDetail) SetReferences(v []CWEReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CweDetail) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRemediationEffort

`func (o *CweDetail) GetRemediationEffort() int32`

GetRemediationEffort returns the RemediationEffort field if non-nil, zero value otherwise.

### GetRemediationEffortOk

`func (o *CweDetail) GetRemediationEffortOk() (*int32, bool)`

GetRemediationEffortOk returns a tuple with the RemediationEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationEffort

`func (o *CweDetail) SetRemediationEffort(v int32)`

SetRemediationEffort sets RemediationEffort field to given value.

### HasRemediationEffort

`func (o *CweDetail) HasRemediationEffort() bool`

HasRemediationEffort returns a boolean if a field has been set.

### GetSeverity

`func (o *CweDetail) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CweDetail) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CweDetail) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CweDetail) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



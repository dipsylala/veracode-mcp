# Annotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The mitigation action that was applied to the finding. | [optional] 
**Comment** | Pointer to **string** | All comments associated with the mitigation action. | [optional] 
**Created** | Pointer to **time.Time** | The date the annotation was added. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] 
**RemainingRisk** | Pointer to **string** | The value in the Remaining Risk field from the Comment column. | [optional] 
**Specifics** | Pointer to **string** | The value in the Specifics field from the Comment column. | [optional] 
**Technique** | Pointer to **string** | The value in the Technique field from the Comment column. | [optional] 
**UserName** | Pointer to **string** | The user who added the comment. | [optional] 
**Verification** | Pointer to **string** | The value of the Verification field in the Comment column. | [optional] 

## Methods

### NewAnnotation

`func NewAnnotation() *Annotation`

NewAnnotation instantiates a new Annotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationWithDefaults

`func NewAnnotationWithDefaults() *Annotation`

NewAnnotationWithDefaults instantiates a new Annotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Annotation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Annotation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Annotation) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Annotation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComment

`func (o *Annotation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Annotation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Annotation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Annotation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreated

`func (o *Annotation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Annotation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Annotation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Annotation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRemainingRisk

`func (o *Annotation) GetRemainingRisk() string`

GetRemainingRisk returns the RemainingRisk field if non-nil, zero value otherwise.

### GetRemainingRiskOk

`func (o *Annotation) GetRemainingRiskOk() (*string, bool)`

GetRemainingRiskOk returns a tuple with the RemainingRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingRisk

`func (o *Annotation) SetRemainingRisk(v string)`

SetRemainingRisk sets RemainingRisk field to given value.

### HasRemainingRisk

`func (o *Annotation) HasRemainingRisk() bool`

HasRemainingRisk returns a boolean if a field has been set.

### GetSpecifics

`func (o *Annotation) GetSpecifics() string`

GetSpecifics returns the Specifics field if non-nil, zero value otherwise.

### GetSpecificsOk

`func (o *Annotation) GetSpecificsOk() (*string, bool)`

GetSpecificsOk returns a tuple with the Specifics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifics

`func (o *Annotation) SetSpecifics(v string)`

SetSpecifics sets Specifics field to given value.

### HasSpecifics

`func (o *Annotation) HasSpecifics() bool`

HasSpecifics returns a boolean if a field has been set.

### GetTechnique

`func (o *Annotation) GetTechnique() string`

GetTechnique returns the Technique field if non-nil, zero value otherwise.

### GetTechniqueOk

`func (o *Annotation) GetTechniqueOk() (*string, bool)`

GetTechniqueOk returns a tuple with the Technique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnique

`func (o *Annotation) SetTechnique(v string)`

SetTechnique sets Technique field to given value.

### HasTechnique

`func (o *Annotation) HasTechnique() bool`

HasTechnique returns a boolean if a field has been set.

### GetUserName

`func (o *Annotation) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Annotation) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Annotation) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Annotation) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVerification

`func (o *Annotation) GetVerification() string`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *Annotation) GetVerificationOk() (*string, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *Annotation) SetVerification(v string)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *Annotation) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



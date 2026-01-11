# Appendix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeSamples** | Pointer to **[]string** | List of code samples associated with a manual finding. | [optional] 
**Description** | Pointer to **string** | Appendix description. | [optional] 
**Screenshots** | Pointer to [**[]Screenshot**](Screenshot.md) | List of screenshots associated with a manual finding. | [optional] 

## Methods

### NewAppendix

`func NewAppendix() *Appendix`

NewAppendix instantiates a new Appendix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendixWithDefaults

`func NewAppendixWithDefaults() *Appendix`

NewAppendixWithDefaults instantiates a new Appendix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeSamples

`func (o *Appendix) GetCodeSamples() []string`

GetCodeSamples returns the CodeSamples field if non-nil, zero value otherwise.

### GetCodeSamplesOk

`func (o *Appendix) GetCodeSamplesOk() (*[]string, bool)`

GetCodeSamplesOk returns a tuple with the CodeSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSamples

`func (o *Appendix) SetCodeSamples(v []string)`

SetCodeSamples sets CodeSamples field to given value.

### HasCodeSamples

`func (o *Appendix) HasCodeSamples() bool`

HasCodeSamples returns a boolean if a field has been set.

### GetDescription

`func (o *Appendix) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Appendix) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Appendix) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Appendix) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScreenshots

`func (o *Appendix) GetScreenshots() []Screenshot`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *Appendix) GetScreenshotsOk() (*[]Screenshot, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *Appendix) SetScreenshots(v []Screenshot)`

SetScreenshots sets Screenshots field to given value.

### HasScreenshots

`func (o *Appendix) HasScreenshots() bool`

HasScreenshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



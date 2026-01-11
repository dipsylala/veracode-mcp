# Screenshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** | The screenshot file data. | [optional] 
**Format** | Pointer to **string** | File format of the screenshot. | [optional] 
**Description** | Pointer to **string** | Description of the screenshot. | [optional] 

## Methods

### NewScreenshot

`func NewScreenshot() *Screenshot`

NewScreenshot instantiates a new Screenshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenshotWithDefaults

`func NewScreenshotWithDefaults() *Screenshot`

NewScreenshotWithDefaults instantiates a new Screenshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Screenshot) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Screenshot) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Screenshot) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *Screenshot) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFormat

`func (o *Screenshot) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Screenshot) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Screenshot) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Screenshot) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDescription

`func (o *Screenshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Screenshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Screenshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Screenshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



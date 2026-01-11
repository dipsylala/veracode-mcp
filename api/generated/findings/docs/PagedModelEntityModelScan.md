# PagedModelEntityModelScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**PagedModelEntityModelScanEmbedded**](PagedModelEntityModelScanEmbedded.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Page** | Pointer to [**PageMetadata**](PageMetadata.md) |  | [optional] 

## Methods

### NewPagedModelEntityModelScan

`func NewPagedModelEntityModelScan() *PagedModelEntityModelScan`

NewPagedModelEntityModelScan instantiates a new PagedModelEntityModelScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedModelEntityModelScanWithDefaults

`func NewPagedModelEntityModelScanWithDefaults() *PagedModelEntityModelScan`

NewPagedModelEntityModelScanWithDefaults instantiates a new PagedModelEntityModelScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *PagedModelEntityModelScan) GetEmbedded() PagedModelEntityModelScanEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PagedModelEntityModelScan) GetEmbeddedOk() (*PagedModelEntityModelScanEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PagedModelEntityModelScan) SetEmbedded(v PagedModelEntityModelScanEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PagedModelEntityModelScan) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *PagedModelEntityModelScan) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedModelEntityModelScan) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedModelEntityModelScan) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedModelEntityModelScan) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPage

`func (o *PagedModelEntityModelScan) GetPage() PageMetadata`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedModelEntityModelScan) GetPageOk() (*PageMetadata, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedModelEntityModelScan) SetPage(v PageMetadata)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedModelEntityModelScan) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



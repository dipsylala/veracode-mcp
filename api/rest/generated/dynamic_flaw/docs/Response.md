# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawBytes** | Pointer to **string** | Base64-encoded raw HTTP response payload. | [optional] [readonly] 

## Methods

### NewResponse

`func NewResponse() *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawBytes

`func (o *Response) GetRawBytes() string`

GetRawBytes returns the RawBytes field if non-nil, zero value otherwise.

### GetRawBytesOk

`func (o *Response) GetRawBytesOk() (*string, bool)`

GetRawBytesOk returns a tuple with the RawBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBytes

`func (o *Response) SetRawBytes(v string)`

SetRawBytes sets RawBytes field to given value.

### HasRawBytes

`func (o *Response) HasRawBytes() bool`

HasRawBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



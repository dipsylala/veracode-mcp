# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL associated with this request. | [optional] 
**RawBytes** | Pointer to **string** | Base64-encoded raw HTTP request payload. | [optional] [readonly] 
**Secure** | Pointer to **bool** | True if this is a secure HTTPS request, false if HTTP. | [optional] [readonly] 
**Port** | Pointer to **int32** | TCP port to which this request was made. | [optional] [readonly] 
**Protocol** | Pointer to **string** | Protocol associated with this request. Typically, HTTP. | [optional] [readonly] 
**Method** | Pointer to **string** | HTTP method of the request (GET, POST, PUT, PATCH, etc.). Can include custom methods. | [optional] [readonly] 
**Path** | Pointer to **string** | Path of the URL associated with this request. For example, no scheme, hostname, port, or parameter information. | [optional] [readonly] 
**Uri** | Pointer to **string** | Relative URI associated with this request, not including the host. | [optional] [readonly] 
**Body** | Pointer to **string** | Body of the request. May be empty. | [optional] [readonly] 
**Referer** | Pointer to **string** | Referer, from the HTTP header, associated with the request. | [optional] [readonly] 
**AttackVectors** | Pointer to [**[]AttackVector**](AttackVector.md) | Attack vector parameters associated with this request. | [optional] 

## Methods

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Request) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Request) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Request) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Request) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRawBytes

`func (o *Request) GetRawBytes() string`

GetRawBytes returns the RawBytes field if non-nil, zero value otherwise.

### GetRawBytesOk

`func (o *Request) GetRawBytesOk() (*string, bool)`

GetRawBytesOk returns a tuple with the RawBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBytes

`func (o *Request) SetRawBytes(v string)`

SetRawBytes sets RawBytes field to given value.

### HasRawBytes

`func (o *Request) HasRawBytes() bool`

HasRawBytes returns a boolean if a field has been set.

### GetSecure

`func (o *Request) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *Request) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *Request) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *Request) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetPort

`func (o *Request) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Request) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Request) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Request) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *Request) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Request) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Request) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Request) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetMethod

`func (o *Request) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Request) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Request) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Request) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *Request) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Request) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Request) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Request) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUri

`func (o *Request) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Request) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Request) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Request) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetBody

`func (o *Request) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Request) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Request) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Request) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetReferer

`func (o *Request) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *Request) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *Request) SetReferer(v string)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *Request) HasReferer() bool`

HasReferer returns a boolean if a field has been set.

### GetAttackVectors

`func (o *Request) GetAttackVectors() []AttackVector`

GetAttackVectors returns the AttackVectors field if non-nil, zero value otherwise.

### GetAttackVectorsOk

`func (o *Request) GetAttackVectorsOk() (*[]AttackVector, bool)`

GetAttackVectorsOk returns a tuple with the AttackVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVectors

`func (o *Request) SetAttackVectors(v []AttackVector)`

SetAttackVectors sets AttackVectors field to given value.

### HasAttackVectors

`func (o *Request) HasAttackVectors() bool`

HasAttackVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApplicationScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalStatus** | Pointer to **string** |  | [optional] 
**ModifiedDate** | Pointer to **time.Time** | The date when the scan results were published. The date/time format is per RFC3339 and ISO-8601, and the timezone is UTC. Example: 2019-04-12T23:20:50.52Z. | [optional] [readonly] 
**ScanType** | Pointer to **string** | Type of scan in which this finding was discovered. | [optional] [readonly] 
**ScanUrl** | Pointer to **string** | Unique path to the latest scan. | [optional] [readonly] 
**Status** | Pointer to **string** | Scan status | [optional] 

## Methods

### NewApplicationScan

`func NewApplicationScan() *ApplicationScan`

NewApplicationScan instantiates a new ApplicationScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationScanWithDefaults

`func NewApplicationScanWithDefaults() *ApplicationScan`

NewApplicationScanWithDefaults instantiates a new ApplicationScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalStatus

`func (o *ApplicationScan) GetInternalStatus() string`

GetInternalStatus returns the InternalStatus field if non-nil, zero value otherwise.

### GetInternalStatusOk

`func (o *ApplicationScan) GetInternalStatusOk() (*string, bool)`

GetInternalStatusOk returns a tuple with the InternalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalStatus

`func (o *ApplicationScan) SetInternalStatus(v string)`

SetInternalStatus sets InternalStatus field to given value.

### HasInternalStatus

`func (o *ApplicationScan) HasInternalStatus() bool`

HasInternalStatus returns a boolean if a field has been set.

### GetModifiedDate

`func (o *ApplicationScan) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ApplicationScan) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ApplicationScan) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ApplicationScan) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetScanType

`func (o *ApplicationScan) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ApplicationScan) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ApplicationScan) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *ApplicationScan) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetScanUrl

`func (o *ApplicationScan) GetScanUrl() string`

GetScanUrl returns the ScanUrl field if non-nil, zero value otherwise.

### GetScanUrlOk

`func (o *ApplicationScan) GetScanUrlOk() (*string, bool)`

GetScanUrlOk returns a tuple with the ScanUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanUrl

`func (o *ApplicationScan) SetScanUrl(v string)`

SetScanUrl sets ScanUrl field to given value.

### HasScanUrl

`func (o *ApplicationScan) HasScanUrl() bool`

HasScanUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationScan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationScan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationScan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationScan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FindingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to [**[]Coordinate**](Coordinate.md) |  | [optional] 
**PolicyVersion** | Pointer to [**PolicyVersion**](PolicyVersion.md) |  | [optional] 
**ScanType** | Pointer to **[]string** | The type of scan on which to enforce the rule. | [optional] 
**Type** | Pointer to **string** | Specify the supported rule types. | [optional] 
**AdvancedOptions** | Pointer to [**FindingRuleAdvancedOptions**](FindingRuleAdvancedOptions.md) |  | [optional] 
**Value** | Pointer to **string** | The value of this specific rule, such as the minimal score value. This value does not apply to the FAIL_ALL rule type. | [optional] 

## Methods

### NewFindingRule

`func NewFindingRule() *FindingRule`

NewFindingRule instantiates a new FindingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingRuleWithDefaults

`func NewFindingRuleWithDefaults() *FindingRule`

NewFindingRuleWithDefaults instantiates a new FindingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *FindingRule) GetCoordinates() []Coordinate`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *FindingRule) GetCoordinatesOk() (*[]Coordinate, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *FindingRule) SetCoordinates(v []Coordinate)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *FindingRule) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetPolicyVersion

`func (o *FindingRule) GetPolicyVersion() PolicyVersion`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *FindingRule) GetPolicyVersionOk() (*PolicyVersion, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *FindingRule) SetPolicyVersion(v PolicyVersion)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *FindingRule) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### GetScanType

`func (o *FindingRule) GetScanType() []string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *FindingRule) GetScanTypeOk() (*[]string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *FindingRule) SetScanType(v []string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *FindingRule) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetType

`func (o *FindingRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindingRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindingRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FindingRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvancedOptions

`func (o *FindingRule) GetAdvancedOptions() FindingRuleAdvancedOptions`

GetAdvancedOptions returns the AdvancedOptions field if non-nil, zero value otherwise.

### GetAdvancedOptionsOk

`func (o *FindingRule) GetAdvancedOptionsOk() (*FindingRuleAdvancedOptions, bool)`

GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedOptions

`func (o *FindingRule) SetAdvancedOptions(v FindingRuleAdvancedOptions)`

SetAdvancedOptions sets AdvancedOptions field to given value.

### HasAdvancedOptions

`func (o *FindingRule) HasAdvancedOptions() bool`

HasAdvancedOptions returns a boolean if a field has been set.

### GetValue

`func (o *FindingRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FindingRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FindingRule) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FindingRule) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



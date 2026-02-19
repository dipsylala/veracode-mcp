# FindingRuleAdvancedOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLicensesMustMeetRequirement** | Pointer to **bool** | Set to true to specify that component licenses must meet all policy rule requirements. | [optional] 
**AllowedNonossLicenses** | Pointer to **bool** | Set to true to allow licenses that are not open-source (OSS). | [optional] 
**FindingRule** | Pointer to [**FindingRule**](FindingRule.md) |  | [optional] 
**IsBlocklist** | Pointer to **bool** | Set to true to add the selected list of licenses to the blocklist. | [optional] 
**SelectedLicenses** | Pointer to [**[]ScaLicenseSummary**](ScaLicenseSummary.md) | List all selected licenses. | [optional] 

## Methods

### NewFindingRuleAdvancedOptions

`func NewFindingRuleAdvancedOptions() *FindingRuleAdvancedOptions`

NewFindingRuleAdvancedOptions instantiates a new FindingRuleAdvancedOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingRuleAdvancedOptionsWithDefaults

`func NewFindingRuleAdvancedOptionsWithDefaults() *FindingRuleAdvancedOptions`

NewFindingRuleAdvancedOptionsWithDefaults instantiates a new FindingRuleAdvancedOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLicensesMustMeetRequirement

`func (o *FindingRuleAdvancedOptions) GetAllLicensesMustMeetRequirement() bool`

GetAllLicensesMustMeetRequirement returns the AllLicensesMustMeetRequirement field if non-nil, zero value otherwise.

### GetAllLicensesMustMeetRequirementOk

`func (o *FindingRuleAdvancedOptions) GetAllLicensesMustMeetRequirementOk() (*bool, bool)`

GetAllLicensesMustMeetRequirementOk returns a tuple with the AllLicensesMustMeetRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLicensesMustMeetRequirement

`func (o *FindingRuleAdvancedOptions) SetAllLicensesMustMeetRequirement(v bool)`

SetAllLicensesMustMeetRequirement sets AllLicensesMustMeetRequirement field to given value.

### HasAllLicensesMustMeetRequirement

`func (o *FindingRuleAdvancedOptions) HasAllLicensesMustMeetRequirement() bool`

HasAllLicensesMustMeetRequirement returns a boolean if a field has been set.

### GetAllowedNonossLicenses

`func (o *FindingRuleAdvancedOptions) GetAllowedNonossLicenses() bool`

GetAllowedNonossLicenses returns the AllowedNonossLicenses field if non-nil, zero value otherwise.

### GetAllowedNonossLicensesOk

`func (o *FindingRuleAdvancedOptions) GetAllowedNonossLicensesOk() (*bool, bool)`

GetAllowedNonossLicensesOk returns a tuple with the AllowedNonossLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNonossLicenses

`func (o *FindingRuleAdvancedOptions) SetAllowedNonossLicenses(v bool)`

SetAllowedNonossLicenses sets AllowedNonossLicenses field to given value.

### HasAllowedNonossLicenses

`func (o *FindingRuleAdvancedOptions) HasAllowedNonossLicenses() bool`

HasAllowedNonossLicenses returns a boolean if a field has been set.

### GetFindingRule

`func (o *FindingRuleAdvancedOptions) GetFindingRule() FindingRule`

GetFindingRule returns the FindingRule field if non-nil, zero value otherwise.

### GetFindingRuleOk

`func (o *FindingRuleAdvancedOptions) GetFindingRuleOk() (*FindingRule, bool)`

GetFindingRuleOk returns a tuple with the FindingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingRule

`func (o *FindingRuleAdvancedOptions) SetFindingRule(v FindingRule)`

SetFindingRule sets FindingRule field to given value.

### HasFindingRule

`func (o *FindingRuleAdvancedOptions) HasFindingRule() bool`

HasFindingRule returns a boolean if a field has been set.

### GetIsBlocklist

`func (o *FindingRuleAdvancedOptions) GetIsBlocklist() bool`

GetIsBlocklist returns the IsBlocklist field if non-nil, zero value otherwise.

### GetIsBlocklistOk

`func (o *FindingRuleAdvancedOptions) GetIsBlocklistOk() (*bool, bool)`

GetIsBlocklistOk returns a tuple with the IsBlocklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocklist

`func (o *FindingRuleAdvancedOptions) SetIsBlocklist(v bool)`

SetIsBlocklist sets IsBlocklist field to given value.

### HasIsBlocklist

`func (o *FindingRuleAdvancedOptions) HasIsBlocklist() bool`

HasIsBlocklist returns a boolean if a field has been set.

### GetSelectedLicenses

`func (o *FindingRuleAdvancedOptions) GetSelectedLicenses() []ScaLicenseSummary`

GetSelectedLicenses returns the SelectedLicenses field if non-nil, zero value otherwise.

### GetSelectedLicensesOk

`func (o *FindingRuleAdvancedOptions) GetSelectedLicensesOk() (*[]ScaLicenseSummary, bool)`

GetSelectedLicensesOk returns a tuple with the SelectedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedLicenses

`func (o *FindingRuleAdvancedOptions) SetSelectedLicenses(v []ScaLicenseSummary)`

SetSelectedLicenses sets SelectedLicenses field to given value.

### HasSelectedLicenses

`func (o *FindingRuleAdvancedOptions) HasSelectedLicenses() bool`

HasSelectedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



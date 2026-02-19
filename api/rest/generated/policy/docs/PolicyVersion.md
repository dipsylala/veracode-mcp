# PolicyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the policy. | [optional] 
**Created** | Pointer to **time.Time** | The date and time the application was created. The date/time is formatted as per RFC3339 and ISO-8601. The timezone is UTC. | [optional] 
**CustomSeverities** | Pointer to [**[]CustomSeverity**](CustomSeverity.md) | A set of severity overrides for use with this policy. | [optional] 
**Description** | Pointer to **string** | A description of the policy. | [optional] 
**EvaluationDate** | Pointer to **time.Time** | The date from which to define the evaluation timeframe, which determines when ﬁndings that violate rules should cause an application to not pass policy. | [optional] 
**EvaluationDateType** | Pointer to **string** | Specify the supported evaluation date type as before or after the specified date. | [optional] 
**FindingRules** | Pointer to [**[]FindingRule**](FindingRule.md) | A set of rules to be evaluated against the scan findings. | [optional] 
**ScaGracePeriods** | Pointer to [**ScaGracePeriods**](ScaGracePeriods.md) |  | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the policy. | [optional] 
**ModifiedBy** | Pointer to **string** | The username of the user who most recently modified the policy. | [optional] 
**Name** | Pointer to **string** | Policy name. | [optional] 
**OrganizationId** | Pointer to **int32** | The organization with which the policy is associated. If no value is provided, the results are publicly visible. | [optional] 
**ScaBlacklistGracePeriod** | Pointer to **int32** | The grace period in number of days permitted for the component blocklist enforcement rule type. | [optional] 
**ScanFrequencyRules** | Pointer to [**[]ScanFrequency**](ScanFrequency.md) | The set of scan frequencies to be evaluated. | [optional] 
**ScoreGracePeriod** | Pointer to **int32** | The number of days grace period allowed for the policy score. | [optional] 
**Sev0GracePeriod** | Pointer to **int32** | The number of days grace period allowed for findings of severity 0. | [optional] 
**Sev1GracePeriod** | Pointer to **int32** | The number of days grace period allowed for findings of severity 1. | [optional] 
**Sev2GracePeriod** | Pointer to **int32** | The number of days grace period allowed for findings of severity 2. | [optional] 
**Sev3GracePeriod** | Pointer to **int32** | The number of days grace period allowed for findings of severity 3. | [optional] 
**Sev4GracePeriod** | Pointer to **int32** | The number of days grace period allowed for findings of severity 4. | [optional] 
**Sev5GracePeriod** | Pointer to **int32** | The number of days grace period allowed for findings of severity 5. | [optional] 
**Type** | Pointer to **string** | The evaluation policy type. | [optional] 
**VendorPolicy** | Pointer to **bool** | Use this flag to indicate if this policy is to be visible and available for policy evaluation by a vendor organization. | [optional] 
**Version** | Pointer to **int32** | The version of this policy. | [optional] 

## Methods

### NewPolicyVersion

`func NewPolicyVersion() *PolicyVersion`

NewPolicyVersion instantiates a new PolicyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyVersionWithDefaults

`func NewPolicyVersionWithDefaults() *PolicyVersion`

NewPolicyVersionWithDefaults instantiates a new PolicyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PolicyVersion) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PolicyVersion) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PolicyVersion) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PolicyVersion) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreated

`func (o *PolicyVersion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PolicyVersion) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PolicyVersion) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PolicyVersion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomSeverities

`func (o *PolicyVersion) GetCustomSeverities() []CustomSeverity`

GetCustomSeverities returns the CustomSeverities field if non-nil, zero value otherwise.

### GetCustomSeveritiesOk

`func (o *PolicyVersion) GetCustomSeveritiesOk() (*[]CustomSeverity, bool)`

GetCustomSeveritiesOk returns a tuple with the CustomSeverities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSeverities

`func (o *PolicyVersion) SetCustomSeverities(v []CustomSeverity)`

SetCustomSeverities sets CustomSeverities field to given value.

### HasCustomSeverities

`func (o *PolicyVersion) HasCustomSeverities() bool`

HasCustomSeverities returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluationDate

`func (o *PolicyVersion) GetEvaluationDate() time.Time`

GetEvaluationDate returns the EvaluationDate field if non-nil, zero value otherwise.

### GetEvaluationDateOk

`func (o *PolicyVersion) GetEvaluationDateOk() (*time.Time, bool)`

GetEvaluationDateOk returns a tuple with the EvaluationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDate

`func (o *PolicyVersion) SetEvaluationDate(v time.Time)`

SetEvaluationDate sets EvaluationDate field to given value.

### HasEvaluationDate

`func (o *PolicyVersion) HasEvaluationDate() bool`

HasEvaluationDate returns a boolean if a field has been set.

### GetEvaluationDateType

`func (o *PolicyVersion) GetEvaluationDateType() string`

GetEvaluationDateType returns the EvaluationDateType field if non-nil, zero value otherwise.

### GetEvaluationDateTypeOk

`func (o *PolicyVersion) GetEvaluationDateTypeOk() (*string, bool)`

GetEvaluationDateTypeOk returns a tuple with the EvaluationDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDateType

`func (o *PolicyVersion) SetEvaluationDateType(v string)`

SetEvaluationDateType sets EvaluationDateType field to given value.

### HasEvaluationDateType

`func (o *PolicyVersion) HasEvaluationDateType() bool`

HasEvaluationDateType returns a boolean if a field has been set.

### GetFindingRules

`func (o *PolicyVersion) GetFindingRules() []FindingRule`

GetFindingRules returns the FindingRules field if non-nil, zero value otherwise.

### GetFindingRulesOk

`func (o *PolicyVersion) GetFindingRulesOk() (*[]FindingRule, bool)`

GetFindingRulesOk returns a tuple with the FindingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingRules

`func (o *PolicyVersion) SetFindingRules(v []FindingRule)`

SetFindingRules sets FindingRules field to given value.

### HasFindingRules

`func (o *PolicyVersion) HasFindingRules() bool`

HasFindingRules returns a boolean if a field has been set.

### GetScaGracePeriods

`func (o *PolicyVersion) GetScaGracePeriods() ScaGracePeriods`

GetScaGracePeriods returns the ScaGracePeriods field if non-nil, zero value otherwise.

### GetScaGracePeriodsOk

`func (o *PolicyVersion) GetScaGracePeriodsOk() (*ScaGracePeriods, bool)`

GetScaGracePeriodsOk returns a tuple with the ScaGracePeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaGracePeriods

`func (o *PolicyVersion) SetScaGracePeriods(v ScaGracePeriods)`

SetScaGracePeriods sets ScaGracePeriods field to given value.

### HasScaGracePeriods

`func (o *PolicyVersion) HasScaGracePeriods() bool`

HasScaGracePeriods returns a boolean if a field has been set.

### GetGuid

`func (o *PolicyVersion) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *PolicyVersion) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *PolicyVersion) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *PolicyVersion) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetModifiedBy

`func (o *PolicyVersion) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PolicyVersion) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PolicyVersion) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *PolicyVersion) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *PolicyVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PolicyVersion) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PolicyVersion) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PolicyVersion) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PolicyVersion) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetScaBlacklistGracePeriod

`func (o *PolicyVersion) GetScaBlacklistGracePeriod() int32`

GetScaBlacklistGracePeriod returns the ScaBlacklistGracePeriod field if non-nil, zero value otherwise.

### GetScaBlacklistGracePeriodOk

`func (o *PolicyVersion) GetScaBlacklistGracePeriodOk() (*int32, bool)`

GetScaBlacklistGracePeriodOk returns a tuple with the ScaBlacklistGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaBlacklistGracePeriod

`func (o *PolicyVersion) SetScaBlacklistGracePeriod(v int32)`

SetScaBlacklistGracePeriod sets ScaBlacklistGracePeriod field to given value.

### HasScaBlacklistGracePeriod

`func (o *PolicyVersion) HasScaBlacklistGracePeriod() bool`

HasScaBlacklistGracePeriod returns a boolean if a field has been set.

### GetScanFrequencyRules

`func (o *PolicyVersion) GetScanFrequencyRules() []ScanFrequency`

GetScanFrequencyRules returns the ScanFrequencyRules field if non-nil, zero value otherwise.

### GetScanFrequencyRulesOk

`func (o *PolicyVersion) GetScanFrequencyRulesOk() (*[]ScanFrequency, bool)`

GetScanFrequencyRulesOk returns a tuple with the ScanFrequencyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFrequencyRules

`func (o *PolicyVersion) SetScanFrequencyRules(v []ScanFrequency)`

SetScanFrequencyRules sets ScanFrequencyRules field to given value.

### HasScanFrequencyRules

`func (o *PolicyVersion) HasScanFrequencyRules() bool`

HasScanFrequencyRules returns a boolean if a field has been set.

### GetScoreGracePeriod

`func (o *PolicyVersion) GetScoreGracePeriod() int32`

GetScoreGracePeriod returns the ScoreGracePeriod field if non-nil, zero value otherwise.

### GetScoreGracePeriodOk

`func (o *PolicyVersion) GetScoreGracePeriodOk() (*int32, bool)`

GetScoreGracePeriodOk returns a tuple with the ScoreGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreGracePeriod

`func (o *PolicyVersion) SetScoreGracePeriod(v int32)`

SetScoreGracePeriod sets ScoreGracePeriod field to given value.

### HasScoreGracePeriod

`func (o *PolicyVersion) HasScoreGracePeriod() bool`

HasScoreGracePeriod returns a boolean if a field has been set.

### GetSev0GracePeriod

`func (o *PolicyVersion) GetSev0GracePeriod() int32`

GetSev0GracePeriod returns the Sev0GracePeriod field if non-nil, zero value otherwise.

### GetSev0GracePeriodOk

`func (o *PolicyVersion) GetSev0GracePeriodOk() (*int32, bool)`

GetSev0GracePeriodOk returns a tuple with the Sev0GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev0GracePeriod

`func (o *PolicyVersion) SetSev0GracePeriod(v int32)`

SetSev0GracePeriod sets Sev0GracePeriod field to given value.

### HasSev0GracePeriod

`func (o *PolicyVersion) HasSev0GracePeriod() bool`

HasSev0GracePeriod returns a boolean if a field has been set.

### GetSev1GracePeriod

`func (o *PolicyVersion) GetSev1GracePeriod() int32`

GetSev1GracePeriod returns the Sev1GracePeriod field if non-nil, zero value otherwise.

### GetSev1GracePeriodOk

`func (o *PolicyVersion) GetSev1GracePeriodOk() (*int32, bool)`

GetSev1GracePeriodOk returns a tuple with the Sev1GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev1GracePeriod

`func (o *PolicyVersion) SetSev1GracePeriod(v int32)`

SetSev1GracePeriod sets Sev1GracePeriod field to given value.

### HasSev1GracePeriod

`func (o *PolicyVersion) HasSev1GracePeriod() bool`

HasSev1GracePeriod returns a boolean if a field has been set.

### GetSev2GracePeriod

`func (o *PolicyVersion) GetSev2GracePeriod() int32`

GetSev2GracePeriod returns the Sev2GracePeriod field if non-nil, zero value otherwise.

### GetSev2GracePeriodOk

`func (o *PolicyVersion) GetSev2GracePeriodOk() (*int32, bool)`

GetSev2GracePeriodOk returns a tuple with the Sev2GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev2GracePeriod

`func (o *PolicyVersion) SetSev2GracePeriod(v int32)`

SetSev2GracePeriod sets Sev2GracePeriod field to given value.

### HasSev2GracePeriod

`func (o *PolicyVersion) HasSev2GracePeriod() bool`

HasSev2GracePeriod returns a boolean if a field has been set.

### GetSev3GracePeriod

`func (o *PolicyVersion) GetSev3GracePeriod() int32`

GetSev3GracePeriod returns the Sev3GracePeriod field if non-nil, zero value otherwise.

### GetSev3GracePeriodOk

`func (o *PolicyVersion) GetSev3GracePeriodOk() (*int32, bool)`

GetSev3GracePeriodOk returns a tuple with the Sev3GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev3GracePeriod

`func (o *PolicyVersion) SetSev3GracePeriod(v int32)`

SetSev3GracePeriod sets Sev3GracePeriod field to given value.

### HasSev3GracePeriod

`func (o *PolicyVersion) HasSev3GracePeriod() bool`

HasSev3GracePeriod returns a boolean if a field has been set.

### GetSev4GracePeriod

`func (o *PolicyVersion) GetSev4GracePeriod() int32`

GetSev4GracePeriod returns the Sev4GracePeriod field if non-nil, zero value otherwise.

### GetSev4GracePeriodOk

`func (o *PolicyVersion) GetSev4GracePeriodOk() (*int32, bool)`

GetSev4GracePeriodOk returns a tuple with the Sev4GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev4GracePeriod

`func (o *PolicyVersion) SetSev4GracePeriod(v int32)`

SetSev4GracePeriod sets Sev4GracePeriod field to given value.

### HasSev4GracePeriod

`func (o *PolicyVersion) HasSev4GracePeriod() bool`

HasSev4GracePeriod returns a boolean if a field has been set.

### GetSev5GracePeriod

`func (o *PolicyVersion) GetSev5GracePeriod() int32`

GetSev5GracePeriod returns the Sev5GracePeriod field if non-nil, zero value otherwise.

### GetSev5GracePeriodOk

`func (o *PolicyVersion) GetSev5GracePeriodOk() (*int32, bool)`

GetSev5GracePeriodOk returns a tuple with the Sev5GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev5GracePeriod

`func (o *PolicyVersion) SetSev5GracePeriod(v int32)`

SetSev5GracePeriod sets Sev5GracePeriod field to given value.

### HasSev5GracePeriod

`func (o *PolicyVersion) HasSev5GracePeriod() bool`

HasSev5GracePeriod returns a boolean if a field has been set.

### GetType

`func (o *PolicyVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyVersion) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyVersion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorPolicy

`func (o *PolicyVersion) GetVendorPolicy() bool`

GetVendorPolicy returns the VendorPolicy field if non-nil, zero value otherwise.

### GetVendorPolicyOk

`func (o *PolicyVersion) GetVendorPolicyOk() (*bool, bool)`

GetVendorPolicyOk returns a tuple with the VendorPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPolicy

`func (o *PolicyVersion) SetVendorPolicy(v bool)`

SetVendorPolicy sets VendorPolicy field to given value.

### HasVendorPolicy

`func (o *PolicyVersion) HasVendorPolicy() bool`

HasVendorPolicy returns a boolean if a field has been set.

### GetVersion

`func (o *PolicyVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PolicyVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PolicyVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PolicyVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



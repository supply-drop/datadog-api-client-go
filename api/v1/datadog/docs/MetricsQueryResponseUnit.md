# MetricsQueryResponseUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | Unit family, allows for conversion between units of the same family, for scaling. | [optional] [readonly] 
**Name** | Pointer to **string** | Unit name | [optional] [readonly] 
**Plural** | Pointer to **string** | Plural form of the unit name. | [optional] [readonly] 
**ScaleFactor** | Pointer to **float64** | Factor for scaling between units of the same family. | [optional] [readonly] 
**ShortName** | Pointer to **string** | Abbreviation of the unit. | [optional] [readonly] 

## Methods

### NewMetricsQueryResponseUnit

`func NewMetricsQueryResponseUnit() *MetricsQueryResponseUnit`

NewMetricsQueryResponseUnit instantiates a new MetricsQueryResponseUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryResponseUnitWithDefaults

`func NewMetricsQueryResponseUnitWithDefaults() *MetricsQueryResponseUnit`

NewMetricsQueryResponseUnitWithDefaults instantiates a new MetricsQueryResponseUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *MetricsQueryResponseUnit) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *MetricsQueryResponseUnit) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *MetricsQueryResponseUnit) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *MetricsQueryResponseUnit) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetName

`func (o *MetricsQueryResponseUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsQueryResponseUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsQueryResponseUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsQueryResponseUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlural

`func (o *MetricsQueryResponseUnit) GetPlural() string`

GetPlural returns the Plural field if non-nil, zero value otherwise.

### GetPluralOk

`func (o *MetricsQueryResponseUnit) GetPluralOk() (*string, bool)`

GetPluralOk returns a tuple with the Plural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlural

`func (o *MetricsQueryResponseUnit) SetPlural(v string)`

SetPlural sets Plural field to given value.

### HasPlural

`func (o *MetricsQueryResponseUnit) HasPlural() bool`

HasPlural returns a boolean if a field has been set.

### GetScaleFactor

`func (o *MetricsQueryResponseUnit) GetScaleFactor() float64`

GetScaleFactor returns the ScaleFactor field if non-nil, zero value otherwise.

### GetScaleFactorOk

`func (o *MetricsQueryResponseUnit) GetScaleFactorOk() (*float64, bool)`

GetScaleFactorOk returns a tuple with the ScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleFactor

`func (o *MetricsQueryResponseUnit) SetScaleFactor(v float64)`

SetScaleFactor sets ScaleFactor field to given value.

### HasScaleFactor

`func (o *MetricsQueryResponseUnit) HasScaleFactor() bool`

HasScaleFactor returns a boolean if a field has been set.

### GetShortName

`func (o *MetricsQueryResponseUnit) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *MetricsQueryResponseUnit) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *MetricsQueryResponseUnit) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *MetricsQueryResponseUnit) HasShortName() bool`

HasShortName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



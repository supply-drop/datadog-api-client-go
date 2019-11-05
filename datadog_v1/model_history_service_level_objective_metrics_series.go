/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
)

// HistoryServiceLevelObjectiveMetricsSeries A representation of `metric` based SLO time series for the provided queries. This is the same response type from `batch_query` endpoint.
type HistoryServiceLevelObjectiveMetricsSeries struct {
	// Count of submitted metrics
	Count *int64 `json:"count,omitempty"`

	Metadata *HistoryServiceLevelObjectiveMetricsSeriesMetadata `json:"metadata,omitempty"`

	// Total Sum of the query
	Sum *float64 `json:"sum,omitempty"`

	// The query timestamps in epoch seconds
	Times *[]int64 `json:"times,omitempty"`

	// The query values
	Values *[]float32 `json:"values,omitempty"`
}

// GetCount returns the Count field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetCountOk() (int64, bool) {
	if o == nil || o.Count == nil {
		var ret int64
		return ret, false
	}
	return *o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetCount(v int64) {
	o.Count = &v
}

// GetMetadata returns the Metadata field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadata() HistoryServiceLevelObjectiveMetricsSeriesMetadata {
	if o == nil || o.Metadata == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeriesMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadataOk() (HistoryServiceLevelObjectiveMetricsSeriesMetadata, bool) {
	if o == nil || o.Metadata == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeriesMetadata
		return ret, false
	}
	return *o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given HistoryServiceLevelObjectiveMetricsSeriesMetadata and assigns it to the Metadata field.
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetMetadata(v HistoryServiceLevelObjectiveMetricsSeriesMetadata) {
	o.Metadata = &v
}

// GetSum returns the Sum field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSum() float64 {
	if o == nil || o.Sum == nil {
		var ret float64
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSumOk() (float64, bool) {
	if o == nil || o.Sum == nil {
		var ret float64
		return ret, false
	}
	return *o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given float64 and assigns it to the Sum field.
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetSum(v float64) {
	o.Sum = &v
}

// GetTimes returns the Times field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetTimes() []int64 {
	if o == nil || o.Times == nil {
		var ret []int64
		return ret
	}
	return *o.Times
}

// GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetTimesOk() ([]int64, bool) {
	if o == nil || o.Times == nil {
		var ret []int64
		return ret, false
	}
	return *o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) HasTimes() bool {
	if o != nil && o.Times != nil {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []int64 and assigns it to the Times field.
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetTimes(v []int64) {
	o.Times = &v
}

// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValues() []float32 {
	if o == nil || o.Values == nil {
		var ret []float32
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValuesOk() ([]float32, bool) {
	if o == nil || o.Values == nil {
		var ret []float32
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetricsSeries) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []float32 and assigns it to the Values field.
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetValues(v []float32) {
	o.Values = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o HistoryServiceLevelObjectiveMetricsSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Sum != nil {
		toSerialize["sum"] = o.Sum
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}
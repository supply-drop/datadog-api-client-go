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

// DowntimeRecurrence struct for DowntimeRecurrence
type DowntimeRecurrence struct {
	Period *int32 `json:"period,omitempty"`

	Type *string `json:"type,omitempty"`

	UntilDate                      *int64    `json:"until_date,omitempty"`
	isExplicitNullUntilDate        bool      `json:"-"`
	UntilOccurrences               *int32    `json:"until_occurrences,omitempty"`
	isExplicitNullUntilOccurrences bool      `json:"-"`
	WeekDays                       *[]string `json:"week_days,omitempty"`
}

// GetPeriod returns the Period field if non-nil, zero value otherwise.
func (o *DowntimeRecurrence) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRecurrence) GetPeriodOk() (int32, bool) {
	if o == nil || o.Period == nil {
		var ret int32
		return ret, false
	}
	return *o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DowntimeRecurrence) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *DowntimeRecurrence) SetPeriod(v int32) {
	o.Period = &v
}

// GetType returns the Type field if non-nil, zero value otherwise.
func (o *DowntimeRecurrence) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRecurrence) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DowntimeRecurrence) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DowntimeRecurrence) SetType(v string) {
	o.Type = &v
}

// GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.
func (o *DowntimeRecurrence) GetUntilDate() int64 {
	if o == nil || o.UntilDate == nil {
		var ret int64
		return ret
	}
	return *o.UntilDate
}

// GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRecurrence) GetUntilDateOk() (int64, bool) {
	if o == nil || o.UntilDate == nil {
		var ret int64
		return ret, false
	}
	return *o.UntilDate, true
}

// HasUntilDate returns a boolean if a field has been set.
func (o *DowntimeRecurrence) HasUntilDate() bool {
	if o != nil && o.UntilDate != nil {
		return true
	}

	return false
}

// SetUntilDate gets a reference to the given int64 and assigns it to the UntilDate field.
func (o *DowntimeRecurrence) SetUntilDate(v int64) {
	o.UntilDate = &v
}

// SetUntilDateExplicitNull (un)sets UntilDate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UntilDate value is set to nil even if false is passed
func (o *DowntimeRecurrence) SetUntilDateExplicitNull(b bool) {
	o.UntilDate = nil
	o.isExplicitNullUntilDate = b
}

// GetUntilOccurrences returns the UntilOccurrences field if non-nil, zero value otherwise.
func (o *DowntimeRecurrence) GetUntilOccurrences() int32 {
	if o == nil || o.UntilOccurrences == nil {
		var ret int32
		return ret
	}
	return *o.UntilOccurrences
}

// GetUntilOccurrencesOk returns a tuple with the UntilOccurrences field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRecurrence) GetUntilOccurrencesOk() (int32, bool) {
	if o == nil || o.UntilOccurrences == nil {
		var ret int32
		return ret, false
	}
	return *o.UntilOccurrences, true
}

// HasUntilOccurrences returns a boolean if a field has been set.
func (o *DowntimeRecurrence) HasUntilOccurrences() bool {
	if o != nil && o.UntilOccurrences != nil {
		return true
	}

	return false
}

// SetUntilOccurrences gets a reference to the given int32 and assigns it to the UntilOccurrences field.
func (o *DowntimeRecurrence) SetUntilOccurrences(v int32) {
	o.UntilOccurrences = &v
}

// SetUntilOccurrencesExplicitNull (un)sets UntilOccurrences to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UntilOccurrences value is set to nil even if false is passed
func (o *DowntimeRecurrence) SetUntilOccurrencesExplicitNull(b bool) {
	o.UntilOccurrences = nil
	o.isExplicitNullUntilOccurrences = b
}

// GetWeekDays returns the WeekDays field if non-nil, zero value otherwise.
func (o *DowntimeRecurrence) GetWeekDays() []string {
	if o == nil || o.WeekDays == nil {
		var ret []string
		return ret
	}
	return *o.WeekDays
}

// GetWeekDaysOk returns a tuple with the WeekDays field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRecurrence) GetWeekDaysOk() ([]string, bool) {
	if o == nil || o.WeekDays == nil {
		var ret []string
		return ret, false
	}
	return *o.WeekDays, true
}

// HasWeekDays returns a boolean if a field has been set.
func (o *DowntimeRecurrence) HasWeekDays() bool {
	if o != nil && o.WeekDays != nil {
		return true
	}

	return false
}

// SetWeekDays gets a reference to the given []string and assigns it to the WeekDays field.
func (o *DowntimeRecurrence) SetWeekDays(v []string) {
	o.WeekDays = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o DowntimeRecurrence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UntilDate == nil {
		if o.isExplicitNullUntilDate {
			toSerialize["until_date"] = o.UntilDate
		}
	} else {
		toSerialize["until_date"] = o.UntilDate
	}
	if o.UntilOccurrences == nil {
		if o.isExplicitNullUntilOccurrences {
			toSerialize["until_occurrences"] = o.UntilOccurrences
		}
	} else {
		toSerialize["until_occurrences"] = o.UntilOccurrences
	}
	if o.WeekDays != nil {
		toSerialize["week_days"] = o.WeekDays
	}
	return json.Marshal(toSerialize)
}

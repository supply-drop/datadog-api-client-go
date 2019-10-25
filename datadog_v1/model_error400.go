/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
	"errors"
)

// Error400 struct for Error400
type Error400 struct {
	Errors *[]string `json:"errors,omitempty"`
}

// GetErrors returns the Errors field if non-nil, zero value otherwise.
func (o *Error400) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Error400) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Error400) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *Error400) SetErrors(v []string) {
	o.Errors = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Error400) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors == nil {
		return nil, errors.New("Errors is required and not nullable, but was not set on Error400")
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

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

// AwsAccountCreateResponse struct for AwsAccountCreateResponse
type AwsAccountCreateResponse struct {
	ExternalId *string `json:"external_id,omitempty"`
}

// GetExternalId returns the ExternalId field if non-nil, zero value otherwise.
func (o *AwsAccountCreateResponse) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsAccountCreateResponse) GetExternalIdOk() (string, bool) {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AwsAccountCreateResponse) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AwsAccountCreateResponse) SetExternalId(v string) {
	o.ExternalId = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o AwsAccountCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	return json.Marshal(toSerialize)
}
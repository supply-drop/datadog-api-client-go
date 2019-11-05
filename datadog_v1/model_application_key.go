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

// ApplicationKey struct for ApplicationKey
type ApplicationKey struct {
	Hash *string `json:"hash,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`
}

// GetHash returns the Hash field if non-nil, zero value otherwise.
func (o *ApplicationKey) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKey) GetHashOk() (string, bool) {
	if o == nil || o.Hash == nil {
		var ret string
		return ret, false
	}
	return *o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApplicationKey) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApplicationKey) SetHash(v string) {
	o.Hash = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *ApplicationKey) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKey) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationKey) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationKey) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field if non-nil, zero value otherwise.
func (o *ApplicationKey) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKey) GetOwnerOk() (string, bool) {
	if o == nil || o.Owner == nil {
		var ret string
		return ret, false
	}
	return *o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ApplicationKey) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ApplicationKey) SetOwner(v string) {
	o.Owner = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o ApplicationKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}
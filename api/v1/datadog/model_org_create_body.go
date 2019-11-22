/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// OrgCreateBody struct for OrgCreateBody
type OrgCreateBody struct {
	Billing      OrgBilling      `json:"billing"`
	Name         string          `json:"name"`
	Subscription OrgSubscription `json:"subscription"`
}

// GetBilling returns the Billing field value
func (o *OrgCreateBody) GetBilling() OrgBilling {
	if o == nil {
		var ret OrgBilling
		return ret
	}

	return o.Billing
}

// SetBilling sets field value
func (o *OrgCreateBody) SetBilling(v OrgBilling) {
	o.Billing = v
}

// GetName returns the Name field value
func (o *OrgCreateBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *OrgCreateBody) SetName(v string) {
	o.Name = v
}

// GetSubscription returns the Subscription field value
func (o *OrgCreateBody) GetSubscription() OrgSubscription {
	if o == nil {
		var ret OrgSubscription
		return ret
	}

	return o.Subscription
}

// SetSubscription sets field value
func (o *OrgCreateBody) SetSubscription(v OrgSubscription) {
	o.Subscription = v
}

type NullableOrgCreateBody struct {
	Value        OrgCreateBody
	ExplicitNull bool
}

func (v NullableOrgCreateBody) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOrgCreateBody) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
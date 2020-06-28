/*
 * Datadog API Collection
 *
 * Collection of all Datadog Public endpoints.
 *
 * API version: 1.0
 * Contact: support@datadoghq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog
// SyntheticsTestsTrigger Object describing the the triggers for a Synthetic test.
type SyntheticsTestsTrigger struct {
	// The test public ID.
	PublicId string `json:"public_id,omitempty"`
	// Allows loading insecure content for an HTTP request.
	AllowInsecureCertificates bool `json:"allowInsecureCertificates,omitempty"`
	BasicAuth SyntheticsTestRequestBasicAuth `json:"basicAuth,omitempty"`
	// Body to include in the test.
	Body string `json:"body,omitempty"`
	// Body to include in the test.
	BodyType string `json:"bodyType,omitempty"`
	// Array with the different device IDs used to run the test.
	DeviceIds []SyntheticsDeviceId `json:"deviceIds,omitempty"`
	// For API SSL test, whether or not the test should follow redirects.
	FollowRedirects bool `json:"followRedirects,omitempty"`
	// Headers to include when performing the test.
	Headers map[string]string `json:"headers,omitempty"`
	// Array of locations used to run the test.
	Locations []string `json:"locations,omitempty"`
	Retry SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	// Starting URL for the browser test.
	StartUrl string `json:"startUrl,omitempty"`
	// Array of variables used for the test.
	Variables []SyntheticsBrowserVariable `json:"variables,omitempty"`
}
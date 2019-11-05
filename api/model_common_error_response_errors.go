/*
 * App Store Connect API
 *
 * The App Store Connect API is a REST API used to build custom workflows as part of your app development life cycle and automate actions you take in App Store Connect. Use of the API requires authorization via JSON Web Tokens (JWT); you obtain keys to create the tokens from your organization’s App Store Connect account. See Creating API Keys for App Store Connect API to create your keys and tokens.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreconnect
// CommonErrorResponseErrors The details about one error that is returned when an API request is not successful.
type CommonErrorResponseErrors struct {
	// (Required) A machine-readable code indicating the type of error. The code is a hierarchical value with levels of specificity separated by the '.' character. This value is parseable for programmatic error handling in code.
	Code string `json:"code"`
	// (Required) The HTTP status code of the error. This status code usually matches the response's status code; however, if the request produces multiple errors, these two codes may differ.
	Status string `json:"status"`
	// The unique ID of a specific instance of an error, request, and response. Use this ID when providing feedback to or debugging issues with Apple.
	Id string `json:"id,omitempty"`
	// (Required) A summary of the error. Do not use this field for programmatic error handling.
	Title string `json:"title"`
	// (Required) A detailed explanation of the error. Do not use this field for programmatic error handling.
	Detail string `json:"detail"`
	// One of two possible types of values. source.parameter, provided when a query parameter produced the error, or source.JsonPointer, provided when a problem with the entity produced the error.
	Source OneOfCommonErrorResponseErrorsJsonPointerCommonErrorResponseErrorsParameter `json:"source,omitempty"`
}
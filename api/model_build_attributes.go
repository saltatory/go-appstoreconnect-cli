/*
 * App Store Connect API
 *
 * The App Store Connect API is a REST API used to build custom workflows as part of your app development life cycle and automate actions you take in App Store Connect. Use of the API requires authorization via JSON Web Tokens (JWT); you obtain keys to create the tokens from your organization’s App Store Connect account. See Creating API Keys for App Store Connect API to create your keys and tokens.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreconnect
// BuildAttributes Attributes that describe a resource.
type BuildAttributes struct {
	// A Boolean value that indicates if the build has expired. An expired build is unavailable for testing.
	Expired bool `json:"expired,omitempty"`
	IconAssetToken ImageAsset `json:"iconAssetToken,omitempty"`
	// The minimum operating system version needed to test a build.
	MinOsVersion string `json:"minOsVersion,omitempty"`
	// The processing state of the build indicating that it is not yet available for testing.
	ProcessingState string `json:"processingState,omitempty"`
	// The version number of the uploaded build.
	Version string `json:"version,omitempty"`
	// A Boolean value that indicates whether the build uses non-exempt encryption.
	UsesNonExemptEncryption bool `json:"usesNonExemptEncryption,omitempty"`
	// The date and time the build was uploaded to App Store Connect.
	UploadedDate string `json:"uploadedDate,omitempty"`
	// The date and time the build will auto-expire and no longer be available for testing.
	ExpirationDate string `json:"expirationDate,omitempty"`
}

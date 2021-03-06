/*
 * App Store Connect API
 *
 * The App Store Connect API is a REST API used to build custom workflows as part of your app development life cycle and automate actions you take in App Store Connect. Use of the API requires authorization via JSON Web Tokens (JWT); you obtain keys to create the tokens from your organization’s App Store Connect account. See Creating API Keys for App Store Connect API to create your keys and tokens.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreconnect
// BetaInviteType An invite type that indicates if a beta tester was invited by an email invite or used a TestFlight public link to join a beta test.
type BetaInviteType string

// List of BetaInviteType
const (
	EMAIL BetaInviteType = "EMAIL"
	PUBLIC_LINK BetaInviteType = "PUBLIC_LINK"
)

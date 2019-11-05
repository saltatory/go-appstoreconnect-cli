# BetaGroupAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInternalGroup** | **bool** | A Boolean value that indicates whether the group is internal. Only existing users of App Store Connect may be added for internal beta testing. | [optional] 
**Name** | **string** | The name for the beta group. | [optional] 
**PublicLink** | **string** | The URL of the public link provided to your app&#39;s beta testers. | [optional] 
**PublicLinkEnabled** | **bool** | A Boolean value that indicates whether a public link is enabled. Enabling a link allows you to invite anyone outside of your team to beta test your app. When you share this link, testers will be able to install the beta version of your app on their devices in TestFlight and share the link with others. | [optional] 
**PublicLinkId** | **string** | The ID as part of the URL of the public link. | [optional] 
**PublicLinkLimit** | **int32** | The maximum number of testers that can join this beta group using the public link. Values must be between 1 and 10,000. | [optional] 
**PublicLinkLimitEnabled** | **bool** | A Boolean value that limits the number of testers who can join the beta group using the public link. | [optional] 
**CreatedDate** | [**time.Time**](time.Time.md) | The creation date of the beta group. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



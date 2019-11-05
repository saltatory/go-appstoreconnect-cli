# \TestFlightApi

All URIs are relative to *http://api.appstoreconnect.apple.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBetaTester**](TestFlightApi.md#AddBetaTester) | **Post** /betaTesters/{id}/relationships/betaGroups | Add a Beta Tester to Beta Groups
[**AssignBetaTesterToBuilds**](TestFlightApi.md#AssignBetaTesterToBuilds) | **Post** /betaTesters/{id}/relationships/builds | Individually Assign a Beta Tester to Builds
[**CreateBetaGroup**](TestFlightApi.md#CreateBetaGroup) | **Post** /betaGroups | Create a Beta Group
[**CreateBetaTester**](TestFlightApi.md#CreateBetaTester) | **Post** /betaTesters | Create a Beta Tester
[**DeleteBetaTester**](TestFlightApi.md#DeleteBetaTester) | **Delete** /betaTesters | Delete a Beta Tester
[**GetAllAppResourceIdsForBetaTester**](TestFlightApi.md#GetAllAppResourceIdsForBetaTester) | **Get** /betaTesters/{id}/relationships/apps | Get All App Resource IDs for a Beta Tester
[**GetAllBuildsAssignedToBetaTester**](TestFlightApi.md#GetAllBuildsAssignedToBetaTester) | **Get** /betaTesters/{id}/relationships/builds | Get All IDs of Builds Individually Assigned to a Beta Tester
[**ListAllAppsForBetaTester**](TestFlightApi.md#ListAllAppsForBetaTester) | **Get** /betaTesters/{id}/apps | List All Apps for a Beta Tester
[**ListAllBuildsAssignedToBetaTester**](TestFlightApi.md#ListAllBuildsAssignedToBetaTester) | **Get** /betaTesters/{id}/builds | List All Builds Individually Assigned to a Beta Tester
[**ListBetaTester**](TestFlightApi.md#ListBetaTester) | **Get** /betaTesters | List Beta Testers
[**ModifyBetaGroup**](TestFlightApi.md#ModifyBetaGroup) | **Patch** /betaGroups/{id} | Modify a Beta Group
[**ReadBetaTester**](TestFlightApi.md#ReadBetaTester) | **Get** /betaTesters/{id} | Read Beta Tester Information
[**RemoveBetaTesterAccessToApps**](TestFlightApi.md#RemoveBetaTesterAccessToApps) | **Delete** /betaTesters/{id}/relationships/apps | Remove a Beta Tester’s Access to Apps
[**RemoveBetaTesterFromBetaGroups**](TestFlightApi.md#RemoveBetaTesterFromBetaGroups) | **Delete** /betaTesters/{id}/relationships/betaGroups | Remove a Beta Tester from Beta Groups
[**SendInvitationToBetaTester**](TestFlightApi.md#SendInvitationToBetaTester) | **Post** /betaTesterInvitations | Send an Invitation to a Beta Tester
[**UnassignBetaTesterFromBuilds**](TestFlightApi.md#UnassignBetaTesterFromBuilds) | **Delete** /betaTesters/{id}/relationships/builds | Individually Unassign a Beta Tester from Builds



## AddBetaTester

> AddBetaTester(ctx, id, optional)

Add a Beta Tester to Beta Groups

Add one or more beta testers to a specific beta group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***AddBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBetaGroupsLinkagesRequest** | [**optional.Interface of BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignBetaTesterToBuilds

> AssignBetaTesterToBuilds(ctx, id, optional)

Individually Assign a Beta Tester to Builds

Individually assign a beta tester to a build.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***AssignBetaTesterToBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssignBetaTesterToBuildsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBuildsLinkagesRequest** | [**optional.Interface of BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBetaGroup

> BetaGroupResponse CreateBetaGroup(ctx, optional)

Create a Beta Group

Create a beta group associated with an app, optionally enabling TestFlight public links.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateBetaGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBetaGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaGroupCreateRequest** | [**optional.Interface of BetaGroupCreateRequest**](BetaGroupCreateRequest.md)|  | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBetaTester

> BetaTesterResponse CreateBetaTester(ctx, betaTesterCreateRequest)

Create a Beta Tester

Create a beta tester assigned to a group, a build, or an app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**betaTesterCreateRequest** | [**BetaTesterCreateRequest**](BetaTesterCreateRequest.md)| The resource data. | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBetaTester

> DeleteBetaTester(ctx, id)

Delete a Beta Tester

Remove a beta tester's ability to test all apps.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAppResourceIdsForBetaTester

> BetaTesterAppsLinkagesResponse GetAllAppResourceIdsForBetaTester(ctx, id, optional)

Get All App Resource IDs for a Beta Tester

Get a list of app resource IDs associated with a beta tester.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***GetAllAppResourceIdsForBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllAppResourceIdsForBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum is 200 | 

### Return type

[**BetaTesterAppsLinkagesResponse**](BetaTesterAppsLinkagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBuildsAssignedToBetaTester

> BetaTesterBuildsLinkagesResponse GetAllBuildsAssignedToBetaTester(ctx, id, optional)

Get All IDs of Builds Individually Assigned to a Beta Tester

Get a list of build resource IDs individually assigned to a specific beta tester.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***GetAllBuildsAssignedToBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllBuildsAssignedToBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of resources to return. Maximum is 200. | 

### Return type

[**BetaTesterBuildsLinkagesResponse**](BetaTesterBuildsLinkagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllAppsForBetaTester

> AppsResponse ListAllAppsForBetaTester(ctx, id, optional)

List All Apps for a Beta Tester

Get a list of apps that a beta tester can test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***ListAllAppsForBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllAppsForBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of resources to return. Maximum is 200. | 
 **fieldsApps** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllBuildsAssignedToBetaTester

> BuildsResponse ListAllBuildsAssignedToBetaTester(ctx, id, optional)

List All Builds Individually Assigned to a Beta Tester

Get a list of builds individually assigned to a specific beta tester.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***ListAllBuildsAssignedToBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllBuildsAssignedToBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of resources to return. Maximum is 200. | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 

### Return type

[**BuildsResponse**](BuildsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBetaTester

> BetaTestersResponse ListBetaTester(ctx, optional)

List Beta Testers

Find and list beta testers for all apps, builds, and beta groups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsApp** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **filterApps** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterBetaGroups** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterBuilds** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterEmail** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterFirstName** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterInviteType** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterLastName** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **include** | [**optional.Interface of []string**](string.md)|  | 
 **limit** | **optional.Int32**| Number of resources to return. Maximum is 200 | 
 **limitApps** | **optional.Int32**| Number of resources to return. Maximum is 50 | 
 **limitBetaGroups** | **optional.Int32**| Number of resources to return. Maximum is 50 | 
 **limitBuilds** | **optional.Int32**| Number of resources to return. Maximum is 50 | 
 **sort** | **optional.String**| Attributes by which to sort. | 

### Return type

[**BetaTestersResponse**](BetaTestersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBetaGroup

> BetaGroupResponse ModifyBetaGroup(ctx, id, optional)

Modify a Beta Group

Modify a beta group's metadata, including changing its Testflight public link status.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***ModifyBetaGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyBetaGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupUpdateRequest** | [**optional.Interface of BetaGroupUpdateRequest**](BetaGroupUpdateRequest.md)|  | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadBetaTester

> BetaTesterResponse ReadBetaTester(ctx, id, optional)

Read Beta Tester Information

Get a specific beta tester.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***ReadBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApp** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **fieldsBetaGroups** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **fieldsBetaTesters** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **fieldsBuilds** | [**optional.Interface of []string**](string.md)| Fields to return for included related types. | 
 **filterApps** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterBetaGroups** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterBuilds** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterEmail** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterFirstName** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterInviteType** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **filterLastName** | [**optional.Interface of []string**](string.md)| Attributes, relationships, and IDs by which to filter. | 
 **include** | [**optional.Interface of []string**](string.md)|  | 
 **limit** | **optional.Int32**| Number of resources to return. Maximum is 200 | 
 **limitApps** | **optional.Int32**| Number of resources to return. Maximum is 50 | 
 **limitBetaGroups** | **optional.Int32**| Number of resources to return. Maximum is 50 | 
 **limitBuilds** | **optional.Int32**| Number of resources to return. Maximum is 50 | 
 **sort** | **optional.String**| Attributes by which to sort. | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBetaTesterAccessToApps

> RemoveBetaTesterAccessToApps(ctx, id, optional)

Remove a Beta Tester’s Access to Apps

Remove a specific beta tester's access to test any builds of one or more apps.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***RemoveBetaTesterAccessToAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveBetaTesterAccessToAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterAppsLinkagesRequest** | [**optional.Interface of BetaTesterAppsLinkagesRequest**](BetaTesterAppsLinkagesRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBetaTesterFromBetaGroups

> RemoveBetaTesterFromBetaGroups(ctx, id, optional)

Remove a Beta Tester from Beta Groups

Remove a specific beta tester from one or more beta groups, revoking their access to test builds associated with those groups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***RemoveBetaTesterFromBetaGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveBetaTesterFromBetaGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBetaGroupsLinkagesRequest** | [**optional.Interface of BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInvitationToBetaTester

> BetaTesterInvitationResponse SendInvitationToBetaTester(ctx, optional)

Send an Invitation to a Beta Tester

Send or resend an invitation to a beta tester to test a specified app.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SendInvitationToBetaTesterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendInvitationToBetaTesterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaTesterInvitationCreateRequest** | [**optional.Interface of BetaTesterInvitationCreateRequest**](BetaTesterInvitationCreateRequest.md)|  | 

### Return type

[**BetaTesterInvitationResponse**](BetaTesterInvitationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignBetaTesterFromBuilds

> UnassignBetaTesterFromBuilds(ctx, id, optional)

Individually Unassign a Beta Tester from Builds

Remove an individually assigned beta tester's ability to test a build.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| (Required) An opaque resource ID that uniquely identifies the resource. | 
 **optional** | ***UnassignBetaTesterFromBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnassignBetaTesterFromBuildsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBuildsLinkagesRequest** | [**optional.Interface of BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


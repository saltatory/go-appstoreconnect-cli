# CommonErrorResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | (Required) A machine-readable code indicating the type of error. The code is a hierarchical value with levels of specificity separated by the &#39;.&#39; character. This value is parseable for programmatic error handling in code. | 
**Status** | **string** | (Required) The HTTP status code of the error. This status code usually matches the response&#39;s status code; however, if the request produces multiple errors, these two codes may differ. | 
**Id** | **string** | The unique ID of a specific instance of an error, request, and response. Use this ID when providing feedback to or debugging issues with Apple. | [optional] 
**Title** | **string** | (Required) A summary of the error. Do not use this field for programmatic error handling. | 
**Detail** | **string** | (Required) A detailed explanation of the error. Do not use this field for programmatic error handling. | 
**Source** | [**OneOfCommonErrorResponseErrorsJsonPointerCommonErrorResponseErrorsParameter**](oneOf&lt;CommonErrorResponse.Errors.JsonPointer,CommonErrorResponse.Errors.Parameter&gt;.md) | One of two possible types of values. source.parameter, provided when a query parameter produced the error, or source.JsonPointer, provided when a problem with the entity produced the error. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



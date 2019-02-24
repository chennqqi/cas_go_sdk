# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExampleGet**](DefaultApi.md#ExampleGet) | **Get** /example | Server example operation
[**PingGet**](DefaultApi.md#PingGet) | **Get** /ping | Server heartbeat operation


# **ExampleGet**
> ExampleGet(ctx, )
Server example operation

This is an example opeartion to show how security is applied to the call.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingGet**
> PingGet(ctx, )
Server heartbeat operation

This operation shows how to override the global security defined above, as we want to open it up for all users.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


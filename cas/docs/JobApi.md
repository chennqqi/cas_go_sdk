# \JobApi

All URIs are relative to *http://cas.ap-beijing.myqcloud.com/-*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VaultsVaultNameJobsGet**](JobApi.md#VaultsVaultNameJobsGet) | **Get** /vaults/{VaultName}/jobs | List Job
[**VaultsVaultNameJobsJobIDGet**](JobApi.md#VaultsVaultNameJobsJobIDGet) | **Get** /vaults/{VaultName}/jobs/&lt;JobID&gt; | Describe Job
[**VaultsVaultNameJobsJobIDOutputGet**](JobApi.md#VaultsVaultNameJobsJobIDOutputGet) | **Get** /vaults/{VaultName}/jobs/&lt;JobID&gt;/output | Get Job Output
[**VaultsVaultNameJobsPost**](JobApi.md#VaultsVaultNameJobsPost) | **Post** /vaults/{VaultName}/jobs | Initiate Job



## VaultsVaultNameJobsGet

> JobsList VaultsVaultNameJobsGet(ctx, vaultName, optional)

List Job

List Jobs 请求实现列出 Vault 的任务，包括正在进行的任务以及最近完成的任务。支持跨账户操作。当操作本账户时，UID为\"-\"。。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
 **optional** | ***VaultsVaultNameJobsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameJobsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completed** | **optional.Bool**|  | 
 **limit** | **optional.Int64**|  | 
 **marker** | **optional.String**|  | 
 **statuscode** | **optional.String**|  | 

### Return type

[**JobsList**](JobsList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameJobsJobIDGet

> OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo VaultsVaultNameJobsJobIDGet(ctx, vaultName, jobID)

Describe Job

Describe Job 请求实现获取Vault的具体任务信息。支持跨账户操作。当操作本账户时，UID为\"-\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**jobID** | **string**|  | 

### Return type

[**OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo**](oneOf&lt;JobArchiveSearchInfo,JobArchiveListSearchInfo,JobArchiveImportInfo,JobArchiveExportInfo&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameJobsJobIDOutputGet

> JobOutput VaultsVaultNameJobsJobIDOutputGet(ctx, vaultName, jobID, optional)

Get Job Output

请求用来输出缓存池中检索出来的 Archive 或Archive列表，缓存池中的内容24小时有效。请求所有数据成功后，返回 200 OK。请求部分数据成功时，返回 206 Partial Content。支持跨账户操作。当操作本账户时，UID 为\"-\"。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**jobID** | **string**|  | 
 **optional** | ***VaultsVaultNameJobsJobIDOutputGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameJobsJobIDOutputGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **optional.String**|  | 

### Return type

[**JobOutput**](JobOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameJobsPost

> VaultsVaultNameJobsPost(ctx, vaultName, optional)

Initiate Job

Describe Job Initiate Job 请求实现将档案或者档案列表取出到缓存池。操作完成后，用户可以通过 Get Job Output 请求读取对应档案或者档案列表。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
 **optional** | ***VaultsVaultNameJobsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameJobsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**optional.Interface of UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \JobApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UIDVaultsVaultNameJobsGet**](JobApi.md#UIDVaultsVaultNameJobsGet) | **Get** /&lt;UID&gt;/vaults/&lt;VaultName&gt;/jobs | List Job
[**UIDVaultsVaultNameJobsJobIDGet**](JobApi.md#UIDVaultsVaultNameJobsJobIDGet) | **Get** /&lt;UID&gt;/vaults/&lt;VaultName&gt;/jobs/&lt;JobID&gt; | Describe Job
[**UIDVaultsVaultNameJobsJobIDOutputGet**](JobApi.md#UIDVaultsVaultNameJobsJobIDOutputGet) | **Get** /&lt;UID&gt;/vaults/&lt;VaultName&gt;/jobs/&lt;JobID&gt;/output | Get Job Output
[**UIDVaultsVaultNameJobsPost**](JobApi.md#UIDVaultsVaultNameJobsPost) | **Post** /&lt;UID&gt;/vaults/&lt;VaultName&gt;/jobs | Initiate Job


# **UIDVaultsVaultNameJobsGet**
> JobsList UIDVaultsVaultNameJobsGet(ctx, uID, vaultName, optional)
List Job

List Jobs 请求实现列出 Vault 的任务，包括正在进行的任务以及最近完成的任务。支持跨账户操作。当操作本账户时，UID为\"-\"。。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**|  | 
  **vaultName** | **string**|  | 
 **optional** | ***UIDVaultsVaultNameJobsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UIDVaultsVaultNameJobsGetOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UIDVaultsVaultNameJobsJobIDGet**
> OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo UIDVaultsVaultNameJobsJobIDGet(ctx, uID, vaultName, jobID)
Describe Job

Describe Job 请求实现获取Vault的具体任务信息。支持跨账户操作。当操作本账户时，UID为\"-\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**|  | 
  **vaultName** | **string**|  | 
  **jobID** | **string**|  | 

### Return type

[**OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo**](oneOf&lt;JobArchiveSearchInfo,JobArchiveListSearchInfo,JobArchiveImportInfo,JobArchiveExportInfo&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UIDVaultsVaultNameJobsJobIDOutputGet**
> JobOutput UIDVaultsVaultNameJobsJobIDOutputGet(ctx, uID, vaultName, jobID, optional)
Get Job Output

请求用来输出缓存池中检索出来的 Archive 或Archive列表，缓存池中的内容24小时有效。请求所有数据成功后，返回 200 OK。请求部分数据成功时，返回 206 Partial Content。支持跨账户操作。当操作本账户时，UID 为\"-\"。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**|  | 
  **vaultName** | **string**|  | 
  **jobID** | **string**|  | 
 **optional** | ***UIDVaultsVaultNameJobsJobIDOutputGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UIDVaultsVaultNameJobsJobIDOutputGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **range_** | **optional.String**|  | 

### Return type

[**JobOutput**](JobOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UIDVaultsVaultNameJobsPost**
> UIDVaultsVaultNameJobsPost(ctx, uID, vaultName, optional)
Initiate Job

Describe Job Initiate Job 请求实现将档案或者档案列表取出到缓存池。操作完成后，用户可以通过 Get Job Output 请求读取对应档案或者档案列表。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**|  | 
  **vaultName** | **string**|  | 
 **optional** | ***UIDVaultsVaultNameJobsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UIDVaultsVaultNameJobsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**optional.Interface of UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


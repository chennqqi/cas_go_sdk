# \ArchiveApi

All URIs are relative to *http://cas.ap-beijing.myqcloud.com/-*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VaultsVaultNameArchivesArchiveIDDelete**](ArchiveApi.md#VaultsVaultNameArchivesArchiveIDDelete) | **Delete** /vaults/{VaultName}/archives/{ArchiveID} | Delete Archive
[**VaultsVaultNameArchivesPost**](ArchiveApi.md#VaultsVaultNameArchivesPost) | **Post** /vaults/{VaultName}/archives | Upload Archive
[**VaultsVaultNameMultipartUploadsGet**](ArchiveApi.md#VaultsVaultNameMultipartUploadsGet) | **Get** /vaults/{VaultName}/multipart-uploads | List Multipart Uploads
[**VaultsVaultNameMultipartUploadsPost**](ArchiveApi.md#VaultsVaultNameMultipartUploadsPost) | **Post** /vaults/{VaultName}/multipart-uploads | Initiate Multipart Upload
[**VaultsVaultNameMultipartUploadsUploadIDDelete**](ArchiveApi.md#VaultsVaultNameMultipartUploadsUploadIDDelete) | **Delete** /vaults/{VaultName}/multipart-uploads/{uploadID} | Abort Multipart Upload
[**VaultsVaultNameMultipartUploadsUploadIDGet**](ArchiveApi.md#VaultsVaultNameMultipartUploadsUploadIDGet) | **Get** /vaults/{VaultName}/multipart-uploads/{uploadID} | List Parts
[**VaultsVaultNameMultipartUploadsUploadIDPost**](ArchiveApi.md#VaultsVaultNameMultipartUploadsUploadIDPost) | **Post** /vaults/{VaultName}/multipart-uploads/{uploadID} | Complete Multipart Upload
[**VaultsVaultNameMultipartUploadsUploadIDPut**](ArchiveApi.md#VaultsVaultNameMultipartUploadsUploadIDPut) | **Put** /vaults/{VaultName}/multipart-uploads/{uploadID} | Upload Part



## VaultsVaultNameArchivesArchiveIDDelete

> VaultsVaultNameArchivesArchiveIDDelete(ctx, vaultName, archiveID)

Delete Archive

Delete Archive 请求实现删除一个 Archive。请求成功以后会返回 x-cas-archive-id 用来表示唯一的 Archive 文件。请求成功返回 204 No Content。在删除 Archive 后，您仍可能成功请求启动对已删除 Archive 的检索任务，但 Archive 检索任务会失败。 在您删除 Archive 时，对相应 Archive ID 正在进行的 Archive 检索可能成功，也可能不成功，具体取决于下面的场景：收到删除 Archive 请求时，Archive 检索任务正在下载 Archive 到缓存池，则 Archive 检索操作可能会失败。 收到删除 Archive 请求时，Archive 检索任务已经下载 Archive 到缓存池，则您将能够下载输出。支持跨账户操作。当操作本账户时，UID 为\"-\"。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**archiveID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameArchivesPost

> VaultsVaultNameArchivesPost(ctx, vaultName, xCasContentSha256, contentLength, xCasSha256TreeHash, body, optional)

Upload Archive

Upload Archive 请求实现上传一个 Archive 到指定 Vault。请求成功以后会返回 x-cas-archive-id 用来表示唯一的Archive 文件。请求成功返回 201 Created。上传文件时，可以指定 x-cas-archive-description 用来做文件内容备注。支持跨账户操作。当操作本账户时，UID为\"-\"。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**xCasContentSha256** | **string**|  | 
**contentLength** | **string**|  | 
**xCasSha256TreeHash** | **string**|  | 
**body** | ***os.File*****os.File**|  | 
 **optional** | ***VaultsVaultNameArchivesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameArchivesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **xCasArchiveDescription** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameMultipartUploadsGet

> VaultsSummary VaultsVaultNameMultipartUploadsGet(ctx, vaultName, optional)

List Multipart Uploads

List Multipart Uploads请求实现列出正在进行中的分段上传

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
 **optional** | ***VaultsVaultNameMultipartUploadsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameMultipartUploadsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**|  | 
 **marker** | **optional.String**|  | 

### Return type

[**VaultsSummary**](VaultsSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameMultipartUploadsPost

> VaultsVaultNameMultipartUploadsPost(ctx, vaultName, xCasPartSize, optional)

Initiate Multipart Upload

Initiate Multipart Upload请求实现初始化分段上传，此请求将返回一个Upload Id用以后续分段上传，此Upload Id有效期24小时。每次分段上传的段大小要求一致（除了最后一个段），且必须为1 MB乘以2的幂次

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**xCasPartSize** | **string**|  | 
 **optional** | ***VaultsVaultNameMultipartUploadsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameMultipartUploadsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCasArchiveDescription** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameMultipartUploadsUploadIDDelete

> VaultsVaultNameMultipartUploadsUploadIDDelete(ctx, vaultName, uploadID)

Abort Multipart Upload

Abort Multipart Upload请求实现终止分段上传。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**uploadID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameMultipartUploadsUploadIDGet

> ListParts VaultsVaultNameMultipartUploadsUploadIDGet(ctx, vaultName, uploadID, optional)

List Parts

List Parts请求实现列出已上传的数据段。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**uploadID** | **string**|  | 
 **optional** | ***VaultsVaultNameMultipartUploadsUploadIDGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameMultipartUploadsUploadIDGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int64**|  | 
 **marker** | **optional.String**|  | 

### Return type

[**ListParts**](ListParts.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameMultipartUploadsUploadIDPost

> VaultsVaultNameMultipartUploadsUploadIDPost(ctx, vaultName, uploadID, xCasSha256TreeHash, xCasArchiveSize)

Complete Multipart Upload

Complete Multipart Upload请求实现结束分段上传，形成文件。发起该请求时必须携带全文件的树形哈希值，服务端将比较用户上传的全文树形哈希和利用已上传分块得到的树形哈希，一致则请求成功，不一致则返回失败

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**uploadID** | **string**|  | 
**xCasSha256TreeHash** | **string**|  | 
**xCasArchiveSize** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameMultipartUploadsUploadIDPut

> VaultsVaultNameMultipartUploadsUploadIDPut(ctx, vaultName, uploadID, contentRange, xCasContentSha256, xCasSha256TreeHash, body, optional)

Upload Part

Upload Part请求实现上传档案的一段数据。支持乱序上传档案段，支持覆盖已上传的数据段。需在请求中标示该数据段在档案的字节范围。此外，支持并行上传数据段，最多可以上传 10000 段。当x-cas-sha256-tree-hash或x-cas-content-sha256与请求体中的真实文件校验和不一致时，请求返回错误。当Content-Length与请求体中的真实文件大小不一致时，请求返回错误。当Content-Range为必须以初始化分块时对应的块大小严格一致。例如，指定 4 194 304 字节 (4MB) 的段大小，则 0 到 4 194 303 字节 (4MB-1) 以及 4 194 304 (4MB) 到 8 388 607 (8MB-1) 为有效的段范围。2097152（ 2MB） 到6291456（ 6MB-1）为非法段范围。成功上传段后，将返回 204 No Content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
**uploadID** | **string**|  | 
**contentRange** | **string**|  | 
**xCasContentSha256** | **string**|  | 
**xCasSha256TreeHash** | **string**|  | 
**body** | ***os.File*****os.File**|  | 
 **optional** | ***VaultsVaultNameMultipartUploadsUploadIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameMultipartUploadsUploadIDPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **contentLength** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


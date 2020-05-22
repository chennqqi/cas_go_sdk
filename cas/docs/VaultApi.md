# \VaultApi

All URIs are relative to *http://cas.ap-beijing.myqcloud.com/-*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVault**](VaultApi.md#CreateVault) | **Put** /vaults/{VaultName} | 创建vault
[**GetVault**](VaultApi.md#GetVault) | **Get** /vaults/{VaultName} | Describe Vault
[**VaultsGet**](VaultApi.md#VaultsGet) | **Get** /vaults | List Vaults
[**VaultsVaultNameAccessPolicyDelete**](VaultApi.md#VaultsVaultNameAccessPolicyDelete) | **Delete** /vaults/{VaultName}/access-policy | Delete Vault Access Policy
[**VaultsVaultNameAccessPolicyGet**](VaultApi.md#VaultsVaultNameAccessPolicyGet) | **Get** /vaults/{VaultName}/access-policy | Get Vault Access Policy 请求读取一个 Vault 的权限
[**VaultsVaultNameAccessPolicyPut**](VaultApi.md#VaultsVaultNameAccessPolicyPut) | **Put** /vaults/{VaultName}/access-policy | Set Vault Access Policy
[**VaultsVaultNameDelete**](VaultApi.md#VaultsVaultNameDelete) | **Delete** /vaults/{VaultName} | 删除vault
[**VaultsVaultNameNotificationConfigurationDelete**](VaultApi.md#VaultsVaultNameNotificationConfigurationDelete) | **Delete** /vaults/{VaultName}/notification-configuration | Delete Vault Notifications
[**VaultsVaultNameNotificationConfigurationGet**](VaultApi.md#VaultsVaultNameNotificationConfigurationGet) | **Get** /vaults/{VaultName}/notification-configuration | Get Vault Notifications
[**VaultsVaultNameNotificationConfigurationPut**](VaultApi.md#VaultsVaultNameNotificationConfigurationPut) | **Put** /vaults/{VaultName}/notification-configuration | Set Vault Access Policy



## CreateVault

> CreateVault(ctx, vaultName)

创建vault

Create Vault 请求实现创建一个 Vault，每个用户支持创建 1000 个 Vault。成功后返回 201 Created。支持跨账户创建。当创建本账户下 valut 时，UID值为\"-\"。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

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


## GetVault

> VaultInfo GetVault(ctx, vaultName)

Describe Vault

Describe Vault 请求实现读取一个 Vault 的属性。档案数与档案总大小，每日盘点更新，非实时数据。请求成功后返回 200 OK。支持跨账户操作。当操作本账户下 valut 时，UID值为\"-\"。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

### Return type

[**VaultInfo**](VaultInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsGet

> VaultsSummary VaultsGet(ctx, optional)

List Vaults

List Vaults 接口实现列出该账户下所有的文件库。档案数与档案总大小，每日盘点更新，非实时数据。支持跨账户操作。当操作本账户时，UID为\"-\"。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VaultsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| 指定要返回的文件库最大数目。该值为正整数，取值1-1000，默认为 1000 | 
 **marker** | **optional.String**| 按字典序，从该 Marker 开始列出 Vault 的 QCS，如果为空则从头列出 。 | 

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


## VaultsVaultNameAccessPolicyDelete

> VaultsVaultNameAccessPolicyDelete(ctx, vaultName)

Delete Vault Access Policy

Delete Vault Access Policy 请求删除 Vault 的权限。只支持所有者操作，对应 UID 值为\"-\"。成功后返回 204 No Content。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

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


## VaultsVaultNameAccessPolicyGet

> Policy VaultsVaultNameAccessPolicyGet(ctx, vaultName)

Get Vault Access Policy 请求读取一个 Vault 的权限

Get Vault Access Policy 请求读取一个 Vault 的权限。只支持所有者操作，对应 UID 值为\"-\"。成功后返回 200 OK。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameAccessPolicyPut

> VaultsVaultNameAccessPolicyPut(ctx, vaultName, optional)

Set Vault Access Policy

Set Vault Access Policy 请求实现为一个 Vault 设置权限。具体策略语法参考『认证与鉴权』-『权限管理』只支持所有者设置权限，对应 UID 值为 \"-\"。成功后返回 204 No Content。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
 **optional** | ***VaultsVaultNameAccessPolicyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameAccessPolicyPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**optional.Interface of Policy**](Policy.md)|  | 

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


## VaultsVaultNameDelete

> VaultsVaultNameDelete(ctx, vaultName)

删除vault

Delete Vault 请求实现删除一个 Vault，删除前要求 Vault 下无 Archive 同时无 Archive 写入。删除 Vault 时同时删除其权限信息。请求成功后返回 204 NoContent。支持跨账户删除。当删除本账户下 Valut 时，UID值为\"-\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

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


## VaultsVaultNameNotificationConfigurationDelete

> VaultsVaultNameNotificationConfigurationDelete(ctx, vaultName)

Delete Vault Notifications

Delete Vault Notifications请求实现删除指定文件库通知回调策略 请求成功，返回 204 No Content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

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


## VaultsVaultNameNotificationConfigurationGet

> NotificationConfiguration VaultsVaultNameNotificationConfigurationGet(ctx, vaultName)

Get Vault Notifications

Get Vault Notifications请求实现读取指定文件库通知回调策略

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 

### Return type

[**NotificationConfiguration**](NotificationConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VaultsVaultNameNotificationConfigurationPut

> VaultsVaultNameNotificationConfigurationPut(ctx, vaultName, optional)

Set Vault Access Policy

Set Vault Access Policy 请求实现为一个 Vault 设置权限。具体策略语法参考『认证与鉴权』-『权限管理』 只支持所有者设置权限，对应 UID 值为 \"-\"。成功后返回 204 No Content。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultName** | **string**|  | 
 **optional** | ***VaultsVaultNameNotificationConfigurationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VaultsVaultNameNotificationConfigurationPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfiguration** | [**optional.Interface of NotificationConfiguration**](NotificationConfiguration.md)|  | 

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


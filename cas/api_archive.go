/*
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * Tecent Cloud Archive Storage Golang SDK.
 *
 * API version: 1.5.0
 * Contact: chennqqi@qq.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ArchiveApiService ArchiveApi service
type ArchiveApiService service

/*
UIDVaultsVaultNameArchivesArchiveIDDelete Delete Archive
Delete Archive 请求实现删除一个 Archive。请求成功以后会返回 x-cas-archive-id 用来表示唯一的 Archive 文件。请求成功返回 204 No Content。在删除 Archive 后，您仍可能成功请求启动对已删除 Archive 的检索任务，但 Archive 检索任务会失败。 在您删除 Archive 时，对相应 Archive ID 正在进行的 Archive 检索可能成功，也可能不成功，具体取决于下面的场景：收到删除 Archive 请求时，Archive 检索任务正在下载 Archive 到缓存池，则 Archive 检索操作可能会失败。 收到删除 Archive 请求时，Archive 检索任务已经下载 Archive 到缓存池，则您将能够下载输出。支持跨账户操作。当操作本账户时，UID 为\&quot;-\&quot;。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param archiveID
*/
func (a *ArchiveApiService) UIDVaultsVaultNameArchivesArchiveIDDelete(ctx _context.Context, uID string, vaultName string, archiveID string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/archives/{ArchiveID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"ArchiveID"+"}", _neturl.QueryEscape(parameterToString(archiveID, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UIDVaultsVaultNameArchivesPostOpts Optional parameters for the method 'UIDVaultsVaultNameArchivesPost'
type UIDVaultsVaultNameArchivesPostOpts struct {
    XCasArchiveDescription optional.String
}

/*
UIDVaultsVaultNameArchivesPost Upload Archive
Upload Archive 请求实现上传一个 Archive 到指定 Vault。请求成功以后会返回 x-cas-archive-id 用来表示唯一的Archive 文件。请求成功返回 201 Created。上传文件时，可以指定 x-cas-archive-description 用来做文件内容备注。支持跨账户操作。当操作本账户时，UID为\&quot;-\&quot;。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param xCasContentSha256
 * @param contentLength
 * @param xCasSha256TreeHash
 * @param body
 * @param optional nil or *UIDVaultsVaultNameArchivesPostOpts - Optional Parameters:
 * @param "XCasArchiveDescription" (optional.String) - 
*/
func (a *ArchiveApiService) UIDVaultsVaultNameArchivesPost(ctx _context.Context, uID string, vaultName string, xCasContentSha256 string, contentLength string, xCasSha256TreeHash string, body *os.File, localVarOptionals *UIDVaultsVaultNameArchivesPostOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/archives"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-cas-content-sha256"] = parameterToString(xCasContentSha256, "")
	localVarHeaderParams["Content-Length"] = parameterToString(contentLength, "")
	localVarHeaderParams["x-cas-sha256-tree-hash"] = parameterToString(xCasSha256TreeHash, "")
	if localVarOptionals != nil && localVarOptionals.XCasArchiveDescription.IsSet() {
		localVarHeaderParams["x-cas-archive-description"] = parameterToString(localVarOptionals.XCasArchiveDescription.Value(), "")
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UIDVaultsVaultNameMultipartUploadsGetOpts Optional parameters for the method 'UIDVaultsVaultNameMultipartUploadsGet'
type UIDVaultsVaultNameMultipartUploadsGetOpts struct {
    Limit optional.Int64
    Marker optional.String
}

/*
UIDVaultsVaultNameMultipartUploadsGet List Multipart Uploads
List Multipart Uploads请求实现列出正在进行中的分段上传
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsGetOpts - Optional Parameters:
 * @param "Limit" (optional.Int64) - 
 * @param "Marker" (optional.String) - 
@return VaultsSummary
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsGet(ctx _context.Context, uID string, vaultName string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsGetOpts) (VaultsSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VaultsSummary
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarHeaderParams["limit"] = parameterToString(localVarOptionals.Limit.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.Marker.IsSet() {
		localVarHeaderParams["marker"] = parameterToString(localVarOptionals.Marker.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UIDVaultsVaultNameMultipartUploadsPostOpts Optional parameters for the method 'UIDVaultsVaultNameMultipartUploadsPost'
type UIDVaultsVaultNameMultipartUploadsPostOpts struct {
    XCasArchiveDescription optional.String
}

/*
UIDVaultsVaultNameMultipartUploadsPost Initiate Multipart Upload
Initiate Multipart Upload请求实现初始化分段上传，此请求将返回一个Upload Id用以后续分段上传，此Upload Id有效期24小时。每次分段上传的段大小要求一致（除了最后一个段），且必须为1 MB乘以2的幂次
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param xCasPartSize
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsPostOpts - Optional Parameters:
 * @param "XCasArchiveDescription" (optional.String) - 
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsPost(ctx _context.Context, uID string, vaultName string, xCasPartSize string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsPostOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-cas-part-size"] = parameterToString(xCasPartSize, "")
	if localVarOptionals != nil && localVarOptionals.XCasArchiveDescription.IsSet() {
		localVarHeaderParams["x-cas-archive-description"] = parameterToString(localVarOptionals.XCasArchiveDescription.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
UIDVaultsVaultNameMultipartUploadsUploadIDDelete Abort Multipart Upload
Abort Multipart Upload请求实现终止分段上传。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDDelete(ctx _context.Context, uID string, vaultName string, uploadID string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/{uploadID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", _neturl.QueryEscape(parameterToString(uploadID, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts Optional parameters for the method 'UIDVaultsVaultNameMultipartUploadsUploadIDGet'
type UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts struct {
    Limit optional.Int64
    Marker optional.String
}

/*
UIDVaultsVaultNameMultipartUploadsUploadIDGet List Parts
List Parts请求实现列出已上传的数据段。
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts - Optional Parameters:
 * @param "Limit" (optional.Int64) - 
 * @param "Marker" (optional.String) - 
@return ListParts
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDGet(ctx _context.Context, uID string, vaultName string, uploadID string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts) (ListParts, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListParts
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/{uploadID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", _neturl.QueryEscape(parameterToString(uploadID, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarHeaderParams["limit"] = parameterToString(localVarOptionals.Limit.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.Marker.IsSet() {
		localVarHeaderParams["marker"] = parameterToString(localVarOptionals.Marker.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
UIDVaultsVaultNameMultipartUploadsUploadIDPost Complete Multipart Upload
Complete Multipart Upload请求实现结束分段上传，形成文件。发起该请求时必须携带全文件的树形哈希值，服务端将比较用户上传的全文树形哈希和利用已上传分块得到的树形哈希，一致则请求成功，不一致则返回失败
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
 * @param xCasSha256TreeHash
 * @param xCasArchiveSize
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDPost(ctx _context.Context, uID string, vaultName string, uploadID string, xCasSha256TreeHash string, xCasArchiveSize string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/{uploadID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", _neturl.QueryEscape(parameterToString(uploadID, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-cas-sha256-tree-hash"] = parameterToString(xCasSha256TreeHash, "")
	localVarHeaderParams["x-cas-archive-size"] = parameterToString(xCasArchiveSize, "")
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts Optional parameters for the method 'UIDVaultsVaultNameMultipartUploadsUploadIDPut'
type UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts struct {
    ContentLength optional.String
}

/*
UIDVaultsVaultNameMultipartUploadsUploadIDPut Upload Part
Upload Part请求实现上传档案的一段数据。支持乱序上传档案段，支持覆盖已上传的数据段。需在请求中标示该数据段在档案的字节范围。此外，支持并行上传数据段，最多可以上传 10000 段。当x-cas-sha256-tree-hash或x-cas-content-sha256与请求体中的真实文件校验和不一致时，请求返回错误。当Content-Length与请求体中的真实文件大小不一致时，请求返回错误。当Content-Range为必须以初始化分块时对应的块大小严格一致。例如，指定 4 194 304 字节 (4MB) 的段大小，则 0 到 4 194 303 字节 (4MB-1) 以及 4 194 304 (4MB) 到 8 388 607 (8MB-1) 为有效的段范围。2097152（ 2MB） 到6291456（ 6MB-1）为非法段范围。成功上传段后，将返回 204 No Content
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
 * @param contentRange
 * @param xCasContentSha256
 * @param xCasSha256TreeHash
 * @param body
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts - Optional Parameters:
 * @param "ContentLength" (optional.String) - 
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDPut(ctx _context.Context, uID string, vaultName string, uploadID string, contentRange string, xCasContentSha256 string, xCasSha256TreeHash string, body *os.File, localVarOptionals *UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/{uploadID}"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", _neturl.QueryEscape(parameterToString(uID, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", _neturl.QueryEscape(parameterToString(vaultName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", _neturl.QueryEscape(parameterToString(uploadID, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Content-Range"] = parameterToString(contentRange, "")
	localVarHeaderParams["x-cas-content-sha256"] = parameterToString(xCasContentSha256, "")
	localVarHeaderParams["x-cas-sha256-tree-hash"] = parameterToString(xCasSha256TreeHash, "")
	if localVarOptionals != nil && localVarOptionals.ContentLength.IsSet() {
		localVarHeaderParams["Content-Length"] = parameterToString(localVarOptionals.ContentLength.Value(), "")
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 4XX {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

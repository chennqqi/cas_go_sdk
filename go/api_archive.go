// Copyright 2019 chennqqi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ArchiveApiService service

/*
ArchiveApiService Delete Archive
Delete Archive 请求实现删除一个 Archive。请求成功以后会返回 x-cas-archive-id 用来表示唯一的 Archive 文件。请求成功返回 204 No Content。在删除 Archive 后，您仍可能成功请求启动对已删除 Archive 的检索任务，但 Archive 检索任务会失败。 在您删除 Archive 时，对相应 Archive ID 正在进行的 Archive 检索可能成功，也可能不成功，具体取决于下面的场景：收到删除 Archive 请求时，Archive 检索任务正在下载 Archive 到缓存池，则 Archive 检索操作可能会失败。 收到删除 Archive 请求时，Archive 检索任务已经下载 Archive 到缓存池，则您将能够下载输出。支持跨账户操作。当操作本账户时，UID 为\&quot;-\&quot;。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param archiveID
*/
func (a *ArchiveApiService) UIDVaultsVaultNameArchivesArchiveIDDelete(ctx context.Context, uID string, vaultName string, archiveID string) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/archives/<ArchiveID>"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ArchiveID"+"}", fmt.Sprintf("%v", archiveID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ArchiveApiService Upload Archive
Upload Archive 请求实现上传一个 Archive 到指定 Vault。请求成功以后会返回 x-cas-archive-id 用来表示唯一的Archive 文件。请求成功返回 201 Created。上传文件时，可以指定 x-cas-archive-description 用来做文件内容备注。支持跨账户操作。当操作本账户时，UID为\&quot;-\&quot;。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param xCasContentSha256
 * @param contentLength
 * @param xCasSha256TreeHash
 * @param optional nil or *UIDVaultsVaultNameArchivesPostOpts - Optional Parameters:
 * @param "XCasArchiveDescription" (optional.String) -
*/

type UIDVaultsVaultNameArchivesPostOpts struct {
	XCasArchiveDescription optional.String
}

func (a *ArchiveApiService) UIDVaultsVaultNameArchivesPost(ctx context.Context, uID string, vaultName string, xCasContentSha256 string, contentLength string, xCasSha256TreeHash string, body io.Reader, localVarOptionals *UIDVaultsVaultNameArchivesPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/archives"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ArchiveApiService List Multipart Uploads
List Multipart Uploads请求实现列出正在进行中的分段上传
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsGetOpts - Optional Parameters:
 * @param "Limit" (optional.Int64) -
 * @param "Marker" (optional.String) -
@return VaultsSummary
*/

type UIDVaultsVaultNameMultipartUploadsGetOpts struct {
	Limit  optional.Int64
	Marker optional.String
}

func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsGet(ctx context.Context, uID string, vaultName string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsGetOpts) (VaultsSummary, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VaultsSummary
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v VaultsSummary
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ArchiveApiService Initiate Multipart Upload
Initiate Multipart Upload请求实现初始化分段上传，此请求将返回一个Upload Id用以后续分段上传，此Upload Id有效期24小时。每次分段上传的段大小要求一致（除了最后一个段），且必须为1 MB乘以2的幂次
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param xCasPartSize
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsPostOpts - Optional Parameters:
 * @param "XCasArchiveDescription" (optional.String) -
*/

type UIDVaultsVaultNameMultipartUploadsPostOpts struct {
	XCasArchiveDescription optional.String
}

func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsPost(ctx context.Context, uID string, vaultName string, xCasPartSize string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ArchiveApiService Abort Multipart Upload
Abort Multipart Upload请求实现终止分段上传。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDDelete(ctx context.Context, uID string, vaultName string, uploadID string) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/<uploadID>"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", fmt.Sprintf("%v", uploadID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ArchiveApiService List Parts
List Parts请求实现列出已上传的数据段。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
 * @param optional nil or *UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts - Optional Parameters:
 * @param "Limit" (optional.Int64) -
 * @param "Marker" (optional.String) -
@return ListParts
*/

type UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts struct {
	Limit  optional.Int64
	Marker optional.String
}

func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDGet(ctx context.Context, uID string, vaultName string, uploadID string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts) (ListParts, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListParts
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/<uploadID>"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", fmt.Sprintf("%v", uploadID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ListParts
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ArchiveApiService Complete Multipart Upload
Complete Multipart Upload请求实现结束分段上传，形成文件。发起该请求时必须携带全文件的树形哈希值，服务端将比较用户上传的全文树形哈希和利用已上传分块得到的树形哈希，一致则请求成功，不一致则返回失败
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param uploadID
 * @param xCasSha256TreeHash
 * @param xCasArchiveSize
*/
func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDPost(ctx context.Context, uID string, vaultName string, uploadID string, xCasSha256TreeHash string, xCasArchiveSize string) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/<uploadID>"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", fmt.Sprintf("%v", uploadID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ArchiveApiService Upload Part
Upload Part请求实现上传档案的一段数据。支持乱序上传档案段，支持覆盖已上传的数据段。需在请求中标示该数据段在档案的字节范围。此外，支持并行上传数据段，最多可以上传 10000 段。当x-cas-sha256-tree-hash或x-cas-content-sha256与请求体中的真实文件校验和不一致时，请求返回错误。当Content-Length与请求体中的真实文件大小不一致时，请求返回错误。当Content-Range为必须以初始化分块时对应的块大小严格一致。例如，指定 4 194 304 字节 (4MB) 的段大小，则 0 到 4 194 303 字节 (4MB-1) 以及 4 194 304 (4MB) 到 8 388 607 (8MB-1) 为有效的段范围。2097152（ 2MB） 到6291456（ 6MB-1）为非法段范围。成功上传段后，将返回 204 No Content
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
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

type UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts struct {
	ContentLength optional.String
}

func (a *ArchiveApiService) UIDVaultsVaultNameMultipartUploadsUploadIDPut(ctx context.Context, uID string, vaultName string, uploadID string, contentRange string, xCasContentSha256 string, xCasSha256TreeHash string, localVarOptionals *UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts, body io.Reader) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/multipart-uploads/<uploadID>"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uploadID"+"}", fmt.Sprintf("%v", uploadID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

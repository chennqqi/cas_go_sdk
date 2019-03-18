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

type JobApiService service

/*
JobApiService List Job
List Jobs 请求实现列出 Vault 的任务，包括正在进行的任务以及最近完成的任务。支持跨账户操作。当操作本账户时，UID为\&quot;-\&quot;。。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param optional nil or *UIDVaultsVaultNameJobsGetOpts - Optional Parameters:
 * @param "Completed" (optional.Bool) -
 * @param "Limit" (optional.Int64) -
 * @param "Marker" (optional.String) -
 * @param "Statuscode" (optional.String) -
@return JobsList
*/

type UIDVaultsVaultNameJobsGetOpts struct {
	Completed  optional.Bool
	Limit      optional.Int64
	Marker     optional.String
	Statuscode optional.String
}

func (a *JobApiService) UIDVaultsVaultNameJobsGet(ctx context.Context, uID string, vaultName string, localVarOptionals *UIDVaultsVaultNameJobsGetOpts) (JobsList, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobsList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/jobs"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Completed.IsSet() {
		localVarQueryParams.Add("completed", parameterToString(localVarOptionals.Completed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Marker.IsSet() {
		localVarQueryParams.Add("marker", parameterToString(localVarOptionals.Marker.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Statuscode.IsSet() {
		localVarQueryParams.Add("statuscode", parameterToString(localVarOptionals.Statuscode.Value(), ""))
	}
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
			var v JobsList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if statusCode4XX(localVarHttpResponse.StatusCode) {
			var v ErrorMessage
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
JobApiService Describe Job
Describe Job 请求实现获取Vault的具体任务信息。支持跨账户操作。当操作本账户时，UID为\&quot;-\&quot;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param jobID
@return OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo
*/
func (a *JobApiService) UIDVaultsVaultNameJobsJobIDGet(ctx context.Context, uID string, vaultName string, jobID string) (OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/jobs/<JobID>"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"JobID"+"}", fmt.Sprintf("%v", jobID), -1)

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
			var v OneOfJobArchiveSearchInfoJobArchiveListSearchInfoJobArchiveImportInfoJobArchiveExportInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if statusCode4XX(localVarHttpResponse.StatusCode) {
			var v ErrorMessage
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
JobApiService Get Job Output
请求用来输出缓存池中检索出来的 Archive 或Archive列表，缓存池中的内容24小时有效。请求所有数据成功后，返回 200 OK。请求部分数据成功时，返回 206 Partial Content。支持跨账户操作。当操作本账户时，UID 为\&quot;-\&quot;。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param jobID
 * @param optional nil or *UIDVaultsVaultNameJobsJobIDOutputGetOpts - Optional Parameters:
 * @param "Range_" (optional.String) -
@return JobOutput/or application/octet-stream
如果返回是application/octet-stream, 应当在函数外面关闭response.Body
*/

type UIDVaultsVaultNameJobsJobIDOutputGetOpts struct {
	Range_ optional.String
}

func (a *JobApiService) UIDVaultsVaultNameJobsJobIDOutputGet(ctx context.Context, uID string, vaultName string, jobID string, localVarOptionals *UIDVaultsVaultNameJobsJobIDOutputGetOpts) (JobOutput, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JobOutput
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/jobs/<JobID>/output"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"JobID"+"}", fmt.Sprintf("%v", jobID), -1)

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
	localVarHttpHeaderAccepts := []string{"application/json", "application/octet-stream"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
		localVarHeaderParams["Range"] = parameterToString(localVarOptionals.Range_.Value(), "")
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
	contentType := localVarHttpResponse.Header.Get("Content-Type")
	if contentType == "application/octet-stream" {
		return localVarReturnValue, localVarHttpResponse, nil
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
			var v JobOutput
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 206 {
			var v JobOutput
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if statusCode4XX(localVarHttpResponse.StatusCode) {
			var v ErrorMessage
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
JobApiService Initiate Job
Describe Job Initiate Job 请求实现将档案或者档案列表取出到缓存池。操作完成后，用户可以通过 Get Job Output 请求读取对应档案或者档案列表。
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uID
 * @param vaultName
 * @param optional nil or *UIDVaultsVaultNameJobsPostOpts - Optional Parameters:
 * @param "UNKNOWNBASETYPE" (optional.Interface of UNKNOWN_BASE_TYPE) -
*/

type UNKNOWN_BASE_TYPE interface {
	//	MarshalJSON() ([]byte, error)
	//	UnMarshalJSON(interface{}) error
}

type UIDVaultsVaultNameJobsPostOpts struct {
	UNKNOWNBASETYPE optional.Interface
}

func (a *JobApiService) UIDVaultsVaultNameJobsPost(ctx context.Context, uID string, vaultName string, localVarOptionals *UIDVaultsVaultNameJobsPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{UID}/vaults/{VaultName}/jobs"
	localVarPath = strings.Replace(localVarPath, "{"+"UID"+"}", fmt.Sprintf("%v", uID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"VaultName"+"}", fmt.Sprintf("%v", vaultName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.UNKNOWNBASETYPE.IsSet() {
		localVarOptionalUNKNOWNBASETYPE, localVarOptionalUNKNOWNBASETYPEok := localVarOptionals.UNKNOWNBASETYPE.Value().(UNKNOWN_BASE_TYPE)
		if !localVarOptionalUNKNOWNBASETYPEok {
			return nil, reportError("uNKNOWNBASETYPE should be UNKNOWN_BASE_TYPE")
		}
		localVarPostBody = &localVarOptionalUNKNOWNBASETYPE
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
		if statusCode4XX(localVarHttpResponse.StatusCode) {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

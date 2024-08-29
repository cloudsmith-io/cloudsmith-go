/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.498.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// VulnerabilitiesApiService VulnerabilitiesApi service
type VulnerabilitiesApiService service

type ApiVulnerabilitiesNamespaceListRequest struct {
	ctx        context.Context
	ApiService *VulnerabilitiesApiService
	owner      string
	page       *int64
	pageSize   *int64
}

// A page number within the paginated result set.
func (r ApiVulnerabilitiesNamespaceListRequest) Page(page int64) ApiVulnerabilitiesNamespaceListRequest {
	r.page = &page
	return r
}

// Number of results to return per page.
func (r ApiVulnerabilitiesNamespaceListRequest) PageSize(pageSize int64) ApiVulnerabilitiesNamespaceListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiVulnerabilitiesNamespaceListRequest) Execute() ([]VulnerabilityScanResultsList, *http.Response, error) {
	return r.ApiService.VulnerabilitiesNamespaceListExecute(r)
}

/*
VulnerabilitiesNamespaceList Lists scan results for a specific namespace.

Lists scan results for a specific namespace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @return ApiVulnerabilitiesNamespaceListRequest
*/
func (a *VulnerabilitiesApiService) VulnerabilitiesNamespaceList(ctx context.Context, owner string) ApiVulnerabilitiesNamespaceListRequest {
	return ApiVulnerabilitiesNamespaceListRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
	}
}

// Execute executes the request
//  @return []VulnerabilityScanResultsList
func (a *VulnerabilitiesApiService) VulnerabilitiesNamespaceListExecute(r ApiVulnerabilitiesNamespaceListRequest) ([]VulnerabilityScanResultsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VulnerabilityScanResultsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.VulnerabilitiesNamespaceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerabilities/{owner}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVulnerabilitiesPackageListRequest struct {
	ctx        context.Context
	ApiService *VulnerabilitiesApiService
	owner      string
	repo       string
	package_   string
	page       *int64
	pageSize   *int64
}

// A page number within the paginated result set.
func (r ApiVulnerabilitiesPackageListRequest) Page(page int64) ApiVulnerabilitiesPackageListRequest {
	r.page = &page
	return r
}

// Number of results to return per page.
func (r ApiVulnerabilitiesPackageListRequest) PageSize(pageSize int64) ApiVulnerabilitiesPackageListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiVulnerabilitiesPackageListRequest) Execute() ([]VulnerabilityScanResultsList, *http.Response, error) {
	return r.ApiService.VulnerabilitiesPackageListExecute(r)
}

/*
VulnerabilitiesPackageList Lists scan results for a specific package.

Lists scan results for a specific package.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @param repo
 @param package_
 @return ApiVulnerabilitiesPackageListRequest
*/
func (a *VulnerabilitiesApiService) VulnerabilitiesPackageList(ctx context.Context, owner string, repo string, package_ string) ApiVulnerabilitiesPackageListRequest {
	return ApiVulnerabilitiesPackageListRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
		package_:   package_,
	}
}

// Execute executes the request
//  @return []VulnerabilityScanResultsList
func (a *VulnerabilitiesApiService) VulnerabilitiesPackageListExecute(r ApiVulnerabilitiesPackageListRequest) ([]VulnerabilityScanResultsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VulnerabilityScanResultsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.VulnerabilitiesPackageList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerabilities/{owner}/{repo}/{package}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"package"+"}", url.PathEscape(parameterValueToString(r.package_, "package_")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVulnerabilitiesReadRequest struct {
	ctx        context.Context
	ApiService *VulnerabilitiesApiService
	owner      string
	repo       string
	package_   string
	identifier string
}

func (r ApiVulnerabilitiesReadRequest) Execute() (*VulnerabilityScanResults, *http.Response, error) {
	return r.ApiService.VulnerabilitiesReadExecute(r)
}

/*
VulnerabilitiesRead Get a scan result.

Get a scan result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @param repo
 @param package_
 @param identifier
 @return ApiVulnerabilitiesReadRequest
*/
func (a *VulnerabilitiesApiService) VulnerabilitiesRead(ctx context.Context, owner string, repo string, package_ string, identifier string) ApiVulnerabilitiesReadRequest {
	return ApiVulnerabilitiesReadRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
		package_:   package_,
		identifier: identifier,
	}
}

// Execute executes the request
//  @return VulnerabilityScanResults
func (a *VulnerabilitiesApiService) VulnerabilitiesReadExecute(r ApiVulnerabilitiesReadRequest) (*VulnerabilityScanResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VulnerabilityScanResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.VulnerabilitiesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerabilities/{owner}/{repo}/{package}/{identifier}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"package"+"}", url.PathEscape(parameterValueToString(r.package_, "package_")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterValueToString(r.identifier, "identifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVulnerabilitiesRepoListRequest struct {
	ctx        context.Context
	ApiService *VulnerabilitiesApiService
	owner      string
	repo       string
	page       *int64
	pageSize   *int64
}

// A page number within the paginated result set.
func (r ApiVulnerabilitiesRepoListRequest) Page(page int64) ApiVulnerabilitiesRepoListRequest {
	r.page = &page
	return r
}

// Number of results to return per page.
func (r ApiVulnerabilitiesRepoListRequest) PageSize(pageSize int64) ApiVulnerabilitiesRepoListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiVulnerabilitiesRepoListRequest) Execute() ([]VulnerabilityScanResultsList, *http.Response, error) {
	return r.ApiService.VulnerabilitiesRepoListExecute(r)
}

/*
VulnerabilitiesRepoList Lists scan results for a specific repository.

Lists scan results for a specific repository.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @param repo
 @return ApiVulnerabilitiesRepoListRequest
*/
func (a *VulnerabilitiesApiService) VulnerabilitiesRepoList(ctx context.Context, owner string, repo string) ApiVulnerabilitiesRepoListRequest {
	return ApiVulnerabilitiesRepoListRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
	}
}

// Execute executes the request
//  @return []VulnerabilityScanResultsList
func (a *VulnerabilitiesApiService) VulnerabilitiesRepoListExecute(r ApiVulnerabilitiesRepoListRequest) ([]VulnerabilityScanResultsList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VulnerabilityScanResultsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.VulnerabilitiesRepoList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerabilities/{owner}/{repo}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterValueToString(r.repo, "repo")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.275.0
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// MetricsApiService MetricsApi service
type MetricsApiService service

type ApiMetricsEntitlementsAccountListRequest struct {
	ctx        context.Context
	ApiService *MetricsApiService
	owner      string
	page       *int64
	pageSize   *int64
	finish     *string
	start      *string
	tokens     *string
}

// A page number within the paginated result set.
func (r ApiMetricsEntitlementsAccountListRequest) Page(page int64) ApiMetricsEntitlementsAccountListRequest {
	r.page = &page
	return r
}

// Number of results to return per page.
func (r ApiMetricsEntitlementsAccountListRequest) PageSize(pageSize int64) ApiMetricsEntitlementsAccountListRequest {
	r.pageSize = &pageSize
	return r
}

// Include metrics upto and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;.
func (r ApiMetricsEntitlementsAccountListRequest) Finish(finish string) ApiMetricsEntitlementsAccountListRequest {
	r.finish = &finish
	return r
}

// Include metrics from and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;.
func (r ApiMetricsEntitlementsAccountListRequest) Start(start string) ApiMetricsEntitlementsAccountListRequest {
	r.start = &start
	return r
}

// A comma seperated list of tokens (slug perm) to include in the results.
func (r ApiMetricsEntitlementsAccountListRequest) Tokens(tokens string) ApiMetricsEntitlementsAccountListRequest {
	r.tokens = &tokens
	return r
}

func (r ApiMetricsEntitlementsAccountListRequest) Execute() (*EntitlementUsageMetrics, *http.Response, error) {
	return r.ApiService.MetricsEntitlementsAccountListExecute(r)
}

/*
MetricsEntitlementsAccountList View for listing entitlement token metrics, across an account.

View for listing entitlement token metrics, across an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @return ApiMetricsEntitlementsAccountListRequest
*/
func (a *MetricsApiService) MetricsEntitlementsAccountList(ctx context.Context, owner string) ApiMetricsEntitlementsAccountListRequest {
	return ApiMetricsEntitlementsAccountListRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
	}
}

// Execute executes the request
//  @return EntitlementUsageMetrics
func (a *MetricsApiService) MetricsEntitlementsAccountListExecute(r ApiMetricsEntitlementsAccountListRequest) (*EntitlementUsageMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntitlementUsageMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.MetricsEntitlementsAccountList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/entitlements/{owner}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.finish != nil {
		localVarQueryParams.Add("finish", parameterToString(*r.finish, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.tokens != nil {
		localVarQueryParams.Add("tokens", parameterToString(*r.tokens, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiMetricsEntitlementsRepoListRequest struct {
	ctx        context.Context
	ApiService *MetricsApiService
	owner      string
	repo       string
	page       *int64
	pageSize   *int64
	finish     *string
	start      *string
	tokens     *string
}

// A page number within the paginated result set.
func (r ApiMetricsEntitlementsRepoListRequest) Page(page int64) ApiMetricsEntitlementsRepoListRequest {
	r.page = &page
	return r
}

// Number of results to return per page.
func (r ApiMetricsEntitlementsRepoListRequest) PageSize(pageSize int64) ApiMetricsEntitlementsRepoListRequest {
	r.pageSize = &pageSize
	return r
}

// Include metrics upto and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;.
func (r ApiMetricsEntitlementsRepoListRequest) Finish(finish string) ApiMetricsEntitlementsRepoListRequest {
	r.finish = &finish
	return r
}

// Include metrics from and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;.
func (r ApiMetricsEntitlementsRepoListRequest) Start(start string) ApiMetricsEntitlementsRepoListRequest {
	r.start = &start
	return r
}

// A comma seperated list of tokens (slug perm) to include in the results.
func (r ApiMetricsEntitlementsRepoListRequest) Tokens(tokens string) ApiMetricsEntitlementsRepoListRequest {
	r.tokens = &tokens
	return r
}

func (r ApiMetricsEntitlementsRepoListRequest) Execute() (*EntitlementUsageMetrics, *http.Response, error) {
	return r.ApiService.MetricsEntitlementsRepoListExecute(r)
}

/*
MetricsEntitlementsRepoList View for listing entitlement token metrics, for a repository.

View for listing entitlement token metrics, for a repository.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @param repo
 @return ApiMetricsEntitlementsRepoListRequest
*/
func (a *MetricsApiService) MetricsEntitlementsRepoList(ctx context.Context, owner string, repo string) ApiMetricsEntitlementsRepoListRequest {
	return ApiMetricsEntitlementsRepoListRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
	}
}

// Execute executes the request
//  @return EntitlementUsageMetrics
func (a *MetricsApiService) MetricsEntitlementsRepoListExecute(r ApiMetricsEntitlementsRepoListRequest) (*EntitlementUsageMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntitlementUsageMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.MetricsEntitlementsRepoList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/entitlements/{owner}/{repo}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterToString(r.repo, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.finish != nil {
		localVarQueryParams.Add("finish", parameterToString(*r.finish, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.tokens != nil {
		localVarQueryParams.Add("tokens", parameterToString(*r.tokens, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiMetricsPackagesListRequest struct {
	ctx        context.Context
	ApiService *MetricsApiService
	owner      string
	repo       string
	page       *int64
	pageSize   *int64
	finish     *string
	packages   *string
	start      *string
}

// A page number within the paginated result set.
func (r ApiMetricsPackagesListRequest) Page(page int64) ApiMetricsPackagesListRequest {
	r.page = &page
	return r
}

// Number of results to return per page.
func (r ApiMetricsPackagesListRequest) PageSize(pageSize int64) ApiMetricsPackagesListRequest {
	r.pageSize = &pageSize
	return r
}

// Include metrics upto and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;.
func (r ApiMetricsPackagesListRequest) Finish(finish string) ApiMetricsPackagesListRequest {
	r.finish = &finish
	return r
}

// A comma seperated list of packages (slug perm) to include in the results.
func (r ApiMetricsPackagesListRequest) Packages(packages string) ApiMetricsPackagesListRequest {
	r.packages = &packages
	return r
}

// Include metrics from and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;.
func (r ApiMetricsPackagesListRequest) Start(start string) ApiMetricsPackagesListRequest {
	r.start = &start
	return r
}

func (r ApiMetricsPackagesListRequest) Execute() (*PackageUsageMetrics, *http.Response, error) {
	return r.ApiService.MetricsPackagesListExecute(r)
}

/*
MetricsPackagesList View for listing package usage metrics, for a repository.

View for listing package usage metrics, for a repository.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @param repo
 @return ApiMetricsPackagesListRequest
*/
func (a *MetricsApiService) MetricsPackagesList(ctx context.Context, owner string, repo string) ApiMetricsPackagesListRequest {
	return ApiMetricsPackagesListRequest{
		ApiService: a,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
	}
}

// Execute executes the request
//  @return PackageUsageMetrics
func (a *MetricsApiService) MetricsPackagesListExecute(r ApiMetricsPackagesListRequest) (*PackageUsageMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PackageUsageMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.MetricsPackagesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/packages/{owner}/{repo}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterToString(r.repo, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.finish != nil {
		localVarQueryParams.Add("finish", parameterToString(*r.finish, ""))
	}
	if r.packages != nil {
		localVarQueryParams.Add("packages", parameterToString(*r.packages, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type WebhookKeysApi interface {

	/*
		CreateWebhookKey Create Webhook Key

		Create a webhook key

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateWebhookKeyRequest
	*/
	CreateWebhookKey(ctx context.Context) ApiCreateWebhookKeyRequest

	// CreateWebhookKeyExecute executes the request
	//  @return WebhookKeyCreateResponse
	CreateWebhookKeyExecute(r ApiCreateWebhookKeyRequest) (*WebhookKeyCreateResponse, *http.Response, error)

	/*
		DeleteWebhookKey Delete Webhook key

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiDeleteWebhookKeyRequest
	*/
	DeleteWebhookKey(ctx context.Context, id string) ApiDeleteWebhookKeyRequest

	// DeleteWebhookKeyExecute executes the request
	//  @return WebhookKeyDeleteResponse
	DeleteWebhookKeyExecute(r ApiDeleteWebhookKeyRequest) (*WebhookKeyDeleteResponse, *http.Response, error)

	/*
		GetWebhookKey Get Webhook Key

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiGetWebhookKeyRequest
	*/
	GetWebhookKey(ctx context.Context, id string) ApiGetWebhookKeyRequest

	// GetWebhookKeyExecute executes the request
	//  @return WebhookKeyResponse
	GetWebhookKeyExecute(r ApiGetWebhookKeyRequest) (*WebhookKeyResponse, *http.Response, error)

	/*
		GetWebhookKeys Get List of Webhook Keys

		Consume the list of webhook keys you have

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetWebhookKeysRequest
	*/
	GetWebhookKeys(ctx context.Context) ApiGetWebhookKeysRequest

	// GetWebhookKeysExecute executes the request
	//  @return GetWebhookKeysResponse
	GetWebhookKeysExecute(r ApiGetWebhookKeysRequest) (*GetWebhookKeysResponse, *http.Response, error)

	/*
		UpdateWebhookKey Update Webhook Key

		updates an existing webhook key

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiUpdateWebhookKeyRequest
	*/
	UpdateWebhookKey(ctx context.Context, id string) ApiUpdateWebhookKeyRequest

	// UpdateWebhookKeyExecute executes the request
	//  @return WebhookKeyResponse
	UpdateWebhookKeyExecute(r ApiUpdateWebhookKeyRequest) (*WebhookKeyResponse, *http.Response, error)
}

// WebhookKeysApiService WebhookKeysApi service
type WebhookKeysApiService service

type ApiCreateWebhookKeyRequest struct {
	ctx               context.Context
	ApiService        WebhookKeysApi
	acceptLanguage    *string
	webhookKeyRequest *WebhookKeyRequest
}

// Use for knowing which language to use
func (r ApiCreateWebhookKeyRequest) AcceptLanguage(acceptLanguage string) ApiCreateWebhookKeyRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiCreateWebhookKeyRequest) WebhookKeyRequest(webhookKeyRequest WebhookKeyRequest) ApiCreateWebhookKeyRequest {
	r.webhookKeyRequest = &webhookKeyRequest
	return r
}

func (r ApiCreateWebhookKeyRequest) Execute() (*WebhookKeyCreateResponse, *http.Response, error) {
	return r.ApiService.CreateWebhookKeyExecute(r)
}

/*
CreateWebhookKey Create Webhook Key

Create a webhook key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateWebhookKeyRequest
*/
func (a *WebhookKeysApiService) CreateWebhookKey(ctx context.Context) ApiCreateWebhookKeyRequest {
	return ApiCreateWebhookKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WebhookKeyCreateResponse
func (a *WebhookKeysApiService) CreateWebhookKeyExecute(r ApiCreateWebhookKeyRequest) (*WebhookKeyCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebhookKeyCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookKeysApiService.CreateWebhookKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.conekta-v2.1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.webhookKeyRequest
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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

type ApiDeleteWebhookKeyRequest struct {
	ctx            context.Context
	ApiService     WebhookKeysApi
	id             string
	acceptLanguage *string
}

// Use for knowing which language to use
func (r ApiDeleteWebhookKeyRequest) AcceptLanguage(acceptLanguage string) ApiDeleteWebhookKeyRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiDeleteWebhookKeyRequest) Execute() (*WebhookKeyDeleteResponse, *http.Response, error) {
	return r.ApiService.DeleteWebhookKeyExecute(r)
}

/*
DeleteWebhookKey Delete Webhook key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiDeleteWebhookKeyRequest
*/
func (a *WebhookKeysApiService) DeleteWebhookKey(ctx context.Context, id string) ApiDeleteWebhookKeyRequest {
	return ApiDeleteWebhookKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return WebhookKeyDeleteResponse
func (a *WebhookKeysApiService) DeleteWebhookKeyExecute(r ApiDeleteWebhookKeyRequest) (*WebhookKeyDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebhookKeyDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookKeysApiService.DeleteWebhookKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook_keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.conekta-v2.1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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

type ApiGetWebhookKeyRequest struct {
	ctx             context.Context
	ApiService      WebhookKeysApi
	id              string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiGetWebhookKeyRequest) AcceptLanguage(acceptLanguage string) ApiGetWebhookKeyRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetWebhookKeyRequest) XChildCompanyId(xChildCompanyId string) ApiGetWebhookKeyRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiGetWebhookKeyRequest) Execute() (*WebhookKeyResponse, *http.Response, error) {
	return r.ApiService.GetWebhookKeyExecute(r)
}

/*
GetWebhookKey Get Webhook Key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiGetWebhookKeyRequest
*/
func (a *WebhookKeysApiService) GetWebhookKey(ctx context.Context, id string) ApiGetWebhookKeyRequest {
	return ApiGetWebhookKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return WebhookKeyResponse
func (a *WebhookKeysApiService) GetWebhookKeyExecute(r ApiGetWebhookKeyRequest) (*WebhookKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebhookKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookKeysApiService.GetWebhookKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook_keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.conekta-v2.1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	if r.xChildCompanyId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Child-Company-Id", r.xChildCompanyId, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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

type ApiGetWebhookKeysRequest struct {
	ctx             context.Context
	ApiService      WebhookKeysApi
	acceptLanguage  *string
	xChildCompanyId *string
	limit           *int32
	search          *string
	next            *string
	previous        *string
}

// Use for knowing which language to use
func (r ApiGetWebhookKeysRequest) AcceptLanguage(acceptLanguage string) ApiGetWebhookKeysRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetWebhookKeysRequest) XChildCompanyId(xChildCompanyId string) ApiGetWebhookKeysRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// The numbers of items to return, the maximum value is 250
func (r ApiGetWebhookKeysRequest) Limit(limit int32) ApiGetWebhookKeysRequest {
	r.limit = &limit
	return r
}

// General order search, e.g. by mail, reference etc.
func (r ApiGetWebhookKeysRequest) Search(search string) ApiGetWebhookKeysRequest {
	r.search = &search
	return r
}

// next page
func (r ApiGetWebhookKeysRequest) Next(next string) ApiGetWebhookKeysRequest {
	r.next = &next
	return r
}

// previous page
func (r ApiGetWebhookKeysRequest) Previous(previous string) ApiGetWebhookKeysRequest {
	r.previous = &previous
	return r
}

func (r ApiGetWebhookKeysRequest) Execute() (*GetWebhookKeysResponse, *http.Response, error) {
	return r.ApiService.GetWebhookKeysExecute(r)
}

/*
GetWebhookKeys Get List of Webhook Keys

Consume the list of webhook keys you have

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetWebhookKeysRequest
*/
func (a *WebhookKeysApiService) GetWebhookKeys(ctx context.Context) ApiGetWebhookKeysRequest {
	return ApiGetWebhookKeysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetWebhookKeysResponse
func (a *WebhookKeysApiService) GetWebhookKeysExecute(r ApiGetWebhookKeysRequest) (*GetWebhookKeysResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetWebhookKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookKeysApiService.GetWebhookKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.next != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "next", r.next, "")
	}
	if r.previous != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "previous", r.previous, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.conekta-v2.1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	if r.xChildCompanyId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Child-Company-Id", r.xChildCompanyId, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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

type ApiUpdateWebhookKeyRequest struct {
	ctx                     context.Context
	ApiService              WebhookKeysApi
	id                      string
	acceptLanguage          *string
	webhookKeyUpdateRequest *WebhookKeyUpdateRequest
}

// Use for knowing which language to use
func (r ApiUpdateWebhookKeyRequest) AcceptLanguage(acceptLanguage string) ApiUpdateWebhookKeyRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiUpdateWebhookKeyRequest) WebhookKeyUpdateRequest(webhookKeyUpdateRequest WebhookKeyUpdateRequest) ApiUpdateWebhookKeyRequest {
	r.webhookKeyUpdateRequest = &webhookKeyUpdateRequest
	return r
}

func (r ApiUpdateWebhookKeyRequest) Execute() (*WebhookKeyResponse, *http.Response, error) {
	return r.ApiService.UpdateWebhookKeyExecute(r)
}

/*
UpdateWebhookKey Update Webhook Key

updates an existing webhook key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiUpdateWebhookKeyRequest
*/
func (a *WebhookKeysApiService) UpdateWebhookKey(ctx context.Context, id string) ApiUpdateWebhookKeyRequest {
	return ApiUpdateWebhookKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return WebhookKeyResponse
func (a *WebhookKeysApiService) UpdateWebhookKeyExecute(r ApiUpdateWebhookKeyRequest) (*WebhookKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebhookKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookKeysApiService.UpdateWebhookKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook_keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.conekta-v2.1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.webhookKeyUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
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

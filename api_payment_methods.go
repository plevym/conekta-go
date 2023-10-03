package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PaymentMethodsApi interface {

	/*
		CreateCustomerPaymentMethods Create Payment Method

		Create a payment method for a customer.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiCreateCustomerPaymentMethodsRequest
	*/
	CreateCustomerPaymentMethods(ctx context.Context, id string) ApiCreateCustomerPaymentMethodsRequest

	// CreateCustomerPaymentMethodsExecute executes the request
	//  @return CreateCustomerPaymentMethodsResponse
	CreateCustomerPaymentMethodsExecute(r ApiCreateCustomerPaymentMethodsRequest) (*CreateCustomerPaymentMethodsResponse, *http.Response, error)

	/*
		DeleteCustomerPaymentMethods Delete Payment Method

		Delete an existing payment method

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param paymentMethodId Identifier of the payment method
		@return ApiDeleteCustomerPaymentMethodsRequest
	*/
	DeleteCustomerPaymentMethods(ctx context.Context, id string, paymentMethodId string) ApiDeleteCustomerPaymentMethodsRequest

	// DeleteCustomerPaymentMethodsExecute executes the request
	//  @return UpdateCustomerPaymentMethodsResponse
	DeleteCustomerPaymentMethodsExecute(r ApiDeleteCustomerPaymentMethodsRequest) (*UpdateCustomerPaymentMethodsResponse, *http.Response, error)

	/*
		GetCustomerPaymentMethods Get Payment Methods

		Get a list of Payment Methods

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiGetCustomerPaymentMethodsRequest
	*/
	GetCustomerPaymentMethods(ctx context.Context, id string) ApiGetCustomerPaymentMethodsRequest

	// GetCustomerPaymentMethodsExecute executes the request
	//  @return GetPaymentMethodResponse
	GetCustomerPaymentMethodsExecute(r ApiGetCustomerPaymentMethodsRequest) (*GetPaymentMethodResponse, *http.Response, error)

	/*
		UpdateCustomerPaymentMethods Update Payment Method

		Gets a payment Method that corresponds to a customer ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param paymentMethodId Identifier of the payment method
		@return ApiUpdateCustomerPaymentMethodsRequest
	*/
	UpdateCustomerPaymentMethods(ctx context.Context, id string, paymentMethodId string) ApiUpdateCustomerPaymentMethodsRequest

	// UpdateCustomerPaymentMethodsExecute executes the request
	//  @return UpdateCustomerPaymentMethodsResponse
	UpdateCustomerPaymentMethodsExecute(r ApiUpdateCustomerPaymentMethodsRequest) (*UpdateCustomerPaymentMethodsResponse, *http.Response, error)
}

// PaymentMethodsApiService PaymentMethodsApi service
type PaymentMethodsApiService service

type ApiCreateCustomerPaymentMethodsRequest struct {
	ctx                                 context.Context
	ApiService                          PaymentMethodsApi
	id                                  string
	createCustomerPaymentMethodsRequest *CreateCustomerPaymentMethodsRequest
	acceptLanguage                      *string
	xChildCompanyId                     *string
}

// requested field for customer payment methods
func (r ApiCreateCustomerPaymentMethodsRequest) CreateCustomerPaymentMethodsRequest(createCustomerPaymentMethodsRequest CreateCustomerPaymentMethodsRequest) ApiCreateCustomerPaymentMethodsRequest {
	r.createCustomerPaymentMethodsRequest = &createCustomerPaymentMethodsRequest
	return r
}

// Use for knowing which language to use
func (r ApiCreateCustomerPaymentMethodsRequest) AcceptLanguage(acceptLanguage string) ApiCreateCustomerPaymentMethodsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiCreateCustomerPaymentMethodsRequest) XChildCompanyId(xChildCompanyId string) ApiCreateCustomerPaymentMethodsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiCreateCustomerPaymentMethodsRequest) Execute() (*CreateCustomerPaymentMethodsResponse, *http.Response, error) {
	return r.ApiService.CreateCustomerPaymentMethodsExecute(r)
}

/*
CreateCustomerPaymentMethods Create Payment Method

Create a payment method for a customer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiCreateCustomerPaymentMethodsRequest
*/
func (a *PaymentMethodsApiService) CreateCustomerPaymentMethods(ctx context.Context, id string) ApiCreateCustomerPaymentMethodsRequest {
	return ApiCreateCustomerPaymentMethodsRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CreateCustomerPaymentMethodsResponse
func (a *PaymentMethodsApiService) CreateCustomerPaymentMethodsExecute(r ApiCreateCustomerPaymentMethodsRequest) (*CreateCustomerPaymentMethodsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateCustomerPaymentMethodsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.CreateCustomerPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/payment_sources"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCustomerPaymentMethodsRequest == nil {
		return localVarReturnValue, nil, reportError("createCustomerPaymentMethodsRequest is required and must be specified")
	}

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
	if r.xChildCompanyId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Child-Company-Id", r.xChildCompanyId, "")
	}
	// body params
	localVarPostBody = r.createCustomerPaymentMethodsRequest
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
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiDeleteCustomerPaymentMethodsRequest struct {
	ctx             context.Context
	ApiService      PaymentMethodsApi
	id              string
	paymentMethodId string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiDeleteCustomerPaymentMethodsRequest) AcceptLanguage(acceptLanguage string) ApiDeleteCustomerPaymentMethodsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiDeleteCustomerPaymentMethodsRequest) XChildCompanyId(xChildCompanyId string) ApiDeleteCustomerPaymentMethodsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiDeleteCustomerPaymentMethodsRequest) Execute() (*UpdateCustomerPaymentMethodsResponse, *http.Response, error) {
	return r.ApiService.DeleteCustomerPaymentMethodsExecute(r)
}

/*
DeleteCustomerPaymentMethods Delete Payment Method

Delete an existing payment method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param paymentMethodId Identifier of the payment method
	@return ApiDeleteCustomerPaymentMethodsRequest
*/
func (a *PaymentMethodsApiService) DeleteCustomerPaymentMethods(ctx context.Context, id string, paymentMethodId string) ApiDeleteCustomerPaymentMethodsRequest {
	return ApiDeleteCustomerPaymentMethodsRequest{
		ApiService:      a,
		ctx:             ctx,
		id:              id,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//
//	@return UpdateCustomerPaymentMethodsResponse
func (a *PaymentMethodsApiService) DeleteCustomerPaymentMethodsExecute(r ApiDeleteCustomerPaymentMethodsRequest) (*UpdateCustomerPaymentMethodsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateCustomerPaymentMethodsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.DeleteCustomerPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/payment_sources/{payment_method_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"payment_method_id"+"}", url.PathEscape(parameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

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
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiGetCustomerPaymentMethodsRequest struct {
	ctx             context.Context
	ApiService      PaymentMethodsApi
	id              string
	acceptLanguage  *string
	xChildCompanyId *string
	limit           *int32
	next            *string
	previous        *string
	search          *string
}

// Use for knowing which language to use
func (r ApiGetCustomerPaymentMethodsRequest) AcceptLanguage(acceptLanguage string) ApiGetCustomerPaymentMethodsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetCustomerPaymentMethodsRequest) XChildCompanyId(xChildCompanyId string) ApiGetCustomerPaymentMethodsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// The numbers of items to return, the maximum value is 250
func (r ApiGetCustomerPaymentMethodsRequest) Limit(limit int32) ApiGetCustomerPaymentMethodsRequest {
	r.limit = &limit
	return r
}

// next page
func (r ApiGetCustomerPaymentMethodsRequest) Next(next string) ApiGetCustomerPaymentMethodsRequest {
	r.next = &next
	return r
}

// previous page
func (r ApiGetCustomerPaymentMethodsRequest) Previous(previous string) ApiGetCustomerPaymentMethodsRequest {
	r.previous = &previous
	return r
}

// General order search, e.g. by mail, reference etc.
func (r ApiGetCustomerPaymentMethodsRequest) Search(search string) ApiGetCustomerPaymentMethodsRequest {
	r.search = &search
	return r
}

func (r ApiGetCustomerPaymentMethodsRequest) Execute() (*GetPaymentMethodResponse, *http.Response, error) {
	return r.ApiService.GetCustomerPaymentMethodsExecute(r)
}

/*
GetCustomerPaymentMethods Get Payment Methods

Get a list of Payment Methods

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiGetCustomerPaymentMethodsRequest
*/
func (a *PaymentMethodsApiService) GetCustomerPaymentMethods(ctx context.Context, id string) ApiGetCustomerPaymentMethodsRequest {
	return ApiGetCustomerPaymentMethodsRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetPaymentMethodResponse
func (a *PaymentMethodsApiService) GetCustomerPaymentMethodsExecute(r ApiGetCustomerPaymentMethodsRequest) (*GetPaymentMethodResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPaymentMethodResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.GetCustomerPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/payment_sources"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.next != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "next", r.next, "")
	}
	if r.previous != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "previous", r.previous, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
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

type ApiUpdateCustomerPaymentMethodsRequest struct {
	ctx                  context.Context
	ApiService           PaymentMethodsApi
	id                   string
	paymentMethodId      string
	updatePaymentMethods *UpdatePaymentMethods
	acceptLanguage       *string
	xChildCompanyId      *string
}

// requested field for customer payment methods
func (r ApiUpdateCustomerPaymentMethodsRequest) UpdatePaymentMethods(updatePaymentMethods UpdatePaymentMethods) ApiUpdateCustomerPaymentMethodsRequest {
	r.updatePaymentMethods = &updatePaymentMethods
	return r
}

// Use for knowing which language to use
func (r ApiUpdateCustomerPaymentMethodsRequest) AcceptLanguage(acceptLanguage string) ApiUpdateCustomerPaymentMethodsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiUpdateCustomerPaymentMethodsRequest) XChildCompanyId(xChildCompanyId string) ApiUpdateCustomerPaymentMethodsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiUpdateCustomerPaymentMethodsRequest) Execute() (*UpdateCustomerPaymentMethodsResponse, *http.Response, error) {
	return r.ApiService.UpdateCustomerPaymentMethodsExecute(r)
}

/*
UpdateCustomerPaymentMethods Update Payment Method

Gets a payment Method that corresponds to a customer ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param paymentMethodId Identifier of the payment method
	@return ApiUpdateCustomerPaymentMethodsRequest
*/
func (a *PaymentMethodsApiService) UpdateCustomerPaymentMethods(ctx context.Context, id string, paymentMethodId string) ApiUpdateCustomerPaymentMethodsRequest {
	return ApiUpdateCustomerPaymentMethodsRequest{
		ApiService:      a,
		ctx:             ctx,
		id:              id,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//
//	@return UpdateCustomerPaymentMethodsResponse
func (a *PaymentMethodsApiService) UpdateCustomerPaymentMethodsExecute(r ApiUpdateCustomerPaymentMethodsRequest) (*UpdateCustomerPaymentMethodsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateCustomerPaymentMethodsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.UpdateCustomerPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/payment_sources/{payment_method_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"payment_method_id"+"}", url.PathEscape(parameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updatePaymentMethods == nil {
		return localVarReturnValue, nil, reportError("updatePaymentMethods is required and must be specified")
	}

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
	if r.xChildCompanyId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Child-Company-Id", r.xChildCompanyId, "")
	}
	// body params
	localVarPostBody = r.updatePaymentMethods
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
		if localVarHTTPResponse.StatusCode == 422 {
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

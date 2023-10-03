package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TaxesApi interface {

	/*
		OrdersCreateTaxes Create Tax

		Create new taxes for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrdersCreateTaxesRequest
	*/
	OrdersCreateTaxes(ctx context.Context, id string) ApiOrdersCreateTaxesRequest

	// OrdersCreateTaxesExecute executes the request
	//  @return UpdateOrderTaxResponse
	OrdersCreateTaxesExecute(r ApiOrdersCreateTaxesRequest) (*UpdateOrderTaxResponse, *http.Response, error)

	/*
		OrdersDeleteTaxes Delete Tax

		Delete taxes for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param taxId identifier
		@return ApiOrdersDeleteTaxesRequest
	*/
	OrdersDeleteTaxes(ctx context.Context, id string, taxId string) ApiOrdersDeleteTaxesRequest

	// OrdersDeleteTaxesExecute executes the request
	//  @return UpdateOrderTaxResponse
	OrdersDeleteTaxesExecute(r ApiOrdersDeleteTaxesRequest) (*UpdateOrderTaxResponse, *http.Response, error)

	/*
		OrdersUpdateTaxes Update Tax

		Update taxes for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param taxId identifier
		@return ApiOrdersUpdateTaxesRequest
	*/
	OrdersUpdateTaxes(ctx context.Context, id string, taxId string) ApiOrdersUpdateTaxesRequest

	// OrdersUpdateTaxesExecute executes the request
	//  @return UpdateOrderTaxResponse
	OrdersUpdateTaxesExecute(r ApiOrdersUpdateTaxesRequest) (*UpdateOrderTaxResponse, *http.Response, error)
}

// TaxesApiService TaxesApi service
type TaxesApiService service

type ApiOrdersCreateTaxesRequest struct {
	ctx             context.Context
	ApiService      TaxesApi
	id              string
	orderTaxRequest *OrderTaxRequest
	acceptLanguage  *string
	xChildCompanyId *string
}

// requested field for a taxes
func (r ApiOrdersCreateTaxesRequest) OrderTaxRequest(orderTaxRequest OrderTaxRequest) ApiOrdersCreateTaxesRequest {
	r.orderTaxRequest = &orderTaxRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersCreateTaxesRequest) AcceptLanguage(acceptLanguage string) ApiOrdersCreateTaxesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersCreateTaxesRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersCreateTaxesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersCreateTaxesRequest) Execute() (*UpdateOrderTaxResponse, *http.Response, error) {
	return r.ApiService.OrdersCreateTaxesExecute(r)
}

/*
OrdersCreateTaxes Create Tax

Create new taxes for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrdersCreateTaxesRequest
*/
func (a *TaxesApiService) OrdersCreateTaxes(ctx context.Context, id string) ApiOrdersCreateTaxesRequest {
	return ApiOrdersCreateTaxesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UpdateOrderTaxResponse
func (a *TaxesApiService) OrdersCreateTaxesExecute(r ApiOrdersCreateTaxesRequest) (*UpdateOrderTaxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateOrderTaxResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.OrdersCreateTaxes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/tax_lines"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderTaxRequest == nil {
		return localVarReturnValue, nil, reportError("orderTaxRequest is required and must be specified")
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
	localVarPostBody = r.orderTaxRequest
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

type ApiOrdersDeleteTaxesRequest struct {
	ctx             context.Context
	ApiService      TaxesApi
	id              string
	taxId           string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiOrdersDeleteTaxesRequest) AcceptLanguage(acceptLanguage string) ApiOrdersDeleteTaxesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersDeleteTaxesRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersDeleteTaxesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersDeleteTaxesRequest) Execute() (*UpdateOrderTaxResponse, *http.Response, error) {
	return r.ApiService.OrdersDeleteTaxesExecute(r)
}

/*
OrdersDeleteTaxes Delete Tax

Delete taxes for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param taxId identifier
	@return ApiOrdersDeleteTaxesRequest
*/
func (a *TaxesApiService) OrdersDeleteTaxes(ctx context.Context, id string, taxId string) ApiOrdersDeleteTaxesRequest {
	return ApiOrdersDeleteTaxesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		taxId:      taxId,
	}
}

// Execute executes the request
//
//	@return UpdateOrderTaxResponse
func (a *TaxesApiService) OrdersDeleteTaxesExecute(r ApiOrdersDeleteTaxesRequest) (*UpdateOrderTaxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateOrderTaxResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.OrdersDeleteTaxes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/tax_lines/{tax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tax_id"+"}", url.PathEscape(parameterValueToString(r.taxId, "taxId")), -1)

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

type ApiOrdersUpdateTaxesRequest struct {
	ctx                   context.Context
	ApiService            TaxesApi
	id                    string
	taxId                 string
	updateOrderTaxRequest *UpdateOrderTaxRequest
	acceptLanguage        *string
	xChildCompanyId       *string
}

// requested field for taxes
func (r ApiOrdersUpdateTaxesRequest) UpdateOrderTaxRequest(updateOrderTaxRequest UpdateOrderTaxRequest) ApiOrdersUpdateTaxesRequest {
	r.updateOrderTaxRequest = &updateOrderTaxRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersUpdateTaxesRequest) AcceptLanguage(acceptLanguage string) ApiOrdersUpdateTaxesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersUpdateTaxesRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersUpdateTaxesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersUpdateTaxesRequest) Execute() (*UpdateOrderTaxResponse, *http.Response, error) {
	return r.ApiService.OrdersUpdateTaxesExecute(r)
}

/*
OrdersUpdateTaxes Update Tax

Update taxes for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param taxId identifier
	@return ApiOrdersUpdateTaxesRequest
*/
func (a *TaxesApiService) OrdersUpdateTaxes(ctx context.Context, id string, taxId string) ApiOrdersUpdateTaxesRequest {
	return ApiOrdersUpdateTaxesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		taxId:      taxId,
	}
}

// Execute executes the request
//
//	@return UpdateOrderTaxResponse
func (a *TaxesApiService) OrdersUpdateTaxesExecute(r ApiOrdersUpdateTaxesRequest) (*UpdateOrderTaxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateOrderTaxResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.OrdersUpdateTaxes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/tax_lines/{tax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tax_id"+"}", url.PathEscape(parameterValueToString(r.taxId, "taxId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrderTaxRequest == nil {
		return localVarReturnValue, nil, reportError("updateOrderTaxRequest is required and must be specified")
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
	localVarPostBody = r.updateOrderTaxRequest
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

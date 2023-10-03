package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ShippingsApi interface {

	/*
		OrdersCreateShipping Create Shipping

		Create new shipping for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrdersCreateShippingRequest
	*/
	OrdersCreateShipping(ctx context.Context, id string) ApiOrdersCreateShippingRequest

	// OrdersCreateShippingExecute executes the request
	//  @return ShippingOrderResponse
	OrdersCreateShippingExecute(r ApiOrdersCreateShippingRequest) (*ShippingOrderResponse, *http.Response, error)

	/*
		OrdersDeleteShipping Delete Shipping

		Delete shipping

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param shippingId identifier
		@return ApiOrdersDeleteShippingRequest
	*/
	OrdersDeleteShipping(ctx context.Context, id string, shippingId string) ApiOrdersDeleteShippingRequest

	// OrdersDeleteShippingExecute executes the request
	//  @return ShippingOrderResponse
	OrdersDeleteShippingExecute(r ApiOrdersDeleteShippingRequest) (*ShippingOrderResponse, *http.Response, error)

	/*
		OrdersUpdateShipping Update Shipping

		Update existing shipping for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param shippingId identifier
		@return ApiOrdersUpdateShippingRequest
	*/
	OrdersUpdateShipping(ctx context.Context, id string, shippingId string) ApiOrdersUpdateShippingRequest

	// OrdersUpdateShippingExecute executes the request
	//  @return ShippingOrderResponse
	OrdersUpdateShippingExecute(r ApiOrdersUpdateShippingRequest) (*ShippingOrderResponse, *http.Response, error)
}

// ShippingsApiService ShippingsApi service
type ShippingsApiService service

type ApiOrdersCreateShippingRequest struct {
	ctx             context.Context
	ApiService      ShippingsApi
	id              string
	shippingRequest *ShippingRequest
	acceptLanguage  *string
	xChildCompanyId *string
}

// requested field for a shipping
func (r ApiOrdersCreateShippingRequest) ShippingRequest(shippingRequest ShippingRequest) ApiOrdersCreateShippingRequest {
	r.shippingRequest = &shippingRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersCreateShippingRequest) AcceptLanguage(acceptLanguage string) ApiOrdersCreateShippingRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersCreateShippingRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersCreateShippingRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersCreateShippingRequest) Execute() (*ShippingOrderResponse, *http.Response, error) {
	return r.ApiService.OrdersCreateShippingExecute(r)
}

/*
OrdersCreateShipping Create Shipping

Create new shipping for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrdersCreateShippingRequest
*/
func (a *ShippingsApiService) OrdersCreateShipping(ctx context.Context, id string) ApiOrdersCreateShippingRequest {
	return ApiOrdersCreateShippingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ShippingOrderResponse
func (a *ShippingsApiService) OrdersCreateShippingExecute(r ApiOrdersCreateShippingRequest) (*ShippingOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingsApiService.OrdersCreateShipping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/shipping_lines"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingRequest == nil {
		return localVarReturnValue, nil, reportError("shippingRequest is required and must be specified")
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
	localVarPostBody = r.shippingRequest
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

type ApiOrdersDeleteShippingRequest struct {
	ctx             context.Context
	ApiService      ShippingsApi
	id              string
	shippingId      string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiOrdersDeleteShippingRequest) AcceptLanguage(acceptLanguage string) ApiOrdersDeleteShippingRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersDeleteShippingRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersDeleteShippingRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersDeleteShippingRequest) Execute() (*ShippingOrderResponse, *http.Response, error) {
	return r.ApiService.OrdersDeleteShippingExecute(r)
}

/*
OrdersDeleteShipping Delete Shipping

Delete shipping

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param shippingId identifier
	@return ApiOrdersDeleteShippingRequest
*/
func (a *ShippingsApiService) OrdersDeleteShipping(ctx context.Context, id string, shippingId string) ApiOrdersDeleteShippingRequest {
	return ApiOrdersDeleteShippingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		shippingId: shippingId,
	}
}

// Execute executes the request
//
//	@return ShippingOrderResponse
func (a *ShippingsApiService) OrdersDeleteShippingExecute(r ApiOrdersDeleteShippingRequest) (*ShippingOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingsApiService.OrdersDeleteShipping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/shipping_lines/{shipping_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shipping_id"+"}", url.PathEscape(parameterValueToString(r.shippingId, "shippingId")), -1)

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
		if localVarHTTPResponse.StatusCode == 428 {
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

type ApiOrdersUpdateShippingRequest struct {
	ctx             context.Context
	ApiService      ShippingsApi
	id              string
	shippingId      string
	shippingRequest *ShippingRequest
	acceptLanguage  *string
	xChildCompanyId *string
}

// requested field for a shipping
func (r ApiOrdersUpdateShippingRequest) ShippingRequest(shippingRequest ShippingRequest) ApiOrdersUpdateShippingRequest {
	r.shippingRequest = &shippingRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersUpdateShippingRequest) AcceptLanguage(acceptLanguage string) ApiOrdersUpdateShippingRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersUpdateShippingRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersUpdateShippingRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersUpdateShippingRequest) Execute() (*ShippingOrderResponse, *http.Response, error) {
	return r.ApiService.OrdersUpdateShippingExecute(r)
}

/*
OrdersUpdateShipping Update Shipping

Update existing shipping for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param shippingId identifier
	@return ApiOrdersUpdateShippingRequest
*/
func (a *ShippingsApiService) OrdersUpdateShipping(ctx context.Context, id string, shippingId string) ApiOrdersUpdateShippingRequest {
	return ApiOrdersUpdateShippingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		shippingId: shippingId,
	}
}

// Execute executes the request
//
//	@return ShippingOrderResponse
func (a *ShippingsApiService) OrdersUpdateShippingExecute(r ApiOrdersUpdateShippingRequest) (*ShippingOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingsApiService.OrdersUpdateShipping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/shipping_lines/{shipping_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shipping_id"+"}", url.PathEscape(parameterValueToString(r.shippingId, "shippingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingRequest == nil {
		return localVarReturnValue, nil, reportError("shippingRequest is required and must be specified")
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
	localVarPostBody = r.shippingRequest
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

package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type OrdersApi interface {

	/*
		CancelOrder Cancel Order

		Cancel an order that has been previously created.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiCancelOrderRequest
	*/
	CancelOrder(ctx context.Context, id string) ApiCancelOrderRequest

	// CancelOrderExecute executes the request
	//  @return OrderResponse
	CancelOrderExecute(r ApiCancelOrderRequest) (*OrderResponse, *http.Response, error)

	/*
		CreateOrder Create order

		Create a new order.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateOrderRequest
	*/
	CreateOrder(ctx context.Context) ApiCreateOrderRequest

	// CreateOrderExecute executes the request
	//  @return OrderResponse
	CreateOrderExecute(r ApiCreateOrderRequest) (*OrderResponse, *http.Response, error)

	/*
		GetOrderById Get Order

		Info for a specific order

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiGetOrderByIdRequest
	*/
	GetOrderById(ctx context.Context, id string) ApiGetOrderByIdRequest

	// GetOrderByIdExecute executes the request
	//  @return OrderResponse
	GetOrderByIdExecute(r ApiGetOrderByIdRequest) (*OrderResponse, *http.Response, error)

	/*
		GetOrders Get a list of Orders

		Get order details in the form of a list

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetOrdersRequest
	*/
	GetOrders(ctx context.Context) ApiGetOrdersRequest

	// GetOrdersExecute executes the request
	//  @return GetOrdersResponse
	GetOrdersExecute(r ApiGetOrdersRequest) (*GetOrdersResponse, *http.Response, error)

	/*
		OrderCancelRefund Cancel Refund

		A refunded order describes the items, amount, and reason an order is being refunded.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param refundId refund identifier
		@return ApiOrderCancelRefundRequest
	*/
	OrderCancelRefund(ctx context.Context, id string, refundId string) ApiOrderCancelRefundRequest

	// OrderCancelRefundExecute executes the request
	//  @return OrderResponse
	OrderCancelRefundExecute(r ApiOrderCancelRefundRequest) (*OrderResponse, *http.Response, error)

	/*
		OrderRefund Refund Order

		A refunded order describes the items, amount, and reason an order is being refunded.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrderRefundRequest
	*/
	OrderRefund(ctx context.Context, id string) ApiOrderRefundRequest

	// OrderRefundExecute executes the request
	//  @return OrderResponse
	OrderRefundExecute(r ApiOrderRefundRequest) (*OrderResponse, *http.Response, error)

	/*
		OrdersCreateCapture Capture Order

		Processes an order that has been previously authorized.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrdersCreateCaptureRequest
	*/
	OrdersCreateCapture(ctx context.Context, id string) ApiOrdersCreateCaptureRequest

	// OrdersCreateCaptureExecute executes the request
	//  @return OrderResponse
	OrdersCreateCaptureExecute(r ApiOrdersCreateCaptureRequest) (*OrderResponse, *http.Response, error)

	/*
		UpdateOrder Update Order

		Update an existing Order.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiUpdateOrderRequest
	*/
	UpdateOrder(ctx context.Context, id string) ApiUpdateOrderRequest

	// UpdateOrderExecute executes the request
	//  @return OrderResponse
	UpdateOrderExecute(r ApiUpdateOrderRequest) (*OrderResponse, *http.Response, error)
}

// OrdersApiService OrdersApi service
type OrdersApiService service

type ApiCancelOrderRequest struct {
	ctx             context.Context
	ApiService      OrdersApi
	id              string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiCancelOrderRequest) AcceptLanguage(acceptLanguage string) ApiCancelOrderRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiCancelOrderRequest) XChildCompanyId(xChildCompanyId string) ApiCancelOrderRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiCancelOrderRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Cancel Order

Cancel an order that has been previously created.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiCancelOrderRequest
*/
func (a *OrdersApiService) CancelOrder(ctx context.Context, id string) ApiCancelOrderRequest {
	return ApiCancelOrderRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) CancelOrderExecute(r ApiCancelOrderRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.CancelOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/cancel"
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
		if localVarHTTPResponse.StatusCode == 402 {
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

type ApiCreateOrderRequest struct {
	ctx             context.Context
	ApiService      OrdersApi
	orderRequest    *OrderRequest
	acceptLanguage  *string
	xChildCompanyId *string
}

// requested field for order
func (r ApiCreateOrderRequest) OrderRequest(orderRequest OrderRequest) ApiCreateOrderRequest {
	r.orderRequest = &orderRequest
	return r
}

// Use for knowing which language to use
func (r ApiCreateOrderRequest) AcceptLanguage(acceptLanguage string) ApiCreateOrderRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiCreateOrderRequest) XChildCompanyId(xChildCompanyId string) ApiCreateOrderRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiCreateOrderRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.CreateOrderExecute(r)
}

/*
CreateOrder Create order

Create a new order.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateOrderRequest
*/
func (a *OrdersApiService) CreateOrder(ctx context.Context) ApiCreateOrderRequest {
	return ApiCreateOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) CreateOrderExecute(r ApiCreateOrderRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.CreateOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderRequest == nil {
		return localVarReturnValue, nil, reportError("orderRequest is required and must be specified")
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
	localVarPostBody = r.orderRequest
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
		if localVarHTTPResponse.StatusCode == 402 {
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

type ApiGetOrderByIdRequest struct {
	ctx             context.Context
	ApiService      OrdersApi
	id              string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiGetOrderByIdRequest) AcceptLanguage(acceptLanguage string) ApiGetOrderByIdRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetOrderByIdRequest) XChildCompanyId(xChildCompanyId string) ApiGetOrderByIdRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiGetOrderByIdRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.GetOrderByIdExecute(r)
}

/*
GetOrderById Get Order

Info for a specific order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiGetOrderByIdRequest
*/
func (a *OrdersApiService) GetOrderById(ctx context.Context, id string) ApiGetOrderByIdRequest {
	return ApiGetOrderByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) GetOrderByIdExecute(r ApiGetOrderByIdRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.GetOrderById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}"
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

type ApiGetOrdersRequest struct {
	ctx             context.Context
	ApiService      OrdersApi
	acceptLanguage  *string
	xChildCompanyId *string
	limit           *int32
	search          *string
	next            *string
	previous        *string
	paymentStatus   *string
}

// Use for knowing which language to use
func (r ApiGetOrdersRequest) AcceptLanguage(acceptLanguage string) ApiGetOrdersRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetOrdersRequest) XChildCompanyId(xChildCompanyId string) ApiGetOrdersRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// The numbers of items to return, the maximum value is 250
func (r ApiGetOrdersRequest) Limit(limit int32) ApiGetOrdersRequest {
	r.limit = &limit
	return r
}

// General order search, e.g. by mail, reference etc.
func (r ApiGetOrdersRequest) Search(search string) ApiGetOrdersRequest {
	r.search = &search
	return r
}

func (r ApiGetOrdersRequest) PaymentStatus(search string) ApiGetOrdersRequest {
	r.paymentStatus = &search
	return r
}

// next page
func (r ApiGetOrdersRequest) Next(next string) ApiGetOrdersRequest {
	r.next = &next
	return r
}

// previous page
func (r ApiGetOrdersRequest) Previous(previous string) ApiGetOrdersRequest {
	r.previous = &previous
	return r
}

func (r ApiGetOrdersRequest) Execute() (*GetOrdersResponse, *http.Response, error) {
	return r.ApiService.GetOrdersExecute(r)
}

/*
GetOrders Get a list of Orders

Get order details in the form of a list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOrdersRequest
*/
func (a *OrdersApiService) GetOrders(ctx context.Context) ApiGetOrdersRequest {
	return ApiGetOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOrdersResponse
func (a *OrdersApiService) GetOrdersExecute(r ApiGetOrdersRequest) (*GetOrdersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrdersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.GetOrders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders"

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
	if r.paymentStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payment_status", r.paymentStatus, "")
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

type ApiOrderCancelRefundRequest struct {
	ctx             context.Context
	ApiService      OrdersApi
	id              string
	refundId        string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiOrderCancelRefundRequest) AcceptLanguage(acceptLanguage string) ApiOrderCancelRefundRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrderCancelRefundRequest) XChildCompanyId(xChildCompanyId string) ApiOrderCancelRefundRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrderCancelRefundRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.OrderCancelRefundExecute(r)
}

/*
OrderCancelRefund Cancel Refund

A refunded order describes the items, amount, and reason an order is being refunded.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param refundId refund identifier
	@return ApiOrderCancelRefundRequest
*/
func (a *OrdersApiService) OrderCancelRefund(ctx context.Context, id string, refundId string) ApiOrderCancelRefundRequest {
	return ApiOrderCancelRefundRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		refundId:   refundId,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) OrderCancelRefundExecute(r ApiOrderCancelRefundRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.OrderCancelRefund")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/refunds/{refund_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"refund_id"+"}", url.PathEscape(parameterValueToString(r.refundId, "refundId")), -1)

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
		if localVarHTTPResponse.StatusCode == 402 {
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

type ApiOrderRefundRequest struct {
	ctx                context.Context
	ApiService         OrdersApi
	id                 string
	orderRefundRequest *OrderRefundRequest
	acceptLanguage     *string
	xChildCompanyId    *string
}

// requested field for a refund
func (r ApiOrderRefundRequest) OrderRefundRequest(orderRefundRequest OrderRefundRequest) ApiOrderRefundRequest {
	r.orderRefundRequest = &orderRefundRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrderRefundRequest) AcceptLanguage(acceptLanguage string) ApiOrderRefundRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrderRefundRequest) XChildCompanyId(xChildCompanyId string) ApiOrderRefundRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrderRefundRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.OrderRefundExecute(r)
}

/*
OrderRefund Refund Order

A refunded order describes the items, amount, and reason an order is being refunded.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrderRefundRequest
*/
func (a *OrdersApiService) OrderRefund(ctx context.Context, id string) ApiOrderRefundRequest {
	return ApiOrderRefundRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) OrderRefundExecute(r ApiOrderRefundRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.OrderRefund")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/refunds"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderRefundRequest == nil {
		return localVarReturnValue, nil, reportError("orderRefundRequest is required and must be specified")
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
	localVarPostBody = r.orderRefundRequest
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
		if localVarHTTPResponse.StatusCode == 402 {
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

type ApiOrdersCreateCaptureRequest struct {
	ctx                 context.Context
	ApiService          OrdersApi
	id                  string
	acceptLanguage      *string
	xChildCompanyId     *string
	orderCaptureRequest *OrderCaptureRequest
}

// Use for knowing which language to use
func (r ApiOrdersCreateCaptureRequest) AcceptLanguage(acceptLanguage string) ApiOrdersCreateCaptureRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersCreateCaptureRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersCreateCaptureRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// requested fields for capture order
func (r ApiOrdersCreateCaptureRequest) OrderCaptureRequest(orderCaptureRequest OrderCaptureRequest) ApiOrdersCreateCaptureRequest {
	r.orderCaptureRequest = &orderCaptureRequest
	return r
}

func (r ApiOrdersCreateCaptureRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.OrdersCreateCaptureExecute(r)
}

/*
OrdersCreateCapture Capture Order

Processes an order that has been previously authorized.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrdersCreateCaptureRequest
*/
func (a *OrdersApiService) OrdersCreateCapture(ctx context.Context, id string) ApiOrdersCreateCaptureRequest {
	return ApiOrdersCreateCaptureRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) OrdersCreateCaptureExecute(r ApiOrdersCreateCaptureRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.OrdersCreateCapture")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/capture"
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
	if r.xChildCompanyId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Child-Company-Id", r.xChildCompanyId, "")
	}
	// body params
	localVarPostBody = r.orderCaptureRequest
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

type ApiUpdateOrderRequest struct {
	ctx                context.Context
	ApiService         OrdersApi
	id                 string
	orderUpdateRequest *OrderUpdateRequest
	acceptLanguage     *string
}

// requested field for an order
func (r ApiUpdateOrderRequest) OrderUpdateRequest(orderUpdateRequest OrderUpdateRequest) ApiUpdateOrderRequest {
	r.orderUpdateRequest = &orderUpdateRequest
	return r
}

// Use for knowing which language to use
func (r ApiUpdateOrderRequest) AcceptLanguage(acceptLanguage string) ApiUpdateOrderRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiUpdateOrderRequest) Execute() (*OrderResponse, *http.Response, error) {
	return r.ApiService.UpdateOrderExecute(r)
}

/*
UpdateOrder Update Order

Update an existing Order.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiUpdateOrderRequest
*/
func (a *OrdersApiService) UpdateOrder(ctx context.Context, id string) ApiUpdateOrderRequest {
	return ApiUpdateOrderRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return OrderResponse
func (a *OrdersApiService) UpdateOrderExecute(r ApiUpdateOrderRequest) (*OrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.UpdateOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("orderUpdateRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.orderUpdateRequest
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

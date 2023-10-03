package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DiscountsApi interface {

	/*
		OrdersCreateDiscountLine Create Discount

		Create discount lines for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrdersCreateDiscountLineRequest
	*/
	OrdersCreateDiscountLine(ctx context.Context, id string) ApiOrdersCreateDiscountLineRequest

	// OrdersCreateDiscountLineExecute executes the request
	//  @return DiscountLinesResponse
	OrdersCreateDiscountLineExecute(r ApiOrdersCreateDiscountLineRequest) (*DiscountLinesResponse, *http.Response, error)

	/*
		OrdersDeleteDiscountLines Delete Discount

		Delete an existing discount lines for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param discountLinesId discount line id identifier
		@return ApiOrdersDeleteDiscountLinesRequest
	*/
	OrdersDeleteDiscountLines(ctx context.Context, id string, discountLinesId string) ApiOrdersDeleteDiscountLinesRequest

	// OrdersDeleteDiscountLinesExecute executes the request
	//  @return DiscountLinesResponse
	OrdersDeleteDiscountLinesExecute(r ApiOrdersDeleteDiscountLinesRequest) (*DiscountLinesResponse, *http.Response, error)

	/*
		OrdersGetDiscountLine Get Discount

		Get an existing discount lines for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param discountLinesId discount line id identifier
		@return ApiOrdersGetDiscountLineRequest
	*/
	OrdersGetDiscountLine(ctx context.Context, id string, discountLinesId string) ApiOrdersGetDiscountLineRequest

	// OrdersGetDiscountLineExecute executes the request
	//  @return DiscountLinesResponse
	OrdersGetDiscountLineExecute(r ApiOrdersGetDiscountLineRequest) (*DiscountLinesResponse, *http.Response, error)

	/*
		OrdersGetDiscountLines Get a List of Discount

		Get discount lines for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrdersGetDiscountLinesRequest
	*/
	OrdersGetDiscountLines(ctx context.Context, id string) ApiOrdersGetDiscountLinesRequest

	// OrdersGetDiscountLinesExecute executes the request
	//  @return GetOrderDiscountLinesResponse
	OrdersGetDiscountLinesExecute(r ApiOrdersGetDiscountLinesRequest) (*GetOrderDiscountLinesResponse, *http.Response, error)

	/*
		OrdersUpdateDiscountLines Update Discount

		Update an existing discount lines for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param discountLinesId discount line id identifier
		@return ApiOrdersUpdateDiscountLinesRequest
	*/
	OrdersUpdateDiscountLines(ctx context.Context, id string, discountLinesId string) ApiOrdersUpdateDiscountLinesRequest

	// OrdersUpdateDiscountLinesExecute executes the request
	//  @return DiscountLinesResponse
	OrdersUpdateDiscountLinesExecute(r ApiOrdersUpdateDiscountLinesRequest) (*DiscountLinesResponse, *http.Response, error)
}

// DiscountsApiService DiscountsApi service
type DiscountsApiService service

type ApiOrdersCreateDiscountLineRequest struct {
	ctx                       context.Context
	ApiService                DiscountsApi
	id                        string
	orderDiscountLinesRequest *OrderDiscountLinesRequest
	acceptLanguage            *string
	xChildCompanyId           *string
}

// requested field for a discount lines
func (r ApiOrdersCreateDiscountLineRequest) OrderDiscountLinesRequest(orderDiscountLinesRequest OrderDiscountLinesRequest) ApiOrdersCreateDiscountLineRequest {
	r.orderDiscountLinesRequest = &orderDiscountLinesRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersCreateDiscountLineRequest) AcceptLanguage(acceptLanguage string) ApiOrdersCreateDiscountLineRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersCreateDiscountLineRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersCreateDiscountLineRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersCreateDiscountLineRequest) Execute() (*DiscountLinesResponse, *http.Response, error) {
	return r.ApiService.OrdersCreateDiscountLineExecute(r)
}

/*
OrdersCreateDiscountLine Create Discount

Create discount lines for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrdersCreateDiscountLineRequest
*/
func (a *DiscountsApiService) OrdersCreateDiscountLine(ctx context.Context, id string) ApiOrdersCreateDiscountLineRequest {
	return ApiOrdersCreateDiscountLineRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DiscountLinesResponse
func (a *DiscountsApiService) OrdersCreateDiscountLineExecute(r ApiOrdersCreateDiscountLineRequest) (*DiscountLinesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiscountLinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountsApiService.OrdersCreateDiscountLine")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/discount_lines"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderDiscountLinesRequest == nil {
		return localVarReturnValue, nil, reportError("orderDiscountLinesRequest is required and must be specified")
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
	localVarPostBody = r.orderDiscountLinesRequest
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

type ApiOrdersDeleteDiscountLinesRequest struct {
	ctx             context.Context
	ApiService      DiscountsApi
	id              string
	discountLinesId string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiOrdersDeleteDiscountLinesRequest) AcceptLanguage(acceptLanguage string) ApiOrdersDeleteDiscountLinesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersDeleteDiscountLinesRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersDeleteDiscountLinesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersDeleteDiscountLinesRequest) Execute() (*DiscountLinesResponse, *http.Response, error) {
	return r.ApiService.OrdersDeleteDiscountLinesExecute(r)
}

/*
OrdersDeleteDiscountLines Delete Discount

Delete an existing discount lines for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param discountLinesId discount line id identifier
	@return ApiOrdersDeleteDiscountLinesRequest
*/
func (a *DiscountsApiService) OrdersDeleteDiscountLines(ctx context.Context, id string, discountLinesId string) ApiOrdersDeleteDiscountLinesRequest {
	return ApiOrdersDeleteDiscountLinesRequest{
		ApiService:      a,
		ctx:             ctx,
		id:              id,
		discountLinesId: discountLinesId,
	}
}

// Execute executes the request
//
//	@return DiscountLinesResponse
func (a *DiscountsApiService) OrdersDeleteDiscountLinesExecute(r ApiOrdersDeleteDiscountLinesRequest) (*DiscountLinesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiscountLinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountsApiService.OrdersDeleteDiscountLines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/discount_lines/{discount_lines_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"discount_lines_id"+"}", url.PathEscape(parameterValueToString(r.discountLinesId, "discountLinesId")), -1)

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

type ApiOrdersGetDiscountLineRequest struct {
	ctx             context.Context
	ApiService      DiscountsApi
	id              string
	discountLinesId string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiOrdersGetDiscountLineRequest) AcceptLanguage(acceptLanguage string) ApiOrdersGetDiscountLineRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersGetDiscountLineRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersGetDiscountLineRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersGetDiscountLineRequest) Execute() (*DiscountLinesResponse, *http.Response, error) {
	return r.ApiService.OrdersGetDiscountLineExecute(r)
}

/*
OrdersGetDiscountLine Get Discount

Get an existing discount lines for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param discountLinesId discount line id identifier
	@return ApiOrdersGetDiscountLineRequest
*/
func (a *DiscountsApiService) OrdersGetDiscountLine(ctx context.Context, id string, discountLinesId string) ApiOrdersGetDiscountLineRequest {
	return ApiOrdersGetDiscountLineRequest{
		ApiService:      a,
		ctx:             ctx,
		id:              id,
		discountLinesId: discountLinesId,
	}
}

// Execute executes the request
//
//	@return DiscountLinesResponse
func (a *DiscountsApiService) OrdersGetDiscountLineExecute(r ApiOrdersGetDiscountLineRequest) (*DiscountLinesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiscountLinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountsApiService.OrdersGetDiscountLine")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/discount_lines/{discount_lines_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"discount_lines_id"+"}", url.PathEscape(parameterValueToString(r.discountLinesId, "discountLinesId")), -1)

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

type ApiOrdersGetDiscountLinesRequest struct {
	ctx             context.Context
	ApiService      DiscountsApi
	id              string
	acceptLanguage  *string
	xChildCompanyId *string
	limit           *int32
	search          *string
	next            *string
	previous        *string
}

// Use for knowing which language to use
func (r ApiOrdersGetDiscountLinesRequest) AcceptLanguage(acceptLanguage string) ApiOrdersGetDiscountLinesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersGetDiscountLinesRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersGetDiscountLinesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// The numbers of items to return, the maximum value is 250
func (r ApiOrdersGetDiscountLinesRequest) Limit(limit int32) ApiOrdersGetDiscountLinesRequest {
	r.limit = &limit
	return r
}

// General order search, e.g. by mail, reference etc.
func (r ApiOrdersGetDiscountLinesRequest) Search(search string) ApiOrdersGetDiscountLinesRequest {
	r.search = &search
	return r
}

// next page
func (r ApiOrdersGetDiscountLinesRequest) Next(next string) ApiOrdersGetDiscountLinesRequest {
	r.next = &next
	return r
}

// previous page
func (r ApiOrdersGetDiscountLinesRequest) Previous(previous string) ApiOrdersGetDiscountLinesRequest {
	r.previous = &previous
	return r
}

func (r ApiOrdersGetDiscountLinesRequest) Execute() (*GetOrderDiscountLinesResponse, *http.Response, error) {
	return r.ApiService.OrdersGetDiscountLinesExecute(r)
}

/*
OrdersGetDiscountLines Get a List of Discount

Get discount lines for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrdersGetDiscountLinesRequest
*/
func (a *DiscountsApiService) OrdersGetDiscountLines(ctx context.Context, id string) ApiOrdersGetDiscountLinesRequest {
	return ApiOrdersGetDiscountLinesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetOrderDiscountLinesResponse
func (a *DiscountsApiService) OrdersGetDiscountLinesExecute(r ApiOrdersGetDiscountLinesRequest) (*GetOrderDiscountLinesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderDiscountLinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountsApiService.OrdersGetDiscountLines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/discount_lines"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiOrdersUpdateDiscountLinesRequest struct {
	ctx                             context.Context
	ApiService                      DiscountsApi
	id                              string
	discountLinesId                 string
	updateOrderDiscountLinesRequest *UpdateOrderDiscountLinesRequest
	acceptLanguage                  *string
	xChildCompanyId                 *string
}

// requested field for a discount lines
func (r ApiOrdersUpdateDiscountLinesRequest) UpdateOrderDiscountLinesRequest(updateOrderDiscountLinesRequest UpdateOrderDiscountLinesRequest) ApiOrdersUpdateDiscountLinesRequest {
	r.updateOrderDiscountLinesRequest = &updateOrderDiscountLinesRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersUpdateDiscountLinesRequest) AcceptLanguage(acceptLanguage string) ApiOrdersUpdateDiscountLinesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersUpdateDiscountLinesRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersUpdateDiscountLinesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersUpdateDiscountLinesRequest) Execute() (*DiscountLinesResponse, *http.Response, error) {
	return r.ApiService.OrdersUpdateDiscountLinesExecute(r)
}

/*
OrdersUpdateDiscountLines Update Discount

Update an existing discount lines for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param discountLinesId discount line id identifier
	@return ApiOrdersUpdateDiscountLinesRequest
*/
func (a *DiscountsApiService) OrdersUpdateDiscountLines(ctx context.Context, id string, discountLinesId string) ApiOrdersUpdateDiscountLinesRequest {
	return ApiOrdersUpdateDiscountLinesRequest{
		ApiService:      a,
		ctx:             ctx,
		id:              id,
		discountLinesId: discountLinesId,
	}
}

// Execute executes the request
//
//	@return DiscountLinesResponse
func (a *DiscountsApiService) OrdersUpdateDiscountLinesExecute(r ApiOrdersUpdateDiscountLinesRequest) (*DiscountLinesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiscountLinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountsApiService.OrdersUpdateDiscountLines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/discount_lines/{discount_lines_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"discount_lines_id"+"}", url.PathEscape(parameterValueToString(r.discountLinesId, "discountLinesId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrderDiscountLinesRequest == nil {
		return localVarReturnValue, nil, reportError("updateOrderDiscountLinesRequest is required and must be specified")
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
	localVarPostBody = r.updateOrderDiscountLinesRequest
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

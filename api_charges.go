package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ChargesApi interface {

	/*
		GetCharges Get A List of Charges

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetChargesRequest
	*/
	GetCharges(ctx context.Context) ApiGetChargesRequest

	// GetChargesExecute executes the request
	//  @return GetChargesResponse
	GetChargesExecute(r ApiGetChargesRequest) (*GetChargesResponse, *http.Response, error)

	/*
		OrdersCreateCharge Create charge

		Create charge for an existing orden

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiOrdersCreateChargeRequest
	*/
	OrdersCreateCharge(ctx context.Context, id string) ApiOrdersCreateChargeRequest

	// OrdersCreateChargeExecute executes the request
	//  @return ChargeOrderResponse
	OrdersCreateChargeExecute(r ApiOrdersCreateChargeRequest) (*ChargeOrderResponse, *http.Response, error)
}

// ChargesApiService ChargesApi service
type ChargesApiService service

type ApiGetChargesRequest struct {
	ctx             context.Context
	ApiService      ChargesApi
	acceptLanguage  *string
	xChildCompanyId *string
	limit           *int32
	search          *string
	next            *string
	previous        *string
}

// Use for knowing which language to use
func (r ApiGetChargesRequest) AcceptLanguage(acceptLanguage string) ApiGetChargesRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetChargesRequest) XChildCompanyId(xChildCompanyId string) ApiGetChargesRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// The numbers of items to return, the maximum value is 250
func (r ApiGetChargesRequest) Limit(limit int32) ApiGetChargesRequest {
	r.limit = &limit
	return r
}

// General order search, e.g. by mail, reference etc.
func (r ApiGetChargesRequest) Search(search string) ApiGetChargesRequest {
	r.search = &search
	return r
}

// next page
func (r ApiGetChargesRequest) Next(next string) ApiGetChargesRequest {
	r.next = &next
	return r
}

// previous page
func (r ApiGetChargesRequest) Previous(previous string) ApiGetChargesRequest {
	r.previous = &previous
	return r
}

func (r ApiGetChargesRequest) Execute() (*GetChargesResponse, *http.Response, error) {
	return r.ApiService.GetChargesExecute(r)
}

/*
GetCharges Get A List of Charges

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetChargesRequest
*/
func (a *ChargesApiService) GetCharges(ctx context.Context) ApiGetChargesRequest {
	return ApiGetChargesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetChargesResponse
func (a *ChargesApiService) GetChargesExecute(r ApiGetChargesRequest) (*GetChargesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChargesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChargesApiService.GetCharges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/charges"

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

type ApiOrdersCreateChargeRequest struct {
	ctx             context.Context
	ApiService      ChargesApi
	id              string
	chargeRequest   *ChargeRequest
	acceptLanguage  *string
	xChildCompanyId *string
}

// requested field for a charge
func (r ApiOrdersCreateChargeRequest) ChargeRequest(chargeRequest ChargeRequest) ApiOrdersCreateChargeRequest {
	r.chargeRequest = &chargeRequest
	return r
}

// Use for knowing which language to use
func (r ApiOrdersCreateChargeRequest) AcceptLanguage(acceptLanguage string) ApiOrdersCreateChargeRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiOrdersCreateChargeRequest) XChildCompanyId(xChildCompanyId string) ApiOrdersCreateChargeRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiOrdersCreateChargeRequest) Execute() (*ChargeOrderResponse, *http.Response, error) {
	return r.ApiService.OrdersCreateChargeExecute(r)
}

/*
OrdersCreateCharge Create charge

Create charge for an existing orden

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiOrdersCreateChargeRequest
*/
func (a *ChargesApiService) OrdersCreateCharge(ctx context.Context, id string) ApiOrdersCreateChargeRequest {
	return ApiOrdersCreateChargeRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ChargeOrderResponse
func (a *ChargesApiService) OrdersCreateChargeExecute(r ApiOrdersCreateChargeRequest) (*ChargeOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ChargeOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChargesApiService.OrdersCreateCharge")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{id}/charges"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.chargeRequest == nil {
		return localVarReturnValue, nil, reportError("chargeRequest is required and must be specified")
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
	localVarPostBody = r.chargeRequest
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

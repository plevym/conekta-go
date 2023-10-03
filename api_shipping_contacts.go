package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ShippingContactsApi interface {

	/*
		CreateCustomerShippingContacts Create a shipping contacts

		Create a shipping contacts for a customer.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiCreateCustomerShippingContactsRequest
	*/
	CreateCustomerShippingContacts(ctx context.Context, id string) ApiCreateCustomerShippingContactsRequest

	// CreateCustomerShippingContactsExecute executes the request
	//  @return CustomerShippingContactsResponse
	CreateCustomerShippingContactsExecute(r ApiCreateCustomerShippingContactsRequest) (*CustomerShippingContactsResponse, *http.Response, error)

	/*
		DeleteCustomerShippingContacts Delete shipping contacts

		Delete shipping contact that corresponds to a customer ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param shippingContactsId identifier
		@return ApiDeleteCustomerShippingContactsRequest
	*/
	DeleteCustomerShippingContacts(ctx context.Context, id string, shippingContactsId string) ApiDeleteCustomerShippingContactsRequest

	// DeleteCustomerShippingContactsExecute executes the request
	//  @return CustomerShippingContactsResponse
	DeleteCustomerShippingContactsExecute(r ApiDeleteCustomerShippingContactsRequest) (*CustomerShippingContactsResponse, *http.Response, error)

	/*
		UpdateCustomerShippingContacts Update shipping contacts

		Update shipping contact that corresponds to a customer ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@param shippingContactsId identifier
		@return ApiUpdateCustomerShippingContactsRequest
	*/
	UpdateCustomerShippingContacts(ctx context.Context, id string, shippingContactsId string) ApiUpdateCustomerShippingContactsRequest

	// UpdateCustomerShippingContactsExecute executes the request
	//  @return CustomerShippingContactsResponse
	UpdateCustomerShippingContactsExecute(r ApiUpdateCustomerShippingContactsRequest) (*CustomerShippingContactsResponse, *http.Response, error)
}

// ShippingContactsApiService ShippingContactsApi service
type ShippingContactsApiService service

type ApiCreateCustomerShippingContactsRequest struct {
	ctx                      context.Context
	ApiService               ShippingContactsApi
	id                       string
	customerShippingContacts *CustomerShippingContacts
	acceptLanguage           *string
	xChildCompanyId          *string
}

// requested field for customer shippings contacts
func (r ApiCreateCustomerShippingContactsRequest) CustomerShippingContacts(customerShippingContacts CustomerShippingContacts) ApiCreateCustomerShippingContactsRequest {
	r.customerShippingContacts = &customerShippingContacts
	return r
}

// Use for knowing which language to use
func (r ApiCreateCustomerShippingContactsRequest) AcceptLanguage(acceptLanguage string) ApiCreateCustomerShippingContactsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiCreateCustomerShippingContactsRequest) XChildCompanyId(xChildCompanyId string) ApiCreateCustomerShippingContactsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiCreateCustomerShippingContactsRequest) Execute() (*CustomerShippingContactsResponse, *http.Response, error) {
	return r.ApiService.CreateCustomerShippingContactsExecute(r)
}

/*
CreateCustomerShippingContacts Create a shipping contacts

Create a shipping contacts for a customer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiCreateCustomerShippingContactsRequest
*/
func (a *ShippingContactsApiService) CreateCustomerShippingContacts(ctx context.Context, id string) ApiCreateCustomerShippingContactsRequest {
	return ApiCreateCustomerShippingContactsRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CustomerShippingContactsResponse
func (a *ShippingContactsApiService) CreateCustomerShippingContactsExecute(r ApiCreateCustomerShippingContactsRequest) (*CustomerShippingContactsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomerShippingContactsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingContactsApiService.CreateCustomerShippingContacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/shipping_contacts"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerShippingContacts == nil {
		return localVarReturnValue, nil, reportError("customerShippingContacts is required and must be specified")
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
	localVarPostBody = r.customerShippingContacts
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

type ApiDeleteCustomerShippingContactsRequest struct {
	ctx                context.Context
	ApiService         ShippingContactsApi
	id                 string
	shippingContactsId string
	acceptLanguage     *string
	xChildCompanyId    *string
}

// Use for knowing which language to use
func (r ApiDeleteCustomerShippingContactsRequest) AcceptLanguage(acceptLanguage string) ApiDeleteCustomerShippingContactsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiDeleteCustomerShippingContactsRequest) XChildCompanyId(xChildCompanyId string) ApiDeleteCustomerShippingContactsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiDeleteCustomerShippingContactsRequest) Execute() (*CustomerShippingContactsResponse, *http.Response, error) {
	return r.ApiService.DeleteCustomerShippingContactsExecute(r)
}

/*
DeleteCustomerShippingContacts Delete shipping contacts

Delete shipping contact that corresponds to a customer ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param shippingContactsId identifier
	@return ApiDeleteCustomerShippingContactsRequest
*/
func (a *ShippingContactsApiService) DeleteCustomerShippingContacts(ctx context.Context, id string, shippingContactsId string) ApiDeleteCustomerShippingContactsRequest {
	return ApiDeleteCustomerShippingContactsRequest{
		ApiService:         a,
		ctx:                ctx,
		id:                 id,
		shippingContactsId: shippingContactsId,
	}
}

// Execute executes the request
//
//	@return CustomerShippingContactsResponse
func (a *ShippingContactsApiService) DeleteCustomerShippingContactsExecute(r ApiDeleteCustomerShippingContactsRequest) (*CustomerShippingContactsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomerShippingContactsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingContactsApiService.DeleteCustomerShippingContacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/shipping_contacts/{shipping_contacts_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shipping_contacts_id"+"}", url.PathEscape(parameterValueToString(r.shippingContactsId, "shippingContactsId")), -1)

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

type ApiUpdateCustomerShippingContactsRequest struct {
	ctx                            context.Context
	ApiService                     ShippingContactsApi
	id                             string
	shippingContactsId             string
	customerUpdateShippingContacts *CustomerUpdateShippingContacts
	acceptLanguage                 *string
	xChildCompanyId                *string
}

// requested field for customer update shippings contacts
func (r ApiUpdateCustomerShippingContactsRequest) CustomerUpdateShippingContacts(customerUpdateShippingContacts CustomerUpdateShippingContacts) ApiUpdateCustomerShippingContactsRequest {
	r.customerUpdateShippingContacts = &customerUpdateShippingContacts
	return r
}

// Use for knowing which language to use
func (r ApiUpdateCustomerShippingContactsRequest) AcceptLanguage(acceptLanguage string) ApiUpdateCustomerShippingContactsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiUpdateCustomerShippingContactsRequest) XChildCompanyId(xChildCompanyId string) ApiUpdateCustomerShippingContactsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiUpdateCustomerShippingContactsRequest) Execute() (*CustomerShippingContactsResponse, *http.Response, error) {
	return r.ApiService.UpdateCustomerShippingContactsExecute(r)
}

/*
UpdateCustomerShippingContacts Update shipping contacts

Update shipping contact that corresponds to a customer ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@param shippingContactsId identifier
	@return ApiUpdateCustomerShippingContactsRequest
*/
func (a *ShippingContactsApiService) UpdateCustomerShippingContacts(ctx context.Context, id string, shippingContactsId string) ApiUpdateCustomerShippingContactsRequest {
	return ApiUpdateCustomerShippingContactsRequest{
		ApiService:         a,
		ctx:                ctx,
		id:                 id,
		shippingContactsId: shippingContactsId,
	}
}

// Execute executes the request
//
//	@return CustomerShippingContactsResponse
func (a *ShippingContactsApiService) UpdateCustomerShippingContactsExecute(r ApiUpdateCustomerShippingContactsRequest) (*CustomerShippingContactsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomerShippingContactsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingContactsApiService.UpdateCustomerShippingContacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}/shipping_contacts/{shipping_contacts_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shipping_contacts_id"+"}", url.PathEscape(parameterValueToString(r.shippingContactsId, "shippingContactsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerUpdateShippingContacts == nil {
		return localVarReturnValue, nil, reportError("customerUpdateShippingContacts is required and must be specified")
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
	localVarPostBody = r.customerUpdateShippingContacts
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

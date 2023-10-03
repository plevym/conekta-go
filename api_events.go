package conekta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type EventsApi interface {

	/*
		GetEvent Get Event

		Returns a single event

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Identifier of the resource
		@return ApiGetEventRequest
	*/
	GetEvent(ctx context.Context, id string) ApiGetEventRequest

	// GetEventExecute executes the request
	//  @return EventResponse
	GetEventExecute(r ApiGetEventRequest) (*EventResponse, *http.Response, error)

	/*
		GetEvents Get list of Events

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetEventsRequest
	*/
	GetEvents(ctx context.Context) ApiGetEventsRequest

	// GetEventsExecute executes the request
	//  @return GetEventsResponse
	GetEventsExecute(r ApiGetEventsRequest) (*GetEventsResponse, *http.Response, error)

	/*
		ResendEvent Resend Event

		Try to send an event

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param eventId event identifier
		@param webhookLogId webhook log identifier
		@return ApiResendEventRequest
	*/
	ResendEvent(ctx context.Context, eventId string, webhookLogId string) ApiResendEventRequest

	// ResendEventExecute executes the request
	//  @return EventsResendResponse
	ResendEventExecute(r ApiResendEventRequest) (*EventsResendResponse, *http.Response, error)
}

// EventsApiService EventsApi service
type EventsApiService service

type ApiGetEventRequest struct {
	ctx             context.Context
	ApiService      EventsApi
	id              string
	acceptLanguage  *string
	xChildCompanyId *string
}

// Use for knowing which language to use
func (r ApiGetEventRequest) AcceptLanguage(acceptLanguage string) ApiGetEventRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetEventRequest) XChildCompanyId(xChildCompanyId string) ApiGetEventRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

func (r ApiGetEventRequest) Execute() (*EventResponse, *http.Response, error) {
	return r.ApiService.GetEventExecute(r)
}

/*
GetEvent Get Event

Returns a single event

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Identifier of the resource
	@return ApiGetEventRequest
*/
func (a *EventsApiService) GetEvent(ctx context.Context, id string) ApiGetEventRequest {
	return ApiGetEventRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return EventResponse
func (a *EventsApiService) GetEventExecute(r ApiGetEventRequest) (*EventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{id}"
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

type ApiGetEventsRequest struct {
	ctx             context.Context
	ApiService      EventsApi
	acceptLanguage  *string
	xChildCompanyId *string
	limit           *int32
	search          *string
	next            *string
	previous        *string
}

// Use for knowing which language to use
func (r ApiGetEventsRequest) AcceptLanguage(acceptLanguage string) ApiGetEventsRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// In the case of a holding company, the company id of the child company to which will process the request.
func (r ApiGetEventsRequest) XChildCompanyId(xChildCompanyId string) ApiGetEventsRequest {
	r.xChildCompanyId = &xChildCompanyId
	return r
}

// The numbers of items to return, the maximum value is 250
func (r ApiGetEventsRequest) Limit(limit int32) ApiGetEventsRequest {
	r.limit = &limit
	return r
}

// General order search, e.g. by mail, reference etc.
func (r ApiGetEventsRequest) Search(search string) ApiGetEventsRequest {
	r.search = &search
	return r
}

// next page
func (r ApiGetEventsRequest) Next(next string) ApiGetEventsRequest {
	r.next = &next
	return r
}

// previous page
func (r ApiGetEventsRequest) Previous(previous string) ApiGetEventsRequest {
	r.previous = &previous
	return r
}

func (r ApiGetEventsRequest) Execute() (*GetEventsResponse, *http.Response, error) {
	return r.ApiService.GetEventsExecute(r)
}

/*
GetEvents Get list of Events

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEventsRequest
*/
func (a *EventsApiService) GetEvents(ctx context.Context) ApiGetEventsRequest {
	return ApiGetEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEventsResponse
func (a *EventsApiService) GetEventsExecute(r ApiGetEventsRequest) (*GetEventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

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

type ApiResendEventRequest struct {
	ctx            context.Context
	ApiService     EventsApi
	eventId        string
	webhookLogId   string
	acceptLanguage *string
}

// Use for knowing which language to use
func (r ApiResendEventRequest) AcceptLanguage(acceptLanguage string) ApiResendEventRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiResendEventRequest) Execute() (*EventsResendResponse, *http.Response, error) {
	return r.ApiService.ResendEventExecute(r)
}

/*
ResendEvent Resend Event

Try to send an event

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventId event identifier
	@param webhookLogId webhook log identifier
	@return ApiResendEventRequest
*/
func (a *EventsApiService) ResendEvent(ctx context.Context, eventId string, webhookLogId string) ApiResendEventRequest {
	return ApiResendEventRequest{
		ApiService:   a,
		ctx:          ctx,
		eventId:      eventId,
		webhookLogId: webhookLogId,
	}
}

// Execute executes the request
//
//	@return EventsResendResponse
func (a *EventsApiService) ResendEventExecute(r ApiResendEventRequest) (*EventsResendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EventsResendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ResendEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{event_id}/webhook_logs/{webhook_log_id}/resend"
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webhook_log_id"+"}", url.PathEscape(parameterValueToString(r.webhookLogId, "webhookLogId")), -1)

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

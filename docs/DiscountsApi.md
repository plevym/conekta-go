# \DiscountsApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersCreateDiscountLine**](DiscountsApi.md#OrdersCreateDiscountLine) | **Post** /orders/{id}/discount_lines | Create Discount
[**OrdersDeleteDiscountLines**](DiscountsApi.md#OrdersDeleteDiscountLines) | **Delete** /orders/{id}/discount_lines/{discount_lines_id} | Delete Discount
[**OrdersUpdateDiscountLines**](DiscountsApi.md#OrdersUpdateDiscountLines) | **Put** /orders/{id}/discount_lines/{discount_lines_id} | Update Discount



## OrdersCreateDiscountLine

> DiscountLinesResponse OrdersCreateDiscountLine(ctx, id).OrderDiscountLinesRequest(orderDiscountLinesRequest).AcceptLanguage(acceptLanguage).Execute()

Create Discount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conekta/conekta-go"
)

func main() {
    id := "6307a60c41de27127515a575" // string | Identifier of the resource
    orderDiscountLinesRequest := *openapiclient.NewOrderDiscountLinesRequest(int64(500), "123", "loyalty") // OrderDiscountLinesRequest | requested field for a discount lines
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiscountsApi.OrdersCreateDiscountLine(context.Background(), id).OrderDiscountLinesRequest(orderDiscountLinesRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.OrdersCreateDiscountLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateDiscountLine`: DiscountLinesResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.OrdersCreateDiscountLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateDiscountLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderDiscountLinesRequest** | [**OrderDiscountLinesRequest**](OrderDiscountLinesRequest.md) | requested field for a discount lines | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**DiscountLinesResponse**](DiscountLinesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersDeleteDiscountLines

> DiscountLinesResponse OrdersDeleteDiscountLines(ctx, id, discountLinesId).AcceptLanguage(acceptLanguage).Execute()

Delete Discount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conekta/conekta-go"
)

func main() {
    id := "6307a60c41de27127515a575" // string | Identifier of the resource
    discountLinesId := "dis_lin_2tQ974hSHcsdeSZHG" // string | identifier
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiscountsApi.OrdersDeleteDiscountLines(context.Background(), id, discountLinesId).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.OrdersDeleteDiscountLines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersDeleteDiscountLines`: DiscountLinesResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.OrdersDeleteDiscountLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**discountLinesId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersDeleteDiscountLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**DiscountLinesResponse**](DiscountLinesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersUpdateDiscountLines

> DiscountLinesResponse OrdersUpdateDiscountLines(ctx, id, discountLinesId).UpdateOrderDiscountLinesRequest(updateOrderDiscountLinesRequest).AcceptLanguage(acceptLanguage).Execute()

Update Discount



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conekta/conekta-go"
)

func main() {
    id := "6307a60c41de27127515a575" // string | Identifier of the resource
    discountLinesId := "dis_lin_2tQ974hSHcsdeSZHG" // string | identifier
    updateOrderDiscountLinesRequest := *openapiclient.NewUpdateOrderDiscountLinesRequest() // UpdateOrderDiscountLinesRequest | requested field for a discount lines
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiscountsApi.OrdersUpdateDiscountLines(context.Background(), id, discountLinesId).UpdateOrderDiscountLinesRequest(updateOrderDiscountLinesRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.OrdersUpdateDiscountLines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersUpdateDiscountLines`: DiscountLinesResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.OrdersUpdateDiscountLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**discountLinesId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersUpdateDiscountLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderDiscountLinesRequest** | [**UpdateOrderDiscountLinesRequest**](UpdateOrderDiscountLinesRequest.md) | requested field for a discount lines | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**DiscountLinesResponse**](DiscountLinesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

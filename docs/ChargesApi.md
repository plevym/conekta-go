# \ChargesApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharges**](ChargesApi.md#GetCharges) | **Get** /charges | Get A List of Charges
[**OrdersCreateCharge**](ChargesApi.md#OrdersCreateCharge) | **Post** /orders/{id}/charges | Create charge



## GetCharges

> GetChargesResponse GetCharges(ctx).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get A List of Charges

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/plevym/conekta-go"
)

func main() {
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
    xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)
    limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
    search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)
    next := "next_example" // string | next page (optional)
    previous := "previous_example" // string | previous page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesApi.GetCharges(context.Background()).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesApi.GetCharges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCharges`: GetChargesResponse
    fmt.Fprintf(os.Stdout, "Response from `ChargesApi.GetCharges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**GetChargesResponse**](GetChargesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersCreateCharge

> ChargeOrderResponse OrdersCreateCharge(ctx, id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create charge



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/plevym/conekta-go"
)

func main() {
    id := "6307a60c41de27127515a575" // string | Identifier of the resource
    chargeRequest := *openapiclient.NewChargeRequest(*openapiclient.NewChargeRequestPaymentMethod("card")) // ChargeRequest | requested field for a charge
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
    xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesApi.OrdersCreateCharge(context.Background(), id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesApi.OrdersCreateCharge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateCharge`: ChargeOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ChargesApi.OrdersCreateCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeRequest** | [**ChargeRequest**](ChargeRequest.md) | requested field for a charge | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ChargeOrderResponse**](ChargeOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

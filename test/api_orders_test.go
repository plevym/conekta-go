/*
Conekta API

Testing OrdersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/plevym/conekta-go"
)

func Test_conekta_OrdersApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test OrdersApiService CancelOrder", func(t *testing.T) {
		orderID := "ord_2tqaGQYZyvBsMKEgs"
		resp, httpRes, err := apiClient.OrdersApi.CancelOrder(context.TODO(), orderID).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService CreateOrder", func(t *testing.T) {
		customerID := "cus_2tYENskzTjjgkGQLt"
		req := conekta.OrderRequest{
			Currency: "MXN",
			CustomerInfo: conekta.OrderRequestCustomerInfo{
				CustomerInfoJustCustomerId: conekta.NewCustomerInfoJustCustomerId(customerID),
			},
			LineItems: []conekta.Product{
				{
					Quantity:  1,
					UnitPrice: 1000,
				},
			},
		}
		resp, httpRes, err := apiClient.OrdersApi.CreateOrder(context.TODO()).OrderRequest(req).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService GetOrderById", func(t *testing.T) {
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		resp, httpRes, err := apiClient.OrdersApi.GetOrderById(context.TODO(), orderID).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService GetOrders", func(t *testing.T) {
		resp, httpRes, err := apiClient.OrdersApi.GetOrders(context.TODO()).Limit(10).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService OrderCancelRefund", func(t *testing.T) {
		orderID := "ord_2tV52JvSom2w3E8bX"
		refundID := "6407b5bee1329a000175ba11"
		resp, httpRes, err := apiClient.OrdersApi.OrderCancelRefund(context.TODO(), orderID, refundID).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService OrderRefund", func(t *testing.T) {
		req := conekta.OrderRefundRequest{
			Amount: 1000,
			Reason: "reason",
		}
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		resp, httpRes, err := apiClient.OrdersApi.OrderRefund(context.TODO(), orderID).
			OrderRefundRequest(req).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService OrdersCreateCapture", func(t *testing.T) {
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		resp, httpRes, err := apiClient.OrdersApi.OrdersCreateCapture(context.TODO(), orderID).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test OrdersApiService UpdateOrder", func(t *testing.T) {
		req := conekta.OrderUpdateRequest{
			Currency: conekta.PtrString("MXN"),
			CustomerInfo: &conekta.OrderUpdateRequestCustomerInfo{
				CustomerInfoJustCustomerId: conekta.NewCustomerInfoJustCustomerId("cus_2tYENskzTjjgkGQLt"),
			},
			LineItems: []conekta.Product{
				{
					Quantity:  1,
					UnitPrice: 1200,
				},
			},
		}
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		resp, httpRes, err := apiClient.OrdersApi.UpdateOrder(context.TODO(), orderID).
			OrderUpdateRequest(req).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

}

/*
Conekta API

Testing DiscountsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/plevym/conekta-go"
)

func Test_conekta_DiscountsApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test DiscountsApiService OrdersCreateDiscountLine", func(t *testing.T) {
		req := conekta.OrderDiscountLinesRequest{
			Amount: 1000,
		}
		resp, httpRes, err := apiClient.DiscountsApi.OrdersCreateDiscountLine(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			OrderDiscountLinesRequest(req).
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

	t.Run("Test DiscountsApiService OrdersDeleteDiscountLines", func(t *testing.T) {
		resp, httpRes, err := apiClient.DiscountsApi.OrdersDeleteDiscountLines(context.TODO(), "ord_2tPAmKCEJqh8RE6nY", "dis_lin_2tQQ58HPgPw7StE8z").
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

	t.Run("Test DiscountsApiService OrdersUpdateDiscountLines", func(t *testing.T) {
		req := conekta.UpdateOrderDiscountLinesRequest{
			Amount: conekta.PtrInt64(2000),
		}
		resp, httpRes, err := apiClient.DiscountsApi.OrdersUpdateDiscountLines(context.TODO(), "ord_2tPAmKCEJqh8RE6nY", "dis_lin_2tQQ58HPgPw7StE8z").
			UpdateOrderDiscountLinesRequest(req).
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

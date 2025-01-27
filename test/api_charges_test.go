/*
Conekta API

Testing ChargesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/plevym/conekta-go"
)

func Test_conekta_ChargesApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test ChargesApiService GetCharges", func(t *testing.T) {
		resp, httpRes, err := apiClient.ChargesApi.GetCharges(context.TODO()).
			Limit(20).
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

	t.Run("Test ChargesApiService OrdersCreateCharge", func(t *testing.T) {
		req := conekta.ChargeRequest{
			Amount: conekta.PtrInt32(1000),
		}
		resp, httpRes, err := apiClient.ChargesApi.OrdersCreateCharge(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			ChargeRequest(req).
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

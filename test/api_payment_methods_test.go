/*
Conekta API

Testing PaymentMethodsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/plevym/conekta-go"
)

func Test_conekta_PaymentMethodsApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test PaymentMethodsApiService CreateCustomerPaymentMethods", func(t *testing.T) {
		req := conekta.CreateCustomerPaymentMethodsRequest{
			PaymentMethodCardRequest: &conekta.PaymentMethodCardRequest{
				Type:    "oxxo_recurrent",
				TokenId: "tokenID",
			},
		}
		resp, httpRes, err := apiClient.PaymentMethodsApi.CreateCustomerPaymentMethods(context.TODO(), "cus_2tXyF9BwPG14UMkkg").
			CreateCustomerPaymentMethodsRequest(req).
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

	t.Run("Test PaymentMethodsApiService DeleteCustomerPaymentMethods", func(t *testing.T) {
		resp, httpRes, err := apiClient.PaymentMethodsApi.DeleteCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt", "src_2tbd5Bgy67RL9oycM").
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

	t.Run("Test PaymentMethodsApiService GetCustomerPaymentMethods", func(t *testing.T) {
		resp, httpRes, err := apiClient.PaymentMethodsApi.GetCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test PaymentMethodsApiService UpdateCustomerPaymentMethods", func(t *testing.T) {
		req := conekta.UpdatePaymentMethods{
			Name: conekta.PtrString("name"),
		}
		resp, httpRes, err := apiClient.PaymentMethodsApi.UpdateCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt", "src_2tbd5Bgy67RL9oycM").
			UpdatePaymentMethods(req).
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

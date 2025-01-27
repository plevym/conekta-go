/*
Conekta API

Testing CustomersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/plevym/conekta-go"
)

func Test_conekta_CustomersApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test CustomersApiService CreateCustomer", func(t *testing.T) {
		req := conekta.Customer{
			Email: "foo@foo.com",
			Name:  "Foo Foo",
		}
		resp, httpRes, err := apiClient.CustomersApi.CreateCustomer(context.TODO()).
			Customer(req).
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

	t.Run("Test CustomersApiService CreateCustomerFiscalEntities", func(t *testing.T) {
		req := conekta.CustomerFiscalEntitiesRequest{
			Email:       conekta.PtrString("foo@foo.com"),
			CompanyName: conekta.PtrString("FooCompany"),
		}
		resp, httpRes, err := apiClient.CustomersApi.CreateCustomerFiscalEntities(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			CustomerFiscalEntitiesRequest(req).
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

	t.Run("Test CustomersApiService DeleteCustomerById", func(t *testing.T) {
		resp, httpRes, err := apiClient.CustomersApi.DeleteCustomerById(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test CustomersApiService GetCustomerById", func(t *testing.T) {
		resp, httpRes, err := apiClient.CustomersApi.GetCustomerById(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test CustomersApiService GetCustomers", func(t *testing.T) {
		resp, httpRes, err := apiClient.CustomersApi.GetCustomers(context.TODO()).
			Limit(10).
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

	t.Run("Test CustomersApiService UpdateCustomer", func(t *testing.T) {
		req := conekta.UpdateCustomer{
			Email: conekta.PtrString("foo@foo.com"),
			Name:  conekta.PtrString("Foo Foo"),
		}
		resp, httpRes, err := apiClient.CustomersApi.UpdateCustomer(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			UpdateCustomer(req).
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

	t.Run("Test CustomersApiService UpdateCustomerFiscalEntities", func(t *testing.T) {
		req := conekta.CustomerUpdateFiscalEntitiesRequest{
			Email:       conekta.PtrString("foo@foo.com"),
			CompanyName: conekta.PtrString("FooCompany"),
		}
		resp, httpRes, err := apiClient.CustomersApi.UpdateCustomerFiscalEntities(context.TODO(), "cus_2tYENskzTjjgkGQLt", "fis_ent_2tYENskzTjjgkGQLr").
			CustomerUpdateFiscalEntitiesRequest(req).
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

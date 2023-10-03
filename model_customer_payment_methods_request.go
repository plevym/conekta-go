package conekta

import (
	"encoding/json"
	"fmt"
)

// CustomerPaymentMethodsRequest - struct for CustomerPaymentMethodsRequest
type CustomerPaymentMethodsRequest struct {
	PaymentMethodCardRequest *PaymentMethodCardRequest
	PaymentMethodCashRequest *PaymentMethodCashRequest
	PaymentMethodSpeiRequest *PaymentMethodSpeiRequest
}

// PaymentMethodCardRequestAsCustomerPaymentMethodsRequest is a convenience function that returns PaymentMethodCardRequest wrapped in CustomerPaymentMethodsRequest
func PaymentMethodCardRequestAsCustomerPaymentMethodsRequest(v *PaymentMethodCardRequest) CustomerPaymentMethodsRequest {
	return CustomerPaymentMethodsRequest{
		PaymentMethodCardRequest: v,
	}
}

// PaymentMethodCashRequestAsCustomerPaymentMethodsRequest is a convenience function that returns PaymentMethodCashRequest wrapped in CustomerPaymentMethodsRequest
func PaymentMethodCashRequestAsCustomerPaymentMethodsRequest(v *PaymentMethodCashRequest) CustomerPaymentMethodsRequest {
	return CustomerPaymentMethodsRequest{
		PaymentMethodCashRequest: v,
	}
}

// PaymentMethodSpeiRequestAsCustomerPaymentMethodsRequest is a convenience function that returns PaymentMethodSpeiRequest wrapped in CustomerPaymentMethodsRequest
func PaymentMethodSpeiRequestAsCustomerPaymentMethodsRequest(v *PaymentMethodSpeiRequest) CustomerPaymentMethodsRequest {
	return CustomerPaymentMethodsRequest{
		PaymentMethodSpeiRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomerPaymentMethodsRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PaymentMethodCardRequest
	err = json.Unmarshal(data, &dst.PaymentMethodCardRequest)
	if err == nil {
		jsonPaymentMethodCardRequest, _ := json.Marshal(dst.PaymentMethodCardRequest)
		if string(jsonPaymentMethodCardRequest) == "{}" { // empty struct
			dst.PaymentMethodCardRequest = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodCardRequest = nil
	}

	// try to unmarshal data into PaymentMethodCashRequest
	err = json.Unmarshal(data, &dst.PaymentMethodCashRequest)
	if err == nil {
		jsonPaymentMethodCashRequest, _ := json.Marshal(dst.PaymentMethodCashRequest)
		if string(jsonPaymentMethodCashRequest) == "{}" { // empty struct
			dst.PaymentMethodCashRequest = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodCashRequest = nil
	}

	// try to unmarshal data into PaymentMethodSpeiRequest
	err = json.Unmarshal(data, &dst.PaymentMethodSpeiRequest)
	if err == nil {
		jsonPaymentMethodSpeiRequest, _ := json.Marshal(dst.PaymentMethodSpeiRequest)
		if string(jsonPaymentMethodSpeiRequest) == "{}" { // empty struct
			dst.PaymentMethodSpeiRequest = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodSpeiRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PaymentMethodCardRequest = nil
		dst.PaymentMethodCashRequest = nil
		dst.PaymentMethodSpeiRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CustomerPaymentMethodsRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CustomerPaymentMethodsRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomerPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodCardRequest != nil {
		return json.Marshal(&src.PaymentMethodCardRequest)
	}

	if src.PaymentMethodCashRequest != nil {
		return json.Marshal(&src.PaymentMethodCashRequest)
	}

	if src.PaymentMethodSpeiRequest != nil {
		return json.Marshal(&src.PaymentMethodSpeiRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CustomerPaymentMethodsRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodCardRequest != nil {
		return obj.PaymentMethodCardRequest
	}

	if obj.PaymentMethodCashRequest != nil {
		return obj.PaymentMethodCashRequest
	}

	if obj.PaymentMethodSpeiRequest != nil {
		return obj.PaymentMethodSpeiRequest
	}

	// all schemas are nil
	return nil
}

type NullableCustomerPaymentMethodsRequest struct {
	value *CustomerPaymentMethodsRequest
	isSet bool
}

func (v NullableCustomerPaymentMethodsRequest) Get() *CustomerPaymentMethodsRequest {
	return v.value
}

func (v *NullableCustomerPaymentMethodsRequest) Set(val *CustomerPaymentMethodsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentMethodsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentMethodsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentMethodsRequest(val *CustomerPaymentMethodsRequest) *NullableCustomerPaymentMethodsRequest {
	return &NullableCustomerPaymentMethodsRequest{value: val, isSet: true}
}

func (v NullableCustomerPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentMethodsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

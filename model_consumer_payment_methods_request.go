package conekta

import (
	"encoding/json"
	"fmt"
)

// ConsumerPaymentMethodsRequest - struct for ConsumerPaymentMethodsRequest
type ConsumerPaymentMethodsRequest struct {
	PaymentMethodCardRequest *PaymentMethodCardRequest
	PaymentMethodCashRequest *PaymentMethodCashRequest
	PaymentMethodSpeiRequest *PaymentMethodSpeiRequest
}

// PaymentMethodCardRequestAsConsumerPaymentMethodsRequest is a convenience function that returns PaymentMethodCardRequest wrapped in ConsumerPaymentMethodsRequest
func PaymentMethodCardRequestAsConsumerPaymentMethodsRequest(v *PaymentMethodCardRequest) ConsumerPaymentMethodsRequest {
	return ConsumerPaymentMethodsRequest{
		PaymentMethodCardRequest: v,
	}
}

// PaymentMethodCashRequestAsConsumerPaymentMethodsRequest is a convenience function that returns PaymentMethodCashRequest wrapped in ConsumerPaymentMethodsRequest
func PaymentMethodCashRequestAsConsumerPaymentMethodsRequest(v *PaymentMethodCashRequest) ConsumerPaymentMethodsRequest {
	return ConsumerPaymentMethodsRequest{
		PaymentMethodCashRequest: v,
	}
}

// PaymentMethodSpeiRequestAsConsumerPaymentMethodsRequest is a convenience function that returns PaymentMethodSpeiRequest wrapped in ConsumerPaymentMethodsRequest
func PaymentMethodSpeiRequestAsConsumerPaymentMethodsRequest(v *PaymentMethodSpeiRequest) ConsumerPaymentMethodsRequest {
	return ConsumerPaymentMethodsRequest{
		PaymentMethodSpeiRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConsumerPaymentMethodsRequest) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ConsumerPaymentMethodsRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ConsumerPaymentMethodsRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConsumerPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
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
func (obj *ConsumerPaymentMethodsRequest) GetActualInstance() interface{} {
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

type NullableConsumerPaymentMethodsRequest struct {
	value *ConsumerPaymentMethodsRequest
	isSet bool
}

func (v NullableConsumerPaymentMethodsRequest) Get() *ConsumerPaymentMethodsRequest {
	return v.value
}

func (v *NullableConsumerPaymentMethodsRequest) Set(val *ConsumerPaymentMethodsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerPaymentMethodsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerPaymentMethodsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerPaymentMethodsRequest(val *ConsumerPaymentMethodsRequest) *NullableConsumerPaymentMethodsRequest {
	return &NullableConsumerPaymentMethodsRequest{value: val, isSet: true}
}

func (v NullableConsumerPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerPaymentMethodsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

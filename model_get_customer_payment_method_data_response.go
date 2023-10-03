package conekta

import (
	"encoding/json"
	"fmt"
)

// GetCustomerPaymentMethodDataResponse - struct for GetCustomerPaymentMethodDataResponse
type GetCustomerPaymentMethodDataResponse struct {
	PaymentMethodCardResponse  *PaymentMethodCardResponse
	PaymentMethodCashResponse  *PaymentMethodCashResponse
	PaymentMethodSpeiRecurrent *PaymentMethodSpeiRecurrent
}

// PaymentMethodCardResponseAsGetCustomerPaymentMethodDataResponse is a convenience function that returns PaymentMethodCardResponse wrapped in GetCustomerPaymentMethodDataResponse
func PaymentMethodCardResponseAsGetCustomerPaymentMethodDataResponse(v *PaymentMethodCardResponse) GetCustomerPaymentMethodDataResponse {
	return GetCustomerPaymentMethodDataResponse{
		PaymentMethodCardResponse: v,
	}
}

// PaymentMethodCashResponseAsGetCustomerPaymentMethodDataResponse is a convenience function that returns PaymentMethodCashResponse wrapped in GetCustomerPaymentMethodDataResponse
func PaymentMethodCashResponseAsGetCustomerPaymentMethodDataResponse(v *PaymentMethodCashResponse) GetCustomerPaymentMethodDataResponse {
	return GetCustomerPaymentMethodDataResponse{
		PaymentMethodCashResponse: v,
	}
}

// PaymentMethodSpeiRecurrentAsGetCustomerPaymentMethodDataResponse is a convenience function that returns PaymentMethodSpeiRecurrent wrapped in GetCustomerPaymentMethodDataResponse
func PaymentMethodSpeiRecurrentAsGetCustomerPaymentMethodDataResponse(v *PaymentMethodSpeiRecurrent) GetCustomerPaymentMethodDataResponse {
	return GetCustomerPaymentMethodDataResponse{
		PaymentMethodSpeiRecurrent: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetCustomerPaymentMethodDataResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'card'
	if jsonDict["type"] == "card" {
		// try to unmarshal JSON data into PaymentMethodCardResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCardResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCardResponse, return on the first match
		} else {
			dst.PaymentMethodCardResponse = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodCardResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cash'
	if jsonDict["type"] == "cash" {
		// try to unmarshal JSON data into PaymentMethodCashResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCashResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCashResponse, return on the first match
		} else {
			dst.PaymentMethodCashResponse = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodCashResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'oxxo_recurrent'
	if jsonDict["type"] == "oxxo_recurrent" {
		// try to unmarshal JSON data into PaymentMethodCashResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCashResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCashResponse, return on the first match
		} else {
			dst.PaymentMethodCashResponse = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodCashResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_card_response'
	if jsonDict["type"] == "payment_method_card_response" {
		// try to unmarshal JSON data into PaymentMethodCardResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCardResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCardResponse, return on the first match
		} else {
			dst.PaymentMethodCardResponse = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodCardResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_cash_response'
	if jsonDict["type"] == "payment_method_cash_response" {
		// try to unmarshal JSON data into PaymentMethodCashResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCashResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCashResponse, return on the first match
		} else {
			dst.PaymentMethodCashResponse = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodCashResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_spei_recurrent'
	if jsonDict["type"] == "payment_method_spei_recurrent" {
		// try to unmarshal JSON data into PaymentMethodSpeiRecurrent
		err = json.Unmarshal(data, &dst.PaymentMethodSpeiRecurrent)
		if err == nil {
			return nil // data stored in dst.PaymentMethodSpeiRecurrent, return on the first match
		} else {
			dst.PaymentMethodSpeiRecurrent = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodSpeiRecurrent: %s", err.Error())
		}
	}

	// check if the discriminator value is 'spei_recurrent'
	if jsonDict["type"] == "spei_recurrent" {
		// try to unmarshal JSON data into PaymentMethodSpeiRecurrent
		err = json.Unmarshal(data, &dst.PaymentMethodSpeiRecurrent)
		if err == nil {
			return nil // data stored in dst.PaymentMethodSpeiRecurrent, return on the first match
		} else {
			dst.PaymentMethodSpeiRecurrent = nil
			return fmt.Errorf("failed to unmarshal GetCustomerPaymentMethodDataResponse as PaymentMethodSpeiRecurrent: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetCustomerPaymentMethodDataResponse) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodCardResponse != nil {
		return json.Marshal(&src.PaymentMethodCardResponse)
	}

	if src.PaymentMethodCashResponse != nil {
		return json.Marshal(&src.PaymentMethodCashResponse)
	}

	if src.PaymentMethodSpeiRecurrent != nil {
		return json.Marshal(&src.PaymentMethodSpeiRecurrent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetCustomerPaymentMethodDataResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodCardResponse != nil {
		return obj.PaymentMethodCardResponse
	}

	if obj.PaymentMethodCashResponse != nil {
		return obj.PaymentMethodCashResponse
	}

	if obj.PaymentMethodSpeiRecurrent != nil {
		return obj.PaymentMethodSpeiRecurrent
	}

	// all schemas are nil
	return nil
}

type NullableGetCustomerPaymentMethodDataResponse struct {
	value *GetCustomerPaymentMethodDataResponse
	isSet bool
}

func (v NullableGetCustomerPaymentMethodDataResponse) Get() *GetCustomerPaymentMethodDataResponse {
	return v.value
}

func (v *NullableGetCustomerPaymentMethodDataResponse) Set(val *GetCustomerPaymentMethodDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerPaymentMethodDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerPaymentMethodDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerPaymentMethodDataResponse(val *GetCustomerPaymentMethodDataResponse) *NullableGetCustomerPaymentMethodDataResponse {
	return &NullableGetCustomerPaymentMethodDataResponse{value: val, isSet: true}
}

func (v NullableGetCustomerPaymentMethodDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomerPaymentMethodDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

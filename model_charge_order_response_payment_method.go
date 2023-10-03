package conekta

import (
	"encoding/json"
	"fmt"
)

// ChargeOrderResponsePaymentMethod - struct for ChargeOrderResponsePaymentMethod
type ChargeOrderResponsePaymentMethod struct {
	PaymentMethodBankTransfer *PaymentMethodBankTransfer
	PaymentMethodCard         *PaymentMethodCard
	PaymentMethodCash         *PaymentMethodCash
}

// PaymentMethodBankTransferAsChargeOrderResponsePaymentMethod is a convenience function that returns PaymentMethodBankTransfer wrapped in ChargeOrderResponsePaymentMethod
func PaymentMethodBankTransferAsChargeOrderResponsePaymentMethod(v *PaymentMethodBankTransfer) ChargeOrderResponsePaymentMethod {
	return ChargeOrderResponsePaymentMethod{
		PaymentMethodBankTransfer: v,
	}
}

// PaymentMethodCardAsChargeOrderResponsePaymentMethod is a convenience function that returns PaymentMethodCard wrapped in ChargeOrderResponsePaymentMethod
func PaymentMethodCardAsChargeOrderResponsePaymentMethod(v *PaymentMethodCard) ChargeOrderResponsePaymentMethod {
	return ChargeOrderResponsePaymentMethod{
		PaymentMethodCard: v,
	}
}

// PaymentMethodCashAsChargeOrderResponsePaymentMethod is a convenience function that returns PaymentMethodCash wrapped in ChargeOrderResponsePaymentMethod
func PaymentMethodCashAsChargeOrderResponsePaymentMethod(v *PaymentMethodCash) ChargeOrderResponsePaymentMethod {
	return ChargeOrderResponsePaymentMethod{
		PaymentMethodCash: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChargeOrderResponsePaymentMethod) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'bank_transfer_payment'
	if jsonDict["object"] == "bank_transfer_payment" {
		// try to unmarshal JSON data into PaymentMethodBankTransfer
		err = json.Unmarshal(data, &dst.PaymentMethodBankTransfer)
		if err == nil {
			return nil // data stored in dst.PaymentMethodBankTransfer, return on the first match
		} else {
			dst.PaymentMethodBankTransfer = nil
			return fmt.Errorf("failed to unmarshal ChargeOrderResponsePaymentMethod as PaymentMethodBankTransfer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'card_payment'
	if jsonDict["object"] == "card_payment" {
		// try to unmarshal JSON data into PaymentMethodCard
		err = json.Unmarshal(data, &dst.PaymentMethodCard)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCard, return on the first match
		} else {
			dst.PaymentMethodCard = nil
			return fmt.Errorf("failed to unmarshal ChargeOrderResponsePaymentMethod as PaymentMethodCard: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cash_payment'
	if jsonDict["object"] == "cash_payment" {
		// try to unmarshal JSON data into PaymentMethodCash
		err = json.Unmarshal(data, &dst.PaymentMethodCash)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCash, return on the first match
		} else {
			dst.PaymentMethodCash = nil
			return fmt.Errorf("failed to unmarshal ChargeOrderResponsePaymentMethod as PaymentMethodCash: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_bank_transfer'
	if jsonDict["object"] == "payment_method_bank_transfer" {
		// try to unmarshal JSON data into PaymentMethodBankTransfer
		err = json.Unmarshal(data, &dst.PaymentMethodBankTransfer)
		if err == nil {
			return nil // data stored in dst.PaymentMethodBankTransfer, return on the first match
		} else {
			dst.PaymentMethodBankTransfer = nil
			return fmt.Errorf("failed to unmarshal ChargeOrderResponsePaymentMethod as PaymentMethodBankTransfer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_card'
	if jsonDict["object"] == "payment_method_card" {
		// try to unmarshal JSON data into PaymentMethodCard
		err = json.Unmarshal(data, &dst.PaymentMethodCard)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCard, return on the first match
		} else {
			dst.PaymentMethodCard = nil
			return fmt.Errorf("failed to unmarshal ChargeOrderResponsePaymentMethod as PaymentMethodCard: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_cash'
	if jsonDict["object"] == "payment_method_cash" {
		// try to unmarshal JSON data into PaymentMethodCash
		err = json.Unmarshal(data, &dst.PaymentMethodCash)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCash, return on the first match
		} else {
			dst.PaymentMethodCash = nil
			return fmt.Errorf("failed to unmarshal ChargeOrderResponsePaymentMethod as PaymentMethodCash: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChargeOrderResponsePaymentMethod) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodBankTransfer != nil {
		return json.Marshal(&src.PaymentMethodBankTransfer)
	}

	if src.PaymentMethodCard != nil {
		return json.Marshal(&src.PaymentMethodCard)
	}

	if src.PaymentMethodCash != nil {
		return json.Marshal(&src.PaymentMethodCash)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChargeOrderResponsePaymentMethod) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodBankTransfer != nil {
		return obj.PaymentMethodBankTransfer
	}

	if obj.PaymentMethodCard != nil {
		return obj.PaymentMethodCard
	}

	if obj.PaymentMethodCash != nil {
		return obj.PaymentMethodCash
	}

	// all schemas are nil
	return nil
}

type NullableChargeOrderResponsePaymentMethod struct {
	value *ChargeOrderResponsePaymentMethod
	isSet bool
}

func (v NullableChargeOrderResponsePaymentMethod) Get() *ChargeOrderResponsePaymentMethod {
	return v.value
}

func (v *NullableChargeOrderResponsePaymentMethod) Set(val *ChargeOrderResponsePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeOrderResponsePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeOrderResponsePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeOrderResponsePaymentMethod(val *ChargeOrderResponsePaymentMethod) *NullableChargeOrderResponsePaymentMethod {
	return &NullableChargeOrderResponsePaymentMethod{value: val, isSet: true}
}

func (v NullableChargeOrderResponsePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeOrderResponsePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
	"fmt"
)

// CheckoutOrderTemplateCustomerInfo - It is the information of the customer who will be created when receiving a new payment.
type CheckoutOrderTemplateCustomerInfo struct {
	CustomerInfo               *CustomerInfo
	CustomerInfoJustCustomerId *CustomerInfoJustCustomerId
}

// CustomerInfoAsCheckoutOrderTemplateCustomerInfo is a convenience function that returns CustomerInfo wrapped in CheckoutOrderTemplateCustomerInfo
func CustomerInfoAsCheckoutOrderTemplateCustomerInfo(v *CustomerInfo) CheckoutOrderTemplateCustomerInfo {
	return CheckoutOrderTemplateCustomerInfo{
		CustomerInfo: v,
	}
}

// CustomerInfoJustCustomerIdAsCheckoutOrderTemplateCustomerInfo is a convenience function that returns CustomerInfoJustCustomerId wrapped in CheckoutOrderTemplateCustomerInfo
func CustomerInfoJustCustomerIdAsCheckoutOrderTemplateCustomerInfo(v *CustomerInfoJustCustomerId) CheckoutOrderTemplateCustomerInfo {
	return CheckoutOrderTemplateCustomerInfo{
		CustomerInfoJustCustomerId: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CheckoutOrderTemplateCustomerInfo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomerInfo
	err = json.Unmarshal(data, &dst.CustomerInfo)
	if err == nil {
		jsonCustomerInfo, _ := json.Marshal(dst.CustomerInfo)
		if string(jsonCustomerInfo) == "{}" { // empty struct
			dst.CustomerInfo = nil
		} else {
			match++
		}
	} else {
		dst.CustomerInfo = nil
	}

	// try to unmarshal data into CustomerInfoJustCustomerId
	err = json.Unmarshal(data, &dst.CustomerInfoJustCustomerId)
	if err == nil {
		jsonCustomerInfoJustCustomerId, _ := json.Marshal(dst.CustomerInfoJustCustomerId)
		if string(jsonCustomerInfoJustCustomerId) == "{}" { // empty struct
			dst.CustomerInfoJustCustomerId = nil
		} else {
			match++
		}
	} else {
		dst.CustomerInfoJustCustomerId = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomerInfo = nil
		dst.CustomerInfoJustCustomerId = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CheckoutOrderTemplateCustomerInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CheckoutOrderTemplateCustomerInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CheckoutOrderTemplateCustomerInfo) MarshalJSON() ([]byte, error) {
	if src.CustomerInfo != nil {
		return json.Marshal(&src.CustomerInfo)
	}

	if src.CustomerInfoJustCustomerId != nil {
		return json.Marshal(&src.CustomerInfoJustCustomerId)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CheckoutOrderTemplateCustomerInfo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomerInfo != nil {
		return obj.CustomerInfo
	}

	if obj.CustomerInfoJustCustomerId != nil {
		return obj.CustomerInfoJustCustomerId
	}

	// all schemas are nil
	return nil
}

type NullableCheckoutOrderTemplateCustomerInfo struct {
	value *CheckoutOrderTemplateCustomerInfo
	isSet bool
}

func (v NullableCheckoutOrderTemplateCustomerInfo) Get() *CheckoutOrderTemplateCustomerInfo {
	return v.value
}

func (v *NullableCheckoutOrderTemplateCustomerInfo) Set(val *CheckoutOrderTemplateCustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutOrderTemplateCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutOrderTemplateCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutOrderTemplateCustomerInfo(val *CheckoutOrderTemplateCustomerInfo) *NullableCheckoutOrderTemplateCustomerInfo {
	return &NullableCheckoutOrderTemplateCustomerInfo{value: val, isSet: true}
}

func (v NullableCheckoutOrderTemplateCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutOrderTemplateCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

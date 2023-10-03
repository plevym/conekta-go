package conekta

import (
	"encoding/json"
	"fmt"
)

// OrderRequestCustomerInfo - Customer information
type OrderRequestCustomerInfo struct {
	CustomerInfo               *CustomerInfo
	CustomerInfoJustCustomerId *CustomerInfoJustCustomerId
}

// CustomerInfoAsOrderRequestCustomerInfo is a convenience function that returns CustomerInfo wrapped in OrderRequestCustomerInfo
func CustomerInfoAsOrderRequestCustomerInfo(v *CustomerInfo) OrderRequestCustomerInfo {
	return OrderRequestCustomerInfo{
		CustomerInfo: v,
	}
}

// CustomerInfoJustCustomerIdAsOrderRequestCustomerInfo is a convenience function that returns CustomerInfoJustCustomerId wrapped in OrderRequestCustomerInfo
func CustomerInfoJustCustomerIdAsOrderRequestCustomerInfo(v *CustomerInfoJustCustomerId) OrderRequestCustomerInfo {
	return OrderRequestCustomerInfo{
		CustomerInfoJustCustomerId: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OrderRequestCustomerInfo) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(OrderRequestCustomerInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OrderRequestCustomerInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OrderRequestCustomerInfo) MarshalJSON() ([]byte, error) {
	if src.CustomerInfo != nil {
		return json.Marshal(&src.CustomerInfo)
	}

	if src.CustomerInfoJustCustomerId != nil {
		return json.Marshal(&src.CustomerInfoJustCustomerId)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OrderRequestCustomerInfo) GetActualInstance() interface{} {
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

type NullableOrderRequestCustomerInfo struct {
	value *OrderRequestCustomerInfo
	isSet bool
}

func (v NullableOrderRequestCustomerInfo) Get() *OrderRequestCustomerInfo {
	return v.value
}

func (v *NullableOrderRequestCustomerInfo) Set(val *OrderRequestCustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRequestCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRequestCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRequestCustomerInfo(val *OrderRequestCustomerInfo) *NullableOrderRequestCustomerInfo {
	return &NullableOrderRequestCustomerInfo{value: val, isSet: true}
}

func (v NullableOrderRequestCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRequestCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
	"fmt"
)

// OrderUpdateRequestCustomerInfo - struct for OrderUpdateRequestCustomerInfo
type OrderUpdateRequestCustomerInfo struct {
	CustomerInfo               *CustomerInfo
	CustomerInfoJustCustomerId *CustomerInfoJustCustomerId
}

// CustomerInfoAsOrderUpdateRequestCustomerInfo is a convenience function that returns CustomerInfo wrapped in OrderUpdateRequestCustomerInfo
func CustomerInfoAsOrderUpdateRequestCustomerInfo(v *CustomerInfo) OrderUpdateRequestCustomerInfo {
	return OrderUpdateRequestCustomerInfo{
		CustomerInfo: v,
	}
}

// CustomerInfoJustCustomerIdAsOrderUpdateRequestCustomerInfo is a convenience function that returns CustomerInfoJustCustomerId wrapped in OrderUpdateRequestCustomerInfo
func CustomerInfoJustCustomerIdAsOrderUpdateRequestCustomerInfo(v *CustomerInfoJustCustomerId) OrderUpdateRequestCustomerInfo {
	return OrderUpdateRequestCustomerInfo{
		CustomerInfoJustCustomerId: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OrderUpdateRequestCustomerInfo) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(OrderUpdateRequestCustomerInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OrderUpdateRequestCustomerInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OrderUpdateRequestCustomerInfo) MarshalJSON() ([]byte, error) {
	if src.CustomerInfo != nil {
		return json.Marshal(&src.CustomerInfo)
	}

	if src.CustomerInfoJustCustomerId != nil {
		return json.Marshal(&src.CustomerInfoJustCustomerId)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OrderUpdateRequestCustomerInfo) GetActualInstance() interface{} {
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

type NullableOrderUpdateRequestCustomerInfo struct {
	value *OrderUpdateRequestCustomerInfo
	isSet bool
}

func (v NullableOrderUpdateRequestCustomerInfo) Get() *OrderUpdateRequestCustomerInfo {
	return v.value
}

func (v *NullableOrderUpdateRequestCustomerInfo) Set(val *OrderUpdateRequestCustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUpdateRequestCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUpdateRequestCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUpdateRequestCustomerInfo(val *OrderUpdateRequestCustomerInfo) *NullableOrderUpdateRequestCustomerInfo {
	return &NullableOrderUpdateRequestCustomerInfo{value: val, isSet: true}
}

func (v NullableOrderUpdateRequestCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUpdateRequestCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

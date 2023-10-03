package conekta

import (
	"encoding/json"
)

// checks if the ChargeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeRequest{}

// ChargeRequest The charges to be made
type ChargeRequest struct {
	Amount *int32 `json:"amount,omitempty"`
	// How many months without interest to apply, it can be 3, 6, 9, 12 or 18
	MonthlyInstallments *int32                     `json:"monthly_installments,omitempty"`
	PaymentMethod       ChargeRequestPaymentMethod `json:"payment_method"`
	// Custom reference to add to the charge
	ReferenceId *string `json:"reference_id,omitempty"`
}

// NewChargeRequest instantiates a new ChargeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeRequest(paymentMethod ChargeRequestPaymentMethod) *ChargeRequest {
	this := ChargeRequest{}
	this.PaymentMethod = paymentMethod
	return &this
}

// NewChargeRequestWithDefaults instantiates a new ChargeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeRequestWithDefaults() *ChargeRequest {
	this := ChargeRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ChargeRequest) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequest) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ChargeRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ChargeRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetMonthlyInstallments returns the MonthlyInstallments field value if set, zero value otherwise.
func (o *ChargeRequest) GetMonthlyInstallments() int32 {
	if o == nil || IsNil(o.MonthlyInstallments) {
		var ret int32
		return ret
	}
	return *o.MonthlyInstallments
}

// GetMonthlyInstallmentsOk returns a tuple with the MonthlyInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequest) GetMonthlyInstallmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.MonthlyInstallments) {
		return nil, false
	}
	return o.MonthlyInstallments, true
}

// HasMonthlyInstallments returns a boolean if a field has been set.
func (o *ChargeRequest) HasMonthlyInstallments() bool {
	if o != nil && !IsNil(o.MonthlyInstallments) {
		return true
	}

	return false
}

// SetMonthlyInstallments gets a reference to the given int32 and assigns it to the MonthlyInstallments field.
func (o *ChargeRequest) SetMonthlyInstallments(v int32) {
	o.MonthlyInstallments = &v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *ChargeRequest) GetPaymentMethod() ChargeRequestPaymentMethod {
	if o == nil {
		var ret ChargeRequestPaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *ChargeRequest) GetPaymentMethodOk() (*ChargeRequestPaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *ChargeRequest) SetPaymentMethod(v ChargeRequestPaymentMethod) {
	o.PaymentMethod = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *ChargeRequest) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargeRequest) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *ChargeRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

func (o ChargeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.MonthlyInstallments) {
		toSerialize["monthly_installments"] = o.MonthlyInstallments
	}
	toSerialize["payment_method"] = o.PaymentMethod
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	return toSerialize, nil
}

type NullableChargeRequest struct {
	value *ChargeRequest
	isSet bool
}

func (v NullableChargeRequest) Get() *ChargeRequest {
	return v.value
}

func (v *NullableChargeRequest) Set(val *ChargeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeRequest(val *ChargeRequest) *NullableChargeRequest {
	return &NullableChargeRequest{value: val, isSet: true}
}

func (v NullableChargeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

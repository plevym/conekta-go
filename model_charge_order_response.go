package conekta

import (
	"encoding/json"
)

// checks if the ChargeOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeOrderResponse{}

// ChargeOrderResponse struct for ChargeOrderResponse
type ChargeOrderResponse struct {
	Amount              *int32                            `json:"amount,omitempty"`
	Channel             *ChargeResponseChannel            `json:"channel,omitempty"`
	CreatedAt           *int64                            `json:"created_at,omitempty"`
	Currency            *string                           `json:"currency,omitempty"`
	CustomerId          *string                           `json:"customer_id,omitempty"`
	Description         *string                           `json:"description,omitempty"`
	DeviceFingerprint   NullableString                    `json:"device_fingerprint,omitempty"`
	FailureCode         *string                           `json:"failure_code,omitempty"`
	FailureMessage      *string                           `json:"failure_message,omitempty"`
	Fee                 *int32                            `json:"fee,omitempty"`
	Id                  *string                           `json:"id,omitempty"`
	Livemode            *bool                             `json:"livemode,omitempty"`
	MonthlyInstallments NullableInt32                     `json:"monthly_installments,omitempty"`
	Object              *string                           `json:"object,omitempty"`
	OrderId             *string                           `json:"order_id,omitempty"`
	PaidAt              NullableInt32                     `json:"paid_at,omitempty"`
	PaymentMethod       *ChargeOrderResponsePaymentMethod `json:"payment_method,omitempty"`
	// Reference ID of the charge
	ReferenceId NullableString           `json:"reference_id,omitempty"`
	Refunds     []map[string]interface{} `json:"refunds,omitempty"`
	Status      *string                  `json:"status,omitempty"`
}

// NewChargeOrderResponse instantiates a new ChargeOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeOrderResponse() *ChargeOrderResponse {
	this := ChargeOrderResponse{}
	return &this
}

// NewChargeOrderResponseWithDefaults instantiates a new ChargeOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeOrderResponseWithDefaults() *ChargeOrderResponse {
	this := ChargeOrderResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ChargeOrderResponse) SetAmount(v int32) {
	o.Amount = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetChannel() ChargeResponseChannel {
	if o == nil || IsNil(o.Channel) {
		var ret ChargeResponseChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetChannelOk() (*ChargeResponseChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given ChargeResponseChannel and assigns it to the Channel field.
func (o *ChargeOrderResponse) SetChannel(v ChargeResponseChannel) {
	o.Channel = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ChargeOrderResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ChargeOrderResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *ChargeOrderResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChargeOrderResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceFingerprint returns the DeviceFingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeOrderResponse) GetDeviceFingerprint() string {
	if o == nil || IsNil(o.DeviceFingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceFingerprint.Get()
}

// GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeOrderResponse) GetDeviceFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceFingerprint.Get(), o.DeviceFingerprint.IsSet()
}

// HasDeviceFingerprint returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasDeviceFingerprint() bool {
	if o != nil && o.DeviceFingerprint.IsSet() {
		return true
	}

	return false
}

// SetDeviceFingerprint gets a reference to the given NullableString and assigns it to the DeviceFingerprint field.
func (o *ChargeOrderResponse) SetDeviceFingerprint(v string) {
	o.DeviceFingerprint.Set(&v)
}

// SetDeviceFingerprintNil sets the value for DeviceFingerprint to be an explicit nil
func (o *ChargeOrderResponse) SetDeviceFingerprintNil() {
	o.DeviceFingerprint.Set(nil)
}

// UnsetDeviceFingerprint ensures that no value is present for DeviceFingerprint, not even an explicit nil
func (o *ChargeOrderResponse) UnsetDeviceFingerprint() {
	o.DeviceFingerprint.Unset()
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetFailureCode() string {
	if o == nil || IsNil(o.FailureCode) {
		var ret string
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetFailureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FailureCode) {
		return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasFailureCode() bool {
	if o != nil && !IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given string and assigns it to the FailureCode field.
func (o *ChargeOrderResponse) SetFailureCode(v string) {
	o.FailureCode = &v
}

// GetFailureMessage returns the FailureMessage field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetFailureMessage() string {
	if o == nil || IsNil(o.FailureMessage) {
		var ret string
		return ret
	}
	return *o.FailureMessage
}

// GetFailureMessageOk returns a tuple with the FailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FailureMessage) {
		return nil, false
	}
	return o.FailureMessage, true
}

// HasFailureMessage returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasFailureMessage() bool {
	if o != nil && !IsNil(o.FailureMessage) {
		return true
	}

	return false
}

// SetFailureMessage gets a reference to the given string and assigns it to the FailureMessage field.
func (o *ChargeOrderResponse) SetFailureMessage(v string) {
	o.FailureMessage = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetFee() int32 {
	if o == nil || IsNil(o.Fee) {
		var ret int32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetFeeOk() (*int32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given int32 and assigns it to the Fee field.
func (o *ChargeOrderResponse) SetFee(v int32) {
	o.Fee = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChargeOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *ChargeOrderResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMonthlyInstallments returns the MonthlyInstallments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeOrderResponse) GetMonthlyInstallments() int32 {
	if o == nil || IsNil(o.MonthlyInstallments.Get()) {
		var ret int32
		return ret
	}
	return *o.MonthlyInstallments.Get()
}

// GetMonthlyInstallmentsOk returns a tuple with the MonthlyInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeOrderResponse) GetMonthlyInstallmentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyInstallments.Get(), o.MonthlyInstallments.IsSet()
}

// HasMonthlyInstallments returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasMonthlyInstallments() bool {
	if o != nil && o.MonthlyInstallments.IsSet() {
		return true
	}

	return false
}

// SetMonthlyInstallments gets a reference to the given NullableInt32 and assigns it to the MonthlyInstallments field.
func (o *ChargeOrderResponse) SetMonthlyInstallments(v int32) {
	o.MonthlyInstallments.Set(&v)
}

// SetMonthlyInstallmentsNil sets the value for MonthlyInstallments to be an explicit nil
func (o *ChargeOrderResponse) SetMonthlyInstallmentsNil() {
	o.MonthlyInstallments.Set(nil)
}

// UnsetMonthlyInstallments ensures that no value is present for MonthlyInstallments, not even an explicit nil
func (o *ChargeOrderResponse) UnsetMonthlyInstallments() {
	o.MonthlyInstallments.Unset()
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ChargeOrderResponse) SetObject(v string) {
	o.Object = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *ChargeOrderResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetPaidAt returns the PaidAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeOrderResponse) GetPaidAt() int32 {
	if o == nil || IsNil(o.PaidAt.Get()) {
		var ret int32
		return ret
	}
	return *o.PaidAt.Get()
}

// GetPaidAtOk returns a tuple with the PaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeOrderResponse) GetPaidAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAt.Get(), o.PaidAt.IsSet()
}

// HasPaidAt returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasPaidAt() bool {
	if o != nil && o.PaidAt.IsSet() {
		return true
	}

	return false
}

// SetPaidAt gets a reference to the given NullableInt32 and assigns it to the PaidAt field.
func (o *ChargeOrderResponse) SetPaidAt(v int32) {
	o.PaidAt.Set(&v)
}

// SetPaidAtNil sets the value for PaidAt to be an explicit nil
func (o *ChargeOrderResponse) SetPaidAtNil() {
	o.PaidAt.Set(nil)
}

// UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
func (o *ChargeOrderResponse) UnsetPaidAt() {
	o.PaidAt.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetPaymentMethod() ChargeOrderResponsePaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret ChargeOrderResponsePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetPaymentMethodOk() (*ChargeOrderResponsePaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ChargeOrderResponsePaymentMethod and assigns it to the PaymentMethod field.
func (o *ChargeOrderResponse) SetPaymentMethod(v ChargeOrderResponsePaymentMethod) {
	o.PaymentMethod = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeOrderResponse) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeOrderResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *ChargeOrderResponse) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *ChargeOrderResponse) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *ChargeOrderResponse) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetRefunds() []map[string]interface{} {
	if o == nil || IsNil(o.Refunds) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetRefundsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Refunds) {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasRefunds() bool {
	if o != nil && !IsNil(o.Refunds) {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given []map[string]interface{} and assigns it to the Refunds field.
func (o *ChargeOrderResponse) SetRefunds(v []map[string]interface{}) {
	o.Refunds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChargeOrderResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChargeOrderResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ChargeOrderResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ChargeOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.DeviceFingerprint.IsSet() {
		toSerialize["device_fingerprint"] = o.DeviceFingerprint.Get()
	}
	if !IsNil(o.FailureCode) {
		toSerialize["failure_code"] = o.FailureCode
	}
	if !IsNil(o.FailureMessage) {
		toSerialize["failure_message"] = o.FailureMessage
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if o.MonthlyInstallments.IsSet() {
		toSerialize["monthly_installments"] = o.MonthlyInstallments.Get()
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if o.PaidAt.IsSet() {
		toSerialize["paid_at"] = o.PaidAt.Get()
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if !IsNil(o.Refunds) {
		toSerialize["refunds"] = o.Refunds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChargeOrderResponse struct {
	value *ChargeOrderResponse
	isSet bool
}

func (v NullableChargeOrderResponse) Get() *ChargeOrderResponse {
	return v.value
}

func (v *NullableChargeOrderResponse) Set(val *ChargeOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeOrderResponse(val *ChargeOrderResponse) *NullableChargeOrderResponse {
	return &NullableChargeOrderResponse{value: val, isSet: true}
}

func (v NullableChargeOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

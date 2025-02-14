package conekta

import (
	"encoding/json"
)

// checks if the ChargeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeResponse{}

// ChargeResponse struct for ChargeResponse
type ChargeResponse struct {
	Amount            *int32                       `json:"amount,omitempty"`
	Channel           *ChargeResponseChannel       `json:"channel,omitempty"`
	CreatedAt         *int64                       `json:"created_at,omitempty"`
	Currency          *string                      `json:"currency,omitempty"`
	CustomerId        *string                      `json:"customer_id,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	DeviceFingerprint *string                      `json:"device_fingerprint,omitempty"`
	FailureCode       *string                      `json:"failure_code,omitempty"`
	FailureMessage    *string                      `json:"failure_message,omitempty"`
	Fee               *int32                       `json:"fee,omitempty"`
	Id                *string                      `json:"id,omitempty"`
	Livemode          *bool                        `json:"livemode,omitempty"`
	Object            *string                      `json:"object,omitempty"`
	OrderId           *string                      `json:"order_id,omitempty"`
	PaidAt            NullableInt32                `json:"paid_at,omitempty"`
	PaymentMethod     *ChargeResponsePaymentMethod `json:"payment_method,omitempty"`
	// Reference ID of the charge
	ReferenceId NullableString                `json:"reference_id,omitempty"`
	Refunds     NullableChargeResponseRefunds `json:"refunds,omitempty"`
	Status      *string                       `json:"status,omitempty"`
}

// NewChargeResponse instantiates a new ChargeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeResponse() *ChargeResponse {
	this := ChargeResponse{}
	return &this
}

// NewChargeResponseWithDefaults instantiates a new ChargeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeResponseWithDefaults() *ChargeResponse {
	this := ChargeResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ChargeResponse) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ChargeResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *ChargeResponse) SetAmount(v int32) {
	o.Amount = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ChargeResponse) GetChannel() ChargeResponseChannel {
	if o == nil || IsNil(o.Channel) {
		var ret ChargeResponseChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetChannelOk() (*ChargeResponseChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ChargeResponse) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given ChargeResponseChannel and assigns it to the Channel field.
func (o *ChargeResponse) SetChannel(v ChargeResponseChannel) {
	o.Channel = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChargeResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChargeResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ChargeResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ChargeResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ChargeResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ChargeResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *ChargeResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *ChargeResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *ChargeResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChargeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChargeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChargeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceFingerprint returns the DeviceFingerprint field value if set, zero value otherwise.
func (o *ChargeResponse) GetDeviceFingerprint() string {
	if o == nil || IsNil(o.DeviceFingerprint) {
		var ret string
		return ret
	}
	return *o.DeviceFingerprint
}

// GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetDeviceFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceFingerprint) {
		return nil, false
	}
	return o.DeviceFingerprint, true
}

// HasDeviceFingerprint returns a boolean if a field has been set.
func (o *ChargeResponse) HasDeviceFingerprint() bool {
	if o != nil && !IsNil(o.DeviceFingerprint) {
		return true
	}

	return false
}

// SetDeviceFingerprint gets a reference to the given string and assigns it to the DeviceFingerprint field.
func (o *ChargeResponse) SetDeviceFingerprint(v string) {
	o.DeviceFingerprint = &v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *ChargeResponse) GetFailureCode() string {
	if o == nil || IsNil(o.FailureCode) {
		var ret string
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetFailureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FailureCode) {
		return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *ChargeResponse) HasFailureCode() bool {
	if o != nil && !IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given string and assigns it to the FailureCode field.
func (o *ChargeResponse) SetFailureCode(v string) {
	o.FailureCode = &v
}

// GetFailureMessage returns the FailureMessage field value if set, zero value otherwise.
func (o *ChargeResponse) GetFailureMessage() string {
	if o == nil || IsNil(o.FailureMessage) {
		var ret string
		return ret
	}
	return *o.FailureMessage
}

// GetFailureMessageOk returns a tuple with the FailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FailureMessage) {
		return nil, false
	}
	return o.FailureMessage, true
}

// HasFailureMessage returns a boolean if a field has been set.
func (o *ChargeResponse) HasFailureMessage() bool {
	if o != nil && !IsNil(o.FailureMessage) {
		return true
	}

	return false
}

// SetFailureMessage gets a reference to the given string and assigns it to the FailureMessage field.
func (o *ChargeResponse) SetFailureMessage(v string) {
	o.FailureMessage = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *ChargeResponse) GetFee() int32 {
	if o == nil || IsNil(o.Fee) {
		var ret int32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetFeeOk() (*int32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *ChargeResponse) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given int32 and assigns it to the Fee field.
func (o *ChargeResponse) SetFee(v int32) {
	o.Fee = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChargeResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *ChargeResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *ChargeResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *ChargeResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ChargeResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ChargeResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ChargeResponse) SetObject(v string) {
	o.Object = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ChargeResponse) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ChargeResponse) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *ChargeResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetPaidAt returns the PaidAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeResponse) GetPaidAt() int32 {
	if o == nil || IsNil(o.PaidAt.Get()) {
		var ret int32
		return ret
	}
	return *o.PaidAt.Get()
}

// GetPaidAtOk returns a tuple with the PaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeResponse) GetPaidAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAt.Get(), o.PaidAt.IsSet()
}

// HasPaidAt returns a boolean if a field has been set.
func (o *ChargeResponse) HasPaidAt() bool {
	if o != nil && o.PaidAt.IsSet() {
		return true
	}

	return false
}

// SetPaidAt gets a reference to the given NullableInt32 and assigns it to the PaidAt field.
func (o *ChargeResponse) SetPaidAt(v int32) {
	o.PaidAt.Set(&v)
}

// SetPaidAtNil sets the value for PaidAt to be an explicit nil
func (o *ChargeResponse) SetPaidAtNil() {
	o.PaidAt.Set(nil)
}

// UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
func (o *ChargeResponse) UnsetPaidAt() {
	o.PaidAt.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ChargeResponse) GetPaymentMethod() ChargeResponsePaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret ChargeResponsePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetPaymentMethodOk() (*ChargeResponsePaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ChargeResponse) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ChargeResponsePaymentMethod and assigns it to the PaymentMethod field.
func (o *ChargeResponse) SetPaymentMethod(v ChargeResponsePaymentMethod) {
	o.PaymentMethod = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeResponse) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargeResponse) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *ChargeResponse) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *ChargeResponse) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *ChargeResponse) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetRefunds returns the Refunds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeResponse) GetRefunds() ChargeResponseRefunds {
	if o == nil || IsNil(o.Refunds.Get()) {
		var ret ChargeResponseRefunds
		return ret
	}
	return *o.Refunds.Get()
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeResponse) GetRefundsOk() (*ChargeResponseRefunds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refunds.Get(), o.Refunds.IsSet()
}

// HasRefunds returns a boolean if a field has been set.
func (o *ChargeResponse) HasRefunds() bool {
	if o != nil && o.Refunds.IsSet() {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given NullableChargeResponseRefunds and assigns it to the Refunds field.
func (o *ChargeResponse) SetRefunds(v ChargeResponseRefunds) {
	o.Refunds.Set(&v)
}

// SetRefundsNil sets the value for Refunds to be an explicit nil
func (o *ChargeResponse) SetRefundsNil() {
	o.Refunds.Set(nil)
}

// UnsetRefunds ensures that no value is present for Refunds, not even an explicit nil
func (o *ChargeResponse) UnsetRefunds() {
	o.Refunds.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChargeResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChargeResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ChargeResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ChargeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeviceFingerprint) {
		toSerialize["device_fingerprint"] = o.DeviceFingerprint
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
	if o.Refunds.IsSet() {
		toSerialize["refunds"] = o.Refunds.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChargeResponse struct {
	value *ChargeResponse
	isSet bool
}

func (v NullableChargeResponse) Get() *ChargeResponse {
	return v.value
}

func (v *NullableChargeResponse) Set(val *ChargeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeResponse(val *ChargeResponse) *NullableChargeResponse {
	return &NullableChargeResponse{value: val, isSet: true}
}

func (v NullableChargeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

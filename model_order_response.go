package conekta

import (
	"encoding/json"
)

// checks if the OrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponse{}

// OrderResponse order response
type OrderResponse struct {
	// The total amount to be collected in cents
	Amount *int32 `json:"amount,omitempty"`
	// The total amount refunded in cents
	AmountRefunded *int32                 `json:"amount_refunded,omitempty"`
	Channel        *ChargeResponseChannel `json:"channel,omitempty"`
	Charges        *OrderResponseCharges  `json:"charges,omitempty"`
	Checkout       *OrderResponseCheckout `json:"checkout,omitempty"`
	// The time at which the object was created in seconds since the Unix epoch
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The three-letter ISO 4217 currency code. The currency of the order.
	Currency      *string                     `json:"currency,omitempty"`
	CustomerInfo  *OrderResponseCustomerInfo  `json:"customer_info,omitempty"`
	DiscountLines *OrderResponseDiscountLines `json:"discount_lines,omitempty"`
	FiscalEntity  *OrderResponseFiscalEntity  `json:"fiscal_entity,omitempty"`
	Id            *string                     `json:"id,omitempty"`
	IsRefundable  *bool                       `json:"is_refundable,omitempty"`
	LineItems     *OrderResponseProducts      `json:"line_items,omitempty"`
	// Whether the object exists in live mode or test mode
	Livemode *bool `json:"livemode,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// String representing the object’s type. Objects of the same type share the same value.
	Object *string `json:"object,omitempty"`
	// The payment status of the order.
	PaymentStatus *string `json:"payment_status,omitempty"`
	// Indicates the processing mode for the order, either ecommerce, recurrent or validation.
	ProcessingMode  *string                       `json:"processing_mode,omitempty"`
	ShippingContact *OrderResponseShippingContact `json:"shipping_contact,omitempty"`
	// The time at which the object was last updated in seconds since the Unix epoch
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

// NewOrderResponse instantiates a new OrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponse() *OrderResponse {
	this := OrderResponse{}
	return &this
}

// NewOrderResponseWithDefaults instantiates a new OrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseWithDefaults() *OrderResponse {
	this := OrderResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderResponse) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *OrderResponse) SetAmount(v int32) {
	o.Amount = &v
}

// GetAmountRefunded returns the AmountRefunded field value if set, zero value otherwise.
func (o *OrderResponse) GetAmountRefunded() int32 {
	if o == nil || IsNil(o.AmountRefunded) {
		var ret int32
		return ret
	}
	return *o.AmountRefunded
}

// GetAmountRefundedOk returns a tuple with the AmountRefunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetAmountRefundedOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountRefunded) {
		return nil, false
	}
	return o.AmountRefunded, true
}

// HasAmountRefunded returns a boolean if a field has been set.
func (o *OrderResponse) HasAmountRefunded() bool {
	if o != nil && !IsNil(o.AmountRefunded) {
		return true
	}

	return false
}

// SetAmountRefunded gets a reference to the given int32 and assigns it to the AmountRefunded field.
func (o *OrderResponse) SetAmountRefunded(v int32) {
	o.AmountRefunded = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *OrderResponse) GetChannel() ChargeResponseChannel {
	if o == nil || IsNil(o.Channel) {
		var ret ChargeResponseChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetChannelOk() (*ChargeResponseChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *OrderResponse) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given ChargeResponseChannel and assigns it to the Channel field.
func (o *OrderResponse) SetChannel(v ChargeResponseChannel) {
	o.Channel = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *OrderResponse) GetCharges() OrderResponseCharges {
	if o == nil || IsNil(o.Charges) {
		var ret OrderResponseCharges
		return ret
	}
	return *o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetChargesOk() (*OrderResponseCharges, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *OrderResponse) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given OrderResponseCharges and assigns it to the Charges field.
func (o *OrderResponse) SetCharges(v OrderResponseCharges) {
	o.Charges = &v
}

// GetCheckout returns the Checkout field value if set, zero value otherwise.
func (o *OrderResponse) GetCheckout() OrderResponseCheckout {
	if o == nil || IsNil(o.Checkout) {
		var ret OrderResponseCheckout
		return ret
	}
	return *o.Checkout
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetCheckoutOk() (*OrderResponseCheckout, bool) {
	if o == nil || IsNil(o.Checkout) {
		return nil, false
	}
	return o.Checkout, true
}

// HasCheckout returns a boolean if a field has been set.
func (o *OrderResponse) HasCheckout() bool {
	if o != nil && !IsNil(o.Checkout) {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given OrderResponseCheckout and assigns it to the Checkout field.
func (o *OrderResponse) SetCheckout(v OrderResponseCheckout) {
	o.Checkout = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *OrderResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OrderResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrderResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *OrderResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerInfo returns the CustomerInfo field value if set, zero value otherwise.
func (o *OrderResponse) GetCustomerInfo() OrderResponseCustomerInfo {
	if o == nil || IsNil(o.CustomerInfo) {
		var ret OrderResponseCustomerInfo
		return ret
	}
	return *o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetCustomerInfoOk() (*OrderResponseCustomerInfo, bool) {
	if o == nil || IsNil(o.CustomerInfo) {
		return nil, false
	}
	return o.CustomerInfo, true
}

// HasCustomerInfo returns a boolean if a field has been set.
func (o *OrderResponse) HasCustomerInfo() bool {
	if o != nil && !IsNil(o.CustomerInfo) {
		return true
	}

	return false
}

// SetCustomerInfo gets a reference to the given OrderResponseCustomerInfo and assigns it to the CustomerInfo field.
func (o *OrderResponse) SetCustomerInfo(v OrderResponseCustomerInfo) {
	o.CustomerInfo = &v
}

// GetDiscountLines returns the DiscountLines field value if set, zero value otherwise.
func (o *OrderResponse) GetDiscountLines() OrderResponseDiscountLines {
	if o == nil || IsNil(o.DiscountLines) {
		var ret OrderResponseDiscountLines
		return ret
	}
	return *o.DiscountLines
}

// GetDiscountLinesOk returns a tuple with the DiscountLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetDiscountLinesOk() (*OrderResponseDiscountLines, bool) {
	if o == nil || IsNil(o.DiscountLines) {
		return nil, false
	}
	return o.DiscountLines, true
}

// HasDiscountLines returns a boolean if a field has been set.
func (o *OrderResponse) HasDiscountLines() bool {
	if o != nil && !IsNil(o.DiscountLines) {
		return true
	}

	return false
}

// SetDiscountLines gets a reference to the given OrderResponseDiscountLines and assigns it to the DiscountLines field.
func (o *OrderResponse) SetDiscountLines(v OrderResponseDiscountLines) {
	o.DiscountLines = &v
}

// GetFiscalEntity returns the FiscalEntity field value if set, zero value otherwise.
func (o *OrderResponse) GetFiscalEntity() OrderResponseFiscalEntity {
	if o == nil || IsNil(o.FiscalEntity) {
		var ret OrderResponseFiscalEntity
		return ret
	}
	return *o.FiscalEntity
}

// GetFiscalEntityOk returns a tuple with the FiscalEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetFiscalEntityOk() (*OrderResponseFiscalEntity, bool) {
	if o == nil || IsNil(o.FiscalEntity) {
		return nil, false
	}
	return o.FiscalEntity, true
}

// HasFiscalEntity returns a boolean if a field has been set.
func (o *OrderResponse) HasFiscalEntity() bool {
	if o != nil && !IsNil(o.FiscalEntity) {
		return true
	}

	return false
}

// SetFiscalEntity gets a reference to the given OrderResponseFiscalEntity and assigns it to the FiscalEntity field.
func (o *OrderResponse) SetFiscalEntity(v OrderResponseFiscalEntity) {
	o.FiscalEntity = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponse) SetId(v string) {
	o.Id = &v
}

// GetIsRefundable returns the IsRefundable field value if set, zero value otherwise.
func (o *OrderResponse) GetIsRefundable() bool {
	if o == nil || IsNil(o.IsRefundable) {
		var ret bool
		return ret
	}
	return *o.IsRefundable
}

// GetIsRefundableOk returns a tuple with the IsRefundable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetIsRefundableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRefundable) {
		return nil, false
	}
	return o.IsRefundable, true
}

// HasIsRefundable returns a boolean if a field has been set.
func (o *OrderResponse) HasIsRefundable() bool {
	if o != nil && !IsNil(o.IsRefundable) {
		return true
	}

	return false
}

// SetIsRefundable gets a reference to the given bool and assigns it to the IsRefundable field.
func (o *OrderResponse) SetIsRefundable(v bool) {
	o.IsRefundable = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *OrderResponse) GetLineItems() OrderResponseProducts {
	if o == nil || IsNil(o.LineItems) {
		var ret OrderResponseProducts
		return ret
	}
	return *o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetLineItemsOk() (*OrderResponseProducts, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *OrderResponse) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given OrderResponseProducts and assigns it to the LineItems field.
func (o *OrderResponse) SetLineItems(v OrderResponseProducts) {
	o.LineItems = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *OrderResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *OrderResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *OrderResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *OrderResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponse) SetObject(v string) {
	o.Object = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *OrderResponse) GetPaymentStatus() string {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret string
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetPaymentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *OrderResponse) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given string and assigns it to the PaymentStatus field.
func (o *OrderResponse) SetPaymentStatus(v string) {
	o.PaymentStatus = &v
}

// GetProcessingMode returns the ProcessingMode field value if set, zero value otherwise.
func (o *OrderResponse) GetProcessingMode() string {
	if o == nil || IsNil(o.ProcessingMode) {
		var ret string
		return ret
	}
	return *o.ProcessingMode
}

// GetProcessingModeOk returns a tuple with the ProcessingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetProcessingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessingMode) {
		return nil, false
	}
	return o.ProcessingMode, true
}

// HasProcessingMode returns a boolean if a field has been set.
func (o *OrderResponse) HasProcessingMode() bool {
	if o != nil && !IsNil(o.ProcessingMode) {
		return true
	}

	return false
}

// SetProcessingMode gets a reference to the given string and assigns it to the ProcessingMode field.
func (o *OrderResponse) SetProcessingMode(v string) {
	o.ProcessingMode = &v
}

// GetShippingContact returns the ShippingContact field value if set, zero value otherwise.
func (o *OrderResponse) GetShippingContact() OrderResponseShippingContact {
	if o == nil || IsNil(o.ShippingContact) {
		var ret OrderResponseShippingContact
		return ret
	}
	return *o.ShippingContact
}

// GetShippingContactOk returns a tuple with the ShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetShippingContactOk() (*OrderResponseShippingContact, bool) {
	if o == nil || IsNil(o.ShippingContact) {
		return nil, false
	}
	return o.ShippingContact, true
}

// HasShippingContact returns a boolean if a field has been set.
func (o *OrderResponse) HasShippingContact() bool {
	if o != nil && !IsNil(o.ShippingContact) {
		return true
	}

	return false
}

// SetShippingContact gets a reference to the given OrderResponseShippingContact and assigns it to the ShippingContact field.
func (o *OrderResponse) SetShippingContact(v OrderResponseShippingContact) {
	o.ShippingContact = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrderResponse) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrderResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *OrderResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o OrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AmountRefunded) {
		toSerialize["amount_refunded"] = o.AmountRefunded
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Checkout) {
		toSerialize["checkout"] = o.Checkout
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CustomerInfo) {
		toSerialize["customer_info"] = o.CustomerInfo
	}
	if !IsNil(o.DiscountLines) {
		toSerialize["discount_lines"] = o.DiscountLines
	}
	if !IsNil(o.FiscalEntity) {
		toSerialize["fiscal_entity"] = o.FiscalEntity
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsRefundable) {
		toSerialize["is_refundable"] = o.IsRefundable
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if !IsNil(o.ProcessingMode) {
		toSerialize["processing_mode"] = o.ProcessingMode
	}
	if !IsNil(o.ShippingContact) {
		toSerialize["shipping_contact"] = o.ShippingContact
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableOrderResponse struct {
	value *OrderResponse
	isSet bool
}

func (v NullableOrderResponse) Get() *OrderResponse {
	return v.value
}

func (v *NullableOrderResponse) Set(val *OrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponse(val *OrderResponse) *NullableOrderResponse {
	return &NullableOrderResponse{value: val, isSet: true}
}

func (v NullableOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

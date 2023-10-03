package conekta

import (
	"encoding/json"
)

// checks if the OrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRequest{}

// OrderRequest a order
type OrderRequest struct {
	// List of [charges](https://developers.conekta.com/v2.1.0/reference/orderscreatecharge) that are applied to the order
	Charges  []ChargeRequest  `json:"charges,omitempty"`
	Checkout *CheckoutRequest `json:"checkout,omitempty"`
	// Currency with which the payment will be made. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217)
	Currency     string                   `json:"currency"`
	CustomerInfo OrderRequestCustomerInfo `json:"customer_info"`
	// List of [discounts](https://developers.conekta.com/v2.1.0/reference/orderscreatediscountline) that are applied to the order. You must have at least one discount.
	DiscountLines []OrderDiscountLinesRequest `json:"discount_lines,omitempty"`
	// List of [products](https://developers.conekta.com/v2.1.0/reference/orderscreateproduct) that are sold in the order. You must have at least one product.
	LineItems []Product `json:"line_items"`
	// Metadata associated with the order
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Allows you to fill out the shipping information at checkout
	NeedsShippingContact *bool `json:"needs_shipping_contact,omitempty"`
	// Indicates whether the order charges must be preauthorized
	PreAuthorize *bool `json:"pre_authorize,omitempty"`
	// Indicates the processing mode for the order, either ecommerce, recurrent or validation.
	ProcessingMode  *string                   `json:"processing_mode,omitempty"`
	ShippingContact *CustomerShippingContacts `json:"shipping_contact,omitempty"`
	// List of [shipping costs](https://developers.conekta.com/v2.1.0/reference/orderscreateshipping). If the online store offers digital products.
	ShippingLines []ShippingRequest `json:"shipping_lines,omitempty"`
	// List of [taxes](https://developers.conekta.com/v2.1.0/reference/orderscreatetaxes) that are applied to the order.
	TaxLines []OrderTaxRequest `json:"tax_lines,omitempty"`
}

// NewOrderRequest instantiates a new OrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRequest(currency string, customerInfo OrderRequestCustomerInfo, lineItems []Product) *OrderRequest {
	this := OrderRequest{}
	this.Currency = currency
	this.CustomerInfo = customerInfo
	this.LineItems = lineItems
	var preAuthorize bool = false
	this.PreAuthorize = &preAuthorize
	return &this
}

// NewOrderRequestWithDefaults instantiates a new OrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRequestWithDefaults() *OrderRequest {
	this := OrderRequest{}
	var preAuthorize bool = false
	this.PreAuthorize = &preAuthorize
	return &this
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *OrderRequest) GetCharges() []ChargeRequest {
	if o == nil || IsNil(o.Charges) {
		var ret []ChargeRequest
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetChargesOk() ([]ChargeRequest, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *OrderRequest) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []ChargeRequest and assigns it to the Charges field.
func (o *OrderRequest) SetCharges(v []ChargeRequest) {
	o.Charges = v
}

// GetCheckout returns the Checkout field value if set, zero value otherwise.
func (o *OrderRequest) GetCheckout() CheckoutRequest {
	if o == nil || IsNil(o.Checkout) {
		var ret CheckoutRequest
		return ret
	}
	return *o.Checkout
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetCheckoutOk() (*CheckoutRequest, bool) {
	if o == nil || IsNil(o.Checkout) {
		return nil, false
	}
	return o.Checkout, true
}

// HasCheckout returns a boolean if a field has been set.
func (o *OrderRequest) HasCheckout() bool {
	if o != nil && !IsNil(o.Checkout) {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given CheckoutRequest and assigns it to the Checkout field.
func (o *OrderRequest) SetCheckout(v CheckoutRequest) {
	o.Checkout = &v
}

// GetCurrency returns the Currency field value
func (o *OrderRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *OrderRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerInfo returns the CustomerInfo field value
func (o *OrderRequest) GetCustomerInfo() OrderRequestCustomerInfo {
	if o == nil {
		var ret OrderRequestCustomerInfo
		return ret
	}

	return o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetCustomerInfoOk() (*OrderRequestCustomerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerInfo, true
}

// SetCustomerInfo sets field value
func (o *OrderRequest) SetCustomerInfo(v OrderRequestCustomerInfo) {
	o.CustomerInfo = v
}

// GetDiscountLines returns the DiscountLines field value if set, zero value otherwise.
func (o *OrderRequest) GetDiscountLines() []OrderDiscountLinesRequest {
	if o == nil || IsNil(o.DiscountLines) {
		var ret []OrderDiscountLinesRequest
		return ret
	}
	return o.DiscountLines
}

// GetDiscountLinesOk returns a tuple with the DiscountLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetDiscountLinesOk() ([]OrderDiscountLinesRequest, bool) {
	if o == nil || IsNil(o.DiscountLines) {
		return nil, false
	}
	return o.DiscountLines, true
}

// HasDiscountLines returns a boolean if a field has been set.
func (o *OrderRequest) HasDiscountLines() bool {
	if o != nil && !IsNil(o.DiscountLines) {
		return true
	}

	return false
}

// SetDiscountLines gets a reference to the given []OrderDiscountLinesRequest and assigns it to the DiscountLines field.
func (o *OrderRequest) SetDiscountLines(v []OrderDiscountLinesRequest) {
	o.DiscountLines = v
}

// GetLineItems returns the LineItems field value
func (o *OrderRequest) GetLineItems() []Product {
	if o == nil {
		var ret []Product
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetLineItemsOk() ([]Product, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *OrderRequest) SetLineItems(v []Product) {
	o.LineItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *OrderRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNeedsShippingContact returns the NeedsShippingContact field value if set, zero value otherwise.
func (o *OrderRequest) GetNeedsShippingContact() bool {
	if o == nil || IsNil(o.NeedsShippingContact) {
		var ret bool
		return ret
	}
	return *o.NeedsShippingContact
}

// GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetNeedsShippingContactOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsShippingContact) {
		return nil, false
	}
	return o.NeedsShippingContact, true
}

// HasNeedsShippingContact returns a boolean if a field has been set.
func (o *OrderRequest) HasNeedsShippingContact() bool {
	if o != nil && !IsNil(o.NeedsShippingContact) {
		return true
	}

	return false
}

// SetNeedsShippingContact gets a reference to the given bool and assigns it to the NeedsShippingContact field.
func (o *OrderRequest) SetNeedsShippingContact(v bool) {
	o.NeedsShippingContact = &v
}

// GetPreAuthorize returns the PreAuthorize field value if set, zero value otherwise.
func (o *OrderRequest) GetPreAuthorize() bool {
	if o == nil || IsNil(o.PreAuthorize) {
		var ret bool
		return ret
	}
	return *o.PreAuthorize
}

// GetPreAuthorizeOk returns a tuple with the PreAuthorize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetPreAuthorizeOk() (*bool, bool) {
	if o == nil || IsNil(o.PreAuthorize) {
		return nil, false
	}
	return o.PreAuthorize, true
}

// HasPreAuthorize returns a boolean if a field has been set.
func (o *OrderRequest) HasPreAuthorize() bool {
	if o != nil && !IsNil(o.PreAuthorize) {
		return true
	}

	return false
}

// SetPreAuthorize gets a reference to the given bool and assigns it to the PreAuthorize field.
func (o *OrderRequest) SetPreAuthorize(v bool) {
	o.PreAuthorize = &v
}

// GetProcessingMode returns the ProcessingMode field value if set, zero value otherwise.
func (o *OrderRequest) GetProcessingMode() string {
	if o == nil || IsNil(o.ProcessingMode) {
		var ret string
		return ret
	}
	return *o.ProcessingMode
}

// GetProcessingModeOk returns a tuple with the ProcessingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetProcessingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessingMode) {
		return nil, false
	}
	return o.ProcessingMode, true
}

// HasProcessingMode returns a boolean if a field has been set.
func (o *OrderRequest) HasProcessingMode() bool {
	if o != nil && !IsNil(o.ProcessingMode) {
		return true
	}

	return false
}

// SetProcessingMode gets a reference to the given string and assigns it to the ProcessingMode field.
func (o *OrderRequest) SetProcessingMode(v string) {
	o.ProcessingMode = &v
}

// GetShippingContact returns the ShippingContact field value if set, zero value otherwise.
func (o *OrderRequest) GetShippingContact() CustomerShippingContacts {
	if o == nil || IsNil(o.ShippingContact) {
		var ret CustomerShippingContacts
		return ret
	}
	return *o.ShippingContact
}

// GetShippingContactOk returns a tuple with the ShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetShippingContactOk() (*CustomerShippingContacts, bool) {
	if o == nil || IsNil(o.ShippingContact) {
		return nil, false
	}
	return o.ShippingContact, true
}

// HasShippingContact returns a boolean if a field has been set.
func (o *OrderRequest) HasShippingContact() bool {
	if o != nil && !IsNil(o.ShippingContact) {
		return true
	}

	return false
}

// SetShippingContact gets a reference to the given CustomerShippingContacts and assigns it to the ShippingContact field.
func (o *OrderRequest) SetShippingContact(v CustomerShippingContacts) {
	o.ShippingContact = &v
}

// GetShippingLines returns the ShippingLines field value if set, zero value otherwise.
func (o *OrderRequest) GetShippingLines() []ShippingRequest {
	if o == nil || IsNil(o.ShippingLines) {
		var ret []ShippingRequest
		return ret
	}
	return o.ShippingLines
}

// GetShippingLinesOk returns a tuple with the ShippingLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetShippingLinesOk() ([]ShippingRequest, bool) {
	if o == nil || IsNil(o.ShippingLines) {
		return nil, false
	}
	return o.ShippingLines, true
}

// HasShippingLines returns a boolean if a field has been set.
func (o *OrderRequest) HasShippingLines() bool {
	if o != nil && !IsNil(o.ShippingLines) {
		return true
	}

	return false
}

// SetShippingLines gets a reference to the given []ShippingRequest and assigns it to the ShippingLines field.
func (o *OrderRequest) SetShippingLines(v []ShippingRequest) {
	o.ShippingLines = v
}

// GetTaxLines returns the TaxLines field value if set, zero value otherwise.
func (o *OrderRequest) GetTaxLines() []OrderTaxRequest {
	if o == nil || IsNil(o.TaxLines) {
		var ret []OrderTaxRequest
		return ret
	}
	return o.TaxLines
}

// GetTaxLinesOk returns a tuple with the TaxLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRequest) GetTaxLinesOk() ([]OrderTaxRequest, bool) {
	if o == nil || IsNil(o.TaxLines) {
		return nil, false
	}
	return o.TaxLines, true
}

// HasTaxLines returns a boolean if a field has been set.
func (o *OrderRequest) HasTaxLines() bool {
	if o != nil && !IsNil(o.TaxLines) {
		return true
	}

	return false
}

// SetTaxLines gets a reference to the given []OrderTaxRequest and assigns it to the TaxLines field.
func (o *OrderRequest) SetTaxLines(v []OrderTaxRequest) {
	o.TaxLines = v
}

func (o OrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Checkout) {
		toSerialize["checkout"] = o.Checkout
	}
	toSerialize["currency"] = o.Currency
	toSerialize["customer_info"] = o.CustomerInfo
	if !IsNil(o.DiscountLines) {
		toSerialize["discount_lines"] = o.DiscountLines
	}
	toSerialize["line_items"] = o.LineItems
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.NeedsShippingContact) {
		toSerialize["needs_shipping_contact"] = o.NeedsShippingContact
	}
	if !IsNil(o.PreAuthorize) {
		toSerialize["pre_authorize"] = o.PreAuthorize
	}
	if !IsNil(o.ProcessingMode) {
		toSerialize["processing_mode"] = o.ProcessingMode
	}
	if !IsNil(o.ShippingContact) {
		toSerialize["shipping_contact"] = o.ShippingContact
	}
	if !IsNil(o.ShippingLines) {
		toSerialize["shipping_lines"] = o.ShippingLines
	}
	if !IsNil(o.TaxLines) {
		toSerialize["tax_lines"] = o.TaxLines
	}
	return toSerialize, nil
}

type NullableOrderRequest struct {
	value *OrderRequest
	isSet bool
}

func (v NullableOrderRequest) Get() *OrderRequest {
	return v.value
}

func (v *NullableOrderRequest) Set(val *OrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRequest(val *OrderRequest) *NullableOrderRequest {
	return &NullableOrderRequest{value: val, isSet: true}
}

func (v NullableOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

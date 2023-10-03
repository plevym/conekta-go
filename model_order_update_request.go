package conekta

import (
	"encoding/json"
)

// checks if the OrderUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderUpdateRequest{}

// OrderUpdateRequest a order
type OrderUpdateRequest struct {
	Charges  []ChargeRequest  `json:"charges,omitempty"`
	Checkout *CheckoutRequest `json:"checkout,omitempty"`
	// Currency with which the payment will be made. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217)
	Currency     *string                         `json:"currency,omitempty"`
	CustomerInfo *OrderUpdateRequestCustomerInfo `json:"customer_info,omitempty"`
	// List of [discounts](https://developers.conekta.com/v2.1.0/reference/orderscreatediscountline) that are applied to the order. You must have at least one discount.
	DiscountLines []OrderDiscountLinesRequest `json:"discount_lines,omitempty"`
	// List of [products](https://developers.conekta.com/v2.1.0/reference/orderscreateproduct) that are sold in the order. You must have at least one product.
	LineItems []Product          `json:"line_items,omitempty"`
	Metadata  *map[string]string `json:"metadata,omitempty"`
	// Indicates whether the order charges must be preauthorized
	PreAuthorize    *bool                     `json:"pre_authorize,omitempty"`
	ShippingContact *CustomerShippingContacts `json:"shipping_contact,omitempty"`
	// List of [shipping costs](https://developers.conekta.com/v2.1.0/reference/orderscreateshipping). If the online store offers digital products.
	ShippingLines []ShippingRequest `json:"shipping_lines,omitempty"`
	TaxLines      []OrderTaxRequest `json:"tax_lines,omitempty"`
}

// NewOrderUpdateRequest instantiates a new OrderUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUpdateRequest() *OrderUpdateRequest {
	this := OrderUpdateRequest{}
	var preAuthorize bool = false
	this.PreAuthorize = &preAuthorize
	return &this
}

// NewOrderUpdateRequestWithDefaults instantiates a new OrderUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUpdateRequestWithDefaults() *OrderUpdateRequest {
	this := OrderUpdateRequest{}
	var preAuthorize bool = false
	this.PreAuthorize = &preAuthorize
	return &this
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetCharges() []ChargeRequest {
	if o == nil || IsNil(o.Charges) {
		var ret []ChargeRequest
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetChargesOk() ([]ChargeRequest, bool) {
	if o == nil || IsNil(o.Charges) {
		return nil, false
	}
	return o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasCharges() bool {
	if o != nil && !IsNil(o.Charges) {
		return true
	}

	return false
}

// SetCharges gets a reference to the given []ChargeRequest and assigns it to the Charges field.
func (o *OrderUpdateRequest) SetCharges(v []ChargeRequest) {
	o.Charges = v
}

// GetCheckout returns the Checkout field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetCheckout() CheckoutRequest {
	if o == nil || IsNil(o.Checkout) {
		var ret CheckoutRequest
		return ret
	}
	return *o.Checkout
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetCheckoutOk() (*CheckoutRequest, bool) {
	if o == nil || IsNil(o.Checkout) {
		return nil, false
	}
	return o.Checkout, true
}

// HasCheckout returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasCheckout() bool {
	if o != nil && !IsNil(o.Checkout) {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given CheckoutRequest and assigns it to the Checkout field.
func (o *OrderUpdateRequest) SetCheckout(v CheckoutRequest) {
	o.Checkout = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *OrderUpdateRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerInfo returns the CustomerInfo field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetCustomerInfo() OrderUpdateRequestCustomerInfo {
	if o == nil || IsNil(o.CustomerInfo) {
		var ret OrderUpdateRequestCustomerInfo
		return ret
	}
	return *o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetCustomerInfoOk() (*OrderUpdateRequestCustomerInfo, bool) {
	if o == nil || IsNil(o.CustomerInfo) {
		return nil, false
	}
	return o.CustomerInfo, true
}

// HasCustomerInfo returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasCustomerInfo() bool {
	if o != nil && !IsNil(o.CustomerInfo) {
		return true
	}

	return false
}

// SetCustomerInfo gets a reference to the given OrderUpdateRequestCustomerInfo and assigns it to the CustomerInfo field.
func (o *OrderUpdateRequest) SetCustomerInfo(v OrderUpdateRequestCustomerInfo) {
	o.CustomerInfo = &v
}

// GetDiscountLines returns the DiscountLines field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetDiscountLines() []OrderDiscountLinesRequest {
	if o == nil || IsNil(o.DiscountLines) {
		var ret []OrderDiscountLinesRequest
		return ret
	}
	return o.DiscountLines
}

// GetDiscountLinesOk returns a tuple with the DiscountLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetDiscountLinesOk() ([]OrderDiscountLinesRequest, bool) {
	if o == nil || IsNil(o.DiscountLines) {
		return nil, false
	}
	return o.DiscountLines, true
}

// HasDiscountLines returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasDiscountLines() bool {
	if o != nil && !IsNil(o.DiscountLines) {
		return true
	}

	return false
}

// SetDiscountLines gets a reference to the given []OrderDiscountLinesRequest and assigns it to the DiscountLines field.
func (o *OrderUpdateRequest) SetDiscountLines(v []OrderDiscountLinesRequest) {
	o.DiscountLines = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetLineItems() []Product {
	if o == nil || IsNil(o.LineItems) {
		var ret []Product
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetLineItemsOk() ([]Product, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []Product and assigns it to the LineItems field.
func (o *OrderUpdateRequest) SetLineItems(v []Product) {
	o.LineItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OrderUpdateRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPreAuthorize returns the PreAuthorize field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetPreAuthorize() bool {
	if o == nil || IsNil(o.PreAuthorize) {
		var ret bool
		return ret
	}
	return *o.PreAuthorize
}

// GetPreAuthorizeOk returns a tuple with the PreAuthorize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetPreAuthorizeOk() (*bool, bool) {
	if o == nil || IsNil(o.PreAuthorize) {
		return nil, false
	}
	return o.PreAuthorize, true
}

// HasPreAuthorize returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasPreAuthorize() bool {
	if o != nil && !IsNil(o.PreAuthorize) {
		return true
	}

	return false
}

// SetPreAuthorize gets a reference to the given bool and assigns it to the PreAuthorize field.
func (o *OrderUpdateRequest) SetPreAuthorize(v bool) {
	o.PreAuthorize = &v
}

// GetShippingContact returns the ShippingContact field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetShippingContact() CustomerShippingContacts {
	if o == nil || IsNil(o.ShippingContact) {
		var ret CustomerShippingContacts
		return ret
	}
	return *o.ShippingContact
}

// GetShippingContactOk returns a tuple with the ShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetShippingContactOk() (*CustomerShippingContacts, bool) {
	if o == nil || IsNil(o.ShippingContact) {
		return nil, false
	}
	return o.ShippingContact, true
}

// HasShippingContact returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasShippingContact() bool {
	if o != nil && !IsNil(o.ShippingContact) {
		return true
	}

	return false
}

// SetShippingContact gets a reference to the given CustomerShippingContacts and assigns it to the ShippingContact field.
func (o *OrderUpdateRequest) SetShippingContact(v CustomerShippingContacts) {
	o.ShippingContact = &v
}

// GetShippingLines returns the ShippingLines field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetShippingLines() []ShippingRequest {
	if o == nil || IsNil(o.ShippingLines) {
		var ret []ShippingRequest
		return ret
	}
	return o.ShippingLines
}

// GetShippingLinesOk returns a tuple with the ShippingLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetShippingLinesOk() ([]ShippingRequest, bool) {
	if o == nil || IsNil(o.ShippingLines) {
		return nil, false
	}
	return o.ShippingLines, true
}

// HasShippingLines returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasShippingLines() bool {
	if o != nil && !IsNil(o.ShippingLines) {
		return true
	}

	return false
}

// SetShippingLines gets a reference to the given []ShippingRequest and assigns it to the ShippingLines field.
func (o *OrderUpdateRequest) SetShippingLines(v []ShippingRequest) {
	o.ShippingLines = v
}

// GetTaxLines returns the TaxLines field value if set, zero value otherwise.
func (o *OrderUpdateRequest) GetTaxLines() []OrderTaxRequest {
	if o == nil || IsNil(o.TaxLines) {
		var ret []OrderTaxRequest
		return ret
	}
	return o.TaxLines
}

// GetTaxLinesOk returns a tuple with the TaxLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateRequest) GetTaxLinesOk() ([]OrderTaxRequest, bool) {
	if o == nil || IsNil(o.TaxLines) {
		return nil, false
	}
	return o.TaxLines, true
}

// HasTaxLines returns a boolean if a field has been set.
func (o *OrderUpdateRequest) HasTaxLines() bool {
	if o != nil && !IsNil(o.TaxLines) {
		return true
	}

	return false
}

// SetTaxLines gets a reference to the given []OrderTaxRequest and assigns it to the TaxLines field.
func (o *OrderUpdateRequest) SetTaxLines(v []OrderTaxRequest) {
	o.TaxLines = v
}

func (o OrderUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Charges) {
		toSerialize["charges"] = o.Charges
	}
	if !IsNil(o.Checkout) {
		toSerialize["checkout"] = o.Checkout
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
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PreAuthorize) {
		toSerialize["pre_authorize"] = o.PreAuthorize
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

type NullableOrderUpdateRequest struct {
	value *OrderUpdateRequest
	isSet bool
}

func (v NullableOrderUpdateRequest) Get() *OrderUpdateRequest {
	return v.value
}

func (v *NullableOrderUpdateRequest) Set(val *OrderUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUpdateRequest(val *OrderUpdateRequest) *NullableOrderUpdateRequest {
	return &NullableOrderUpdateRequest{value: val, isSet: true}
}

func (v NullableOrderUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CheckoutOrderTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutOrderTemplate{}

// CheckoutOrderTemplate It maintains the attributes with which the order will be created when receiving a new payment.
type CheckoutOrderTemplate struct {
	// It is the currency in which the order will be created. It must be a valid ISO 4217 currency code.
	Currency     string                             `json:"currency"`
	CustomerInfo *CheckoutOrderTemplateCustomerInfo `json:"customer_info,omitempty"`
	// They are the products to buy. Each contains the \"unit price\" and \"quantity\" parameters that are used to calculate the total amount of the order.
	LineItems []Product `json:"line_items"`
	// It is a set of key-value pairs that you can attach to the order. It can be used to store additional information about the order in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewCheckoutOrderTemplate instantiates a new CheckoutOrderTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutOrderTemplate(currency string, lineItems []Product) *CheckoutOrderTemplate {
	this := CheckoutOrderTemplate{}
	this.Currency = currency
	this.LineItems = lineItems
	return &this
}

// NewCheckoutOrderTemplateWithDefaults instantiates a new CheckoutOrderTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutOrderTemplateWithDefaults() *CheckoutOrderTemplate {
	this := CheckoutOrderTemplate{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *CheckoutOrderTemplate) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CheckoutOrderTemplate) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CheckoutOrderTemplate) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerInfo returns the CustomerInfo field value if set, zero value otherwise.
func (o *CheckoutOrderTemplate) GetCustomerInfo() CheckoutOrderTemplateCustomerInfo {
	if o == nil || IsNil(o.CustomerInfo) {
		var ret CheckoutOrderTemplateCustomerInfo
		return ret
	}
	return *o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderTemplate) GetCustomerInfoOk() (*CheckoutOrderTemplateCustomerInfo, bool) {
	if o == nil || IsNil(o.CustomerInfo) {
		return nil, false
	}
	return o.CustomerInfo, true
}

// HasCustomerInfo returns a boolean if a field has been set.
func (o *CheckoutOrderTemplate) HasCustomerInfo() bool {
	if o != nil && !IsNil(o.CustomerInfo) {
		return true
	}

	return false
}

// SetCustomerInfo gets a reference to the given CheckoutOrderTemplateCustomerInfo and assigns it to the CustomerInfo field.
func (o *CheckoutOrderTemplate) SetCustomerInfo(v CheckoutOrderTemplateCustomerInfo) {
	o.CustomerInfo = &v
}

// GetLineItems returns the LineItems field value
func (o *CheckoutOrderTemplate) GetLineItems() []Product {
	if o == nil {
		var ret []Product
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *CheckoutOrderTemplate) GetLineItemsOk() ([]Product, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *CheckoutOrderTemplate) SetLineItems(v []Product) {
	o.LineItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CheckoutOrderTemplate) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutOrderTemplate) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckoutOrderTemplate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CheckoutOrderTemplate) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CheckoutOrderTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutOrderTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.CustomerInfo) {
		toSerialize["customer_info"] = o.CustomerInfo
	}
	toSerialize["line_items"] = o.LineItems
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableCheckoutOrderTemplate struct {
	value *CheckoutOrderTemplate
	isSet bool
}

func (v NullableCheckoutOrderTemplate) Get() *CheckoutOrderTemplate {
	return v.value
}

func (v *NullableCheckoutOrderTemplate) Set(val *CheckoutOrderTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutOrderTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutOrderTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutOrderTemplate(val *CheckoutOrderTemplate) *NullableCheckoutOrderTemplate {
	return &NullableCheckoutOrderTemplate{value: val, isSet: true}
}

func (v NullableCheckoutOrderTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutOrderTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

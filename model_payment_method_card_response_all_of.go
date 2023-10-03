package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCardResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCardResponseAllOf{}

// PaymentMethodCardResponseAllOf use for card responses
type PaymentMethodCardResponseAllOf struct {
	Last4               *string `json:"last4,omitempty"`
	Bin                 *string `json:"bin,omitempty"`
	CardType            *string `json:"card_type,omitempty"`
	ExpMonth            *string `json:"exp_month,omitempty"`
	ExpYear             *string `json:"exp_year,omitempty"`
	Brand               *string `json:"brand,omitempty"`
	Name                *string `json:"name,omitempty"`
	Default             *bool   `json:"default,omitempty"`
	VisibleOnCheckout   *bool   `json:"visible_on_checkout,omitempty"`
	PaymentSourceStatus *string `json:"payment_source_status,omitempty"`
}

// NewPaymentMethodCardResponseAllOf instantiates a new PaymentMethodCardResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCardResponseAllOf() *PaymentMethodCardResponseAllOf {
	this := PaymentMethodCardResponseAllOf{}
	return &this
}

// NewPaymentMethodCardResponseAllOfWithDefaults instantiates a new PaymentMethodCardResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCardResponseAllOfWithDefaults() *PaymentMethodCardResponseAllOf {
	this := PaymentMethodCardResponseAllOf{}
	return &this
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetLast4() string {
	if o == nil || IsNil(o.Last4) {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetLast4Ok() (*string, bool) {
	if o == nil || IsNil(o.Last4) {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasLast4() bool {
	if o != nil && !IsNil(o.Last4) {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PaymentMethodCardResponseAllOf) SetLast4(v string) {
	o.Last4 = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *PaymentMethodCardResponseAllOf) SetBin(v string) {
	o.Bin = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *PaymentMethodCardResponseAllOf) SetCardType(v string) {
	o.CardType = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetExpMonth() string {
	if o == nil || IsNil(o.ExpMonth) {
		var ret string
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetExpMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpMonth) {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasExpMonth() bool {
	if o != nil && !IsNil(o.ExpMonth) {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given string and assigns it to the ExpMonth field.
func (o *PaymentMethodCardResponseAllOf) SetExpMonth(v string) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetExpYear() string {
	if o == nil || IsNil(o.ExpYear) {
		var ret string
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetExpYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpYear) {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasExpYear() bool {
	if o != nil && !IsNil(o.ExpYear) {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given string and assigns it to the ExpYear field.
func (o *PaymentMethodCardResponseAllOf) SetExpYear(v string) {
	o.ExpYear = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *PaymentMethodCardResponseAllOf) SetBrand(v string) {
	o.Brand = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodCardResponseAllOf) SetName(v string) {
	o.Name = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethodCardResponseAllOf) SetDefault(v bool) {
	o.Default = &v
}

// GetVisibleOnCheckout returns the VisibleOnCheckout field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetVisibleOnCheckout() bool {
	if o == nil || IsNil(o.VisibleOnCheckout) {
		var ret bool
		return ret
	}
	return *o.VisibleOnCheckout
}

// GetVisibleOnCheckoutOk returns a tuple with the VisibleOnCheckout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetVisibleOnCheckoutOk() (*bool, bool) {
	if o == nil || IsNil(o.VisibleOnCheckout) {
		return nil, false
	}
	return o.VisibleOnCheckout, true
}

// HasVisibleOnCheckout returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasVisibleOnCheckout() bool {
	if o != nil && !IsNil(o.VisibleOnCheckout) {
		return true
	}

	return false
}

// SetVisibleOnCheckout gets a reference to the given bool and assigns it to the VisibleOnCheckout field.
func (o *PaymentMethodCardResponseAllOf) SetVisibleOnCheckout(v bool) {
	o.VisibleOnCheckout = &v
}

// GetPaymentSourceStatus returns the PaymentSourceStatus field value if set, zero value otherwise.
func (o *PaymentMethodCardResponseAllOf) GetPaymentSourceStatus() string {
	if o == nil || IsNil(o.PaymentSourceStatus) {
		var ret string
		return ret
	}
	return *o.PaymentSourceStatus
}

// GetPaymentSourceStatusOk returns a tuple with the PaymentSourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponseAllOf) GetPaymentSourceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentSourceStatus) {
		return nil, false
	}
	return o.PaymentSourceStatus, true
}

// HasPaymentSourceStatus returns a boolean if a field has been set.
func (o *PaymentMethodCardResponseAllOf) HasPaymentSourceStatus() bool {
	if o != nil && !IsNil(o.PaymentSourceStatus) {
		return true
	}

	return false
}

// SetPaymentSourceStatus gets a reference to the given string and assigns it to the PaymentSourceStatus field.
func (o *PaymentMethodCardResponseAllOf) SetPaymentSourceStatus(v string) {
	o.PaymentSourceStatus = &v
}

func (o PaymentMethodCardResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCardResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Last4) {
		toSerialize["last4"] = o.Last4
	}
	if !IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.ExpMonth) {
		toSerialize["exp_month"] = o.ExpMonth
	}
	if !IsNil(o.ExpYear) {
		toSerialize["exp_year"] = o.ExpYear
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.VisibleOnCheckout) {
		toSerialize["visible_on_checkout"] = o.VisibleOnCheckout
	}
	if !IsNil(o.PaymentSourceStatus) {
		toSerialize["payment_source_status"] = o.PaymentSourceStatus
	}
	return toSerialize, nil
}

type NullablePaymentMethodCardResponseAllOf struct {
	value *PaymentMethodCardResponseAllOf
	isSet bool
}

func (v NullablePaymentMethodCardResponseAllOf) Get() *PaymentMethodCardResponseAllOf {
	return v.value
}

func (v *NullablePaymentMethodCardResponseAllOf) Set(val *PaymentMethodCardResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCardResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCardResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCardResponseAllOf(val *PaymentMethodCardResponseAllOf) *NullablePaymentMethodCardResponseAllOf {
	return &NullablePaymentMethodCardResponseAllOf{value: val, isSet: true}
}

func (v NullablePaymentMethodCardResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCardResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

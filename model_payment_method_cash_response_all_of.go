package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCashResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCashResponseAllOf{}

// PaymentMethodCashResponseAllOf use for cash responses
type PaymentMethodCashResponseAllOf struct {
	Reference  *string `json:"reference,omitempty"`
	Barcode    *string `json:"barcode,omitempty"`
	BarcodeUrl *string `json:"barcode_url,omitempty"`
	ExpiresAt  *int64  `json:"expires_at,omitempty"`
	Provider   *string `json:"provider,omitempty"`
}

// NewPaymentMethodCashResponseAllOf instantiates a new PaymentMethodCashResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCashResponseAllOf() *PaymentMethodCashResponseAllOf {
	this := PaymentMethodCashResponseAllOf{}
	return &this
}

// NewPaymentMethodCashResponseAllOfWithDefaults instantiates a new PaymentMethodCashResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCashResponseAllOfWithDefaults() *PaymentMethodCashResponseAllOf {
	this := PaymentMethodCashResponseAllOf{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodCashResponseAllOf) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponseAllOf) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodCashResponseAllOf) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodCashResponseAllOf) SetReference(v string) {
	o.Reference = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *PaymentMethodCashResponseAllOf) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponseAllOf) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *PaymentMethodCashResponseAllOf) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *PaymentMethodCashResponseAllOf) SetBarcode(v string) {
	o.Barcode = &v
}

// GetBarcodeUrl returns the BarcodeUrl field value if set, zero value otherwise.
func (o *PaymentMethodCashResponseAllOf) GetBarcodeUrl() string {
	if o == nil || IsNil(o.BarcodeUrl) {
		var ret string
		return ret
	}
	return *o.BarcodeUrl
}

// GetBarcodeUrlOk returns a tuple with the BarcodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponseAllOf) GetBarcodeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BarcodeUrl) {
		return nil, false
	}
	return o.BarcodeUrl, true
}

// HasBarcodeUrl returns a boolean if a field has been set.
func (o *PaymentMethodCashResponseAllOf) HasBarcodeUrl() bool {
	if o != nil && !IsNil(o.BarcodeUrl) {
		return true
	}

	return false
}

// SetBarcodeUrl gets a reference to the given string and assigns it to the BarcodeUrl field.
func (o *PaymentMethodCashResponseAllOf) SetBarcodeUrl(v string) {
	o.BarcodeUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodCashResponseAllOf) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponseAllOf) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodCashResponseAllOf) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PaymentMethodCashResponseAllOf) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PaymentMethodCashResponseAllOf) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponseAllOf) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PaymentMethodCashResponseAllOf) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PaymentMethodCashResponseAllOf) SetProvider(v string) {
	o.Provider = &v
}

func (o PaymentMethodCashResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCashResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.BarcodeUrl) {
		toSerialize["barcode_url"] = o.BarcodeUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullablePaymentMethodCashResponseAllOf struct {
	value *PaymentMethodCashResponseAllOf
	isSet bool
}

func (v NullablePaymentMethodCashResponseAllOf) Get() *PaymentMethodCashResponseAllOf {
	return v.value
}

func (v *NullablePaymentMethodCashResponseAllOf) Set(val *PaymentMethodCashResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCashResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCashResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCashResponseAllOf(val *PaymentMethodCashResponseAllOf) *NullablePaymentMethodCashResponseAllOf {
	return &NullablePaymentMethodCashResponseAllOf{value: val, isSet: true}
}

func (v NullablePaymentMethodCashResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCashResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

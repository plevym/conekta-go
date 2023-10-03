package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCash{}

// PaymentMethodCash struct for PaymentMethodCash
type PaymentMethodCash struct {
	Type        *string        `json:"type,omitempty"`
	Object      string         `json:"object"`
	AuthCode    NullableInt32  `json:"auth_code,omitempty"`
	CashierId   NullableString `json:"cashier_id,omitempty"`
	Reference   *string        `json:"reference,omitempty"`
	BarcodeUrl  *string        `json:"barcode_url,omitempty"`
	ExpiresAt   *int64         `json:"expires_at,omitempty"`
	ServiceName *string        `json:"service_name,omitempty"`
	Store       NullableString `json:"store,omitempty"`
	StoreName   *string        `json:"store_name,omitempty"`
}

// NewPaymentMethodCash instantiates a new PaymentMethodCash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCash(object string) *PaymentMethodCash {
	this := PaymentMethodCash{}
	this.Object = object
	return &this
}

// NewPaymentMethodCashWithDefaults instantiates a new PaymentMethodCash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCashWithDefaults() *PaymentMethodCash {
	this := PaymentMethodCash{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodCash) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodCash) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value
func (o *PaymentMethodCash) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodCash) SetObject(v string) {
	o.Object = v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodCash) GetAuthCode() int32 {
	if o == nil || IsNil(o.AuthCode.Get()) {
		var ret int32
		return ret
	}
	return *o.AuthCode.Get()
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodCash) GetAuthCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthCode.Get(), o.AuthCode.IsSet()
}

// HasAuthCode returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasAuthCode() bool {
	if o != nil && o.AuthCode.IsSet() {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given NullableInt32 and assigns it to the AuthCode field.
func (o *PaymentMethodCash) SetAuthCode(v int32) {
	o.AuthCode.Set(&v)
}

// SetAuthCodeNil sets the value for AuthCode to be an explicit nil
func (o *PaymentMethodCash) SetAuthCodeNil() {
	o.AuthCode.Set(nil)
}

// UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
func (o *PaymentMethodCash) UnsetAuthCode() {
	o.AuthCode.Unset()
}

// GetCashierId returns the CashierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodCash) GetCashierId() string {
	if o == nil || IsNil(o.CashierId.Get()) {
		var ret string
		return ret
	}
	return *o.CashierId.Get()
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodCash) GetCashierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CashierId.Get(), o.CashierId.IsSet()
}

// HasCashierId returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasCashierId() bool {
	if o != nil && o.CashierId.IsSet() {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given NullableString and assigns it to the CashierId field.
func (o *PaymentMethodCash) SetCashierId(v string) {
	o.CashierId.Set(&v)
}

// SetCashierIdNil sets the value for CashierId to be an explicit nil
func (o *PaymentMethodCash) SetCashierIdNil() {
	o.CashierId.Set(nil)
}

// UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
func (o *PaymentMethodCash) UnsetCashierId() {
	o.CashierId.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodCash) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodCash) SetReference(v string) {
	o.Reference = &v
}

// GetBarcodeUrl returns the BarcodeUrl field value if set, zero value otherwise.
func (o *PaymentMethodCash) GetBarcodeUrl() string {
	if o == nil || IsNil(o.BarcodeUrl) {
		var ret string
		return ret
	}
	return *o.BarcodeUrl
}

// GetBarcodeUrlOk returns a tuple with the BarcodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetBarcodeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BarcodeUrl) {
		return nil, false
	}
	return o.BarcodeUrl, true
}

// HasBarcodeUrl returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasBarcodeUrl() bool {
	if o != nil && !IsNil(o.BarcodeUrl) {
		return true
	}

	return false
}

// SetBarcodeUrl gets a reference to the given string and assigns it to the BarcodeUrl field.
func (o *PaymentMethodCash) SetBarcodeUrl(v string) {
	o.BarcodeUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodCash) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PaymentMethodCash) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *PaymentMethodCash) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *PaymentMethodCash) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetStore returns the Store field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodCash) GetStore() string {
	if o == nil || IsNil(o.Store.Get()) {
		var ret string
		return ret
	}
	return *o.Store.Get()
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodCash) GetStoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Store.Get(), o.Store.IsSet()
}

// HasStore returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasStore() bool {
	if o != nil && o.Store.IsSet() {
		return true
	}

	return false
}

// SetStore gets a reference to the given NullableString and assigns it to the Store field.
func (o *PaymentMethodCash) SetStore(v string) {
	o.Store.Set(&v)
}

// SetStoreNil sets the value for Store to be an explicit nil
func (o *PaymentMethodCash) SetStoreNil() {
	o.Store.Set(nil)
}

// UnsetStore ensures that no value is present for Store, not even an explicit nil
func (o *PaymentMethodCash) UnsetStore() {
	o.Store.Unset()
}

// GetStoreName returns the StoreName field value if set, zero value otherwise.
func (o *PaymentMethodCash) GetStoreName() string {
	if o == nil || IsNil(o.StoreName) {
		var ret string
		return ret
	}
	return *o.StoreName
}

// GetStoreNameOk returns a tuple with the StoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCash) GetStoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.StoreName) {
		return nil, false
	}
	return o.StoreName, true
}

// HasStoreName returns a boolean if a field has been set.
func (o *PaymentMethodCash) HasStoreName() bool {
	if o != nil && !IsNil(o.StoreName) {
		return true
	}

	return false
}

// SetStoreName gets a reference to the given string and assigns it to the StoreName field.
func (o *PaymentMethodCash) SetStoreName(v string) {
	o.StoreName = &v
}

func (o PaymentMethodCash) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["object"] = o.Object
	if o.AuthCode.IsSet() {
		toSerialize["auth_code"] = o.AuthCode.Get()
	}
	if o.CashierId.IsSet() {
		toSerialize["cashier_id"] = o.CashierId.Get()
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.BarcodeUrl) {
		toSerialize["barcode_url"] = o.BarcodeUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.Store.IsSet() {
		toSerialize["store"] = o.Store.Get()
	}
	if !IsNil(o.StoreName) {
		toSerialize["store_name"] = o.StoreName
	}
	return toSerialize, nil
}

type NullablePaymentMethodCash struct {
	value *PaymentMethodCash
	isSet bool
}

func (v NullablePaymentMethodCash) Get() *PaymentMethodCash {
	return v.value
}

func (v *NullablePaymentMethodCash) Set(val *PaymentMethodCash) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCash) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCash(val *PaymentMethodCash) *NullablePaymentMethodCash {
	return &NullablePaymentMethodCash{value: val, isSet: true}
}

func (v NullablePaymentMethodCash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

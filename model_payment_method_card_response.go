package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCardResponse{}

// PaymentMethodCardResponse struct for PaymentMethodCardResponse
type PaymentMethodCardResponse struct {
	Type                string  `json:"type"`
	Id                  string  `json:"id"`
	Object              string  `json:"object"`
	CreatedAt           int64   `json:"created_at"`
	ParentId            *string `json:"parent_id,omitempty"`
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

// NewPaymentMethodCardResponse instantiates a new PaymentMethodCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCardResponse(type_ string, id string, object string, createdAt int64) *PaymentMethodCardResponse {
	this := PaymentMethodCardResponse{}
	this.Type = type_
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewPaymentMethodCardResponseWithDefaults instantiates a new PaymentMethodCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCardResponseWithDefaults() *PaymentMethodCardResponse {
	this := PaymentMethodCardResponse{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodCardResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodCardResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PaymentMethodCardResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodCardResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PaymentMethodCardResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodCardResponse) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethodCardResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethodCardResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *PaymentMethodCardResponse) SetParentId(v string) {
	o.ParentId = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetLast4() string {
	if o == nil || IsNil(o.Last4) {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetLast4Ok() (*string, bool) {
	if o == nil || IsNil(o.Last4) {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasLast4() bool {
	if o != nil && !IsNil(o.Last4) {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PaymentMethodCardResponse) SetLast4(v string) {
	o.Last4 = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *PaymentMethodCardResponse) SetBin(v string) {
	o.Bin = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *PaymentMethodCardResponse) SetCardType(v string) {
	o.CardType = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetExpMonth() string {
	if o == nil || IsNil(o.ExpMonth) {
		var ret string
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetExpMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpMonth) {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasExpMonth() bool {
	if o != nil && !IsNil(o.ExpMonth) {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given string and assigns it to the ExpMonth field.
func (o *PaymentMethodCardResponse) SetExpMonth(v string) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetExpYear() string {
	if o == nil || IsNil(o.ExpYear) {
		var ret string
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetExpYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpYear) {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasExpYear() bool {
	if o != nil && !IsNil(o.ExpYear) {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given string and assigns it to the ExpYear field.
func (o *PaymentMethodCardResponse) SetExpYear(v string) {
	o.ExpYear = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *PaymentMethodCardResponse) SetBrand(v string) {
	o.Brand = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodCardResponse) SetName(v string) {
	o.Name = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethodCardResponse) SetDefault(v bool) {
	o.Default = &v
}

// GetVisibleOnCheckout returns the VisibleOnCheckout field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetVisibleOnCheckout() bool {
	if o == nil || IsNil(o.VisibleOnCheckout) {
		var ret bool
		return ret
	}
	return *o.VisibleOnCheckout
}

// GetVisibleOnCheckoutOk returns a tuple with the VisibleOnCheckout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetVisibleOnCheckoutOk() (*bool, bool) {
	if o == nil || IsNil(o.VisibleOnCheckout) {
		return nil, false
	}
	return o.VisibleOnCheckout, true
}

// HasVisibleOnCheckout returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasVisibleOnCheckout() bool {
	if o != nil && !IsNil(o.VisibleOnCheckout) {
		return true
	}

	return false
}

// SetVisibleOnCheckout gets a reference to the given bool and assigns it to the VisibleOnCheckout field.
func (o *PaymentMethodCardResponse) SetVisibleOnCheckout(v bool) {
	o.VisibleOnCheckout = &v
}

// GetPaymentSourceStatus returns the PaymentSourceStatus field value if set, zero value otherwise.
func (o *PaymentMethodCardResponse) GetPaymentSourceStatus() string {
	if o == nil || IsNil(o.PaymentSourceStatus) {
		var ret string
		return ret
	}
	return *o.PaymentSourceStatus
}

// GetPaymentSourceStatusOk returns a tuple with the PaymentSourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardResponse) GetPaymentSourceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentSourceStatus) {
		return nil, false
	}
	return o.PaymentSourceStatus, true
}

// HasPaymentSourceStatus returns a boolean if a field has been set.
func (o *PaymentMethodCardResponse) HasPaymentSourceStatus() bool {
	if o != nil && !IsNil(o.PaymentSourceStatus) {
		return true
	}

	return false
}

// SetPaymentSourceStatus gets a reference to the given string and assigns it to the PaymentSourceStatus field.
func (o *PaymentMethodCardResponse) SetPaymentSourceStatus(v string) {
	o.PaymentSourceStatus = &v
}

func (o PaymentMethodCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
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

type NullablePaymentMethodCardResponse struct {
	value *PaymentMethodCardResponse
	isSet bool
}

func (v NullablePaymentMethodCardResponse) Get() *PaymentMethodCardResponse {
	return v.value
}

func (v *NullablePaymentMethodCardResponse) Set(val *PaymentMethodCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCardResponse(val *PaymentMethodCardResponse) *NullablePaymentMethodCardResponse {
	return &NullablePaymentMethodCardResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

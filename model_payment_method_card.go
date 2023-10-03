package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCard{}

// PaymentMethodCard struct for PaymentMethodCard
type PaymentMethodCard struct {
	Type        *string `json:"type,omitempty"`
	Object      string  `json:"object"`
	AccountType *string `json:"account_type,omitempty"`
	AuthCode    *string `json:"auth_code,omitempty"`
	Brand       *string `json:"brand,omitempty"`
	// Id sent for recurrent charges.
	ContractId      *string       `json:"contract_id,omitempty"`
	Country         *string       `json:"country,omitempty"`
	ExpMonth        *string       `json:"exp_month,omitempty"`
	ExpYear         *string       `json:"exp_year,omitempty"`
	FraudIndicators []interface{} `json:"fraud_indicators,omitempty"`
	Issuer          *string       `json:"issuer,omitempty"`
	Last4           *string       `json:"last4,omitempty"`
	Name            *string       `json:"name,omitempty"`
}

// NewPaymentMethodCard instantiates a new PaymentMethodCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCard(object string) *PaymentMethodCard {
	this := PaymentMethodCard{}
	this.Object = object
	return &this
}

// NewPaymentMethodCardWithDefaults instantiates a new PaymentMethodCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCardWithDefaults() *PaymentMethodCard {
	this := PaymentMethodCard{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodCard) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value
func (o *PaymentMethodCard) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodCard) SetObject(v string) {
	o.Object = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *PaymentMethodCard) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetAuthCode() string {
	if o == nil || IsNil(o.AuthCode) {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetAuthCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthCode) {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasAuthCode() bool {
	if o != nil && !IsNil(o.AuthCode) {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *PaymentMethodCard) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *PaymentMethodCard) SetBrand(v string) {
	o.Brand = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *PaymentMethodCard) SetContractId(v string) {
	o.ContractId = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PaymentMethodCard) SetCountry(v string) {
	o.Country = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetExpMonth() string {
	if o == nil || IsNil(o.ExpMonth) {
		var ret string
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetExpMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpMonth) {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasExpMonth() bool {
	if o != nil && !IsNil(o.ExpMonth) {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given string and assigns it to the ExpMonth field.
func (o *PaymentMethodCard) SetExpMonth(v string) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetExpYear() string {
	if o == nil || IsNil(o.ExpYear) {
		var ret string
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetExpYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpYear) {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasExpYear() bool {
	if o != nil && !IsNil(o.ExpYear) {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given string and assigns it to the ExpYear field.
func (o *PaymentMethodCard) SetExpYear(v string) {
	o.ExpYear = &v
}

// GetFraudIndicators returns the FraudIndicators field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetFraudIndicators() []interface{} {
	if o == nil || IsNil(o.FraudIndicators) {
		var ret []interface{}
		return ret
	}
	return o.FraudIndicators
}

// GetFraudIndicatorsOk returns a tuple with the FraudIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetFraudIndicatorsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.FraudIndicators) {
		return nil, false
	}
	return o.FraudIndicators, true
}

// HasFraudIndicators returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasFraudIndicators() bool {
	if o != nil && !IsNil(o.FraudIndicators) {
		return true
	}

	return false
}

// SetFraudIndicators gets a reference to the given []interface{} and assigns it to the FraudIndicators field.
func (o *PaymentMethodCard) SetFraudIndicators(v []interface{}) {
	o.FraudIndicators = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *PaymentMethodCard) SetIssuer(v string) {
	o.Issuer = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetLast4() string {
	if o == nil || IsNil(o.Last4) {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetLast4Ok() (*string, bool) {
	if o == nil || IsNil(o.Last4) {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasLast4() bool {
	if o != nil && !IsNil(o.Last4) {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PaymentMethodCard) SetLast4(v string) {
	o.Last4 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodCard) SetName(v string) {
	o.Name = &v
}

func (o PaymentMethodCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.AuthCode) {
		toSerialize["auth_code"] = o.AuthCode
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.ContractId) {
		toSerialize["contract_id"] = o.ContractId
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.ExpMonth) {
		toSerialize["exp_month"] = o.ExpMonth
	}
	if !IsNil(o.ExpYear) {
		toSerialize["exp_year"] = o.ExpYear
	}
	if !IsNil(o.FraudIndicators) {
		toSerialize["fraud_indicators"] = o.FraudIndicators
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Last4) {
		toSerialize["last4"] = o.Last4
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePaymentMethodCard struct {
	value *PaymentMethodCard
	isSet bool
}

func (v NullablePaymentMethodCard) Get() *PaymentMethodCard {
	return v.value
}

func (v *NullablePaymentMethodCard) Set(val *PaymentMethodCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCard(val *PaymentMethodCard) *NullablePaymentMethodCard {
	return &NullablePaymentMethodCard{value: val, isSet: true}
}

func (v NullablePaymentMethodCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

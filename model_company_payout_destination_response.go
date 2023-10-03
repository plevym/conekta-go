package conekta

import (
	"encoding/json"
)

// checks if the CompanyPayoutDestinationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyPayoutDestinationResponse{}

// CompanyPayoutDestinationResponse Company payout destination model
type CompanyPayoutDestinationResponse struct {
	// The resource's type
	Object *string `json:"object,omitempty"`
	// currency of the receiving account
	Currency *string `json:"currency,omitempty"`
	// Name of the account holder
	AccountHolderName *string `json:"account_holder_name,omitempty"`
	// Name of the bank
	Bank *string `json:"bank,omitempty"`
	// Type of the payout destination
	Type *string `json:"type,omitempty"`
	// Account number of the receiving account
	AccountNumber *string `json:"account_number,omitempty"`
}

// NewCompanyPayoutDestinationResponse instantiates a new CompanyPayoutDestinationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyPayoutDestinationResponse() *CompanyPayoutDestinationResponse {
	this := CompanyPayoutDestinationResponse{}
	return &this
}

// NewCompanyPayoutDestinationResponseWithDefaults instantiates a new CompanyPayoutDestinationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyPayoutDestinationResponseWithDefaults() *CompanyPayoutDestinationResponse {
	this := CompanyPayoutDestinationResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CompanyPayoutDestinationResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyPayoutDestinationResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CompanyPayoutDestinationResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CompanyPayoutDestinationResponse) SetObject(v string) {
	o.Object = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CompanyPayoutDestinationResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyPayoutDestinationResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CompanyPayoutDestinationResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CompanyPayoutDestinationResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetAccountHolderName returns the AccountHolderName field value if set, zero value otherwise.
func (o *CompanyPayoutDestinationResponse) GetAccountHolderName() string {
	if o == nil || IsNil(o.AccountHolderName) {
		var ret string
		return ret
	}
	return *o.AccountHolderName
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyPayoutDestinationResponse) GetAccountHolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountHolderName) {
		return nil, false
	}
	return o.AccountHolderName, true
}

// HasAccountHolderName returns a boolean if a field has been set.
func (o *CompanyPayoutDestinationResponse) HasAccountHolderName() bool {
	if o != nil && !IsNil(o.AccountHolderName) {
		return true
	}

	return false
}

// SetAccountHolderName gets a reference to the given string and assigns it to the AccountHolderName field.
func (o *CompanyPayoutDestinationResponse) SetAccountHolderName(v string) {
	o.AccountHolderName = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *CompanyPayoutDestinationResponse) GetBank() string {
	if o == nil || IsNil(o.Bank) {
		var ret string
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyPayoutDestinationResponse) GetBankOk() (*string, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *CompanyPayoutDestinationResponse) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given string and assigns it to the Bank field.
func (o *CompanyPayoutDestinationResponse) SetBank(v string) {
	o.Bank = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CompanyPayoutDestinationResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyPayoutDestinationResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CompanyPayoutDestinationResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CompanyPayoutDestinationResponse) SetType(v string) {
	o.Type = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *CompanyPayoutDestinationResponse) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyPayoutDestinationResponse) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *CompanyPayoutDestinationResponse) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *CompanyPayoutDestinationResponse) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

func (o CompanyPayoutDestinationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyPayoutDestinationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.AccountHolderName) {
		toSerialize["account_holder_name"] = o.AccountHolderName
	}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	return toSerialize, nil
}

type NullableCompanyPayoutDestinationResponse struct {
	value *CompanyPayoutDestinationResponse
	isSet bool
}

func (v NullableCompanyPayoutDestinationResponse) Get() *CompanyPayoutDestinationResponse {
	return v.value
}

func (v *NullableCompanyPayoutDestinationResponse) Set(val *CompanyPayoutDestinationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyPayoutDestinationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyPayoutDestinationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyPayoutDestinationResponse(val *CompanyPayoutDestinationResponse) *NullableCompanyPayoutDestinationResponse {
	return &NullableCompanyPayoutDestinationResponse{value: val, isSet: true}
}

func (v NullableCompanyPayoutDestinationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyPayoutDestinationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

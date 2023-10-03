package conekta

import (
	"encoding/json"
)

// checks if the TransferMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferMethodResponse{}

// TransferMethodResponse Method used to make the transfer.
type TransferMethodResponse struct {
	// Name of the account holder.
	AccountHolder *string `json:"account_holder,omitempty"`
	// Account number of the bank account.
	AccountNumber *string `json:"account_number,omitempty"`
	// Name of the bank.
	Bank *string `json:"bank,omitempty"`
	// Date and time of creation of the transfer.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unique identifier of the transfer.
	Id *string `json:"id,omitempty"`
	// Object name, which is bank_transfer_payout_method.
	Object *string `json:"object,omitempty"`
	// Unique identifier of the payee.
	PayeeId *string `json:"payee_id,omitempty"`
	// Type of the payee.
	Type *string `json:"type,omitempty"`
}

// NewTransferMethodResponse instantiates a new TransferMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferMethodResponse() *TransferMethodResponse {
	this := TransferMethodResponse{}
	return &this
}

// NewTransferMethodResponseWithDefaults instantiates a new TransferMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferMethodResponseWithDefaults() *TransferMethodResponse {
	this := TransferMethodResponse{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetAccountHolder() string {
	if o == nil || IsNil(o.AccountHolder) {
		var ret string
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetAccountHolderOk() (*string, bool) {
	if o == nil || IsNil(o.AccountHolder) {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasAccountHolder() bool {
	if o != nil && !IsNil(o.AccountHolder) {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given string and assigns it to the AccountHolder field.
func (o *TransferMethodResponse) SetAccountHolder(v string) {
	o.AccountHolder = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *TransferMethodResponse) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetBank() string {
	if o == nil || IsNil(o.Bank) {
		var ret string
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetBankOk() (*string, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given string and assigns it to the Bank field.
func (o *TransferMethodResponse) SetBank(v string) {
	o.Bank = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *TransferMethodResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferMethodResponse) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TransferMethodResponse) SetObject(v string) {
	o.Object = &v
}

// GetPayeeId returns the PayeeId field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetPayeeId() string {
	if o == nil || IsNil(o.PayeeId) {
		var ret string
		return ret
	}
	return *o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetPayeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeId) {
		return nil, false
	}
	return o.PayeeId, true
}

// HasPayeeId returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasPayeeId() bool {
	if o != nil && !IsNil(o.PayeeId) {
		return true
	}

	return false
}

// SetPayeeId gets a reference to the given string and assigns it to the PayeeId field.
func (o *TransferMethodResponse) SetPayeeId(v string) {
	o.PayeeId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferMethodResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMethodResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferMethodResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferMethodResponse) SetType(v string) {
	o.Type = &v
}

func (o TransferMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountHolder) {
		toSerialize["account_holder"] = o.AccountHolder
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.PayeeId) {
		toSerialize["payee_id"] = o.PayeeId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransferMethodResponse struct {
	value *TransferMethodResponse
	isSet bool
}

func (v NullableTransferMethodResponse) Get() *TransferMethodResponse {
	return v.value
}

func (v *NullableTransferMethodResponse) Set(val *TransferMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferMethodResponse(val *TransferMethodResponse) *NullableTransferMethodResponse {
	return &NullableTransferMethodResponse{value: val, isSet: true}
}

func (v NullableTransferMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

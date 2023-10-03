package conekta

import (
	"encoding/json"
)

// checks if the TransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionResponse{}

// TransactionResponse The Transaction object represents the actions or steps of an order. Statuses can be: unprocessed, pending, available, owen, paid_out, voided, capture, capture_reversal, liquidation, liquidation_reversal, payout, payout_reversal, refund, refund_reversal, chargeback, chargeback_reversal, rounding_adjustment, won_chargeback, transferred, and transferred.
type TransactionResponse struct {
	// The amount of the transaction.
	Amount int64 `json:"amount"`
	// Randomly assigned unique order identifier associated with the charge.
	Charge string `json:"charge"`
	// Date and time of creation of the transaction in Unix format.
	CreatedAt int64 `json:"created_at"`
	// The currency of the transaction. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217)
	Currency string `json:"currency"`
	// The amount to be deducted for taxes and commissions.
	Fee int64 `json:"fee"`
	// Unique identifier of the transaction.
	Id string `json:"id"`
	// Indicates whether the transaction was created in live mode or test mode.
	Livemode bool `json:"livemode"`
	// The net amount after deducting commissions and taxes.
	Net int64 `json:"net"`
	// Object name, which is transaction.
	Object string `json:"object"`
	// Code indicating transaction status.
	Status string `json:"status"`
	// Transaction Type
	Type string `json:"type"`
}

// NewTransactionResponse instantiates a new TransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResponse(amount int64, charge string, createdAt int64, currency string, fee int64, id string, livemode bool, net int64, object string, status string, type_ string) *TransactionResponse {
	this := TransactionResponse{}
	this.Amount = amount
	this.Charge = charge
	this.CreatedAt = createdAt
	this.Currency = currency
	this.Fee = fee
	this.Id = id
	this.Livemode = livemode
	this.Net = net
	this.Object = object
	this.Status = status
	this.Type = type_
	return &this
}

// NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResponseWithDefaults() *TransactionResponse {
	this := TransactionResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TransactionResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCharge returns the Charge field value
func (o *TransactionResponse) GetCharge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Charge
}

// GetChargeOk returns a tuple with the Charge field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetChargeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Charge, true
}

// SetCharge sets field value
func (o *TransactionResponse) SetCharge(v string) {
	o.Charge = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TransactionResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TransactionResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value
func (o *TransactionResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetFee returns the Fee field value
func (o *TransactionResponse) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *TransactionResponse) SetFee(v int64) {
	o.Fee = v
}

// GetId returns the Id field value
func (o *TransactionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionResponse) SetId(v string) {
	o.Id = v
}

// GetLivemode returns the Livemode field value
func (o *TransactionResponse) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *TransactionResponse) SetLivemode(v bool) {
	o.Livemode = v
}

// GetNet returns the Net field value
func (o *TransactionResponse) GetNet() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Net
}

// GetNetOk returns a tuple with the Net field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetNetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Net, true
}

// SetNet sets field value
func (o *TransactionResponse) SetNet(v int64) {
	o.Net = v
}

// GetObject returns the Object field value
func (o *TransactionResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *TransactionResponse) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *TransactionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *TransactionResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionResponse) SetType(v string) {
	o.Type = v
}

func (o TransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["charge"] = o.Charge
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["currency"] = o.Currency
	toSerialize["fee"] = o.Fee
	toSerialize["id"] = o.Id
	toSerialize["livemode"] = o.Livemode
	toSerialize["net"] = o.Net
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransactionResponse struct {
	value *TransactionResponse
	isSet bool
}

func (v NullableTransactionResponse) Get() *TransactionResponse {
	return v.value
}

func (v *NullableTransactionResponse) Set(val *TransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResponse(val *TransactionResponse) *NullableTransactionResponse {
	return &NullableTransactionResponse{value: val, isSet: true}
}

func (v NullableTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

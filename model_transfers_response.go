package conekta

import (
	"encoding/json"
)

// checks if the TransfersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransfersResponse{}

// TransfersResponse A transfer represents the action of sending an amount to a business bank account including the status, amount and method used to make the transfer.
type TransfersResponse struct {
	// Amount in cents of the transfer.
	Amount *int64 `json:"amount,omitempty"`
	// Date and time of creation of the transfer.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The currency of the transfer. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217)
	Currency *string `json:"currency,omitempty"`
	// Unique identifier of the transfer.
	Id *string `json:"id,omitempty"`
	// Indicates whether the transfer was created in live mode or test mode.
	Livemode *bool                   `json:"livemode,omitempty"`
	Method   *TransferMethodResponse `json:"method,omitempty"`
	// Object name, which is transfer.
	Object *string `json:"object,omitempty"`
	// Description of the transfer.
	StatementDescription *string `json:"statement_description,omitempty"`
	// Reference number of the transfer.
	StatementReference *string `json:"statement_reference,omitempty"`
	// Code indicating transfer status.
	Status *string `json:"status,omitempty"`
}

// NewTransfersResponse instantiates a new TransfersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfersResponse() *TransfersResponse {
	this := TransfersResponse{}
	return &this
}

// NewTransfersResponseWithDefaults instantiates a new TransfersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransfersResponseWithDefaults() *TransfersResponse {
	this := TransfersResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransfersResponse) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransfersResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *TransfersResponse) SetAmount(v int64) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransfersResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransfersResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *TransfersResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TransfersResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransfersResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TransfersResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransfersResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransfersResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransfersResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *TransfersResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *TransfersResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *TransfersResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TransfersResponse) GetMethod() TransferMethodResponse {
	if o == nil || IsNil(o.Method) {
		var ret TransferMethodResponse
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetMethodOk() (*TransferMethodResponse, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TransfersResponse) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given TransferMethodResponse and assigns it to the Method field.
func (o *TransfersResponse) SetMethod(v TransferMethodResponse) {
	o.Method = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TransfersResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TransfersResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TransfersResponse) SetObject(v string) {
	o.Object = &v
}

// GetStatementDescription returns the StatementDescription field value if set, zero value otherwise.
func (o *TransfersResponse) GetStatementDescription() string {
	if o == nil || IsNil(o.StatementDescription) {
		var ret string
		return ret
	}
	return *o.StatementDescription
}

// GetStatementDescriptionOk returns a tuple with the StatementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetStatementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StatementDescription) {
		return nil, false
	}
	return o.StatementDescription, true
}

// HasStatementDescription returns a boolean if a field has been set.
func (o *TransfersResponse) HasStatementDescription() bool {
	if o != nil && !IsNil(o.StatementDescription) {
		return true
	}

	return false
}

// SetStatementDescription gets a reference to the given string and assigns it to the StatementDescription field.
func (o *TransfersResponse) SetStatementDescription(v string) {
	o.StatementDescription = &v
}

// GetStatementReference returns the StatementReference field value if set, zero value otherwise.
func (o *TransfersResponse) GetStatementReference() string {
	if o == nil || IsNil(o.StatementReference) {
		var ret string
		return ret
	}
	return *o.StatementReference
}

// GetStatementReferenceOk returns a tuple with the StatementReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetStatementReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.StatementReference) {
		return nil, false
	}
	return o.StatementReference, true
}

// HasStatementReference returns a boolean if a field has been set.
func (o *TransfersResponse) HasStatementReference() bool {
	if o != nil && !IsNil(o.StatementReference) {
		return true
	}

	return false
}

// SetStatementReference gets a reference to the given string and assigns it to the StatementReference field.
func (o *TransfersResponse) SetStatementReference(v string) {
	o.StatementReference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransfersResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransfersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransfersResponse) SetStatus(v string) {
	o.Status = &v
}

func (o TransfersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransfersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.StatementDescription) {
		toSerialize["statement_description"] = o.StatementDescription
	}
	if !IsNil(o.StatementReference) {
		toSerialize["statement_reference"] = o.StatementReference
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableTransfersResponse struct {
	value *TransfersResponse
	isSet bool
}

func (v NullableTransfersResponse) Get() *TransfersResponse {
	return v.value
}

func (v *NullableTransfersResponse) Set(val *TransfersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfersResponse(val *TransfersResponse) *NullableTransfersResponse {
	return &NullableTransfersResponse{value: val, isSet: true}
}

func (v NullableTransfersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

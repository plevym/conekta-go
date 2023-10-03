package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodBankTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodBankTransfer{}

// PaymentMethodBankTransfer struct for PaymentMethodBankTransfer
type PaymentMethodBankTransfer struct {
	Type                       *string        `json:"type,omitempty"`
	Object                     string         `json:"object"`
	Bank                       *string        `json:"bank,omitempty"`
	Clabe                      *string        `json:"clabe,omitempty"`
	Description                NullableString `json:"description,omitempty"`
	ExecutedAt                 NullableInt32  `json:"executed_at,omitempty"`
	ExpiresAt                  *int64         `json:"expires_at,omitempty"`
	IssuingAccountBank         NullableString `json:"issuing_account_bank,omitempty"`
	IssuingAccountNumber       NullableString `json:"issuing_account_number,omitempty"`
	IssuingAccountHolderName   NullableString `json:"issuing_account_holder_name,omitempty"`
	IssuingAccountTaxId        NullableString `json:"issuing_account_tax_id,omitempty"`
	PaymentAttempts            []interface{}  `json:"payment_attempts,omitempty"`
	ReceivingAccountHolderName NullableString `json:"receiving_account_holder_name,omitempty"`
	ReceivingAccountNumber     *string        `json:"receiving_account_number,omitempty"`
	ReceivingAccountBank       *string        `json:"receiving_account_bank,omitempty"`
	ReceivingAccountTaxId      NullableString `json:"receiving_account_tax_id,omitempty"`
	ReferenceNumber            NullableString `json:"reference_number,omitempty"`
	TrackingCode               NullableString `json:"tracking_code,omitempty"`
}

// NewPaymentMethodBankTransfer instantiates a new PaymentMethodBankTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodBankTransfer(object string) *PaymentMethodBankTransfer {
	this := PaymentMethodBankTransfer{}
	this.Object = object
	return &this
}

// NewPaymentMethodBankTransferWithDefaults instantiates a new PaymentMethodBankTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodBankTransferWithDefaults() *PaymentMethodBankTransfer {
	this := PaymentMethodBankTransfer{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodBankTransfer) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value
func (o *PaymentMethodBankTransfer) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodBankTransfer) SetObject(v string) {
	o.Object = v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetBank() string {
	if o == nil || IsNil(o.Bank) {
		var ret string
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetBankOk() (*string, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given string and assigns it to the Bank field.
func (o *PaymentMethodBankTransfer) SetBank(v string) {
	o.Bank = &v
}

// GetClabe returns the Clabe field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetClabe() string {
	if o == nil || IsNil(o.Clabe) {
		var ret string
		return ret
	}
	return *o.Clabe
}

// GetClabeOk returns a tuple with the Clabe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetClabeOk() (*string, bool) {
	if o == nil || IsNil(o.Clabe) {
		return nil, false
	}
	return o.Clabe, true
}

// HasClabe returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasClabe() bool {
	if o != nil && !IsNil(o.Clabe) {
		return true
	}

	return false
}

// SetClabe gets a reference to the given string and assigns it to the Clabe field.
func (o *PaymentMethodBankTransfer) SetClabe(v string) {
	o.Clabe = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PaymentMethodBankTransfer) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PaymentMethodBankTransfer) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetDescription() {
	o.Description.Unset()
}

// GetExecutedAt returns the ExecutedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetExecutedAt() int32 {
	if o == nil || IsNil(o.ExecutedAt.Get()) {
		var ret int32
		return ret
	}
	return *o.ExecutedAt.Get()
}

// GetExecutedAtOk returns a tuple with the ExecutedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetExecutedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutedAt.Get(), o.ExecutedAt.IsSet()
}

// HasExecutedAt returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasExecutedAt() bool {
	if o != nil && o.ExecutedAt.IsSet() {
		return true
	}

	return false
}

// SetExecutedAt gets a reference to the given NullableInt32 and assigns it to the ExecutedAt field.
func (o *PaymentMethodBankTransfer) SetExecutedAt(v int32) {
	o.ExecutedAt.Set(&v)
}

// SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil
func (o *PaymentMethodBankTransfer) SetExecutedAtNil() {
	o.ExecutedAt.Set(nil)
}

// UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetExecutedAt() {
	o.ExecutedAt.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PaymentMethodBankTransfer) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetIssuingAccountBank returns the IssuingAccountBank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetIssuingAccountBank() string {
	if o == nil || IsNil(o.IssuingAccountBank.Get()) {
		var ret string
		return ret
	}
	return *o.IssuingAccountBank.Get()
}

// GetIssuingAccountBankOk returns a tuple with the IssuingAccountBank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetIssuingAccountBankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuingAccountBank.Get(), o.IssuingAccountBank.IsSet()
}

// HasIssuingAccountBank returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasIssuingAccountBank() bool {
	if o != nil && o.IssuingAccountBank.IsSet() {
		return true
	}

	return false
}

// SetIssuingAccountBank gets a reference to the given NullableString and assigns it to the IssuingAccountBank field.
func (o *PaymentMethodBankTransfer) SetIssuingAccountBank(v string) {
	o.IssuingAccountBank.Set(&v)
}

// SetIssuingAccountBankNil sets the value for IssuingAccountBank to be an explicit nil
func (o *PaymentMethodBankTransfer) SetIssuingAccountBankNil() {
	o.IssuingAccountBank.Set(nil)
}

// UnsetIssuingAccountBank ensures that no value is present for IssuingAccountBank, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetIssuingAccountBank() {
	o.IssuingAccountBank.Unset()
}

// GetIssuingAccountNumber returns the IssuingAccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetIssuingAccountNumber() string {
	if o == nil || IsNil(o.IssuingAccountNumber.Get()) {
		var ret string
		return ret
	}
	return *o.IssuingAccountNumber.Get()
}

// GetIssuingAccountNumberOk returns a tuple with the IssuingAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetIssuingAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuingAccountNumber.Get(), o.IssuingAccountNumber.IsSet()
}

// HasIssuingAccountNumber returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasIssuingAccountNumber() bool {
	if o != nil && o.IssuingAccountNumber.IsSet() {
		return true
	}

	return false
}

// SetIssuingAccountNumber gets a reference to the given NullableString and assigns it to the IssuingAccountNumber field.
func (o *PaymentMethodBankTransfer) SetIssuingAccountNumber(v string) {
	o.IssuingAccountNumber.Set(&v)
}

// SetIssuingAccountNumberNil sets the value for IssuingAccountNumber to be an explicit nil
func (o *PaymentMethodBankTransfer) SetIssuingAccountNumberNil() {
	o.IssuingAccountNumber.Set(nil)
}

// UnsetIssuingAccountNumber ensures that no value is present for IssuingAccountNumber, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetIssuingAccountNumber() {
	o.IssuingAccountNumber.Unset()
}

// GetIssuingAccountHolderName returns the IssuingAccountHolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetIssuingAccountHolderName() string {
	if o == nil || IsNil(o.IssuingAccountHolderName.Get()) {
		var ret string
		return ret
	}
	return *o.IssuingAccountHolderName.Get()
}

// GetIssuingAccountHolderNameOk returns a tuple with the IssuingAccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetIssuingAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuingAccountHolderName.Get(), o.IssuingAccountHolderName.IsSet()
}

// HasIssuingAccountHolderName returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasIssuingAccountHolderName() bool {
	if o != nil && o.IssuingAccountHolderName.IsSet() {
		return true
	}

	return false
}

// SetIssuingAccountHolderName gets a reference to the given NullableString and assigns it to the IssuingAccountHolderName field.
func (o *PaymentMethodBankTransfer) SetIssuingAccountHolderName(v string) {
	o.IssuingAccountHolderName.Set(&v)
}

// SetIssuingAccountHolderNameNil sets the value for IssuingAccountHolderName to be an explicit nil
func (o *PaymentMethodBankTransfer) SetIssuingAccountHolderNameNil() {
	o.IssuingAccountHolderName.Set(nil)
}

// UnsetIssuingAccountHolderName ensures that no value is present for IssuingAccountHolderName, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetIssuingAccountHolderName() {
	o.IssuingAccountHolderName.Unset()
}

// GetIssuingAccountTaxId returns the IssuingAccountTaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetIssuingAccountTaxId() string {
	if o == nil || IsNil(o.IssuingAccountTaxId.Get()) {
		var ret string
		return ret
	}
	return *o.IssuingAccountTaxId.Get()
}

// GetIssuingAccountTaxIdOk returns a tuple with the IssuingAccountTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetIssuingAccountTaxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuingAccountTaxId.Get(), o.IssuingAccountTaxId.IsSet()
}

// HasIssuingAccountTaxId returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasIssuingAccountTaxId() bool {
	if o != nil && o.IssuingAccountTaxId.IsSet() {
		return true
	}

	return false
}

// SetIssuingAccountTaxId gets a reference to the given NullableString and assigns it to the IssuingAccountTaxId field.
func (o *PaymentMethodBankTransfer) SetIssuingAccountTaxId(v string) {
	o.IssuingAccountTaxId.Set(&v)
}

// SetIssuingAccountTaxIdNil sets the value for IssuingAccountTaxId to be an explicit nil
func (o *PaymentMethodBankTransfer) SetIssuingAccountTaxIdNil() {
	o.IssuingAccountTaxId.Set(nil)
}

// UnsetIssuingAccountTaxId ensures that no value is present for IssuingAccountTaxId, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetIssuingAccountTaxId() {
	o.IssuingAccountTaxId.Unset()
}

// GetPaymentAttempts returns the PaymentAttempts field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetPaymentAttempts() []interface{} {
	if o == nil || IsNil(o.PaymentAttempts) {
		var ret []interface{}
		return ret
	}
	return o.PaymentAttempts
}

// GetPaymentAttemptsOk returns a tuple with the PaymentAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetPaymentAttemptsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.PaymentAttempts) {
		return nil, false
	}
	return o.PaymentAttempts, true
}

// HasPaymentAttempts returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasPaymentAttempts() bool {
	if o != nil && !IsNil(o.PaymentAttempts) {
		return true
	}

	return false
}

// SetPaymentAttempts gets a reference to the given []interface{} and assigns it to the PaymentAttempts field.
func (o *PaymentMethodBankTransfer) SetPaymentAttempts(v []interface{}) {
	o.PaymentAttempts = v
}

// GetReceivingAccountHolderName returns the ReceivingAccountHolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetReceivingAccountHolderName() string {
	if o == nil || IsNil(o.ReceivingAccountHolderName.Get()) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountHolderName.Get()
}

// GetReceivingAccountHolderNameOk returns a tuple with the ReceivingAccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetReceivingAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivingAccountHolderName.Get(), o.ReceivingAccountHolderName.IsSet()
}

// HasReceivingAccountHolderName returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasReceivingAccountHolderName() bool {
	if o != nil && o.ReceivingAccountHolderName.IsSet() {
		return true
	}

	return false
}

// SetReceivingAccountHolderName gets a reference to the given NullableString and assigns it to the ReceivingAccountHolderName field.
func (o *PaymentMethodBankTransfer) SetReceivingAccountHolderName(v string) {
	o.ReceivingAccountHolderName.Set(&v)
}

// SetReceivingAccountHolderNameNil sets the value for ReceivingAccountHolderName to be an explicit nil
func (o *PaymentMethodBankTransfer) SetReceivingAccountHolderNameNil() {
	o.ReceivingAccountHolderName.Set(nil)
}

// UnsetReceivingAccountHolderName ensures that no value is present for ReceivingAccountHolderName, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetReceivingAccountHolderName() {
	o.ReceivingAccountHolderName.Unset()
}

// GetReceivingAccountNumber returns the ReceivingAccountNumber field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetReceivingAccountNumber() string {
	if o == nil || IsNil(o.ReceivingAccountNumber) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountNumber
}

// GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetReceivingAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountNumber) {
		return nil, false
	}
	return o.ReceivingAccountNumber, true
}

// HasReceivingAccountNumber returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasReceivingAccountNumber() bool {
	if o != nil && !IsNil(o.ReceivingAccountNumber) {
		return true
	}

	return false
}

// SetReceivingAccountNumber gets a reference to the given string and assigns it to the ReceivingAccountNumber field.
func (o *PaymentMethodBankTransfer) SetReceivingAccountNumber(v string) {
	o.ReceivingAccountNumber = &v
}

// GetReceivingAccountBank returns the ReceivingAccountBank field value if set, zero value otherwise.
func (o *PaymentMethodBankTransfer) GetReceivingAccountBank() string {
	if o == nil || IsNil(o.ReceivingAccountBank) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountBank
}

// GetReceivingAccountBankOk returns a tuple with the ReceivingAccountBank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBankTransfer) GetReceivingAccountBankOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountBank) {
		return nil, false
	}
	return o.ReceivingAccountBank, true
}

// HasReceivingAccountBank returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasReceivingAccountBank() bool {
	if o != nil && !IsNil(o.ReceivingAccountBank) {
		return true
	}

	return false
}

// SetReceivingAccountBank gets a reference to the given string and assigns it to the ReceivingAccountBank field.
func (o *PaymentMethodBankTransfer) SetReceivingAccountBank(v string) {
	o.ReceivingAccountBank = &v
}

// GetReceivingAccountTaxId returns the ReceivingAccountTaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetReceivingAccountTaxId() string {
	if o == nil || IsNil(o.ReceivingAccountTaxId.Get()) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountTaxId.Get()
}

// GetReceivingAccountTaxIdOk returns a tuple with the ReceivingAccountTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetReceivingAccountTaxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivingAccountTaxId.Get(), o.ReceivingAccountTaxId.IsSet()
}

// HasReceivingAccountTaxId returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasReceivingAccountTaxId() bool {
	if o != nil && o.ReceivingAccountTaxId.IsSet() {
		return true
	}

	return false
}

// SetReceivingAccountTaxId gets a reference to the given NullableString and assigns it to the ReceivingAccountTaxId field.
func (o *PaymentMethodBankTransfer) SetReceivingAccountTaxId(v string) {
	o.ReceivingAccountTaxId.Set(&v)
}

// SetReceivingAccountTaxIdNil sets the value for ReceivingAccountTaxId to be an explicit nil
func (o *PaymentMethodBankTransfer) SetReceivingAccountTaxIdNil() {
	o.ReceivingAccountTaxId.Set(nil)
}

// UnsetReceivingAccountTaxId ensures that no value is present for ReceivingAccountTaxId, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetReceivingAccountTaxId() {
	o.ReceivingAccountTaxId.Unset()
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber.Get()
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetReferenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceNumber.Get(), o.ReferenceNumber.IsSet()
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasReferenceNumber() bool {
	if o != nil && o.ReferenceNumber.IsSet() {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given NullableString and assigns it to the ReferenceNumber field.
func (o *PaymentMethodBankTransfer) SetReferenceNumber(v string) {
	o.ReferenceNumber.Set(&v)
}

// SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil
func (o *PaymentMethodBankTransfer) SetReferenceNumberNil() {
	o.ReferenceNumber.Set(nil)
}

// UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetReferenceNumber() {
	o.ReferenceNumber.Unset()
}

// GetTrackingCode returns the TrackingCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodBankTransfer) GetTrackingCode() string {
	if o == nil || IsNil(o.TrackingCode.Get()) {
		var ret string
		return ret
	}
	return *o.TrackingCode.Get()
}

// GetTrackingCodeOk returns a tuple with the TrackingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodBankTransfer) GetTrackingCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingCode.Get(), o.TrackingCode.IsSet()
}

// HasTrackingCode returns a boolean if a field has been set.
func (o *PaymentMethodBankTransfer) HasTrackingCode() bool {
	if o != nil && o.TrackingCode.IsSet() {
		return true
	}

	return false
}

// SetTrackingCode gets a reference to the given NullableString and assigns it to the TrackingCode field.
func (o *PaymentMethodBankTransfer) SetTrackingCode(v string) {
	o.TrackingCode.Set(&v)
}

// SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil
func (o *PaymentMethodBankTransfer) SetTrackingCodeNil() {
	o.TrackingCode.Set(nil)
}

// UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
func (o *PaymentMethodBankTransfer) UnsetTrackingCode() {
	o.TrackingCode.Unset()
}

func (o PaymentMethodBankTransfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodBankTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.Clabe) {
		toSerialize["clabe"] = o.Clabe
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ExecutedAt.IsSet() {
		toSerialize["executed_at"] = o.ExecutedAt.Get()
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.IssuingAccountBank.IsSet() {
		toSerialize["issuing_account_bank"] = o.IssuingAccountBank.Get()
	}
	if o.IssuingAccountNumber.IsSet() {
		toSerialize["issuing_account_number"] = o.IssuingAccountNumber.Get()
	}
	if o.IssuingAccountHolderName.IsSet() {
		toSerialize["issuing_account_holder_name"] = o.IssuingAccountHolderName.Get()
	}
	if o.IssuingAccountTaxId.IsSet() {
		toSerialize["issuing_account_tax_id"] = o.IssuingAccountTaxId.Get()
	}
	if !IsNil(o.PaymentAttempts) {
		toSerialize["payment_attempts"] = o.PaymentAttempts
	}
	if o.ReceivingAccountHolderName.IsSet() {
		toSerialize["receiving_account_holder_name"] = o.ReceivingAccountHolderName.Get()
	}
	if !IsNil(o.ReceivingAccountNumber) {
		toSerialize["receiving_account_number"] = o.ReceivingAccountNumber
	}
	if !IsNil(o.ReceivingAccountBank) {
		toSerialize["receiving_account_bank"] = o.ReceivingAccountBank
	}
	if o.ReceivingAccountTaxId.IsSet() {
		toSerialize["receiving_account_tax_id"] = o.ReceivingAccountTaxId.Get()
	}
	if o.ReferenceNumber.IsSet() {
		toSerialize["reference_number"] = o.ReferenceNumber.Get()
	}
	if o.TrackingCode.IsSet() {
		toSerialize["tracking_code"] = o.TrackingCode.Get()
	}
	return toSerialize, nil
}

type NullablePaymentMethodBankTransfer struct {
	value *PaymentMethodBankTransfer
	isSet bool
}

func (v NullablePaymentMethodBankTransfer) Get() *PaymentMethodBankTransfer {
	return v.value
}

func (v *NullablePaymentMethodBankTransfer) Set(val *PaymentMethodBankTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodBankTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodBankTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodBankTransfer(val *PaymentMethodBankTransfer) *NullablePaymentMethodBankTransfer {
	return &NullablePaymentMethodBankTransfer{value: val, isSet: true}
}

func (v NullablePaymentMethodBankTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodBankTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

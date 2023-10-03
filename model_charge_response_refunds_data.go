package conekta

import (
	"encoding/json"
)

// checks if the ChargeResponseRefundsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeResponseRefundsData{}

// ChargeResponseRefundsData struct for ChargeResponseRefundsData
type ChargeResponseRefundsData struct {
	Amount    int64   `json:"amount"`
	AuthCode  *string `json:"auth_code,omitempty"`
	CreatedAt int64   `json:"created_at"`
	// refund expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	Id        string `json:"id"`
	Object    string `json:"object"`
	// refund status
	Status *string `json:"status,omitempty"`
}

// NewChargeResponseRefundsData instantiates a new ChargeResponseRefundsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeResponseRefundsData(amount int64, createdAt int64, id string, object string) *ChargeResponseRefundsData {
	this := ChargeResponseRefundsData{}
	this.Amount = amount
	this.CreatedAt = createdAt
	this.Id = id
	this.Object = object
	return &this
}

// NewChargeResponseRefundsDataWithDefaults instantiates a new ChargeResponseRefundsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeResponseRefundsDataWithDefaults() *ChargeResponseRefundsData {
	this := ChargeResponseRefundsData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ChargeResponseRefundsData) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ChargeResponseRefundsData) SetAmount(v int64) {
	o.Amount = v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *ChargeResponseRefundsData) GetAuthCode() string {
	if o == nil || IsNil(o.AuthCode) {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetAuthCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthCode) {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *ChargeResponseRefundsData) HasAuthCode() bool {
	if o != nil && !IsNil(o.AuthCode) {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *ChargeResponseRefundsData) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChargeResponseRefundsData) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChargeResponseRefundsData) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ChargeResponseRefundsData) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ChargeResponseRefundsData) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *ChargeResponseRefundsData) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *ChargeResponseRefundsData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChargeResponseRefundsData) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *ChargeResponseRefundsData) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ChargeResponseRefundsData) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChargeResponseRefundsData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChargeResponseRefundsData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ChargeResponseRefundsData) SetStatus(v string) {
	o.Status = &v
}

func (o ChargeResponseRefundsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeResponseRefundsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.AuthCode) {
		toSerialize["auth_code"] = o.AuthCode
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChargeResponseRefundsData struct {
	value *ChargeResponseRefundsData
	isSet bool
}

func (v NullableChargeResponseRefundsData) Get() *ChargeResponseRefundsData {
	return v.value
}

func (v *NullableChargeResponseRefundsData) Set(val *ChargeResponseRefundsData) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeResponseRefundsData) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeResponseRefundsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeResponseRefundsData(val *ChargeResponseRefundsData) *NullableChargeResponseRefundsData {
	return &NullableChargeResponseRefundsData{value: val, isSet: true}
}

func (v NullableChargeResponseRefundsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeResponseRefundsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

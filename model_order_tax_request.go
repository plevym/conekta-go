package conekta

import (
	"encoding/json"
)

// checks if the OrderTaxRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTaxRequest{}

// OrderTaxRequest create new taxes for an existing order
type OrderTaxRequest struct {
	// The amount to be collected for tax in cents
	Amount int64 `json:"amount"`
	// description or tax's name
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// NewOrderTaxRequest instantiates a new OrderTaxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTaxRequest(amount int64, description string) *OrderTaxRequest {
	this := OrderTaxRequest{}
	this.Amount = amount
	this.Description = description
	return &this
}

// NewOrderTaxRequestWithDefaults instantiates a new OrderTaxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTaxRequestWithDefaults() *OrderTaxRequest {
	this := OrderTaxRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OrderTaxRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrderTaxRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrderTaxRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *OrderTaxRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrderTaxRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OrderTaxRequest) SetDescription(v string) {
	o.Description = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderTaxRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderTaxRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *OrderTaxRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o OrderTaxRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTaxRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["description"] = o.Description
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableOrderTaxRequest struct {
	value *OrderTaxRequest
	isSet bool
}

func (v NullableOrderTaxRequest) Get() *OrderTaxRequest {
	return v.value
}

func (v *NullableOrderTaxRequest) Set(val *OrderTaxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTaxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTaxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTaxRequest(val *OrderTaxRequest) *NullableOrderTaxRequest {
	return &NullableOrderTaxRequest{value: val, isSet: true}
}

func (v NullableOrderTaxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTaxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

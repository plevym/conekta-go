package conekta

import (
	"encoding/json"
)

// checks if the UpdateOrderTaxRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderTaxRequest{}

// UpdateOrderTaxRequest create new taxes for an existing order
type UpdateOrderTaxRequest struct {
	// The amount to be collected for tax in cents
	Amount *int64 `json:"amount,omitempty"`
	// description or tax's name
	Description *string                           `json:"description,omitempty"`
	Metadata    map[string]map[string]interface{} `json:"metadata,omitempty"`
}

// NewUpdateOrderTaxRequest instantiates a new UpdateOrderTaxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderTaxRequest() *UpdateOrderTaxRequest {
	this := UpdateOrderTaxRequest{}
	return &this
}

// NewUpdateOrderTaxRequestWithDefaults instantiates a new UpdateOrderTaxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderTaxRequestWithDefaults() *UpdateOrderTaxRequest {
	this := UpdateOrderTaxRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UpdateOrderTaxRequest) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxRequest) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UpdateOrderTaxRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UpdateOrderTaxRequest) SetAmount(v int64) {
	o.Amount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateOrderTaxRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateOrderTaxRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateOrderTaxRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateOrderTaxRequest) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxRequest) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateOrderTaxRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateOrderTaxRequest) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

func (o UpdateOrderTaxRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderTaxRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableUpdateOrderTaxRequest struct {
	value *UpdateOrderTaxRequest
	isSet bool
}

func (v NullableUpdateOrderTaxRequest) Get() *UpdateOrderTaxRequest {
	return v.value
}

func (v *NullableUpdateOrderTaxRequest) Set(val *UpdateOrderTaxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderTaxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderTaxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderTaxRequest(val *UpdateOrderTaxRequest) *NullableUpdateOrderTaxRequest {
	return &NullableUpdateOrderTaxRequest{value: val, isSet: true}
}

func (v NullableUpdateOrderTaxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderTaxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseFiscalEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseFiscalEntity{}

// OrderResponseFiscalEntity struct for OrderResponseFiscalEntity
type OrderResponseFiscalEntity struct {
	Address *OrderResponseFiscalEntityAddress `json:"address,omitempty"`
	TaxId   *string                           `json:"tax_id,omitempty"`
	Id      *string                           `json:"id,omitempty"`
	Object  *string                           `json:"object,omitempty"`
}

// NewOrderResponseFiscalEntity instantiates a new OrderResponseFiscalEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseFiscalEntity() *OrderResponseFiscalEntity {
	this := OrderResponseFiscalEntity{}
	return &this
}

// NewOrderResponseFiscalEntityWithDefaults instantiates a new OrderResponseFiscalEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseFiscalEntityWithDefaults() *OrderResponseFiscalEntity {
	this := OrderResponseFiscalEntity{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntity) GetAddress() OrderResponseFiscalEntityAddress {
	if o == nil || IsNil(o.Address) {
		var ret OrderResponseFiscalEntityAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntity) GetAddressOk() (*OrderResponseFiscalEntityAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntity) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given OrderResponseFiscalEntityAddress and assigns it to the Address field.
func (o *OrderResponseFiscalEntity) SetAddress(v OrderResponseFiscalEntityAddress) {
	o.Address = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntity) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntity) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntity) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *OrderResponseFiscalEntity) SetTaxId(v string) {
	o.TaxId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponseFiscalEntity) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntity) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntity) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntity) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseFiscalEntity) SetObject(v string) {
	o.Object = &v
}

func (o OrderResponseFiscalEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseFiscalEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableOrderResponseFiscalEntity struct {
	value *OrderResponseFiscalEntity
	isSet bool
}

func (v NullableOrderResponseFiscalEntity) Get() *OrderResponseFiscalEntity {
	return v.value
}

func (v *NullableOrderResponseFiscalEntity) Set(val *OrderResponseFiscalEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseFiscalEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseFiscalEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseFiscalEntity(val *OrderResponseFiscalEntity) *NullableOrderResponseFiscalEntity {
	return &NullableOrderResponseFiscalEntity{value: val, isSet: true}
}

func (v NullableOrderResponseFiscalEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseFiscalEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

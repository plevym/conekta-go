package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseShippingContactAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseShippingContactAllOf{}

// OrderResponseShippingContactAllOf struct for OrderResponseShippingContactAllOf
type OrderResponseShippingContactAllOf struct {
	CreatedAt *int64  `json:"created_at,omitempty"`
	Id        *string `json:"id,omitempty"`
	Object    *string `json:"object,omitempty"`
}

// NewOrderResponseShippingContactAllOf instantiates a new OrderResponseShippingContactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseShippingContactAllOf() *OrderResponseShippingContactAllOf {
	this := OrderResponseShippingContactAllOf{}
	return &this
}

// NewOrderResponseShippingContactAllOfWithDefaults instantiates a new OrderResponseShippingContactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseShippingContactAllOfWithDefaults() *OrderResponseShippingContactAllOf {
	this := OrderResponseShippingContactAllOf{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderResponseShippingContactAllOf) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContactAllOf) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderResponseShippingContactAllOf) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *OrderResponseShippingContactAllOf) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponseShippingContactAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContactAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponseShippingContactAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponseShippingContactAllOf) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseShippingContactAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContactAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseShippingContactAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseShippingContactAllOf) SetObject(v string) {
	o.Object = &v
}

func (o OrderResponseShippingContactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseShippingContactAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableOrderResponseShippingContactAllOf struct {
	value *OrderResponseShippingContactAllOf
	isSet bool
}

func (v NullableOrderResponseShippingContactAllOf) Get() *OrderResponseShippingContactAllOf {
	return v.value
}

func (v *NullableOrderResponseShippingContactAllOf) Set(val *OrderResponseShippingContactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseShippingContactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseShippingContactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseShippingContactAllOf(val *OrderResponseShippingContactAllOf) *NullableOrderResponseShippingContactAllOf {
	return &NullableOrderResponseShippingContactAllOf{value: val, isSet: true}
}

func (v NullableOrderResponseShippingContactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseShippingContactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

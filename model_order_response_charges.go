package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseCharges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseCharges{}

// OrderResponseCharges The charges associated with the order
type OrderResponseCharges struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string                `json:"object"`
	Data   []ChargesDataResponse `json:"data,omitempty"`
}

// NewOrderResponseCharges instantiates a new OrderResponseCharges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseCharges(hasMore bool, object string) *OrderResponseCharges {
	this := OrderResponseCharges{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewOrderResponseChargesWithDefaults instantiates a new OrderResponseCharges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseChargesWithDefaults() *OrderResponseCharges {
	this := OrderResponseCharges{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *OrderResponseCharges) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *OrderResponseCharges) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *OrderResponseCharges) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *OrderResponseCharges) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *OrderResponseCharges) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *OrderResponseCharges) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseCharges) GetData() []ChargesDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []ChargesDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCharges) GetDataOk() ([]ChargesDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseCharges) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ChargesDataResponse and assigns it to the Data field.
func (o *OrderResponseCharges) SetData(v []ChargesDataResponse) {
	o.Data = v
}

func (o OrderResponseCharges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseCharges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderResponseCharges struct {
	value *OrderResponseCharges
	isSet bool
}

func (v NullableOrderResponseCharges) Get() *OrderResponseCharges {
	return v.value
}

func (v *NullableOrderResponseCharges) Set(val *OrderResponseCharges) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseCharges) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseCharges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseCharges(val *OrderResponseCharges) *NullableOrderResponseCharges {
	return &NullableOrderResponseCharges{value: val, isSet: true}
}

func (v NullableOrderResponseCharges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseCharges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

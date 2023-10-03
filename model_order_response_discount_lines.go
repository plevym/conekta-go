package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseDiscountLines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseDiscountLines{}

// OrderResponseDiscountLines struct for OrderResponseDiscountLines
type OrderResponseDiscountLines struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string                      `json:"object"`
	Data   []DiscountLinesDataResponse `json:"data,omitempty"`
}

// NewOrderResponseDiscountLines instantiates a new OrderResponseDiscountLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDiscountLines(hasMore bool, object string) *OrderResponseDiscountLines {
	this := OrderResponseDiscountLines{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewOrderResponseDiscountLinesWithDefaults instantiates a new OrderResponseDiscountLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDiscountLinesWithDefaults() *OrderResponseDiscountLines {
	this := OrderResponseDiscountLines{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *OrderResponseDiscountLines) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *OrderResponseDiscountLines) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *OrderResponseDiscountLines) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *OrderResponseDiscountLines) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *OrderResponseDiscountLines) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *OrderResponseDiscountLines) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseDiscountLines) GetData() []DiscountLinesDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []DiscountLinesDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDiscountLines) GetDataOk() ([]DiscountLinesDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseDiscountLines) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DiscountLinesDataResponse and assigns it to the Data field.
func (o *OrderResponseDiscountLines) SetData(v []DiscountLinesDataResponse) {
	o.Data = v
}

func (o OrderResponseDiscountLines) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseDiscountLines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderResponseDiscountLines struct {
	value *OrderResponseDiscountLines
	isSet bool
}

func (v NullableOrderResponseDiscountLines) Get() *OrderResponseDiscountLines {
	return v.value
}

func (v *NullableOrderResponseDiscountLines) Set(val *OrderResponseDiscountLines) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDiscountLines) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDiscountLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDiscountLines(val *OrderResponseDiscountLines) *NullableOrderResponseDiscountLines {
	return &NullableOrderResponseDiscountLines{value: val, isSet: true}
}

func (v NullableOrderResponseDiscountLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDiscountLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

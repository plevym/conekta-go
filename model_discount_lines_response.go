package conekta

import (
	"encoding/json"
)

// checks if the DiscountLinesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountLinesResponse{}

// DiscountLinesResponse struct for DiscountLinesResponse
type DiscountLinesResponse struct {
	// The amount to be deducted from the total sum of all payments, in cents.
	Amount int64 `json:"amount"`
	// Discount code.
	Code string `json:"code"`
	// It can be 'loyalty', 'campaign', 'coupon' o 'sign'
	Type string `json:"type"`
	// The discount line id
	Id string `json:"id"`
	// The object name
	Object string `json:"object"`
	// The order id
	ParentId string `json:"parent_id"`
}

// NewDiscountLinesResponse instantiates a new DiscountLinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountLinesResponse(amount int64, code string, type_ string, id string, object string, parentId string) *DiscountLinesResponse {
	this := DiscountLinesResponse{}
	this.Amount = amount
	this.Code = code
	this.Type = type_
	this.Id = id
	this.Object = object
	this.ParentId = parentId
	return &this
}

// NewDiscountLinesResponseWithDefaults instantiates a new DiscountLinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountLinesResponseWithDefaults() *DiscountLinesResponse {
	this := DiscountLinesResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *DiscountLinesResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DiscountLinesResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCode returns the Code field value
func (o *DiscountLinesResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DiscountLinesResponse) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *DiscountLinesResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiscountLinesResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DiscountLinesResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscountLinesResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *DiscountLinesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DiscountLinesResponse) SetObject(v string) {
	o.Object = v
}

// GetParentId returns the ParentId field value
func (o *DiscountLinesResponse) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponse) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *DiscountLinesResponse) SetParentId(v string) {
	o.ParentId = v
}

func (o DiscountLinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountLinesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["code"] = o.Code
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["parent_id"] = o.ParentId
	return toSerialize, nil
}

type NullableDiscountLinesResponse struct {
	value *DiscountLinesResponse
	isSet bool
}

func (v NullableDiscountLinesResponse) Get() *DiscountLinesResponse {
	return v.value
}

func (v *NullableDiscountLinesResponse) Set(val *DiscountLinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountLinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountLinesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountLinesResponse(val *DiscountLinesResponse) *NullableDiscountLinesResponse {
	return &NullableDiscountLinesResponse{value: val, isSet: true}
}

func (v NullableDiscountLinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountLinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

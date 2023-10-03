package conekta

import (
	"encoding/json"
)

// checks if the DiscountLinesDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountLinesDataResponse{}

// DiscountLinesDataResponse struct for DiscountLinesDataResponse
type DiscountLinesDataResponse struct {
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

// NewDiscountLinesDataResponse instantiates a new DiscountLinesDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountLinesDataResponse(amount int64, code string, type_ string, id string, object string, parentId string) *DiscountLinesDataResponse {
	this := DiscountLinesDataResponse{}
	this.Amount = amount
	this.Code = code
	this.Type = type_
	this.Id = id
	this.Object = object
	this.ParentId = parentId
	return &this
}

// NewDiscountLinesDataResponseWithDefaults instantiates a new DiscountLinesDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountLinesDataResponseWithDefaults() *DiscountLinesDataResponse {
	this := DiscountLinesDataResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *DiscountLinesDataResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesDataResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DiscountLinesDataResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCode returns the Code field value
func (o *DiscountLinesDataResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesDataResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DiscountLinesDataResponse) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *DiscountLinesDataResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesDataResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiscountLinesDataResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DiscountLinesDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscountLinesDataResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *DiscountLinesDataResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesDataResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DiscountLinesDataResponse) SetObject(v string) {
	o.Object = v
}

// GetParentId returns the ParentId field value
func (o *DiscountLinesDataResponse) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesDataResponse) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *DiscountLinesDataResponse) SetParentId(v string) {
	o.ParentId = v
}

func (o DiscountLinesDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountLinesDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["code"] = o.Code
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["parent_id"] = o.ParentId
	return toSerialize, nil
}

type NullableDiscountLinesDataResponse struct {
	value *DiscountLinesDataResponse
	isSet bool
}

func (v NullableDiscountLinesDataResponse) Get() *DiscountLinesDataResponse {
	return v.value
}

func (v *NullableDiscountLinesDataResponse) Set(val *DiscountLinesDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountLinesDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountLinesDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountLinesDataResponse(val *DiscountLinesDataResponse) *NullableDiscountLinesDataResponse {
	return &NullableDiscountLinesDataResponse{value: val, isSet: true}
}

func (v NullableDiscountLinesDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountLinesDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

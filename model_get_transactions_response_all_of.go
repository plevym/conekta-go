package conekta

import (
	"encoding/json"
)

// checks if the GetTransactionsResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransactionsResponseAllOf{}

// GetTransactionsResponseAllOf struct for GetTransactionsResponseAllOf
type GetTransactionsResponseAllOf struct {
	// Transactions
	Data []TransactionResponse `json:"data,omitempty"`
}

// NewGetTransactionsResponseAllOf instantiates a new GetTransactionsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionsResponseAllOf() *GetTransactionsResponseAllOf {
	this := GetTransactionsResponseAllOf{}
	return &this
}

// NewGetTransactionsResponseAllOfWithDefaults instantiates a new GetTransactionsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionsResponseAllOfWithDefaults() *GetTransactionsResponseAllOf {
	this := GetTransactionsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTransactionsResponseAllOf) GetData() []TransactionResponse {
	if o == nil || IsNil(o.Data) {
		var ret []TransactionResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionsResponseAllOf) GetDataOk() ([]TransactionResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTransactionsResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TransactionResponse and assigns it to the Data field.
func (o *GetTransactionsResponseAllOf) SetData(v []TransactionResponse) {
	o.Data = v
}

func (o GetTransactionsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransactionsResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetTransactionsResponseAllOf struct {
	value *GetTransactionsResponseAllOf
	isSet bool
}

func (v NullableGetTransactionsResponseAllOf) Get() *GetTransactionsResponseAllOf {
	return v.value
}

func (v *NullableGetTransactionsResponseAllOf) Set(val *GetTransactionsResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionsResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionsResponseAllOf(val *GetTransactionsResponseAllOf) *NullableGetTransactionsResponseAllOf {
	return &NullableGetTransactionsResponseAllOf{value: val, isSet: true}
}

func (v NullableGetTransactionsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

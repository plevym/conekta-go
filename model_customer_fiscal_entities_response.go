package conekta

import (
	"encoding/json"
)

// checks if the CustomerFiscalEntitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerFiscalEntitiesResponse{}

// CustomerFiscalEntitiesResponse struct for CustomerFiscalEntitiesResponse
type CustomerFiscalEntitiesResponse struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string                               `json:"object"`
	Data   []CustomerFiscalEntitiesDataResponse `json:"data,omitempty"`
}

// NewCustomerFiscalEntitiesResponse instantiates a new CustomerFiscalEntitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFiscalEntitiesResponse(hasMore bool, object string) *CustomerFiscalEntitiesResponse {
	this := CustomerFiscalEntitiesResponse{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewCustomerFiscalEntitiesResponseWithDefaults instantiates a new CustomerFiscalEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFiscalEntitiesResponseWithDefaults() *CustomerFiscalEntitiesResponse {
	this := CustomerFiscalEntitiesResponse{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *CustomerFiscalEntitiesResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CustomerFiscalEntitiesResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *CustomerFiscalEntitiesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerFiscalEntitiesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesResponse) GetData() []CustomerFiscalEntitiesDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerFiscalEntitiesDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesResponse) GetDataOk() ([]CustomerFiscalEntitiesDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerFiscalEntitiesDataResponse and assigns it to the Data field.
func (o *CustomerFiscalEntitiesResponse) SetData(v []CustomerFiscalEntitiesDataResponse) {
	o.Data = v
}

func (o CustomerFiscalEntitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerFiscalEntitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerFiscalEntitiesResponse struct {
	value *CustomerFiscalEntitiesResponse
	isSet bool
}

func (v NullableCustomerFiscalEntitiesResponse) Get() *CustomerFiscalEntitiesResponse {
	return v.value
}

func (v *NullableCustomerFiscalEntitiesResponse) Set(val *CustomerFiscalEntitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFiscalEntitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFiscalEntitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFiscalEntitiesResponse(val *CustomerFiscalEntitiesResponse) *NullableCustomerFiscalEntitiesResponse {
	return &NullableCustomerFiscalEntitiesResponse{value: val, isSet: true}
}

func (v NullableCustomerFiscalEntitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFiscalEntitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

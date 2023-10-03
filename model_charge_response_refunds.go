package conekta

import (
	"encoding/json"
)

// checks if the ChargeResponseRefunds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeResponseRefunds{}

// ChargeResponseRefunds struct for ChargeResponseRefunds
type ChargeResponseRefunds struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string `json:"object"`
	// URL of the next page.
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Url of the previous page.
	PreviousPageUrl NullableString `json:"previous_page_url,omitempty"`
	// refunds
	Data []ChargeResponseRefundsData `json:"data,omitempty"`
}

// NewChargeResponseRefunds instantiates a new ChargeResponseRefunds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeResponseRefunds(hasMore bool, object string) *ChargeResponseRefunds {
	this := ChargeResponseRefunds{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewChargeResponseRefundsWithDefaults instantiates a new ChargeResponseRefunds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeResponseRefundsWithDefaults() *ChargeResponseRefunds {
	this := ChargeResponseRefunds{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *ChargeResponseRefunds) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefunds) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ChargeResponseRefunds) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *ChargeResponseRefunds) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefunds) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ChargeResponseRefunds) SetObject(v string) {
	o.Object = v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeResponseRefunds) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeResponseRefunds) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *ChargeResponseRefunds) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *ChargeResponseRefunds) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *ChargeResponseRefunds) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *ChargeResponseRefunds) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPreviousPageUrl returns the PreviousPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeResponseRefunds) GetPreviousPageUrl() string {
	if o == nil || IsNil(o.PreviousPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPageUrl.Get()
}

// GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeResponseRefunds) GetPreviousPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPageUrl.Get(), o.PreviousPageUrl.IsSet()
}

// HasPreviousPageUrl returns a boolean if a field has been set.
func (o *ChargeResponseRefunds) HasPreviousPageUrl() bool {
	if o != nil && o.PreviousPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousPageUrl gets a reference to the given NullableString and assigns it to the PreviousPageUrl field.
func (o *ChargeResponseRefunds) SetPreviousPageUrl(v string) {
	o.PreviousPageUrl.Set(&v)
}

// SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil
func (o *ChargeResponseRefunds) SetPreviousPageUrlNil() {
	o.PreviousPageUrl.Set(nil)
}

// UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
func (o *ChargeResponseRefunds) UnsetPreviousPageUrl() {
	o.PreviousPageUrl.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ChargeResponseRefunds) GetData() []ChargeResponseRefundsData {
	if o == nil || IsNil(o.Data) {
		var ret []ChargeResponseRefundsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefunds) GetDataOk() ([]ChargeResponseRefundsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ChargeResponseRefunds) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ChargeResponseRefundsData and assigns it to the Data field.
func (o *ChargeResponseRefunds) SetData(v []ChargeResponseRefundsData) {
	o.Data = v
}

func (o ChargeResponseRefunds) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeResponseRefunds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if o.NextPageUrl.IsSet() {
		toSerialize["next_page_url"] = o.NextPageUrl.Get()
	}
	if o.PreviousPageUrl.IsSet() {
		toSerialize["previous_page_url"] = o.PreviousPageUrl.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableChargeResponseRefunds struct {
	value *ChargeResponseRefunds
	isSet bool
}

func (v NullableChargeResponseRefunds) Get() *ChargeResponseRefunds {
	return v.value
}

func (v *NullableChargeResponseRefunds) Set(val *ChargeResponseRefunds) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeResponseRefunds) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeResponseRefunds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeResponseRefunds(val *ChargeResponseRefunds) *NullableChargeResponseRefunds {
	return &NullableChargeResponseRefunds{value: val, isSet: true}
}

func (v NullableChargeResponseRefunds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeResponseRefunds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

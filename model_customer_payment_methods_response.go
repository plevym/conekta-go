package conekta

import (
	"encoding/json"
)

// checks if the CustomerPaymentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPaymentMethodsResponse{}

// CustomerPaymentMethodsResponse struct for CustomerPaymentMethodsResponse
type CustomerPaymentMethodsResponse struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string `json:"object"`
	// URL of the next page.
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Url of the previous page.
	PreviousPageUrl NullableString               `json:"previous_page_url,omitempty"`
	Data            []CustomerPaymentMethodsData `json:"data,omitempty"`
}

// NewCustomerPaymentMethodsResponse instantiates a new CustomerPaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentMethodsResponse(hasMore bool, object string) *CustomerPaymentMethodsResponse {
	this := CustomerPaymentMethodsResponse{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewCustomerPaymentMethodsResponseWithDefaults instantiates a new CustomerPaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentMethodsResponseWithDefaults() *CustomerPaymentMethodsResponse {
	this := CustomerPaymentMethodsResponse{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *CustomerPaymentMethodsResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentMethodsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CustomerPaymentMethodsResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *CustomerPaymentMethodsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentMethodsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerPaymentMethodsResponse) SetObject(v string) {
	o.Object = v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerPaymentMethodsResponse) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPaymentMethodsResponse) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *CustomerPaymentMethodsResponse) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *CustomerPaymentMethodsResponse) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *CustomerPaymentMethodsResponse) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *CustomerPaymentMethodsResponse) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPreviousPageUrl returns the PreviousPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerPaymentMethodsResponse) GetPreviousPageUrl() string {
	if o == nil || IsNil(o.PreviousPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPageUrl.Get()
}

// GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPaymentMethodsResponse) GetPreviousPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPageUrl.Get(), o.PreviousPageUrl.IsSet()
}

// HasPreviousPageUrl returns a boolean if a field has been set.
func (o *CustomerPaymentMethodsResponse) HasPreviousPageUrl() bool {
	if o != nil && o.PreviousPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousPageUrl gets a reference to the given NullableString and assigns it to the PreviousPageUrl field.
func (o *CustomerPaymentMethodsResponse) SetPreviousPageUrl(v string) {
	o.PreviousPageUrl.Set(&v)
}

// SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil
func (o *CustomerPaymentMethodsResponse) SetPreviousPageUrlNil() {
	o.PreviousPageUrl.Set(nil)
}

// UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
func (o *CustomerPaymentMethodsResponse) UnsetPreviousPageUrl() {
	o.PreviousPageUrl.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerPaymentMethodsResponse) GetData() []CustomerPaymentMethodsData {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerPaymentMethodsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentMethodsResponse) GetDataOk() ([]CustomerPaymentMethodsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerPaymentMethodsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerPaymentMethodsData and assigns it to the Data field.
func (o *CustomerPaymentMethodsResponse) SetData(v []CustomerPaymentMethodsData) {
	o.Data = v
}

func (o CustomerPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPaymentMethodsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableCustomerPaymentMethodsResponse struct {
	value *CustomerPaymentMethodsResponse
	isSet bool
}

func (v NullableCustomerPaymentMethodsResponse) Get() *CustomerPaymentMethodsResponse {
	return v.value
}

func (v *NullableCustomerPaymentMethodsResponse) Set(val *CustomerPaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentMethodsResponse(val *CustomerPaymentMethodsResponse) *NullableCustomerPaymentMethodsResponse {
	return &NullableCustomerPaymentMethodsResponse{value: val, isSet: true}
}

func (v NullableCustomerPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

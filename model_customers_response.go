package conekta

import (
	"encoding/json"
)

// checks if the CustomersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomersResponse{}

// CustomersResponse struct for CustomersResponse
type CustomersResponse struct {
	Data []CustomerResponse `json:"data,omitempty"`
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string `json:"object"`
	// URL of the next page.
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Url of the previous page.
	PreviousPageUrl NullableString `json:"previous_page_url,omitempty"`
}

// NewCustomersResponse instantiates a new CustomersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomersResponse(hasMore bool, object string) *CustomersResponse {
	this := CustomersResponse{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewCustomersResponseWithDefaults instantiates a new CustomersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomersResponseWithDefaults() *CustomersResponse {
	this := CustomersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomersResponse) GetData() []CustomerResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomersResponse) GetDataOk() ([]CustomerResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomersResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerResponse and assigns it to the Data field.
func (o *CustomersResponse) SetData(v []CustomerResponse) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *CustomersResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CustomersResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CustomersResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *CustomersResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomersResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomersResponse) SetObject(v string) {
	o.Object = v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomersResponse) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomersResponse) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *CustomersResponse) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *CustomersResponse) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *CustomersResponse) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *CustomersResponse) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPreviousPageUrl returns the PreviousPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomersResponse) GetPreviousPageUrl() string {
	if o == nil || IsNil(o.PreviousPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPageUrl.Get()
}

// GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomersResponse) GetPreviousPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPageUrl.Get(), o.PreviousPageUrl.IsSet()
}

// HasPreviousPageUrl returns a boolean if a field has been set.
func (o *CustomersResponse) HasPreviousPageUrl() bool {
	if o != nil && o.PreviousPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousPageUrl gets a reference to the given NullableString and assigns it to the PreviousPageUrl field.
func (o *CustomersResponse) SetPreviousPageUrl(v string) {
	o.PreviousPageUrl.Set(&v)
}

// SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil
func (o *CustomersResponse) SetPreviousPageUrlNil() {
	o.PreviousPageUrl.Set(nil)
}

// UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
func (o *CustomersResponse) UnsetPreviousPageUrl() {
	o.PreviousPageUrl.Unset()
}

func (o CustomersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if o.NextPageUrl.IsSet() {
		toSerialize["next_page_url"] = o.NextPageUrl.Get()
	}
	if o.PreviousPageUrl.IsSet() {
		toSerialize["previous_page_url"] = o.PreviousPageUrl.Get()
	}
	return toSerialize, nil
}

type NullableCustomersResponse struct {
	value *CustomersResponse
	isSet bool
}

func (v NullableCustomersResponse) Get() *CustomersResponse {
	return v.value
}

func (v *NullableCustomersResponse) Set(val *CustomersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomersResponse(val *CustomersResponse) *NullableCustomersResponse {
	return &NullableCustomersResponse{value: val, isSet: true}
}

func (v NullableCustomersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

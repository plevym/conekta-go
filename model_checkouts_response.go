package conekta

import (
	"encoding/json"
)

// checks if the CheckoutsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutsResponse{}

// CheckoutsResponse struct for CheckoutsResponse
type CheckoutsResponse struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string `json:"object"`
	// URL of the next page.
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Url of the previous page.
	PreviousPageUrl NullableString     `json:"previous_page_url,omitempty"`
	Data            []CheckoutResponse `json:"data,omitempty"`
}

// NewCheckoutsResponse instantiates a new CheckoutsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutsResponse(hasMore bool, object string) *CheckoutsResponse {
	this := CheckoutsResponse{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewCheckoutsResponseWithDefaults instantiates a new CheckoutsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutsResponseWithDefaults() *CheckoutsResponse {
	this := CheckoutsResponse{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *CheckoutsResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CheckoutsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CheckoutsResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *CheckoutsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CheckoutsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CheckoutsResponse) SetObject(v string) {
	o.Object = v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutsResponse) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutsResponse) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *CheckoutsResponse) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *CheckoutsResponse) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *CheckoutsResponse) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *CheckoutsResponse) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPreviousPageUrl returns the PreviousPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutsResponse) GetPreviousPageUrl() string {
	if o == nil || IsNil(o.PreviousPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPageUrl.Get()
}

// GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutsResponse) GetPreviousPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPageUrl.Get(), o.PreviousPageUrl.IsSet()
}

// HasPreviousPageUrl returns a boolean if a field has been set.
func (o *CheckoutsResponse) HasPreviousPageUrl() bool {
	if o != nil && o.PreviousPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousPageUrl gets a reference to the given NullableString and assigns it to the PreviousPageUrl field.
func (o *CheckoutsResponse) SetPreviousPageUrl(v string) {
	o.PreviousPageUrl.Set(&v)
}

// SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil
func (o *CheckoutsResponse) SetPreviousPageUrlNil() {
	o.PreviousPageUrl.Set(nil)
}

// UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
func (o *CheckoutsResponse) UnsetPreviousPageUrl() {
	o.PreviousPageUrl.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckoutsResponse) GetData() []CheckoutResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CheckoutResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutsResponse) GetDataOk() ([]CheckoutResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckoutsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CheckoutResponse and assigns it to the Data field.
func (o *CheckoutsResponse) SetData(v []CheckoutResponse) {
	o.Data = v
}

func (o CheckoutsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableCheckoutsResponse struct {
	value *CheckoutsResponse
	isSet bool
}

func (v NullableCheckoutsResponse) Get() *CheckoutsResponse {
	return v.value
}

func (v *NullableCheckoutsResponse) Set(val *CheckoutsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutsResponse(val *CheckoutsResponse) *NullableCheckoutsResponse {
	return &NullableCheckoutsResponse{value: val, isSet: true}
}

func (v NullableCheckoutsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the GetWebhookKeysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhookKeysResponse{}

// GetWebhookKeysResponse struct for GetWebhookKeysResponse
type GetWebhookKeysResponse struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string `json:"object"`
	// URL of the next page.
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Url of the previous page.
	PreviousPageUrl NullableString       `json:"previous_page_url,omitempty"`
	Data            []WebhookKeyResponse `json:"data,omitempty"`
}

// NewGetWebhookKeysResponse instantiates a new GetWebhookKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookKeysResponse(hasMore bool, object string) *GetWebhookKeysResponse {
	this := GetWebhookKeysResponse{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewGetWebhookKeysResponseWithDefaults instantiates a new GetWebhookKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookKeysResponseWithDefaults() *GetWebhookKeysResponse {
	this := GetWebhookKeysResponse{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *GetWebhookKeysResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *GetWebhookKeysResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *GetWebhookKeysResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *GetWebhookKeysResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GetWebhookKeysResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GetWebhookKeysResponse) SetObject(v string) {
	o.Object = v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetWebhookKeysResponse) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetWebhookKeysResponse) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *GetWebhookKeysResponse) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *GetWebhookKeysResponse) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *GetWebhookKeysResponse) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *GetWebhookKeysResponse) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPreviousPageUrl returns the PreviousPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetWebhookKeysResponse) GetPreviousPageUrl() string {
	if o == nil || IsNil(o.PreviousPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPageUrl.Get()
}

// GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetWebhookKeysResponse) GetPreviousPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPageUrl.Get(), o.PreviousPageUrl.IsSet()
}

// HasPreviousPageUrl returns a boolean if a field has been set.
func (o *GetWebhookKeysResponse) HasPreviousPageUrl() bool {
	if o != nil && o.PreviousPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousPageUrl gets a reference to the given NullableString and assigns it to the PreviousPageUrl field.
func (o *GetWebhookKeysResponse) SetPreviousPageUrl(v string) {
	o.PreviousPageUrl.Set(&v)
}

// SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil
func (o *GetWebhookKeysResponse) SetPreviousPageUrlNil() {
	o.PreviousPageUrl.Set(nil)
}

// UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
func (o *GetWebhookKeysResponse) UnsetPreviousPageUrl() {
	o.PreviousPageUrl.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWebhookKeysResponse) GetData() []WebhookKeyResponse {
	if o == nil || IsNil(o.Data) {
		var ret []WebhookKeyResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhookKeysResponse) GetDataOk() ([]WebhookKeyResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWebhookKeysResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []WebhookKeyResponse and assigns it to the Data field.
func (o *GetWebhookKeysResponse) SetData(v []WebhookKeyResponse) {
	o.Data = v
}

func (o GetWebhookKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhookKeysResponse) ToMap() (map[string]interface{}, error) {
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

type NullableGetWebhookKeysResponse struct {
	value *GetWebhookKeysResponse
	isSet bool
}

func (v NullableGetWebhookKeysResponse) Get() *GetWebhookKeysResponse {
	return v.value
}

func (v *NullableGetWebhookKeysResponse) Set(val *GetWebhookKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookKeysResponse(val *GetWebhookKeysResponse) *NullableGetWebhookKeysResponse {
	return &NullableGetWebhookKeysResponse{value: val, isSet: true}
}

func (v NullableGetWebhookKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

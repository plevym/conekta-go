package conekta

import (
	"encoding/json"
)

// checks if the ApiKeyCreateResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyCreateResponseAllOf{}

// ApiKeyCreateResponseAllOf struct for ApiKeyCreateResponseAllOf
type ApiKeyCreateResponseAllOf struct {
	// It is occupied as a user when authenticated with basic authentication, with a blank password. This value will only appear once, in the request to create a new key. Copy and save it in a safe place.
	AuthenticationToken *string `json:"authentication_token,omitempty"`
}

// NewApiKeyCreateResponseAllOf instantiates a new ApiKeyCreateResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyCreateResponseAllOf() *ApiKeyCreateResponseAllOf {
	this := ApiKeyCreateResponseAllOf{}
	return &this
}

// NewApiKeyCreateResponseAllOfWithDefaults instantiates a new ApiKeyCreateResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyCreateResponseAllOfWithDefaults() *ApiKeyCreateResponseAllOf {
	this := ApiKeyCreateResponseAllOf{}
	return &this
}

// GetAuthenticationToken returns the AuthenticationToken field value if set, zero value otherwise.
func (o *ApiKeyCreateResponseAllOf) GetAuthenticationToken() string {
	if o == nil || IsNil(o.AuthenticationToken) {
		var ret string
		return ret
	}
	return *o.AuthenticationToken
}

// GetAuthenticationTokenOk returns a tuple with the AuthenticationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponseAllOf) GetAuthenticationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationToken) {
		return nil, false
	}
	return o.AuthenticationToken, true
}

// HasAuthenticationToken returns a boolean if a field has been set.
func (o *ApiKeyCreateResponseAllOf) HasAuthenticationToken() bool {
	if o != nil && !IsNil(o.AuthenticationToken) {
		return true
	}

	return false
}

// SetAuthenticationToken gets a reference to the given string and assigns it to the AuthenticationToken field.
func (o *ApiKeyCreateResponseAllOf) SetAuthenticationToken(v string) {
	o.AuthenticationToken = &v
}

func (o ApiKeyCreateResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyCreateResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationToken) {
		toSerialize["authentication_token"] = o.AuthenticationToken
	}
	return toSerialize, nil
}

type NullableApiKeyCreateResponseAllOf struct {
	value *ApiKeyCreateResponseAllOf
	isSet bool
}

func (v NullableApiKeyCreateResponseAllOf) Get() *ApiKeyCreateResponseAllOf {
	return v.value
}

func (v *NullableApiKeyCreateResponseAllOf) Set(val *ApiKeyCreateResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyCreateResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyCreateResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyCreateResponseAllOf(val *ApiKeyCreateResponseAllOf) *NullableApiKeyCreateResponseAllOf {
	return &NullableApiKeyCreateResponseAllOf{value: val, isSet: true}
}

func (v NullableApiKeyCreateResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyCreateResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

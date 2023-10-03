package conekta

import (
	"encoding/json"
)

// checks if the SmsCheckoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsCheckoutRequest{}

// SmsCheckoutRequest struct for SmsCheckoutRequest
type SmsCheckoutRequest struct {
	Phonenumber string `json:"phonenumber"`
}

// NewSmsCheckoutRequest instantiates a new SmsCheckoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsCheckoutRequest(phonenumber string) *SmsCheckoutRequest {
	this := SmsCheckoutRequest{}
	this.Phonenumber = phonenumber
	return &this
}

// NewSmsCheckoutRequestWithDefaults instantiates a new SmsCheckoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsCheckoutRequestWithDefaults() *SmsCheckoutRequest {
	this := SmsCheckoutRequest{}
	return &this
}

// GetPhonenumber returns the Phonenumber field value
func (o *SmsCheckoutRequest) GetPhonenumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phonenumber
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value
// and a boolean to check if the value has been set.
func (o *SmsCheckoutRequest) GetPhonenumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phonenumber, true
}

// SetPhonenumber sets field value
func (o *SmsCheckoutRequest) SetPhonenumber(v string) {
	o.Phonenumber = v
}

func (o SmsCheckoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsCheckoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phonenumber"] = o.Phonenumber
	return toSerialize, nil
}

type NullableSmsCheckoutRequest struct {
	value *SmsCheckoutRequest
	isSet bool
}

func (v NullableSmsCheckoutRequest) Get() *SmsCheckoutRequest {
	return v.value
}

func (v *NullableSmsCheckoutRequest) Set(val *SmsCheckoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsCheckoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsCheckoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsCheckoutRequest(val *SmsCheckoutRequest) *NullableSmsCheckoutRequest {
	return &NullableSmsCheckoutRequest{value: val, isSet: true}
}

func (v NullableSmsCheckoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsCheckoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

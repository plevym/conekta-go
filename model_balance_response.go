package conekta

import (
	"encoding/json"
)

// checks if the BalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceResponse{}

// BalanceResponse balance model
type BalanceResponse struct {
	// The balance's available
	Available []BalanceCommonField `json:"available,omitempty"`
	// The balance's cashout retention amount
	CashoutRetentionAmount []BalanceCommonField `json:"cashout_retention_amount,omitempty"`
	// The balance's conekta retention
	ConektaRetention []BalanceCommonField `json:"conekta_retention,omitempty"`
	// The balance's gateway
	Gateway []BalanceCommonField `json:"gateway,omitempty"`
	// The balance's pending
	Pending []BalanceCommonField `json:"pending,omitempty"`
	// The balance's retained
	Retained []BalanceCommonField `json:"retained,omitempty"`
	// The balance's retention amount
	RetentionAmount []BalanceCommonField `json:"retention_amount,omitempty"`
	// The balance's target collateral amount
	TargetCollateralAmount map[string]interface{} `json:"target_collateral_amount,omitempty"`
	// The balance's target retention amount
	TargetRetentionAmount []BalanceCommonField `json:"target_retention_amount,omitempty"`
	// The balance's temporarily retained
	TemporarilyRetained []BalanceCommonField `json:"temporarily_retained,omitempty"`
}

// NewBalanceResponse instantiates a new BalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceResponse() *BalanceResponse {
	this := BalanceResponse{}
	return &this
}

// NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceResponseWithDefaults() *BalanceResponse {
	this := BalanceResponse{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *BalanceResponse) GetAvailable() []BalanceCommonField {
	if o == nil || IsNil(o.Available) {
		var ret []BalanceCommonField
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetAvailableOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *BalanceResponse) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given []BalanceCommonField and assigns it to the Available field.
func (o *BalanceResponse) SetAvailable(v []BalanceCommonField) {
	o.Available = v
}

// GetCashoutRetentionAmount returns the CashoutRetentionAmount field value if set, zero value otherwise.
func (o *BalanceResponse) GetCashoutRetentionAmount() []BalanceCommonField {
	if o == nil || IsNil(o.CashoutRetentionAmount) {
		var ret []BalanceCommonField
		return ret
	}
	return o.CashoutRetentionAmount
}

// GetCashoutRetentionAmountOk returns a tuple with the CashoutRetentionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetCashoutRetentionAmountOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.CashoutRetentionAmount) {
		return nil, false
	}
	return o.CashoutRetentionAmount, true
}

// HasCashoutRetentionAmount returns a boolean if a field has been set.
func (o *BalanceResponse) HasCashoutRetentionAmount() bool {
	if o != nil && !IsNil(o.CashoutRetentionAmount) {
		return true
	}

	return false
}

// SetCashoutRetentionAmount gets a reference to the given []BalanceCommonField and assigns it to the CashoutRetentionAmount field.
func (o *BalanceResponse) SetCashoutRetentionAmount(v []BalanceCommonField) {
	o.CashoutRetentionAmount = v
}

// GetConektaRetention returns the ConektaRetention field value if set, zero value otherwise.
func (o *BalanceResponse) GetConektaRetention() []BalanceCommonField {
	if o == nil || IsNil(o.ConektaRetention) {
		var ret []BalanceCommonField
		return ret
	}
	return o.ConektaRetention
}

// GetConektaRetentionOk returns a tuple with the ConektaRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetConektaRetentionOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.ConektaRetention) {
		return nil, false
	}
	return o.ConektaRetention, true
}

// HasConektaRetention returns a boolean if a field has been set.
func (o *BalanceResponse) HasConektaRetention() bool {
	if o != nil && !IsNil(o.ConektaRetention) {
		return true
	}

	return false
}

// SetConektaRetention gets a reference to the given []BalanceCommonField and assigns it to the ConektaRetention field.
func (o *BalanceResponse) SetConektaRetention(v []BalanceCommonField) {
	o.ConektaRetention = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *BalanceResponse) GetGateway() []BalanceCommonField {
	if o == nil || IsNil(o.Gateway) {
		var ret []BalanceCommonField
		return ret
	}
	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetGatewayOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *BalanceResponse) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given []BalanceCommonField and assigns it to the Gateway field.
func (o *BalanceResponse) SetGateway(v []BalanceCommonField) {
	o.Gateway = v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *BalanceResponse) GetPending() []BalanceCommonField {
	if o == nil || IsNil(o.Pending) {
		var ret []BalanceCommonField
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetPendingOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *BalanceResponse) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given []BalanceCommonField and assigns it to the Pending field.
func (o *BalanceResponse) SetPending(v []BalanceCommonField) {
	o.Pending = v
}

// GetRetained returns the Retained field value if set, zero value otherwise.
func (o *BalanceResponse) GetRetained() []BalanceCommonField {
	if o == nil || IsNil(o.Retained) {
		var ret []BalanceCommonField
		return ret
	}
	return o.Retained
}

// GetRetainedOk returns a tuple with the Retained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetRetainedOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.Retained) {
		return nil, false
	}
	return o.Retained, true
}

// HasRetained returns a boolean if a field has been set.
func (o *BalanceResponse) HasRetained() bool {
	if o != nil && !IsNil(o.Retained) {
		return true
	}

	return false
}

// SetRetained gets a reference to the given []BalanceCommonField and assigns it to the Retained field.
func (o *BalanceResponse) SetRetained(v []BalanceCommonField) {
	o.Retained = v
}

// GetRetentionAmount returns the RetentionAmount field value if set, zero value otherwise.
func (o *BalanceResponse) GetRetentionAmount() []BalanceCommonField {
	if o == nil || IsNil(o.RetentionAmount) {
		var ret []BalanceCommonField
		return ret
	}
	return o.RetentionAmount
}

// GetRetentionAmountOk returns a tuple with the RetentionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetRetentionAmountOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.RetentionAmount) {
		return nil, false
	}
	return o.RetentionAmount, true
}

// HasRetentionAmount returns a boolean if a field has been set.
func (o *BalanceResponse) HasRetentionAmount() bool {
	if o != nil && !IsNil(o.RetentionAmount) {
		return true
	}

	return false
}

// SetRetentionAmount gets a reference to the given []BalanceCommonField and assigns it to the RetentionAmount field.
func (o *BalanceResponse) SetRetentionAmount(v []BalanceCommonField) {
	o.RetentionAmount = v
}

// GetTargetCollateralAmount returns the TargetCollateralAmount field value if set, zero value otherwise.
func (o *BalanceResponse) GetTargetCollateralAmount() map[string]interface{} {
	if o == nil || IsNil(o.TargetCollateralAmount) {
		var ret map[string]interface{}
		return ret
	}
	return o.TargetCollateralAmount
}

// GetTargetCollateralAmountOk returns a tuple with the TargetCollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetTargetCollateralAmountOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TargetCollateralAmount) {
		return map[string]interface{}{}, false
	}
	return o.TargetCollateralAmount, true
}

// HasTargetCollateralAmount returns a boolean if a field has been set.
func (o *BalanceResponse) HasTargetCollateralAmount() bool {
	if o != nil && !IsNil(o.TargetCollateralAmount) {
		return true
	}

	return false
}

// SetTargetCollateralAmount gets a reference to the given map[string]interface{} and assigns it to the TargetCollateralAmount field.
func (o *BalanceResponse) SetTargetCollateralAmount(v map[string]interface{}) {
	o.TargetCollateralAmount = v
}

// GetTargetRetentionAmount returns the TargetRetentionAmount field value if set, zero value otherwise.
func (o *BalanceResponse) GetTargetRetentionAmount() []BalanceCommonField {
	if o == nil || IsNil(o.TargetRetentionAmount) {
		var ret []BalanceCommonField
		return ret
	}
	return o.TargetRetentionAmount
}

// GetTargetRetentionAmountOk returns a tuple with the TargetRetentionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetTargetRetentionAmountOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.TargetRetentionAmount) {
		return nil, false
	}
	return o.TargetRetentionAmount, true
}

// HasTargetRetentionAmount returns a boolean if a field has been set.
func (o *BalanceResponse) HasTargetRetentionAmount() bool {
	if o != nil && !IsNil(o.TargetRetentionAmount) {
		return true
	}

	return false
}

// SetTargetRetentionAmount gets a reference to the given []BalanceCommonField and assigns it to the TargetRetentionAmount field.
func (o *BalanceResponse) SetTargetRetentionAmount(v []BalanceCommonField) {
	o.TargetRetentionAmount = v
}

// GetTemporarilyRetained returns the TemporarilyRetained field value if set, zero value otherwise.
func (o *BalanceResponse) GetTemporarilyRetained() []BalanceCommonField {
	if o == nil || IsNil(o.TemporarilyRetained) {
		var ret []BalanceCommonField
		return ret
	}
	return o.TemporarilyRetained
}

// GetTemporarilyRetainedOk returns a tuple with the TemporarilyRetained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetTemporarilyRetainedOk() ([]BalanceCommonField, bool) {
	if o == nil || IsNil(o.TemporarilyRetained) {
		return nil, false
	}
	return o.TemporarilyRetained, true
}

// HasTemporarilyRetained returns a boolean if a field has been set.
func (o *BalanceResponse) HasTemporarilyRetained() bool {
	if o != nil && !IsNil(o.TemporarilyRetained) {
		return true
	}

	return false
}

// SetTemporarilyRetained gets a reference to the given []BalanceCommonField and assigns it to the TemporarilyRetained field.
func (o *BalanceResponse) SetTemporarilyRetained(v []BalanceCommonField) {
	o.TemporarilyRetained = v
}

func (o BalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.CashoutRetentionAmount) {
		toSerialize["cashout_retention_amount"] = o.CashoutRetentionAmount
	}
	if !IsNil(o.ConektaRetention) {
		toSerialize["conekta_retention"] = o.ConektaRetention
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !IsNil(o.Retained) {
		toSerialize["retained"] = o.Retained
	}
	if !IsNil(o.RetentionAmount) {
		toSerialize["retention_amount"] = o.RetentionAmount
	}
	if !IsNil(o.TargetCollateralAmount) {
		toSerialize["target_collateral_amount"] = o.TargetCollateralAmount
	}
	if !IsNil(o.TargetRetentionAmount) {
		toSerialize["target_retention_amount"] = o.TargetRetentionAmount
	}
	if !IsNil(o.TemporarilyRetained) {
		toSerialize["temporarily_retained"] = o.TemporarilyRetained
	}
	return toSerialize, nil
}

type NullableBalanceResponse struct {
	value *BalanceResponse
	isSet bool
}

func (v NullableBalanceResponse) Get() *BalanceResponse {
	return v.value
}

func (v *NullableBalanceResponse) Set(val *BalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceResponse(val *BalanceResponse) *NullableBalanceResponse {
	return &NullableBalanceResponse{value: val, isSet: true}
}

func (v NullableBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

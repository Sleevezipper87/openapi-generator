/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore
import (
	"encoding/json"
)
// NumberOnly struct for NumberOnly
type NumberOnly struct {
	JustNumber *float32 `json:"JustNumber,omitempty"`

}

// GetJustNumber returns the JustNumber field if non-nil, zero value otherwise.
func (o *NumberOnly) GetJustNumber() float32 {
	if o == nil || o.JustNumber == nil {
		var ret float32
		return ret
	}
	return *o.JustNumber
}

// GetJustNumberOk returns a tuple with the JustNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NumberOnly) GetJustNumberOk() (float32, bool) {
	if o == nil || o.JustNumber == nil {
		var ret float32
		return ret, false
	}
	return *o.JustNumber, true
}

// HasJustNumber returns a boolean if a field has been set.
func (o *NumberOnly) HasJustNumber() bool {
	if o != nil && o.JustNumber != nil {
		return true
	}

	return false
}

// SetJustNumber gets a reference to the given float32 and assigns it to the JustNumber field.
func (o *NumberOnly) SetJustNumber(v float32) {
	o.JustNumber = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o NumberOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JustNumber != nil {
		toSerialize["JustNumber"] = o.JustNumber
	}
	return json.Marshal(toSerialize)
}


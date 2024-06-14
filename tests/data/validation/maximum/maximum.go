// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type Maximum struct {
	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger int `json:"myInteger" yaml:"myInteger" mapstructure:"myInteger"`

	// MyNullableInteger corresponds to the JSON schema field "myNullableInteger".
	MyNullableInteger *int `json:"myNullableInteger,omitempty" yaml:"myNullableInteger,omitempty" mapstructure:"myNullableInteger,omitempty"`

	// MyNullableNumber corresponds to the JSON schema field "myNullableNumber".
	MyNullableNumber *float64 `json:"myNullableNumber,omitempty" yaml:"myNullableNumber,omitempty" mapstructure:"myNullableNumber,omitempty"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber float64 `json:"myNumber" yaml:"myNumber" mapstructure:"myNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Maximum) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["myInteger"]; raw != nil && !ok {
		return fmt.Errorf("field myInteger in Maximum: required")
	}
	if _, ok := raw["myNumber"]; raw != nil && !ok {
		return fmt.Errorf("field myNumber in Maximum: required")
	}
	type Plain Maximum
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.MyInteger > 5 {
		return fmt.Errorf("field %s: violates maximum %s", "myInteger", "5")
	}
	if plain.MyNullableInteger != nil && *plain.MyNullableInteger > 10 {
		return fmt.Errorf("field %s: violates maximum %s", "myNullableInteger", "10")
	}
	if plain.MyNullableNumber != nil && *plain.MyNullableNumber > 10.3 {
		return fmt.Errorf("field %s: violates maximum %s", "myNullableNumber", "10.3")
	}
	if plain.MyNumber > 5.6 {
		return fmt.Errorf("field %s: violates maximum %s", "myNumber", "5.6")
	}
	*j = Maximum(plain)
	return nil
}

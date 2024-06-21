package validators

import (
	"fmt"
	"testing"
)

type FieldValidations struct {
	IsString       string   `json:"isString"       validates:"required;isString;minLength=2;maxLength=5"`
	IsNullString1  *string  `json:"isNullString1"  validates:"isNullString;minLength=2;maxLength=5"`
	IsNullString2  *string  `json:"isNullString2"  validates:"isNullString;minLength=2;maxLength=5"`
	Email          string   `json:"email"          validates:"required;email"`
	IsNumber1      float64  `json:"isNumber1"      validates:"required;isNumber;min=2;max=5"`
	IsNumber2      float64  `json:"isNumber2"      validates:"required;isNumber;min=2;max=5"`
	IsNullNumber1  *float64 `json:"isNullNumber1"  validates:"isNullNumber;min=2;max=5"`
	IsNullNumber2  *float64 `json:"isNullNumber2"  validates:"isNullNumber;min=2;max=5"`
	IsNullNumber3  *float64 `json:"isNullNumber3"  validates:"isNullNumber;min=2;max=5"`
	IsInt          int      `json:"isInt"          validates:"required;isInt;min=2;max=5"`
	IsNullInt1     *int     `json:"isNullInt1"     validates:"isNullInt;min=2;max=5"`
	IsNullInt2     *int     `json:"isNullInt2"     validates:"isNullInt;min=2;max=5"`
	IsFloat        float64  `json:"isFloat"        validates:"required;isFloat;min=2;max=5"`
	IsNullFloat1   *float64 `json:"isNullFloat1"   validates:"isNullFloat;min=2;max=5"`
	IsNullFloat2   *float64 `json:"isNullFloat2"   validates:"isNullFloat;min=2;max=5"`
	IsBoolean      bool     `json:"isBoolean"      validates:"required;isBool"`
	IsNullBoolean1 *bool    `json:"isNullBoolean1" validates:"isNullBool"`
	IsNullBoolean2 *bool    `json:"isNullBoolean2" validates:"isNullBool"`
	StartsWith     string   `json:"startsWith"     validates:"required;startsWith=na"`
	StartsNotWith  string   `json:"startsNotWith"  validates:"required;startsNotWith=ne"`
	EndsWith       string   `json:"endsWith"       validates:"required;endsWith=me"`
	EndsNotWith    string   `json:"endsNotWith"    validates:"required;endsNotWith=ma"`
}

const values = `{
	"isString":       "name",
	"isNullString2":  "name",
	"email":          "emailexample@gmail.com",
	"isNumber1":      3,
	"isNumber2":      3.3,
	"isNullNumber2":  3,
	"isNullNumber3":  3.3,
	"isInt":          3,
	"isNullInt2":     3,
	"isFloat":        3.3,
	"isNullFloat2":   3.3,
	"isBoolean":      true,
	"isNullBoolean2": true,
	"startsWith":     "name",
	"startsNotWith":  "name",
	"endsWith":       "name",
	"endsNotWith":    "name"
}`

func TestValidate(t *testing.T) {
	/// - Successes
	// Test 1
	fieldValidations := FieldValidations{}

	err := Validate(&fieldValidations, values)

	if err != nil {
		t.Errorf(
			fmt.Sprintf("received: %s; expected: nil", err.Error()),
		)
	}
}

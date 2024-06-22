package validators

import (
	"fmt"
	"strings"
	"testing"
)

type FieldValidations struct {
	Email          string   `json:"email"          validates:"required;email"`
	EndsWith       string   `json:"endsWith"       validates:"required;isString;endsWith=me"`
	EndsNotWith    string   `json:"endsNotWith"    validates:"required;isString;endsNotWith=ma"`
	IsAlpha        string   `json:"isAlpha"        validates:"required;isString;isAlpha"`
	IsAlphaSpace1  string   `json:"isAlphaSpace1"  validates:"isString;isAlphaSpace"`
	IsAlphaSpace2  string   `json:"isAlphaSpace2"  validates:"required;isString;isAlphaSpace"`
	IsArray        []string `json:"isArray"        validates:"required;minLength=1;maxLength=2;isArray;isString"`
	IsBoolean      bool     `json:"isBoolean"      validates:"required;isBool"`
	IsFloat        float64  `json:"isFloat"        validates:"required;isFloat;min=2;max=5"`
	IsInt          int      `json:"isInt"          validates:"required;isInt;min=2;max=5"`
	IsNullArray1   []string `json:"isNullArray1"   validates:"minLength=1;maxLength=2;isNullArray;isString"`
	IsNullArray2   []string `json:"isNullArray2"   validates:"minLength=1;maxLength=2;isNullArray;isString"`
	IsNullBoolean1 *bool    `json:"isNullBoolean1" validates:"isNullBool"`
	IsNullBoolean2 *bool    `json:"isNullBoolean2" validates:"isNullBool"`
	IsNullFloat1   *float64 `json:"isNullFloat1"   validates:"isNullFloat;min=2;max=5"`
	IsNullFloat2   *float64 `json:"isNullFloat2"   validates:"isNullFloat;min=2;max=5"`
	IsNullInt1     *int     `json:"isNullInt1"     validates:"isNullInt;min=2;max=5"`
	IsNullInt2     *int     `json:"isNullInt2"     validates:"isNullInt;min=2;max=5"`
	IsNullNumber1  *float64 `json:"isNullNumber1"  validates:"isNullNumber;min=2;max=5"`
	IsNullNumber2  *float64 `json:"isNullNumber2"  validates:"isNullNumber;min=2;max=5"`
	IsNullNumber3  *float64 `json:"isNullNumber3"  validates:"isNullNumber;min=2;max=5"`
	IsNullString1  *string  `json:"isNullString1"  validates:"isNullString;minLength=2;maxLength=5"`
	IsNullString2  *string  `json:"isNullString2"  validates:"isNullString;minLength=2;maxLength=5"`
	IsNumber1      float64  `json:"isNumber1"      validates:"required;isNumber;min=2;max=5"`
	IsNumber2      float64  `json:"isNumber2"      validates:"required;isNumber;min=2;max=5"`
	IsString       string   `json:"isString"       validates:"required;isString;minLength=2;maxLength=5"`
	Password       string   `json:"password"       validates:"required;password"`
	Regex          string   `json:"regex"          validates:"required;isString;regex=^[0-9]+$"`
	StartsWith     string   `json:"startsWith"     validates:"required;isString;startsWith=na"`
	StartsNotWith  string   `json:"startsNotWith"  validates:"required;isString;startsNotWith=ne"`
	Url            string   `json:"url"            validates:"required;url"`
}

const values = `{
	"email":          "emailexample@gmail.com",
	"endsWith":       "name",
	"endsNotWith":    "name",
	"isAlpha":        "abcDEFáÉ",
	"isAlphaSpace1":  "abcDEFáÉ",
	"isAlphaSpace2":  "abcDEFáÉ ",
	"isArray":        ["name"],
	"isBoolean":      true,
	"isFloat":        3.3,
	"isInt":          3,
	"isNullArray2":   ["name"],
	"isNullBoolean2": true,
	"isNullFloat2":   3.3,
	"isNullInt2":     3,
	"isNullNumber2":  3,
	"isNullNumber3":  3.3,
	"isNullString2":  "name",
	"isNumber1":      3,
	"isNumber2":      3.3,
	"isString":       "name",
	"password":       "aA1!",
	"regex":          "0123",
	"startsWith":     "name",
	"startsNotWith":  "name",
	"url":            "https://example.com"
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

	/// - Errors
	/// Required
	// Test 1
	type RequiredStruct1 struct {
		Value string `json:"value" validates:"required"`
	}
	requiredStruct1 := RequiredStruct1{}
	json := `{}`
	err = Validate(&requiredStruct1, json)

	if err == nil {
		t.Errorf("Required - received: nil; expected: \"[error message]\"")
	}

	// Test 2
	type RequiredStruct2 struct {
		Value string `json:"value" validates:"required=error"`
	}
	requiredStruct2 := RequiredStruct2{}
	err = Validate(&requiredStruct2, json)

	if err == nil || !strings.Contains(err.Error(), errCustomMessage) {
		t.Errorf(setValidateError("Required - received: %v; expected: \"error\""))
	}

	/// IsString
	// Test 1
	type IsStringStruct1 struct {
		Value string `json:"value" validates:"isString"`
	}
	isStringStruct1 := IsStringStruct1{}
	json = `{
		"value": 0,
	}`
	err = Validate(&isStringStruct1, json)

	if err == nil {
		t.Errorf("IsString - received: nil; expected: \"[error message]\"")
	}

	// Test 2
	type IsStringStruct2 struct {
		Value string `json:"value" validates:"isString=error"`
	}
	isStringStruct2 := IsStringStruct2{}
	err = Validate(&isStringStruct2, json)

	if err == nil || !strings.Contains(err.Error(), errCustomMessage) {
		t.Errorf(setValidateError("IsString - received: %v; expected: \"error\""))
	}

	/// IsInt
	// Test 1
	type IsIntStruct1 struct {
		Value string `json:"value" validates:"isInt"`
	}
	isIntStruct1 := IsIntStruct1{}
	json = `{
		"value": "name",
	}`
	err = Validate(&isIntStruct1, json)

	if err == nil {
		t.Errorf("IsInt - received: nil; expected: \"[error message]\"")
	}

	// Test 2
	type IsIntStruct2 struct {
		Value string `json:"value" validates:"isInt=error"`
	}
	isIntStruct2 := IsIntStruct2{}
	err = Validate(&isIntStruct2, json)

	if err == nil || !strings.Contains(err.Error(), errCustomMessage) {
		t.Errorf(setValidateError("IsInt - received: %v; expected: \"error\""))
	}
}

package validators

import (
	"fmt"
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
	OneOf          string   `json:"oneOf"          validates:"required;oneOf=EASY MEDIUM HARD"`
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
	"url":            "https://example.com",
	"oneOf":          "EASY"
}`

func TestValidate(t *testing.T) {
	/// - Successes
	// Test 1
	fieldValidations := FieldValidations{}

	err = Validate(&fieldValidations, values)

	if err != nil {
		t.Errorf(
			fmt.Sprintf("received: %s; expected: nil", err.Error()),
		)
	}

	/// - Errors
	// Test 1
	type ValidateStruct struct {
		Value string `json:"value"`
	}
	jsonValue = `{}`
	validateStruct := ValidateStruct{}
	err := Validate(&validateStruct, jsonValue)
	if err == nil {
		t.Errorf("received: nil; expected: \"[error message]\"")
	}

	/// Required
	// Test 1
	type RequiredStruct1 struct {
		Value string `json:"value" validates:"required"`
	}
	initValidateTest(t, validateTestDefault, "Required", &RequiredStruct1{})

	// Test 2
	type RequiredStruct2 struct {
		Value string `json:"value" validates:"required=error"`
	}
	initValidateTest(t, validateTestCustom, "Required", &RequiredStruct2{})

	/// IsString
	// Test 1
	type IsStringStruct1 struct {
		Value string `json:"value" validates:"isString"`
	}
	jsonValue = `{
		"value": 0
	}`
	initValidateTest(t, validateTestDefault, "IsString", &IsStringStruct1{})

	// Test 2
	type IsStringStruct2 struct {
		Value string `json:"value" validates:"isString=error"`
	}
	initValidateTest(t, validateTestCustom, "IsString", &IsStringStruct2{})

	/// IsNumber
	// Test 1
	type IsNumberStruct1 struct {
		Value string `json:"value" validates:"isNumber"`
	}
	jsonValue = `{
		"value": "name"
	}`
	initValidateTest(t, validateTestDefault, "IsNumber", &IsNumberStruct1{})

	// Test 2
	type IsNumberStruct2 struct {
		Value string `json:"value" validates:"isNumber=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNumber", &IsNumberStruct2{})

	/// IsInt
	// Test 1
	type IsIntStruct1 struct {
		Value string `json:"value" validates:"isInt"`
	}
	initValidateTest(t, validateTestDefault, "IsInt", &IsIntStruct1{})

	// Test 2
	type IsIntStruct2 struct {
		Value string `json:"value" validates:"isInt=error"`
	}
	initValidateTest(t, validateTestCustom, "IsInt", &IsIntStruct2{})

	/// IsFloat
	// Test 1
	type IsFloatStruct1 struct {
		Value string `json:"value" validates:"isFloat"`
	}
	initValidateTest(t, validateTestDefault, "IsFloat", &IsFloatStruct1{})

	// Test 2
	type IsFloatStruct2 struct {
		Value string `json:"value" validates:"isFloat=error"`
	}
	initValidateTest(t, validateTestCustom, "IsFloat", &IsFloatStruct2{})

	/// IsBool
	// Test 1
	type IsBoolStruct1 struct {
		Value string `json:"value" validates:"isBool"`
	}
	initValidateTest(t, validateTestDefault, "IsBool", &IsBoolStruct1{})

	// Test 2
	type IsBoolStruct2 struct {
		Value string `json:"value" validates:"isBool=error"`
	}
	initValidateTest(t, validateTestCustom, "IsBool", &IsBoolStruct2{})

	/// IsArray
	// Test 1
	type IsArrayStruct1 struct {
		Value string `json:"value" validates:"isArray"`
	}
	initValidateTest(t, validateTestDefault, "IsArray", &IsArrayStruct1{})

	// Test 2
	type IsArrayStruct2 struct {
		Value string `json:"value" validates:"isArray=error"`
	}
	initValidateTest(t, validateTestCustom, "IsArray", &IsArrayStruct2{})

	/// IsNullString
	// Test 1
	type IsNullStringStruct1 struct {
		Value string `json:"value" validates:"isNullString"`
	}
	jsonValue = `{
		"value": 0
	}`
	initValidateTest(t, validateTestDefault, "IsNullString", &IsNullStringStruct1{})

	// Test 2
	type IsNullStringStruct2 struct {
		Value string `json:"value" validates:"isNullString=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNullString", &IsNullStringStruct2{})

	/// IsNullNumber
	// Test 1
	type IsNullNumberStruct1 struct {
		Value string `json:"value" validates:"isNullNumber"`
	}
	jsonValue = `{
		"value": "name"
	}`
	initValidateTest(t, validateTestDefault, "IsNullNumber", &IsNullNumberStruct1{})

	// Test 2
	type IsNullNumberStruct2 struct {
		Value string `json:"value" validates:"isNullNumber=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNullNumber", &IsNullNumberStruct2{})

	/// IsNullInt
	// Test 1
	type IsNullIntStruct1 struct {
		Value string `json:"value" validates:"isNullInt"`
	}
	initValidateTest(t, validateTestDefault, "IsNullInt", &IsNullIntStruct1{})

	// Test 2
	type IsNullIntStruct2 struct {
		Value string `json:"value" validates:"isNullInt=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNullInt", &IsNullIntStruct2{})

	/// IsNullFloat
	// Test 1
	type IsNullFloatStruct1 struct {
		Value string `json:"value" validates:"isNullFloat"`
	}
	initValidateTest(t, validateTestDefault, "IsNullFloat", &IsNullFloatStruct1{})

	// Test 2
	type IsNullFloatStruct2 struct {
		Value string `json:"value" validates:"isNullFloat=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNullFloat", &IsNullFloatStruct2{})

	/// IsNullBool
	// Test 1
	type IsNullBoolStruct1 struct {
		Value string `json:"value" validates:"isNullBool"`
	}
	initValidateTest(t, validateTestDefault, "IsNullBool", &IsNullBoolStruct1{})

	// Test 2
	type IsNullBoolStruct2 struct {
		Value string `json:"value" validates:"isNullBool=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNullBool", &IsNullBoolStruct2{})

	/// IsNullArray
	// Test 1
	type IsNullArrayStruct1 struct {
		Value string `json:"value" validates:"isNullArray"`
	}
	initValidateTest(t, validateTestDefault, "IsNullArray", &IsNullArrayStruct1{})

	// Test 2
	type IsNullArrayStruct2 struct {
		Value string `json:"value" validates:"isNullArray=error"`
	}
	initValidateTest(t, validateTestCustom, "IsNullArray", &IsNullArrayStruct2{})

	/// Password
	// Test 1
	type PasswordStruct1 struct {
		Value string `json:"value" validates:"password"`
	}
	initValidateTest(t, validateTestDefault, "Password", &PasswordStruct1{})

	// Test 2
	type PasswordStruct2 struct {
		Value string `json:"value" validates:"password=error"`
	}
	initValidateTest(t, validateTestCustom, "Password", &PasswordStruct2{})

	/// Email
	// Test 1
	type EmailStruct1 struct {
		Value string `json:"value" validates:"email"`
	}
	initValidateTest(t, validateTestDefault, "Email", &EmailStruct1{})

	// Test 2
	type EmailStruct2 struct {
		Value string `json:"value" validates:"email=error"`
	}
	initValidateTest(t, validateTestCustom, "Email", &EmailStruct2{})

	/// Min
	// Test 1
	type MinStruct1 struct {
		Value string `json:"value" validates:"min=5"`
	}
	jsonValue = `{
		"value": 3
	}`
	initValidateTest(t, validateTestDefault, "Min", &MinStruct1{})

	// Test 2
	type MinStruct2 struct {
		Value string `json:"value" validates:"min=5,error"`
	}
	initValidateTest(t, validateTestCustom, "Min", &MinStruct2{})

	/// Max
	// Test 1
	type MaxStruct1 struct {
		Value string `json:"value" validates:"max=2"`
	}
	initValidateTest(t, validateTestDefault, "Max", &MaxStruct1{})

	// Test 2
	type MaxStruct2 struct {
		Value string `json:"value" validates:"max=2,error"`
	}
	initValidateTest(t, validateTestCustom, "Max", &MaxStruct2{})

	/// MinLength
	// Test 1
	type MinLengthStruct1 struct {
		Value string `json:"value" validates:"minLength=5"`
	}
	jsonValue = `{
		"value": "name"
	}`
	initValidateTest(t, validateTestDefault, "MinLength", &MinLengthStruct1{})

	// Test 2
	type MinLengthStruct2 struct {
		Value string `json:"value" validates:"minLength=5,error"`
	}
	initValidateTest(t, validateTestCustom, "MinLength", &MinLengthStruct2{})

	/// MaxLength
	// Test 1
	type MaxLengthStruct1 struct {
		Value string `json:"value" validates:"maxLength=2"`
	}
	initValidateTest(t, validateTestDefault, "MaxLength", &MaxLengthStruct1{})

	// Test 2
	type MaxLengthStruct2 struct {
		Value string `json:"value" validates:"maxLength=2,error"`
	}
	initValidateTest(t, validateTestCustom, "MaxLength", &MaxLengthStruct2{})

	/// IsAlpha
	// Test 1
	type IsAlphaStruct1 struct {
		Value string `json:"value" validates:"isAlpha"`
	}
	jsonValue = `{
		"value": "0"
	}`
	initValidateTest(t, validateTestDefault, "IsAlpha", &IsAlphaStruct1{})

	// Test 2
	type IsAlphaStruct2 struct {
		Value string `json:"value" validates:"isAlpha=error"`
	}
	initValidateTest(t, validateTestCustom, "IsAlpha", &IsAlphaStruct2{})

	/// IsAlphaSpace
	// Test 1
	type IsAlphaSpaceStruct1 struct {
		Value string `json:"value" validates:"isAlphaSpace"`
	}
	jsonValue = `{
		"value": "0"
	}`
	initValidateTest(t, validateTestDefault, "IsAlphaSpace", &IsAlphaSpaceStruct1{})

	// Test 2
	type IsAlphaSpaceStruct2 struct {
		Value string `json:"value" validates:"isAlphaSpace=error"`
	}
	initValidateTest(t, validateTestCustom, "IsAlphaSpace", &IsAlphaSpaceStruct2{})

	/// StartsWith
	// Test 1
	type StartsWithStruct1 struct {
		Value string `json:"value" validates:"startsWith=ne"`
	}
	initValidateTest(t, validateTestDefault, "StartsWith", &StartsWithStruct1{})

	// Test 2
	type StartsWithStruct2 struct {
		Value string `json:"value" validates:"startsWith=ne,error"`
	}
	initValidateTest(t, validateTestCustom, "StartsWith", &StartsWithStruct2{})

	/// StartsNotWith
	// Test 1
	type StartsNotWithStruct1 struct {
		Value string `json:"value" validates:"startsNotWith=na"`
	}
	initValidateTest(t, validateTestDefault, "StartsNotWith", &StartsNotWithStruct1{})

	// Test 2
	type StartsNotWithStruct2 struct {
		Value string `json:"value" validates:"startsNotWith=na,error"`
	}
	initValidateTest(t, validateTestCustom, "StartsNotWith", &StartsNotWithStruct2{})

	/// EndsWith
	// Test 1
	type EndsWithStruct1 struct {
		Value string `json:"value" validates:"endsWith=ma"`
	}
	initValidateTest(t, validateTestDefault, "EndsWith", &EndsWithStruct1{})

	// Test 2
	type EndsWithStruct2 struct {
		Value string `json:"value" validates:"endsWith=ma,error"`
	}
	initValidateTest(t, validateTestCustom, "EndsWith", &EndsWithStruct2{})

	/// EndsNotWith
	// Test 1
	type EndsNotWithStruct1 struct {
		Value string `json:"value" validates:"endsNotWith=me"`
	}
	initValidateTest(t, validateTestDefault, "EndsNotWith", &EndsNotWithStruct1{})

	// Test 2
	type EndsNotWithStruct2 struct {
		Value string `json:"value" validates:"endsNotWith=me,error"`
	}
	initValidateTest(t, validateTestCustom, "EndsNotWith", &EndsNotWithStruct2{})

	/// Regex
	// Test 1
	type RegexStruct1 struct {
		Value string `json:"value" validates:"regex=^[0-9]+$"`
	}
	initValidateTest(t, validateTestDefault, "Regex", &RegexStruct1{})

	// Test 2
	type RegexStruct2 struct {
		Value string `json:"value" validates:"regex=^[0-9]+$,error"`
	}
	initValidateTest(t, validateTestCustom, "Regex", &RegexStruct2{})

	/// Url
	// Test 1
	type UrlStruct1 struct {
		Value string `json:"value" validates:"url"`
	}
	initValidateTest(t, validateTestDefault, "Url", &UrlStruct1{})

	// Test 2
	type UrlStruct2 struct {
		Value string `json:"value" validates:"url=error"`
	}
	initValidateTest(t, validateTestCustom, "Url", &UrlStruct2{})

	/// OneOf
	// Test 1
	type OneOfStruct1 struct {
		Value string `json:"value" validates:"oneOf=EASY MEDIUM HARD"`
	}
	initValidateTest(t, validateTestDefault, "OneOf", &OneOfStruct1{})

	/// OneOf
	// Test 1
	type OneOfStruct2 struct {
		Value string `json:"value" validates:"oneOf=EASY MEDIUM HARD, error"`
	}
	initValidateTest(t, validateTestDefault, "OneOf", &OneOfStruct2{})
}

package validators

import "testing"

func TestIsArray(t *testing.T) {
	customErrorMessage := "error"

	/// String Array
	// Errors
	errorMessage, stopLoop = IsArray(
		StringArray,
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	StringArray,
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		StringArray,
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	StringArray,
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		StringArray,
		[]Validator{},
	)([]string{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	StringArray,
	[]Validator{},
)([]string{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		StringArray,
		[]Validator{},
		customErrorMessage,
	)([]string{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	StringArray,
	[]Validator{},
	"error",
)([]string{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}
}

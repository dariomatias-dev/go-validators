package validators

import "testing"

func TestIsArray(t *testing.T) {
	customErrorMessage := "error"

	/// Interface Array
	// Errors
	errorMessage, stopLoop = IsArray(
		InterfaceArray,
		[]Validator{},
	)([]string{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	InterfaceArray,
	[]Validator{},
)([]string{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		InterfaceArray,
		[]Validator{},
		customErrorMessage,
	)([]string{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	InterfaceArray,
	[]Validator{},
	"error",
)([]string{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		InterfaceArray,
		[]Validator{},
	)([]interface{}{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	InterfaceArray,
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		InterfaceArray,
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	InterfaceArray,
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

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

	/// Int Array
	// Errors
	errorMessage, stopLoop = IsArray(
		IntArray,
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	IntArray,
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		IntArray,
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	IntArray,
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		IntArray,
		[]Validator{},
	)([]int{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	IntArray,
	[]Validator{},
)([]int{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		IntArray,
		[]Validator{},
		customErrorMessage,
	)([]int{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	IntArray,
	[]Validator{},
	"error",
)([]int{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}
}

package validators

import (
	"testing"

	arraytype "github.com/dariomatias-dev/go-validators/array_type"
)

func TestIsNullArray(t *testing.T) {
	customErrorMessage := "error"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{},
		[]Validator{},
	)(nil)

	if errorMessage != nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{},
	[]Validator{},
)(nil) = %v, %t; expected: nil, true
			`,
			getArgs()...,
		)
	}

	// Test 2
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{},
		[]Validator{},
		customErrorMessage,
	)(nil)

	if errorMessage != nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{},
	[]Validator{},
	"error",
)(nil) = %v, %t; expected: nil, true
			`,
			getArgs()...,
		)
	}

	// Test 3
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	// Test 4
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{},
	[]Validator{},
	"error",
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	// Test 5
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{
		MinLength: 3,
	},
	[]Validator{},
)([]string{"a", "b", "c"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	// Test 6
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{
			MinLength: 3,
		},
		[]Validator{},
		customErrorMessage,
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{
		MinLength: 3,
	},
	[]Validator{},
	"error",
)([]string{"a", "b", "c"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
			`,
			getArgs()...,
		)
	}

	// Test 2
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
			`,
			getArgs()...,
		)
	}

	// Test 3
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage == nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{
		MinLength: 3,
	},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: "[error message]", true
			`,
			getArgs()...,
		)
	}

	// Test 4
	errorMessage, stopLoop = IsNullArray(
		arraytype.String,
		Array{
			MinLength:             3,
			MinLengthErrorMessage: customErrorMessage,
		},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	arraytype.String,
	Array{
		MinLength: 3,
		MinLengthErrorMessage: "error",,
	},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: "error", true
			`,
			getArgs()...,
		)
	}
}

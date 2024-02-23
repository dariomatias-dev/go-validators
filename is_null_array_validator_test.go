package validators

import (
	"testing"
)

func TestIsNullArray(t *testing.T) {
	customErrorMessage := "error"

	// Errors
	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage == nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{
		MinLength: 3,
	},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: "[error message]", true
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{
			MinLength: 3,
		},
		[]Validator{},
		customErrorMessage,
	)([]string{"a", "b"})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{
		MinLength: 3,
	},
	[]Validator{},
	"error",
)([]string{"a", "b"}) = %v, %t; expected: "error", true
			`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{},
		[]Validator{},
	)(nil)

	if errorMessage != nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{},
	[]Validator{},
)(nil) = %v, %t; expected: nil, true
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)(nil)

	if errorMessage != nil || !stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{},
	[]Validator{},
	"error",
)(nil) = %v, %t; expected: nil, true
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{},
	[]Validator{},
	"error",
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsNullArray(
	StringArray,
	Array{
		MinLength: 3,
	},
	[]Validator{},
)([]string{"a", "b", "c"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsNullArray(
		StringArray,
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
	StringArray,
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
}

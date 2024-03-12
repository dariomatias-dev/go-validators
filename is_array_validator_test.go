package validators

import (
	"testing"

	arraytype "github.com/dariomatias-dev/go-validators/array_type"
)

func TestIsArray(t *testing.T) {
	customErrorMessage := "error"

	/// AllowEmpty
	// Success
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{
			AllowEmpty: true,
		},
		[]Validator{},
	)([]string{})

	if errorMessage != nil || stopLoop {
		t.Errorf(`
IsArray(
	arraytype.String,
	Array{
		AllowEmpty: true,
	},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	/// MinLength
	// Error
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
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

	// Success
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(`
IsArray(
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

	/// MaxLength
	// Error
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{
			MaxLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c", "d"})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.String,
	Array{
		MaxLength: 3,
	},
	[]Validator{},
)([]string{"a", "b", "c", "d"}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	// Success
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{
			MaxLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(`
IsArray(
	arraytype.String,
	Array{
		MaxLength: 3,
	},
	[]Validator{},
)([]string{"a", "b", "c"}) = %v, %t; expected: nil, false
		`,
			getArgs()...,
		)
	}

	/// interface Array
	// Errors
	errorMessage, stopLoop = IsArray(
		arraytype.Interface,
		Array{},
		[]Validator{},
	)([]string{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	InterfaceArray,
	Array{},
	[]Validator{},
)([]string{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Interface,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]string{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	InterfaceArray,
	Array{},
	[]Validator{},
	"error",
)([]string{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		arraytype.Interface,
		Array{},
		[]Validator{},
	)([]interface{}{"a", nil, "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	InterfaceArray,
	Array{},
	[]Validator{},
)([]interface{}{"a", nil, "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Interface,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{"a", nil, "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	InterfaceArray,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{"a", nil, "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// string Array
	// Errors
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.String,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.String,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.String,
	Array{},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.String,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.String,
	Array{},
	[]Validator{},
	"error",
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// int Array
	// Errors
	errorMessage, stopLoop = IsArray(
		arraytype.Int,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Int,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Int,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Int,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		arraytype.Int,
		Array{},
		[]Validator{},
	)([]int{0, 1, 2})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Int,
	Array{},
	[]Validator{},
)([]int{0, 1, 2}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Int,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]int{0, 1, 2})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Int,
	Array{},
	[]Validator{},
	"error",
)([]int{0, 1, 2}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// float64 Array
	// Errors
	errorMessage, stopLoop = IsArray(
		arraytype.Float64,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Float64,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Float64,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Float64,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		arraytype.Float64,
		Array{},
		[]Validator{},
	)([]float64{0.0, 1.0, 2.0})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Float64,
	Array{},
	[]Validator{},
)([]float64{0.0, 1.0, 2.0}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Float64,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]float64{0.0, 1.0, 2.0})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Float64,
	Array{},
	[]Validator{},
	"error",
)([]float64{0.0, 1.0, 2.0}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// bool Array
	// Errors
	errorMessage, stopLoop = IsArray(
		arraytype.Bool,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Bool,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Bool,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Bool,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		arraytype.Bool,
		Array{},
		[]Validator{},
	)([]bool{true, false})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Bool,
	Array{},
	[]Validator{},
)([]bool{true, false}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Bool,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]bool{true, false})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Bool,
	Array{},
	[]Validator{},
	"error",
)([]bool{true, false}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// any array
	type Struct struct {
		Name string
	}

	// Errors
	errorMessage, stopLoop = IsArray(
		arraytype.Any,
		Array{},
		[]Validator{},
	)("")

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Any,
	Array{},
	[]Validator{},
)("") = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Any,
		Array{},
		[]Validator{},
		customErrorMessage,
	)("")

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	arraytype.Any,
	Array{},
	[]Validator{},
	"error",
)("") = %v, %t; expected: "error", true
		`,
			getArgs()...,
		)
	}

	// Successes
	errorMessage, stopLoop = IsArray(
		arraytype.Any,
		Array{},
		[]Validator{},
	)([]Struct{{Name: "Name"}})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Any,
	Array{},
	[]Validator{},
)([]Struct{{Name: "Name"}}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		arraytype.Any,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]Struct{{Name: "Name"}})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	arraytype.Any,
	Array{},
	[]Validator{},
	"error",
)([]Struct{{Name: "Name"}}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}
}

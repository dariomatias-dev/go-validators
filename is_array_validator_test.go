package validators

import "testing"

func TestIsArray(t *testing.T) {
	customErrorMessage := "error"

	/// AllowEmpty
	// Success
	errorMessage, stopLoop = IsArray(
		StringArray,
		Array{
			AllowEmpty: true,
		},
		[]Validator{},
	)([]string{})

	if errorMessage != nil || stopLoop {
		t.Errorf(`
IsArray(
	StringArray,
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
		StringArray,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
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

	// Success
	errorMessage, stopLoop = IsArray(
		StringArray,
		Array{
			MinLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(`
IsArray(
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

	/// MaxLength
	// Error
	errorMessage, stopLoop = IsArray(
		StringArray,
		Array{
			MaxLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c", "d"})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	StringArray,
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
		StringArray,
		Array{
			MaxLength: 3,
		},
		[]Validator{},
	)([]string{"a", "b", "c"})

	if errorMessage != nil || stopLoop {
		t.Errorf(`
IsArray(
	StringArray,
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
		InterfaceArray,
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
		InterfaceArray,
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
		InterfaceArray,
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
		InterfaceArray,
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
		StringArray,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	StringArray,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		StringArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	StringArray,
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
		StringArray,
		Array{},
		[]Validator{},
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	StringArray,
	Array{},
	[]Validator{},
)([]string{"a", "b"}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		StringArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]string{"a", "b"})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	StringArray,
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
		IntArray,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	IntArray,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		IntArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	IntArray,
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
		IntArray,
		Array{},
		[]Validator{},
	)([]int{0, 1, 2})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	IntArray,
	Array{},
	[]Validator{},
)([]int{0, 1, 2}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		IntArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]int{0, 1, 2})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	IntArray,
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
		Float64Array,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	Float64Array,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		Float64Array,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	Float64Array,
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
		Float64Array,
		Array{},
		[]Validator{},
	)([]float64{0.0, 1.0, 2.0})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	Float64Array,
	Array{},
	[]Validator{},
)([]float64{0.0, 1.0, 2.0}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		Float64Array,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]float64{0.0, 1.0, 2.0})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	Float64Array,
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
		BoolArray,
		Array{},
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	BoolArray,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		BoolArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	BoolArray,
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
		BoolArray,
		Array{},
		[]Validator{},
	)([]bool{true, false})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	BoolArray,
	Array{},
	[]Validator{},
)([]bool{true, false}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		BoolArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]bool{true, false})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	BoolArray,
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
		AnyArray,
		Array{},
		[]Validator{},
	)("")

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
IsArray(
	AnyArray,
	Array{},
	[]Validator{},
)("") = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		AnyArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)("")

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
IsArray(
	AnyArray,
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
		AnyArray,
		Array{},
		[]Validator{},
	)([]Struct{{Name: "Name"}})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	AnyArray,
	Array{},
	[]Validator{},
)([]Struct{{Name: "Name"}}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		AnyArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]Struct{{Name: "Name"}})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
IsArray(
	AnyArray,
	Array{},
	[]Validator{},
	"error",
)([]Struct{{Name: "Name"}}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}
}

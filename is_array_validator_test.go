package validators

import "testing"

func TestIsArray(t *testing.T) {
	customErrorMessage := "error"

	/// interface Array
	// Errors
	errorMessage, stopLoop = IsArray(
		InterfaceArray,
		Array{},
		[]Validator{},
	)([]string{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
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
errorMessage, stopLoop = IsArray(
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
	)([]interface{}{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	InterfaceArray,
	Array{},
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		InterfaceArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	InterfaceArray,
	Array{},
	[]Validator{},
	"error",
)([]interface{}{}) = %v, %t; expected: nil, false
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
errorMessage, stopLoop = IsArray(
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
errorMessage, stopLoop = IsArray(
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
	)([]string{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	StringArray,
	Array{},
	[]Validator{},
)([]string{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		StringArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]string{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	StringArray,
	Array{},
	[]Validator{},
	"error",
)([]string{}) = %v, %t; expected: nil, false
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
errorMessage, stopLoop = IsArray(
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
errorMessage, stopLoop = IsArray(
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
	)([]int{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	IntArray,
	Array{},
	[]Validator{},
)([]int{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		IntArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]int{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	IntArray,
	Array{},
	[]Validator{},
	"error",
)([]int{}) = %v, %t; expected: nil, false
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
errorMessage, stopLoop = IsArray(
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
errorMessage, stopLoop = IsArray(
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
	)([]float64{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	Float64Array,
	Array{},
	[]Validator{},
)([]float64{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		Float64Array,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]float64{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	Float64Array,
	Array{},
	[]Validator{},
	"error",
)([]float64{}) = %v, %t; expected: nil, false
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
errorMessage, stopLoop = IsArray(
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
errorMessage, stopLoop = IsArray(
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
	)([]bool{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	BoolArray,
	Array{},
	[]Validator{},
)([]bool{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		BoolArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]bool{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	BoolArray,
	Array{},
	[]Validator{},
	"error",
)([]bool{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	/// any array
	type Struct struct{}

	// Errors
	errorMessage, stopLoop = IsArray(
		AnyArray,
		Array{},
		[]Validator{},
	)("")

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
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
errorMessage, stopLoop = IsArray(
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
	)([]Struct{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	AnyArray,
	Array{},
	[]Validator{},
)([]Struct{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		AnyArray,
		Array{},
		[]Validator{},
		customErrorMessage,
	)([]Struct{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	AnyArray,
	Array{},
	[]Validator{},
	"error",
)([]Struct{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}
}

package validators

import "testing"

func TestIsArray(t *testing.T) {
	customErrorMessage := "error"

	/// interface Array
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

	/// string Array
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

	/// int Array
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

	/// float64 Array
	// Errors
	errorMessage, stopLoop = IsArray(
		Float64Array,
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	Float64Array,
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		Float64Array,
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	Float64Array,
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
		[]Validator{},
	)([]float64{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	Float64Array,
	[]Validator{},
)([]float64{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		Float64Array,
		[]Validator{},
		customErrorMessage,
	)([]float64{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	Float64Array,
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
		[]Validator{},
	)([]interface{}{})

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	BoolArray,
	[]Validator{},
)([]interface{}{}) = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		BoolArray,
		[]Validator{},
		customErrorMessage,
	)([]interface{}{})

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	BoolArray,
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
		[]Validator{},
	)([]bool{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	BoolArray,
	[]Validator{},
)([]bool{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		BoolArray,
		[]Validator{},
		customErrorMessage,
	)([]bool{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	BoolArray,
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
		[]Validator{},
	)("")

	if errorMessage == nil || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	AnyArray,
	[]Validator{},
)("") = %v, %t; expected: "[error message]", true
		`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		AnyArray,
		[]Validator{},
		customErrorMessage,
	)("")

	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf(`
errorMessage, stopLoop = IsArray(
	AnyArray,
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
		[]Validator{},
	)([]Struct{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	AnyArray,
	[]Validator{},
)([]Struct{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}

	errorMessage, stopLoop = IsArray(
		AnyArray,
		[]Validator{},
		customErrorMessage,
	)([]Struct{})

	if errorMessage != nil || stopLoop {
		t.Errorf(
			`
errorMessage, stopLoop = IsArray(
	AnyArray,
	[]Validator{},
	"error",
)([]Struct{}) = %v, %t; expected: nil, false
			`,
			getArgs()...,
		)
	}
}

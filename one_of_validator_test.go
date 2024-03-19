package validators

import (
	"testing"

	arraytype "github.com/dariomatias-dev/go-validators/array_type"
)

func TestOneOf(t *testing.T) {
	stringOptions := []string{"one", "two", "three"}
	intOptions := []int{1, 2, 3}
	float64Options := []float64{1.0, 2.0, 3.0}

	/// - Successes
	// Test 1
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions)("one")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"one\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions, customErrorMessage)("one")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"one\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = OneOf(arraytype.Int, intOptions)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.Int, []int{1, 2, 3})(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = OneOf(arraytype.Int, intOptions, customErrorMessage)(1)
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.Int, []int{1, 2, 3})(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = OneOf(arraytype.Float64, float64Options)(1.0)
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.Float64, []float64{1.0, 2.0, 3.0})(1.0) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = OneOf(arraytype.Float64, float64Options, customErrorMessage)(1.0)
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.Float64, []float64{1.0, 2.0, 3.0})(1.0) = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions)("four")
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"four\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions, customErrorMessage)("four")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"four\") = %v, %t; expected: \"error\"], false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = OneOf(arraytype.Int, intOptions)(4)
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf(arraytype.Int, []int{1, 2, 3})(4) = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = OneOf(arraytype.Int, intOptions, customErrorMessage)(4)
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf(arraytype.Int, []int{1, 2, 3})(4) = %v, %t; expected: \"error\", false", getArgs()...)
	}

	// Test 5
	errorMessage, stopLoop = OneOf(arraytype.Float64, float64Options)(4.0)
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf(arraytype.Float64, []float64{1.0, 2.0, 3.0})(4.0) = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 6
	errorMessage, stopLoop = OneOf(arraytype.Float64, float64Options, customErrorMessage)(4.0)
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf(arraytype.Float64, []float64{1.0, 2.0, 3.0})(4.0) = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

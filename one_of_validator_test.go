package validators

import (
	"testing"

	arraytype "github.com/dariomatias-dev/go-validators/array_type"
)

func TestOneOf(t *testing.T) {
	stringOptions := []string{"one", "two", "three"}

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
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions)("three")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"three\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions, customErrorMessage)("three")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"three\") = %v, %t; expected: nil, false", getArgs()...)
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
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions)("five")
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"five\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = OneOf(arraytype.String, stringOptions, customErrorMessage)("five")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("OneOf(arraytype.String, []string{\"one\", \"two\", \"three\"})(\"five\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

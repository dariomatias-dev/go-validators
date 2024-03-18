package validators

import "testing"

func TestOneOf(t *testing.T) {
	options := []string{"one", "two", "three"}

	/// - Successes
	// Test 1
	errorMessage, stopLoop = OneOf(options)("one")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"one\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = OneOf(options, customErrorMessage)("one")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"one\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = OneOf(options)("three")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"three\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = OneOf(options, customErrorMessage)("three")
	if errorMessage != nil || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"three\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = OneOf(options)("four")
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"four\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = OneOf(options, customErrorMessage)("four")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"four\") = %v, %t; expected: \"error\"], false", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = OneOf(options)("five")
	if errorMessage == nil || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"five\") = %v, %t; expected: [error message], false", getArgs()...)
	}

	// Test 4
	errorMessage, stopLoop = OneOf(options, customErrorMessage)("five")
	if errorMessage == nil || *errorMessage != customErrorMessage || stopLoop {
		t.Errorf("OneOf([]string{\"one\", \"two\", \"three\"})(\"five\") = %v, %t; expected: \"error\", false", getArgs()...)
	}
}

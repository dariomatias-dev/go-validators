package validators

import "testing"

func TestIsString(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsString()("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsString()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsString(customErrorMessage)("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsString(\"error\")(\"aA\") = %v, %t; expected: nil, true", getArgs()...)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = IsString()(0)
	if errorMessage == nil || !stopLoop {
		t.Errorf("IsString()(0) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsString(customErrorMessage)(0)
	if errorMessage == nil || *errorMessage != customErrorMessage || !stopLoop {
		t.Errorf("IsString(\"error\")(0) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	errorMessage, stopLoop = IsString(valueErrorMessage)(0)
	if errorMessage == nil || formattingPlaceholdersPattern.MatchString(*errorMessage) || !stopLoop {
		t.Errorf("IsString(\"error: value is %%T\")(0) = %v, %t; expected: \"error: value is int\", true", getArgs()...)
	}
}

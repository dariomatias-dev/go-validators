package validators

import "testing"

func TestIsOptional(t *testing.T) {
	/// - Successes
	// Test 1
	errorMessage, stopLoop = IsOptional()(nil)
	if errorMessage != nil || !stopLoop {
		t.Errorf("IsOptional()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsOptional()("")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsOptional()(\"\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	errorMessage, stopLoop = IsOptional()("aA")
	if errorMessage != nil || stopLoop {
		t.Errorf("IsOptional()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}
}

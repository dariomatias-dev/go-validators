package validators

import "testing"

func TestIsOptional(t *testing.T) {
	/// - Successes
	// Test 1
	err, stopLoop = IsOptional()(nil)
	if err != nil || !stopLoop {
		t.Errorf("IsOptional()(nil) = %v, %t; expected: nil, true", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsOptional()("")
	if err != nil || stopLoop {
		t.Errorf("IsOptional()(\"\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, stopLoop = IsOptional()("aA")
	if err != nil || stopLoop {
		t.Errorf("IsOptional()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}
}

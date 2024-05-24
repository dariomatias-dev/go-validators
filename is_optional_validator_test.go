package validators

import "testing"

func TestIsOptional(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsOptional()(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("IsOptional()(nil) = %v, %t; expected: nil, true"))
	}

	// Test 2
	err, abortValidation = IsOptional()("")
	if err != nil || abortValidation {
		t.Errorf(setError("IsOptional()(\"\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsOptional()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("IsOptional()(\"aA\") = %v, %t; expected: nil, false"))
	}
}

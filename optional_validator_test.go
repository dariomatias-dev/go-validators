package validators

import (
	"testing"
)

func TestOptional(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Optional()(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("Optional()(nil) = %v, %t; expected: nil, true"))
	}

	// Test 2
	err, abortValidation = Optional()("")
	if err != nil || abortValidation {
		t.Errorf(setError("Optional()(\"\") = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Optional()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("Optional()(\"aA\") = %v, %t; expected: nil, false"))
	}
}

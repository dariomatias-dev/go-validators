package validators

import (
	"testing"
)

func TestMax(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Max(2)(1)
	if err != nil || abortValidation {
		t.Errorf(setError("Max(2)(1) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Max(2, errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf(setError("Max(2, \"error\")(1) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = Max(2.1)(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("Max(2.1)(1.1) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Max(2.1, errCustomMessage)(1.1)
	if err != nil || abortValidation {
		t.Errorf(setError("Max(2.1, \"error\")(1.1) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Max(1)(2)
	if err == nil || abortValidation {
		t.Errorf(setError("Max(1)(2) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = Max(1, errCustomMessage)(2)
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Max(1, \"error\")(2) = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = Max(1.1)(1.2)
	if err == nil || abortValidation {
		t.Errorf(setError("Max(1.1)(1.2) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = Max(1.1, errCustomMessage)(1.2)
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Max(1.1, \"error\")(1.2) = %v, %t; expected: \"error\", false"))
	}
}

package validators

import (
	"testing"
)

func TestIsInt(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsInt()(1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsInt()(1) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsInt(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsInt(\"error\")(1) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsInt()(nil)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsInt()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsInt(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsInt(\"error\")(nil) = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = IsInt()("")
	if err == nil || !abortValidation {
		t.Errorf(setError("IsInt()(\"\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = IsInt(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsInt(\"error\")(\"\") = %v, %t; expected: \"error\", true"))
	}

	// Test 5
	err, abortValidation = IsInt()(1.1)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsInt()(1.1) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 6
	err, abortValidation = IsInt(errCustomMessage)(1.1)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsInt(\"error\")(1.1) = %v, %t; expected: \"error\", true"))
	}
}

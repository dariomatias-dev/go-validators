package validators

import (
	"testing"
)

func TestRequired(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Required()(1)
	if err != nil || abortValidation {
		t.Errorf(setError("Required()(1) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Required(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf(setError("Required(\"error\")(1) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = Required()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("Required()(\"aA\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Required(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("Required(\"error\")(\"aA\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Required()(nil)
	if err == nil || !abortValidation {
		t.Errorf(setError("Required()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = Required(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("Required(\"error\")(nil) = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = Required()("")
	if err == nil || !abortValidation {
		t.Errorf(setError("Required()(\"\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = Required(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("Required(\"error\")(\"\") = %v, %t; expected: \"error\", true"))
	}

	// Test 5
	err, abortValidation = Required()("  ")
	if err == nil || !abortValidation {
		t.Errorf(setError("Required()(\"  \") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 6
	err, abortValidation = Required(errCustomMessage)("  ")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("Required(\"error\")(\"  \") = %v, %t; expected: \"error\", true"))
	}
}

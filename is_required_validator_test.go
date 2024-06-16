package validators

import (
	"testing"
)

func TestIsRequired(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsRequired()(1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsRequired()(1) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = IsRequired(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf(setError("IsRequired(\"error\")(1) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = IsRequired()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("IsRequired()(\"aA\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = IsRequired(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("IsRequired(\"error\")(\"aA\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsRequired()(nil)
	if err == nil || !abortValidation {
		t.Errorf(setError("IsRequired()(nil) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = IsRequired(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsRequired(\"error\")(nil) = %v, %t; expected: \"error\", true"))
	}

	// Test 3
	err, abortValidation = IsRequired()("")
	if err == nil || !abortValidation {
		t.Errorf(setError("IsRequired()(\"\") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 4
	err, abortValidation = IsRequired(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsRequired(\"error\")(\"\") = %v, %t; expected: \"error\", true"))
	}

	// Test 5
	err, abortValidation = IsRequired()("  ")
	if err == nil || !abortValidation {
		t.Errorf(setError("IsRequired()(\"  \") = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 6
	err, abortValidation = IsRequired(errCustomMessage)("  ")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("IsRequired(\"error\")(\"  \") = %v, %t; expected: \"error\", true"))
	}
}

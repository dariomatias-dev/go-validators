package validators

import "testing"

func TestIsRequired(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = IsRequired()(1)
	if err != nil || abortValidation {
		t.Errorf("IsRequired()(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsRequired(errCustomMessage)(1)
	if err != nil || abortValidation {
		t.Errorf("IsRequired(\"error\")(1) = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsRequired()("aA")
	if err != nil || abortValidation {
		t.Errorf("IsRequired()(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsRequired(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf("IsRequired(\"error\")(\"aA\") = %v, %t; expected: nil, false", getArgs()...)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsRequired()(nil)
	if err == nil || !abortValidation {
		t.Errorf("IsRequired()(nil) = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 2
	err, abortValidation = IsRequired(errCustomMessage)(nil)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsRequired(\"error\")(nil) = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 3
	err, abortValidation = IsRequired()("")
	if err == nil || !abortValidation {
		t.Errorf("IsRequired()(\"\") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 4
	err, abortValidation = IsRequired(errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsRequired(\"error\")(\"\") = %v, %t; expected: \"error\", true", getArgs()...)
	}

	// Test 5
	err, abortValidation = IsRequired()("  ")
	if err == nil || !abortValidation {
		t.Errorf("IsRequired()(\"  \") = %v, %t; expected: \"[error message]\", true", getArgs()...)
	}

	// Test 6
	err, abortValidation = IsRequired(errCustomMessage)("  ")
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf("IsRequired(\"error\")(\"  \") = %v, %t; expected: \"error\", true", getArgs()...)
	}
}

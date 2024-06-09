package validators

import (
	"testing"
)

func TestIsArray(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, abortValidation = IsArray()([]string{})

	if err != nil || abortValidation {
		t.Errorf(
			`IsArray()([]string{}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	// Test 2
	err, abortValidation = IsArray()([1]string{""})

	if err != nil || abortValidation {
		t.Errorf(
			`IsArray()([1]string{""}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsArray()(nil)

	if err == nil || !abortValidation {
		t.Errorf(
			"IsArray()(nil) = %v, %t; expected: \"[error message]\", true",
			getArgs()...,
		)
	}

	// Test 2
	err, abortValidation = IsArray(errCustomMessage)(nil)

	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(
			"IsArray(\"error\")(nil) = %v, %t; expected: \"error\", true",
			getArgs()...,
		)
	}

	// Test 3
	err, abortValidation = IsArray()("")

	if err == nil || !abortValidation {
		t.Errorf(
			"IsArray()(\"\") = %v, %t; expected: \"[error message]\", true",
			getArgs()...,
		)
	}

	// Test 4
	err, abortValidation = IsArray(errCustomMessage)("")

	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(
			"IsArray(\"error\")(\"\") = %v, %t; expected: \"error\", true",
			getArgs()...,
		)
	}
}

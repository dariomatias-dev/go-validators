package validators

import (
	"testing"
)

func TestNullValidator(t *testing.T) {
	isNullString := func(errorMessage ...string) Validator {
		return isNullValidator(
			IsString,
			"The value is not a string or null: value is %T.",
			errorMessage,
		)
	}

	/// - Successes
	// Test 1
	err, abortValidation = isNullString()(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("isNullString()(nil) = %v, %t; expected: nil, true"))
	}

	// Test 2
	err, abortValidation = isNullString(errCustomMessage)(nil)
	if err != nil || !abortValidation {
		t.Errorf(setError("isNullString(\"error\")(nil) = %v, %t; expected: nil, true"))
	}

	// Test 3
	err, abortValidation = isNullString()("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("isNullString()(\"aA\") = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = isNullString(errCustomMessage)("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("isNullString(\"error\")(\"aA\") = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = isNullString()(0)
	if err == nil || !abortValidation {
		t.Errorf(setError("isNullString()(0) = %v, %t; expected: \"[error message]\", true"))
	}

	// Test 2
	err, abortValidation = isNullString(errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(setError("isNullString(\"error\")(0) = %v, %t; expected: \"error\", true"))
	}
}

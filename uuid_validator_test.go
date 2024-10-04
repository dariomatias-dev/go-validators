package validators

import (
	"testing"
)

func TestUuid(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = Uuid(2)("11d50d80-807c-22a4-bb46-bb149cf3b146")
	if err != nil || abortValidation {
		t.Errorf(setError("Uuid(2)(\"11d50d80-807c-22a4-bb46-bb149cf3b146\") == %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = Uuid(2, errCustomMessage)("11d50d80-807c-22a4-bb46-bb149cf3b146")
	if err != nil || abortValidation {
		t.Errorf(setError("Uuid(2, \"error\")(\"11d50d80-807c-22a4-bb46-bb149cf3b146\") == %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = Uuid(3)("6fa459ea-ee8a-3ca4-894e-db77e160355e")
	if err != nil || abortValidation {
		t.Errorf(setError("Uuid(3)(\"6fa459ea-ee8a-3ca4-894e-db77e160355e\") == %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Uuid(3, errCustomMessage)("6fa459ea-ee8a-3ca4-894e-db77e160355e")
	if err != nil || abortValidation {
		t.Errorf(setError("Uuid(3, \"error\")(\"6fa459ea-ee8a-3ca4-894e-db77e160355e\") == %v, %t; expected: nil, false"))
	}

	// Test 5
	err, abortValidation = Uuid(4)("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err != nil || abortValidation {
		t.Errorf(setError("Uuid(4)(\"f47ac10b-58cc-4372-a567-0e02b2c3d479\") == %v, %t; expected: nil, false"))
	}

	// Test 6
	err, abortValidation = Uuid(4, errCustomMessage)("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err != nil || abortValidation {
		t.Errorf(setError("Uuid(4, \"error\")(\"f47ac10b-58cc-4372-a567-0e02b2c3d479\") == %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Uuid(2)("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err == nil || abortValidation {
		t.Errorf(setError("Uuid(2)(\"f47ac10b-58cc-4372-a567-0e02b2c3d479\") == %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = Uuid(2, errCustomMessage)("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Uuid(2, \"error\")(\"f47ac10b-58cc-4372-a567-0e02b2c3d479\") == %v, %t; expected: \"error\", false"))
	}
}

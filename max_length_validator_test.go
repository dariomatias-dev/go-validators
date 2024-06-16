package validators

import (
	"testing"
)

func TestMaxLength(t *testing.T) {
	/// - Successes
	/// String size
	// Test 1
	err, abortValidation = MaxLength(5)("aA")
	if err != nil || abortValidation {
		t.Errorf(setError("MaxLength(5)(\"aA\") = %v, %t; expected: nil, false"))
	}

	/// Array Size
	// Test 2
	err, abortValidation = MaxLength(5)([]int{0, 1, 2, 3, 4})
	if err != nil || abortValidation {
		t.Errorf(setError("MaxLength(5)([]int{0, 1, 2, 3, 4}) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = MaxLength(5)([5]int{0, 1, 2, 3, 4})
	if err != nil || abortValidation {
		t.Errorf(setError("MaxLength(5)([5]int{0, 1, 2, 3, 4}) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = MaxLength(5)([5]int{0, 1, 2, 3})
	if err != nil || abortValidation {
		t.Errorf(setError("MaxLength(5)([5]int{0, 1, 2, 3}) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	/// String size
	// Test 1
	err, abortValidation = MaxLength(5)("aAbBcC")
	if err == nil || abortValidation {
		t.Errorf(setError("MaxLength(5)(\"aAbBcC\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = MaxLength(5, errCustomMessage)("aAbBcC")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MaxLength(5, \"error\")(\"aAbBcC\") = %v, %t; expected: \"error\", false"))
	}

	/// Array Size
	// Test 3
	err, abortValidation = MaxLength(5)([]int{0, 1, 2, 3, 4, 5})
	if err == nil || abortValidation {
		t.Errorf(setError("MaxLength(5)([]int{0, 1, 2, 3, 4, 5}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = MaxLength(5, errCustomMessage)([]int{0, 1, 2, 3, 4, 5})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MaxLength(5, \"error\")([]int{0, 1, 2, 3, 4, 5}) = %v, %t; expected: \"error\", false"))
	}

	// Test 5
	err, abortValidation = MaxLength(5)([6]int{0, 1, 2, 3, 4, 5})
	if err == nil || abortValidation {
		t.Errorf(setError("MaxLength(5)([6]int{0, 1, 2, 3, 4, 5}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 6
	err, abortValidation = MaxLength(5, errCustomMessage)([6]int{0, 1, 2, 3, 4, 5})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MaxLength(5, \"error\")([6]int{0, 1, 2, 3, 4, 5}) = %v, %t; expected: \"error\", false"))
	}

	// Test 7
	err, abortValidation = MaxLength(5)([6]int{0, 1, 2, 3})
	if err == nil || abortValidation {
		t.Errorf(setError("MaxLength(5)([6]int{0, 1, 2, 3}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 8
	err, abortValidation = MaxLength(5, errCustomMessage)([6]int{0, 1, 2, 3})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MaxLength(5, \"error\")([6]int{0, 1, 2, 3}) = %v, %t; expected: \"error\", false"))
	}
}

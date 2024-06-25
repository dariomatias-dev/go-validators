package validators

import (
	"testing"
)

func TestMinLength(t *testing.T) {
	/// - Successes
	/// String size
	// Test 1
	err, abortValidation = MinLength(5)("aAbBcC")
	if err != nil || abortValidation {
		t.Errorf(setError("MinLength(5)(\"aAbBcC\") = %v, %t; expected: nil, false"))
	}

	/// Slice Size
	// Test 2
	err, abortValidation = MinLength(5)([]int{0, 1, 2, 3, 4})
	if err != nil || abortValidation {
		t.Errorf(setError("MinLength(5)([]int{0, 1, 2, 3, 4}) = %v, %t; expected: nil, false"))
	}

	/// Array Size
	// Test 3
	err, abortValidation = MinLength(5)([5]int{0, 1, 2, 3})
	if err != nil || abortValidation {
		t.Errorf(setError("MinLength(5)([5]int{0, 1, 2, 3}) = %v, %t; expected: nil, false"))
	}

	// Test 3
	err, abortValidation = MinLength(5)([5]int{0, 1, 2, 3, 4})
	if err != nil || abortValidation {
		t.Errorf(setError("MinLength(5)([5]int{0, 1, 2, 3, 4}) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	/// String size
	// Test 1
	err, abortValidation = MinLength(5)("")
	if err == nil || abortValidation {
		t.Errorf(setError("MinLength(5)(\"\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = MinLength(5, errCustomMessage)("")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MinLength(5, \"error\")(\"\") = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = MinLength(5)("aA")
	if err == nil || abortValidation {
		t.Errorf(setError("MinLength(5)(\"aA\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = MinLength(5, errCustomMessage)("aA")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MinLength(5, \"error\")(\"aA\") = %v, %t; expected: \"error\", false"))
	}

	/// Slice Size
	// Test 1
	err, abortValidation = MinLength(5)([]int{})
	if err == nil || abortValidation {
		t.Errorf(setError("MinLength(5)([]int{}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = MinLength(5, errCustomMessage)([]int{})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MinLength(5, \"error\")([]int{}) = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = MinLength(5)([]int{0, 1, 2, 3})
	if err == nil || abortValidation {
		t.Errorf(setError("MinLength(5)([]int{0, 1, 2, 3}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = MinLength(5, errCustomMessage)([]int{0, 1, 2, 3})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MinLength(5, \"error\")([]int{0, 1, 2, 3}) = %v, %t; expected: \"error\", false"))
	}

	/// Array Size
	// Test 5
	err, abortValidation = MinLength(5)([4]int{0, 1, 2, 3})
	if err == nil || abortValidation {
		t.Errorf(setError("MinLength(5)([4]int{0, 1, 2, 3}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 6
	err, abortValidation = MinLength(5, errCustomMessage)([4]int{0, 1, 2, 3})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("MinLength(5, \"error\")([4]int{0, 1, 2, 3}) = %v, %t; expected: \"error\", false"))
	}
}

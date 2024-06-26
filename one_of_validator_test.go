package validators

import (
	"testing"
)

func TestOneOf(t *testing.T) {
	/// - Successes
	// Test 1
	err, abortValidation = OneOf([]int{1, 2, 3})(2)
	if err != nil || abortValidation {
		t.Errorf(setError("OneOf([]int{1, 2, 3})(2) = %v, %t; expected: nil, false"))
	}

	// Test 2
	err, abortValidation = OneOf([3]int{1, 2, 3})(2)
	if err != nil || abortValidation {
		t.Errorf(setError("OneOf([3]int{1, 2, 3})(2) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	// Test 1
	err, abortValidation = OneOf(nil)(0)
	if err == nil || abortValidation {
		t.Errorf(setError("OneOf([]int{1, 2, 3})(0) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = OneOf([]int{1, 2, 3})(0)
	if err == nil || abortValidation {
		t.Errorf(setError("OneOf([]int{1, 2, 3})(0) = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = OneOf([]int{1, 2, 3}, errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("OneOf([]int{1, 2, 3}, \"error\")(0) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = OneOf([3]int{1, 2, 3})(0)
	if err == nil || abortValidation {
		t.Errorf(setError("OneOf([3]int{1, 2, 3})(0) = %v, %t; expected: \"error\", false"))
	}

	// Test 3
	err, abortValidation = OneOf([3]int{1, 2, 3}, errCustomMessage)(0)
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("OneOf([3]int{1, 2, 3}, \"error\")(0) = %v, %t; expected: \"[error message]\", false"))
	}
}

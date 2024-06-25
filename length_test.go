package validators

import (
	"testing"
)

func TestLength(t *testing.T) {
	/// - Successes
	/// String size
	// Test 1
	err, abortValidation = Length(5)("aAbBc")
	if err != nil || abortValidation {
		t.Errorf(setError("Length(5)(\"aAbBc\") = %v, %t; expected: nil, false"))
	}

	/// Slice Size
	// Test 2
	err, abortValidation = Length(5)([]int{0, 1, 2, 3, 4})
	if err != nil || abortValidation {
		t.Errorf(setError("Length(5)([]int{0, 1, 2, 3, 4}) = %v, %t; expected: nil, false"))
	}

	/// Array Size
	// Test 3
	err, abortValidation = Length(5)([5]int{0, 1, 2, 3, 4})
	if err != nil || abortValidation {
		t.Errorf(setError("Length(5)([5]int{0, 1, 2, 3, 4}) = %v, %t; expected: nil, false"))
	}

	// Test 4
	err, abortValidation = Length(5)([5]int{0, 1, 2, 3})
	if err != nil || abortValidation {
		t.Errorf(setError("Length(5)([5]int{0, 1, 2, 3}) = %v, %t; expected: nil, false"))
	}

	/// - Errors
	/// String size
	// Test 1
	err, abortValidation = Length(5)("aAbB")
	if err == nil || abortValidation {
		t.Errorf(setError("Length(5)(\"aAbB\") = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 2
	err, abortValidation = Length(5, errCustomMessage)("aAbB")
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Length(5, \"error\")(\"aAbB\") = %v, %t; expected: \"error\", false"))
	}

	/// Slice Size
	// Test 3
	err, abortValidation = Length(5)([]int{0, 1, 2, 3})
	if err == nil || abortValidation {
		t.Errorf(setError("Length(5)([]int{0, 1, 2, 3}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 4
	err, abortValidation = Length(5, errCustomMessage)([]int{0, 1, 2, 3})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Length(5, \"error\")([]int{0, 1, 2, 3}) = %v, %t; expected: \"error\", false"))
	}

	/// Array Size
	// Test 5
	err, abortValidation = Length(5)([4]int{0, 1, 2, 3})
	if err == nil || abortValidation {
		t.Errorf(setError("Length(5)([4]int{0, 1, 2, 3}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 6
	err, abortValidation = Length(5, errCustomMessage)([4]int{0, 1, 2, 3})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Length(5, \"error\")([4]int{0, 1, 2, 3}) = %v, %t; expected: \"error\", false"))
	}

	// Test 7
	err, abortValidation = Length(5)([6]int{0, 1, 2, 3, 4})
	if err == nil || abortValidation {
		t.Errorf(setError("Length(5)([6]int{0, 1, 2, 3, 4}) = %v, %t; expected: \"[error message]\", false"))
	}

	// Test 8
	err, abortValidation = Length(5, errCustomMessage)([6]int{0, 1, 2, 3, 4})
	if err == nil || err.Error() != errCustom.Error() || abortValidation {
		t.Errorf(setError("Length(5, \"error\")([6]int{0, 1, 2, 3, 4}) = %v, %t; expected: \"error\", false"))
	}
}

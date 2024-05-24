package validators

import "testing"

func TestRegex(t *testing.T) {
	numbersPattern := "^[0-9]+$"
	lettersPattern := "^[a-zA-Z]+$"

	/// - Successes
	// Test 1
	err, abortValidation = Regex(
		numbersPattern,
		"Invalid value",
	)("0123")
	if err != nil || abortValidation {
		t.Errorf(setError(`
Regex(
	"^[0-9]+$",
	"Invalid value",
)("0123") = %s, %t; expected: nil, false
		`,
		))
	}

	// Test 2
	lettersPattern = "^[a-zA-Z]+$"
	err, abortValidation = Regex(
		lettersPattern,
		"Invalid value",
	)("aAbB")
	if err != nil || abortValidation {
		t.Errorf(setError(`
Regex(
	"^[a-zA-Z]+$",
	"Invalid value",
)("aAbB") = %s, %t; expected: nil, false
		`,
		))
	}

	/// - Errors
	// Test 1
	err, abortValidation = Regex(
		numbersPattern,
		"Invalid value",
	)("aA01")
	if err == nil || abortValidation {
		t.Errorf(setError(`
Regex(
	"^[0-9]+$",
	"Invalid value",
)("aA01") = %s, %t; expected: "Invalid value", false
		`,
		))
	}

	// Test 2
	err, abortValidation = Regex(
		lettersPattern,
		"Invalid value",
	)("aA01")
	if err == nil || abortValidation {
		t.Errorf(setError(`
Regex(
	"^[a-zA-Z]+$",
	"Invalid value",
)("aA01") = %s, %t; expected: "Invalid value", false
		`,
		))
	}
}

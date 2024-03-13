package validators

import "testing"

func TestRegex(t *testing.T) {
	numbersPattern := "^[0-9]+$"
	lettersPattern := "^[a-zA-Z]+$"

	/// - Successes
	// Test 1
	errorMessage, stopLoop = Regex(
		numbersPattern,
		"Invalid value",
	)("0123")
	if errorMessage != nil || stopLoop {
		t.Errorf(`
Regex(
	"^[0-9]+$",
	"Invalid value",
)("0123") = %s, %t; expected: nil, false
		`,
			getArgs()...,
		)
	}

	// Test 2
	lettersPattern = "^[a-zA-Z]+$"
	errorMessage, stopLoop = Regex(
		lettersPattern,
		"Invalid value",
	)("aAbB")
	if errorMessage != nil || stopLoop {
		t.Errorf(`
Regex(
	"^[a-zA-Z]+$",
	"Invalid value",
)("aAbB") = %s, %t; expected: nil, false
		`,
			getArgs()...,
		)
	}

	/// - Errors
	// Test 1
	errorMessage, stopLoop = Regex(
		numbersPattern,
		"Invalid value",
	)("aA01")
	if errorMessage == nil || stopLoop {
		t.Errorf(`
Regex(
	"^[0-9]+$",
	"Invalid value",
)("aA01") = %s, %t; expected: "[error message]", false
		`,
			getArgs()...,
		)
	}

	// Test 2
	errorMessage, stopLoop = Regex(
		lettersPattern,
		"Invalid value",
	)("aA01")
	if errorMessage == nil || stopLoop {
		t.Errorf(`
Regex(
	"^[a-zA-Z]+$",
	"Invalid value",
)("aA01") = %s, %t; expected: "[error message]", false
		`,
			getArgs()...,
		)
	}
}

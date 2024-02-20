package validators

import "testing"

func TestRegex(t *testing.T) {
	numbersPattern := "^[0-9]+$"
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

	lettersPattern := "^[a-zA-Z]+$"
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
}

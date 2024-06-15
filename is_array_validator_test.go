package validators

import (
	"testing"
)

func TestIsArray(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
	)([]string{})

	if err != nil || abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)([]string{}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	// Test 2
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
	)([1]string{""})

	if err != nil || abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)([1]string{""}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	// Test 3
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
	)([]string{"Item1", "Item2"})

	if err != nil || abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)([]string{"Item1", "Item2"}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	// Test 4
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
	)([1]string{""})

	if err != nil || abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)([2]string{"Item1", "Item2"}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
	)(nil)

	if err == nil || !abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)(nil) = %v, %t; expected: "[error message]", true`,
			getArgs()...,
		)
	}

	// Test 2
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
		errCustomMessage,
	)(nil)

	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
	"error",
)(nil) = %v, %t; expected: "error", true`,
			getArgs()...,
		)
	}

	// Test 3
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
	)("")

	if err == nil || !abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)("") = %v, %t; expected: "[error message]", true`,
			getArgs()...,
		)
	}

	// Test 4
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
		},
		errCustomMessage,
	)("")

	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
	"error",
)("") = %v, %t; expected: "error", true`,
			getArgs()...,
		)
	}

	// Test 5
	err, abortValidation = IsArray(
		[]Validator{
			IsString(),
			MaxLength(3),
		},
	)([]string{"Item1", "Item2"})

	if err == nil || !abortValidation {
		t.Errorf(
			`
IsArray(
	[]Validator{
		IsString(),
	},
)([]string{"Item1", "Item2"}) = %v, %t; expected: "[error message]", true`,
			getArgs()...,
		)
	}
}

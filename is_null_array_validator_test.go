package validators

import (
	"testing"
)

func TestIsNullArray(t *testing.T) {
	/// - Succcesses
	// Test 1
	err, abortValidation = IsNullArray(
		[]Validator{
			IsString(),
		},
	)(nil)

	if err != nil || !abortValidation {
		t.Errorf(
			`
IsNullArray(
	[]Validator{
		IsString(),
	},
)(nil) = %v, %t; expected: nil, true`,
			getArgs()...,
		)
	}

	// Test 2
	err, abortValidation = IsNullArray(
		[]Validator{
			IsString(),
		},
	)([]string{})

	if err != nil || abortValidation {
		t.Errorf(
			`
IsNullArray(
	[]Validator{
		IsString(),
	},
)([]string{}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	// Test 3
	err, abortValidation = IsNullArray(
		[]Validator{
			IsString(),
		},
	)([1]string{""})

	if err != nil || abortValidation {
		t.Errorf(
			`
IsNullArray(
	[]Validator{
		IsString(),
	},
)([1]string{""}) = %v, %t; expected: nil, false`,
			getArgs()...,
		)
	}

	/// - Errors
	// Test 1
	err, abortValidation = IsNullArray(
		[]Validator{
			IsString(),
		},
	)("")

	if err == nil || !abortValidation {
		t.Errorf(
			`
IsNullArray(
	[]Validator{
		IsString(),
	},
)(\"\") = %v, %t; expected: \"[error message]\", true`,
			getArgs()...,
		)
	}

	// Test 2
	err, abortValidation = IsNullArray(
		[]Validator{
			IsString(),
		},
		errCustomMessage,
	)("")

	if err == nil || err.Error() != errCustom.Error() || !abortValidation {
		t.Errorf(
			`
IsNullArray(
	[]Validator{
		IsString(),
	},
	"error",
)(\"\") = %v, %t; expected: \"error\", true`,
			getArgs()...,
		)
	}
}

package validators

import (
	"testing"
)

func TestValidate(t *testing.T) {
	extractValue := func(value *[]string) any {
		if value == nil {
			return nil
		} else {
			return *value
		}
	}

	// Test 1
	test1 := func(
		value any,
		expected any,
		condition func(errorsLen int) bool,
	) {
		errors := Validate(
			value,
			IsInt(),
			Min(3),
			Max(10),
		)

		if condition(len(*errors)) {
			t.Errorf(`
Validate(
	%v,
	IsInt(),
	Min(3),
	Max(10),
) = %v; expected: %v
			`,
				value,
				extractValue(errors),
				expected,
			)
		}
	}

	/// Errors
	test1(
		"4",
		"[error message(s)]",
		func(errorsLen int) bool {
			return errorsLen == 0
		},
	)

	test1(
		2,
		"[error message(s)]",
		func(errorsLen int) bool {
			return errorsLen == 0
		},
	)

	/// Success
	test1(
		4,
		nil,
		func(errorsLen int) bool {
			return errorsLen != 0
		},
	)

	// Test 2
	test2 := func(
		value any,
		expected any,
		condition func(errorsLen int) bool,
	) {
		errors := Validate(
			value,
			IsString(),
			MinLength(3),
			MaxLength(10),
		)

		if condition(len(*errors)) {
			t.Errorf(`
errors := Validate(
	%v,
	IsString(),
	MinLength(3),
	MaxLength(10),
) = %v; expected: %v
			`,
				value,
				extractValue(errors),
				expected,
			)
		}
	}

	/// Errors
	test2(
		nil,
		"[error message(s)]",
		func(errorsLen int) bool {
			return errorsLen == 0
		},
	)

	test2(
		"Na",
		"[error message(s)]",
		func(errorsLen int) bool {
			return errorsLen == 0
		},
	)

	/// Success
	test2(
		"Name",
		nil,
		func(errorsLen int) bool {
			return errorsLen != 0
		},
	)

}

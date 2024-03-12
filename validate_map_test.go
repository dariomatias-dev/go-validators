package validators

import (
	"fmt"
	"testing"
)

func TestValidateMap(t *testing.T) {
	validations := Validators{
		"name": []Validator{
			IsString(),
			MinLength(3),
			MaxLength(20),
		},
		"age": []Validator{
			IsInt(),
			Min(18),
			Max(100),
		},
		"email": []Validator{
			Email(),
		},
	}

	extractValue := func(value *map[string][]string) any {
		if value == nil {
			return nil
		} else {
			return *value
		}
	}

	generateErrorMessage := func(
		errors *map[string][]string,
		expected any,
	) string {
		return fmt.Sprintf(
			`
Validators{
	"name": []Validator{
		IsString(),
		MinLength(3),
		MaxLength(20),
	},
	"age": []Validator{
		IsInt(),
		Min(18),
		Max(100),
	},
	"email": []Validator{
		Email(),
	},
} = %v; expected: %v
			`,
			extractValue(errors),
			expected,
		)
	}

	data := map[string]interface{}{
		"name":  "Name",
		"age":   15,
		"email": "emailexample@gmail.com",
	}

	// Errors
	errors := ValidateMap(validations)(data)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[error message(s)]"),
		)
	}

	data["email"] = "emailexample"
	errors = ValidateMap(validations)(data)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[error message(s)]"),
		)
	}

	// Success
	data["age"] = 18
	data["email"] = "emailexample@gmail.com"
	errors = ValidateMap(validations)(data)
	if errors != nil {
		t.Error(
			generateErrorMessage(errors, "nil"),
		)
	}
}

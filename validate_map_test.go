package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

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

	user := User{
		Name:  "Name",
		Age:   15,
		Email: "exampleemail@gmail.com",
	}

	errors = ValidateMap(validations)(user)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[error message(s)]"),
		)
	}

	user.Email = "emailexample"

	errors = ValidateMap(validations)(user)
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

	user.Age = 18
	user.Email = "emailexample@gmail.com"
	if errors != nil {
		t.Error(
			generateErrorMessage(errors, "nil"),
		)
	}
}

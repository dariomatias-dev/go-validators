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

	user := User{
		Name:  "Name",
		Age:   18,
		Email: "exampleemail@gmail.com",
	}

	data := map[string]interface{}{
		"name":  "Name",
		"age":   18,
		"email": "emailexample@gmail.com",
	}

	/// - Successes
	// Test 1
	errors := ValidateMap(validations)(user)
	if errors != nil {
		t.Error(
			generateErrorMessage(errors, "nil"),
		)
	}

	// Test 2
	errors = ValidateMap(validations)(data)
	if errors != nil {
		t.Error(
			generateErrorMessage(errors, "nil"),
		)
	}

	/// - Errors
	// Test 1
	data["age"] = 15
	errors = ValidateMap(validations)(data)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[ error message(s) ]"),
		)
	}

	// Test 2
	data["age"] = 18
	data["email"] = "emailexample"
	errors = ValidateMap(validations)(data)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[ error message(s) ]"),
		)
	}

	// Test 3
	user.Age = 15
	errors = ValidateMap(validations)(user)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[ error message(s) ]"),
		)
	}

	// Test 4
	user.Age = 18
	user.Email = "emailexample"
	errors = ValidateMap(validations)(user)
	if errors == nil {
		t.Error(
			generateErrorMessage(errors, "[ error message(s) ]"),
		)
	}
}

package validators

// Package validators provides configurations for the IsArray and IsNullArray validators.
type Array struct {
	// AllowEmpty indicates whether the field can be empty or not.
	AllowEmpty bool
	// AllowEmptyErrorMessage is a custom error message to be displayed when the field is empty and not allowed.
	AllowEmptyErrorMessage string

	// MinLength specifies the minimum number of elements in the array.
	MinLength int
	// MinLengthErrorMessage is a custom error message to be displayed when the array has fewer elements than allowed.
	MinLengthErrorMessage string

	// MaxLength specifies the maximum number of elements in the array.
	MaxLength int
	// MaxLengthErrorMessage is a custom error message to be displayed when the array has more elements than allowed.
	MaxLengthErrorMessage string
}

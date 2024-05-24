package validators

// Represents a validation map where the keys are the names of the fields to be validated, with each field having slices of associated validators.
type Validators map[string][]Validator

// A function that takes a value to be validated and returns an error message along with a stop indicator for further validations.
// If the validation is successful, the error message should be nil and the stop indicator should be false. Otherwise, an error message is returned along with a boolean indicating whether to stop subsequent validations or not.
type Validator func(value any) (
	err error,
	abortValidation bool,
)

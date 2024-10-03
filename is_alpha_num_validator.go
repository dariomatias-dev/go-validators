package validators

var (
	alphaNumRegex = `^[a-zA-ZÀ-ÖØ-öø-ÿ0-9]+$`
)

// Checks if the value contains only letters.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "abcABC"
//	v.IsAlpha()(value)     // Output: nil, false
//
//	value = "abcABC "
//	v.IsAlpha()(value)     // Output: [error message], false
//
//	value = "abcABC0123!@"
//	v.IsAlpha()(value)     // Output: [error message], false
func IsAlphaNum(
	errorMessage ...string,
) Validator {
	message := "Contains characters that are not letters and numbers."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return Regex(alphaNumRegex, message)
}

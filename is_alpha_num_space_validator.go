package validators

var (
	alphaNumSpaceRegex = `^[a-zA-ZÀ-ÖØ-öø-ÿ0-9\s]+$`
)

// Checks if the value contains only letters, numbers and spaces.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "abcABC"
//	v.IsAlphaNumSpace()(value) // Output: nil, false
//
//	value = "abcABC "
//	v.IsAlphaNumSpace()(value) // Output: nil, false
//
//	value = "abcABC012!@"
//	v.IsAlphaNumSpace()(value) // Output: [error message], false
func IsAlphaNumSpace(
	errorMessage ...string,
) Validator {
	message := "Contains characters other than letters and spaces."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return Regex(alphaNumSpaceRegex, message)
}

package validators

var (
	alphaSpaceRegex = `^[a-zA-ZÀ-ÖØ-öø-ÿ\s]+$`
)

// Checks if the value contains only letters and spaces.
//
// Configuration parameters:
//   - errorMessage (string): custom error message (optional).
//
// Input value (string): value to be validated.
//
// Usage examples:
//
//	value := "abcABC"
//	v.IsAlphaSpace()(value) // Output: nil, false
//
//	value = "abcABC "
//	v.IsAlphaSpace()(value) // Output: nil, false
//
//	value = "abcABC0123!@"
//	v.IsAlphaSpace()(value) // Output: [error message], false
func IsAlphaSpace(
	errorMessage ...string,
) Validator {
	message := "Contains characters other than letters and spaces."
	if len(errorMessage) != 0 {
		message = errorMessage[0]
	}

	return Regex(alphaSpaceRegex, message)
}

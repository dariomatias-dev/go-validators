package validators

type Array struct {
	AllowEmpty             bool
	AllowEmptyErrorMessage string
	MinLength              int
	MinLengthErrorMessage  string
	MaxLength              int
	MaxLengthErrorMessage  string
}

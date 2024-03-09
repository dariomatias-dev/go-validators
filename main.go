package validators

type Validators map[string][]Validator
type Validator func(value interface{}) (
	errorMessage *string,
	stopLoop bool,
)

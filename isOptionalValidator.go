package validators

func IsOptional() Validator {
	return func(value interface{}) (*string, bool) {
		if value == nil {
			return nil, true
		}

		return nil, false
	}
}

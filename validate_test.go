package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name" validates:"minLength=5;maxLength=15"`
	Age   int32  `json:"age" validates:"isRequired;isInt;min=5"`
	Email string `json:"email" validates:"minLength=5"`
}

func TestValidate(t *testing.T) {
	user := User{
		Name:  "Dário Matias",
		Age:   19,
		Email: "matiasdario75@gmail.com",
	}

	// user := &User{}

	data := `{
		"name":  "Dário Matias",
		"age":   19,
		"email": "matiasdario75@gmail.com"
	}`

	fmt.Println(
		Validate(&user, &data),
	)
}

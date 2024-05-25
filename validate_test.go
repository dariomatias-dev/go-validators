package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name" validates:"isNullString"`
	Age   int64  `json:"age" validates:"isRequired;isInt;min=5"`
	Email string `json:"email" validates:"minLength=5"`
}

func TestValidate(t *testing.T) {
	// user := User{
	// 	Name:  "Dário Matias",
	// 	Age:   19,
	// 	Email: "matiasdario75@gmail.com",
	// }

	user := &User{}

	data := `{
		"name":  0,
		"age": 19,
		"email": "matiasdario75@gmail.com"
	}`

	fmt.Println(
		Validate(user, &data),
	)
}

package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string `json:"name" validates:"isOptional;isNullString"`
	Age     int64  `json:"age" validates:"isRequired;isInt;min=5"`
	Email   string `json:"email" validates:"minLength=5"`
	Product Product
}

type Product struct {
	Id   string `json:"id" validates:"minLength=2"`
	Name string `json:"name" validates:"minLength=2"`
}

func TestValidate(t *testing.T) {
	user := User{
		Name:  "Dário Matias",
		Age:   19,
		Email: "matiasdario75@gmail.com",
		Product: Product{
			Id:   "1234",
			Name: "Smartphone",
		},
	}

	// user := &User{}

	// data := `{
	// 	"name": false,
	// 	"age": 19,
	// 	"email": "matiasdario75@gmail.com"
	// }`

	fmt.Println(
		Validate(user, nil),
	)
}

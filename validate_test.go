package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string  `json:"name" validates:"isOptional;isNullString"`
	Age     int64   `json:"age" validates:"isRequired;isInt;min=5"`
	Email   string  `json:"email" validates:"minLength=5"`
	Product Product `json:"product"`
}

type Product struct {
	Id   string `json:"id" validates:"minLength=5"`
	Name string `json:"name" validates:"minLength=5"`
}

func TestValidate(t *testing.T) {
	user := User{
		Name:  "Dário Matias",
		Age:   19,
		Email: "matiasdario75@gmail.com",
		Product: Product{
			Id: "56789",
			Name: "Smartphone",
		},
	}

	// user := &User{}

	// data := `{
	// 	"name": "Dário Matias",
	// 	"age": 19,
	// 	"email": "matiasdario75@gmail.com",
	// 	"product": {
	// 		"id": "56789",
	// 		"name": "Smartphone"
	// 	}
	// }`

	fmt.Println(
		Validate(user, nil),
	)
	// fmt.Println(
	// 	Validate(user, &data),
	// )
}

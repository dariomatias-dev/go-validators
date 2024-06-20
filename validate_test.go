package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name" validates:"isRequired;isString;minLength=5;maxLength=50"`
	Age   int64  `json:"age" validates:"isRequired;isInt;min=18;max=120"`
	Email string `json:"email" validates:"isRequired"`
	// Product  Product   `json:"product" validates:"isRequired"`
	// Products []Product `json:"products" validates:"isRequired"`
	Ids *[]int `json:"ids" validates:"minLength=2;isNullArray=Error Message;isRequired=Error message;min=2"`
}

type Product struct {
	Id   string `json:"id" validates:"isRequired"`
	Name string `json:"name" validates:"isRequired;isString;minLength=5;maxLength=50"`
}

func TestValidate(t *testing.T) {
	user := &User{}

	data := `{
		"name": "Dário Matias",
		"age": 19,
		"email": "matiasdario75@gmail.com",
		"ids": [0, 7, 3, 10]
	}`

	fmt.Println(
		Validate(user, &data),
	)
}

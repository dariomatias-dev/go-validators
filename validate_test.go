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
	Ids []int `json:"ids" validates:"isArray=Error Message;isRequired=Error message;min=2"`
}

type Product struct {
	Id   string `json:"id" validates:"isRequired"`
	Name string `json:"name" validates:"isRequired;isString;minLength=5;maxLength=50"`
}

func TestValidate(t *testing.T) {
	user := User{
		Name:  "Dário Matias",
		Age:   19,
		Email: "matiasdario75@gmail.com",
		Ids:   []int{0, 4, 1, 5},
		// Product: Product{
		// 	Id:   "56789",
		// 	Name: "Smartphone",
		// },
		// Products: []Product{
		// 	{
		// 		Id:   "6547346739",
		// 		Name: "Not",
		// 	},
		// 	{
		// 		Id:   "0275949305",
		// 		Name: "Smartphone",
		// 	},
		// },
	}

	// user := &User{}

	// data := `{
	// 	"name": "Dário Matias",
	// 	"age": 19,
	// 	"email": "matiasdario75@gmail.com",
	// 	"products": [
	// 		{
	// 			"id": "6547346739",
	// 			"name": "Notebook"
	// 		}
	// 	]
	// }`

	// fmt.Println(
	// 	Validate(user, &data),
	// )
	fmt.Println(
		Validate(user, nil),
	)
}

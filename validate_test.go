package validators

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name" validates:"required;minLength=5;maxLength=15"`
	Age   int32    `json:"age" validates:"min=5"`
	Email string `json:"email" validates:"minLength=5"`
}

func TestValidate(t *testing.T) {
	user := User{
		Name: "Dário Matias",
		Age: 19,
		Email: "matiasdario75@gmail.com",
	}

	fmt.Println(Validate(user))
}

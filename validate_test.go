package validators

import "testing"

type User struct {
	Name  string `json:"name" validates:"required;minLength=5;maxLength=10"`
	Age   int    `json:"age" validates:"min=5,Mensagem de erro"`
	Email string `json:"email" validates:"maxLength=5;minLength=5"`
}

func TestValidate(t *testing.T) {
	user := User{
		Name: "Dário Matias",
		Age: 19,
		Email: "matiasdario75@gmail.com",
	}

	Validate(user)
}

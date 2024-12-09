package main

import "fmt"

type User struct {
	Name string
	Email string
	Address Address // composição
}

type Address struct {
	Street string
	City string
}

func (u *User) Apresentar() (string) {
	return fmt.Sprintf("Olá, meu nome é %s e meu e-mail é %s e minha rua é %s e minha cidade é %s", u.Name, u.Email, u.Address.Street, u.Address.City)
}

func main() {

	user := User{
		Name: "Allison",
		Email: "gierdiaz@hotmail",
		Address: Address{
			Street: "Rua Olinto da Gama Botelho",
			City: "Rio de Janeiro",
		},
	}

	println(user.Apresentar())
	// println(user.Name, user.Email, user.Address.Street, user.Address.City)
}
package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "Allison",
		"email": "gierdiaz@hotmail",
	}

	fmt.Println(user)

	person := map[string]map[string]string{
		"name": {
			"first": "Allison",
			"last": "Gierdiaz",
		},
		"curso": {
			"nome": "Análise e Desenvolvimento de Sistemas",
			"specialization": "Sistemas Distrubuidos e Arquitetura de Microserviços",
		},
		
	}

	fmt.Println(person)

	person["signo"] = map[string]string{
		"signo": "Touro",
	}

	fmt.Println(person)
}
package main

import "fmt"

func clousure() func() {
	text := "Dentro da função closure"

	funcao := func() { // função que retorna uma função
		fmt.Println(text)
	}

	return funcao
}

func main() {
  texto := "Dentro da função main"
  fmt.Println(texto)

  funcaoNova := clousure()
  funcaoNova()
}
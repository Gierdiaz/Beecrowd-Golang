package main

// go build -> cria o binário
// ./modulo -> executa o binário

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do arquivo main")

	// chamando o pacote e acessandogo o método ou a função
	aux := auxiliar.Escrever()
	fmt.Println(aux)

	auxiliar.Lendo()

	// Verificando se o email é válido ou não
	email := "gierdiaz@hotmail.com"
	err := checkmail.ValidateFormat(email)
	if err != nil {
		fmt.Printf("O e-mail '%s' é inválido: %s\n", email, err)
	} else {
		fmt.Printf("O e-mail '%s' é válido.\n", email)
	}
}
package auxiliar

import "fmt"

// A visibilidade dos métodos e funções é se começa com letra maiúscula (public) ou minúscula (private)

// retornando uma string
func Escrever() string {
	return "escrevendo do auxiliar"
}

// não retorna nada
func Lendo() {
	fmt.Println("lendo do auxiliar")
}
package main

import (
	"fmt"
	"introducao-testes/endereco"
)

func main() {
	TipoEndereco := endereco.TipoEndereco("Avenida Paulista")
	fmt.Println(TipoEndereco)
}
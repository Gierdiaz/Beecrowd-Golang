package main

import "fmt"

func main() {
	// Referencia de memória
	var variavel3 int = 27
	var ponteiro *int = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // Desreferenciação
}
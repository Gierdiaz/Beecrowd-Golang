package main

import "fmt"

func main(){

	var a, b, soma int

	fmt.Scan(&a, &b)
	soma = soma_simples(a, b)
	fmt.Println("SOMA =", soma)
}

func soma_simples(a, b int) int {
	return a + b
}
package main

import "fmt"

func main() {

	var produto, a, b int

	fmt.Scan(&a, &b)

	produto = produto_simples(a, b)
	fmt.Println("PROD =", produto)
}

func produto_simples(a, b int) int {
	return a * b
}
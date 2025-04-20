package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	multiplos(a, b)
}

func multiplos(a, b int) {
	if a % b == 0 || b % a == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}
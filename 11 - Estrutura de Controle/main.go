package main

import "fmt"

func main() {
	number := 10

	if number > 15 {
		fmt.Println("Maior que 15")
	} else if number < 15 {
		fmt.Println("Menor que 15")
	} else {
		fmt.Println("Igual a 15")
	}

	if outroNumero := number; outroNumero > 0 {
		fmt.Println("Maior que 0")
	}
}
package main

import "fmt"

func soma(numeros ...int) int { // ...int = slice por isso podemos usar o for range
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)
}
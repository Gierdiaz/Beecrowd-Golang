package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	contador := pares_entre_cinco_numeros([]int{a, b, c, d, e})
	fmt.Printf("%d valores pares\n", contador)
}

func pares_entre_cinco_numeros(valores []int) int {
	contador := 0

	for _, numero := range valores {
		if (numero % 2 == 0) {
			contador++
		}
	}

	return contador
}
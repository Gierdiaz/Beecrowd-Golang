package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	pares, impares, positivos, negativos := pares_impares_positivos_negativos([]int{a, b, c, d, e})

	fmt.Printf("%d valor(es) par(es)\n", pares)
	fmt.Printf("%d valor(es) impar(es)\n", impares)
	fmt.Printf("%d valor(es) positivo(s)\n", positivos)
	fmt.Printf("%d valor(es) negativo(s)\n", negativos)
}

func pares_impares_positivos_negativos(valores []int) (int, int, int, int) {
	pares := 0
	impares := 0
	positivos := 0
	negativos := 0

	for _, numero := range valores {
		if numero % 2 == 0 {
			pares++
		} else {
			impares++
		}

		if numero > 0 {
			positivos++
		} else if numero < 0 {
			negativos++
		}
	}

	return pares, impares, positivos, negativos 
}
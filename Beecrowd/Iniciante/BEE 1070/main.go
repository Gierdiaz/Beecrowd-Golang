package main

import "fmt"

// func main() {
// 	var x int
// 	fmt.Scan(&x)

// 	seis_numeros_impares(x)
// }

// func seis_numeros_impares(valor int) int{
// 	contador := 0

// 	for iteracao := valor; contador < 6; iteracao ++ {
// 		if (iteracao % 2 != 0) {
// 			fmt.Println(iteracao)
// 			contador++
// 		}
// 	}

// 	return contador
// }

// Uma segunda abordagem mais "Go-like"

func main() {
	var valor int
	fmt.Scan(&valor)

	impares := gerarImpares(valor)

	for _, numero := range impares {
		fmt.Println(numero)
	}
}

func gerarImpares(valor int) []int {
	impares := make([]int, 0, 6) // tamanho 0 e capacidade 6

	for iteracao := valor; len(impares) < 6; iteracao++ {
		if (iteracao % 2 != 0) {
			impares = append(impares, iteracao)
		}
	}

	return impares
}
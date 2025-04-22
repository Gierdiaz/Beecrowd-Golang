package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	pares := quadrado_pares(n)
	mostra_quadrados(pares)
}

func mostra_quadrados(pares []int) {
	for _, par := range pares  {
		fmt.Printf("%d^2 = %d\n", par, par*par)
	}
}

func quadrado_pares(n int) []int {
	pares := make([]int, 0)

	for iteracao := 1; iteracao <= n; iteracao++ {
		if iteracao % 2 == 0 {
			pares = append(pares, iteracao)
		}
	}

	return pares
}
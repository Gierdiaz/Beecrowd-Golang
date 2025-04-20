package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	sort_simples(a, b, c)

	fmt.Println()
	numeros_em_ordem_exibidos([]float64{a, b, c})

}

func sort_simples(n1, n2, n3 float64) []float64 {
	numeros := []float64{n1, n2, n3}
	sort.Float64s(numeros)

	for _, numero := range numeros {
		fmt.Println(int(numero))
	}
	
	return numeros
}

func numeros_em_ordem_exibidos(numeros []float64) {
	for _, numero := range numeros {
		fmt.Println(int(numero))
	}
}
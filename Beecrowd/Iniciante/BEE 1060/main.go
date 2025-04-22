package main

import "fmt"

func main() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	contador := numeros_positivos([]float64{a, b, c, d, e, f})
	fmt.Printf("%d valores positivos\n", contador)
}

func numeros_positivos(valores []float64) int {
	contador := 0

	for _, numeros := range valores {
		if (numeros > 0) {
			contador++
		}
	}

	return contador
}
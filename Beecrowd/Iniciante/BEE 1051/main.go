package main

import "fmt"

func main() {
	var valor float64

	fmt.Scan(&valor)

	valorImposto, isento  := calculo_imposto_de_renda(valor)
	
	if isento {
		fmt.Println("Isento")
	} else {
		fmt.Printf("R$ %.2f\n", valorImposto)
	}
}

func calculo_imposto_de_renda(valor float64) (float64, bool) {

	if valor <= 2000.00 {
		return 0.0, true
	}

	var valorImposto float64

	switch {
		case valor <= 3000.00:
			valorImposto = (valor - 2000.00) * 0.08
		case valor <= 4500.00:
			valorImposto = (1000.00 * 0.08) + (valor - 3000.00) * 0.18
		default:
			valorImposto = (1000.00 * 0.08) + (1500.00 * 0.18) + (valor - 4500.00) * 0.28
	}

	return valorImposto, false

}
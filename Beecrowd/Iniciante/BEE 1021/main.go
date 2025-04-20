package main

import (
	"fmt"
	"math"
)

func main() {
	var valor float64
	fmt.Scan(&valor)
	decomporValor(valor)
}

func decomporValor(valor float64) {
	notas := []float64{100.00, 50.00, 20.00, 10.00, 5.00, 2.00}
	moedas := []float64{1.00, 0.50, 0.25, 0.10, 0.05, 0.01}

	fmt.Println("NOTAS:")
	valor = calcularNotas(valor, notas)

	fmt.Println("MOEDAS:")
	calcularMoedas(valor, moedas)

}

func calcularNotas(valor float64, notas []float64) float64 {
	for _, nota := range notas {
		quantidade := int(valor / nota)
		fmt.Printf("%d nota(s) de R$ %.2f\n", quantidade, nota)
		valor = math.Mod(valor, nota)
	}

	return valor
}

func calcularMoedas(valor float64, moedas []float64) {
	for _, moeda := range moedas {
		quantidade := int(math.Round(valor * 100) / math.Round(moeda * 100))
		fmt.Printf("%d moeda(s) de R$ %.2f\n", quantidade, moeda)
		valor = math.Mod(math.Round(valor * 100), math.Round(moeda * 100)) / 100
	}
}
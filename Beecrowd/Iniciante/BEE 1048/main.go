package main

import "fmt"

func main() {
	var valor float64
	fmt.Scan(&valor)
	calcular_aumento_salario(valor)
}

func calcular_aumento_salario(valor float64) {

	var percentual float64

	switch {
		case valor >= 0 && valor <= 400.00:
			percentual = 0.15
		case valor >= 400.01 && valor <= 800.00:
			percentual = 0.12
		case valor >= 800.01 && valor <= 1200.00:
			percentual = 0.10
		case valor >= 1200.01 && valor <= 2000.00:
			percentual = 0.07
		case valor > 2000.00:
			percentual = 0.04
	}

	aumento := valor * percentual
	novoSalario := valor + aumento

	fmt.Printf("Novo salario: %.2f\n", novoSalario)
	fmt.Printf("Reajuste ganho: %.2f\n", aumento)
	fmt.Printf("Em percentual: %.0f %%\n", percentual*100)
}
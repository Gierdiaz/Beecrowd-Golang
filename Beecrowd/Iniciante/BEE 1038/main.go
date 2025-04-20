package main

import (
	"fmt"
)

func main() {
	var codigo, quantidade int
	fmt.Scan(&codigo, &quantidade)

	pedido := lanche(codigo, quantidade)
	fmt.Printf("Total: R$ %.2f\n", pedido)
}

func lanche(codigo, quantidade int) float64 {
	produto := map[int]map[string]float64{
		1: {"Cachorro Quente": 4.00},
		2: {"X-Salada": 4.50},
		3: {"X-Bacon": 5.00},
		4: {"Torrada simples": 2.00},
		5: {"Refrigerante": 1.50},
	}


	if item, existe := produto[codigo]; existe {
		for _, preco := range item {
			return float64(quantidade) * preco
		}
	}

	return 0.0 // Caso o código do produto nao seja encontrado
}
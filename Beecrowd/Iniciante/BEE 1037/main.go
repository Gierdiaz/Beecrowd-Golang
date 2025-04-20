package main

import "fmt"

func main() {
	var valor float64

	fmt.Scan(&valor)
	resultado := intervalo(valor)
	fmt.Println(resultado)
}

func intervalo(valor float64) string {
	intervalos := [][]float64{{0, 25}, {25, 50}, {50, 75}, {75, 100}}

	if valor >= intervalos[0][0] && valor <= intervalos[0][1] {
		return "Intervalo [0,25]"
	} else if valor > intervalos[1][0] && valor <= intervalos[1][1] {
		return "Intervalo (25,50]"
	} else if valor > intervalos[2][0] && valor <= intervalos[2][1] {
		return "Intervalo (50,75]"
	} else if valor > intervalos[3][0] && valor <= intervalos[3][1] {
		return "Intervalo (75,100]"
	} else {
		return "Fora de intervalo"
	}

	// Uma segunda forma mais simples de entender do que duplo slice
	// func verificarIntervalo(valor float64) string {
	// 	if valor >= 0 && valor <= 25 {
	// 		return "Intervalo [0,25]"
	// 	} else if valor > 25 && valor <= 50 {
	// 		return "Intervalo (25,50]"
	// 	} else if valor > 50 && valor <= 75 {
	// 		return "Intervalo (50,75]"
	// 	} else if valor > 75 && valor <= 100 {
	// 		return "Intervalo (75,100]"
	// 	} else {
	// 		return "Fora de intervalo"
	// 	}
	// }
}




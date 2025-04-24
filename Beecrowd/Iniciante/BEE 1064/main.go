package main

import "fmt"

func main() {
	var valores [6]float64

	for i := 0; i < 6; i++ {
		fmt.Scan(&valores[i])
	}

	quantidadePositivos, media := positivos_e_media(valores[:])

	fmt.Printf("%d valores positivos\n", quantidadePositivos)
	fmt.Printf("%.1f\n", media)
}

func positivos_e_media(valores []float64) (int, float64) {
	var soma float64
	var count int

	for _, valor := range valores {
		if valor > 0 {
			soma = soma + valor
			count++
		}
	}

	if count > 0 {
		return count, soma / float64(count)
	}

	return count, 0.0
}
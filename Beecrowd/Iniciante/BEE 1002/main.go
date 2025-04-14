package main

import "fmt"

func main() {
	var raio float64

	fmt.Print("Digite o valor do raio: ")
	fmt.Scan(&raio)

	fmt.Printf("A=%.4f\n", area_do_circulo(raio))
}

func area_do_circulo(raio float64) float64 {
	return 3.14159 * raio * raio
}
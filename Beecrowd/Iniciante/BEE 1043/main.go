package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	calcular_forma_triangulo(a, b, c)
}

func calcular_forma_triangulo(a, b, c float64) float64 {

	// a < (b + c)
	// b < (a + c)
	// c < (a + b)

	if a < (b + c) && b < (a + c) && c < (a + b) {
		perimetro :=  perimetro_do_triangulo(a, b, c)
		fmt.Printf("Perimetro = %.1f\n", perimetro)
		return perimetro
		
	} else {
		trapezio := area_do_trapezio(a, b, c)
		fmt.Printf("Area = %.1f\n", trapezio)
		return trapezio
	}
}

// Perímetro de um triangulo: p = l + l + l
func perimetro_do_triangulo(a, b, c float64) float64 {
	return (a + b + c)
}

// Area do trapézio (B + b) * h / 2
func area_do_trapezio(a, b, c float64) float64 {
	return (a + b) * c / 2
}
package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z float64
	fmt.Scan(&x, &y, &z)

	valores := []float64{x, y, z}
	sort.Sort(sort.Reverse(sort.Float64Slice(valores)))

	a, b, c := valores[0], valores[1], valores[2]

	tipo_de_triangulos(a, b, c)
}

func tipo_de_triangulos(a, b, c float64) {
	if (a >= b + c) {
		fmt.Println("NAO FORMA TRIANGULO")
		return
	}

	// Teste de ângulo
	if a * a == b * b + c * c {
		fmt.Println("TRIANGULO RETANGULO")
	} else if (a * a > b * b + c * c) {
		fmt.Println("TRIANGULO OBTUSANGULO")
	} else if (a * a < b * b + c * c) {
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	// Teste de lados
	if (a == b && b == c) {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if (a == b || b == c || a == c) {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}

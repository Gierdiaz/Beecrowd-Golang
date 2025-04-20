package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	coordenadas(a, b)
}

func coordenadas(x, y float64) {
	switch {
		case x < 0 && y > 0:
			fmt.Println("Q2")
		case x > 0 && y > 0:
			fmt.Println("Q1")
		case x > 0 && y < 0:
			fmt.Println("Q4")
		case x < 0 && y < 0:
			fmt.Println("Q3")
		case x == 0 && y == 0:
			fmt.Println("Origem")
		case x == 0:
			fmt.Println("Eixo Y")
		case y == 0:
			fmt.Println("Eixo X")
		default:
			fmt.Println("Indefinido")
	}
}
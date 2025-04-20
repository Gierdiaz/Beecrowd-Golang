package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	teste_selecao_1(a, b, c, d)
}

func soma(x, y int) int {
	return (x + y)
}

func teste_selecao_1(a, b, c, d int) {
	if b > c && d > a && soma(c, d) > soma(a, b) && c > 0 && d > 0 && a % 2 == 0 {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}
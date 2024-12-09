package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	resto := 10 % 3

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	// Atribuição
	numero := 10
	numero += 10
	numero -= 10
	numero *= 10
	numero /= 10

	fmt.Println(numero)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Unários
	x := 10
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	// Ternários
	numero = 10
	resultado := "Maior que 10"
	if numero > 10 {
		resultado = "Menor que 10"
	} else {
		resultado = "Maior que 10"
	}
	fmt.Println(resultado)
}
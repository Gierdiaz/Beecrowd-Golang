package main

import (
	"fmt"
)

// func main() {
// 	var x, y int

// 	fmt.Scan(&x, &y)
// 	mostrar_soma_dos_valores_entre_impares_consecutivos(x, y)
// }



// func mostrar_soma_dos_valores_entre_impares_consecutivos(x, y int) int {
// 	if x > y {
// 		x, y = y, x // troca os valores se estiverem na ordem errada
// 	}

// 	soma := 0

// 	for iteracao := x + 1; iteracao < y; iteracao++ {
// 		if (iteracao % 2 != 0) {
// 			soma = soma + 1
// 		}
// 	}

// 	return soma
// }

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	resultado := somaDosImparesEntre(x, y)
	fmt.Println(resultado)
}

func somaDosImparesEntre(x, y int) int {
	soma := 0

	// determina o menor e maior valor entre os dois
	inicio := min(x, y) + 1
	fim := max(x, y)

	for i := inicio; i < fim; i++ {
		if i%2 != 0 {
			soma += i
		}
	}

	return soma
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
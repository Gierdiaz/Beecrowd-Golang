package main

import "fmt"

// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	calcular_numeros_impares(x)

// }

// func calcular_numeros_impares(x int) {

// 	for iteracao := 1; iteracao <= x; iteracao++ {
// 		if (iteracao % 2 != 0) {
// 			fmt.Println(iteracao)
// 		}
// 	}
// }

func main() {
	var x int
	fmt.Scan(&x)
	
	for _, impar := range imparesMenoresQue(x) {
		fmt.Println(impar)
	}
}

func imparesMenoresQue(x int) []int {
	impares := []int{}
	for iteracao := 1; iteracao <= x; iteracao++ {
		if iteracao % 2 != 0 {
			impares = append(impares, iteracao)
		}
	}
	return impares
}
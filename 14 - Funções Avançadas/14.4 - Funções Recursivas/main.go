package main

import "fmt"

func somaSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	return slice[0] + somaSlice(slice[1:])
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	for posicao := uint(3); posicao < 10; posicao++ {
		fmt.Println(fibonacci(posicao))
	}

	slice := []int{1, 2, 3, 4, 5}

	result := somaSlice(slice)

	fmt.Printf("A soma dos elementos é: %d\n", result)
}
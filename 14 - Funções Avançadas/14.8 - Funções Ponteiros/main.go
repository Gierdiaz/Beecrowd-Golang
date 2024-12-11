package main

import "fmt"

func inverterSinalComPonteiro(num *int ) int{ // passando uma referência de memória
	*num = *num * - 1
	return *num
}

func inverterSinalSemPonteiro(num int) int { // passando um parâmetro por cópia de um valor
	return num * - 1
}

func main() {
	numero := 20
	numeroInvertidoComPonteiro := inverterSinalComPonteiro(&numero)
	fmt.Println(numeroInvertidoComPonteiro)
	fmt.Println(numero)

	novoNumero := 40
	numeroInvertidoSemPonteiro := inverterSinalSemPonteiro(novoNumero)
	fmt.Println(numeroInvertidoSemPonteiro)
	fmt.Println(novoNumero)
}
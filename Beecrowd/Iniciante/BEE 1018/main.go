package main

import "fmt"

func main() {
	var valor int

	fmt.Scan(&valor)

	fmt.Println(valor)

	cedulas := calcule_cedulas(valor)

	notas := []int{100, 50, 20, 10, 5, 2, 1}

	for _, nota := range notas {
		quantidade := cedulas[nota]
		fmt.Printf("%d nota(s) de R$ %d,00\n", quantidade, nota)
	}


}

func calcule_cedulas(valor int) map[int]int {
	notas := []int{100, 50, 20, 10, 5, 2, 1}
	contagemNotas := make(map[int]int) // criando um map vazio
	valorRestante := valor

	for _, nota := range notas {
		quantidade := valorRestante / nota
		if quantidade > 0 {
			contagemNotas[nota] = quantidade
			valorRestante -= quantidade * nota
		}
	}

	return contagemNotas
}
package main

import "fmt"

func recuperarExecução() {
	if r := recover(); r != nil {
		fmt.Println("Tentando recuperar a execução.")
	}
	
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecução()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6!")
}

func main() {
	result := alunoEstaAprovado(6, 6)
	fmt.Println("Pós execução!")
	fmt.Println(result)
}
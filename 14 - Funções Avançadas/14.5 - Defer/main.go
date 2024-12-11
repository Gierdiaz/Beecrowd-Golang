package main

import "fmt"

func função1() {
	fmt.Println("Executando a função 1")
}

func função2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("A média calculada será retornada.")
	fmt.Println("Entrando na função para verifica se o aluno está aprovado ou não.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer função1() // Executa a função 1 por último

	função2() // Primeiro executa a função 2, depois executa a função 1

	restult := alunoEstaAprovado(7, 8)
	fmt.Println(restult)
}
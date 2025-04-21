package main

import "fmt"

func main() {
	var horaInicial, horaFinal int
	fmt.Scan(&horaInicial, &horaFinal)
	calcular_duracao_do_jogo(horaInicial, horaFinal)
}

func calcular_duracao_do_jogo(horaInicial, horaFinal int) {
	if (horaInicial == horaFinal) {
		fmt.Println("O JOGO DUROU 24 HORA(S)")
	} else if (horaInicial < horaFinal) {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", (horaFinal - horaInicial))
	} else {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", (24 - horaInicial) + horaFinal)
	}
}
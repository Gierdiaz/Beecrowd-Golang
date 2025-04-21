package main

import (
	"fmt"
)

func main() {
	var horaInicial, minutoInicial, horaFinal, minutoFinal int
	fmt.Scan(&horaInicial, &minutoInicial, &horaFinal, &minutoFinal)

	horas, minutos := calcularTempoDoJogo(horaInicial, minutoInicial, horaFinal, minutoFinal)

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", horas, minutos)
}

func calcularTempoDoJogo(horaInicial, minutoInicial, horaFinal, minutoFinal int) (int, int) {
	inicial := horaInicial * 60 + minutoInicial
	final := horaFinal * 60 + minutoFinal

	// Verifica se o jogo começa e termina no mesmo horário
	if horaInicial == horaFinal && minutoInicial == minutoFinal {
		return 24, 0
	}

	// Caso o final seja antes do início (passou da meia-noite)
	if (final < inicial) {
		final = final + 24 * 60
	}

	duracao := final - inicial
	horas := duracao / 60
	minutos := duracao % 60

	return horas, minutos
}
package main

// Nao gostei desse exercicio.

import "fmt"

func main() {
	var dia1, dia2 int
	var h1, m1, s1 int
	var h2, m2, s2 int

	// Entrada
	fmt.Scanf("Dia %d", &dia1)
	fmt.Scanf("%d : %d : %d", &h1, &m1, &s1)

	fmt.Scanf("Dia %d", &dia2)
	fmt.Scanf("%d : %d : %d", &h2, &m2, &s2)

	// Chamada da função
	duracaoDoEvento(dia1, h1, m1, s1, dia2, h2, m2, s2)
}

// Função que calcula e imprime a duração do evento
func duracaoDoEvento(d1, h1, m1, s1, d2, h2, m2, s2 int) {
	// Converter tudo para segundos
	segundosInicial := d1*86400 + h1*3600 + m1*60 + s1
	segundosFinal := d2*86400 + h2*3600 + m2*60 + s2

	// Diferença total em segundos
	total := segundosFinal - segundosInicial

	// Quebrar em dias, horas, minutos e segundos
	dias := total / 86400
	total %= 86400

	horas := total / 3600
	total %= 3600

	minutos := total / 60
	segundos := total % 60

	// Saída formatada
	fmt.Printf("%d dia(s)\n", dias)
	fmt.Printf("%d hora(s)\n", horas)
	fmt.Printf("%d minuto(s)\n", minutos)
	fmt.Printf("%d segundo(s)\n", segundos)
}

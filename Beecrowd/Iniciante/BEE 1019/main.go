package main

import (
	"fmt"
	"time"
)

func main() {
	var valor int

	fmt.Scan(&valor)

	h, m, s := conversao_de_tempo(valor)

	fmt.Printf("%d:%d:%d\n", h, m, s)
}

func conversao_de_tempo(segundos int) (hora, minuto, segundo int) {
	duracao := time.Duration(segundos) * time.Second
	hora = int(duracao.Hours())
	minuto = int(duracao.Minutes()) % 60
	segundo = int(duracao.Seconds()) % 60
	return
}

package main

import "fmt"

func main() {
	var valor int

	fmt.Scan(&valor)

	ano, mes, dia := idade_em_dias(valor)
	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", ano, mes, dia)
}

func idade_em_dias(dias int) (ano, mes, dia int) {
	ano = (dias / 365)
	mes = (dias % 365) / 30 
	dia = (dias % 365) % 30
	return
}
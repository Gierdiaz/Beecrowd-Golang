package main

func main() {
	resultado := Somar(10, 20)
	println(resultado)

	var f = func (txt string) string {
		return txt
	}
	f("Funcão anonima")

	_, subtracao := Calculo(10, 20)
	println("Subtração: ", subtracao)
}

func Somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func Calculo(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
package main

import "fmt"

func diaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func diaSemana2(dia int) string {
	switch {
		case dia == 1:
			return "Domingo"
			fallthrough // joga pra dentro da próxima condição
		case dia == 2:
			return "Segunda-Feira"
		case dia == 3:
			return "Terça-Feira"
		case dia == 4:
			return "Quarta-Feira"
		case dia == 5:
			return "Quinta-Feira"
		case dia == 6:
			return "Sexta-Feira"
		case dia == 7:
			return "Sabado"
		default:
			return "Dia inválido"
	}
}

func main() {

	fmt.Println(diaSemana(0))

	fmt.Println(diaSemana2(0))
}
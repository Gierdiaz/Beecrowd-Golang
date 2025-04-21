package main

import "fmt"

func main() {
	var ddd int
	fmt.Scan(&ddd)
	discagem_direta_a_distancia(ddd)
}

func discagem_direta_a_distancia(ddd int) {
	codigo := map[int]string{
		61: "Brasilia",
		71: "Salvador",
		11: "Sao Paulo",
		21: "Rio de Janeiro",
		32: "Juiz de Fora",
		19: "Campinas",
		27: "Vitoria",
		31: "Belo Horizonte",
	}

	// if cidade, ok := codigo[ddd]; ok {
	// 	fmt.Println(cidade)
	// } else {
	// 	fmt.Println("DDD nao cadastrado")
	// }

	for chave, cidade := range codigo {
		if chave == ddd {
			fmt.Println(cidade)
			return
		}
	}

	fmt.Println("DDD nao cadastrado")
}
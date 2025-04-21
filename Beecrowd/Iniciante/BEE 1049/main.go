package main

import "fmt"

func main() {
	var opcao1, opcao2, opcao3 string
	fmt.Scan(&opcao1, &opcao2, &opcao3)

	catalogo_animal(opcao1, opcao2, opcao3)
}

func catalogo_animal(opcao1, opcao2, opcao3 string) {
	catalogo := map[string]map[string]map[string]string{
		"vertebrado": {
			"ave": {
				"carnivoro": "aguia",
				"onivoro":   "pomba",
			},
			"mamifero": {
				"onivoro":   "homem",
				"herbivoro": "vaca",
			},
		},
		"invertebrado": {
			"inseto": {
				"hematofago": "pulga",
				"herbivoro":  "lagarta",
			},
			"anelideo": {
				"hematofago": "sanguessuga",
				"onivoro":    "minhoca",
			},
		},
	}

	if resultado, ok := catalogo[opcao1][opcao2][opcao3]; ok {
		fmt.Println(resultado)
	} else {
		fmt.Println("Animal não encontrado")
	}

}
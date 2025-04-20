package main

import "fmt"

func main() {
	var n1, n2, n3, n4 float64
	fmt.Scan(&n1, &n2, &n3, &n4)

	media := calculo_media_3(n1, n2, n3, n4)
	verificacao_aprovacao(media)
}

func calculo_media_3(n1, n2, n3, n4 float64) float64 {
	var peso1 int = 2
	var peso2 int = 3
	var peso3 int = 4
	var peso4 int = 1

	media := (float64(peso1) * n1 + float64(peso2) * n2 + float64(peso3) * n3 + float64(peso4) * n4) / (float64(peso1) + float64(peso2) + float64(peso3) + float64(peso4))

	return media
}

func verificacao_aprovacao(media float64) {
	fmt.Printf("Media: %.1f\n", media)

	if (media >= 7.0) {
		fmt.Println("Aluno aprovado.")
	} else if (media < 5.0) {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")

		var notaExame float64
		fmt.Scan(&notaExame)
		fmt.Printf("Nota do exame: %.1f\n", notaExame)

		mediaFinal := pontuacao_exame(media, notaExame)

		if (mediaFinal >= 5.0) {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", mediaFinal)
	}
}

func pontuacao_exame(media, notaExame float64) float64 {
	return (media + notaExame) / 2
}
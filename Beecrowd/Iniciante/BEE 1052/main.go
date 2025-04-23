package main

import "fmt"

func main() {
	var valor int
	fmt.Scan(&valor)

	calcular_mes(valor)
}

func calcular_mes(mes int) {
	meses := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	if (mes >= 1 && mes <= 12) {
		fmt.Println(meses[mes-1])
	} else {
		fmt.Println("Invalid month")
	}
}
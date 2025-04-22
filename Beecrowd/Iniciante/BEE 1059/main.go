package main

import "fmt"

func main() {
	numeros_pares()
}

func numeros_pares() {
	for	numero := 1; numero <= 100; numero++ {
		if numero %2 == 0 {
			fmt.Println(numero)
		}
	}
}
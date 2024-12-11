package main

import "fmt"

func generica(i interface{}) { // Interface genérica aceitam qualquer tipo de dado
	fmt.Println(i)
}

func main() {
	generica(1)
	generica("string")
	generica(true)
}
package main

import "fmt"

func main() {
	canal := make(chan string, 2) // capacidade 2 com buffer
	canal <- "Olá"
	canal <- "Programando em Go"

	mensagem := <-canal
	fmt.Println(mensagem)
}
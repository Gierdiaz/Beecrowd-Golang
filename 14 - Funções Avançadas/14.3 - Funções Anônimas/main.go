package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função anonima")
	}()

	func(text string) {
		fmt.Println(text)
	}("Função anonima com parâmetro")

	func(text string) string {
		return text
	}("Função anonima com retorno")
}
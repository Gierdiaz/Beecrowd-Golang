package main

import "fmt"

func main() {
	fmt.Println("FUnção main sendo executada")
}

func init() { // Só pode ter uma função init por arquivo
	fmt.Println("Executando a função Init")
}
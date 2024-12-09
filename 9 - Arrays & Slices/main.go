package main

import "fmt"

func main() {
	// Arrays são inflexiveis
	fmt.Println("----------------------")

	var arr [5]int
	fmt.Println(arr)

	arr2 := [3]string{"Allison", "Maria", "Joao"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr3)

	// Slices são flexiveis
	fmt.Println("----------------------")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice2 := []string{"Allison", "Maria", "Joao"}
	slice2 = append(slice2, "Pedro")
	slice2 = slice2[1:3] // corta do primeiro e vai até o terceiro
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("----------------------")

	slice3 := make([]float32, 10, 15) // tamanho 10, capacidade 15. Se eu alterar de 10 para 16 ele ultrapassa nosso limite que é um array
	fmt.Println(len(slice3), cap(slice3))
	fmt.Println(slice3)

}
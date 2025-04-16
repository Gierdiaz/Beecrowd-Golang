package main

import "fmt"

func main() {
	// variável explícita
	var name string = "Állison"
	fmt.Println(name)


	// o tipo de dado int representa um inteiro 
	var numeroInt int = 10000 // Tamanho da arquitetura do computador
	var numeroInt16 int16 = 10000
	var numeroInt32 int32 = 10000
	var numeroInt64 int64 = 10000

	fmt.Printf("Tipo: %T, Valor: %v\n", numeroInt, numeroInt)
	fmt.Printf("Tipo: %T, Valor: %v\n", numeroInt16, numeroInt16)
	fmt.Printf("Tipo: %T, Valor: %v\n", numeroInt32, numeroInt32)
	fmt.Printf("Tipo: %T, Valor: %v\n", numeroInt64, numeroInt64)

	// alias

	// rune = int32
	var numeroRune rune = 12344
	// byte = uint8
	var numeroByte byte = 123

	fmt.Printf("Tipo: %T, Valor: %v\n", numeroRune, numeroRune)
	fmt.Printf("Tipo: %T, Valor: %v\n", numeroByte, numeroByte)

	var numeroFloat32 float32 = 123.45
	var numeroFloat64 float64 = 123.45
	fmt.Printf("Tipo: %T, Valor: %v\n", numeroFloat32, numeroFloat32)
	fmt.Printf("Tipo: %T, Valor: %v\n", numeroFloat64, numeroFloat64)

	var boolean bool = true
	fmt.Printf("Tipo: %T, Valor: %v\n", boolean, boolean)

	var err error
	fmt.Printf("Tipo: %T, Valor: %v\n", err, err)
}


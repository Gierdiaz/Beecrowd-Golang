package main

import "fmt"

func main(){
	var a, b, x int
	fmt.Scan(&a, &b)
	x = extremamente_facil(a, b)
	fmt.Println("X =", x)
}

func extremamente_facil(a, b int) int {
	return a + b
}
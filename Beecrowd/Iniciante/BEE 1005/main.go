package main

import "fmt"

func main() {
	var a, b float64

	fmt.Scan(&a, &b)
	media := media_1(a, b)
	fmt.Printf("MEDIA = %.5f\n", media)
}

func media_1(a, b float64) float64 {
	pesoA := 3.5
	pesoB := 7.5
	return ((a * pesoA) + (b * pesoB)) / (pesoA + pesoB)
}
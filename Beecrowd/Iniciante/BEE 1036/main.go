package main

import (
	"fmt"
	"math"
)

// Δ = b² - 4ac 
// x = -b +/- sqrt(Δ) / 2a
 
func main() {
    var a, b, c float64

	fmt.Scan(&a, &b, &c)

	if a == 0 {
		fmt.Println("Impossivel calcular")
		return
	}

	delta := calculo_discriminante(a, b, c)
	raizes := calculo_raizes(a, b, delta)

	if len(raizes) == 2 {
		fmt.Printf("R1 = %.5f\n", raizes[0])
		fmt.Printf("R2 = %.5f\n", raizes[1])
	} else {
		fmt.Println("Impossivel calcular")
	}
}

func calculo_discriminante(a, b, c float64) float64 {
	delta := math.Pow(float64(b), 2.0) - 4 * float64(a) * float64(c)
	return delta
}

func calculo_raizes(a, b, delta float64) []float64{
	if delta < 0 {
		return []float64{}
	}

	x1 := (- b + math.Sqrt(delta)) / (2 * a) 
	x2 := (- b - math.Sqrt(delta)) / (2 * a)

	return []float64{x1, x2}
}
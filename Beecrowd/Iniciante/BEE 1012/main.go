package main
 
import (
    "fmt"
    "math"
)
 
func main() {
    var a, b, c float64
    
    fmt.Scan(&a, &b, &c)
    triangulo := calculo_area_triangulo(a, c)
    circulo := calculo_area_circulo(c)
    trapezio := calculo_area_trapezio(a, b, c)
    quadrado := calculo_area_quadrado(b)
    retangulo := calculo_area_retangulo(a, b)
    
    fmt.Printf("TRIANGULO: %.3f\n", triangulo)
    fmt.Printf("CIRCULO: %.3f\n", circulo)
    fmt.Printf("TRAPEZIO: %.3f\n", trapezio)
    fmt.Printf("QUADRADO: %.3f\n", quadrado)
    fmt.Printf("RETANGULO: %.3f\n", retangulo)
}

func calculo_area_triangulo(a, c float64) float64 {
    return (a * c) / 2
}

func calculo_area_circulo(c float64) float64 {
    var pi float64 = 3.14159
	return pi * math.Pow(c, 2) 
}

func calculo_area_trapezio(a, b, c float64) float64 {
        return ((a + b) * c) / 2
}

func calculo_area_quadrado(b float64) float64 {
    return math.Pow(b, 2.0)
}

func calculo_area_retangulo(a, b float64) float64 {
    return a * b
}


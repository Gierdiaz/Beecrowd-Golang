package main
 
import (
    "fmt"
    "math"
)
 
func main() {
    var raio float64
    
    fmt.Scan(&raio)
    esfera := calculo_esfera(raio)
    fmt.Printf("VOLUME = %.3f\n", esfera)
}

func calculo_esfera(raio float64) float64 {
    var pi float64 = 3.14159
    var cubo float64 = math.Pow(raio, 3.0)
    
    return ((4.0 / 3) * pi * cubo) 
}
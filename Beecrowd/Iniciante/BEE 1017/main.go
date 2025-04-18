package main
 
import (
    "fmt"
)
 
func main() {
    var tempo, velocidade int
    fmt.Scan(&tempo, &velocidade)
    
    gasto := calculo_gasto_combustivel(tempo, velocidade)
    fmt.Printf("%.3f\n", gasto)
}

func calculo_gasto_combustivel(tempo, velocidade int) float64 {
    distancia := float64(velocidade) * float64(tempo)
    return distancia / 12
}
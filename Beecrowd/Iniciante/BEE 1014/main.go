package main
 
import (
    "fmt"
)
 
func main() {
    var distancia, litro float64
    
    fmt.Scan(&distancia, &litro)
    
    consumo_medio := calculo_consumo_medio(distancia, litro)
    
    fmt.Printf("%.3f km/l\n", consumo_medio)
}

func calculo_consumo_medio(distancia, litro float64) float64 {
    return distancia / litro
}

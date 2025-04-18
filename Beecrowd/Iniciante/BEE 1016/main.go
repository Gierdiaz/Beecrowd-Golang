package main
 
import (
    "fmt"
)
 
func main() {
    var d int
    
    fmt.Scan(&d)
    distancia := calculo_distancia(d)
    fmt.Println(distancia, "minutos")
}


func calculo_distancia(distancia int) int {
	return distancia * 2
}
package main
 
import (
    "fmt"
)
 
func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    media := media_2(a, b, c)
    fmt.Printf("MEDIA = %.1f\n", media)
    
}

func media_2(a, b, c float64) float64 {
    var pesoA float64 = 2
    var pesoB float64 = 3
    var pesoC float64 = 5
    
    return ((pesoA*a) + (pesoB*b) + (pesoC*c)) / (pesoA + pesoB + pesoC)
}
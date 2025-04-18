package main
 
import (
    "fmt"
    "math"
)

type Ponto struct {
	X float64
	Y float64 
}
 
func main() {
    var x1, y1, x2, y2 float64

	fmt.Scan(&x1, &y1, &x2, &y2)

	p1 := Ponto{X: x1, Y: y1}
	p2 := Ponto{X: x2, Y: y2}

	distancia := distancia_entre_dois_pontos(p1, p2)
	fmt.Printf("%.4f\n", distancia)
}

func distancia_entre_dois_pontos(p1, p2 Ponto) float64 {
	dx := p2.X - p1.X
    dy := p2.Y - p1.Y
	
	return math.Sqrt(dx*dx + dy*dy)
}
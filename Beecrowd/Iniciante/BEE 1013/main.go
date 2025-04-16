package main
 
import (
    "fmt"
)
 
func main() {
    var a, b, c int
    
    fmt.Scan(&a, &b, &c)
    
    maiorAB := maior_valor(a, b)
	maiorValor := maior_valor(maiorAB, c)
    
    fmt.Println(maiorValor ,"eh o maior")
}
// Usa a fórmula do enunciado: (a + b + abs(a - b)) / 2
func maior_valor(a, b int) int {
	return (a + b + abs(a - b)) / 2
}

// Retorna o valor absoluto
func abs(s int) int {
	if s < 0 {
		return -s
	}
	return s
}

package main
 
import (
    "fmt"
)
 
func main() {
    var funcionarios, hora int64
	var  valor float64
    fmt.Scan(&funcionarios, &hora, &valor)
    
    sal := salario(hora, valor)
    
    fmt.Println("NUMBER =", funcionarios)
	fmt.Printf("SALARY = U$ %.2f\n", sal)
}

func salario(hora int64, valor float64) float64 {
	return float64(hora) * valor
}
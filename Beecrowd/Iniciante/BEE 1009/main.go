package main
 
import (
    "fmt"
)
 
func main() {
    var vendedor string
    var salario, venda float64
    
    fmt.Scan(&vendedor, &salario, &venda)

    total := salario_com_bonus(salario, venda)
    fmt.Printf("TOTAL = R$ %.2f\n", total)
}

func salario_com_bonus(salario, venda float64) float64 {
    return salario + (venda * 0.15)
}
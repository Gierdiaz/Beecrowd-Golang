package main
 
import (
    "fmt"
)
 
func main() {
    var a, b, c, d int64
    fmt.Scan(&a, &b, &c, &d)
    diff := diferenca(a, b, c, d)
    fmt.Println("DIFERENCA =", diff)
}

func diferenca(a, b, c, d int64 ) int64  {
    return (a*b - c*d)
}


package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	par_impar(n)
}

func par_impar(n int) {
	for i := 0; i < n; i++ { // para ler N vezes
		var x int
		fmt.Scan(&x)

		if x == 0 {
			fmt.Println("NULL")
		} else if x%2 == 0 && x > 0 {
			fmt.Println("EVEN POSITIVE")
		} else if x%2 == 0 && x < 0 {
			fmt.Println("EVEN NEGATIVE")
		} else if x%2 != 0 && x > 0 {
			fmt.Println("ODD POSITIVE")
		} else if x%2 != 0 && x < 0 {
			fmt.Println("ODD NEGATIVE")
		}
	}
}
package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	
	in, out := contarIntervalo(n)

	fmt.Printf("%d in\n", in)
	fmt.Printf("%d out\n", out)
}

func contarIntervalo(n int) (int, int) {

	var x, in, out int

	for iteracao := 0; iteracao < n; iteracao++ {
		fmt.Scan(&x)
		if (x >= 10 && x <= 20) {
			in++
		} else {
			out++
		}
	}

	return in, out
}
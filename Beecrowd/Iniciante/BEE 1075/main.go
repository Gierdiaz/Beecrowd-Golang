package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	resto_2(n)
}

func resto_2(n int) {
	for i := 1; i < 10000; i++ {
		if i % n == 2 {
			fmt.Println(i)
		}
	}
}
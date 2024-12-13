package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	go escrever("Primeiro")
	escrever("Programanndo em Go!")

}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
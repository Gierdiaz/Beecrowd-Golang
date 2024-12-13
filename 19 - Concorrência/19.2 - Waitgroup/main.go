package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitgroup sync.WaitGroup

	waitgroup.Add(2) // passanndo a quantidade de goroutines que ele deve esperar

	go func() {
		escrever("Primeiro")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitgroup.Done()
	}()

	waitgroup.Wait() // espera a contnagem chegar em zero
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
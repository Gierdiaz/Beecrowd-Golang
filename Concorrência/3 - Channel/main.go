package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go escrever("Primeira Goroutine", channel)

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Fim da execução.")
}

func escrever(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // canal está recendo o texto. Estou enviando um valor para o canal
		time.Sleep(time.Second)
	}

	close(channel)
}
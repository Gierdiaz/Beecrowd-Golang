package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i < 5 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep((time.Second * 1))
	}

	for j := 0; j < 5; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Allison", "Maria", "Joao"}  
	for indice, name := range names {
		fmt.Println(indice, name)
	}

	for _, letra := range "GOLANG" {
		fmt.Println(string(letra))
	}

	user := map[string]string{
		"name": "Allison",
		"email": "gierdiaz@hotmail",
	}

	for _, value := range user {
		fmt.Println(value)
	}
}
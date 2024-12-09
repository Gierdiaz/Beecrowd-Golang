package main

import "fmt"

type Student struct{
	Name 			string
	Registration 	rune
	Specialty
}

type Specialty struct {
	Course 			string
	Specialization 	string
}

func main() {
	s := Student{
		Name: "Allison",
		Registration: 2021085758,
		Specialty: Specialty{
			Course: "Análise e Desenvolvimento de Sistemas",
			Specialization: "Sistemas Distrubuidos e Arquitetura de Microserviços",
		},
	}

	fmt.Println(s)
}
package main

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float32
}

type Retangulo struct {
	altura float32
	largura float32
}

type Circulo struct {
	raio float32
}

func (c Circulo) aera() float32 {
	return 2 * c.raio * math.Pi
}

func (r Retangulo) area() float32 {
	return r.altura * r.largura
}


func escreverArea(r Forma) {
	fmt.Printf("O cálculo da área é: %.2f\n", r.area())
}

func escreverRaio(c Circulo) {
	fmt.Printf("O raio do círculo é: %.2f\n", c.aera())
}

func main() {
	ret := Retangulo{
		altura: 10.5,
		largura: 11.6,
	}

	cir := Circulo{
		raio: 10,
	}

	escreverArea(ret)
	escreverRaio(cir)
	
}
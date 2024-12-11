package main

import "fmt"

type Pessoa struct {
	Name string
	Age int
}

func (p Pessoa) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no Banco de Dados.\n", p.Name)
}

func (p Pessoa) maiorIdade() bool {
	return p.Age >= 18
}

func (p *Pessoa) atualizaNome(name string) {
	p.Name = name
	fmt.Printf("Nome atualizado para: %s\n", p.Name)
}

func main() {
	pessoa := Pessoa{"Allison", 21}
	pessoa.salvar()

	pessoa2 := Pessoa{"Maria", 22}
	pessoa2.salvar()
	maiorDeIdade := pessoa2.maiorIdade()
	fmt.Println(maiorDeIdade)

	pessoa2.atualizaNome("Pâmela")

}

package endereco

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinusculo := strings.ToLower(endereco)

	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			return cases.Title(language.BrazilianPortuguese).String(tipo)
		}
	}


	return "Tipo de endereço inválido"
}

package endereco

import "testing"

type enderecoTest struct {
	endereco string
	tipoEsperado string
}

func TestTipoEndereco(t *testing.T) {
	testes := []enderecoTest{
		{"Avenida Paulista", "Avenida"},
		{"Rua das Laranjeiras", "Rua"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praceta da Liberdade", "Praceta"},
		{"", "Tipo de endereço inválido"},
	}

	for _, teste := range testes {
		tipoRecebido := TipoEndereco(teste.endereco)
		if tipoRecebido != teste.tipoEsperado {
			t.Errorf("Tipo recebido: %s, Esperado: %s", tipoRecebido, teste.tipoEsperado)
		}
	}
}
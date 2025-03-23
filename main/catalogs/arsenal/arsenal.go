package arsenal

import (
	"aprendendogolang/catalogs/dados"
)

type Arma struct {
	Nome string
	Dano int
}

var (
	Arco = Arma{
		Nome: "Arco Curto",
		Dano: dados.RolarDados(1, 4, 0),
	}
	MachadoSimples = Arma{
		Nome: "Machado Simples",
		Dano: dados.RolarDados(1, 6, 0),
	}
	Adaga = Arma{
		Nome: "Adaga",
		Dano: dados.RolarDados(1, 4, 0),
	}
)

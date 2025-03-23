package arsenal

import (
	"aprendendogolang/catalogs/dados"
)

type Arma struct {
	Nome string
	Dano func() int
}

var (
	Arco = Arma{
		Nome: "Arco Curto",
		Dano: func() int { return dados.RolarDados(1, 4) },
	}
	MachadoSimples = Arma{
		Nome: "Machado Simples",
		Dano: func() int { return dados.RolarDados(1, 6) },
	}
	Adaga = Arma{
		Nome: "Adaga",
		Dano: func() int { return dados.RolarDados(1, 4) },
	}
	Soco = Arma{
		Nome: "Soco",
		Dano: func() int { return dados.RolarDados(1, 4) },
	}
	Mordida = Arma{
		Nome: "Mordida",
		Dano: func() int { return dados.RolarDados(1, 4) },
	}
)

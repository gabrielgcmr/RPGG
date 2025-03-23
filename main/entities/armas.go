package entities

import (
	"aprendendogolang/catalogs/dados"
	"aprendendogolang/types"
)

var (
	Arco = types.Arma{
		Nome: "Arco Curto",
		Dano: dados.D6,
	}
	MachadoSimples = types.Arma{
		Nome: "Machado Simples",
		Dano: dados.D8,
	}
	Adaga = types.Arma{
		Nome: "Adaga",
		Dano: dados.D4}
)

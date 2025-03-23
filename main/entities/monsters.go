package entities

import (
	"aprendendogolang/catalogs/dados"
	"aprendendogolang/factories"
)

var (
	RatoGigante   = factories.NovoMonstro("Rato Gigante", nil, 6, 4, 6, 3, 3, 2, dados.D4())
	AranhaGigante = factories.NovoMonstro("Aranha Gigante", nil, 6, 6, 6, 2, 5, 3, dados.D4())
	Lobo          = factories.NovoMonstro("Lobo", nil, 8, 13, 8, 5, 5, 5, dados.D8())
	Ogro          = factories.NovoMonstro("Ogro", &MachadoSimples, 12, 8, 12, 5, 6, 7, dados.D12())
)

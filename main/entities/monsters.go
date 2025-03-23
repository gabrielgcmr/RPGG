package entities

import (
	"aprendendogolang/catalogs/arsenal"
	"aprendendogolang/catalogs/dados"
	"aprendendogolang/factories"
)

var (
	RatoGigante   = factories.NovoMonstro("Rato Gigante", nil, 6, 4, 6, 3, 3, 2, dados.RolarDados(1, 4, 0))
	AranhaGigante = factories.NovoMonstro("Aranha Gigante", nil, 6, 6, 6, 2, 5, 3, dados.RolarDados(1, 4, 2))
	Lobo          = factories.NovoMonstro("Lobo", nil, 8, 13, 8, 5, 5, 5, dados.RolarDados(1, 8, 2))
	Ogro          = factories.NovoMonstro("Ogro", &arsenal.MachadoSimples, 12, 8, 12, 5, 6, 7, dados.RolarDados(1, 12, 2))
)

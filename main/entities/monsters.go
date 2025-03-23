package entities

import (
	"aprendendogolang/catalogs/arsenal"
	"aprendendogolang/catalogs/dados"
	"aprendendogolang/factories"
)

var (
	RatoGigante   = factories.NovoMonstro("Rato Gigante", &arsenal.Mordida, 6, 4, 8, 3, 3, 2, dados.RolarDados(1, 4)+1)
	AranhaGigante = factories.NovoMonstro("Aranha Gigante", &arsenal.Mordida, 6, 6, 8, 3, 5, 3, dados.RolarDados(1, 4)+2)
	Lobo          = factories.NovoMonstro("Lobo", &arsenal.Mordida, 8, 13, 9, 9, 5, 5, dados.RolarDados(1, 8)+2)
	Ogro          = factories.NovoMonstro("Ogro", &arsenal.MachadoSimples, 12, 8, 12, 5, 6, 7, dados.RolarDados(1, 12)+2)
)

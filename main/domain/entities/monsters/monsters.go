package monstros

import (
	"main/catalogo/armas"
	"main/domain/monstro"
)

var (
	RatoGigante = monstro.NovoMonstro(
		"Rato Gigante",
		&armas.Mordida,
		monstro.Atributos{FOR: 6, DES: 7, CON: 8, INT: 3, SAB: 3, CAR: 2},
		6,
		40,
	)

	AranhaGigante = monstro.NovoMonstro(
		"Aranha Gigante",
		&armas.Mordida,
		monstro.Atributos{FOR: 6, DES: 8, CON: 8, INT: 3, SAB: 5, CAR: 3},
		8,
		50,
	)

	Lobo = monstro.NovoMonstro(
		"Lobo",
		&armas.Mordida,
		monstro.Atributos{FOR: 8, DES: 13, CON: 9, INT: 9, SAB: 5, CAR: 5},
		10,
		80,
	)

	Ogro = monstro.NovoMonstro(
		"Ogro",
		&armas.MachadoSimples,
		monstro.Atributos{FOR: 12, DES: 8, CON: 12, INT: 5, SAB: 6, CAR: 7},
		20,
		120,
	)
)

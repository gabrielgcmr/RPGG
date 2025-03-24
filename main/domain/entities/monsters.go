package entities

import (
	"main/catalogo/armas"
	"main/domain/monstro"
	"main/game/dado"
)

var (
	RatoGigante = monstro.NovoMonstro(
		"Rato Gigante",
		armas.Mordida,
		monstro.Atributos{FOR: 6, DES: 4, CON: 8, INT: 3, SAB: 3, CAR: 2},
		func() int { return dado.Rolar(1, 4) + 1 },
		40,
	)

	AranhaGigante = monstro.NovoMonstro(
		"Aranha Gigante",
		&armas.Mordida,
		monstro.Atributos{FOR: 6, DES: 6, CON: 8, INT: 3, SAB: 5, CAR: 3},
		func() int { return dado.Rolar(1, 4) + 2 },
		50,
	)

	Lobo = monstro.NovoMonstro(
		"Lobo",
		&armas.Mordida,
		monstro.Atributos{FOR: 8, DES: 13, CON: 9, INT: 9, SAB: 5, CAR: 5},
		func() int { return dado.Rolar(1, 8) + 2 },
		80,
	)

	Ogro = monstro.NovoMonstro(
		"Ogro",
		&armas.MachadoSimples,
		monstro.Atributos{FOR: 12, DES: 8, CON: 12, INT: 5, SAB: 6, CAR: 7},
		func() int { return dado.Rolar(1, 12) + 2 },
		120,
	)
)

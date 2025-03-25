// domain/monstro/type.go
package monstro

import (
	"main/catalogo/armas"
)

type Monstro struct {
	Nome   string
	Equipe string

	Atributos struct {
		FOR int
		DES int
		CON int
		INT int
		SAB int
		CAR int
	}

	CA   int
	HP   int
	XP   int
	Arma *armas.Arma
}

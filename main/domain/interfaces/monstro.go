package interfaces

import (
	"main/catalogs"
	"main/catalogs/arsenal"
)

type Monstro struct {
	Nome      string
	Entidade  string
	Tipo      string
	Atributos catalogs.Atributos

	FOR int
	DES int
	CON int
	INT int
	SAB int
	CAR int

	ClasseDeArmadura int
	PontosDeVida     int
	Arma             *arsenal.Arma
}
